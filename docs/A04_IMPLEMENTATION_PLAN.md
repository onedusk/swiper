# A04 Implementation Plan: Absolute Path Comparison for Directory Skip

**Feature ID:** A04
**Priority:** RESILIENCE | SECURITY
**Target Files:** `internal/scanner/scanner.go`
**Branch:** `feature/a04-absolute-path-comparison`
**Estimated Completion:** Single session
**Created:** 2025-10-06

---

## Executive Summary

Implement robust directory skipping using absolute path comparison to prevent:
- Infinite recursion when copyDir is inside scanDir
- Directory explosion (copying output into itself repeatedly)
- Relative/absolute path mismatches causing failures

**Current Issue:** Line 87-89 in `scanner.go` uses simple string comparison `path == s.copyDir` which fails when paths differ in representation (relative vs absolute, `.`, `..`, symlinks).

**Solution:** Resolve both paths to absolute canonical form using `filepath.Abs()` and compare, plus add subtree prefix checking.

---

## Documentation Sources

### Go Standard Library References
- **filepath.Abs()**: https://pkg.go.dev/path/filepath#Abs
  - Converts path to absolute form
  - Resolves `.` and `..` components
  - Does NOT follow symlinks (use `filepath.EvalSymlinks` for that)
  - Returns error if current directory cannot be determined

- **filepath.SkipDir**: https://pkg.go.dev/path/filepath#SkipDir
  - Special error value to skip directory in Walk/WalkDir
  - Must be returned exactly (not wrapped)

- **filepath.Separator**: https://pkg.go.dev/path/filepath#Separator
  - OS-specific separator ('/' on Unix, '\\' on Windows)
  - Use for building path prefixes

- **strings.HasPrefix()**: https://pkg.go.dev/strings#HasPrefix
  - Check if path is within directory tree
  - Pattern: `strings.HasPrefix(path, dir+string(filepath.Separator))`

### Project Standards
- **Error Handling**: Wrap errors with `fmt.Errorf("context: %w", err)`
- **Logging**: Use `s.logAsync()` for async non-blocking logs
- **Field Naming**: camelCase for private fields, document public fields
- **Testing**: Table-driven tests, integration tests in separate files

---

## Component Relationships

### Current Architecture
```
PDFScanner struct (scanner.go:18-25)
├── scanDir string          (user-provided, may be relative)
├── copyDir string          (user-provided, may be relative)
├── logChan chan string     (async logging)
├── bufferManager           (buffer pool)
└── metricsCollector        (metrics)

New() function (scanner.go:28)
├── Validates scanDir
├── Creates copyDir with os.MkdirAll
├── Initializes PDFScanner
└── Starts asyncLogger goroutine

FindPDFs() function (scanner.go:76)
├── Calls filepath.Walk
├── Filters for .pdf extension
├── Skips copyDir (CURRENTLY BROKEN)
└── Returns []string of PDF paths
```

### Proposed Architecture
```
PDFScanner struct (scanner.go:18-27)
├── scanDir string          (user-provided, preserved for display)
├── copyDir string          (user-provided, preserved for display)
├── scanDirAbs string       (NEW: canonical absolute path)
├── copyDirAbs string       (NEW: canonical absolute path)
├── logChan chan string
├── bufferManager
└── metricsCollector

New() function (scanner.go:28)
├── Validates scanDir
├── Resolves scanDir → scanDirAbs (filepath.Abs)
├── Resolves copyDir → copyDirAbs (filepath.Abs)
├── Creates copyDirAbs with os.MkdirAll
├── Initializes PDFScanner (with new fields)
└── Starts asyncLogger goroutine

FindPDFs() function (scanner.go:76)
├── Calls filepath.WalkDir (or filepath.Walk in current)
├── For each directory entry:
│   ├── Skip if path == copyDirAbs (exact match)
│   ├── Skip if path is subtree of copyDirAbs (prefix check)
│   └── Log skip action
├── Filters for .pdf extension
└── Returns []string of PDF paths
```

