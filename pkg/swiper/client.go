package swiper

import (
	"context"
	"fmt"

	"swiper/internal/batch"
	"swiper/internal/config"
	"swiper/internal/extractor"
	"swiper/internal/scanner"
)

// Client provides high-level PDF extraction API
type Client struct {
	config *config.Options
}

// NewClient creates a new Swiper client
func NewClient(opts ...Option) (*Client, error) {
	c := &Client{
		config: &config.Options{},
	}

	// Apply options
	for _, opt := range opts {
		opt(c)
	}

	// Set defaults
	c.config.SetDefaults()

	// Validate configuration
	if err := c.config.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

// ExtractSingle extracts a single PDF to markdown
func (c *Client) ExtractSingle(ctx context.Context, pdfPath string) (*Result, error) {
	outputDir := c.config.OutputDir

	ext, err := extractor.New(pdfPath, outputDir, c.config.ProcessCount)
	if err != nil {
		return nil, fmt.Errorf("failed to create extractor: %w", err)
	}
	defer ext.Cleanup()

	if err := ext.ExtractPages(); err != nil {
		return nil, fmt.Errorf("extraction failed: %w", err)
	}

	return &Result{
		PDFPath:   pdfPath,
		OutputDir: outputDir,
		Success:   true,
	}, nil
}

// ExtractBatch extracts multiple PDFs concurrently
func (c *Client) ExtractBatch(ctx context.Context, pdfDir string) ([]*Result, error) {
	outputDir := c.config.OutputDir

	processor, err := batch.New(pdfDir, outputDir, c.config.ProcessCount)
	if err != nil {
		return nil, fmt.Errorf("failed to create batch processor: %w", err)
	}

	if err := processor.ProcessAll(); err != nil {
		return nil, fmt.Errorf("batch processing failed: %w", err)
	}

	// For now, return a single result for the batch
	return []*Result{
		{
			PDFPath:   pdfDir,
			OutputDir: outputDir,
			Success:   true,
		},
	}, nil
}

// ScanAndCopy recursively scans for PDFs and copies them
func (c *Client) ScanAndCopy(ctx context.Context, scanDir, copyDir string) error {
	if scanDir == "" {
		scanDir = c.config.ScanDir
	}
	if copyDir == "" {
		copyDir = c.config.CopyDir
	}

	scan, err := scanner.New(scanDir, copyDir)
	if err != nil {
		return fmt.Errorf("failed to create scanner: %w", err)
	}

	if err := scan.ScanAndCopy(); err != nil {
		return fmt.Errorf("scan and copy failed: %w", err)
	}

	return nil
}