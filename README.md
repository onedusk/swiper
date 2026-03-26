# Swiper

High-performance PDF to markdown converter with parallel processing.

## Installation

```bash
go install github.com/onedusk/swiper/cmd/swiper@latest
```

Requires [poppler-utils](https://poppler.freedesktop.org/) for PDF processing:

```bash
# macOS
brew install poppler

# Debian/Ubuntu
sudo apt install poppler-utils

# Fedora
sudo dnf install poppler-utils
```

## Quick Start

```bash
# Extract a single PDF
swiper -file document.pdf

# Extract specific pages
swiper -file document.pdf -pages 1-10,50

# Batch process a directory of PDFs
swiper -dir /path/to/pdfs -output extracted

# Resume interrupted batch processing
swiper -dir /path/to/pdfs -output extracted -resume

# Scan and copy PDFs from a directory tree
swiper -scan /path/to/search -copydir pdf-collection
```

## Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-file <path>` | Path to a single PDF file | |
| `-dir <path>` | Directory of PDFs for batch processing | |
| `-scan <path>` | Scan directory tree for PDFs | |
| `-output <dir>` | Output directory | PDF name or `extracted-pdfs` |
| `-pages <range>` | Page range filter (e.g., `1-10,50-60,75`) | all pages |
| `-processes <n>` | Number of worker processes | CPU count |
| `-q` | Suppress non-error output | `false` |
| `-resume` | Resume interrupted batch processing | `false` |
| `-copydir <dir>` | Directory to copy scanned PDFs | `pdf-docs` |
| `-cache` | Enable result caching | `false` |
| `-config <path>` | YAML configuration file | |
| `-profile cpu\|mem` | Enable CPU or memory profiling | |
| `-benchmark` | Detailed performance metrics | `false` |

## Output Format

Each PDF produces a directory with markdown files and extracted images:

```
output/
├── index.md          # Links to all page files
├── page_1.md         # Markdown for page 1
├── page_2.md         # Markdown for page 2
└── images/
    ├── page_1_img_001.jpg
    └── page_2_img_001.png
```

## Requirements

- Go 1.24+
- poppler-utils (`pdfinfo`, `pdftotext`, `pdfimages`)