---

## Dependency Mapping

### External Dependencies
```go
// Required imports in scanner.go
import (
    "fmt"          // Error formatting
    "io"           // File I/O
    "log"          // Logging (fallback)
    "os"           // File operations, stat
    "path/filepath" // Path manipulation, Abs, Separator
    "strings"      // String operations, HasPrefix
    "sync"         // Mutex, atomic
    "sync/atomic"  // Atomic operations
    "time"         // Time tracking

    "swiper/internal/metrics" // Metrics collection
    "swiper/internal/pool"    // Buffer pool
)
```

### Internal Dependencies
- **metrics.Collector**: Already imported and used
- **pool.BufferPoolManager**: Already imported and used
- No new dependencies required

### Dependency Graph
```
scanner.go
├── DEPENDS ON: path/filepath (stdlib)
│   ├── Abs()
│   ├── Separator
│   └── SkipDir
├── DEPENDS ON: strings (stdlib)
│   └── HasPrefix()
├── DEPENDS ON: fmt (stdlib)
│   └── Errorf()
└── DEPENDS ON: os (stdlib)
    └── MkdirAll()

No packages depend on scanner.go changes (internal only)
```

---

## Imports/Exports Analysis

### Current Exports (Public API)
```go
// scanner.go exports (unchanged):

type PDFScanner struct {
    // All fields are private (lowercase)
}

func New(scanDir, copyDir string) (*PDFScanner, error)
    // Creates new scanner instance
    // SIGNATURE UNCHANGED

func (s *PDFScanner) FindPDFs() ([]string, error)
    // Finds all PDFs in scanDir
    // SIGNATURE UNCHANGED
    // BEHAVIOR IMPROVED (no breaking changes)

func (s *PDFScanner) ScanAndCopy() error
    // Main entry point
    // SIGNATURE UNCHANGED

// Private methods (not exported):
// - asyncLogger()
// - logAsync()
// - copyPDFWithProgress()
```

### Impact on External Callers
```
cmd/swiper/main.go (runScanner function)
└── Calls scanner.New(scanDir, copyDir)
    └── NO CHANGES REQUIRED (signature unchanged)

No other packages import internal/scanner
└── Internal package, not part of public API
```

**Conclusion:** Zero breaking changes. All modifications are internal implementation details.

---

## Target Hitlist

### Files to Modify
1. ✅ `internal/scanner/scanner.go`
   - Lines 18-25: Add new fields to PDFScanner struct
   - Lines 28-58: Update New() to resolve absolute paths
   - Lines 76-104: Update FindPDFs() with absolute path logic

### Files to Create
2. ✅ `testdata/a04-nested/test.pdf`
   - Test fixture for integration testing
   - Can be minimal valid PDF

3. ✅ `testdata/a04-nested/output/.gitkeep`
   - Ensures output directory structure in git

### Files to Update (Testing)
4. ✅ `docs/CHANGELOG.md`
   - Add entry under [Unreleased] → Changed section

5. ⚠️ `internal/scanner/scanner_test.go` (optional, may not exist)
   - If exists: verify tests still pass
   - If not exists: consider creating basic test

---

## Detailed Code Changes

### Change 1: Add Absolute Path Fields
**File:** `internal/scanner/scanner.go`
**Location:** Lines 18-25 (struct definition)

**BEFORE:**
```go
// PDFScanner handles scanning and copying PDF files
type PDFScanner struct {
	scanDir          string
	copyDir          string
	logChan          chan string
	bufferManager    *pool.BufferPoolManager
	metricsCollector *metrics.Collector
}
```

**AFTER:**
```go
// PDFScanner handles scanning and copying PDF files
type PDFScanner struct {
	scanDir          string  // User-provided scan directory (may be relative)
	copyDir          string  // User-provided copy directory (may be relative)
	scanDirAbs       string  // Absolute path of scan directory
	copyDirAbs       string  // Absolute path of copy directory
	logChan          chan string
	bufferManager    *pool.BufferPoolManager
	metricsCollector *metrics.Collector
}
```

