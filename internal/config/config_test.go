package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestSetDefaults_ProcessCount(t *testing.T) {
	opts := &Options{}
	opts.SetDefaults()
	if opts.ProcessCount != runtime.NumCPU() {
		t.Fatalf("expected ProcessCount %d, got %d", runtime.NumCPU(), opts.ProcessCount)
	}
}

func TestSetDefaults_CopyDir(t *testing.T) {
	opts := &Options{ScanDir: "/some/dir"}
	opts.SetDefaults()
	if opts.CopyDir != "pdf-docs" {
		t.Fatalf("expected CopyDir 'pdf-docs', got %q", opts.CopyDir)
	}
}

func TestSetDefaults_OutputDir(t *testing.T) {
	opts := &Options{InputDir: "/some/dir"}
	opts.SetDefaults()
	if opts.OutputDir != "extracted-pdfs" {
		t.Fatalf("expected OutputDir 'extracted-pdfs', got %q", opts.OutputDir)
	}
}

func TestSetDefaults_NoSideEffects(t *testing.T) {
	opts := &Options{}
	opts.SetDefaults()
	if opts.CopyDir != "" {
		t.Fatalf("expected empty CopyDir without ScanDir, got %q", opts.CopyDir)
	}
	if opts.OutputDir != "" {
		t.Fatalf("expected empty OutputDir without InputDir, got %q", opts.OutputDir)
	}
}

func TestLoadFromFile_Valid(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yml")
	yaml := `pdf_file: test.pdf
output_dir: /tmp/out
process_count: 4
cache_results: true
`
	if err := os.WriteFile(path, []byte(yaml), 0644); err != nil {
		t.Fatal(err)
	}

	opts, err := LoadFromFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.PdfFile != "test.pdf" {
		t.Fatalf("expected PdfFile 'test.pdf', got %q", opts.PdfFile)
	}
	if opts.OutputDir != "/tmp/out" {
		t.Fatalf("expected OutputDir '/tmp/out', got %q", opts.OutputDir)
	}
	if opts.ProcessCount != 4 {
		t.Fatalf("expected ProcessCount 4, got %d", opts.ProcessCount)
	}
	if !opts.CacheResults {
		t.Fatal("expected CacheResults true")
	}
}

func TestLoadFromFile_InvalidYAML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.yml")
	if err := os.WriteFile(path, []byte(":::invalid:::yaml{{{"), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := LoadFromFile(path)
	if err == nil {
		t.Fatal("expected error for invalid YAML")
	}
}

