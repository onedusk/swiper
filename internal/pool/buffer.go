package pool

import (
	"sync"
)

// BufferPoolManager manages multiple buffer pools with different sizes
type BufferPoolManager struct {
	smallPool  *sync.Pool  // 32KB buffers
	mediumPool *sync.Pool  // 128KB buffers
	largePool  *sync.Pool  // 256KB buffers
	xlargePool *sync.Pool  // 1MB buffers
	metrics    MetricsRecorder
}

// MetricsRecorder interface for metrics recording
type MetricsRecorder interface {
	RecordBufferPoolHit()
	RecordBufferPoolMiss()
	RecordBufferPoolCreated(size int)
}

// NewBufferPoolManager creates a new buffer pool manager
func NewBufferPoolManager(metrics MetricsRecorder) *BufferPoolManager {
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
		metrics: metrics,
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
	if m.metrics != nil {
		m.metrics.RecordBufferPoolHit()
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