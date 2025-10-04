package main

import (
    "bufio"
    "bytes"
    "context"
    "crypto/rand"
    "encoding/hex"
    "flag"
    "fmt"
    "gopkg.in/yaml.v2"
    "io"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "runtime/pprof"
    "sort"
    "strconv"
    "strings"
    "sync"
    "sync/atomic"
    "time"
    "unsafe"
)

// Options holds configuration options loaded from flags or a YAML config.
type Options struct {
	PdfFile      string `yaml:"pdf_file"`
	OutputDir    string `yaml:"output_dir"`
	ProcessCount int    `yaml:"process_count"`
	ScanDir      string `yaml:"scan_dir"`
	CopyDir      string `yaml:"copy_dir"`
	Profile      string `yaml:"profile"`      // CPU or memory profiling
	CacheResults bool   `yaml:"cache_results"` // Cache extracted text/images
}

// MetricsCollector collects performance metrics
type MetricsCollector struct {
	mu              sync.Mutex
	pagesProcessed  int64
	textExtracted   int64
	imagesExtracted int64
	bytesProcessed  int64
	processingTimes []time.Duration
	pdfSizes        map[string]int64
	// Cache performance metrics
	cacheHits       int64
	cacheMisses     int64
	// Buffer pool metrics
	bufferPoolHits    int64
	bufferPoolMisses  int64
	bufferPoolCreated int64
	// Page processing metrics
	pageQueueDepth    []int
	workerUtilization map[int]time.Duration
}

// NewMetricsCollector creates a new metrics collector
func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{
		processingTimes:   make([]time.Duration, 0),
		pdfSizes:          make(map[string]int64),
		pageQueueDepth:    make([]int, 0),
		workerUtilization: make(map[int]time.Duration),
	}
}

// RecordPageProcessed records a processed page
func (m *MetricsCollector) RecordPageProcessed() {
	atomic.AddInt64(&m.pagesProcessed, 1)
}

// RecordTextExtracted records text extraction
func (m *MetricsCollector) RecordTextExtracted(bytes int) {
	atomic.AddInt64(&m.textExtracted, int64(bytes))
}

// RecordImagesExtracted records image extraction
func (m *MetricsCollector) RecordImagesExtracted(count int) {
	atomic.AddInt64(&m.imagesExtracted, int64(count))
}

// RecordBytesProcessed records bytes processed for I/O operations
func (m *MetricsCollector) RecordBytesProcessed(bytes int64) {
	atomic.AddInt64(&m.bytesProcessed, bytes)
}

// RecordProcessingTime records processing time
func (m *MetricsCollector) RecordProcessingTime(d time.Duration) {
	m.mu.Lock()
	m.processingTimes = append(m.processingTimes, d)
	m.mu.Unlock()
}

// RecordCacheHit records a cache hit
func (m *MetricsCollector) RecordCacheHit() {
	atomic.AddInt64(&m.cacheHits, 1)
}

// RecordCacheMiss records a cache miss
func (m *MetricsCollector) RecordCacheMiss() {
	atomic.AddInt64(&m.cacheMisses, 1)
}

// RecordBufferPoolHit records a buffer pool hit
func (m *MetricsCollector) RecordBufferPoolHit() {
	atomic.AddInt64(&m.bufferPoolHits, 1)
}

// RecordBufferPoolMiss records a buffer pool miss
func (m *MetricsCollector) RecordBufferPoolMiss() {
	atomic.AddInt64(&m.bufferPoolMisses, 1)
}

// RecordBufferPoolCreated records buffer creation
func (m *MetricsCollector) RecordBufferPoolCreated(size int) {
	atomic.AddInt64(&m.bufferPoolCreated, int64(size))
}

// RecordPageQueueDepth records the current page queue depth
func (m *MetricsCollector) RecordPageQueueDepth(depth int) {
	m.mu.Lock()
	m.pageQueueDepth = append(m.pageQueueDepth, depth)
	m.mu.Unlock()
}

// RecordWorkerTime records time spent by a worker
func (m *MetricsCollector) RecordWorkerTime(workerID int, duration time.Duration) {
	m.mu.Lock()
	m.workerUtilization[workerID] += duration
	m.mu.Unlock()
}

// GetCacheHitRate returns the cache hit rate as a percentage
func (m *MetricsCollector) GetCacheHitRate() float64 {
	hits := atomic.LoadInt64(&m.cacheHits)
	misses := atomic.LoadInt64(&m.cacheMisses)
	total := hits + misses
	if total == 0 {
		return 0.0
	}
	return float64(hits) / float64(total) * 100.0
}