**Rationale:** Store both user-provided paths (for display/error messages) and canonical absolute paths (for comparison).

---

### Change 2: Resolve Absolute Paths in New()
**File:** `internal/scanner/scanner.go`
**Location:** Lines 28-58 (New function)

**BEFORE:**
```go
func New(scanDir, copyDir string) (*PDFScanner, error) {
	// Use current directory if not specified
	if scanDir == "" {
		scanDir = "."
	}

	// Default copy directory
	if copyDir == "" {
		copyDir = "pdf-docs"
	}

	// Create the copy directory if it doesn't exist
	if err := os.MkdirAll(copyDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create copy directory: %v", err)
	}

	metricsCollector := metrics.NewCollector()
	logChan := make(chan string, 100)
	scanner := &PDFScanner{
		scanDir:          scanDir,
		copyDir:          copyDir,
		logChan:          logChan,
		bufferManager:    pool.NewBufferPoolManager(metricsCollector),
		metricsCollector: metricsCollector,
	}

	// Start async logger
	go scanner.asyncLogger()

	return scanner, nil
}
```

**AFTER:**
```go
func New(scanDir, copyDir string) (*PDFScanner, error) {
	// Use current directory if not specified
	if scanDir == "" {
		scanDir = "."
	}

	// Default copy directory
	if copyDir == "" {
		copyDir = "pdf-docs"
	}

	// Resolve to absolute paths for robust comparison
	scanDirAbs, err := filepath.Abs(scanDir)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve scan directory: %w", err)
	}
	copyDirAbs, err := filepath.Abs(copyDir)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve copy directory: %w", err)
	}

	// Create the copy directory if it doesn't exist
	if err := os.MkdirAll(copyDirAbs, 0755); err != nil {
		return nil, fmt.Errorf("failed to create copy directory: %v", err)
	}

	metricsCollector := metrics.NewCollector()
	logChan := make(chan string, 100)
	scanner := &PDFScanner{
		scanDir:          scanDir,
		copyDir:          copyDir,
		scanDirAbs:       scanDirAbs,
		copyDirAbs:       copyDirAbs,
		logChan:          logChan,
		bufferManager:    pool.NewBufferPoolManager(metricsCollector),
		metricsCollector: metricsCollector,
	}

	// Start async logger
	go scanner.asyncLogger()

	return scanner, nil
}
```

**Changes:**
- Added `filepath.Abs(scanDir)` → `scanDirAbs` with error handling
- Added `filepath.Abs(copyDir)` → `copyDirAbs` with error handling
- Changed `os.MkdirAll(copyDir, ...)` → `os.MkdirAll(copyDirAbs, ...)`
- Added `scanDirAbs` and `copyDirAbs` to struct initialization

---

### Change 3: Update FindPDFs() Directory Skip Logic
**File:** `internal/scanner/scanner.go`
**Location:** Lines 76-104 (FindPDFs function)

**BEFORE:**
```go
// FindPDFs recursively finds all PDF files in the scan directory
func (s *PDFScanner) FindPDFs() ([]string, error) {
	var pdfFiles []string

	err := filepath.Walk(s.scanDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			s.logAsync("Error accessing path %s: %v", path, err)
			return nil // Continue walking
		}

		// Skip the copy directory to avoid recursion
		if info.IsDir() && path == s.copyDir {
			return filepath.SkipDir
		}

		// Check if it's a PDF file
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".pdf") {
			pdfFiles = append(pdfFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %v", err)
	}

	return pdfFiles, nil
}
```

