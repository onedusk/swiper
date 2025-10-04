#!/usr/bin/env bash

# Parallel PPM to JPG/PNG converter using ImageMagick
# Supports recursive directory traversal and parallel processing

set -euo pipefail

# Default values
OUTPUT_FORMAT="jpg"
QUALITY=90
PARALLEL_JOBS=$(nproc 2>/dev/null || sysctl -n hw.ncpu 2>/dev/null || echo 4)
PRESERVE_STRUCTURE=true
VERBOSE=false
DRY_RUN=false
SOURCE_DIR="."
OUTPUT_DIR=""

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to display usage
usage() {
    cat << EOF
Usage: $(basename "$0") [OPTIONS]

Recursively convert PPM images to JPG or PNG format using ImageMagick in parallel.

OPTIONS:
    -s, --source DIR        Source directory containing PPM files (default: current directory)
    -o, --output DIR        Output directory (default: same as source for each file)
    -f, --format FORMAT     Output format: jpg or png (default: jpg)
    -q, --quality NUM       JPEG quality 1-100 (default: 90, ignored for PNG)
    -j, --jobs NUM          Number of parallel jobs (default: number of CPU cores)
    -F, --flat              Don't preserve directory structure (output all to single directory)
    -v, --verbose           Enable verbose output
    -n, --dry-run           Show what would be converted without actually converting
    -h, --help              Show this help message

EXAMPLES:
    # Convert all PPM files to JPG with default settings
    $(basename "$0")

    # Convert to PNG format with 8 parallel jobs
    $(basename "$0") -f png -j 8

    # Convert with custom quality and output directory
    $(basename "$0") -o ./converted -q 95 -f jpg

    # Dry run to see what would be converted
    $(basename "$0") -n -v

EOF
    exit 0
}

# Function to print colored messages
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1" >&2
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_verbose() {
    if [ "$VERBOSE" = true ]; then
        echo -e "[VERBOSE] $1"
    fi
}

# Parse command line arguments
parse_args() {
    while [[ $# -gt 0 ]]; do
        case $1 in
            -s|--source)
                SOURCE_DIR="$2"
                shift 2
                ;;
            -o|--output)
                OUTPUT_DIR="$2"
                shift 2
                ;;
            -f|--format)
                OUTPUT_FORMAT="$2"
                case "$OUTPUT_FORMAT" in
                    jpg|jpeg|png)
                        ;;
                    *)
                        log_error "Invalid format: $OUTPUT_FORMAT. Use 'jpg' or 'png'"
                        exit 1
                        ;;
                esac
                # Normalize jpeg to jpg
                if [ "$OUTPUT_FORMAT" = "jpeg" ]; then
                    OUTPUT_FORMAT="jpg"
                fi
                shift 2
                ;;
            -q|--quality)
                QUALITY="$2"
                if ! [[ "$QUALITY" =~ ^[1-9][0-9]?$|^100$ ]]; then
                    log_error "Quality must be between 1 and 100"
                    exit 1
                fi
                shift 2
                ;;
            -j|--jobs)
                PARALLEL_JOBS="$2"
                if ! [[ "$PARALLEL_JOBS" =~ ^[1-9][0-9]*$ ]]; then
                    log_error "Number of jobs must be a positive integer"
                    exit 1
                fi
                shift 2
                ;;
            -F|--flat)
                PRESERVE_STRUCTURE=false
                shift
                ;;
            -v|--verbose)
                VERBOSE=true
                shift
                ;;
            -n|--dry-run)
                DRY_RUN=true
                shift
                ;;
            -h|--help)
                usage
                ;;
            *)
                log_error "Unknown option: $1"
                usage
                ;;
        esac
    done
}

# Check dependencies
check_dependencies() {
    local deps_missing=false
    
    if ! command -v convert &> /dev/null; then
        log_error "ImageMagick 'convert' command not found. Please install ImageMagick."
        deps_missing=true
    fi
    
    if ! command -v find &> /dev/null; then
        log_error "'find' command not found."
        deps_missing=true
    fi
    
    # Check for parallel processing tool
    if command -v parallel &> /dev/null; then
        USE_GNU_PARALLEL=true
        log_verbose "Using GNU parallel for processing"
    elif command -v xargs &> /dev/null; then
        USE_GNU_PARALLEL=false
        log_verbose "Using xargs for parallel processing"
    else
        log_error "Neither GNU parallel nor xargs found. Please install GNU parallel for best performance."
        deps_missing=true
    fi
    
    if [ "$deps_missing" = true ]; then
        exit 1
    fi
}

# Validate directories
validate_directories() {
    if [ ! -d "$SOURCE_DIR" ]; then
        log_error "Source directory does not exist: $SOURCE_DIR"
        exit 1
    fi
    
    # Convert to absolute path
    SOURCE_DIR=$(cd "$SOURCE_DIR" && pwd -P)
    
    # If output directory is specified
    if [ -n "$OUTPUT_DIR" ]; then
        if [ "$DRY_RUN" = false ]; then
            if [ ! -d "$OUTPUT_DIR" ]; then
                log_info "Creating output directory: $OUTPUT_DIR"
                mkdir -p "$OUTPUT_DIR"
            fi
        fi
        OUTPUT_DIR=$(cd "$OUTPUT_DIR" && pwd -P)
    fi
}

