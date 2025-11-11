package metrics

import (
	"sync"
	"sync/atomic"
	"time"
)

// Collector collects performance metrics
type Collector struct {
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

// NewCollector creates a new metrics collector
func NewCollector() *Collector {
	return &Collector{
		processingTimes:   make([]time.Duration, 0),
		pdfSizes:          make(map[string]int64),
		pageQueueDepth:    make([]int, 0),
		workerUtilization: make(map[int]time.Duration),
	}
}

// RecordPageProcessed records a processed page
func (m *Collector) RecordPageProcessed() {
	atomic.AddInt64(&m.pagesProcessed, 1)
}

// RecordTextExtracted records text extraction
func (m *Collector) RecordTextExtracted(bytes int) {
	atomic.AddInt64(&m.textExtracted, int64(bytes))
}

// RecordImagesExtracted records image extraction
func (m *Collector) RecordImagesExtracted(count int) {
	atomic.AddInt64(&m.imagesExtracted, int64(count))
}

// RecordBytesProcessed records bytes processed for I/O operations
func (m *Collector) RecordBytesProcessed(bytes int64) {
	atomic.AddInt64(&m.bytesProcessed, bytes)
}

// RecordProcessingTime records processing time
func (m *Collector) RecordProcessingTime(d time.Duration) {
	m.mu.Lock()
	m.processingTimes = append(m.processingTimes, d)
	m.mu.Unlock()
}

// RecordCacheHit records a cache hit
func (m *Collector) RecordCacheHit() {
	atomic.AddInt64(&m.cacheHits, 1)
}

// RecordCacheMiss records a cache miss
func (m *Collector) RecordCacheMiss() {
	atomic.AddInt64(&m.cacheMisses, 1)
}

// RecordBufferPoolHit records a buffer pool hit
func (m *Collector) RecordBufferPoolHit() {
	atomic.AddInt64(&m.bufferPoolHits, 1)
}

// RecordBufferPoolMiss records a buffer pool miss
func (m *Collector) RecordBufferPoolMiss() {
	atomic.AddInt64(&m.bufferPoolMisses, 1)
}

// RecordBufferPoolCreated records buffer creation
func (m *Collector) RecordBufferPoolCreated(size int) {
	atomic.AddInt64(&m.bufferPoolCreated, int64(size))
}

// RecordPageQueueDepth records the current page queue depth
func (m *Collector) RecordPageQueueDepth(depth int) {
	m.mu.Lock()
	m.pageQueueDepth = append(m.pageQueueDepth, depth)
	m.mu.Unlock()
}

// RecordWorkerTime records time spent by a worker
func (m *Collector) RecordWorkerTime(workerID int, duration time.Duration) {
	m.mu.Lock()
	m.workerUtilization[workerID] += duration
	m.mu.Unlock()
}

// GetCacheHitRate returns the cache hit rate as a percentage
func (m *Collector) GetCacheHitRate() float64 {
	hits := atomic.LoadInt64(&m.cacheHits)
	misses := atomic.LoadInt64(&m.cacheMisses)
	total := hits + misses
	if total == 0 {
		return 0.0
	}
	return float64(hits) / float64(total) * 100.0
}

// GetPagesProcessed returns the total pages processed
func (m *Collector) GetPagesProcessed() int64 {
	return atomic.LoadInt64(&m.pagesProcessed)
}

// GetTextExtracted returns the total text extracted in bytes
func (m *Collector) GetTextExtracted() int64 {
	return atomic.LoadInt64(&m.textExtracted)
}

// GetImagesExtracted returns the total images extracted
func (m *Collector) GetImagesExtracted() int64 {
	return atomic.LoadInt64(&m.imagesExtracted)
}

// GetBytesProcessed returns the total bytes processed
func (m *Collector) GetBytesProcessed() int64 {
	return atomic.LoadInt64(&m.bytesProcessed)
}

// GetCacheHits returns the total cache hits
func (m *Collector) GetCacheHits() int64 {
	return atomic.LoadInt64(&m.cacheHits)
}

// GetCacheMisses returns the total cache misses
func (m *Collector) GetCacheMisses() int64 {
	return atomic.LoadInt64(&m.cacheMisses)
}

// GetBufferPoolHits returns the total buffer pool hits
func (m *Collector) GetBufferPoolHits() int64 {
	return atomic.LoadInt64(&m.bufferPoolHits)
}

// GetBufferPoolMisses returns the total buffer pool misses
func (m *Collector) GetBufferPoolMisses() int64 {
	return atomic.LoadInt64(&m.bufferPoolMisses)
}

// GetBufferPoolCreated returns the total buffer memory created
func (m *Collector) GetBufferPoolCreated() int64 {
	return atomic.LoadInt64(&m.bufferPoolCreated)
}

// GetProcessingTimes returns a copy of processing times
func (m *Collector) GetProcessingTimes() []time.Duration {
	m.mu.Lock()
	defer m.mu.Unlock()

	times := make([]time.Duration, len(m.processingTimes))
	copy(times, m.processingTimes)
	return times
}

// GetPageQueueDepths returns a copy of page queue depths
func (m *Collector) GetPageQueueDepths() []int {
	m.mu.Lock()
	defer m.mu.Unlock()

	depths := make([]int, len(m.pageQueueDepth))
	copy(depths, m.pageQueueDepth)
	return depths
}

// GetWorkerUtilization returns a copy of worker utilization map
func (m *Collector) GetWorkerUtilization() map[int]time.Duration {
	m.mu.Lock()
	defer m.mu.Unlock()

	util := make(map[int]time.Duration)
	for k, v := range m.workerUtilization {
		util[k] = v
	}
	return util
}

// RecordPDFSize records the size of a PDF file
func (m *Collector) RecordPDFSize(name string, size int64) {
	m.mu.Lock()
	m.pdfSizes[name] = size
	m.mu.Unlock()
}

// GetPDFSizes returns a copy of PDF sizes map
func (m *Collector) GetPDFSizes() map[string]int64 {
	m.mu.Lock()
	defer m.mu.Unlock()

	sizes := make(map[string]int64)
	for k, v := range m.pdfSizes {
		sizes[k] = v
	}
	return sizes
}