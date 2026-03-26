package extractor

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/onedusk/swiper/internal/cache"
	"github.com/onedusk/swiper/internal/config"
	alog "github.com/onedusk/swiper/internal/log"
	"github.com/onedusk/swiper/internal/metrics"
	"github.com/onedusk/swiper/internal/pool"
)

// Extractor encapsulates the PDF extraction logic
type Extractor struct {
	pdfPath          string
	outputDir        string
	imageDir         string
	processCount     int
	bufferPool       *sync.Pool
	bufferManager    *pool.BufferPoolManager
	logger           *alog.AsyncLogger
	tempDirPool      *pool.TempDirPool
	pageCount        int
	pageCountMu      sync.Mutex
	resultCache      *cache.ResultCache
	metricsCollector *metrics.Collector
	ctx              context.Context
	cancel           context.CancelFunc
	commandPool      *pool.CommandPool
	cleanupOnce      sync.Once
	pageResults      []*PageResult
	pageResultsMu    sync.Mutex
	quiet            bool
	pageRanges       []config.PageRange
	perPageImageDirs bool
}

// New creates a new extractor instance
func New(pdfPath, outputDir string, processCount int, opts ...Option) (*Extractor, error) {
	// Ensure the PDF file exists
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("PDF file not found: %s", pdfPath)
	}

	// If no output directory is provided, use the PDF filename (without extension)
	if outputDir == "" {
		base := filepath.Base(pdfPath)
		ext := filepath.Ext(pdfPath)
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

	// Pre-create temp directories pool
	tempPoolSize := processCount * 2
	if processCount > 8 {
		tempPoolSize = processCount + 4
	}
	tempDirPool := pool.NewTempDirPool(tempPoolSize)

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())

	extractor := &Extractor{
		pdfPath:          pdfPath,
		outputDir:        outputDir,
		processCount:     processCount,
		imageDir:         imageDir,
		bufferPool:       bufferPool,
		bufferManager:    bufferManager,
		tempDirPool:      tempDirPool,
		pageCount:        -1, // Cache -1 means not yet fetched
		resultCache:      cache.NewResultCache(),
		metricsCollector: metricsCollector,
		ctx:              ctx,
		cancel:           cancel,
		commandPool:      pool.NewCommandPool(ctx, 30*time.Second),
	}

	// Apply options
	for _, opt := range opts {
		opt(extractor)
	}

	// Create logger after options (quiet flag may be set)
	if extractor.logger == nil {
		extractor.logger = alog.New(logChannelSize, extractor.quiet)
	}

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

// WithLogger sets a custom async logger (for sharing across components)
func WithLogger(l *alog.AsyncLogger) Option {
	return func(e *Extractor) {
		e.logger = l
	}
}

// WithQuiet enables quiet mode (suppresses non-error output)
func WithQuiet(q bool) Option {
	return func(e *Extractor) {
		e.quiet = q
	}
}

// WithPageRanges sets page ranges to extract (nil = all pages)
func WithPageRanges(ranges []config.PageRange) Option {
	return func(e *Extractor) {
		e.pageRanges = ranges
	}
}

// WithPerPageImageDirs organizes images into per-page subdirectories
func WithPerPageImageDirs(enabled bool) Option {
	return func(e *Extractor) {
		e.perPageImageDirs = enabled
	}
}

// ExtractPages processes all pages concurrently
func (e *Extractor) ExtractPages() error {
	totalPages, err := e.getPageCount()
	if err != nil {
		return err
	}

	processes := e.calculateWorkerCount(totalPages)

	// Determine which pages to extract
	var pagesToExtract []int
	if e.pageRanges != nil {
		pagesToExtract = config.ExpandPages(e.pageRanges, totalPages)
	} else {
		pagesToExtract = config.ExpandPages(nil, totalPages)
	}
	extractCount := len(pagesToExtract)

	e.logAsync("Extracting %d pages from %s using %d processes", extractCount, e.pdfPath, processes)
	startTime := time.Now()

	e.pageResults = make([]*PageResult, 0, extractCount)

	// Run worker pool
	successCount, errorCount := e.runWorkerPool(processes, extractCount, pagesToExtract, startTime)

	if err := e.createMainMarkdown(); err != nil {
		return err
	}

	e.reportExtractionSummary(startTime, successCount, errorCount, extractCount)

	return nil
}

// calculateWorkerCount determines the optimal number of workers based on
// page count, CPU count, and available memory.
func (e *Extractor) calculateWorkerCount(totalPages int) int {
	processes := e.processCount
	if totalPages < processes {
		processes = totalPages
	}

	// Adaptive scaling for very large PDFs
	if totalPages > 500 && processes > 8 {
		processes = 8 + (processes-8)/2
	}

	return processes
}

