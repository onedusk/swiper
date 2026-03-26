package batch

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/onedusk/swiper/internal/extractor"
	alog "github.com/onedusk/swiper/internal/log"
	"github.com/onedusk/swiper/internal/metrics"
	"github.com/onedusk/swiper/internal/pool"
)

// Processor handles batch processing of multiple PDF files
type Processor struct {
	inputDir         string
	outputDir        string
	processCount     int
	logger           *alog.AsyncLogger
	concurrentPDFs   int // Number of PDFs to process simultaneously
	workerPool       chan struct{}
	metricsCollector *metrics.Collector
	bufferManager    *pool.BufferPoolManager
	resume           bool
}

// progressState tracks batch processing progress for resume support.
type progressState struct {
	Completed []string `json:"completed"`
	Failed    []string `json:"failed"`
	Timestamp string   `json:"timestamp"`
}

const progressFileName = ".swiper-progress"

// New creates a new batch processor
func New(inputDir, outputDir string, processCount int) (*Processor, error) {
	if inputDir == "" {
		return nil, fmt.Errorf("input directory is required")
	}

	// Check if input directory exists
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("input directory not found: %s", inputDir)
	}

	// Default output directory
	if outputDir == "" {
		outputDir = "extracted-pdfs"
	}

	// Create main output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %v", err)
	}

	if processCount <= 0 {
		if n := runtime.NumCPU(); n > 0 {
			processCount = n
		} else {
			processCount = 4
		}
	}

	// Calculate optimal concurrent PDFs based on available cores
	concurrentPDFs := calculateOptimalConcurrentPDFs(processCount)

	// Adaptive channel sizing
	logChannelSize := 100
	if concurrentPDFs > 4 {
		logChannelSize = concurrentPDFs * 50
	}

	logger := alog.New(logChannelSize, false)
	workerPool := make(chan struct{}, concurrentPDFs)

	metricsCollector := metrics.NewCollector()
	batch := &Processor{
		inputDir:         inputDir,
		outputDir:        outputDir,
		processCount:     processCount,
		logger:           logger,
		concurrentPDFs:   concurrentPDFs,
		workerPool:       workerPool,
		metricsCollector: metricsCollector,
		bufferManager:    pool.NewBufferPoolManager(metricsCollector),
	}

	return batch, nil
}

// logAsync sends a log message asynchronously
func (b *Processor) logAsync(format string, v ...interface{}) {
	b.logger.Log(format, v...)
}

// FindPDFs finds all PDF files in the input directory
func (b *Processor) FindPDFs() ([]string, error) {
	var pdfFiles []string

	err := filepath.Walk(b.inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			b.logAsync("Error accessing path %s: %v", path, err)
			return nil // Continue walking
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".pdf") {
			pdfFiles = append(pdfFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %v", err)
	}

	return pdfFiles, nil
}

// processPDF processes a single PDF file
func (b *Processor) processPDF(pdfPath string, index, total int) error {
	// Get the PDF filename without extension for output directory
	baseName := filepath.Base(pdfPath)
	nameWithoutExt := strings.TrimSuffix(baseName, filepath.Ext(baseName))
	pdfOutputDir := filepath.Join(b.outputDir, nameWithoutExt)

	b.logAsync("[%d/%d] Processing: %s", index, total, baseName)
	startTime := time.Now()

	// Get file size for metrics
	if info, err := os.Stat(pdfPath); err == nil {
		b.metricsCollector.RecordPDFSize(baseName, info.Size())
	}

	// Create extractor for this PDF with adjusted process count
	processesPerPDF := b.processCount / b.concurrentPDFs
	if processesPerPDF < 1 {
		processesPerPDF = 1
	}

	ext, err := extractor.New(pdfPath, pdfOutputDir, processesPerPDF,
		extractor.WithMetrics(b.metricsCollector))
	if err != nil {
		return fmt.Errorf("failed to initialize extractor: %v", err)
	}
	defer ext.Cleanup()

	// Extract pages
	if err := ext.ExtractPages(); err != nil {
		return fmt.Errorf("extraction failed: %v", err)
	}

	duration := time.Since(startTime)
	b.metricsCollector.RecordProcessingTime(duration)
	b.logAsync("[%d/%d] Completed %s in %.2f seconds", index, total, baseName, duration.Seconds())

	return nil
}

// SetResume enables resume mode for batch processing.
func (b *Processor) SetResume(enabled bool) {
	b.resume = enabled
}

func (b *Processor) progressFilePath() string {
	return filepath.Join(b.outputDir, progressFileName)
}

func (b *Processor) loadProgress() *progressState {
	data, err := os.ReadFile(b.progressFilePath())
	if err != nil {
		return &progressState{}
	}
	var state progressState
	if err := json.Unmarshal(data, &state); err != nil {
		return &progressState{}
	}
	return &state
}

func (b *Processor) saveProgress(state *progressState) {
	state.Timestamp = time.Now().Format(time.RFC3339)
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return
	}
	if err := os.WriteFile(b.progressFilePath(), data, 0644); err != nil {
		b.logAsync("Warning: failed to save progress file: %v", err)
	}
}

