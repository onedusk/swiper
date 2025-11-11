package cache

import (
	"sync"
	"time"
)

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

// MetricsRecorder interface for metrics recording
type MetricsRecorder interface {
	RecordCacheHit()
	RecordCacheMiss()
}

// NewResultCache creates a new result cache
func NewResultCache() *ResultCache {
	return &ResultCache{
		cache: make(map[string]*CachedResult),
	}
}

// Get retrieves a cached result
func (c *ResultCache) Get(key string, metrics MetricsRecorder) (*CachedResult, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	result, exists := c.cache[key]
	if exists && metrics != nil {
		metrics.RecordCacheHit()
	} else if metrics != nil {
		metrics.RecordCacheMiss()
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

// Clear removes all cached results
func (c *ResultCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache = make(map[string]*CachedResult)
}

// Size returns the number of cached items
func (c *ResultCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.cache)
}