// PrintSummary prints a summary of metrics with context-aware labels
func (m *MetricsCollector) PrintSummary(context string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.processingTimes) == 0 {
		return
	}

	log.Println("\n" + strings.Repeat("=", 50))
	log.Println("PERFORMANCE METRICS")
	log.Println(strings.Repeat("=", 50))
	log.Printf("Pages processed: %d", atomic.LoadInt64(&m.pagesProcessed))
	log.Printf("Text extracted: %.2f MB", float64(atomic.LoadInt64(&m.textExtracted))/(1024*1024))
	log.Printf("Images extracted: %d", atomic.LoadInt64(&m.imagesExtracted))
	log.Printf("Total bytes processed: %.2f MB", float64(atomic.LoadInt64(&m.bytesProcessed))/(1024*1024))

	if len(m.processingTimes) > 0 {
		var total time.Duration
		for _, t := range m.processingTimes {
			total += t
		}
		avg := total / time.Duration(len(m.processingTimes))

	// Context-aware timing labels
	if context == "batch" {
		log.Printf("Average processing time per PDF: %v", avg)
		log.Printf("Total processing time across all PDFs: %v", total)
	} else {
		log.Printf("Average processing time per page: %v", avg)
		log.Printf("Total processing time: %v", total)
	}
	}

	// Cache performance metrics
	cacheHits := atomic.LoadInt64(&m.cacheHits)
	cacheMisses := atomic.LoadInt64(&m.cacheMisses)
	totalCacheAccess := cacheHits + cacheMisses
	if totalCacheAccess > 0 {
		hitRate := float64(cacheHits) / float64(totalCacheAccess) * 100.0
		log.Printf("Cache hits: %d, Cache misses: %d", cacheHits, cacheMisses)
		log.Printf("Cache hit rate: %.2f%%", hitRate)
	}

	// Buffer pool metrics
	bpHits := atomic.LoadInt64(&m.bufferPoolHits)
	bpMisses := atomic.LoadInt64(&m.bufferPoolMisses)
	totalBPAccess := bpHits + bpMisses
	if totalBPAccess > 0 {
		bpHitRate := float64(bpHits) / float64(totalBPAccess) * 100.0
		log.Printf("Buffer pool hits: %d, misses: %d (%.2f%% hit rate)", bpHits, bpMisses, bpHitRate)
		log.Printf("Buffers created: %.2f MB", float64(atomic.LoadInt64(&m.bufferPoolCreated))/(1024*1024))
	}

	// Worker utilization
	if len(m.workerUtilization) > 0 {
		var totalWorkerTime time.Duration
		for _, t := range m.workerUtilization {
			totalWorkerTime += t
		}
		avgWorkerTime := totalWorkerTime / time.Duration(len(m.workerUtilization))
		log.Printf("Average worker utilization: %v", avgWorkerTime)
	}

	// Page queue depth analysis
	if len(m.pageQueueDepth) > 0 {
		var totalDepth int
		maxDepth := 0
		for _, d := range m.pageQueueDepth {
			totalDepth += d
			if d > maxDepth {
				maxDepth = d
			}
		}
		avgDepth := float64(totalDepth) / float64(len(m.pageQueueDepth))
		log.Printf("Page queue - Avg depth: %.2f, Max depth: %d", avgDepth, maxDepth)
	}
}

// ResultCache caches extraction results
type ResultCache struct {
	mu    sync.RWMutex
	cache map[string]*CachedResult
}

// CachedResult holds cached extraction results
type CachedResult struct {
	Text      string
	Images    []string
	Timestamp time.Time
}

// NewResultCache creates a new result cache
func NewResultCache() *ResultCache {
	return &ResultCache{
		cache: make(map[string]*CachedResult),
	}
}

// Get retrieves a cached result
func (c *ResultCache) Get(key string, metricsCollector *MetricsCollector) (*CachedResult, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	result, exists := c.cache[key]
	if exists && metricsCollector != nil {
		metricsCollector.RecordCacheHit()
	} else if metricsCollector != nil {
		metricsCollector.RecordCacheMiss()
	}
	return result, exists
}

// Set stores a result in cache
func (c *ResultCache) Set(key string, result *CachedResult) {
	c.mu.Lock()
	defer c.mu.Unlock()
	result.Timestamp = time.Now()
	c.cache[key] = result
}

// BufferPoolManager manages multiple buffer pools with different sizes
type BufferPoolManager struct {
	smallPool  *sync.Pool  // 32KB buffers
	mediumPool *sync.Pool  // 128KB buffers
	largePool  *sync.Pool  // 256KB buffers
	xlargePool *sync.Pool  // 1MB buffers
	metricsCollector *MetricsCollector
}

// NewBufferPoolManager creates a new buffer pool manager
func NewBufferPoolManager(metrics *MetricsCollector) *BufferPoolManager {
	return &BufferPoolManager{
		smallPool: &sync.Pool{
			New: func() interface{} {
				if metrics != nil {
					metrics.RecordBufferPoolCreated(32 * 1024)
				}
				return make([]byte, 32*1024)
			},
		},
		mediumPool: &sync.Pool{
			New: func() interface{} {
				if metrics != nil {
					metrics.RecordBufferPoolCreated(128 * 1024)
				}
				return make([]byte, 128*1024)
			},
		},
		largePool: &sync.Pool{
			New: func() interface{} {
				if metrics != nil {
					metrics.RecordBufferPoolCreated(256 * 1024)
				}
				return make([]byte, 256*1024)
			},
		},
		xlargePool: &sync.Pool{
			New: func() interface{} {
				if metrics != nil {
					metrics.RecordBufferPoolCreated(1024 * 1024)
				}
				return make([]byte, 1024*1024)
			},
		},
		metricsCollector: metrics,
	}
}

