package swiper

import (
	"time"

	"github.com/onedusk/swiper/internal/config"
)

// PageSummary represents the extraction outcome for a single page.
type PageSummary struct {
	Page       int
	HasText    bool
	ImageCount int
	Errors     []string
	Duration   time.Duration
}

// Result represents the result of a PDF extraction
type Result struct {
	PDFPath     string
	OutputDir   string
	PageCount   int
	Success     bool
	Error       error
	PageResults []PageSummary
	Duration    time.Duration
}

// Option configures the Client
type Option func(*Client)

// WithProcessCount sets the number of worker processes
func WithProcessCount(n int) Option {
	return func(c *Client) {
		c.config.ProcessCount = n
	}
}

// WithOutputDir sets the output directory
func WithOutputDir(dir string) Option {
	return func(c *Client) {
		c.config.OutputDir = dir
	}
}

// WithCache enables result caching
func WithCache(enabled bool) Option {
	return func(c *Client) {
		c.config.CacheResults = enabled
	}
}

// WithConfig sets the full configuration
func WithConfig(cfg *config.Options) Option {
	return func(c *Client) {
		c.config = cfg
	}
}