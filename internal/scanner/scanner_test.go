package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

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

func TestNew_ValidDirs(t *testing.T) {
	scanDir := t.TempDir()
	copyDir := t.TempDir()

	s, err := New(scanDir, copyDir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil scanner")
	}
}

func TestNew_CreatesCopyDir(t *testing.T) {
	scanDir := t.TempDir()
	copyDir := filepath.Join(t.TempDir(), "newsubdir", "copies")

	_, err := New(scanDir, copyDir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := os.Stat(copyDir); os.IsNotExist(err) {
		t.Fatal("expected copyDir to be created")
	}
}

func TestFindPDFs(t *testing.T) {
	scanDir := t.TempDir()
	copyFile(t, testdataPath("simple.pdf"), filepath.Join(scanDir, "doc.pdf"))
	os.WriteFile(filepath.Join(scanDir, "readme.txt"), []byte("not a pdf"), 0644)

	s, err := New(scanDir, t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

	pdfs, err := s.FindPDFs()
	if err != nil {
		t.Fatalf("FindPDFs error: %v", err)
	}
	if len(pdfs) != 1 {
		t.Fatalf("expected 1 PDF, got %d", len(pdfs))
	}
}

func TestFindPDFs_Recursive(t *testing.T) {
	scanDir := t.TempDir()
	subDir := filepath.Join(scanDir, "subdir")
	os.MkdirAll(subDir, 0755)

	copyFile(t, testdataPath("simple.pdf"), filepath.Join(scanDir, "root.pdf"))
	copyFile(t, testdataPath("simple.pdf"), filepath.Join(subDir, "nested.pdf"))

	s, err := New(scanDir, t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

	pdfs, err := s.FindPDFs()
	if err != nil {
		t.Fatal(err)
	}
	if len(pdfs) != 2 {
		t.Fatalf("expected 2 PDFs, got %d", len(pdfs))
	}
}

func TestFindPDFs_NoPDFs(t *testing.T) {
	scanDir := t.TempDir()

	s, err := New(scanDir, t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

	pdfs, err := s.FindPDFs()
	if err != nil {
		t.Fatal(err)
	}
	if len(pdfs) != 0 {
		t.Fatalf("expected 0 PDFs, got %d", len(pdfs))
	}
}

func TestScanAndCopy(t *testing.T) {
	scanDir := t.TempDir()
	copyDir := t.TempDir()

	copyFile(t, testdataPath("simple.pdf"), filepath.Join(scanDir, "doc.pdf"))

	s, err := New(scanDir, copyDir)
	if err != nil {
		t.Fatal(err)
	}

	if err := s.ScanAndCopy(); err != nil {
		t.Fatalf("ScanAndCopy error: %v", err)
	}

	entries, err := os.ReadDir(copyDir)
	if err != nil {
		t.Fatal(err)
	}

	pdfCount := 0
	for _, e := range entries {
		if filepath.Ext(e.Name()) == ".pdf" {
			pdfCount++
		}
	}
	if pdfCount < 1 {
		t.Fatal("expected at least 1 PDF copied")
	}
}

func TestScanAndCopy_CollisionHandling(t *testing.T) {
	scanDir := t.TempDir()
	copyDir := t.TempDir()

	// Pre-place a file in copyDir with same name
	copyFile(t, testdataPath("simple.pdf"), filepath.Join(scanDir, "doc.pdf"))
	copyFile(t, testdataPath("simple.pdf"), filepath.Join(copyDir, "doc.pdf"))

	s, err := New(scanDir, copyDir)
	if err != nil {
		t.Fatal(err)
	}

	if err := s.ScanAndCopy(); err != nil {
		t.Fatalf("ScanAndCopy error: %v", err)
	}

	entries, err := os.ReadDir(copyDir)
	if err != nil {
		t.Fatal(err)
	}

	pdfCount := 0
	for _, e := range entries {
		if filepath.Ext(e.Name()) == ".pdf" {
			pdfCount++
		}
	}
	if pdfCount < 2 {
		t.Fatalf("expected at least 2 PDFs after collision, got %d", pdfCount)
	}
}
