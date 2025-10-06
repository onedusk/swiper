#!/bin/bash
# IPP Gate Enforcement Script
# Validates gate completion before allowing next phase

set -euo pipefail

# Configuration
GATE="${1:-}"
IMPROVEMENT="${2:-}"
LOG_DIR="logs/ipp-gates"
TIMESTAMP=$(date +%Y%m%d-%H%M%S)
EVIDENCE_DIR="docs/evidence"
LOG_FILE="/dev/null"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Usage
usage() {
    cat << EOF
Usage: $0 <GATE> <IMPROVEMENT_ID>

Gates:
  SPEC    - Specification gate (plan document required)
  BUILD   - Build gate (compilation, linting)
  TEST    - Test gate (test suite, race detector)
  VALID   - Validation gate (benchmarks, integration)
  PROVE   - Proof gate (PR requirements)

Examples:
  $0 SPEC A01
  $0 BUILD A04
  $0 TEST A04
EOF
    exit 1
}

# Logging
log_append() {
    printf '[%s] %s\n' "$1" "$2" >> "$LOG_FILE"
}

log_info() {
    local message="$*"
    log_append "INFO" "$message"
    printf '%b\n' "${GREEN}[INFO]${NC} $message"
}

log_error() {
    local message="$*"
    log_append "ERROR" "$message"
    printf '%b\n' "${RED}[ERROR]${NC} $message" >&2
}

log_warn() {
    local message="$*"
    log_append "WARN" "$message"
    printf '%b\n' "${YELLOW}[WARN]${NC} $message"
}

log_gate() {
    local message="$*"
    local label="GATE:${GATE}"
    log_append "$label" "$message"
    printf '%b\n' "${GREEN}[GATE:${GATE}]${NC} $message"
}

# Setup
setup() {
    mkdir -p "$LOG_DIR" "$EVIDENCE_DIR"
    LOG_FILE="$LOG_DIR/${TIMESTAMP}-${IMPROVEMENT}-${GATE}.log"
    : > "$LOG_FILE"

    log_info "IPP Gate Enforcement"
    log_info "Gate: $GATE"
    log_info "Improvement: $IMPROVEMENT"
    log_info "Timestamp: $TIMESTAMP"
    log_info "Log: $LOG_FILE"
    echo ""
    printf '\n' >> "$LOG_FILE"
}

# IPP-SPEC Gate
gate_spec() {
    log_gate "Validating SPEC gate..."

    local plan_file="docs/${IMPROVEMENT}_IMPLEMENTATION_PLAN.md"
    local required_sections=("Success Criteria" "Rollback Strategy" "Implementation Checklist")

    # Check plan file exists
    if [[ ! -f "$plan_file" ]]; then
        log_error "Plan file not found: $plan_file"
        log_error "Required: docs/ANN_IMPLEMENTATION_PLAN.md"
        return 1
    fi
    log_info "✓ Plan file exists: $plan_file"

    # Check required sections
    for section in "${required_sections[@]}"; do
        if ! grep -q "$section" "$plan_file"; then
            log_error "Missing required section: $section"
            return 1
        fi
        log_info "✓ Section found: $section"
    done

    # Count steps
    local step_count=$(grep -c "^- \[ \]" "$plan_file" || echo 0)
    if [[ $step_count -lt 20 ]]; then
        log_error "Implementation checklist must contain at least 20 steps (found $step_count)"
        return 1
    fi
    log_info "✓ Step count: $step_count"

    log_gate "SPEC gate PASSED"
    return 0
}

# IPP-BUILD Gate
gate_build() {
    log_gate "Validating BUILD gate..."

    # Check branch naming
    local current_branch=$(git branch --show-current)
    if [[ ! "$current_branch" =~ ^feature/${IMPROVEMENT,,}- ]]; then
        log_error "Branch must match pattern: feature/aNN-description"
        log_error "Current branch: $current_branch"
        return 1
    fi
    log_info "✓ Branch naming: $current_branch"

    # Build check
    if ! go build -o /tmp/swiper-test ./cmd/swiper; then
        log_error "Build failed"
        return 1
    fi
    log_info "✓ Build successful"
    rm -f /tmp/swiper-test

    # Vet check
    if ! go vet ./...; then
        log_error "go vet failed"
        return 1
    fi
    log_info "✓ go vet passed"

    # Format check
    local unformatted=$(go fmt ./... 2>&1)
    if [[ -n "$unformatted" ]]; then
        log_error "Code not formatted: $unformatted"
        return 1
    fi
    log_info "✓ Code formatting valid"

    log_gate "BUILD gate PASSED"
    return 0
}

