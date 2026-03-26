package cache

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// testCacheMetrics is a minimal MetricsRecorder implementation for testing.
type testCacheMetrics struct {
	hits   int64
	misses int64
}

func (m *testCacheMetrics) RecordCacheHit()  { atomic.AddInt64(&m.hits, 1) }
func (m *testCacheMetrics) RecordCacheMiss() { atomic.AddInt64(&m.misses, 1) }

func TestNewResultCache(t *testing.T) {
	c := NewResultCache()
	if c == nil {
		t.Fatal("NewResultCache returned nil")
	}
	if c.Size() != 0 {
		t.Fatalf("expected size 0, got %d", c.Size())
	}
}

func TestSetAndGet(t *testing.T) {
	c := NewResultCache()
	m := &testCacheMetrics{}

	c.Set("key1", &CachedResult{
		Text:   "hello",
		Images: []string{"img1.jpg", "img2.png"},
	})

	result, ok := c.Get("key1", m)
	if !ok {
		t.Fatal("expected cache hit")
	}
	if result.Text != "hello" {
		t.Fatalf("expected text 'hello', got %q", result.Text)
	}
	if len(result.Images) != 2 {
		t.Fatalf("expected 2 images, got %d", len(result.Images))
	}
	if result.Images[0] != "img1.jpg" || result.Images[1] != "img2.png" {
		t.Fatalf("unexpected images: %v", result.Images)
	}
	if result.Timestamp.IsZero() {
		t.Fatal("expected non-zero timestamp")
	}
}

func TestGetMiss(t *testing.T) {
	c := NewResultCache()
	m := &testCacheMetrics{}

	result, ok := c.Get("nonexistent", m)
	if ok {
		t.Fatal("expected cache miss")
	}
	if result != nil {
		t.Fatal("expected nil result on miss")
	}
	if atomic.LoadInt64(&m.misses) != 1 {
		t.Fatalf("expected 1 miss, got %d", atomic.LoadInt64(&m.misses))
	}
}

func TestClear(t *testing.T) {
	c := NewResultCache()
	m := &testCacheMetrics{}

	for i := 0; i < 3; i++ {
		c.Set(fmt.Sprintf("key%d", i), &CachedResult{Text: "data"})
	}
	if c.Size() != 3 {
		t.Fatalf("expected size 3, got %d", c.Size())
	}

	c.Clear()
	if c.Size() != 0 {
		t.Fatalf("expected size 0 after clear, got %d", c.Size())
	}

	_, ok := c.Get("key0", m)
	if ok {
		t.Fatal("expected miss after clear")
	}
}

func TestSize(t *testing.T) {
	c := NewResultCache()

	for i := 0; i < 5; i++ {
		c.Set(fmt.Sprintf("key%d", i), &CachedResult{Text: "data"})
	}
	if c.Size() != 5 {
		t.Fatalf("expected size 5, got %d", c.Size())
	}

	// Overwrite existing key — size should stay 5
	c.Set("key0", &CachedResult{Text: "updated"})
	if c.Size() != 5 {
		t.Fatalf("expected size 5 after overwrite, got %d", c.Size())
	}
}

func TestMetricsRecording(t *testing.T) {
	c := NewResultCache()
	m := &testCacheMetrics{}

	// Miss
	c.Get("miss", m)
	if atomic.LoadInt64(&m.misses) != 1 {
		t.Fatalf("expected 1 miss, got %d", atomic.LoadInt64(&m.misses))
	}
	if atomic.LoadInt64(&m.hits) != 0 {
		t.Fatalf("expected 0 hits, got %d", atomic.LoadInt64(&m.hits))
	}

	// Hit
	c.Set("hit", &CachedResult{Text: "data"})
	c.Get("hit", m)
	if atomic.LoadInt64(&m.hits) != 1 {
		t.Fatalf("expected 1 hit, got %d", atomic.LoadInt64(&m.hits))
	}
}

func TestGetWithNilMetrics(t *testing.T) {
	c := NewResultCache()
	c.Set("key", &CachedResult{Text: "data"})

	// Should not panic
	result, ok := c.Get("key", nil)
	if !ok {
		t.Fatal("expected hit")
	}
	if result.Text != "data" {
		t.Fatalf("expected 'data', got %q", result.Text)
	}
}

func TestConcurrentAccess(t *testing.T) {
	c := NewResultCache()
	m := &testCacheMetrics{}

	var wg sync.WaitGroup
	// 50 writers
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			c.Set(fmt.Sprintf("key%d", id), &CachedResult{Text: fmt.Sprintf("val%d", id)})
		}(i)
	}
	// 50 readers
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			c.Get(fmt.Sprintf("key%d", id%50), m)
		}(i)
	}
	wg.Wait()

	if c.Size() != 50 {
		t.Fatalf("expected size 50, got %d", c.Size())
	}
}
