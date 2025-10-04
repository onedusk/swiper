# Gemini Project Context: Swiper

This document provides context for the "Swiper" project, a command-line tool written in Go for converting PDF documents to markdown with embedded images.

## Project Overview

Swiper is a high-performance PDF to markdown converter. It leverages parallel processing to speed up the extraction of text and images from PDF files.

**Key Features:**

*   **Parallel Processing:** Utilizes multiple CPU cores to process PDF pages concurrently.
*   **Batch Processing:** Can process an entire directory of PDF files.
*   **Result Caching:** Caches results to avoid reprocessing the same file.
*   **Performance Metrics:** Provides detailed performance metrics in benchmark mode.
*   **Go Library:** Can be used as a library in other Go projects.

**Architecture:**

The project follows a modular architecture with clear separation of concerns:

*   `cmd/swiper/main.go`: The main entry point of the application, handling command-line arguments and orchestrating the different modules.
*   `internal/`: Contains the core logic of the application.
    *   `extractor`: Handles the extraction of pages from a single PDF file.
    *   `scanner`: Scans a directory for PDF files.
    *   `batch`: Manages the batch processing of multiple PDF files.
    *   `config`: Manages the application's configuration.
*   `pkg/swiper`: Provides a public API for using Swiper as a Go library.

## Building and Running

The project uses standard Go tools for building and dependency management. It also has a dependency on `poppler-utils`.

**Dependencies:**

*   Go 1.24+
*   `poppler-utils` (`pdfinfo`, `pdftotext`, `pdfimages`)

**Build Commands:**

The `scripts/verify-and-build.sh` script is the recommended way to build the project. It verifies dependencies, runs checks, and builds the binary.

```bash
# Automated verification and build
./scripts/verify-and-build.sh
```

Alternatively, you can build manually:

```bash
# Build the modular version
go build -o swiper-new cmd/swiper/main.go
```

**Running the Application:**

```bash
# Process a single PDF
./swiper-new -file document.pdf

# Batch process a directory of PDFs
./swiper-new -dir /path/to/pdfs

# Scan a directory for PDFs and copy them
./swiper-new -scan /path/to/scan
```

## Development Conventions

*   **Coding Style:** The code follows standard Go formatting and conventions.
*   **Testing:** The project has a `scripts/verify-and-build.sh` script that includes a step for running tests, although no tests were found during the analysis.
*   **Dependencies:** Dependencies are managed using Go modules (`go.mod`).
