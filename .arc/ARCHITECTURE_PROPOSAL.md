# Swiper Architecture Proposal

## Executive Summary

This document proposes a refactored architecture for Swiper that follows Go's module documentation standards, enhances maintainability, improves testability, and enables better performance optimization through clearer separation of concerns.

**Current State:** 1,776-line monolithic `main.go`
**Proposed State:** Modular package-based architecture with clear boundaries

---

## Current Architecture Analysis

### Strengths
- High-performance concurrent page processing
- Comprehensive metrics collection and reporting
- Advanced buffer pool management with multiple sizes
- Robust temp directory pooling
- Async logging to prevent worker blocking
- Result caching for performance optimization
- Context-based cancellation support

### Issues
1. **Maintainability:** Single file makes navigation and understanding difficult
2. **Testability:** No clear test boundaries, hard to unit test components
3. **Extensibility:** Adding new features requires modifying monolithic file
4. **Reusability:** Cannot use Swiper as a library in other Go projects
5. **Documentation:** Code organization doesn't reflect logical boundaries
6. **Collaboration:** Multiple developers would conflict on same file

---

## Proposed Architecture

### Directory Structure

```
swiper/
├── go.mod
├── go.sum
├── README.md
├── CLAUDE.md
├── ARCHITECTURE_PROPOSAL.md (this file)
├── verify-and-build.sh
│
├── cmd/
│   └── swiper/
│       └── main.go                 # CLI entry point (50-100 lines)
│
├── internal/                       # Private implementation packages
│   ├── extractor/
│   │   ├── extractor.go           # Core PDF extraction orchestration
│   │   ├── page.go                # Page processing logic
│   │   ├── image.go               # Image extraction (pdfimages wrapper)
│   │   ├── text.go                # Text extraction (pdftotext wrapper)
│   │   └── markdown.go            # Markdown generation
│   │
│   ├── scanner/
│   │   ├── scanner.go             # PDF file discovery and scanning
│   │   └── copy.go                # File copying with collision handling
│   │
│   ├── batch/
│   │   ├── processor.go           # Batch processing orchestration
│   │   └── worker.go              # Worker pool management
│   │
│   ├── pool/
│   │   ├── buffer.go              # Buffer pool with size-based management
│   │   ├── tempdir.go             # Temp directory pooling
│   │   └── command.go             # Command execution pooling
│   │
│   ├── metrics/
│   │   ├── collector.go           # Metrics collection logic
│   │   ├── reporter.go            # Metrics reporting and formatting
│   │   └── types.go               # Metric type definitions
│   │
│   ├── cache/
│   │   ├── cache.go               # Result caching implementation
│   │   └── types.go               # Cache type definitions
│   │
│   └── config/
│       ├── config.go              # Configuration management
│       ├── options.go             # Option types and validation
│       └── loader.go              # YAML config loading
│
├── pkg/                            # Public API packages
│   └── swiper/
│       ├── client.go              # Public API for library use
│       ├── types.go               # Public type definitions
│       ├── errors.go              # Exported error types
│       └── doc.go                 # Package documentation
│
├── docs/
│   ├── implementation_proof_protocol/
│   ├── architecture.md            # Architecture documentation
│   ├── api.md                     # Public API documentation
│   ├── performance.md             # Performance benchmarks
│   └── migration.md               # Migration guide from monolithic
│
└── tests/
    ├── integration/               # Integration tests
    │   ├── single_pdf_test.go
    │   ├── batch_test.go
    │   └── scanner_test.go
    ├── fixtures/                  # Test PDF files
    │   └── samples/
    └── benchmarks/                # Performance benchmarks
        └── extraction_bench_test.go
```

---

## Module Organization Principles

### 1. Internal vs Public Packages

