package extractor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"io"

	"github.com/onedusk/swiper/internal/cache"
)

// processPage extracts text and images from a page and writes a Markdown file
func (e *Extractor) processPage(page int) *PageResult {
	start := time.Now()
	defer func() {
		e.metricsCollector.RecordPageProcessed()
		e.metricsCollector.RecordProcessingTime(time.Since(start))
	}()

	result := &PageResult{Page: page}

	// Check if context is cancelled
	select {
	case <-e.ctx.Done():
		result.TextErr = e.ctx.Err()
		result.Duration = time.Since(start)
		return result
	default:
	}

	// Extract text and images in parallel
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		result.Text, result.TextErr = e.extractTextFromPage(page)
	}()
	var imgErrors []error
	go func() {
		defer wg.Done()
		result.Images, imgErrors = e.extractImagesFromPage(page)
	}()
	wg.Wait()
	result.ImageErrors = imgErrors

	if result.TextErr != nil {
		e.logAsync("Error extracting text from page %d: %v", page, result.TextErr)
	}
	if len(result.ImageErrors) > 0 {
		e.logAsync("Image errors on page %d: %s", page, result.ErrorSummary())
	}

	// Write markdown (even for partial results)
	if err := e.writePageMarkdown(result); err != nil {
		e.logAsync("Failed to write markdown for page %d: %v", page, err)
		if result.TextErr == nil {
			result.TextErr = err
		}
	}

	result.Duration = time.Since(start)
	return result
}

// writePageMarkdown writes the markdown file for a processed page
func (e *Extractor) writePageMarkdown(result *PageResult) error {
	buf := e.bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer e.bufferPool.Put(buf)

	// Embed error comment if any extraction errors occurred
	buf.WriteString(result.MarkdownErrorComment())

	buf.WriteString(fmt.Sprintf("# Page %d\n\n", result.Page))

	if strings.TrimSpace(result.Text) != "" {
		buf.WriteString("## Text Content\n\n")
		buf.WriteString("```\n")
		buf.WriteString(result.Text)
		buf.WriteString("\n```\n\n")
	}

	if len(result.Images) > 0 {
		buf.WriteString("## Images\n\n")
		for _, img := range result.Images {
			relativePath := filepath.ToSlash(filepath.Join("images", img))
			buf.WriteString(fmt.Sprintf("![Image from page %d](%s)\n\n", result.Page, relativePath))
		}
	}

	outputFilename := filepath.Join(e.outputDir, fmt.Sprintf("page_%d.md", result.Page))

	file, err := os.Create(outputFilename)
	if err != nil {
		return fmt.Errorf("failed to create markdown file for page %d: %v", result.Page, err)
	}
	defer file.Close()

	if _, err := io.Copy(file, bytes.NewReader(buf.Bytes())); err != nil {
		return fmt.Errorf("failed to write markdown file for page %d: %v", result.Page, err)
	}

	e.logAsync("Saved page %d to %s", result.Page, outputFilename)
	return nil
}

// extractTextFromPage uses "pdftotext" to extract text from a specific page
func (e *Extractor) extractTextFromPage(page int) (string, error) {
	// Check cache first
	cacheKey := fmt.Sprintf("%s:text:%d", e.pdfFile, page)
	if cached, exists := e.resultCache.Get(cacheKey, e.metricsCollector); exists {
		e.logAsync("Cache hit for text on page %d", page)
		return cached.Text, nil
	}

	tempFile, err := os.CreateTemp("", "page_*.txt")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %v", err)
	}
	tempFilePath := tempFile.Name()
	tempFile.Close()
	defer os.Remove(tempFilePath)

	// Use command pool for better command execution
	cmd := exec.CommandContext(e.ctx, "pdftotext", "-f", strconv.Itoa(page), "-l", strconv.Itoa(page), e.pdfFile, tempFilePath)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("pdftotext failed for page %d: %w", page, err)
	}

	content, err := os.ReadFile(tempFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read temporary file: %v", err)
	}

	text := string(content)

	// Cache the result
	e.resultCache.Set(cacheKey, &cache.CachedResult{Text: text})
	e.metricsCollector.RecordTextExtracted(len(text))

	return text, nil
}

