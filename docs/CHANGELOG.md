# Changelog

All notable changes to the Swiper PDF extraction tool will be documented in this file.

## [Unreleased] - 2026-02-27

### Changed

- **Module path qualified**: Changed `module swiper` to `module github.com/onedusk/swiper` in `go.mod` to enable cross-module imports
- Updated 14 import paths across 7 files (`cmd/swiper/main.go`, `pkg/swiper/client.go`, `pkg/swiper/types.go`, `internal/batch/processor.go`, `internal/extractor/extractor.go`, `internal/extractor/page.go`, `internal/scanner/scanner.go`)

## [Unreleased] - 2025-09-26

### Added

- **BufferPoolManager**: Centralized buffer pool management with multiple size tiers (32KB, 128KB, 256KB, 1MB) for adaptive buffer selection based on file size hints
- **Enhanced Performance Metrics**:
  - Buffer pool metrics tracking (hits, misses, created bytes)
  - Worker utilization tracking
  - Page queue depth analysis
  - More comprehensive reporting in PrintSummary
- **Smart Channel Sizing Heuristics**:
  - `calculateOptimalBufferSize` function with adaptive sizing based on PDF size
  - `calculateOptimalConcurrentPDFs` function that considers memory availability
  - Dynamic channel sizing in NewExtractor based on worker count
  - Smart page queue buffering in extractPages

### Changed

- **Refactored extractImagesFromPage** (main.go:565-677):
  - Eliminated unnecessary result slice rebuilding
  - Pre-allocates result slice to avoid multiple allocations
  - Uses semaphore pattern to limit concurrent file operations
  - Improved error handling with sync.Once pattern
  - Maintains image order with indexed results
- **Optimized pdfinfo Parsing** (main.go:449-495):
  - Uses buffered reading with bufio.Scanner
  - Line-by-line streaming instead of loading entire output
  - Atomic operations for cache access
  - Eliminated redundant string operations
- **Streamlined Concurrency Handling**:
  - PDFScanner now uses concurrent file copying with semaphore pattern (main.go:1111-1208)
  - Enhanced extractPages with producer-consumer pattern and batch yielding
  - Worker performance tracking and adaptive scaling for large PDFs
  - Improved error handling with atomic counters

### Performance Improvements

- **Memory Efficiency**: Shared buffer pools reduce memory allocation overhead
- **Reduced Lock Contention**: Use of atomic operations where possible
- **Better Resource Utilization**: Adaptive worker scaling based on PDF size and system resources
- **Improved Throughput**: Smart channel buffering and semaphore-based concurrency control
- **Enhanced Observability**: Comprehensive metrics for monitoring performance bottlenecks

### Technical Details

- Integrated metrics tracking across all components (Extractor, PDFScanner, BatchProcessor)
- Shared BufferPoolManager instance across workflows for better resource management
- Backward compatibility maintained while significantly improving performance for large-scale PDF processing operations