// GetBuffer returns an appropriately sized buffer from the pool
func (m *BufferPoolManager) GetBuffer(sizeHint int64) []byte {
	var pool *sync.Pool
	if sizeHint < 64*1024 {
		pool = m.smallPool
	} else if sizeHint < 256*1024 {
		pool = m.mediumPool
	} else if sizeHint < 512*1024 {
		pool = m.largePool
	} else {
		pool = m.xlargePool
	}

	buffer := pool.Get().([]byte)
	if m.metricsCollector != nil {
		m.metricsCollector.RecordBufferPoolHit()
	}
	return buffer
}

// PutBuffer returns a buffer to the appropriate pool
func (m *BufferPoolManager) PutBuffer(buffer []byte) {
	size := len(buffer)
	var pool *sync.Pool

	switch size {
	case 32 * 1024:
		pool = m.smallPool
	case 128 * 1024:
		pool = m.mediumPool
	case 256 * 1024:
		pool = m.largePool
	case 1024 * 1024:
		pool = m.xlargePool
	default:
		// Buffer doesn't match any pool size, don't return it
		return
	}

	pool.Put(buffer)
}

// CommandPool manages a pool of reusable command executors
type CommandPool struct {
	mu       sync.Mutex
	cmdCache map[string]*exec.Cmd
	ctx      context.Context
}

// NewCommandPool creates a new command pool
func NewCommandPool(ctx context.Context) *CommandPool {
	return &CommandPool{
		cmdCache: make(map[string]*exec.Cmd),
		ctx:      ctx,
	}
}

// ExecuteCommand executes a command with timeout and caching
func (p *CommandPool) ExecuteCommand(name string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(p.ctx, 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	return cmd.Output()
}

// Extractor encapsulates the PDF extraction logic.
type Extractor struct {
	pdfFile          string
	outputDir        string
	processCount     int
	imageDir        string
	bufferPool       *sync.Pool
	bufferManager    *BufferPoolManager // Managed buffer pools
	logChan          chan string
	tempDirPool      chan string
	pageCount        int
	pageCountMu      sync.Mutex
	resultCache      *ResultCache
	metricsCollector *MetricsCollector
	ctx              context.Context
	cancel           context.CancelFunc
	commandPool      *CommandPool
}

// NewExtractor creates a new extractor instance.
func NewExtractor(pdfFile, outputDir string, processCount int) (*Extractor, error) {
	// Ensure the PDF file exists.
	if _, err := os.Stat(pdfFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("PDF file not found: %s", pdfFile)
	}
	// If no output directory is provided, use the PDF filename (without extension).
	if outputDir == "" {
		base := filepath.Base(pdfFile)
		ext := filepath.Ext(pdfFile)
		outputDir = strings.TrimSuffix(base, ext)
	}
	// Create the output directory.
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, err
	}
	imageDir := filepath.Join(outputDir, "images")
	if err := os.MkdirAll(imageDir, 0755); err != nil {
		return nil, err
	}
	// Use provided process count or default to the number of CPUs (or 4).
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
	metricsCollector := NewMetricsCollector()

	// Create managed buffer pool manager
	bufferManager := NewBufferPoolManager(metricsCollector)

	// Calculate optimal channel buffer sizes based on process count
	logChannelSize := 200
	if processCount > 8 {
		logChannelSize = processCount * 50 // Scale with workers
	}

	// Create async logging channel with adaptive buffer
	logChan := make(chan string, logChannelSize)

	// Pre-create temp directories pool with adaptive sizing
	tempPoolSize := processCount * 2
	if processCount > 8 {
		tempPoolSize = processCount + 4 // Don't scale linearly for high counts
	}
	tempDirPool := make(chan string, tempPoolSize)

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
		resultCache:      NewResultCache(),
		metricsCollector: metricsCollector,
		ctx:              ctx,
		cancel:           cancel,
		commandPool:      NewCommandPool(ctx),
	}

	// Start async logger
	go extractor.asyncLogger()

	// Pre-create temp directories
	go extractor.initTempDirPool()

	return extractor, nil
}

