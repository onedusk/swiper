# Swiper Aggressive Refactor Plan

## Overview

Complete architectural refactor of Swiper from a 1,776-line monolithic `main.go` to a modular, maintainable Go project following best practices.

**Constraint:** No backward compatibility required - this is an internal tool
**Timeline:** 5 days
**Approach:** Aggressive, break-everything-first refactor

---

## Current State

**File:** `main.go` (1,776 lines)

**Components (all in one file):**
- Options struct and config loading (28-1151)
- MetricsCollector (40-224)
- ResultCache (227-265)
- BufferPoolManager (268-355)
- CommandPool (358-379)
- Extractor (381-1044)
- PDFScanner (1154-1405)
- BatchProcessor (1408-1641)
- Main CLI (1643-1775)

**Strengths:**
- High performance concurrent processing
- Comprehensive metrics
- Advanced buffer pooling
- Robust error handling

**Problems:**
- Impossible to navigate
- Cannot unit test components
- Hard to extend
- Not reusable as library
- Difficult collaboration

---

## Target Architecture

```
swiper/
├── cmd/swiper/main.go          (50-100 lines)
├── internal/
│   ├── metrics/                (collector.go, reporter.go, types.go)
│   ├── pool/                   (buffer.go, tempdir.go, command.go)
│   ├── cache/                  (cache.go, types.go)
│   ├── config/                 (config.go, options.go, loader.go)
│   ├── extractor/              (extractor.go, page.go, image.go, text.go, markdown.go)
│   ├── scanner/                (scanner.go, copy.go)
│   └── batch/                  (processor.go, worker.go)
├── pkg/swiper/                 (client.go, types.go, errors.go, doc.go)
└── tests/
    ├── integration/
    ├── fixtures/
    └── benchmarks/
```

---

## 5-Day Implementation Plan

### Day 1: Extraction Blitz (Break Everything)

**Goal:** Rip apart monolith into packages

**Morning: Directory Setup (2 hours)**
```bash
mkdir -p cmd/swiper
mkdir -p internal/{metrics,pool,cache,config,extractor,scanner,batch}
mkdir -p pkg/swiper
mkdir -p tests/{integration,fixtures,benchmarks}
```

**Afternoon: Extract Support Packages (4 hours)**

