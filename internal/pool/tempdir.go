package pool

import (
	"os"
	"path/filepath"
	"sync"
)

// TempDirPool manages a pool of temporary directories
type TempDirPool struct {
	pool    chan string
	initWg  sync.WaitGroup
	closed  bool
	closeMu sync.Mutex
}

// NewTempDirPool creates a new temporary directory pool
func NewTempDirPool(size int) *TempDirPool {
	pool := &TempDirPool{
		pool: make(chan string, size),
	}
	pool.initWg.Add(1)
	go pool.init(size)
	return pool
}

// init pre-creates temporary directories for reuse
func (p *TempDirPool) init(size int) {
	defer p.initWg.Done()
	for i := 0; i < size; i++ {
		tempDir, err := os.MkdirTemp("", "pdf_images_*")
		if err != nil {
			continue
		}
		p.closeMu.Lock()
		if p.closed {
			p.closeMu.Unlock()
			os.RemoveAll(tempDir)
			return
		}
		p.closeMu.Unlock()
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
		// Guard against use after Cleanup
		p.closeMu.Lock()
		if p.closed {
			p.closeMu.Unlock()
			os.RemoveAll(dir)
			return os.MkdirTemp("", "pdf_images_*")
		}
		p.closeMu.Unlock()

		// Clear directory contents before reuse
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
	p.closeMu.Lock()
	if p.closed {
		p.closeMu.Unlock()
		os.RemoveAll(dir)
		return
	}
	p.closeMu.Unlock()
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
	p.closeMu.Lock()
	p.closed = true
	p.closeMu.Unlock()

	// Wait for init goroutine to finish
	p.initWg.Wait()

	close(p.pool)
	for dir := range p.pool {
		os.RemoveAll(dir)
	}
}
