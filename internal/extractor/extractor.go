package extractor

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/onedusk/swiper/internal/cache"
	"github.com/onedusk/swiper/internal/metrics"
	"github.com/onedusk/swiper/internal/pool"
)

// Extractor encapsulates the PDF extraction logic
type Extractor struct {
	pdfFile          string
	outputDir        string
	imageDir         string
	processCount     int
	bufferPool       *sync.Pool
	bufferManager    *pool.BufferPoolManager
	logChan          chan string
	tempDirPool      *pool.TempDirPool
	pageCount        int
	pageCountMu      sync.Mutex
	resultCache      *cache.ResultCache
	metricsCollector *metrics.Collector
	ctx              context.Context
	cancel           context.CancelFunc
	commandPool      *pool.CommandPool
}

// New creates a new extractor instance
func New(pdfFile, outputDir string, processCount int, opts ...Option) (*Extractor, error) {
	// Ensure the PDF file exists
	if _, err := os.Stat(pdfFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("PDF file not found: %s", pdfFile)
	}

	// If no output directory is provided, use the PDF filename (without extension)
	if outputDir == "" {
		base := filepath.Base(pdfFile)
		ext := filepath.Ext(pdfFile)
		outputDir = strings.TrimSuffix(base, ext)
	}

	// Create the output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, err
	}

	imageDir := filepath.Join(outputDir, "images")
	if err := os.MkdirAll(imageDir, 0755); err != nil {
		return nil, err
	}

	// Use provided process count or default to the number of CPUs
	if processCount <= 0 {
		if n := runtime.NumCPU(); n > 0 {
			processCount = n
		} else {
			processCount = 4
		}
	}

	// Create buffer pool for reusing bytes.Buffer instances
	bufferPool := &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	// Create metrics collector
	metricsCollector := metrics.NewCollector()

	// Create managed buffer pool manager
	bufferManager := pool.NewBufferPoolManager(metricsCollector)

	// Calculate optimal channel buffer sizes
	logChannelSize := 200
	if processCount > 8 {
		logChannelSize = processCount * 50
	}

	// Create async logging channel
	logChan := make(chan string, logChannelSize)

	// Pre-create temp directories pool
	tempPoolSize := processCount * 2
	if processCount > 8 {
		tempPoolSize = processCount + 4
	}
	tempDirPool := pool.NewTempDirPool(tempPoolSize)

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())

	extractor := &Extractor{
		pdfFile:          pdfFile,
		outputDir:        outputDir,
		processCount:     processCount,
		imageDir:         imageDir,
		bufferPool:       bufferPool,
		bufferManager:    bufferManager,
		logChan:          logChan,
		tempDirPool:      tempDirPool,
		pageCount:        -1, // Cache -1 means not yet fetched
		resultCache:      cache.NewResultCache(),
		metricsCollector: metricsCollector,
		ctx:              ctx,
		cancel:           cancel,
		commandPool:      pool.NewCommandPool(ctx),
	}

	// Apply options
	for _, opt := range opts {
		opt(extractor)
	}

	// Start async logger
	go extractor.asyncLogger()

	return extractor, nil
}

// Option is a functional option for configuring the Extractor
type Option func(*Extractor)

// WithMetrics sets a custom metrics collector
func WithMetrics(m *metrics.Collector) Option {
	return func(e *Extractor) {
		e.metricsCollector = m
	}
}

// WithCache sets a custom result cache
func WithCache(c *cache.ResultCache) Option {
	return func(e *Extractor) {
		e.resultCache = c
	}
}

