package pool

import (
	"sync/atomic"
	"testing"
)

type testBufferMetrics struct {
	hits    int64
	misses  int64
	created int64
}

func (m *testBufferMetrics) RecordBufferPoolHit()             { atomic.AddInt64(&m.hits, 1) }
func (m *testBufferMetrics) RecordBufferPoolMiss()            { atomic.AddInt64(&m.misses, 1) }
func (m *testBufferMetrics) RecordBufferPoolCreated(size int) { atomic.AddInt64(&m.created, int64(size)) }

func TestNewBufferPoolManager_NilMetrics(t *testing.T) {
	m := NewBufferPoolManager(nil)
	if m == nil {
		t.Fatal("expected non-nil manager")
	}
	// Should work without panic
	buf := m.GetBuffer(100)
	if len(buf) != 32*1024 {
		t.Fatalf("expected 32KB buffer, got %d", len(buf))
	}
	m.PutBuffer(buf)
}

func TestGetBuffer_TierSelection(t *testing.T) {
	m := NewBufferPoolManager(nil)

	cases := []struct {
		name     string
		sizeHint int64
		wantLen  int
	}{
		{"zero", 0, SmallBufferSize},
		{"1 byte", 1, SmallBufferSize},
		{"31KB", 31 * 1024, SmallBufferSize},
		{"exact 32KB", SmallBufferSize, SmallBufferSize},
		{"32KB+1", SmallBufferSize + 1, MediumBufferSize},
		{"64KB", 64 * 1024, MediumBufferSize},
		{"127KB", 127 * 1024, MediumBufferSize},
		{"exact 128KB", MediumBufferSize, MediumBufferSize},
		{"128KB+1", MediumBufferSize + 1, LargeBufferSize},
		{"200KB", 200 * 1024, LargeBufferSize},
		{"255KB", 255 * 1024, LargeBufferSize},
		{"exact 256KB", LargeBufferSize, LargeBufferSize},
		{"256KB+1", LargeBufferSize + 1, XLargeBufferSize},
		{"512KB", 512 * 1024, XLargeBufferSize},
		{"2MB", 2 * 1024 * 1024, XLargeBufferSize},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buf := m.GetBuffer(tc.sizeHint)
			if len(buf) != tc.wantLen {
				t.Fatalf("sizeHint=%d: expected len %d, got %d", tc.sizeHint, tc.wantLen, len(buf))
			}
			m.PutBuffer(buf)
		})
	}
}

func TestPutBuffer_RoundTrip(t *testing.T) {
	m := NewBufferPoolManager(nil)

	buf := m.GetBuffer(1)
	if len(buf) != 32*1024 {
		t.Fatalf("expected 32KB, got %d", len(buf))
	}
	m.PutBuffer(buf)

	buf2 := m.GetBuffer(1)
	if len(buf2) != 32*1024 {
		t.Fatalf("expected 32KB on reuse, got %d", len(buf2))
	}
	m.PutBuffer(buf2)
}

func TestPutBuffer_MismatchedSize(t *testing.T) {
	m := NewBufferPoolManager(nil)

	// Create a buffer that doesn't match any pool size
	oddBuf := make([]byte, 50000)
	// Should not panic
	m.PutBuffer(oddBuf)

	// Next get should still return correct pool size
	buf := m.GetBuffer(1)
	if len(buf) != 32*1024 {
		t.Fatalf("expected 32KB, got %d", len(buf))
	}
	m.PutBuffer(buf)
}

func TestBufferMetricsRecording(t *testing.T) {
	metrics := &testBufferMetrics{}
	m := NewBufferPoolManager(metrics)

	// First GetBuffer triggers pool.New (created) + always records hit
	buf := m.GetBuffer(1)
	m.PutBuffer(buf)

	if atomic.LoadInt64(&metrics.hits) < 1 {
		t.Fatal("expected at least 1 hit recorded")
	}
	if atomic.LoadInt64(&metrics.created) < 1 {
		t.Fatal("expected at least 1 created recorded")
	}
}