# Function to convert a single PPM file
convert_ppm() {
    local input_file="$1"
    local source_base="$2"
    local output_base="$3"
    local format="$4"
    local quality="$5"
    local preserve="$6"
    local dry_run="$7"
    local verbose="$8"
    
    # Get the relative path from source base
    local rel_path="${input_file#$source_base/}"
    local dir_path=$(dirname "$rel_path")
    local base_name=$(basename "$input_file" .ppm)
    
    # Determine output path
    local output_file
    if [ -n "$output_base" ]; then
        if [ "$preserve" = "true" ] && [ "$dir_path" != "." ]; then
            output_file="$output_base/$dir_path/${base_name}.${format}"
        else
            output_file="$output_base/${base_name}.${format}"
        fi
    else
        # Output to same directory as input
        output_file="${input_file%.ppm}.${format}"
    fi
    
    # Create output directory if needed
    local output_dir_path=$(dirname "$output_file")
    if [ "$dry_run" = "false" ] && [ ! -d "$output_dir_path" ]; then
        mkdir -p "$output_dir_path"
    fi
    
    # Execute conversion based on format
    
    if [ "$verbose" = "true" ]; then
        echo "[CONVERT] $input_file -> $output_file"
    fi
    
    if [ "$dry_run" = "false" ]; then
        local exit_code
        if [ "$format" = "jpg" ]; then
            convert "$input_file" -quality "$quality" "$output_file" 2>/dev/null
            exit_code=$?
        elif [ "$format" = "png" ]; then
            # PNG compression level (0-9, where 9 is maximum compression)
            # Map quality percentage to PNG compression inversely
            local png_compression=$((9 - (quality - 1) / 11))
            convert "$input_file" -quality "${png_compression}0" "$output_file" 2>/dev/null
            exit_code=$?
        fi
        if [ $exit_code -eq 0 ]; then
            echo "✓ $rel_path"
            return 0
        else
            echo "✗ $rel_path" >&2
            return 1
        fi
    else
        echo "[DRY-RUN] Would convert: $input_file -> $output_file"
    fi
}

# Export functions for parallel execution (bash-specific)
if [ -n "$BASH_VERSION" ]; then
    export -f convert_ppm
    export -f log_verbose
fi

# Main conversion process
main() {
    parse_args "$@"
    check_dependencies
    validate_directories
    
    # Count PPM files
    local ppm_count=$(find "$SOURCE_DIR" -type f -iname "*.ppm" 2>/dev/null | wc -l)
    
    if [ "$ppm_count" -eq 0 ]; then
        log_warning "No PPM files found in $SOURCE_DIR"
        exit 0
    fi
    
    log_info "Found $ppm_count PPM file(s) to convert"
    log_info "Output format: $(echo "$OUTPUT_FORMAT" | tr '[:lower:]' '[:upper:]')"
    if [ "$OUTPUT_FORMAT" = "jpg" ]; then
        log_info "JPEG quality: $QUALITY"
    fi
    log_info "Parallel jobs: $PARALLEL_JOBS"
    
    if [ -n "$OUTPUT_DIR" ]; then
        log_info "Output directory: $OUTPUT_DIR"
        if [ "$PRESERVE_STRUCTURE" = false ]; then
            log_info "Flat output structure (no subdirectories)"
        else
            log_info "Preserving directory structure"
        fi
    else
        log_info "Output: Same directory as each source file"
    fi
    
    if [ "$DRY_RUN" = true ]; then
        log_warning "DRY RUN MODE - No files will be converted"
    fi
    
    echo ""
    
    # Prepare for tracking results
    local start_time=$(date +%s)
    local success_count=0
    local fail_count=0
    
    # Process files in parallel
    if [ "$USE_GNU_PARALLEL" = true ]; then
        # Using GNU parallel
        find "$SOURCE_DIR" -type f -iname "*.ppm" -print0 | \
        parallel -0 -j "$PARALLEL_JOBS" --bar \
            convert_ppm {} "$SOURCE_DIR" "$OUTPUT_DIR" "$OUTPUT_FORMAT" \
            "$QUALITY" "$PRESERVE_STRUCTURE" "$DRY_RUN" "$VERBOSE"
    else
        # Using xargs
        find "$SOURCE_DIR" -type f -iname "*.ppm" -print0 | \
        xargs -0 -P "$PARALLEL_JOBS" -I {} bash -c \
            'convert_ppm "$@"' _ {} "$SOURCE_DIR" "$OUTPUT_DIR" "$OUTPUT_FORMAT" \
            "$QUALITY" "$PRESERVE_STRUCTURE" "$DRY_RUN" "$VERBOSE"
    fi
    
    local exit_code=$?
    local end_time=$(date +%s)
    local duration=$((end_time - start_time))
    
    echo ""
    log_info "Conversion completed in $duration seconds"
    
    if [ "$DRY_RUN" = false ]; then
        # Count actual converted files for summary
        if [ -n "$OUTPUT_DIR" ]; then
            local converted_count=$(find "$OUTPUT_DIR" -type f -name "*.${OUTPUT_FORMAT}" 2>/dev/null | wc -l)
            log_info "Successfully converted $converted_count file(s)"
        fi
    fi
    
    exit $exit_code
}

# Run main function
main "$@"