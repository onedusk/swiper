# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

Swiper is a high-performance PDF extraction tool written in Go that converts PDF documents into markdown files with embedded images. It supports parallel processing for efficient extraction of text and images from PDF files, with capabilities for single PDF processing, batch processing of multiple PDFs, and PDF discovery/collection.

## Key Commands

### Build & Run
```bash
# Build the binary
go build -o swiper ./cmd/swiper/

# Run with a single PDF
./swiper -file document.pdf [-output output_dir] [-processes 8]

# Batch process multiple PDFs
./swiper -dir /path/to/pdfs [-output extracted-pdfs] [-processes 8]

# Scan and copy PDFs to collection
./swiper -scan /path/to/scan [-copydir pdf-collection]

# Run tests
go test ./...

# Format code
go fmt ./...
```

### Development Dependencies
```bash
# Install Go module dependencies
go mod download

# Update dependencies
go mod tidy

# Install required system tools (macOS/Linux)
# Required: pdfinfo, pdftotext, pdfimages (from poppler-utils)
brew install poppler  # macOS
apt-get install poppler-utils  # Ubuntu/Debian
```

## Architecture Overview

### Core Components

1. **Extractor** (`internal/extractor/`): Main PDF processing engine
   - `extractor.go` — manages concurrent page extraction with worker pools, async logging, temp directory pooling
   - `page.go` — per-page text extraction and image processing

2. **PDFScanner** (`internal/scanner/scanner.go`): PDF discovery and collection
   - Recursively scans directories for PDF files
   - Copies PDFs with collision handling (timestamps)
   - Progress tracking with file size statistics

3. **BatchProcessor** (`internal/batch/processor.go`): Batch PDF processing
   - Processes multiple PDFs sequentially
   - Each PDF internally uses parallel page processing
   - Comprehensive error tracking and reporting

4. **Supporting packages:**
   - `internal/config/` — CLI flag parsing, YAML config, defaults, validation
   - `internal/pool/` — buffer pooling (`sync.Pool`), command execution, temp directory pooling
   - `internal/cache/` — page count caching
   - `internal/metrics/` — performance metrics collection and reporting
   - `pkg/swiper/` — public API (`client.go`) and types (`types.go`)

### Performance Optimizations

- **Buffer Pooling**: Reuses `bytes.Buffer` instances via `sync.Pool`
- **Temp Directory Pool**: Pre-creates and reuses temp directories
- **Async Logging**: Non-blocking log channel prevents worker stalls
- **Parallel Processing**: Configurable worker count (defaults to CPU count)
- **Page Count Caching**: Single pdfinfo call per PDF

### Concurrency Model

- Uses goroutine worker pools for page processing
- Channel-based work distribution
- Mutex-protected shared state (success counts, page count cache)
- Graceful cleanup of resources (temp dirs, channels)

## Project Structure

```
swiper/
├── cmd/swiper/main.go          # CLI entry point
├── internal/
│   ├── batch/processor.go      # Batch PDF processing
│   ├── cache/cache.go          # Page count caching
│   ├── config/config.go        # Configuration and flag parsing
│   ├── extractor/
│   │   ├── extractor.go        # Core extraction engine
│   │   └── page.go             # Per-page processing
│   ├── metrics/
│   │   ├── collector.go        # Metrics collection
│   │   └── reporter.go         # Metrics reporting
│   ├── pool/
│   │   ├── buffer.go           # Buffer pooling (sync.Pool)
│   │   ├── command.go          # Command execution pool
│   │   └── tempdir.go          # Temp directory pooling
│   └── scanner/scanner.go      # PDF discovery and collection
├── pkg/swiper/
│   ├── client.go               # Public API
│   └── types.go                # Exported types
├── go.mod
├── go.sum
└── CLAUDE.md
```

Output structure (created automatically):
```
output_dir/           # Named after PDF or custom
├── index.md          # Main index linking all pages
├── page_1.md         # Markdown for page 1
├── page_2.md         # Markdown for page 2
└── images/           # Extracted images directory
    ├── page_1_img_1.jpg
    └── page_1_img_2.png
```

## Key Technical Details

### External Dependencies
- **gopkg.in/yaml.v2**: YAML configuration parsing
- **poppler-utils**: Required system utilities (pdfinfo, pdftotext, pdfimages)

### Command-Line Interface
- Uses Go's `flag` package for argument parsing
- Supports both flags and YAML config files
- Config file values overridden by CLI flags

### Error Handling
- Graceful degradation on page extraction failures
- Comprehensive error reporting with context
- Non-fatal errors logged, processing continues
- Summary statistics show success/failure counts

### Resource Management
- Automatic cleanup of temp directories
- Channel closure and flush before exit
- File handle management with proper defer statements

## Development Guidelines

### Making Changes
1. Modular architecture — entry point at `cmd/swiper/main.go`, packages in `internal/`
2. Maintain backward compatibility with existing CLI flags
3. Test with various PDF types (text-heavy, image-heavy, mixed)
4. Ensure poppler-utils commands remain compatible

### Testing Considerations
- Test with PDFs of different sizes (1 page to 1000+ pages)
- Verify temp directory cleanup
- Check concurrent processing with different worker counts
- Validate output markdown structure