package batch

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func requirePoppler(t *testing.T) {
	t.Helper()
	if _, err := exec.LookPath("pdfinfo"); err != nil {
		t.Skip("poppler-utils not found — install poppler to run batch tests")
	}
}

func testdataPath(name string) string {
	return filepath.Join("..", "..", "testdata", name)
}

func copyFile(t *testing.T, src, dst string) {
	t.Helper()
	data, err := os.ReadFile(src)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(dst, data, 0644); err != nil {
		t.Fatal(err)
	}
}

func TestNew_ValidInputDir(t *testing.T) {
	inputDir := t.TempDir()
	outputDir := t.TempDir()

	p, err := New(inputDir, outputDir, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p == nil {
		t.Fatal("expected non-nil processor")
	}
}

func TestNew_EmptyInputDir(t *testing.T) {
	_, err := New("", t.TempDir(), 2)
	if err == nil {
		t.Fatal("expected error for empty input dir")
	}
	if !strings.Contains(err.Error(), "required") {
		t.Fatalf("expected 'required' in error, got %q", err.Error())
	}
}

func TestNew_NonExistentInputDir(t *testing.T) {
	_, err := New("/nonexistent/path/to/dir", t.TempDir(), 2)
	if err == nil {
		t.Fatal("expected error for nonexistent input dir")
	}
	if !strings.Contains(err.Error(), "not found") {
		t.Fatalf("expected 'not found' in error, got %q", err.Error())
	}
}

func TestFindPDFs(t *testing.T) {
	inputDir := t.TempDir()
	copyFile(t, testdataPath("simple.pdf"), filepath.Join(inputDir, "doc.pdf"))
	os.WriteFile(filepath.Join(inputDir, "readme.txt"), []byte("not a pdf"), 0644)

	p, err := New(inputDir, t.TempDir(), 2)
	if err != nil {
		t.Fatal(err)
	}

	pdfs, err := p.FindPDFs()
	if err != nil {
		t.Fatalf("FindPDFs error: %v", err)
	}
	if len(pdfs) != 1 {
		t.Fatalf("expected 1 PDF, got %d", len(pdfs))
	}
}

func TestProcessAll_TwoPDFs(t *testing.T) {
	requirePoppler(t)

	inputDir := t.TempDir()
	outputDir := t.TempDir()

	copyFile(t, testdataPath("simple.pdf"), filepath.Join(inputDir, "doc1.pdf"))
	copyFile(t, testdataPath("simple.pdf"), filepath.Join(inputDir, "doc2.pdf"))

	p, err := New(inputDir, outputDir, 2)
	if err != nil {
		t.Fatal(err)
	}

	if err := p.ProcessAll(); err != nil {
		t.Fatalf("ProcessAll error: %v", err)
	}

	// Verify two output subdirectories
	entries, err := os.ReadDir(outputDir)
	if err != nil {
		t.Fatal(err)
	}

	dirCount := 0
	for _, e := range entries {
		if e.IsDir() {
			dirCount++
			// Each should have an index.md
			indexPath := filepath.Join(outputDir, e.Name(), "index.md")
			if _, err := os.Stat(indexPath); os.IsNotExist(err) {
				t.Fatalf("expected index.md in %s", e.Name())
			}
		}
	}
	if dirCount < 2 {
		t.Fatalf("expected at least 2 output directories, got %d", dirCount)
	}
}

func TestProcessAll_EmptyDirectory(t *testing.T) {
	inputDir := t.TempDir()
	outputDir := t.TempDir()

	p, err := New(inputDir, outputDir, 2)
	if err != nil {
		t.Fatal(err)
	}

	// Should not error on empty directory
	if err := p.ProcessAll(); err != nil {
		t.Fatalf("expected no error for empty dir, got %v", err)
	}
}
