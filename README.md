# swiper

A blazingly fast PDF-to-markdown converter written in Go. Extract text and images from single files or batch process entire directories with parallel processing.

## Overview

`swiper` is a command-line tool designed for efficient conversion of PDF files to markdown format with embedded images. Built in Go, it leverages worker pools for parallel processing to maximize performance, making it ideal for large-scale document conversion tasks.

## Features

- **High Performance**: Parallel page processing at ~8 pages/second
- **Advanced Buffer Pooling**: 4-tier buffer management (32KB-1MB) for optimal memory usage
- **Flexible Input**: Process a single PDF file or an entire directory recursively
- **Result Caching**: Optional caching for faster re-processing
- **Performance Metrics**: Built-in benchmarking and profiling capabilities
- **Async Logging**: Non-blocking log operations for better performance
- **Cross-Platform**: Works on Linux, macOS, and Windows

## Quick Start

Extract markdown from a single PDF:

```bash
swiper -file ./document.pdf
```

Batch process a directory of PDFs:

```bash
swiper -dir ./pdfs -output ./markdown
```

## Installation

### Prerequisites

Swiper requires `poppler-utils` for PDF processing.

**macOS:**
```bash
brew install poppler
```

**Ubuntu/Debian:**
```bash
sudo apt-get update
sudo apt-get install poppler-utils
```

**Red Hat/CentOS/Fedora:**
```bash
sudo yum install poppler-utils
```

