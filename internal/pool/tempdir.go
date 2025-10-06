package pool

import (
	"os"
	"path/filepath"
	"sync"
)

// TempDirPool manages a pool of temporary directories
type TempDirPool struct {
	pool      chan string
	closeOnce sync.Once
	closed    bool
	mu        sync.RWMutex
}

// NewTempDirPool creates a new temporary directory pool
func NewTempDirPool(size int) *TempDirPool {
	pool := &TempDirPool{
		pool: make(chan string, size),
	}
	go pool.init(size)
	return pool
}

// init pre-creates temporary directories for reuse
func (p *TempDirPool) init(size int) {
	for i := 0; i < size; i++ {
		tempDir, err := os.MkdirTemp("", "pdf_images_*")
		if err != nil {
			continue
		}

		// Check if pool is closed before sending
		p.mu.RLock()
		isClosed := p.closed
		p.mu.RUnlock()

		if isClosed {
			os.RemoveAll(tempDir)
			return
		}

		select {
		case p.pool <- tempDir:
		default:
			os.RemoveAll(tempDir)
		}
	}
}

// GetTempDir gets a temp directory from the pool or creates a new one
func (p *TempDirPool) GetTempDir() (string, error) {
	select {
	case dir := <-p.pool:
		// Clear the directory before reuse
		entries, _ := os.ReadDir(dir)
		for _, entry := range entries {
			os.RemoveAll(filepath.Join(dir, entry.Name()))
		}
		return dir, nil
	default:
		// Pool empty, create new one
		return os.MkdirTemp("", "pdf_images_*")
	}
}

// ReturnTempDir returns a temp directory to the pool or removes it
func (p *TempDirPool) ReturnTempDir(dir string) {
	select {
	case p.pool <- dir:
		// Successfully returned to pool
	default:
		// Pool full, remove the directory
		os.RemoveAll(dir)
	}
}

// Cleanup removes all temp directories in the pool
func (p *TempDirPool) Cleanup() {
	p.closeOnce.Do(func() {
		// Mark as closed to prevent init from sending
		p.mu.Lock()
		p.closed = true
		p.mu.Unlock()

		close(p.pool)
		for dir := range p.pool {
			os.RemoveAll(dir)
		}
	})
}