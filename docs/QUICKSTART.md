# Swiper Quick Start Guide

## What is Swiper?

Swiper is a high-performance PDF extraction tool written in Go that converts PDF documents into markdown files with embedded images. It supports parallel processing for efficient extraction of text and images from PDF files.

## Quick Setup

### 1. Install Dependencies

**macOS:**
```bash
brew install poppler
```

**Ubuntu/Debian:**
```bash
sudo apt-get install poppler-utils
```

**CentOS/RHEL:**
```bash
sudo yum install poppler-utils
```

### 2. Build Swiper

```bash
cd /Users/macadelic/dusk-labs/utils/swiper
./verify-and-build.sh
```

The script will:
- Verify Go installation (requires Go 1.24+)
- Check for required system dependencies
- Download Go dependencies
- Run tests
- Build the optimized binary

## Usage Examples

### Extract a Single PDF

```bash
./swiper -file document.pdf
```

Output: Creates `document/` directory with:
- `index.md` - Main index linking to all pages
- `page_1.md`, `page_2.md`, etc. - Individual page markdown files
- `images/` - Extracted images

### Extract with Custom Output Directory

```bash
./swiper -file report.pdf -output my-output
```

### Batch Process Multiple PDFs

```bash
./swiper -dir /path/to/pdfs -output extracted-pdfs
```

### Scan and Collect PDFs

```bash
./swiper -scan /path/to/scan -copydir pdf-collection
```

Recursively finds all PDFs and copies them to `pdf-collection/`.

### Performance Options

```bash
# Use 8 parallel workers
./swiper -file document.pdf -processes 8

# Enable result caching
./swiper -file document.pdf -cache

# Run with performance benchmarking
./swiper -file document.pdf -benchmark

# CPU profiling
./swiper -file document.pdf -profile cpu
```

## Performance Tips

1. **Worker Count:** Use `-processes N` to match your CPU cores
   - Default: Auto-detected based on CPU count
   - More workers = faster for large PDFs
   - Diminishing returns beyond 8-16 workers

2. **Caching:** Enable `-cache` for repeated extractions
   - Useful when re-processing same PDFs
   - Stores text and image extraction results

3. **Large PDFs:** Swiper automatically optimizes for PDF size
   - Small PDFs (<50 pages): Buffer all pages
   - Large PDFs (>500 pages): Adaptive worker scaling
   - Huge PDFs (>1000 pages): Maximum buffering with capping

## Architecture

### Current Structure
- **Single binary:** All code in `main.go` (1,776 lines)
- **High performance:** Concurrent processing, buffer pooling, metrics collection
- **Status:** Production-ready, fully functional

### Proposed Architecture (Under Review)
A modular refactor has been proposed but **not yet approved**:
- **8 internal packages:** Focused, testable components
- **Public API:** Optional library usage (`pkg/swiper`)
- **Better testing:** Unit and integration tests
- **10x maintainability:** Easier to understand and extend

**See PROPOSAL.md for complete details**

**Status:** 🟡 PENDING APPROVAL - No implementation yet

## Output Format

### Directory Structure
```
document/
├── index.md              # Main index
├── page_1.md            # Page 1 content
├── page_2.md            # Page 2 content
├── page_3.md            # ...
└── images/              # Extracted images
    ├── page_1_img_001.jpg
    ├── page_1_img_002.png
    └── page_2_img_001.jpg
```

### Markdown Format

**index.md:**
```markdown
# document.pdf - PDF Extract

This document contains the extracted content from `/path/to/document.pdf`.

## Pages

- [Page 1](page_1.md)
- [Page 2](page_2.md)
- [Page 3](page_3.md)
```

**page_X.md:**
```markdown
# Page 1

## Text Content

```
[Extracted text here]
```

## Images

![Image from page 1](images/page_1_img_001.jpg)
![Image from page 1](images/page_1_img_002.png)
```

## Performance Metrics

Swiper automatically tracks and reports:
- Pages processed
- Text extracted (MB)
- Images extracted (count)
- Total bytes processed
- Processing time per page/PDF
- Cache hit rates
- Buffer pool statistics
- Worker utilization
- Queue depth analysis

Example output:
```
==================================================
PERFORMANCE METRICS
==================================================
Pages processed: 150
Text extracted: 2.34 MB
Images extracted: 45
Total bytes processed: 15.67 MB
Average processing time per page: 125ms
Total processing time: 18.75s
Cache hit rate: 0.00%
Buffer pool hits: 1,234, misses: 56 (95.65% hit rate)
Pages per second: 8.00
```

## Troubleshooting

### "Error: Missing required dependencies"
Install poppler-utils (see Install Dependencies section)

### "Error: PDF file not found"
Verify the PDF path is correct and file exists

### "Error: failed to run pdfinfo"
Ensure poppler-utils is installed and in your PATH

### Performance Issues
- Increase worker count: `-processes 16`
- Enable caching: `-cache`
- Check available memory for large PDFs
- Use SSD for better I/O performance

## Next Steps

1. **Try it out:** Extract your first PDF
2. **Review proposal:** Read `PROPOSAL.md` for proposed modular architecture
3. **Provide feedback:** Share thoughts on proposed changes
4. **Stay tuned:** Implementation will begin after approval

**Note:** A modular architecture has been proposed but **not yet implemented**. The current monolithic version is production-ready and will remain unchanged until/unless the proposal is approved.

## Resources

- **Architecture Proposal:** `PROPOSAL.md` (pending approval)
- **Build Script:** `verify-and-build.sh`
- **Documentation Index:** `INDEX.md`
- **Module Documentation:** `docs/doc.md`
- **Implementation Protocol:** `docs/implementation_proof_protocol/`

## Version Information

- **Go Version:** 1.24.0
- **Binary Size:** ~2.8MB
- **Current Status:** Monolithic, production-ready
- **Proposal Status:** Modular refactor proposed, awaiting approval
- **License:** Check repository root

---

**Questions?** Refer to `CLAUDE.md` for development guidelines.