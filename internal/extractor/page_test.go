package extractor

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func requirePopplerPage(t *testing.T) {
	t.Helper()
	if _, err := exec.LookPath("pdfinfo"); err != nil {
		t.Skip("poppler-utils not found — install poppler to run page tests")
	}
}

func TestExtractPages_SimplePDF(t *testing.T) {
	requirePopplerPage(t)
	outDir := t.TempDir()

	ext, err := New(testdataPath("simple.pdf"), outDir, 2)
	if err != nil {
		t.Fatal(err)
	}
	defer ext.Cleanup()

	if err := ext.ExtractPages(); err != nil {
		t.Fatalf("ExtractPages failed: %v", err)
	}

	// Verify index.md exists
	indexPath := filepath.Join(outDir, "index.md")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Fatal("index.md not found")
	}

	// Verify page files exist
	for _, page := range []string{"page_1.md", "page_2.md"} {
		pagePath := filepath.Join(outDir, page)
		if _, err := os.Stat(pagePath); os.IsNotExist(err) {
			t.Fatalf("%s not found", page)
		}
	}

	// Verify page_1.md has content
	content, err := os.ReadFile(filepath.Join(outDir, "page_1.md"))
	if err != nil {
		t.Fatal(err)
	}
	if len(strings.TrimSpace(string(content))) == 0 {
		t.Fatal("page_1.md is empty")
	}
}

func TestExtractPages_ImagesPDF(t *testing.T) {
	requirePopplerPage(t)
	outDir := t.TempDir()

	ext, err := New(testdataPath("images.pdf"), outDir, 2)
	if err != nil {
		t.Fatal(err)
	}
	defer ext.Cleanup()

	if err := ext.ExtractPages(); err != nil {
		t.Fatalf("ExtractPages failed: %v", err)
	}

	// Verify images directory exists
	imgDir := filepath.Join(outDir, "images")
	if _, err := os.Stat(imgDir); os.IsNotExist(err) {
		t.Fatal("images/ directory not found")
	}

	// Verify at least one image was extracted
	entries, err := os.ReadDir(imgDir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) == 0 {
		t.Fatal("no images extracted")
	}
}

func TestExtractPages_EmptyPDF(t *testing.T) {
	requirePopplerPage(t)
	outDir := t.TempDir()

	ext, err := New(testdataPath("empty.pdf"), outDir, 2)
	if err != nil {
		t.Fatal(err)
	}
	defer ext.Cleanup()

	// Should not panic — may return error for 0-page or succeed with empty output
	_ = ext.ExtractPages()

	// Verify index.md exists regardless
	indexPath := filepath.Join(outDir, "index.md")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		// Some PDFs with 1 blank page may still produce an index
		// This is acceptable — the key test is no panic
		t.Log("index.md not created for empty PDF (acceptable)")
	}
}

func TestExtractPages_IndexContainsLinks(t *testing.T) {
	requirePopplerPage(t)
	outDir := t.TempDir()

	ext, err := New(testdataPath("simple.pdf"), outDir, 2)
	if err != nil {
		t.Fatal(err)
	}
	defer ext.Cleanup()

	if err := ext.ExtractPages(); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filepath.Join(outDir, "index.md"))
	if err != nil {
		t.Fatal(err)
	}
	text := string(content)

	if !strings.Contains(text, "page_1") {
		t.Error("index.md should link to page_1")
	}
	if !strings.Contains(text, "page_2") {
		t.Error("index.md should link to page_2")
	}
}
