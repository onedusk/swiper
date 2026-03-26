package metrics

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func captureLog(fn func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(log.LstdFlags)
	}()
	fn()
	return buf.String()
}

func TestPrintSummary_EmptyCollector(t *testing.T) {
	c := NewCollector()
	output := captureLog(func() {
		c.PrintSummary("single")
	})
	// No processingTimes → returns early with no output
	if output != "" {
		t.Fatalf("expected no output for empty collector, got %q", output)
	}
}

func TestPrintSummary_SingleContext(t *testing.T) {
	c := NewCollector()
	c.RecordPageProcessed()
	c.RecordPageProcessed()
	c.RecordTextExtracted(1024)
	c.RecordImagesExtracted(3)
	c.RecordProcessingTime(100 * time.Millisecond)

	output := captureLog(func() {
		c.PrintSummary("single")
	})

	checks := []string{
		"PERFORMANCE METRICS",
		"Pages processed: 2",
		"Images extracted: 3",
		"Average processing time per page:",
	}
	for _, want := range checks {
		if !strings.Contains(output, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
	if strings.Contains(output, "per PDF") {
		t.Error("single context should not contain 'per PDF'")
	}
}

func TestPrintSummary_BatchContext(t *testing.T) {
	c := NewCollector()
	c.RecordProcessingTime(200 * time.Millisecond)

	output := captureLog(func() {
		c.PrintSummary("batch")
	})

	if !strings.Contains(output, "Average processing time per PDF:") {
		t.Error("batch context should contain 'per PDF'")
	}
}

func TestPrintSummary_WithCacheMetrics(t *testing.T) {
	c := NewCollector()
	c.RecordProcessingTime(50 * time.Millisecond)
	c.RecordCacheHit()
	c.RecordCacheHit()
	c.RecordCacheMiss()

	output := captureLog(func() {
		c.PrintSummary("single")
	})

	if !strings.Contains(output, "Cache hits: 2") {
		t.Error("expected cache hits in output")
	}
	if !strings.Contains(output, "Cache hit rate:") {
		t.Error("expected cache hit rate in output")
	}
}

func TestPrintSummary_WithWorkerUtilization(t *testing.T) {
	c := NewCollector()
	c.RecordProcessingTime(50 * time.Millisecond)
	c.RecordWorkerTime(0, 100*time.Millisecond)
	c.RecordWorkerTime(1, 200*time.Millisecond)

	output := captureLog(func() {
		c.PrintSummary("single")
	})

	if !strings.Contains(output, "Average worker utilization:") {
		t.Error("expected worker utilization in output")
	}
}

func TestPrintSummary_WithQueueDepth(t *testing.T) {
	c := NewCollector()
	c.RecordProcessingTime(50 * time.Millisecond)
	c.RecordPageQueueDepth(5)
	c.RecordPageQueueDepth(10)

	output := captureLog(func() {
		c.PrintSummary("single")
	})

	if !strings.Contains(output, "Page queue") {
		t.Error("expected page queue info in output")
	}
}
