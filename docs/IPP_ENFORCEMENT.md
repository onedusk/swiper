# IPP Gate Enforcement

*Last Updated: 2025-10-06*

## Quick Start

```bash
# Validate each gate in sequence
./scripts/ipp-enforce.sh SPEC A01
./scripts/ipp-enforce.sh BUILD A01
./scripts/ipp-enforce.sh TEST A01
./scripts/ipp-enforce.sh VALID A01
./scripts/ipp-enforce.sh PROVE A01
```

## Gate Sequence

Each improvement (A01-A10) must pass through 5 gates:

1. **SPEC** - Plan document exists with success criteria
2. **BUILD** - Code compiles, passes vet and fmt
3. **TEST** - Test suite passes with race detector
4. **VALID** - Benchmarks validated, CHANGELOG updated
5. **PROVE** - PR created with evidence, CI passing

**Hard Rule:** Cannot proceed to next gate until previous passes.

## SPEC Gate

**Validates:**
- `docs/ANN_IMPLEMENTATION_PLAN.md` exists
- Contains "Success Criteria" section
- Contains "Rollback Strategy" section
- Contains "Implementation Checklist"
- Minimum 20 steps recommended

**Example:**
```bash
./scripts/ipp-enforce.sh SPEC A04
```

## BUILD Gate

**Validates:**
- Branch named `feature/aNN-description`
- `go build ./cmd/swiper` succeeds
- `go vet ./...` zero warnings
- `go fmt` all files formatted

**Example:**
```bash
git checkout -b feature/a04-absolute-paths
./scripts/ipp-enforce.sh BUILD A04
```

## TEST Gate

**Validates:**
- `go test ./...` passes
- `go test -race ./...` clean
- Test logs saved to `/tmp/ipp-test-*.log`
- Evidence copied to `docs/evidence/`

**Example:**
```bash
./scripts/ipp-enforce.sh TEST A04
```

## VALID Gate

**Validates:**
- `docs/CHANGELOG.md` updated (diff from main)
- Integration test artifacts present
- Benchmarks run successfully (optional)

**Example:**
```bash
# After updating CHANGELOG
./scripts/ipp-enforce.sh VALID A04
```

## PROVE Gate

**Validates:**
- PR exists for current branch
- PR description includes evidence
- CI checks passing or pending
- Rollback test successful

**Example:**
```bash
gh pr create --title "feat: A04 implementation"
./scripts/ipp-enforce.sh PROVE A04
```

## Output

**Pass:**
```
=========================================
Gate: SPEC - PASSED ✓
Improvement: A04
You may proceed to next phase
=========================================
```

**Fail:**
```
=========================================
Gate: BUILD - FAILED ✗
Improvement: A04
Fix issues before proceeding
=========================================
```

## Logs

All gate executions logged to: `logs/ipp-gates/TIMESTAMP-IMP-GATE.log`

Evidence artifacts: `docs/evidence/IMP-artifact-TIMESTAMP.log`

## Rollback on Failure

If any gate fails:

```bash
# Discard changes
git reset --hard origin/main
git clean -fd

# Or revert specific commit
git revert <sha>
```

## Automation Integration

Use in daily automation:

```bash
#!/bin/bash
# Daily A01-A10 implementation

for gate in SPEC BUILD TEST VALID PROVE; do
    if ! ./scripts/ipp-enforce.sh "$gate" "$IMPROVEMENT"; then
        echo "Gate $gate failed - aborting"
        git reset --hard origin/main
        exit 1
    fi
done
```

## Bypass (Emergency Only)

No bypass mechanism. If gate incorrectly fails, fix the validation logic in `scripts/ipp-enforce.sh`.

Human review required at PROVE gate before merge.