**AFTER:**
```go
// FindPDFs recursively finds all PDF files in the scan directory
func (s *PDFScanner) FindPDFs() ([]string, error) {
	var pdfFiles []string

	err := filepath.Walk(s.scanDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			s.logAsync("Error accessing path %s: %v", path, err)
			return nil // Continue walking
		}

		// Skip the copy directory to avoid recursion (absolute path comparison)
		if info.IsDir() {
			// Resolve current path to absolute for comparison
			pathAbs, absErr := filepath.Abs(path)
			if absErr == nil {
				// Check exact match (copyDir itself)
				if pathAbs == s.copyDirAbs {
					s.logAsync("Skipping destination directory: %s", path)
					return filepath.SkipDir
				}
				// Check if path is within copyDir subtree
				if strings.HasPrefix(pathAbs, s.copyDirAbs+string(filepath.Separator)) {
					s.logAsync("Skipping destination subdirectory: %s", path)
					return filepath.SkipDir
				}
			}
		}

		// Check if it's a PDF file
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".pdf") {
			pdfFiles = append(pdfFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %v", err)
	}

	return pdfFiles, nil
}
```

**Changes:**
- Added `filepath.Abs(path)` for current directory being walked
- Changed comparison from `path == s.copyDir` → `pathAbs == s.copyDirAbs`
- Added prefix check: `strings.HasPrefix(pathAbs, s.copyDirAbs+separator)`
  - Ensures we skip not just copyDir but all subdirectories under it
