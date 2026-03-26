package pool

import (
	"sync"
)

// Buffer size constants
const (
	SmallBufferSize  = 32 * 1024
	MediumBufferSize = 128 * 1024
	LargeBufferSize  = 256 * 1024
	XLargeBufferSize = 1024 * 1024
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
					metrics.RecordBufferPoolCreated(SmallBufferSize)
				}
				return make([]byte, SmallBufferSize)
			},
		},
		mediumPool: &sync.Pool{
			New: func() interface{} {
				if metrics != nil {
					metrics.RecordBufferPoolCreated(MediumBufferSize)
				}
				return make([]byte, MediumBufferSize)
			},
		},
		largePool: &sync.Pool{
			New: func() interface{} {
				if metrics != nil {
					metrics.RecordBufferPoolCreated(LargeBufferSize)
				}
				return make([]byte, LargeBufferSize)
			},
		},
		xlargePool: &sync.Pool{
			New: func() interface{} {
				if metrics != nil {
					metrics.RecordBufferPoolCreated(XLargeBufferSize)
				}
				return make([]byte, XLargeBufferSize)
			},
		},
		metrics: metrics,
	}
}

// GetBuffer returns an appropriately sized buffer from the pool
func (m *BufferPoolManager) GetBuffer(sizeHint int64) []byte {
	var pool *sync.Pool
	if sizeHint <= SmallBufferSize {
		pool = m.smallPool
	} else if sizeHint <= MediumBufferSize {
		pool = m.mediumPool
	} else if sizeHint <= LargeBufferSize {
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
	case SmallBufferSize:
		pool = m.smallPool
	case MediumBufferSize:
		pool = m.mediumPool
	case LargeBufferSize:
		pool = m.largePool
	case XLargeBufferSize:
		pool = m.xlargePool
	default:
		// Buffer doesn't match any pool size, don't return it
		return
	}

	pool.Put(buffer)
}