package scanner

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	alog "github.com/onedusk/swiper/internal/log"
	"github.com/onedusk/swiper/internal/metrics"
	"github.com/onedusk/swiper/internal/pool"
)

// PDFScanner handles scanning and copying PDF files
type PDFScanner struct {
	scanDir          string
	copyDir          string
	logger           *alog.AsyncLogger
	bufferManager    *pool.BufferPoolManager
	metricsCollector *metrics.Collector
}

// New creates a new PDF scanner instance
func New(scanDir, copyDir string) (*PDFScanner, error) {
	// Use current directory if not specified
	if scanDir == "" {
		scanDir = "."
	}

	// Default copy directory
	if copyDir == "" {
		copyDir = "pdf-docs"
	}

	// Create the copy directory if it doesn't exist
	if err := os.MkdirAll(copyDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create copy directory: %v", err)
	}

	metricsCollector := metrics.NewCollector()
	scanner := &PDFScanner{
		scanDir:          scanDir,
		copyDir:          copyDir,
		logger:           alog.New(100, false),
		bufferManager:    pool.NewBufferPoolManager(metricsCollector),
		metricsCollector: metricsCollector,
	}

	return scanner, nil
}

// logAsync sends a log message asynchronously
func (s *PDFScanner) logAsync(format string, v ...interface{}) {
	s.logger.Log(format, v...)
}

// FindPDFs recursively finds all PDF files in the scan directory
func (s *PDFScanner) FindPDFs() ([]string, error) {
	var pdfFiles []string

	err := filepath.Walk(s.scanDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			s.logAsync("Error accessing path %s: %v", path, err)
			return nil // Continue walking
		}

		// Skip the copy directory to avoid recursion
		if info.IsDir() && path == s.copyDir {
			return filepath.SkipDir
		}

		// Check if it's a PDF file
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

// copyPDFWithProgress copies a PDF file with progress indication
func (s *PDFScanner) copyPDFWithProgress(src string, totalFiles, currentFile int) error {
	// Get just the filename
	baseName := filepath.Base(src)
	dst := filepath.Join(s.copyDir, baseName)

	// Check if file already exists
	if _, err := os.Stat(dst); err == nil {
		// File exists, add timestamp to make unique
		timestamp := time.Now().Format("20060102_150405")
		ext := filepath.Ext(baseName)
		nameWithoutExt := strings.TrimSuffix(baseName, ext)
		baseName = fmt.Sprintf("%s_%s%s", nameWithoutExt, timestamp, ext)
		dst = filepath.Join(s.copyDir, baseName)
	}

	s.logAsync("[%d/%d] Copying: %s -> %s", currentFile, totalFiles, src, dst)

	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()

	// Get file info for size
	fileInfo, err := srcFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat source file: %v", err)
	}

	// Create destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dstFile.Close()

	// Get optimized buffer from pool manager
	fileSize := fileInfo.Size()
	buffer := s.bufferManager.GetBuffer(fileSize)
	defer s.bufferManager.PutBuffer(buffer)

	copied, err := io.CopyBuffer(dstFile, srcFile, buffer)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}

	if copied != fileSize {
		return fmt.Errorf("copy size mismatch: expected %d, got %d", fileSize, copied)
	}

	// Record metrics
	s.metricsCollector.RecordBytesProcessed(copied)

	// Sync to ensure data is written
	if err := dstFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync file: %v", err)
	}

	s.logAsync("[%d/%d] Successfully copied: %s (%.2f MB)", currentFile, totalFiles, baseName, float64(copied)/(1024*1024))
	return nil
}

// ScanAndCopy scans for PDFs and copies them with concurrent processing
func (s *PDFScanner) ScanAndCopy() error {
	s.logAsync("Scanning directory: %s", s.scanDir)

	// Find all PDF files
	pdfFiles, err := s.FindPDFs()
	if err != nil {
		return err
	}

	if len(pdfFiles) == 0 {
		s.logAsync("No PDF files found in %s", s.scanDir)
		return nil
	}

	s.logAsync("Found %d PDF files", len(pdfFiles))
	startTime := time.Now()

	// Determine optimal concurrency
	maxWorkers := 4
	if len(pdfFiles) < maxWorkers {
		maxWorkers = len(pdfFiles)
	}

	// Create work channel and semaphore
	type workItem struct {
		path  string
		index int
	}
	workChan := make(chan workItem, len(pdfFiles))
	sem := make(chan struct{}, maxWorkers)

	// Queue all work items
	for i, pdfPath := range pdfFiles {
		workChan <- workItem{path: pdfPath, index: i + 1}
	}
	close(workChan)

	// Process files concurrently
	var wg sync.WaitGroup
	var successCount int64
	var totalSize int64
	var failedFiles []string
	var mu sync.Mutex

	for work := range workChan {
		wg.Add(1)
		sem <- struct{}{} // Acquire semaphore

		go func(item workItem) {
			defer wg.Done()
			defer func() { <-sem }() // Release semaphore

			if err := s.copyPDFWithProgress(item.path, len(pdfFiles), item.index); err != nil {
				s.logAsync("Failed to copy %s: %v", item.path, err)
				mu.Lock()
				failedFiles = append(failedFiles, filepath.Base(item.path))
				mu.Unlock()
				return
			}

			atomic.AddInt64(&successCount, 1)

			// Get file size for stats
			if info, err := os.Stat(item.path); err == nil {
				atomic.AddInt64(&totalSize, info.Size())
			}
		}(work)
	}

	wg.Wait()

	duration := time.Since(startTime).Seconds()
	successful := atomic.LoadInt64(&successCount)
	total := atomic.LoadInt64(&totalSize)

	s.logAsync("\nScan and copy completed in %.2f seconds", duration)
	s.logAsync("Successfully copied %d out of %d PDF files", successful, len(pdfFiles))
	s.logAsync("Total size copied: %.2f MB", float64(total)/(1024*1024))
	s.logAsync("Average copy speed: %.2f MB/s", float64(total)/(1024*1024*duration))

	if len(failedFiles) > 0 {
		s.logAsync("\nFailed to copy:")
		for _, file := range failedFiles {
			s.logAsync("  - %s", file)
		}
	}

	s.logAsync("PDF files copied to: %s", s.copyDir)

	// Print metrics if available
	if s.metricsCollector != nil {
		s.metricsCollector.PrintSummary("scan")
	}

	// Close log channel
	s.logger.Close()

	return nil
}