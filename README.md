# Swiper

> Convert PDF documents to markdown with embedded images using parallel processing.

## Documentation

- **[QUICKSTART.md](docs/QUICKSTART.md)** - Complete user guide
- **[PROPOSAL.md](docs/PROPOSAL.md)** - Architecture proposal (IMPLEMENTED)
- **[INDEX.md](docs/INDEX.md)** - Documentation index
- **[CLAUDE.md](CLAUDE.md)** - Development guidelines

## Features

- Parallel page processing
- Advanced buffer pooling
- Performance metrics
- Result caching
- Batch processing
- PDF scanning and copying

## Performance

- ~8 pages/second
- Configurable worker count
- 4-tier buffer management (32KB-1MB)
- Async logging

## Quick Start

### Install Dependencies

**macOS:**
```bash
brew install poppler
```

**Ubuntu/Debian:**
```bash
sudo apt-get install poppler-utils
```

### Build

```bash
# Build modular version
go build -o swiper-new cmd/swiper/main.go

# Or build monolithic version (preserved)
go build -o swiper main.go
```

### Use

```bash
# Using new modular version
./swiper-new -file document.pdf

# Batch processing
./swiper-new -dir /path/to/pdfs

# Scan and copy PDFs
./swiper-new -scan /path/to/scan
```

## Requirements

- Go 1.24+
- poppler-utils (pdfinfo, pdftotext, pdfimages)

## Building

```bash
# Automated verification and build
./verify-and-build.sh

# Manual build
go build -o swiper main.go
```

## Library Usage

Now available as a Go library:

```go
import "swiper/pkg/swiper"

func main() {
    client, err := swiper.NewClient(
        swiper.WithProcessCount(8),
        swiper.WithOutputDir("output"),
    )

    ctx := context.Background()
    result, err := client.ExtractSingle(ctx, "document.pdf")
}
```

## Contributing

The codebase is now modular and much easier to contribute to. Each package has clear responsibilities and can be tested independently.

## License

Check repository root for license information.

---

**Questions?** See [QUICKSTART.md](QUICKSTART.md) or [INDEX.md](INDEX.md)
