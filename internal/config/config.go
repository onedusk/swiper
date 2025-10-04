package config

import (
	"os"
	"runtime"

	"gopkg.in/yaml.v2"
)

// Options holds configuration options loaded from flags or a YAML config.
type Options struct {
	PdfFile      string `yaml:"pdf_file"`
	OutputDir    string `yaml:"output_dir"`
	ProcessCount int    `yaml:"process_count"`
	ScanDir      string `yaml:"scan_dir"`
	CopyDir      string `yaml:"copy_dir"`
	Profile      string `yaml:"profile"`      // CPU or memory profiling
	CacheResults bool   `yaml:"cache_results"` // Cache extracted text/images
	InputDir     string `yaml:"input_dir"`     // For batch processing
}

// LoadFromFile reads a YAML configuration file.
func LoadFromFile(path string) (*Options, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var opts Options
	if err := yaml.Unmarshal(data, &opts); err != nil {
		return nil, err
	}
	return &opts, nil
}

// SetDefaults sets default values for unspecified options
func (opts *Options) SetDefaults() {
	// Default process count to CPU count
	if opts.ProcessCount <= 0 {
		if n := runtime.NumCPU(); n > 0 {
			opts.ProcessCount = n
		} else {
			opts.ProcessCount = 4
		}
	}

	// Default copy directory for scanning
	if opts.CopyDir == "" && opts.ScanDir != "" {
		opts.CopyDir = "pdf-docs"
	}

	// Default output directory for batch processing
	if opts.OutputDir == "" && opts.InputDir != "" {
		opts.OutputDir = "extracted-pdfs"
	}
}

// Validate checks if the configuration is valid
func (opts *Options) Validate() error {
	// Add validation logic as needed
	return nil
}

// Merge merges two Options structs, with values from 'other' taking precedence
func (opts *Options) Merge(other *Options) {
	if other.PdfFile != "" {
		opts.PdfFile = other.PdfFile
	}
	if other.OutputDir != "" {
		opts.OutputDir = other.OutputDir
	}
	if other.ProcessCount > 0 {
		opts.ProcessCount = other.ProcessCount
	}
	if other.ScanDir != "" {
		opts.ScanDir = other.ScanDir
	}
	if other.CopyDir != "" {
		opts.CopyDir = other.CopyDir
	}
	if other.InputDir != "" {
		opts.InputDir = other.InputDir
	}
	if other.Profile != "" {
		opts.Profile = other.Profile
	}
	if other.CacheResults {
		opts.CacheResults = other.CacheResults
	}
}