// extractImagesFromPage uses "pdfimages" to extract images from a specific page
func (e *Extractor) extractImagesFromPage(page int) ([]string, []error) {
	// Check cache first
	cacheKey := fmt.Sprintf("%s:images:%d", e.pdfFile, page)
	if cached, exists := e.resultCache.Get(cacheKey, e.metricsCollector); exists {
		e.logAsync("Cache hit for images on page %d", page)
		return cached.Images, nil
	}

	// Get a temporary directory from pool
	tempDir, err := e.tempDirPool.GetTempDir()
	if err != nil {
		return nil, []error{fmt.Errorf("failed to get temporary directory: %v", err)}
	}
	defer e.tempDirPool.ReturnTempDir(tempDir)

	// Build and run the command
	outputPrefix := filepath.Join(tempDir, "img")
	cmd := exec.CommandContext(e.ctx, "pdfimages", "-j", "-f", strconv.Itoa(page), "-l", strconv.Itoa(page), e.pdfFile, outputPrefix)
	if out, err := cmd.CombinedOutput(); err != nil {
		return nil, []error{fmt.Errorf("pdfimages failed for page %d: %w (%s)", page, err, strings.TrimSpace(string(out)))}
	}

	// List files in the temporary directory
	files, err := os.ReadDir(tempDir)
	if err != nil {
		return nil, []error{fmt.Errorf("failed to list files in temp dir: %v", err)}
	}

	if len(files) == 0 {
		// No images extracted, cache empty result
		e.resultCache.Set(cacheKey, &cache.CachedResult{Images: []string{}})
		return []string{}, nil
	}

	// Pre-allocate result slice
	resultImages := make([]string, 0, len(files))

	// Process images with limited concurrency
	maxConcurrent := 5
	if len(files) < maxConcurrent {
		maxConcurrent = len(files)
	}
	sem := make(chan struct{}, maxConcurrent)

	type imageResult struct {
		index int
		name  string
		err   error
	}
	resultChan := make(chan imageResult, len(files))

	var wg sync.WaitGroup

	fileIdx := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		wg.Add(1)
		sem <- struct{}{}

		go func(fileInfo os.DirEntry, index int) {
			defer wg.Done()
			defer func() { <-sem }()

			srcPath := filepath.Join(tempDir, fileInfo.Name())
			ext := filepath.Ext(fileInfo.Name())
			newName := fmt.Sprintf("page_%d_img_%03d%s", page, index+1, ext)
			destPath := filepath.Join(e.imageDir, newName)

			// Use optimized copy
			if err := e.copyFileOptimized(srcPath, destPath); err != nil {
				e.logAsync("Failed to copy image file %s: %v", srcPath, err)
				resultChan <- imageResult{index: index, err: fmt.Errorf("failed to copy %s: %w", fileInfo.Name(), err)}
				return
			}

			resultChan <- imageResult{index: index, name: newName}
		}(file, fileIdx)
		fileIdx++
	}

	wg.Wait()
	close(resultChan)

	// Collect results in order
	results := make([]imageResult, 0, len(resultChan))
	for result := range resultChan {
		results = append(results, result)
	}

	// Sort by index to maintain order
	sort.Slice(results, func(i, j int) bool {
		return results[i].index < results[j].index
	})

	// Separate successes and failures
	var imageErrors []error
	for _, result := range results {
		if result.err != nil {
			imageErrors = append(imageErrors, result.err)
		} else {
			resultImages = append(resultImages, result.name)
		}
	}

	// Cache successful results (even partial)
	if len(resultImages) > 0 {
		e.resultCache.Set(cacheKey, &cache.CachedResult{Images: resultImages})
	}
	e.metricsCollector.RecordImagesExtracted(len(resultImages))
	e.logAsync("Extracted %d images from page %d (%d errors)", len(resultImages), page, len(imageErrors))

	return resultImages, imageErrors
}

// copyFileOptimized copies a file with optimized buffer size using buffer pool manager
func (e *Extractor) copyFileOptimized(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get file info for adaptive buffer sizing
	fileInfo, err := in.Stat()
	if err != nil {
		// Fallback to default buffer
		buffer := e.bufferManager.GetBuffer(64 * 1024)
		defer e.bufferManager.PutBuffer(buffer)
		if _, err = io.CopyBuffer(out, in, buffer); err != nil {
			return err
		}
		return out.Sync()
	}

	// Get appropriately sized buffer from manager
	fileSize := fileInfo.Size()
	buffer := e.bufferManager.GetBuffer(fileSize)
	defer e.bufferManager.PutBuffer(buffer)

	bytesCopied, err := io.CopyBuffer(out, in, buffer)
	if err != nil {
		return err
	}
	if bytesCopied > 0 {
		e.metricsCollector.RecordBytesProcessed(bytesCopied)
	}
	return out.Sync()
}

// createMainMarkdown generates an index Markdown file linking to all pages
func (e *Extractor) createMainMarkdown() error {
	totalPages, err := e.getPageCount()
	if err != nil {
		return err
	}
	var content strings.Builder
	baseName := filepath.Base(e.pdfFile)
	content.WriteString(fmt.Sprintf("# %s - PDF Extract\n\n", baseName))
	content.WriteString(fmt.Sprintf("This document contains the extracted content from `%s`.\n\n", e.pdfFile))
	content.WriteString("## Pages\n\n")
	for i := 1; i <= totalPages; i++ {
		content.WriteString(fmt.Sprintf("- [Page %d](page_%d.md)\n", i, i))
	}
	mainMdPath := filepath.Join(e.outputDir, "index.md")
	if err := os.WriteFile(mainMdPath, []byte(content.String()), 0644); err != nil {
		return fmt.Errorf("failed to write main markdown file: %v", err)
	}
	e.logAsync("Created main index file at %s", mainMdPath)
	return nil
}