// ProcessAll processes all PDF files in the directory
func (b *Processor) ProcessAll() error {
	b.logAsync("Scanning directory: %s", b.inputDir)

	pdfFiles, err := b.FindPDFs()
	if err != nil {
		return err
	}

	if len(pdfFiles) == 0 {
		b.logAsync("No PDF files found in %s", b.inputDir)
		return nil
	}

	// Resume support: skip already-completed PDFs
	var state *progressState
	if b.resume {
		state = b.loadProgress()
		completedSet := make(map[string]bool, len(state.Completed))
		for _, name := range state.Completed {
			completedSet[name] = true
		}
		var remaining []string
		for _, path := range pdfFiles {
			if !completedSet[filepath.Base(path)] {
				remaining = append(remaining, path)
			}
		}
		if len(remaining) < len(pdfFiles) {
			b.logAsync("Resuming: %d/%d PDFs already completed, %d remaining", len(pdfFiles)-len(remaining), len(pdfFiles), len(remaining))
		}
		pdfFiles = remaining
		state.Failed = nil // Reset failed list for this run
	} else {
		state = &progressState{}
	}

	if len(pdfFiles) == 0 {
		b.logAsync("All PDFs already processed")
		return nil
	}

	b.logAsync("Found %d PDF files to process", len(pdfFiles))
	b.logAsync("Output directory: %s", b.outputDir)
	b.logAsync("Using %d processes per PDF, %d concurrent PDFs\n", b.processCount, b.concurrentPDFs)

	overallStart := time.Now()

	var wg sync.WaitGroup
	var mu sync.Mutex
	successCount := 0
	failedPDFs := []string{}

	for i, pdfPath := range pdfFiles {
		wg.Add(1)
		b.workerPool <- struct{}{}

		go func(path string, index int) {
			defer wg.Done()
			defer func() { <-b.workerPool }()

			baseName := filepath.Base(path)
			if err := b.processPDF(path, index+1, len(pdfFiles)); err != nil {
				b.logAsync("Failed to process %s: %v", baseName, err)
				mu.Lock()
				failedPDFs = append(failedPDFs, baseName)
				if b.resume {
					state.Failed = append(state.Failed, baseName)
					b.saveProgress(state)
				}
				mu.Unlock()
			} else {
				mu.Lock()
				successCount++
				if b.resume {
					state.Completed = append(state.Completed, baseName)
					b.saveProgress(state)
				}
				mu.Unlock()
			}
		}(pdfPath, i)
	}

	wg.Wait()

	overallDuration := time.Since(overallStart).Seconds()

	b.logAsync("%s", "\n"+strings.Repeat("=", 50))
	b.logAsync("BATCH PROCESSING COMPLETE")
	b.logAsync("%s", strings.Repeat("=", 50))
	b.logAsync("Total time: %.2f seconds", overallDuration)
	b.logAsync("Successfully processed: %d/%d PDFs", successCount, len(pdfFiles))
	b.logAsync("Average time per PDF: %.2f seconds", overallDuration/float64(len(pdfFiles)))

	// Print performance metrics
	b.metricsCollector.PrintSummary("batch")

	if len(failedPDFs) > 0 {
		b.logAsync("\nFailed PDFs:")
		for _, pdf := range failedPDFs {
			b.logAsync("  - %s", pdf)
		}
	}

	b.logAsync("\nAll extracted content saved to: %s", b.outputDir)

	// Clean completion: remove progress file
	if b.resume && len(failedPDFs) == 0 {
		os.Remove(b.progressFilePath())
	}

	// Flush logger
	b.logger.Close()

	return nil
}

// calculateOptimalConcurrentPDFs determines optimal PDF concurrency
func calculateOptimalConcurrentPDFs(processCount int) int {
	concurrentPDFs := processCount / 4

	// Apply min/max bounds
	if concurrentPDFs < 2 {
		concurrentPDFs = 2
	} else if concurrentPDFs > 8 {
		concurrentPDFs = 8
	}

	return concurrentPDFs
}