**Windows:**
Download and install poppler from [poppler for Windows](https://github.com/oschwartz10612/poppler-windows/releases) and add to PATH.

### Install Swiper

**Option 1: Install with Go**
```bash
go install github.com/onedusk/swiper@latest
```

**Option 2: Download Pre-built Binary**
Download from the [Releases](https://github.com/onedusk/swiper/releases) page for your platform.

**Option 3: Build from Source**
```bash
git clone https://github.com/onedusk/swiper.git
cd swiper
go build -o swiper cmd/swiper/main.go
```

### Verify Installation

```bash
swiper -help
```

Should display the help menu with all available flags.

## Usage

### `-file`

Extract content from a single PDF file and convert it to markdown.

**Syntax:**
```bash
swiper -file <path/to/input.pdf>
```

By default, output is saved in the same directory as the input file with `.md` extension. Combine with `-output` to specify a custom output directory.

**Examples:**
```bash
# Extract single PDF
swiper -file report.pdf

# Extract with custom output location
swiper -file report.pdf -output ./extracted/
```

**Output:** Creates `report.md` with embedded images extracted to a subdirectory.

### `-dir`

Batch process all PDF files in a directory recursively.

**Syntax:**
```bash
swiper -dir <path/to/directory>
```

Swiper will:
- Recursively find all `.pdf` files
- Process them in parallel using worker pools
- Handle individual file errors gracefully
- Report progress and statistics

**Examples:**
```bash
# Process all PDFs in directory
swiper -dir ./documents

# Process with custom output directory
swiper -dir ./documents -output ./markdown-output
```

**Performance:** Processes ~8 pages/second with parallel workers.

### `-output`

Specify the output directory for extracted markdown and images.

**Syntax:**
```bash
swiper -output <path/to/output>
```

**Behavior:**
- For single files: Creates markdown file in the output directory
- For directory batch: Preserves directory structure in output
- Creates subdirectories for images automatically
- Output directory is created if it doesn't exist

**Examples:**
```bash
# Single file with custom output
swiper -file doc.pdf -output ./results

# Batch processing with organized output
swiper -dir ./pdfs -output ./markdown
```

### `-processes`

Control the number of parallel worker goroutines for processing.

**Syntax:**
```bash
swiper -processes <count>
```

**Default:** Automatically set to `runtime.NumCPU()` (number of CPU cores)

**When to Adjust:**
- **Increase (8-16+):** High-performance servers with fast storage
- **Decrease (1-4):** Low-memory systems or resource-constrained environments
- **Single-threaded (1):** For debugging or minimal resource usage

**Examples:**
```bash
# Use 4 workers
swiper -dir ./pdfs -processes 4

# Maximum parallelism (12 cores)
swiper -dir ./large-dataset -processes 12

# Single-threaded processing
swiper -dir ./pdfs -processes 1
```

**Performance Impact:** Higher worker counts increase throughput but consume more memory (approximately 32KB-1MB per worker for buffer pools).

## Comparison with Alternatives

| Feature                  | swiper (Go)          | pdftotext (C++)      | Apache Tika (Java)   | PyPDF2 (Python)      | mupdf (C)            |
| ------------------------ | -------------------- | -------------------- | -------------------- | -------------------- | -------------------- |
| **Language**             | Go                   | C++                  | Java                 | Python               | C                    |
| **Dependencies**         | Poppler-utils        | Poppler-utils        | JVM                  | None (pure Python)   | None                 |
| **Speed**                | **Very High**        | High                 | Moderate             | Low                  | Very High            |
| **Ease of Use**          | **Very High**        | High                 | Moderate             | High                 | Moderate             |
| **Parallel Processing**  | **Built-in**         | Manual (via shell)   | Yes (configurable)   | Manual (via script)  | Manual (via shell)   |
| **Output Formats**       | `markdown`, `text`   | `text`               | `text`, `html`, `xml` | `text`               | `text`, `html`, `svg` |
| **Installation**         | Single binary        | Package manager      | JAR download / Maven | `pip install`        | Build from source    |
| **Markdown Support**     | **Native**           | No                   | No                   | No                   | No                   |

**Key Advantages:**
- Native markdown output with embedded images
- Built-in parallel processing without shell scripting
- Single binary deployment with no runtime dependencies
- Advanced buffer pooling for optimal memory usage

## Advanced Batch Processing Examples

### 1. Process a Directory to Individual Markdown Files

Process all PDFs in the `invoices/` directory and save markdown files to `invoices_markdown/`:

```bash
swiper -dir ./invoices -output ./invoices_markdown
```

### 2. Process with Custom Worker Count

Use only 2 workers for resource-constrained environments:

```bash
swiper -dir ./large_archive -output ./archive_markdown -processes 2
```

### 3. Enable Caching for Faster Re-Processing

Use result caching to speed up repeated conversions:

```bash
swiper -dir ./reports -output ./reports_markdown -cache
```

### 4. Benchmark Mode for Performance Testing

Run with detailed performance metrics:

```bash
swiper -dir ./test_pdfs -benchmark
```

### 5. Real-World Pipeline Integration

Integrate swiper into a document processing pipeline:

```bash
#!/bin/bash

INPUT_DIR="./new_documents"
OUTPUT_DIR="./processed_markdown"
LOG_FILE="processing.log"

echo "Starting PDF conversion..."
# Run swiper with progress logging
swiper -dir "$INPUT_DIR" -output "$OUTPUT_DIR" -processes 8 2> "$LOG_FILE"

echo "Conversion complete. Indexing markdown files..."
# Index the converted files
find "$OUTPUT_DIR" -name "*.md" > markdown_index.txt

echo "Done. See $LOG_FILE for processing details."
```

### 6. Error Handling and Logging

Separate stdout and stderr for better error tracking:

```bash
# Process directory and log errors separately
swiper -dir ./mixed_quality_pdfs -output ./extracted 2> errors.log

# Check for failed conversions
if [ -s errors.log ]; then
    echo "Some PDFs failed to convert. See errors.log"
else
    echo "All PDFs converted successfully"
fi
```

## Library Usage

Swiper is also available as a Go library:

```go
import "github.com/onedusk/swiper/pkg/swiper"

func main() {
    client, err := swiper.NewClient(
        swiper.WithProcessCount(8),
        swiper.WithOutputDir("output"),
        swiper.WithCache(true),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()
    result, err := client.ExtractSingle(ctx, "document.pdf")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Converted: %s\n", result.OutputPath)
}
```

## Performance

- **Processing Speed**: ~8 pages/second on modern hardware
- **Memory Usage**: Configurable with buffer pools (32KB-1MB per worker)
- **Concurrency**: Scales with CPU cores (configurable via `-processes`)
- **Benchmark Mode**: Use `-benchmark` flag for detailed performance metrics

## Requirements

- Go 1.24+ (for building from source)
- poppler-utils (pdfinfo, pdftotext, pdfimages)

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for:
- How to submit pull requests
- Code style guidelines
- Testing requirements
- Development setup instructions

## Code of Conduct

This project adheres to the Contributor Covenant Code of Conduct. See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details.

## License

Check repository root for license information.

## Documentation

- **[QUICKSTART.md](docs/QUICKSTART.md)** - Complete user guide
- **[PROPOSAL.md](docs/PROPOSAL.md)** - Architecture proposal (IMPLEMENTED)
- **[INDEX.md](docs/INDEX.md)** - Documentation index
- **[CLAUDE.md](CLAUDE.md)** - Development guidelines

---

**Questions?** Open an issue or see the documentation links above.