- Added logging for both skip cases
- Gracefully handles `Abs()` errors (continue walking if can't resolve)

**Edge Cases Handled:**
1. `scanDir="."` and `copyDir="output"` → Both resolved to absolute
2. `scanDir="/path/to/scan"` and `copyDir="output"` → copyDir resolved relative to CWD
3. `copyDir` is subdirectory of `scanDir` → Detected and skipped with subtree check
4. Symlinks in path → Not followed by `Abs()`, but compared correctly
5. Windows path separators → `filepath.Separator` is platform-aware

---

## Test Plan

### Unit Test Scenarios
1. **Exact Path Match**
   ```bash
   scanDir="/tmp/scan"
   copyDir="/tmp/scan/output"
   Expected: output directory skipped, logged
   ```

2. **Relative Path Resolution**
   ```bash
   scanDir="testdata/a04-nested"
   copyDir="testdata/a04-nested/output"
   Expected: Resolves to absolute, skips correctly
   ```

3. **Current Directory**
   ```bash
   cd testdata/a04-nested
   scanDir="."
   copyDir="./output"
   Expected: Both resolve, skip works
   ```

4. **Absolute vs Relative Mix**
   ```bash
   scanDir="$(pwd)/testdata"
   copyDir="testdata/output"
   Expected: Both resolve to same base, comparison works
   ```

### Integration Test Commands
```bash
# Setup
mkdir -p testdata/a04-nested/output
echo "test" > testdata/a04-nested/test.pdf

# Test 1: Relative paths
./swiper -scan testdata/a04-nested -copydir testdata/a04-nested/output
# Expected: Finds test.pdf, copies once, skips output dir

# Test 2: Absolute paths
./swiper -scan "$(pwd)/testdata/a04-nested" -copydir "$(pwd)/testdata/a04-nested/output"
# Expected: Same as Test 1

# Test 3: Mixed paths
./swiper -scan "$(pwd)/testdata/a04-nested" -copydir testdata/a04-nested/output
# Expected: Same as Test 1

# Validation
count=$(find testdata/a04-nested/output -name "*.pdf" | wc -l)
if [ "$count" -eq 1 ]; then
    echo "PASS: Exactly 1 PDF copied"
else
    echo "FAIL: Expected 1, got $count"
fi
```

### Performance Test
```bash
# Verify no infinite loop
timeout 5s ./swiper -scan testdata/a04-nested -copydir testdata/a04-nested/output
# Expected: Completes in <2s, exits 0
```

---

## Validation Checklist

### Code Quality
- [ ] No compilation errors: `go build ./cmd/swiper`
- [ ] No vet warnings: `go vet ./internal/scanner`
- [ ] No race conditions: `go test -race ./internal/scanner`
- [ ] Consistent formatting: `go fmt ./internal/scanner`

### Functionality
- [ ] Relative paths work correctly
- [ ] Absolute paths work correctly
- [ ] Mixed relative/absolute paths work
- [ ] Logs show "Skipping destination directory" message
- [ ] No infinite recursion (<2s completion)
- [ ] Exactly 1 copy of each PDF (no duplicates)

### Backwards Compatibility
- [ ] Existing tests pass: `go test ./...`
- [ ] API signatures unchanged
- [ ] No breaking changes to public interfaces

### Documentation
- [ ] CHANGELOG.md updated with change description
- [ ] Code comments added for new fields
- [ ] This implementation plan marked complete

---

## Rollback Strategy

### If Tests Fail
```bash
# Discard all changes
git checkout -- internal/scanner/scanner.go

# Or revert specific commit
git revert <commit-sha>
```

### If Performance Regresses
```bash
# filepath.Abs() is fast (~100ns), but if issues arise:
# - Cache absolute paths in struct (already doing this)
# - Skip Abs() call for already-absolute paths (optimization)
# - Profile with: go test -cpuprofile=cpu.prof -bench=.
```

### If Bugs Found in Production
```bash
# Emergency fix: revert to simple comparison
if info.IsDir() && path == s.copyDir {
    return filepath.SkipDir
}
# Then investigate why Abs() comparison failed
```

---

## Success Criteria

1. ✅ All tests pass
2. ✅ No infinite recursion demonstrated
3. ✅ Logs show skip messages
4. ✅ Performance unchanged (<1% regression acceptable)
5. ✅ PR approved and merged
6. ✅ CHANGELOG updated

---

## Implementation Checklist (22 Steps)

- [ ] 1. Commit CHANGELOG update for PR #2
- [ ] 2. Create feature branch: feature/a04-absolute-path-comparison
- [ ] 3. Add scanDirAbs and copyDirAbs fields to PDFScanner struct
- [ ] 4. Add filepath.Abs() resolution in New() for scanDir
- [ ] 5. Add filepath.Abs() resolution in New() for copyDir
- [ ] 6. Update New() to use copyDirAbs for os.MkdirAll
- [ ] 7. Add absolute path comparison in FindPDFs() for exact match
- [ ] 8. Add prefix check in FindPDFs() to skip copyDir subtree
- [ ] 9. Add logging for skipped destination directory
- [ ] 10. Build and verify no compilation errors
- [ ] 11. Create test directory structure: testdata/a04-nested/output
- [ ] 12. Create test PDF file in testdata/a04-nested/
- [ ] 13. Test: scan with copydir as subdirectory (relative paths)
- [ ] 14. Test: scan with copydir as subdirectory (absolute paths)
- [ ] 15. Test: verify log shows 'Skipping destination directory' message
- [ ] 16. Test: verify no infinite recursion (completes in <2s)
- [ ] 17. Test: verify output contains only 1 copy of test PDF
- [ ] 18. Run go test ./internal/scanner
- [ ] 19. Create commit with descriptive message
- [ ] 20. Push feature branch to remote
- [ ] 21. Create PR for A04 implementation
- [ ] 22. Merge and verify A04 PR

---

## References

- **Atomic Improvement Plan**: `/Users/macadelic/dracos/utils/swiper/docs/ATOMIC_IMPROVEMENT_PLAN.md`
- **Go filepath Package**: https://pkg.go.dev/path/filepath
- **Project CLAUDE.md**: `/Users/macadelic/dracos/utils/swiper/CLAUDE.md`
- **Previous PR #2**: https://github.com/onedusk/swiper/pull/2

---

**Status:** Ready for implementation
**Estimated Time:** 30-45 minutes
**Risk Level:** Low
**Breaking Changes:** None