func TestLoadFromFile_MissingFile(t *testing.T) {
	_, err := LoadFromFile("/nonexistent/path/config.yml")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestMerge_OverlayTakesPrecedence(t *testing.T) {
	base := &Options{
		PdfFile:      "base.pdf",
		OutputDir:    "/base/out",
		ProcessCount: 2,
	}
	overlay := &Options{
		PdfFile:      "overlay.pdf",
		ProcessCount: 8,
	}
	base.Merge(overlay)

	if base.PdfFile != "overlay.pdf" {
		t.Fatalf("expected PdfFile 'overlay.pdf', got %q", base.PdfFile)
	}
	if base.ProcessCount != 8 {
		t.Fatalf("expected ProcessCount 8, got %d", base.ProcessCount)
	}
	if base.OutputDir != "/base/out" {
		t.Fatalf("expected OutputDir unchanged, got %q", base.OutputDir)
	}
}

func TestMerge_ZeroValuesDoNotOverwrite(t *testing.T) {
	base := &Options{
		PdfFile:      "base.pdf",
		ProcessCount: 4,
	}
	overlay := &Options{} // all zero values
	base.Merge(overlay)

	if base.PdfFile != "base.pdf" {
		t.Fatalf("expected PdfFile unchanged, got %q", base.PdfFile)
	}
	if base.ProcessCount != 4 {
		t.Fatalf("expected ProcessCount unchanged, got %d", base.ProcessCount)
	}
}

func TestMerge_BoolField(t *testing.T) {
	base := &Options{CacheResults: false}
	overlay := &Options{CacheResults: true}
	base.Merge(overlay)

	if !base.CacheResults {
		t.Fatal("expected CacheResults true after merge")
	}
}

func TestValidate_NoInput(t *testing.T) {
	opts := Options{}
	err := opts.Validate()
	if err == nil {
		t.Fatal("expected error for empty options")
	}
	if !strings.Contains(err.Error(), "no input specified") {
		t.Fatalf("expected 'no input specified' error, got: %v", err)
	}
}

func TestValidate_ValidSingleFile(t *testing.T) {
	dir := t.TempDir()
	pdfPath := filepath.Join(dir, "test.pdf")
	if err := os.WriteFile(pdfPath, []byte("%PDF-1.4"), 0644); err != nil {
		t.Fatal(err)
	}
	opts := Options{PdfFile: pdfPath}
	if err := opts.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidate_ValidScanDir(t *testing.T) {
	dir := t.TempDir()
	opts := Options{ScanDir: dir}
	if err := opts.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidate_ValidInputDir(t *testing.T) {
	dir := t.TempDir()
	opts := Options{InputDir: dir}
	if err := opts.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidate_ConflictingModes(t *testing.T) {
	dir := t.TempDir()
	pdfPath := filepath.Join(dir, "test.pdf")
	os.WriteFile(pdfPath, []byte("%PDF-1.4"), 0644)

	opts := Options{PdfFile: pdfPath, ScanDir: dir}
	err := opts.Validate()
	if err == nil {
		t.Fatal("expected error for conflicting modes")
	}
	if !strings.Contains(err.Error(), "conflicting") {
		t.Fatalf("expected 'conflicting' error, got: %v", err)
	}
}

func TestValidate_NonExistentPDF(t *testing.T) {
	opts := Options{PdfFile: "/nonexistent/path.pdf"}
	err := opts.Validate()
	if err == nil {
		t.Fatal("expected error for nonexistent PDF")
	}
}

func TestValidate_NegativeProcessCount(t *testing.T) {
	dir := t.TempDir()
	opts := Options{ScanDir: dir, ProcessCount: -1}
	err := opts.Validate()
	if err == nil {
		t.Fatal("expected error for negative process count")
	}
}

func TestValidate_ZeroProcessCountIsValid(t *testing.T) {
	dir := t.TempDir()
	opts := Options{ScanDir: dir, ProcessCount: 0}
	if err := opts.Validate(); err != nil {
		t.Fatalf("process count 0 should be valid: %v", err)
	}
}

func TestSetDefaults_ExistingProcessCountPreserved(t *testing.T) {
	opts := &Options{ProcessCount: 16}
	opts.SetDefaults()
	if opts.ProcessCount != 16 {
		t.Fatalf("expected ProcessCount 16 preserved, got %d", opts.ProcessCount)
	}
}

func TestLoadFromFile_AllFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "full.yml")
	yaml := `pdf_file: doc.pdf
output_dir: /out
process_count: 8
scan_dir: /scan
copy_dir: /copy
profile: fast
cache_results: true
input_dir: /input
`
	if err := os.WriteFile(path, []byte(yaml), 0644); err != nil {
		t.Fatal(err)
	}

	opts, err := LoadFromFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	checks := []struct {
		name string
		got  string
		want string
	}{
		{"PdfFile", opts.PdfFile, "doc.pdf"},
		{"OutputDir", opts.OutputDir, "/out"},
		{"ScanDir", opts.ScanDir, "/scan"},
		{"CopyDir", opts.CopyDir, "/copy"},
		{"Profile", opts.Profile, "fast"},
		{"InputDir", opts.InputDir, "/input"},
	}
	for _, c := range checks {
		if !strings.EqualFold(c.got, c.want) {
			t.Errorf("%s: expected %q, got %q", c.name, c.want, c.got)
		}
	}
	if opts.ProcessCount != 8 {
		t.Errorf("ProcessCount: expected 8, got %d", opts.ProcessCount)
	}
}
