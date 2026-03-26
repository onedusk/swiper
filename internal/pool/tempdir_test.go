package pool

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewTempDirPool_PreCreates(t *testing.T) {
	p := NewTempDirPool(3)
	defer p.Cleanup()

	// Give the init goroutine time to populate
	time.Sleep(200 * time.Millisecond)

	// Should be able to get 3 pre-created dirs without blocking
	var dirs []string
	for i := 0; i < 3; i++ {
		dir, err := p.GetTempDir()
		if err != nil {
			t.Fatalf("GetTempDir %d: %v", i, err)
		}
		dirs = append(dirs, dir)
	}

	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Fatalf("pre-created dir does not exist: %s", dir)
		}
		os.RemoveAll(dir)
	}
}

func TestGetTempDir_ValidDirectory(t *testing.T) {
	p := NewTempDirPool(1)
	defer p.Cleanup()

	time.Sleep(100 * time.Millisecond)

	dir, err := p.GetTempDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("dir does not exist: %v", err)
	}
	if !info.IsDir() {
		t.Fatalf("expected directory, got file")
	}
	os.RemoveAll(dir)
}

func TestGetTempDir_ReturnClearsContents(t *testing.T) {
	p := NewTempDirPool(1)
	defer p.Cleanup()

	time.Sleep(100 * time.Millisecond)

	dir, err := p.GetTempDir()
	if err != nil {
		t.Fatal(err)
	}

	// Create a marker file
	marker := filepath.Join(dir, "marker.txt")
	if err := os.WriteFile(marker, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	// Return and get again
	p.ReturnTempDir(dir)
	dir2, err := p.GetTempDir()
	if err != nil {
		t.Fatal(err)
	}

	// If we got the same dir back, marker should be cleared
	if dir2 == dir {
		if _, err := os.Stat(marker); !os.IsNotExist(err) {
			t.Fatal("expected marker file to be cleared on reuse")
		}
	}
	os.RemoveAll(dir2)
}

func TestGetTempDir_PoolExhaustion(t *testing.T) {
	p := NewTempDirPool(1)
	defer p.Cleanup()

	time.Sleep(100 * time.Millisecond)

	// Get the one pooled dir
	dir1, err := p.GetTempDir()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir1)

	// Pool is now empty — should create a new dir
	dir2, err := p.GetTempDir()
	if err != nil {
		t.Fatal("expected new dir creation when pool empty")
	}
	defer os.RemoveAll(dir2)

	if dir1 == dir2 {
		t.Fatal("expected different directories")
	}
}

func TestCleanup(t *testing.T) {
	p := NewTempDirPool(2)
	time.Sleep(200 * time.Millisecond)

	dir1, err := p.GetTempDir()
	if err != nil {
		t.Fatal(err)
	}
	dir2, err := p.GetTempDir()
	if err != nil {
		t.Fatal(err)
	}

	// Return both to pool
	p.ReturnTempDir(dir1)
	p.ReturnTempDir(dir2)

	// Cleanup should remove all
	p.Cleanup()

	for _, dir := range []string{dir1, dir2} {
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			t.Fatalf("expected dir removed after Cleanup: %s", dir)
		}
	}
}
