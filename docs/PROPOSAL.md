# Swiper Modular Architecture Proposal

## Purpose

Propose a modular architecture for Swiper that follows Go best practices. This is a **proposal only** - no implementation will begin until approved.

---

## Current State

**Structure:** Single file monolith
**File:** `main.go` (1,776 lines)
**Status:** Working, production-ready, high-performance

**Components (all in one file):**
- Configuration and options
- Metrics collection
- Result caching
- Buffer pool management
- Command pooling
- PDF extraction (core logic)
- PDF scanning and copying
- Batch processing
- CLI interface

**What Works Well:**
- High-performance concurrent processing
- Advanced resource pooling
- Comprehensive metrics
- Robust error handling
- Production-ready reliability

**What Could Be Better:**
- Hard to navigate (1,776 lines)
- Cannot unit test components
- Not reusable as library
- Difficult to extend
- No clear boundaries

---

## Proposed Architecture

### Package Structure

```
swiper/
├── go.mod
├── go.sum
│
├── cmd/
│   └── swiper/
│       └── main.go              # CLI entry point (minimal)
│
├── internal/                    # Private implementation
│   ├── metrics/                 # Performance tracking
│   │   ├── collector.go
│   │   └── types.go
│   │
│   ├── pool/                    # Resource pooling
│   │   ├── buffer.go
│   │   ├── tempdir.go
│   │   └── command.go
│   │
│   ├── cache/                   # Result caching
│   │   └── cache.go
│   │
│   ├── config/                  # Configuration
│   │   ├── config.go
│   │   └── loader.go
│   │
│   ├── extractor/               # PDF extraction
│   │   ├── extractor.go
│   │   ├── page.go
│   │   ├── image.go
│   │   └── text.go
│   │
│   ├── scanner/                 # PDF discovery
│   │   └── scanner.go
│   │
│   └── batch/                   # Batch processing
│       └── processor.go
│
├── pkg/                         # Public API (optional)
│   └── swiper/
│       ├── client.go
│       ├── types.go
│       └── errors.go
│
└── tests/
    ├── integration/
    └── benchmarks/
```

### Package Responsibilities