1. **internal/metrics/** (200 lines)
   - Extract MetricsCollector (lines 40-224)
   - Create collector.go, reporter.go, types.go
   - No dependencies on other packages

2. **internal/pool/** (180 lines)
   - Extract BufferPoolManager (lines 268-355)
   - Extract CommandPool (lines 358-379)
   - Extract temp dir pooling functions
   - Create buffer.go, tempdir.go, command.go

3. **internal/cache/** (80 lines)
   - Extract ResultCache (lines 227-265)
   - Create cache.go, types.go

4. **internal/config/** (120 lines)
   - Extract Options (lines 28-37)
   - Extract loadConfig (lines 1141-1151)
   - Create config.go, options.go, loader.go

**Exit:** Compilation broken, packages created

---

### Day 2: Core Logic Extraction (Make it Compile)

**Goal:** Extract domain packages and wire up

**Morning: Domain Packages (4 hours)**

1. **internal/extractor/** (400+ lines)
   - Extract Extractor struct and methods (lines 381-1044)
   - Split into:
     - extractor.go - core orchestration
     - page.go - processPage
     - image.go - extractImagesFromPage
     - text.go - extractTextFromPage
     - markdown.go - createMainMarkdown

2. **internal/scanner/** (250 lines)
   - Extract PDFScanner (lines 1154-1405)
   - Split into scanner.go, copy.go

3. **internal/batch/** (230 lines)
   - Extract BatchProcessor (lines 1408-1641)
   - Split into processor.go, worker.go

**Afternoon: Public API & CLI (3 hours)**

1. **pkg/swiper/** (150 lines)
   ```go
   // client.go
   type Client struct { ... }
   func NewClient(opts ...Option) (*Client, error)
   func (c *Client) ExtractSingle(...) (*Result, error)
   func (c *Client) ExtractBatch(...) ([]*Result, error)
   func (c *Client) ScanAndCopy(...) error

   // types.go
   type Result struct { ... }
   type Option func(*Client)

   // errors.go
   var ErrPDFNotFound = errors.New(...)
   ```

2. **cmd/swiper/main.go** (50-100 lines)
   ```go
   func main() {
       // Parse flags
       // Create swiper.Client
       // Call appropriate method
       // Handle errors
   }
   ```

**Exit:** Code compiles (may have runtime issues)

---

### Day 3: Integration & Fixing (Make it Work)

**Goal:** Fix all runtime issues, make workflows functional

**Morning: Compilation Fixes (3 hours)**
1. Resolve import cycles
2. Fix missing exports (uppercase first letter)
3. Update all cross-package references
4. Ensure proper initialization

**Afternoon: Runtime Testing (3 hours)**
1. Test: `./swiper -file test.pdf`
   - Fix crashes
   - Fix extraction errors
   - Verify output

2. Test: `./swiper -dir test-pdfs/`
   - Fix batch processing
   - Verify all PDFs extract

3. Test: `./swiper -scan . -copydir pdfs`
   - Fix scanner
   - Verify file copying

**Exit:** All three workflows work end-to-end

---

### Day 4: Testing & Optimization (Make it Right)

**Goal:** Ensure quality and performance

**Morning: Unit Tests (4 hours)**

1. **internal/metrics/collector_test.go**
   ```go
   func TestMetricsCollector_RecordPageProcessed(t *testing.T)
   func TestMetricsCollector_GetCacheHitRate(t *testing.T)
   ```

2. **internal/pool/buffer_test.go**
   ```go
   func TestBufferPoolManager_GetBuffer(t *testing.T)
   func TestBufferPoolManager_SizeSelection(t *testing.T)
   ```

3. **internal/cache/cache_test.go**
   ```go
   func TestResultCache_GetSet(t *testing.T)
   func TestResultCache_Concurrency(t *testing.T)
   ```

**Afternoon: Integration & Benchmarks (3 hours)**

1. **tests/integration/single_pdf_test.go**
   ```go
   func TestExtractSinglePDF(t *testing.T)
   ```

2. **tests/benchmarks/extraction_bench_test.go**
   ```go
   func BenchmarkPageExtraction(b *testing.B)
   func BenchmarkBatchProcessing(b *testing.B)
   ```

3. Run benchmarks, compare with baseline
4. Profile if performance regression
5. Optimize hot paths

**Exit:** Tests pass, performance maintained

---

### Day 5: Documentation & Polish (Make it Shine)

**Goal:** Production ready

**Morning: Documentation (3 hours)**

1. **pkg/swiper/doc.go**
   ```go
   // Package swiper provides high-performance PDF extraction.
   //
   // Example usage:
   //   client, _ := swiper.NewClient()
   //   result, _ := client.ExtractSingle(ctx, "doc.pdf")
   ```

2. Add godoc comments to all exported functions

3. **Update CLAUDE.md**
   - New directory structure
   - Development commands
   - Testing instructions

4. **Update README.md**
   - Installation
   - Usage examples
   - Library usage

**Afternoon: Final Verification (2 hours)**

1. Run `./verify-and-build.sh`
2. Test all examples in documentation
3. Verify godoc looks good: `go doc -all swiper`
4. Create example programs in `examples/`
5. Git commit and push

**Exit:** Fully documented, production ready

---

## Code Migration Checklist

### Types to Extract
- [ ] Options → internal/config/
- [ ] MetricsCollector → internal/metrics/
- [ ] ResultCache → internal/cache/
- [ ] BufferPoolManager → internal/pool/
- [ ] CommandPool → internal/pool/
- [ ] Extractor → internal/extractor/
- [ ] PDFScanner → internal/scanner/
- [ ] BatchProcessor → internal/batch/

### Functions to Extract
- [ ] loadConfig → internal/config/
- [ ] copyFile → internal/scanner/ or internal/extractor/
- [ ] generateRandomHex → internal/extractor/
- [ ] calculateOptimalBufferSize → internal/extractor/
- [ ] calculateOptimalConcurrentPDFs → internal/batch/
- [ ] min → internal/extractor/ or utils

### Global Variables
- [ ] bufferPool → internal/pool/
- [ ] Context handling → per package

---

## Package Dependency Rules

**Dependency Flow (allowed):**
```
cmd/swiper/
  └─> pkg/swiper/
       └─> internal/batch/
            └─> internal/extractor/
                 ├─> internal/pool/
                 ├─> internal/cache/
                 ├─> internal/metrics/
                 └─> internal/config/

internal/scanner/
  ├─> internal/pool/
  ├─> internal/metrics/
  └─> internal/config/
```

**No circular dependencies allowed**

**External dependencies:**
- gopkg.in/yaml.v2 (config only)
- Standard library everywhere

---

## Testing Strategy

### Unit Tests (70% coverage target)
- internal/metrics/ - 90%
- internal/pool/ - 85%
- internal/cache/ - 90%
- internal/config/ - 80%
- internal/extractor/ - 65%
- internal/scanner/ - 70%
- internal/batch/ - 70%

### Integration Tests
- Single PDF extraction
- Batch processing
- Scanner workflow
- Error scenarios

### Benchmark Tests
- Page extraction performance
- Batch processing throughput
- Memory usage
- CPU profiling

---

## Performance Targets

**Must maintain or improve:**
- Pages per second
- Memory usage
- CPU utilization
- I/O throughput

**Baseline (from current monolith):**
- Run: `./swiper -file large.pdf -benchmark`
- Record: pages/sec, memory, time

**After refactor:**
- Run same benchmark
- Compare results
- Profile if regression
- Optimize if needed

---

## Git Strategy

### Before Starting
```bash
git add -A
git commit -m "Checkpoint: Monolithic version before refactor"
git tag v0.9-monolithic
git push origin v0.9-monolithic
```

### During Refactor
```bash
# Day 1
git checkout -b refactor-modular
git commit -m "Day 1: Extract support packages"

# Day 2
git commit -m "Day 2: Extract domain logic"

# Day 3
git commit -m "Day 3: Fix integration"

# Day 4
git commit -m "Day 4: Add tests and benchmarks"

# Day 5
git commit -m "Day 5: Documentation and polish"
git tag v1.0.0
```

### Rollback Plan
If everything breaks:
```bash
git checkout v0.9-monolithic
```

---

## Success Criteria

- [ ] Binary compiles
- [ ] All three workflows functional (single, batch, scan)
- [ ] Tests pass (>70% coverage)
- [ ] Performance maintained
- [ ] Public API documented
- [ ] Example programs work
- [ ] CLAUDE.md updated
- [ ] README.md updated
- [ ] Code maintainability 10x better

---

## Risk Mitigation

**Risk:** Circular dependencies
**Mitigation:** Strict dependency flow, interfaces

**Risk:** Performance regression
**Mitigation:** Benchmark before/after, profile

**Risk:** Breaking changes
**Mitigation:** Accepted - not a risk

**Risk:** Lost functionality
**Mitigation:** Integration tests

**Risk:** Incomplete migration
**Mitigation:** 5-day deadline, focused scope

---

## Resources

**Required Tools:**
- Go 1.24+
- poppler-utils
- git
- profiling tools (pprof)

**Reference Documentation:**
- Go module documentation (docs/doc.md)
- Implementation Proof Protocol (docs/implementation_proof_protocol/)
- Current CLAUDE.md

**Inspiration:**
- Standard library structure
- Go project layout (https://github.com/golang-standards/project-layout)
- Clean Architecture principles

---

## Post-Refactor Benefits

**Immediate:**
- 10x better code navigation
- 100% testable components
- Clear package boundaries
- Professional structure

**Medium-term:**
- Add features without breaking things
- Optimize individual components
- Library usage in other projects
- Team collaboration possible

**Long-term:**
- Extensible plugin architecture
- Alternative implementations (cloud, GPU)
- Community contributions
- Battle-tested patterns

---

**Document Version:** 1.0
**Created:** 2025-09-29
**Status:** Ready to Execute
**Commitment:** 5 days, no compromises, no backward compatibility