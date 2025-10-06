# Changelog

All notable changes to the Swiper PDF extraction tool will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned
- Context propagation through all public APIs (A01)
- filepath.WalkDir migration for faster directory traversal (A02)
- Atomic file copy with temp+rename pattern (A03)

## [0.1.0] - 2025-10-06

### Added
- **Absolute path comparison for directory skip** ([#3](https://github.com/onedusk/swiper/pull/3))
  - Added `scanDirAbs` and `copyDirAbs` fields to PDFScanner struct
  - Resolve paths to absolute form using `filepath.Abs()` in `New()`
  - Updated `FindPDFs()` to use absolute path comparison and prefix checking
  - Prevents infinite recursion when copyDir is located within scanDir
  - Added logging for skipped directories
  - Implements A04 from ATOMIC_IMPROVEMENT_PLAN.md

### Technical Details
- `internal/scanner/scanner.go`: Robust directory skipping with canonical path comparison
- `testdata/a04-nested/`: Test fixtures for validation

## [0.0.2] - 2025-10-06

### Fixed
- **Critical**: Resolved double-close channel panic in batch processing ([#2](https://github.com/onedusk/swiper/pull/2))
  - Made `Cleanup()` idempotent using `sync.Once` to prevent duplicate cleanup calls
  - Removed duplicate cleanup from `ExtractPages()` method
  - Fixed race condition in `TempDirPool` initialization goroutine
  - Added closed flag with mutex protection to prevent sending to closed channels
  - Tested successfully with 61 PDFs (1,514 pages) in batch mode

### Technical Details
- `internal/extractor/extractor.go`: Added `cleanupOnce sync.Once` field for idempotent cleanup
- `internal/pool/tempdir.go`: Added `closed bool` flag and `mu sync.RWMutex` for thread-safe state management

## [0.0.1] - 2025-09-26

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