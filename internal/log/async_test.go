package alog

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
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

func TestMessageDelivery(t *testing.T) {
	output := captureLog(func() {
		l := New(100, false)
		for i := 0; i < 10; i++ {
			l.Log("msg %d", i)
		}
		l.Close()
	})

	for i := 0; i < 10; i++ {
		want := fmt.Sprintf("msg %d", i)
		if !strings.Contains(output, want) {
			t.Fatalf("expected %q in output", want)
		}
	}
}

func TestQuietMode(t *testing.T) {
	output := captureLog(func() {
		l := New(100, true)
		l.Log("should not appear")
		l.Log("also hidden")
		l.Close()
	})

	if output != "" {
		t.Fatalf("expected no output in quiet mode, got %q", output)
	}
}

func TestBackpressure(t *testing.T) {
	// Buffer size 1 — second message should fallback to sync
	output := captureLog(func() {
		l := New(1, false)
		// Fill the channel
		l.Log("first")
		// This may or may not block depending on timing, but should not deadlock
		l.Log("second")
		l.Close()
	})

	if !strings.Contains(output, "first") {
		t.Fatal("expected 'first' in output")
	}
	if !strings.Contains(output, "second") {
		t.Fatal("expected 'second' in output (sync fallback)")
	}
}

func TestCloseDrainsAll(t *testing.T) {
	output := captureLog(func() {
		l := New(200, false)
		for i := 0; i < 100; i++ {
			l.Log("drain-%d", i)
		}
		l.Close()
	})

	count := strings.Count(output, "drain-")
	if count != 100 {
		t.Fatalf("expected 100 messages drained, got %d", count)
	}
}

func TestConcurrentLog(t *testing.T) {
	l := New(200, true) // quiet to avoid log contention
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			l.Log("goroutine %d", id)
		}(i)
	}
	wg.Wait()
	l.Close()
}

func TestCloseIdempotency(t *testing.T) {
	l := New(10, false)
	l.Log("test")

	// Should not panic — internal sync.Once guard
	l.Close()
	l.Close()
}
