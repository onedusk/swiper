#!/bin/bash

# Swiper Verification and Build Script
# Verifies dependencies, runs checks, and builds the binary

set -e  # Exit on error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}Swiper Verification and Build${NC}"
echo -e "${BLUE}========================================${NC}\n"

# Step 1: Check Go installation
echo -e "${YELLOW}[1/7] Checking Go installation...${NC}"
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed${NC}"
    exit 1
fi
GO_VERSION=$(go version)
echo -e "${GREEN}✓ $GO_VERSION${NC}\n"

# Step 2: Check required system dependencies
echo -e "${YELLOW}[2/7] Checking poppler-utils (pdfinfo, pdftotext, pdfimages)...${NC}"
MISSING_DEPS=()

if ! command -v pdfinfo &> /dev/null; then
    MISSING_DEPS+=("pdfinfo")
fi
if ! command -v pdftotext &> /dev/null; then
    MISSING_DEPS+=("pdftotext")
fi
if ! command -v pdfimages &> /dev/null; then
    MISSING_DEPS+=("pdfimages")
fi

if [ ${#MISSING_DEPS[@]} -gt 0 ]; then
    echo -e "${RED}Error: Missing required dependencies: ${MISSING_DEPS[*]}${NC}"
    echo -e "${YELLOW}Install with:${NC}"
    if [[ "$OSTYPE" == "darwin"* ]]; then
        echo -e "  brew install poppler"
    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
        echo -e "  apt-get install poppler-utils  # Ubuntu/Debian"
        echo -e "  yum install poppler-utils      # CentOS/RHEL"
    fi
    exit 1
fi
echo -e "${GREEN}✓ All poppler-utils commands available${NC}\n"

# Step 3: Verify go.mod
echo -e "${YELLOW}[3/7] Verifying go.mod...${NC}"
if [ ! -f "go.mod" ]; then
    echo -e "${RED}Error: go.mod not found${NC}"
    exit 1
fi
echo -e "${GREEN}✓ go.mod found${NC}"
cat go.mod
echo ""

# Step 4: Download dependencies
echo -e "${YELLOW}[4/7] Downloading dependencies...${NC}"
go mod download
echo -e "${GREEN}✓ Dependencies downloaded${NC}\n"

# Step 5: Verify dependencies
echo -e "${YELLOW}[5/7] Verifying dependencies...${NC}"
go mod verify
echo -e "${GREEN}✓ Dependencies verified${NC}\n"

# Step 6: Run tests (if any exist)
echo -e "${YELLOW}[6/7] Running tests...${NC}"
if ls *_test.go 1> /dev/null 2>&1; then
    go test -v ./...
    echo -e "${GREEN}✓ Tests passed${NC}\n"
else
    echo -e "${YELLOW}! No test files found, skipping tests${NC}\n"
fi

# Step 7: Build the binary
echo -e "${YELLOW}[7/7] Building swiper binary...${NC}"
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S_UTC')
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
GO_VERSION_SHORT=$(go version | awk '{print $3}')

# Build with optimizations
go build -v \
    -ldflags="-s -w -X 'main.Version=dev' -X 'main.BuildTime=$BUILD_TIME' -X 'main.GitCommit=$GIT_COMMIT'" \
    -o swiper \
    main.go

if [ -f "swiper" ]; then
    BINARY_SIZE=$(du -h swiper | cut -f1)
    echo -e "${GREEN}✓ Build successful${NC}"
    echo -e "${GREEN}  Binary: ./swiper${NC}"
    echo -e "${GREEN}  Size: $BINARY_SIZE${NC}"
    echo -e "${GREEN}  Go Version: $GO_VERSION_SHORT${NC}"
    echo -e "${GREEN}  Git Commit: $GIT_COMMIT${NC}"
    echo -e "${GREEN}  Build Time: $BUILD_TIME${NC}\n"
else
    echo -e "${RED}Error: Build failed${NC}"
    exit 1
fi

# Verify the binary works
echo -e "${YELLOW}Verifying binary...${NC}"
if ./swiper -help &> /dev/null; then
    echo -e "${GREEN}✓ Binary executes successfully${NC}\n"
else
    echo -e "${RED}Error: Binary verification failed${NC}"
    exit 1
fi

# Summary
echo -e "${BLUE}========================================${NC}"
echo -e "${GREEN}Build Complete!${NC}"
echo -e "${BLUE}========================================${NC}"
echo -e "\n${YELLOW}Usage examples:${NC}"
echo -e "  ./swiper -file document.pdf"
echo -e "  ./swiper -dir /path/to/pdfs"
echo -e "  ./swiper -scan . -copydir pdf-collection"
echo -e "  ./swiper -help\n"