// runWorkerPool starts worker goroutines, enqueues pages, and waits for completion.
// Returns success and error counts.
func (e *Extractor) runWorkerPool(processes, extractCount int, pagesToExtract []int, startTime time.Time) (int64, int64) {
	var wg sync.WaitGroup

	pageBufferSize := calculateOptimalBufferSize(extractCount, processes)
	pages := make(chan int, pageBufferSize)

	var successCount int64
	var errorCount int64
	var completedCount int64
	var lastProgressTime int64

	for i := 0; i < processes; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			workerStart := time.Now()
			processedPages := 0

			for page := range pages {
				pageStart := time.Now()

				if processedPages%10 == 0 {
					e.metricsCollector.RecordPageQueueDepth(len(pages))
				}

				result := e.processPage(page)

				e.pageResultsMu.Lock()
				e.pageResults = append(e.pageResults, result)
				e.pageResultsMu.Unlock()

				if result.Success() {
					atomic.AddInt64(&successCount, 1)
				} else if errors.Is(result.TextErr, context.Canceled) {
					atomic.AddInt64(&errorCount, 1)
					atomic.AddInt64(&completedCount, 1)
					return
				} else {
					atomic.AddInt64(&errorCount, 1)
					e.logAsync("Worker %d: page %d: %s", workerID, page, result.ErrorSummary())
				}

				// Progress reporting
				completed := atomic.AddInt64(&completedCount, 1)
				e.reportProgress(completed, int64(extractCount), page, result.Duration, startTime, &lastProgressTime)

				e.metricsCollector.RecordWorkerTime(workerID, time.Since(pageStart))
				processedPages++
			}

			e.metricsCollector.RecordWorkerTime(workerID, time.Since(workerStart))
		}(i)
	}

	go func() {
		defer close(pages)
		for _, p := range pagesToExtract {
			select {
			case pages <- p:
			case <-e.ctx.Done():
				return
			}
		}
	}()

	wg.Wait()
	return successCount, errorCount
}

// reportProgress logs extraction progress, rate-limited for large PDFs.
func (e *Extractor) reportProgress(completed, total int64, page int, pageDur time.Duration, startTime time.Time, lastProgressTime *int64) {
	now := time.Now().UnixNano()
	lastProg := atomic.LoadInt64(lastProgressTime)
	shouldReport := false
	if total <= 10 {
		shouldReport = true
	} else if now-lastProg > 2e9 {
		shouldReport = true
	} else if completed*100/total%5 == 0 {
		shouldReport = true
	}
	if shouldReport && atomic.CompareAndSwapInt64(lastProgressTime, lastProg, now) {
		pct := completed * 100 / total
		remaining := total - completed
		avgDur := time.Since(startTime) / time.Duration(completed)
		eta := time.Duration(remaining) * avgDur
		e.logAsync("[%d/%d] %d%% - Page %d (%.1fs) ETA %s", completed, total, pct, page, pageDur.Seconds(), eta.Round(time.Second))
	}
}

// reportExtractionSummary logs the final extraction statistics.
func (e *Extractor) reportExtractionSummary(startTime time.Time, successCount, errorCount int64, extractCount int) {
	duration := time.Since(startTime)

	e.logAsync("")
	e.logAsync("Extraction completed in %.2f seconds", duration.Seconds())
	e.logAsync("Successfully extracted %d out of %d pages", successCount, extractCount)
	if errorCount > 0 {
		e.logAsync("Failed to extract %d pages", errorCount)
	}
	e.logAsync("Pages per second: %.2f", float64(successCount)/duration.Seconds())
	e.logAsync("Output directory: %s", e.outputDir)

	e.metricsCollector.PrintSummary("single")
}

// Cleanup releases resources. Safe to call multiple times.
func (e *Extractor) Cleanup() {
	e.cleanupOnce.Do(func() {
		e.cancel()
		e.tempDirPool.Cleanup()
		e.logger.Close()
	})
}

// Results returns the per-page extraction results, sorted by page number.
// Must be called after ExtractPages completes.
func (e *Extractor) Results() []*PageResult {
	e.pageResultsMu.Lock()
	defer e.pageResultsMu.Unlock()

	results := make([]*PageResult, len(e.pageResults))
	copy(results, e.pageResults)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Page < results[j].Page
	})
	return results
}

// getPageCount returns the total number of pages in the PDF
func (e *Extractor) getPageCount() (int, error) {
	e.pageCountMu.Lock()
	defer e.pageCountMu.Unlock()

	if e.pageCount > 0 {
		return e.pageCount, nil
	}

	cmd := exec.CommandContext(e.ctx, "pdfinfo", e.pdfPath)
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

// logAsync sends a log message asynchronously
func (e *Extractor) logAsync(format string, v ...interface{}) {
	e.logger.Log(format, v...)
}

// calculateOptimalBufferSize determines the best channel buffer size
func calculateOptimalBufferSize(totalPages, processes int) int {
	var bufferSize int

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