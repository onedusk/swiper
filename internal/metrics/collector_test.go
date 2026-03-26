package metrics

import (
	"math"
	"sync"
	"testing"
	"time"
)

func TestNewCollector(t *testing.T) {
	c := NewCollector()
	if c == nil {
		t.Fatal("expected non-nil collector")
	}
	if c.GetPagesProcessed() != 0 {
		t.Fatalf("expected 0 pages, got %d", c.GetPagesProcessed())
	}
	if c.GetTextExtracted() != 0 {
		t.Fatalf("expected 0 text, got %d", c.GetTextExtracted())
	}
	if c.GetImagesExtracted() != 0 {
		t.Fatalf("expected 0 images, got %d", c.GetImagesExtracted())
	}
	if c.GetBytesProcessed() != 0 {
		t.Fatalf("expected 0 bytes, got %d", c.GetBytesProcessed())
	}
}

func TestRecordPageProcessed(t *testing.T) {
	c := NewCollector()
	c.RecordPageProcessed()
	c.RecordPageProcessed()
	c.RecordPageProcessed()
	if c.GetPagesProcessed() != 3 {
		t.Fatalf("expected 3, got %d", c.GetPagesProcessed())
	}
}

func TestRecordTextExtracted(t *testing.T) {
	c := NewCollector()
	c.RecordTextExtracted(100)
	c.RecordTextExtracted(200)
	if c.GetTextExtracted() != 300 {
		t.Fatalf("expected 300, got %d", c.GetTextExtracted())
	}
}

func TestRecordImagesExtracted(t *testing.T) {
	c := NewCollector()
	c.RecordImagesExtracted(5)
	if c.GetImagesExtracted() != 5 {
		t.Fatalf("expected 5, got %d", c.GetImagesExtracted())
	}
}

func TestRecordBytesProcessed(t *testing.T) {
	c := NewCollector()
	c.RecordBytesProcessed(1024)
	if c.GetBytesProcessed() != 1024 {
		t.Fatalf("expected 1024, got %d", c.GetBytesProcessed())
	}
}

func TestRecordCacheHitMiss(t *testing.T) {
	c := NewCollector()
	c.RecordCacheHit()
	c.RecordCacheHit()
	c.RecordCacheMiss()
	if c.GetCacheHits() != 2 {
		t.Fatalf("expected 2 hits, got %d", c.GetCacheHits())
	}
	if c.GetCacheMisses() != 1 {
		t.Fatalf("expected 1 miss, got %d", c.GetCacheMisses())
	}
}

func TestGetCacheHitRate(t *testing.T) {
	cases := []struct {
		name   string
		hits   int
		misses int
		want   float64
	}{
		{"no data", 0, 0, 0.0},
		{"all hits", 1, 0, 100.0},
		{"half", 1, 1, 50.0},
		{"75 percent", 3, 1, 75.0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCollector()
			for i := 0; i < tc.hits; i++ {
				c.RecordCacheHit()
			}
			for i := 0; i < tc.misses; i++ {
				c.RecordCacheMiss()
			}
			got := c.GetCacheHitRate()
			if math.Abs(got-tc.want) > 0.01 {
				t.Fatalf("expected %.2f, got %.2f", tc.want, got)
			}
		})
	}
}

func TestRecordBufferPool(t *testing.T) {
	c := NewCollector()
	c.RecordBufferPoolHit()
	c.RecordBufferPoolHit()
	c.RecordBufferPoolHit()
	c.RecordBufferPoolMiss()
	c.RecordBufferPoolCreated(32768)

	if c.GetBufferPoolHits() != 3 {
		t.Fatalf("expected 3 hits, got %d", c.GetBufferPoolHits())
	}
	if c.GetBufferPoolMisses() != 1 {
		t.Fatalf("expected 1 miss, got %d", c.GetBufferPoolMisses())
	}
	if c.GetBufferPoolCreated() != 32768 {
		t.Fatalf("expected 32768 created bytes, got %d", c.GetBufferPoolCreated())
	}
}

func TestRecordProcessingTime(t *testing.T) {
	c := NewCollector()
	durations := []time.Duration{100 * time.Millisecond, 200 * time.Millisecond, 300 * time.Millisecond}
	for _, d := range durations {
		c.RecordProcessingTime(d)
	}
	times := c.GetProcessingTimes()
	if len(times) != 3 {
		t.Fatalf("expected 3 times, got %d", len(times))
	}
	for i, d := range durations {
		if times[i] != d {
			t.Fatalf("time[%d]: expected %v, got %v", i, d, times[i])
		}
	}
}

func TestRecordWorkerTime(t *testing.T) {
	c := NewCollector()
	c.RecordWorkerTime(0, 100*time.Millisecond)
	c.RecordWorkerTime(1, 200*time.Millisecond)
	c.RecordWorkerTime(0, 50*time.Millisecond)

	util := c.GetWorkerUtilization()
	if len(util) != 2 {
		t.Fatalf("expected 2 workers, got %d", len(util))
	}
	if util[0] != 150*time.Millisecond {
		t.Fatalf("worker 0: expected 150ms, got %v", util[0])
	}
	if util[1] != 200*time.Millisecond {
		t.Fatalf("worker 1: expected 200ms, got %v", util[1])
	}
}

func TestRecordPDFSize(t *testing.T) {
	c := NewCollector()
	c.RecordPDFSize("doc1.pdf", 1024)
	c.RecordPDFSize("doc2.pdf", 2048)

	sizes := c.GetPDFSizes()
	if len(sizes) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(sizes))
	}
	if sizes["doc1.pdf"] != 1024 {
		t.Fatalf("doc1: expected 1024, got %d", sizes["doc1.pdf"])
	}
	if sizes["doc2.pdf"] != 2048 {
		t.Fatalf("doc2: expected 2048, got %d", sizes["doc2.pdf"])
	}
}

func TestRecordPageQueueDepth(t *testing.T) {
	c := NewCollector()
	c.RecordPageQueueDepth(5)
	c.RecordPageQueueDepth(10)
	c.RecordPageQueueDepth(3)

	depths := c.GetPageQueueDepths()
	if len(depths) != 3 {
		t.Fatalf("expected 3 depths, got %d", len(depths))
	}
	expected := []int{5, 10, 3}
	for i, want := range expected {
		if depths[i] != want {
			t.Fatalf("depth[%d]: expected %d, got %d", i, want, depths[i])
		}
	}
}

func TestConcurrentAccess(t *testing.T) {
	c := NewCollector()
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.RecordPageProcessed()
			c.RecordTextExtracted(10)
			c.RecordImagesExtracted(1)
			c.RecordBytesProcessed(100)
			c.RecordCacheHit()
			c.RecordCacheMiss()
			c.RecordBufferPoolHit()
			c.RecordProcessingTime(1 * time.Millisecond)
		}()
	}
	wg.Wait()

	if c.GetPagesProcessed() != 50 {
		t.Fatalf("expected 50 pages, got %d", c.GetPagesProcessed())
	}
	if c.GetTextExtracted() != 500 {
		t.Fatalf("expected 500 text bytes, got %d", c.GetTextExtracted())
	}
	if c.GetCacheHits() != 50 {
		t.Fatalf("expected 50 cache hits, got %d", c.GetCacheHits())
	}
}