**internal/metrics/**
- Collect performance metrics
- Track statistics
- Generate reports
- No external dependencies

**internal/pool/**
- Buffer pool management (4 size tiers)
- Temp directory pooling
- Command execution pooling
- Shared resource management

**internal/cache/**
- Result caching
- Cache key management
- Hit/miss tracking

**internal/config/**
- Configuration loading
- Option validation
- YAML parsing

**internal/extractor/**
- Core PDF extraction logic
- Page processing
- Image extraction
- Text extraction

**internal/scanner/**
- PDF file discovery
- Recursive scanning
- File copying

**internal/batch/**
- Batch orchestration
- Worker pool management
- Concurrent processing

**pkg/swiper/** (Optional)
- Public library API
- Client interface
- Allows usage as Go library

**cmd/swiper/main.go**
- CLI argument parsing
- Command dispatch
- Error handling
- Minimal glue code

---

## Benefits

### Maintainability
- Small, focused files (100-300 lines each)
- Clear responsibilities
- Easy to navigate
- Reduced cognitive load

### Testability
- Unit test each package
- Mock dependencies
- Test in isolation
- Faster test runs

### Extensibility
- Add features without breaking existing code
- Swap implementations
- Plugin architecture possible
- Clear extension points

### Reusability
- Use as Go library (pkg/swiper)
- Share packages across projects
- Common patterns

### Performance
- Easier to profile by package
- Optimize specific components
- Better resource locality
- Clear bottleneck identification

---

## Migration Strategy

### Approach
Since backward compatibility is not required, we can:
- Completely refactor structure
- Break everything temporarily
- Focus on ideal architecture
- No legacy constraints

### High-Level Steps
1. Create directory structure
2. Extract packages one by one
3. Wire up dependencies
4. Add tests
5. Verify functionality
6. Document

### Risk Mitigation
- Git history preserves old version
- Can rollback anytime
- Integration tests ensure functionality
- Performance benchmarks prevent regression

---

## What This Proposal Includes

✅ Complete package structure design
✅ Responsibility boundaries
✅ Migration strategy
✅ Benefits and trade-offs
✅ Example API design

## What This Proposal Does NOT Include

❌ Actual implementation
❌ Code changes
❌ Timeline commitments
❌ Resource allocation
❌ Breaking changes (yet)

---

## Example Public API (If Approved)

```go
package swiper

// Client provides PDF extraction capabilities
type Client struct {
    config *config.Config
}

// NewClient creates a new Swiper client
func NewClient(opts ...Option) (*Client, error) {
    // Implementation
}

// ExtractSingle extracts a single PDF
func (c *Client) ExtractSingle(ctx context.Context, pdfPath string) (*Result, error) {
    // Implementation
}

// ExtractBatch extracts multiple PDFs
func (c *Client) ExtractBatch(ctx context.Context, pdfDir string) ([]*Result, error) {
    // Implementation
}

// ScanAndCopy scans and copies PDFs
func (c *Client) ScanAndCopy(ctx context.Context, scanDir, copyDir string) error {
    // Implementation
}
```

### Library Usage Example
```go
import "swiper/pkg/swiper"

func main() {
    client, _ := swiper.NewClient(
        swiper.WithProcessCount(8),
        swiper.WithCache(true),
    )

    result, _ := client.ExtractSingle(context.Background(), "doc.pdf")
    fmt.Printf("Extracted %d pages\n", result.PageCount)
}
```

---

## Dependencies

**External:**
- gopkg.in/yaml.v2 (configuration)
- Standard library

**System:**
- poppler-utils (pdfinfo, pdftotext, pdfimages)

**Go Version:**
- 1.24+ required

---

## Performance Expectations

### Must Maintain
- Current pages/second throughput
- Memory usage characteristics
- CPU utilization patterns
- I/O performance

### Strategy
- Benchmark before refactoring
- Benchmark after refactoring
- Profile if any regression
- Optimize as needed

### No Expected Degradation
Package boundaries should not impact performance. May even improve due to better code locality.

---

## Testing Approach

### Unit Tests
- Test each package independently
- Mock dependencies
- 70%+ coverage target

### Integration Tests
- End-to-end workflows
- Single PDF extraction
- Batch processing
- Scanner functionality

### Benchmark Tests
- Performance verification
- Regression detection
- Bottleneck identification

---

## Decision Required

### Options

**Option 1: Approve Proposal**
- Proceed with modular refactor
- Timeline TBD after approval
- Breaking changes accepted

**Option 2: Modify Proposal**
- Suggest changes to package structure
- Adjust scope or approach
- Re-submit for approval

**Option 3: Keep Current Structure**
- No refactoring
- Maintain monolithic design
- Continue as-is

---

## Questions for Decision Makers

1. **Is library usage (pkg/swiper) desired?**
   - Yes → Include public API
   - No → Keep internal only

2. **What is acceptable timeline?**
   - Days, weeks, or months?
   - Resource availability?

3. **Are breaking changes truly acceptable?**
   - CLI interface changes
   - Output format changes
   - Configuration changes

4. **What is priority?**
   - Maintainability
   - Library usage
   - Performance
   - Extensibility

5. **Testing requirements?**
   - Coverage percentage?
   - Test types needed?
   - CI/CD integration?

---

## Next Steps (If Approved)

1. **Create detailed implementation plan**
   - Break down tasks
   - Estimate timeline
   - Identify dependencies

2. **Set up git branch**
   - Tag current version
   - Create refactor branch
   - Plan commits

3. **Begin implementation**
   - Follow approved plan
   - Regular check-ins
   - Iterative approach

4. **Testing and validation**
   - Write tests as we go
   - Benchmark continuously
   - Document everything

5. **Review and merge**
   - Code review
   - Final testing
   - Merge to main

---

## Appendices

### Related Documents
- **QUICKSTART.md** - Current usage guide
- **CLAUDE.md** - Development guidelines
- **verify-and-build.sh** - Build automation
- **docs/doc.md** - Go modules reference
- **docs/implementation_proof_protocol/** - IPP framework

### References
- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Go Modules Reference](https://go.dev/ref/mod)
- [Effective Go](https://go.dev/doc/effective_go)

---

## Proposal Status

**Version:** 1.0
**Date:** 2025-09-29
**Status:** 🟡 PENDING APPROVAL
**Submitted By:** Architecture Planning
**Decision Needed By:** Project Owner

**This is a proposal only. No implementation will begin without explicit approval.**