// getPageCount runs "pdfinfo" and parses the page count with caching.
func (e *Extractor) getPageCount() (int, error) {
	// Check cache first using atomic load for better performance
	if count := atomic.LoadInt32((*int32)(unsafe.Pointer(&e.pageCount))); count > 0 {
		return int(count), nil
	}

	// Use buffered reading for better performance
	cmd := exec.Command("pdfinfo", e.pdfFile)
	output, err := cmd.StdoutPipe()
	if err != nil {
		return 0, fmt.Errorf("failed to create stdout pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return 0, fmt.Errorf("failed to run pdfinfo: %v", err)
	}

	// Read and parse output line by line for efficiency
	scanner := bufio.NewScanner(output)
	var pageCount int
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Pages:") {
			// Extract page count from this line
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				if count, err := strconv.Atoi(parts[1]); err == nil {
					pageCount = count
					found = true
					break
				}
			}
		}
	}

	if err := cmd.Wait(); err != nil {
		return 0, fmt.Errorf("pdfinfo command failed: %v", err)
	}

	if !found || pageCount == 0 {
		return 0, fmt.Errorf("failed to parse page count from pdfinfo for %s", e.pdfFile)
	}

	// Cache the result using atomic store
	atomic.StoreInt32((*int32)(unsafe.Pointer(&e.pageCount)), int32(pageCount))

	return pageCount, nil
}

// asyncLogger handles async logging to prevent blocking workers
func (e *Extractor) asyncLogger() {
	for msg := range e.logChan {
		log.Print(msg)
	}
}

// initTempDirPool pre-creates temporary directories for reuse
func (e *Extractor) initTempDirPool() {
	for i := 0; i < cap(e.tempDirPool); i++ {
		tempDir, err := os.MkdirTemp("", "pdf_images_*")
		if err != nil {
			e.logAsync("Failed to pre-create temp dir: %v", err)
			continue
		}
		select {
		case e.tempDirPool <- tempDir:
		default:
			os.RemoveAll(tempDir)
		}
	}
}

// getTempDir gets a temp directory from the pool or creates a new one
func (e *Extractor) getTempDir() (string, error) {
	select {
	case dir := <-e.tempDirPool:
		// Clear the directory before reuse
		entries, _ := os.ReadDir(dir)
		for _, entry := range entries {
			os.RemoveAll(filepath.Join(dir, entry.Name()))
		}
		return dir, nil
	default:
		// Pool empty, create new one
		return os.MkdirTemp("", "pdf_images_*")
	}
}

// returnTempDir returns a temp directory to the pool or removes it
func (e *Extractor) returnTempDir(dir string) {
	select {
	case e.tempDirPool <- dir:
		// Successfully returned to pool
	default:
		// Pool full, remove the directory
		os.RemoveAll(dir)
	}
}

// logAsync sends a log message asynchronously
func (e *Extractor) logAsync(format string, v ...interface{}) {
	select {
	case e.logChan <- fmt.Sprintf(format, v...):
	default:
		// Channel full, log synchronously as fallback
		log.Printf(format, v...)
	}
}

