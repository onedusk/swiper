# Contributing to swiper

First off, thank you for considering contributing to swiper! Every contribution helps make this PDF-to-markdown converter better.

## How to Contribute

We welcome contributions in the form of bug reports, feature requests, documentation improvements, and pull requests.

### Reporting Issues

If you encounter a bug or have a feature idea, please [open an issue](https://github.com/onedusk/swiper/issues) on our GitHub repository. Provide as much detail as possible, including:

- A clear and descriptive title
- Steps to reproduce the bug (if applicable)
- The expected behavior and what actually happened
- Your operating system and swiper version
- Sample PDF files (if the issue is conversion-related)

### Submitting Pull Requests

1. **Fork the repository** and create a new branch from `main`
2. **Set up your local development environment**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/swiper.git
   cd swiper
   go mod download
   ```
3. **Make your changes** and ensure they follow the project's style
4. **Format your code** using `gofmt`:
   ```bash
   gofmt -w .
   ```
5. **Add or update tests** to cover your changes:
   ```bash
   go test ./...
   ```
6. **Commit your changes** using the Conventional Commits format (see below)
7. **Push your branch** to your fork and open a pull request

## Local Development Setup

### Prerequisites

- Go 1.24 or later
- poppler-utils installed (`brew install poppler` on macOS)
- Git

### Building

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/swiper.git
cd swiper

# Install dependencies
go mod download

# Build the binary
go build -o swiper cmd/swiper/main.go

# Run your local build
./swiper -help
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

### Testing Your Changes

Before submitting a PR, test your changes with:

```bash
# Single file conversion
./swiper -file testdata/sample.pdf

# Directory processing
./swiper -dir testdata/pdfs -output ./test-output

# Benchmark mode
./swiper -file testdata/sample.pdf -benchmark
```

## Code Style

We follow the standard Go community style guidelines:

- Use `gofmt` for formatting (run before committing)
- Follow Go naming conventions (PascalCase for exported, camelCase for private)
- Write clear, descriptive variable and function names
- Add comments for exported functions and types
- Keep functions focused and concise

### Code Organization

The codebase is modular:

- `cmd/swiper/` - CLI entry point
- `pkg/swiper/` - Public library interface
- `internal/` - Internal implementation packages
  - `internal/extractor/` - PDF extraction logic
  - `internal/buffer/` - Buffer pool management
  - `internal/worker/` - Worker pool implementation

## Commit Message Format

We use the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification. Each commit message should be prefixed with a type:

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, no logic change)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks
- `perf`: Performance improvements

**Examples:**

```
feat: add JSON output format for batch processing
fix: correct handling of unicode characters in PDF text
docs: update installation instructions for Windows
refactor: extract buffer pool logic into separate package
test: add integration tests for directory processing
perf: optimize worker pool allocation
```

**Format:**
```
<type>: <short description>

<optional longer description>

<optional footer>
```

## Testing Guidelines

### Unit Tests

- Test files should be named `*_test.go` and placed alongside the source files
- Use table-driven tests for multiple test cases
- Test both success and failure scenarios
- Mock external dependencies (like poppler-utils calls) when appropriate

**Example:**
```go
func TestExtractSingle(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        wantErr bool
    }{
        {"valid PDF", "testdata/valid.pdf", false},
        {"invalid file", "nonexistent.pdf", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test logic
        })
    }
}
```

### Integration Tests

- Place integration tests in separate files or directories
- Use real PDF files from `testdata/`
- Test end-to-end workflows

## Pull Request Process

1. **Update documentation** if you're adding features or changing behavior
2. **Add tests** for new functionality
3. **Ensure all tests pass** (`go test ./...`)
4. **Update CHANGELOG.md** if applicable (for significant changes)
5. **Request review** from maintainers
6. **Address feedback** promptly and professionally
7. **Squash commits** if requested before merging

### PR Title Format

Use the same Conventional Commits format for PR titles:

```
feat: add support for encrypted PDFs
fix: resolve memory leak in worker pool
docs: improve batch processing examples
```

## Development Workflow

### Feature Development

```bash
# Create a feature branch
git checkout -b feat/my-new-feature

# Make changes and commit
git add .
git commit -m "feat: add my new feature"

# Push and create PR
git push -u origin feat/my-new-feature
```

### Bug Fixes

```bash
# Create a bugfix branch
git checkout -b fix/issue-123

# Fix the bug and add test
git add .
git commit -m "fix: resolve issue with XYZ

Closes #123"

# Push and create PR
git push -u origin fix/issue-123
```

## Performance Considerations

When contributing performance-related changes:

- Use the `-benchmark` flag to measure impact
- Profile with `-profile cpu` or `-profile memory`
- Document performance improvements in the PR description
- Consider memory usage alongside speed

## Documentation

When adding features:

- Update the README.md if it affects user-facing behavior
- Add inline code comments for complex logic
- Update relevant documentation in `docs/`
- Include usage examples

## Getting Help

If you need help or have questions:

- Check existing [issues](https://github.com/onedusk/swiper/issues)
- Review the [documentation](docs/)
- Ask in your PR or issue thread
- Be patient and respectful

## Code of Conduct

This project adheres to the Contributor Covenant Code of Conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

## License

By contributing, you agree that your contributions will be licensed under the same license as the project (see LICENSE file).

---

Thank you for contributing to swiper!
