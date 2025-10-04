package pool

import (
	"os"
	"path/filepath"
)

// TempDirPool manages a pool of temporary directories
type TempDirPool struct {
	pool chan string
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
	close(p.pool)
	for dir := range p.pool {
		os.RemoveAll(dir)
	}
}