# Changelog

All notable changes to the Swiper PDF extraction tool will be documented in this file.

## [Unreleased] - 2026-03-26

### Added

- **Test infrastructure**: Unit tests for all internal packages (cache, config, pool, metrics, extractor, scanner, batch) with `-race` support and test PDF fixtures
- **PageResult struct**: Structured per-page extraction outcome tracking with `Success()`, `PartialSuccess()`, `ErrorSummary()`, and `MarkdownErrorComment()` methods
- **Real config validation**: Mode exclusivity enforcement, file/directory existence checks, page range syntax validation via `errors.Join`
- **Shared AsyncLogger package** (`internal/log`): Replaces 3 duplicate `logChan`/`asyncLogger()` implementations with a single reusable logger supporting quiet mode and backpressure fallback
- **Page range filtering**: `-pages 1-10,50` flag with `ParsePageRanges()` and `ExpandPages()` for selective page extraction
- **Quiet mode**: `-q` flag suppresses non-error output via `AsyncLogger` quiet parameter
- **Poppler install hints**: Platform-specific install instructions shown when poppler binaries are missing
- **Progress reporter**: Rate-limited `[N/M] X% - Page N (Xs) ETA Xs` output during extraction
- **Per-page image subdirectories**: `WithPerPageImageDirs(true)` option organizes images into `images/page_N/` subdirectories
- **Batch resume**: `-resume` flag with `.swiper-progress` JSON state file for resuming interrupted batch processing
- **Page-range text batching**: `extractTextBatch()` method and benchmarks for multi-page `pdftotext` calls (evaluated, single-page remains default)
- **PageSummary in public API**: `Result.PageResults` and `Result.Duration` fields populated by `Client.ExtractSingle`
- **Configurable CommandPool timeout**: `NewCommandPool(ctx, timeout)` replaces hardcoded 30s
- **Buffer size constants**: `SmallBufferSize`, `MediumBufferSize`, `LargeBufferSize`, `XLargeBufferSize`
- **Comprehensive README**: Installation, usage, flag reference, output format documentation

### Changed

- **Architecture consolidation**: Deleted monolithic root `main.go` (1,775 lines) and `devtmp/` experimental directory; modular `cmd/swiper/` + `internal/` layout is now canonical
- **CLI restructured**: All modes (single, batch, scan) go through unified config -> SetDefaults -> Validate -> dispatch path
- **Help text rebranded**: "PDF Tool" -> "Swiper - High-performance PDF to markdown converter"
- **CLAUDE.md updated**: Reflects modular architecture, correct build command (`go build -o swiper ./cmd/swiper/`), updated project structure
- **Buffer pool tier boundaries fixed**: `<` changed to `<=` so exact-boundary files get the matching tier instead of being bumped to the next larger pool
- **Context propagation**: All `exec.Command` calls replaced with `exec.CommandContext` for proper cancellation support
- **Function decomposition**: `ExtractPages` split into `calculateWorkerCount`, `runWorkerPool`, `reportProgress`, `reportExtractionSummary`; `extractImagesFromPage` split with `copyImagesFromDir`
- **Path naming normalized**: `pdfFile` -> `pdfPath` across extractor package
- **`ExtractPages` no longer calls `Cleanup()` internally**: Callers own cleanup via `defer`

### Fixed

- **Swallowed errors in `extractTextFromPage`**: Now returns `fmt.Errorf("pdftotext failed: %w", err)` instead of `("", nil)`
- **Swallowed errors in `extractImagesFromPage`**: Returns `([]string, []error)` with per-image error tracking instead of `(nil, nil)` or discarding partial results
- **TempDirPool init/Cleanup race**: Added `sync.WaitGroup` and `closeMu` guard to prevent `init()` goroutine from sending on closed channel
- **TempDirPool `GetTempDir` after `Cleanup`**: Added `closed` guard to prevent use after cleanup
- **AsyncLogger double-close**: Internal `sync.Once` guard in `Close()` method
- **`-copydir` flag default bleeding through `Merge()`**: Changed from `"pdf-docs"` to empty default
- **`context.Canceled` comparison**: Changed `==` to `errors.Is()` for wrapped errors
- **Dead semaphore close**: Removed `close(sem)` on semaphore channel in image extraction
- **Test directory leak**: `TestNew_EmptyOutputDir` now cleans up created directory
- **Validate early return**: Returns immediately on conflicting modes instead of continuing file checks
- **Dead code removed**: `calculateOptimalBufferSize` initial assignment, unreliable `MemStats.Sys` memory scaling
- **Resume state**: `saveProgress` now logs write errors; failed PDF list reset on each resume run
- **`createMainMarkdown`**: Now uses `pageResults` instead of `1..totalPages`, respecting `--pages` filter

## [Unreleased] - 2026-02-27

### Changed

- **Module path qualified**: Changed `module swiper` to `module github.com/onedusk/swiper` in `go.mod` to enable cross-module imports
- Updated 14 import paths across 7 files (`cmd/swiper/main.go`, `pkg/swiper/client.go`, `pkg/swiper/types.go`, `internal/batch/processor.go`, `internal/extractor/extractor.go`, `internal/extractor/page.go`, `internal/scanner/scanner.go`)

### Fixed

- **Double-close panic in Extractor**: `Cleanup()` now uses `sync.Once` to prevent panic when called multiple times (e.g. deferred `Cleanup()` after `ExtractPages()` already cleaned up resources)

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