// ExtractPages processes all pages concurrently
func (e *Extractor) ExtractPages() error {
	totalPages, err := e.getPageCount()
	if err != nil {
		return err
	}

	// Calculate optimal worker count based on PDF size
	processes := e.processCount
	if totalPages < processes {
		processes = totalPages
	}

	// Adaptive worker scaling for very large PDFs
	if totalPages > 500 && processes > 8 {
		processes = 8 + (processes-8)/2
	}

	e.logAsync("Extracting %d pages from %s using %d processes", totalPages, e.pdfFile, processes)
	startTime := time.Now()

	var wg sync.WaitGroup

	// Smart channel sizing based on PDF characteristics
	pageBufferSize := calculateOptimalBufferSize(totalPages, processes)
	pages := make(chan int, pageBufferSize)

	// Track metrics
	var successCount int64
	var errorCount int64

	// Start worker goroutines
	for i := 0; i < processes; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			workerStart := time.Now()
			processedPages := 0

			for page := range pages {
				pageStart := time.Now()

				// Record queue depth periodically
				if processedPages%10 == 0 {
					e.metricsCollector.RecordPageQueueDepth(len(pages))
				}

				if err := e.processPage(page); err == nil {
					atomic.AddInt64(&successCount, 1)
				} else if err == context.Canceled {
					return // Exit early if cancelled
				} else {
					atomic.AddInt64(&errorCount, 1)
					e.logAsync("Worker %d: Error processing page %d: %v", workerID, page, err)
				}

				e.metricsCollector.RecordWorkerTime(workerID, time.Since(pageStart))
				processedPages++
			}

			e.metricsCollector.RecordWorkerTime(workerID, time.Since(workerStart))
		}(i)
	}

	// Producer: enqueue page numbers
	go func() {
		defer close(pages)
		for i := 1; i <= totalPages; i++ {
			select {
			case pages <- i:
			case <-e.ctx.Done():
				return
			}
		}
	}()

	wg.Wait()

	if err := e.createMainMarkdown(); err != nil {
		return err
	}

	duration := time.Since(startTime)
	successful := atomic.LoadInt64(&successCount)
	failed := atomic.LoadInt64(&errorCount)

	e.logAsync("")
	e.logAsync("Extraction completed in %.2f seconds", duration.Seconds())
	e.logAsync("Successfully extracted %d out of %d pages", successful, totalPages)
	if failed > 0 {
		e.logAsync("Failed to extract %d pages", failed)
	}
	e.logAsync("Pages per second: %.2f", float64(successful)/duration.Seconds())
	e.logAsync("Output directory: %s", e.outputDir)

	// Print metrics
	e.metricsCollector.PrintSummary("single")

	// Cleanup
	e.cancel()
	e.tempDirPool.Cleanup()

	// Close log channel and ensure all logs are flushed
	close(e.logChan)
	time.Sleep(100 * time.Millisecond)

	return nil
}

// Cleanup releases resources
func (e *Extractor) Cleanup() {
	e.cancel()
	e.tempDirPool.Cleanup()
	close(e.logChan)
}

// getPageCount returns the total number of pages in the PDF
func (e *Extractor) getPageCount() (int, error) {
	e.pageCountMu.Lock()
	defer e.pageCountMu.Unlock()

	if e.pageCount > 0 {
		return e.pageCount, nil
	}

	cmd := exec.Command("pdfinfo", e.pdfFile)
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("failed to run pdfinfo: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Pages:") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				if count, err := strconv.Atoi(parts[1]); err == nil {
					e.pageCount = count
					return count, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("failed to parse page count from pdfinfo")
}

// asyncLogger handles async logging
func (e *Extractor) asyncLogger() {
	for msg := range e.logChan {
		log.Print(msg)
	}
}

// logAsync sends a log message asynchronously
func (e *Extractor) logAsync(format string, v ...interface{}) {
	select {
	case e.logChan <- fmt.Sprintf(format, v...):
	default:
		log.Printf(format, v...)
	}
}

// calculateOptimalBufferSize determines the best channel buffer size
func calculateOptimalBufferSize(totalPages, processes int) int {
	bufferSize := min(totalPages, processes*4)

	if totalPages <= 10 {
		bufferSize = totalPages
	} else if totalPages <= 50 {
		bufferSize = min(totalPages, processes*2)
	} else if totalPages <= 200 {
		bufferSize = processes * 4
	} else if totalPages <= 1000 {
		bufferSize = processes * 6
	} else {
		bufferSize = processes * 8
		if bufferSize > 100 {
			bufferSize = 100
		}
	}

	return bufferSize
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}