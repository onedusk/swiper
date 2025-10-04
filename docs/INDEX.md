# Swiper Documentation Index

Complete guide to Swiper PDF extraction tool and its architectural refactor.

## Quick Links

- **Get Started:** [QUICKSTART.md](QUICKSTART.md) - Installation and usage
- **Build Tool:** [verify-and-build.sh](verify-and-build.sh) - Automated build script
- **🔥 Architecture Proposal:** [PROPOSAL.md](PROPOSAL.md) - Modular refactor proposal (PENDING APPROVAL)
- **Development:** [CLAUDE.md](CLAUDE.md) - Development guidelines

## Document Overview

### User Documentation

**QUICKSTART.md** (6KB)
- Installation instructions (macOS, Ubuntu, CentOS)
- Usage examples for all three modes
- Performance optimization tips
- Troubleshooting guide
- Output format specifications

**Audience:** End users, first-time users

### Build & Development

**verify-and-build.sh** (Executable)
- Automated verification and build
- Dependency checking
- Go module verification
- Binary compilation with metadata
- Validation testing

**Usage:** `./verify-and-build.sh`

**CLAUDE.md** (4KB)
- Project development guidelines
- Key commands and workflows
- Architecture notes
- Prohibited practices
- AI-assisted development best practices

**Audience:** Developers, AI assistants

### Architecture Proposal

**PROPOSAL.md** (9KB) 🔥
- Current state analysis
- Proposed modular structure (8 packages)
- Benefits and trade-offs
- Example public API design
- Migration strategy
- Decision questions
- **Status: PENDING APPROVAL**

**Audience:** Decision makers, stakeholders, reviewers

**Key Point:** This is a proposal only - no implementation will begin without explicit approval.

## Current Status

**Binary:** ✅ Built successfully (2.8MB)
**Go Version:** 1.24.0
**Architecture:** Monolithic (1,776 lines in main.go)
**Refactor Status:** 🟡 PROPOSAL SUBMITTED - AWAITING APPROVAL

## Next Step

**Review PROPOSAL.md** and decide:
- ✅ Approve proposal → Implementation begins
- 🔄 Modify proposal → Adjust and resubmit
- ❌ Keep current → No changes

## Architecture at a Glance

### Current (Monolithic)
```
main.go (1,776 lines)
└── Everything in one file
```

### Proposed (Modular)
```
swiper/
├── cmd/swiper/main.go      (50-100 lines)
├── internal/               (8 packages)
│   ├── metrics/
│   ├── pool/
│   ├── cache/
│   ├── config/
│   ├── extractor/
│   ├── scanner/
│   └── batch/
└── pkg/swiper/             (Public API)
```

## Key Features

### Current Capabilities
- Single PDF extraction to markdown
- Batch processing of multiple PDFs
- Recursive PDF scanning and copying
- Parallel page processing
- Advanced buffer pooling
- Performance metrics
- Result caching
- Async logging

### Planned Improvements
- Modular architecture
- Unit testable components
- Public library API
- 10x better maintainability
- Professional structure
- Extensible design

## Performance Characteristics

**Processing Speed:** ~8 pages/second
**Concurrency:** Configurable worker count
**Buffer Management:** 4-tier pooling (32KB, 128KB, 256KB, 1MB)
**Cache Hit Rate:** Configurable caching
**Resource Usage:** Optimized temp directory pooling

## Usage Examples

### Extract Single PDF
```bash
./swiper -file document.pdf
```

### Batch Process
```bash
./swiper -dir /path/to/pdfs -output extracted
```

### Scan and Copy
```bash
./swiper -scan . -copydir pdf-collection
```

### Performance Mode
```bash
./swiper -file doc.pdf -processes 8 -cache -benchmark
```

## Development Commands

### Build
```bash
./verify-and-build.sh
```

### Manual Build
```bash
go build -o swiper main.go
```

### Run
```bash
./swiper -file test.pdf
```

### Test (Future)
```bash
go test ./...
go test -cover ./...
```

## Project Structure (Future)

### Internal Packages
- **metrics/** - Performance tracking (200 lines)
- **pool/** - Resource pooling (180 lines)
- **cache/** - Result caching (80 lines)
- **config/** - Configuration (120 lines)
- **extractor/** - PDF extraction (400+ lines)
- **scanner/** - File discovery (250 lines)
- **batch/** - Batch processing (230 lines)

### Public API
- **pkg/swiper/** - Library interface (150 lines)

### CLI
- **cmd/swiper/** - Command-line tool (50-100 lines)

## Documentation Philosophy

### Layered Approach
1. **Quick Start** - Get running fast
2. **Architecture** - Understand the design
3. **Implementation** - Execute the plan
4. **Summary** - See the big picture
5. **Index** - Navigate everything

### Audience-Focused
- **Users:** QUICKSTART.md
- **Developers:** CLAUDE.md, REFACTOR_PLAN.md
- **Architects:** ARCHITECTURE_PROPOSAL.md
- **Stakeholders:** IMPLEMENTATION_SUMMARY.md
- **Everyone:** INDEX.md (this file)

## Success Criteria

### Must Have
- [ ] Binary compiles
- [ ] All workflows functional
- [ ] Performance maintained
- [ ] Documentation updated

### Should Have
- [ ] >70% test coverage
- [ ] Public API documented
- [ ] Examples work
- [ ] Benchmarks recorded

### Nice to Have
- [ ] >80% test coverage
- [ ] Performance improvements
- [ ] Additional examples
- [ ] CI/CD pipeline

## Contributing

This is an internal tool, but architectural feedback is welcome.

### Areas for Input
- Package structure
- Dependency organization
- API design
- Testing strategy
- Performance optimization

## Related Documentation

### In This Repository
- `docs/doc.md` - Go Modules Reference (62KB)
- `docs/implementation_proof_protocol/` - IPP framework
- `docs/swiper.performance.txt` - Performance notes

### External Resources
- [Go Modules Reference](https://go.dev/ref/mod)
- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://go.dev/doc/effective_go)

## Version History

**v0.9-monolithic** (Current)
- Monolithic architecture
- Full feature set
- High performance
- Production ready

**v1.0.0** (Planned)
- Modular architecture
- Public API
- Comprehensive tests
- Enhanced maintainability

## Contact & Support

For questions about:
- **Usage:** See QUICKSTART.md
- **Development:** See CLAUDE.md
- **Architecture:** See ARCHITECTURE_PROPOSAL.md
- **Implementation:** See REFACTOR_PLAN.md

## License

Check repository root for license information.

---

**Last Updated:** 2025-09-29
**Current Status:** Proposal Submitted, Awaiting Approval
**Next Action:** Review PROPOSAL.md and approve/modify/reject