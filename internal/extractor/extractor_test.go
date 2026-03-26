package extractor

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/onedusk/swiper/internal/cache"
	"github.com/onedusk/swiper/internal/metrics"
)

func requirePoppler(t *testing.T) {
	t.Helper()
	if _, err := exec.LookPath("pdfinfo"); err != nil {
		t.Skip("poppler-utils not found — install poppler to run extractor tests")
	}
}

func testdataPath(name string) string {
	return filepath.Join("..", "..", "testdata", name)
}

func TestNew_ValidPDF(t *testing.T) {
	requirePoppler(t)
	outDir := t.TempDir()

	ext, err := New(testdataPath("simple.pdf"), outDir, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer ext.Cleanup()

	if ext == nil {
		t.Fatal("expected non-nil extractor")
	}
}

func TestNew_EmptyOutputDir(t *testing.T) {
	requirePoppler(t)

	// Empty outputDir should create a directory named after the PDF
	ext, err := New(testdataPath("simple.pdf"), "", 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer ext.Cleanup()
	defer os.RemoveAll(ext.outputDir)

	if ext.outputDir != "simple" {
		t.Fatalf("expected outputDir 'simple', got %q", ext.outputDir)
	}
}

func TestNew_NonExistentPDF(t *testing.T) {
	_, err := New("/nonexistent/path.pdf", t.TempDir(), 2)
	if err == nil {
		t.Fatal("expected error for nonexistent PDF")
	}
}

func TestNew_DefaultProcessCount(t *testing.T) {
	requirePoppler(t)

	ext, err := New(testdataPath("simple.pdf"), t.TempDir(), 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer ext.Cleanup()

	if ext.processCount <= 0 {
		t.Fatalf("expected positive processCount, got %d", ext.processCount)
	}
}

func TestWithMetrics(t *testing.T) {
	requirePoppler(t)

	m := metrics.NewCollector()
	ext, err := New(testdataPath("simple.pdf"), t.TempDir(), 2, WithMetrics(m))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer ext.Cleanup()

	if ext.metricsCollector != m {
		t.Fatal("expected custom metrics collector to be set")
	}
}

func TestWithCache(t *testing.T) {
	requirePoppler(t)

	c := cache.NewResultCache()
	ext, err := New(testdataPath("simple.pdf"), t.TempDir(), 2, WithCache(c))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer ext.Cleanup()

	if ext.resultCache != c {
		t.Fatal("expected custom cache to be set")
	}
}

func TestCleanup_Idempotent(t *testing.T) {
	requirePoppler(t)

	ext, err := New(testdataPath("simple.pdf"), t.TempDir(), 2)
	if err != nil {
		t.Fatal(err)
	}

	// Should not panic on double cleanup
	ext.Cleanup()
	ext.Cleanup()
}

func TestExtractPages_ResultsAccessor(t *testing.T) {
	requirePoppler(t)
	outDir := t.TempDir()

	ext, err := New(testdataPath("simple.pdf"), outDir, 2)
	if err != nil {
		t.Fatal(err)
	}
	defer ext.Cleanup()

	if err := ext.ExtractPages(); err != nil {
		t.Fatal(err)
	}

	results := ext.Results()
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	// Results should be sorted by page number
	for i, r := range results {
		if r.Page != i+1 {
			t.Fatalf("result[%d]: expected page %d, got %d", i, i+1, r.Page)
		}
		if !r.Success() {
			t.Fatalf("result[%d]: expected Success(), got errors: %s", i, r.ErrorSummary())
		}
		if r.Duration == 0 {
			t.Fatalf("result[%d]: expected non-zero Duration", i)
		}
	}
}

func TestGetPageCount(t *testing.T) {
	requirePoppler(t)

	ext, err := New(testdataPath("simple.pdf"), t.TempDir(), 2)
	if err != nil {
		t.Fatal(err)
	}
	defer ext.Cleanup()

	count, err := ext.getPageCount()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 2 {
		t.Fatalf("expected 2 pages, got %d", count)
	}

	// Second call should return cached value
	count2, err := ext.getPageCount()
	if err != nil {
		t.Fatal(err)
	}
	if count2 != count {
		t.Fatalf("expected cached value %d, got %d", count, count2)
	}
}