// extractImagesFromPage uses "pdfimages" to extract images from a specific page.
func (e *Extractor) extractImagesFromPage(page int) ([]string, error) {
	// Check cache first
	cacheKey := fmt.Sprintf("%s:images:%d", e.pdfFile, page)
	if cached, exists := e.resultCache.Get(cacheKey, e.metricsCollector); exists {
		e.logAsync("Cache hit for images on page %d", page)
		return cached.Images, nil
	}

	// Get a temporary directory from pool.
	tempDir, err := e.getTempDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get temporary directory: %v", err)
	}
	defer e.returnTempDir(tempDir)

	// Build and run the command with timeout
	outputPrefix := filepath.Join(tempDir, "img")
	out, err := e.commandPool.ExecuteCommand("pdfimages", "-j", "-f", strconv.Itoa(page), "-l", strconv.Itoa(page), e.pdfFile, outputPrefix)
	if err != nil {
		e.logAsync("Error extracting images from page %d: %s", page, string(out))
		return nil, nil // Return empty slice on error.
	}

	// List files in the temporary directory.
	files, err := os.ReadDir(tempDir)
	if err != nil {
		return nil, fmt.Errorf("failed to list files in temp dir: %v", err)
	}

	if len(files) == 0 {
		// No images extracted, return empty result
		e.resultCache.Set(cacheKey, &CachedResult{Images: []string{}})
		return []string{}, nil
	}

	// Pre-allocate result slice to avoid rebuilding
	resultImages := make([]string, 0, len(files))

	// Use semaphore to limit concurrent file operations
	maxConcurrent := 5
	if len(files) < maxConcurrent {
		maxConcurrent = len(files)
	}
	sem := make(chan struct{}, maxConcurrent)

	type imageResult struct {
		index int
		name  string
	}
	resultChan := make(chan imageResult, len(files))

	var wg sync.WaitGroup
	var copyErr error
	var errOnce sync.Once

	fileIdx := 0
	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		wg.Add(1)
		sem <- struct{}{} // Acquire semaphore

		go func(fileInfo os.DirEntry, index int) {
			defer wg.Done()
			defer func() { <-sem }() // Release semaphore

			srcPath := filepath.Join(tempDir, fileInfo.Name())
			ext := filepath.Ext(fileInfo.Name())
			newName := fmt.Sprintf("page_%d_img_%03d%s", page, index+1, ext)
			destPath := filepath.Join(e.imageDir, newName)

			// Use optimized copy with buffer manager
			if err := e.copyFileOptimized(srcPath, destPath); err != nil {
				errOnce.Do(func() {
					copyErr = fmt.Errorf("failed to copy %s: %v", fileInfo.Name(), err)
				})
				e.logAsync("Failed to copy image file %s: %v", srcPath, err)
				return
			}

			resultChan <- imageResult{index: index, name: newName}
		}(file, fileIdx)
		fileIdx++
	}

	wg.Wait()
	close(sem)
	close(resultChan)

	// Check for errors
	if copyErr != nil {
		return nil, copyErr
	}

	// Collect results in order
	results := make([]imageResult, 0, len(resultChan))
	for result := range resultChan {
		results = append(results, result)
	}

	// Sort by index to maintain order
	sort.Slice(results, func(i, j int) bool {
		return results[i].index < results[j].index
	})

	// Build final image list
	for _, result := range results {
		resultImages = append(resultImages, result.name)
	}

	// Cache the result
	e.resultCache.Set(cacheKey, &CachedResult{Images: resultImages})
	e.metricsCollector.RecordImagesExtracted(len(resultImages))
	e.logAsync("Extracted %d images from page %d", len(resultImages), page)
	return resultImages, nil
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
	defer func() {
		_ = out.Close()
	}()

	// Get file info for adaptive buffer sizing
	fileInfo, err := in.Stat()
	if err != nil {
		// Fallback to default buffer if we can't get file info
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

// extractTextFromPage uses "pdftotext" to extract text from a specific page.
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
	out, err := e.commandPool.ExecuteCommand("pdftotext", "-f", strconv.Itoa(page), "-l", strconv.Itoa(page), e.pdfFile, tempFilePath)
	if err != nil {
		e.logAsync("Error extracting text from page %d: %s", page, string(out))
		return "", nil
	}

	// Use buffered reading for better I/O
	file, err := os.Open(tempFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to open temporary file: %v", err)
	}
	defer file.Close()

	var content bytes.Buffer
	bufReader := make([]byte, 32*1024) // 32KB buffer
	for {
		n, err := file.Read(bufReader)
		if n > 0 {
			content.Write(bufReader[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to read temporary file: %v", err)
		}
	}

	text := content.String()

	// Cache the result
	e.resultCache.Set(cacheKey, &CachedResult{Text: text})
	e.metricsCollector.RecordTextExtracted(len(text))

	return text, nil
}

// processPage extracts text and images from a page and writes a Markdown file.
func (e *Extractor) processPage(page int) error {
	start := time.Now()
	defer func() {
		e.metricsCollector.RecordPageProcessed()
		e.metricsCollector.RecordProcessingTime(time.Since(start))
	}()

	// Check if context is cancelled
	select {
	case <-e.ctx.Done():
		return e.ctx.Err()
	default:
	}

	// Extract text and images in parallel
	var text string
	var images []string
	var textErr, imgErr error
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		text, textErr = e.extractTextFromPage(page)
	}()
	go func() {
		defer wg.Done()
		images, imgErr = e.extractImagesFromPage(page)
	}()
	wg.Wait()

	if textErr != nil {
		e.logAsync("Error extracting text from page %d: %v", page, textErr)
	}
	if imgErr != nil {
		e.logAsync("Error extracting images from page %d: %v", page, imgErr)
	}

	// Get buffer from pool
	buf := e.bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer e.bufferPool.Put(buf)

	buf.WriteString(fmt.Sprintf("# Page %d\n\n", page))

	if strings.TrimSpace(text) != "" {
		buf.WriteString("## Text Content\n\n")
		buf.WriteString("```\n")
		buf.WriteString(text)
		buf.WriteString("\n```\n\n")
	}

	if len(images) > 0 {
		buf.WriteString("## Images\n\n")
		for _, img := range images {
			relativePath := filepath.ToSlash(filepath.Join("images", img))
			buf.WriteString(fmt.Sprintf("![Image from page %d](%s)\n\n", page, relativePath))
		}
		e.metricsCollector.RecordImagesExtracted(len(images))
	}

	outputFilename := filepath.Join(e.outputDir, fmt.Sprintf("page_%d.md", page))

	// Write with buffered I/O
	file, err := os.Create(outputFilename)
	if err != nil {
		return fmt.Errorf("failed to create markdown file for page %d: %v", page, err)
	}
	defer file.Close()

	bufWriter := make([]byte, 64*1024) // 64KB write buffer
	n, err := io.CopyBuffer(file, bytes.NewReader(buf.Bytes()), bufWriter)
	if err != nil {
		return fmt.Errorf("failed to write markdown file for page %d: %v", page, err)
	}
	if n != int64(buf.Len()) {
		return fmt.Errorf("incomplete write for page %d", page)
	}

	e.logAsync("Saved page %d to %s", page, outputFilename)
	return nil
}

// createMainMarkdown generates an index Markdown file linking to all pages.
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

// extractPages processes all pages concurrently.
func (e *Extractor) extractPages() error {
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
		// Scale back workers for very large PDFs to avoid resource exhaustion
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
	workerTimes := make([]time.Duration, processes)

	// Start worker goroutines with performance tracking
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

			workerTimes[workerID] = time.Since(workerStart)
		}(i)
	}

	// Producer: enqueue page numbers with batching for better performance
	go func() {
		defer close(pages)
		batchSize := 10
		if totalPages > 100 {
			batchSize = 20
		}

		for i := 1; i <= totalPages; i++ {
			select {
			case pages <- i:
				// Yield occasionally for very large PDFs
				if i%batchSize == 0 && totalPages > 100 {
					runtime.Gosched()
				}
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
	e.cancel() // Cancel context
	close(e.tempDirPool)
	for dir := range e.tempDirPool {
		os.RemoveAll(dir)
	}

	// Close log channel and ensure all logs are flushed
	close(e.logChan)
	time.Sleep(100 * time.Millisecond) // Give logger time to flush

	return nil
}

// calculateOptimalBufferSize determines the best channel buffer size
func calculateOptimalBufferSize(totalPages, processes int) int {
	// Base calculation
	bufferSize := min(totalPages, processes*4)

	// Adjust based on PDF size
	if totalPages <= 10 {
		bufferSize = totalPages // Small PDFs: buffer all pages
	} else if totalPages <= 50 {
		bufferSize = min(totalPages, processes*2) // Medium PDFs: moderate buffering
	} else if totalPages <= 200 {
		bufferSize = processes * 4 // Large PDFs: standard buffering
	} else if totalPages <= 1000 {
		bufferSize = processes * 6 // Very large PDFs: increased buffering
	} else {
		bufferSize = processes * 8 // Huge PDFs: maximum buffering
		// Cap at reasonable maximum
		if bufferSize > 100 {
			bufferSize = 100
		}
	}

	return bufferSize
}

// calculateOptimalConcurrentPDFs determines optimal PDF concurrency
func calculateOptimalConcurrentPDFs(processCount int) int {
	concurrentPDFs := processCount / 4

	// Apply min/max bounds based on system capabilities
	if concurrentPDFs < 2 {
		concurrentPDFs = 2
	} else if concurrentPDFs > 8 {
		// Cap at 8 to prevent resource exhaustion
		concurrentPDFs = 8
	}

	// Adjust based on available memory
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	availableMemGB := float64(memStats.Sys) / (1024 * 1024 * 1024)

	if availableMemGB < 2 {
		// Low memory: reduce concurrency
		if concurrentPDFs > 2 {
			concurrentPDFs = 2
		}
	} else if availableMemGB < 4 {
		// Medium memory: moderate concurrency
		if concurrentPDFs > 4 {
			concurrentPDFs = 4
		}
	}

	return concurrentPDFs
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// copyFile copies a file from src to dst.
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}

// generateRandomHex creates a random hex string of n bytes.
func generateRandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// loadConfig reads a YAML configuration file.
func loadConfig(path string) (*Options, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var opts Options
	if err := yaml.Unmarshal(data, &opts); err != nil {
		return nil, err
	}
	return &opts, nil
}

// PDFScanner handles scanning and copying PDF files
type PDFScanner struct {
	scanDir       string
	copyDir       string
	logChan       chan string
	bufferManager *BufferPoolManager
	metricsCollector *MetricsCollector
}

// NewPDFScanner creates a new PDF scanner instance
func NewPDFScanner(scanDir, copyDir string) (*PDFScanner, error) {
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

	metricsCollector := NewMetricsCollector()
	logChan := make(chan string, 100)
	scanner := &PDFScanner{
		scanDir:          scanDir,
		copyDir:          copyDir,
		logChan:          logChan,
		bufferManager:    NewBufferPoolManager(metricsCollector),
		metricsCollector: metricsCollector,
	}

	// Start async logger
	go scanner.asyncLogger()

	return scanner, nil
}

// asyncLogger handles async logging for scanner
func (s *PDFScanner) asyncLogger() {
	for msg := range s.logChan {
		log.Print(msg)
	}
}

// logAsync sends a log message asynchronously
func (s *PDFScanner) logAsync(format string, v ...interface{}) {
	select {
	case s.logChan <- fmt.Sprintf(format, v...):
	default:
		log.Printf(format, v...)
	}
}

// findPDFs recursively finds all PDF files in the scan directory
func (s *PDFScanner) findPDFs() ([]string, error) {
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

// scanAndCopy scans for PDFs and copies them with concurrent processing
func (s *PDFScanner) scanAndCopy() error {
	s.logAsync("Scanning directory: %s", s.scanDir)

	// Find all PDF files
	pdfFiles, err := s.findPDFs()
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
	close(s.logChan)
	time.Sleep(100 * time.Millisecond) // Give logger time to flush

	return nil
}

// BatchProcessor handles batch processing of multiple PDF files
type BatchProcessor struct {
	inputDir        string
	outputDir       string
	processCount    int
	logChan         chan string
	concurrentPDFs  int    // Number of PDFs to process simultaneously
	workerPool      chan struct{} // Semaphore for limiting concurrent PDFs
	metricsCollector *MetricsCollector
	bufferManager   *BufferPoolManager
}

// NewBatchProcessor creates a new batch processor
func NewBatchProcessor(inputDir, outputDir string, processCount int) (*BatchProcessor, error) {
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

	logChan := make(chan string, logChannelSize)
	workerPool := make(chan struct{}, concurrentPDFs)

	metricsCollector := NewMetricsCollector()
	batch := &BatchProcessor{
		inputDir:        inputDir,
		outputDir:       outputDir,
		processCount:    processCount,
		logChan:         logChan,
		concurrentPDFs:  concurrentPDFs,
		workerPool:      workerPool,
		metricsCollector: metricsCollector,
		bufferManager:   NewBufferPoolManager(metricsCollector),
	}

	// Start async logger
	go batch.asyncLogger()

	return batch, nil
}

// asyncLogger handles async logging for batch processor
func (b *BatchProcessor) asyncLogger() {
	for msg := range b.logChan {
		log.Print(msg)
	}
}

// logAsync sends a log message asynchronously
func (b *BatchProcessor) logAsync(format string, v ...interface{}) {
	select {
	case b.logChan <- fmt.Sprintf(format, v...):
	default:
		log.Printf(format, v...)
	}
}

// findPDFs finds all PDF files in the input directory
func (b *BatchProcessor) findPDFs() ([]string, error) {
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
func (b *BatchProcessor) processPDF(pdfPath string, index, total int) error {
	// Get the PDF filename without extension for output directory
	baseName := filepath.Base(pdfPath)
	nameWithoutExt := strings.TrimSuffix(baseName, filepath.Ext(baseName))
	pdfOutputDir := filepath.Join(b.outputDir, nameWithoutExt)

	b.logAsync("[%d/%d] Processing: %s", index, total, baseName)
	startTime := time.Now()

	// Get file size for metrics
	if info, err := os.Stat(pdfPath); err == nil {
		b.metricsCollector.mu.Lock()
		b.metricsCollector.pdfSizes[baseName] = info.Size()
		b.metricsCollector.mu.Unlock()
	}

	// Create extractor for this PDF with adjusted process count
	processesPerPDF := b.processCount / b.concurrentPDFs
	if processesPerPDF < 1 {
		processesPerPDF = 1
	}

	extractor, err := NewExtractor(pdfPath, pdfOutputDir, processesPerPDF)
	if err != nil {
		return fmt.Errorf("failed to initialize extractor: %v", err)
	}

	// Share metrics collector and buffer manager
	extractor.metricsCollector = b.metricsCollector
	extractor.bufferManager = b.bufferManager

	// Extract pages
	if err := extractor.extractPages(); err != nil {
		return fmt.Errorf("extraction failed: %v", err)
	}

	duration := time.Since(startTime)
	b.metricsCollector.RecordProcessingTime(duration)
	b.logAsync("[%d/%d] Completed %s in %.2f seconds", index, total, baseName, duration.Seconds())

	return nil
}

// processAll processes all PDF files in the directory
func (b *BatchProcessor) processAll() error {
	b.logAsync("Scanning directory: %s", b.inputDir)

	// Find all PDF files
	pdfFiles, err := b.findPDFs()
	if err != nil {
		return err
	}

	if len(pdfFiles) == 0 {
		b.logAsync("No PDF files found in %s", b.inputDir)
		return nil
	}

	b.logAsync("Found %d PDF files to process", len(pdfFiles))
	b.logAsync("Output directory: %s", b.outputDir)
	b.logAsync("Using %d processes per PDF, %d concurrent PDFs\n", b.processCount, b.concurrentPDFs)

	overallStart := time.Now()

	// Use WaitGroup for tracking completion
	var wg sync.WaitGroup
	// Mutex for protecting shared state
	var mu sync.Mutex
	successCount := 0
	failedPDFs := []string{}

	// Process PDFs concurrently with a semaphore to limit parallelism
	for i, pdfPath := range pdfFiles {
		wg.Add(1)
		b.workerPool <- struct{}{} // Acquire semaphore

		go func(path string, index int) {
			defer wg.Done()
			defer func() { <-b.workerPool }() // Release semaphore

			if err := b.processPDF(path, index+1, len(pdfFiles)); err != nil {
				b.logAsync("Failed to process %s: %v", filepath.Base(path), err)
				mu.Lock()
				failedPDFs = append(failedPDFs, filepath.Base(path))
				mu.Unlock()
			} else {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}(pdfPath, i)
	}

	// Wait for all PDFs to complete
	wg.Wait()

	overallDuration := time.Since(overallStart).Seconds()

	b.logAsync("\n" + strings.Repeat("=", 50))
	b.logAsync("BATCH PROCESSING COMPLETE")
	b.logAsync(strings.Repeat("=", 50))
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

	// Close log channel
	close(b.logChan)
	time.Sleep(100 * time.Millisecond)

	return nil
}

func main() {
	// Define command-line flags.
	pdfFileFlag := flag.String("file", "", "Path to a single PDF file")
	inputDirFlag := flag.String("dir", "", "Directory containing PDF files to process")
	outputDirFlag := flag.String("output", "", "Directory to store extracted pages")
	processCountFlag := flag.Int("processes", 0, "Number of processes to use")
	configFlag := flag.String("config", "", "Path to a YAML configuration file")
	scanDirFlag := flag.String("scan", "", "Scan directory for PDF files and copy to pdf-docs")
	copyDirFlag := flag.String("copydir", "pdf-docs", "Directory to copy PDFs to (default: pdf-docs)")
	profileFlag := flag.String("profile", "", "Enable profiling (cpu or memory)")
	cacheFlag := flag.Bool("cache", false, "Enable result caching")
	benchmarkFlag := flag.Bool("benchmark", false, "Run in benchmark mode with detailed metrics")
	helpFlag := flag.Bool("help", false, "Prints help")
	flag.Parse()

	if *helpFlag {
		fmt.Println("PDF Tool - Extract pages from PDFs or scan/copy PDFs")
		fmt.Println("\nUsage:")
		fmt.Println("  Single PDF:    -file <pdf> [-output <dir>] [-processes <n>]")
		fmt.Println("  Batch mode:    -dir <directory> [-output <dir>] [-processes <n>]")
		fmt.Println("  Scan mode:     -scan <dir> [-copydir <dir>]")
		fmt.Println("\nPerformance Options:")
		fmt.Println("  -cache         Enable result caching for repeated extractions")
		fmt.Println("  -profile cpu   Enable CPU profiling")
		fmt.Println("  -profile mem   Enable memory profiling")
		fmt.Println("  -benchmark     Run with detailed performance metrics")
		fmt.Println("\nExamples:")
		fmt.Println("  ./swiper -file document.pdf")
		fmt.Println("  ./swiper -dir /path/to/pdfs -output extracted")
		fmt.Println("  ./swiper -scan . -copydir pdf-collection")
		fmt.Println("  ./swiper -file doc.pdf -profile cpu -benchmark")
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Setup profiling if requested
	if *profileFlag == "cpu" {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal("Could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("Could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
		log.Println("CPU profiling enabled, output: cpu.prof")
	} else if *profileFlag == "mem" {
		defer func() {
			f, err := os.Create("mem.prof")
			if err != nil {
				log.Fatal("Could not create memory profile: ", err)
			}
			defer f.Close()
			runtime.GC() // Force GC before heap profile
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal("Could not write memory profile: ", err)
			}
			log.Println("Memory profile written to mem.prof")
		}()
	}

	if *benchmarkFlag {
		log.Println("Running in benchmark mode with detailed metrics")
		runtime.GOMAXPROCS(runtime.NumCPU()) // Ensure all CPUs are utilized
		log.Printf("Using %d CPU cores", runtime.NumCPU())
	}
	
	// Check if scan mode is requested
	if *scanDirFlag != "" {
		scanner, err := NewPDFScanner(*scanDirFlag, *copyDirFlag)
		if err != nil {
			log.Fatalf("Failed to initialize scanner: %v", err)
		}
		if err := scanner.scanAndCopy(); err != nil {
			log.Fatalf("Scan failed: %v", err)
		}
		return
	}
	
	// Check if batch processing mode is requested
	if *inputDirFlag != "" {
		batch, err := NewBatchProcessor(*inputDirFlag, *outputDirFlag, *processCountFlag)
		if err != nil {
			log.Fatalf("Failed to initialize batch processor: %v", err)
		}
		if err := batch.processAll(); err != nil {
			log.Fatalf("Batch processing failed: %v", err)
		}
		return
	}

	// Load configuration from YAML file if provided.
	opts := Options{
		PdfFile:      *pdfFileFlag,
		OutputDir:    *outputDirFlag,
		ProcessCount: *processCountFlag,
		ScanDir:      *scanDirFlag,
		CopyDir:      *copyDirFlag,
		Profile:      *profileFlag,
		CacheResults: *cacheFlag,
	}
	if *configFlag != "" {
		configOpts, err := loadConfig(*configFlag)
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}
		// Merge CLI options (which take precedence) with config file.
		if opts.PdfFile == "" {
			opts.PdfFile = configOpts.PdfFile
		}
		if opts.OutputDir == "" {
			opts.OutputDir = configOpts.OutputDir
		}
		if opts.ProcessCount == 0 {
			opts.ProcessCount = configOpts.ProcessCount
		}
	}

	if opts.PdfFile == "" {
		log.Fatal("Error: No input specified. Use -file for single PDF, -dir for batch processing, or -scan to copy PDFs")
	}

	extractor, err := NewExtractor(opts.PdfFile, opts.OutputDir, opts.ProcessCount)
	if err != nil {
		log.Fatalf("Failed to initialize extractor: %v", err)
	}

	if err := extractor.extractPages(); err != nil {
		log.Fatalf("Extraction failed: %v", err)
	}
}