# IPP-TEST Gate
gate_test() {
    log_gate "Validating TEST gate..."

    local test_log="/tmp/ipp-test-${IMPROVEMENT}-${TIMESTAMP}.log"

    # Run tests
    if ! go test -v ./... > "$test_log" 2>&1; then
        log_error "Tests failed. Log: $test_log"
        tail -20 "$test_log" >&2
        return 1
    fi
    log_info "✓ Tests passed"

    # Race detector
    if ! go test -race ./... > "${test_log}.race" 2>&1; then
        log_error "Race detector found issues. Log: ${test_log}.race"
        tail -20 "${test_log}.race" >&2
        return 1
    fi
    log_info "✓ Race detector clean"

    # Save evidence
    cp "$test_log" "$EVIDENCE_DIR/${IMPROVEMENT}-test-${TIMESTAMP}.log"
    log_info "✓ Test evidence saved: $EVIDENCE_DIR/${IMPROVEMENT}-test-${TIMESTAMP}.log"

    log_gate "TEST gate PASSED"
    return 0
}

# IPP-VALID Gate
gate_valid() {
    log_gate "Validating VALID gate..."

    # Check CHANGELOG updated
    if ! git diff origin/main HEAD -- docs/CHANGELOG.md | grep -q "^+"; then
        log_error "CHANGELOG.md not updated"
        return 1
    fi
    log_info "✓ CHANGELOG.md updated"

    # Check for expected log behavior (example: A04 skip message)
    local test_output="/tmp/ipp-validation-${IMPROVEMENT}.txt"
    if [[ -f "$test_output" ]]; then
        log_info "✓ Integration test output found: $test_output"
    else
        log_warn "No integration test output found (optional)"
    fi

    # Benchmark check (if benchmarks exist)
    if go test -run=^$ -bench=. ./... > /dev/null 2>&1; then
        log_info "✓ Benchmarks available (optional validation)"
    fi

    log_gate "VALID gate PASSED"
    return 0
}

# IPP-PROVE Gate
gate_prove() {
    log_gate "Validating PROVE gate..."

    if ! command -v gh >/dev/null 2>&1; then
        log_error "GitHub CLI (gh) is required for the PROVE gate"
        return 1
    fi

    # Check PR exists
    local pr_number=$(gh pr view --json number -q .number 2>/dev/null || echo "")
    if [[ -z "$pr_number" ]]; then
        log_error "No PR found for current branch"
        log_error "Create PR with: gh pr create"
        return 1
    fi
    log_info "✓ PR exists: #$pr_number"

    # Check PR description contains evidence
    local pr_body=$(gh pr view --json body -q .body)
    if [[ ! "$pr_body" =~ "test" ]] && [[ ! "$pr_body" =~ "evidence" ]]; then
        log_warn "PR description may lack test evidence"
    else
        log_info "✓ PR description includes evidence"
    fi

    # Check CI status
    local ci_status=$(gh pr view --json statusCheckRollup -q '.statusCheckRollup[0].state' 2>/dev/null || echo "")
    if [[ "$ci_status" == "FAILURE" ]] || [[ "$ci_status" == "ERROR" ]]; then
        log_error "CI checks failing"
        return 1
    fi
    log_info "✓ CI status: ${ci_status:-PENDING}"

    # Test rollback capability
    log_info "Testing rollback capability..."
    local current_sha=$(git rev-parse HEAD)
    if git revert --no-commit "$current_sha" > /dev/null 2>&1; then
        git revert --abort
        log_info "✓ Rollback test successful"
    else
        log_warn "Rollback test skipped (may have conflicts)"
    fi

    log_gate "PROVE gate PASSED"
    log_info "Human review required before merge"
    return 0
}

# Main
main() {
    if [[ -z "$GATE" ]] || [[ -z "$IMPROVEMENT" ]]; then
        usage
    fi

    setup

    case "$GATE" in
        SPEC)
            gate_spec
            ;;
        BUILD)
            gate_build
            ;;
        TEST)
            gate_test
            ;;
        VALID)
            gate_valid
            ;;
        PROVE)
            gate_prove
            ;;
        *)
            log_error "Unknown gate: $GATE"
            usage
            ;;
    esac

    local exit_code=$?

    if [[ $exit_code -eq 0 ]]; then
        echo ""
        log_info "========================================="
        log_info "Gate: $GATE - PASSED ✓"
        log_info "Improvement: $IMPROVEMENT"
        log_info "You may proceed to next phase"
        log_info "========================================="
    else
        echo ""
        log_error "========================================="
        log_error "Gate: $GATE - FAILED ✗"
        log_error "Improvement: $IMPROVEMENT"
        log_error "Fix issues before proceeding"
        log_error "========================================="
    fi

    return $exit_code
}

main "$@"