**internal/** - Private implementation details
- Cannot be imported by external projects
- Free to change without breaking compatibility
- Contains all core business logic

**pkg/** - Public API
- Stable interface for library consumers
- Semantic versioning applies
- Comprehensive documentation required
- Minimal surface area

### 2. Package Responsibilities

#### cmd/swiper/main.go (CLI Entry Point)
**Responsibility:** CLI argument parsing, orchestration
**Lines of Code:** ~50-100
**Dependencies:** internal packages, pkg/swiper

```go
package main

import (
    "flag"
    "log"
    "os"

    "swiper/internal/config"
    "swiper/pkg/swiper"
)

func main() {
    // Parse flags
    // Load config
    // Create client
    // Execute command
    // Handle errors
}
```

#### internal/extractor/
**Responsibility:** PDF extraction orchestration
**Key Types:** `Extractor`, `Page`, `ExtractionResult`
**Key Functions:**
- `NewExtractor(pdfPath, outputDir string, opts ...Option) (*Extractor, error)`
- `ExtractPages(ctx context.Context) error`
- `GetPageCount() (int, error)`

#### internal/pool/
**Responsibility:** Resource pooling for performance
**Key Types:** `BufferPoolManager`, `TempDirPool`, `CommandPool`
**Benefits:**
- Reusable across extractor, scanner, batch processor
- Testable in isolation
- Performance tuning without affecting other components

#### internal/metrics/
**Responsibility:** Performance monitoring and reporting
**Key Types:** `Collector`, `Report`, `Metric`
**Benefits:**
- Pluggable metric backends (console, file, Prometheus)
- Test metric collection separately
- Add new metrics without changing core logic

#### internal/cache/
**Responsibility:** Result caching
**Key Types:** `ResultCache`, `CacheKey`, `CachedResult`
**Benefits:**
- Swap cache implementations (in-memory, Redis, disk)
- Test cache behavior independently
- Configure cache strategies per use case

#### pkg/swiper/
**Responsibility:** Public API for library usage
**Key Types:** `Client`, `Options`, `Result`, `Error`

```go
package swiper

// Client provides high-level API for PDF extraction
type Client struct {
    config *config.Config
}

// NewClient creates a new Swiper client
func NewClient(opts ...Option) (*Client, error)

// ExtractSingle extracts a single PDF
func (c *Client) ExtractSingle(ctx context.Context, pdfPath string) (*Result, error)

// ExtractBatch extracts multiple PDFs
func (c *Client) ExtractBatch(ctx context.Context, pdfDir string) ([]*Result, error)

// ScanAndCopy scans for PDFs and copies them
func (c *Client) ScanAndCopy(ctx context.Context, scanDir, copyDir string) error
```

---

## Implementation Strategy

### Aggressive Single-Week Refactor

Since backward compatibility is not required, we can execute a complete refactor rapidly.

### Day 1: Extraction Blitz
**Goal:** Rip apart monolith into packages (break everything)

**Morning:**
1. Create complete directory structure
2. Extract all type definitions to appropriate packages
3. Copy functions to new package files
4. Don't worry about compilation yet

**Afternoon:**
1. Extract `internal/metrics/` - complete package
2. Extract `internal/pool/` - complete package
3. Extract `internal/cache/` - complete package
4. Extract `internal/config/` - complete package

**Exit Criteria:**
- All support packages created
- Types and functions moved
- Compilation broken (expected)

### Day 2: Core Logic Extraction
**Goal:** Extract domain logic packages

**Morning:**
1. Extract `internal/extractor/` - core extraction
2. Extract `internal/scanner/` - PDF scanning
3. Extract `internal/batch/` - batch processing

**Afternoon:**
1. Create `pkg/swiper/` public API
2. Rewrite `cmd/swiper/main.go` (50-100 lines)
3. Wire up all dependencies

**Exit Criteria:**
- All packages created
- Main.go minimal
- Code compiles (even if broken)

### Day 3: Integration & Fixing
**Goal:** Make it work

**Morning:**
1. Fix compilation errors
2. Resolve circular dependencies
3. Clean up imports

**Afternoon:**
1. Test single PDF extraction
2. Test batch processing
3. Test scanning mode
4. Fix runtime errors

**Exit Criteria:**
- Binary compiles
- Core workflows functional
- No crashes

### Day 4: Testing & Optimization
**Goal:** Ensure quality and performance

**Morning:**
1. Write unit tests for critical packages
2. Create integration tests
3. Add benchmark tests

**Afternoon:**
1. Run performance benchmarks
2. Profile for bottlenecks
3. Optimize if needed
4. Verify metrics still work

**Exit Criteria:**
- Tests pass
- Performance maintained or improved
- Metrics functional

### Day 5: Documentation & Polish
**Goal:** Production ready

**Morning:**
1. Document public API (godoc)
2. Update CLAUDE.md
3. Update README.md
4. Write migration notes

**Afternoon:**
1. Create example programs
2. Test build script
3. Final verification
4. Git commit and push

**Exit Criteria:**
- Documentation complete
- Examples work
- Build verified
- Ready to use

---

## Benefits of Proposed Architecture

### Maintainability
- **Smaller files:** Each file focuses on single responsibility
- **Clear boundaries:** Package structure reflects logical organization
- **Easier navigation:** Developers quickly find relevant code
- **Reduced cognitive load:** Understand one package at a time

### Testability
- **Unit testing:** Test packages in isolation
- **Mock interfaces:** Inject dependencies for testing
- **Test organization:** Tests alongside implementation
- **Faster tests:** Run subset of tests during development

### Performance
- **Profiling:** Identify bottlenecks by package
- **Optimization:** Optimize pools without touching extraction logic
- **Resource management:** Clear ownership of resources
- **Metrics:** Package-level performance metrics

### Extensibility
- **Plugin architecture:** New extractors (e.g., OCR) as plugins
- **Alternative implementations:** Swap cache or pool implementations
- **Feature flags:** Enable/disable features per package
- **API versioning:** Stable public API with internal flexibility

### Reusability
- **Library usage:** Other Go projects can import pkg/swiper
- **Component reuse:** Use buffer pools in other projects
- **Shareable config:** Config package usable elsewhere
- **Common patterns:** Establish patterns for other tools

### Collaboration
- **Parallel development:** Multiple developers work on different packages
- **Code review:** Review smaller, focused changes
- **Git conflicts:** Reduced file conflicts
- **Ownership:** Clear package ownership

---

## Migration Path

### No Backward Compatibility Required
- This is an internal tool - breaking changes acceptable
- We can completely refactor without constraints
- Opportunity for clean-slate redesign
- Can optimize for best practices, not legacy support

### Aggressive Migration Strategy
1. **Day 1-2:** Complete package extraction, break everything
2. **Day 3:** Wire up new architecture, make it compile
3. **Day 4:** Fix integration, verify core functionality
4. **Day 5:** Testing, benchmarking, documentation

### Approach
- Burn the ships: delete monolithic code once extracted
- No parallel maintenance of old structure
- Focus on ideal architecture, not compatibility
- Fast iteration without migration overhead

### Risk Mitigation
- Git history preserves old version if needed
- Comprehensive integration tests ensure functionality
- Performance benchmarks ensure no regression
- Single developer ownership allows bold moves

---

## Performance Considerations

### Current Performance Characteristics
- Parallel page processing: configurable worker count
- Buffer pooling: 4 size tiers (32KB, 128KB, 256KB, 1MB)
- Temp directory pooling: 2x worker count
- Async logging: non-blocking channel
- Result caching: in-memory with metrics

### Performance Preservation Strategy
1. **Benchmark before refactoring:** Establish baseline
2. **Benchmark after each phase:** Ensure no regression
3. **Profile regularly:** Identify optimization opportunities
4. **Monitor metrics:** Track changes in performance

### Expected Performance Impact
- **No degradation:** Refactoring preserves logic
- **Potential improvements:** Better resource locality
- **Easier optimization:** Clear bottleneck identification
- **Better scaling:** Worker pools more configurable

---

## Go Module Best Practices

### Module Path
```go
module github.com/dusk-labs/swiper

go 1.24

require (
    gopkg.in/yaml.v2 v2.4.0
)
```

### Versioning Strategy
- **v0.x.x:** Current development (breaking changes allowed)
- **v1.0.0:** Stable public API release
- **v1.x.x:** Backward-compatible additions
- **v2.0.0:** Major breaking changes (if needed)

### Internal Package Protection
- Use `internal/` to prevent external imports
- Only `pkg/` packages are public
- Clear contract between internal and public

### Documentation Standards
- Every public function has godoc comment
- Package documentation in doc.go
- Examples in _test.go files
- Architecture docs in docs/

---

## Testing Strategy

### Unit Tests
```
internal/pool/buffer_test.go
internal/metrics/collector_test.go
internal/cache/cache_test.go
```

### Integration Tests
```
tests/integration/single_pdf_test.go
tests/integration/batch_test.go
tests/integration/scanner_test.go
```

### Benchmark Tests
```
tests/benchmarks/extraction_bench_test.go
```

### Test Coverage Goals
- **Internal packages:** 80%+ coverage
- **Public API:** 90%+ coverage
- **Integration tests:** All workflows
- **Benchmark tests:** All critical paths

---

## Metrics and Monitoring

### Current Metrics
- Pages processed
- Text extracted (bytes)
- Images extracted (count)
- Bytes processed (I/O)
- Processing times
- Cache hit/miss rates
- Buffer pool statistics
- Worker utilization
- Queue depth

### Enhanced Metrics (Future)
- Per-package metrics
- Resource usage tracking
- Error rates and types
- Latency percentiles (p50, p95, p99)
- Throughput metrics

---

## Conclusion

This architecture proposal transforms Swiper from a monolithic tool into a well-structured, maintainable, and reusable Go project that follows industry best practices and Go module conventions.

### Key Outcomes
1. **Improved maintainability:** Clear package structure
2. **Enhanced testability:** Isolated components
3. **Better performance:** Easier optimization
4. **Increased reusability:** Library API
5. **Scalable development:** Team collaboration
6. **Clean slate:** No legacy baggage

### Next Steps
1. Commit current working state to git (backup)
2. Execute Day 1: Extraction Blitz
3. Execute Day 2: Core Logic Extraction
4. Execute Day 3: Integration & Fixing
5. Execute Day 4: Testing & Optimization
6. Execute Day 5: Documentation & Polish

### Success Criteria
- Binary functionality intact (behavior may differ)
- Performance maintained or improved
- Test coverage >70% (initial target)
- Public API documented
- Developer experience dramatically improved
- Code maintainability 10x better

---

**Document Version:** 2.0
**Date:** 2025-09-29
**Author:** Architecture Planning
**Status:** Approved - Ready for Implementation
**Breaking Changes:** Accepted - No backward compatibility required