package metrics

import (
	"log"
	"strings"
	"time"
)

// PrintSummary prints a summary of metrics with context-aware labels
func (m *Collector) PrintSummary(context string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.processingTimes) == 0 {
		return
	}

	log.Println("\n" + strings.Repeat("=", 50))
	log.Println("PERFORMANCE METRICS")
	log.Println(strings.Repeat("=", 50))
	log.Printf("Pages processed: %d", m.GetPagesProcessed())
	log.Printf("Text extracted: %.2f MB", float64(m.GetTextExtracted())/(1024*1024))
	log.Printf("Images extracted: %d", m.GetImagesExtracted())
	log.Printf("Total bytes processed: %.2f MB", float64(m.GetBytesProcessed())/(1024*1024))

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
	cacheHits := m.GetCacheHits()
	cacheMisses := m.GetCacheMisses()
	totalCacheAccess := cacheHits + cacheMisses
	if totalCacheAccess > 0 {
		hitRate := float64(cacheHits) / float64(totalCacheAccess) * 100.0
		log.Printf("Cache hits: %d, Cache misses: %d", cacheHits, cacheMisses)
		log.Printf("Cache hit rate: %.2f%%", hitRate)
	}

	// Buffer pool metrics
	bpHits := m.GetBufferPoolHits()
	bpMisses := m.GetBufferPoolMisses()
	totalBPAccess := bpHits + bpMisses
	if totalBPAccess > 0 {
		bpHitRate := float64(bpHits) / float64(totalBPAccess) * 100.0
		log.Printf("Buffer pool hits: %d, misses: %d (%.2f%% hit rate)", bpHits, bpMisses, bpHitRate)
		log.Printf("Buffers created: %.2f MB", float64(m.GetBufferPoolCreated())/(1024*1024))
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