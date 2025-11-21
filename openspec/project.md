# Project Context

## Purpose
Spekka is a CLI-based development workflow tool that helps software engineering teams build better products through spec-driven development. It provides a structured, AI-assisted approach to planning, specification writing, and implementation that maximizes the value of LLM context windows. The tool integrates with popular issue tracking systems via webhooks and uses a 3-layer context system (standards, product vision, specifications) to maintain clarity and consistency throughout the development lifecycle.

**Core Goals:**
- Enable spec-driven development workflows with AI assistance
- Optimize LLM context window usage through structured information layers
- Keep specifications version-controlled and synchronized with code
- Enforce team coding standards automatically during AI-assisted development
- Integrate seamlessly with existing issue trackers and development tools

## Tech Stack
- **Language:** Go (latest stable version)
- **CLI Framework:** Cobra (for command structure and argument parsing)
- **Configuration:** Viper (configuration management)
- **Markdown Processing:** goldmark or blackfriday (spec/doc parsing)
- **HTTP Server:** net/http with gorilla/mux or chi (for webhook server and web UI)
- **Testing:** Go testing package with testify assertions, table-driven tests, gomock when needed
- **Linting:** golangci-lint (meta-linter for comprehensive code quality)
- **Build:** Standard Go toolchain (go build, go mod)

## Project Conventions

### Code Style
- **Formatting:** Use `gofmt` and `goimports` for all code (enforced in CI)
- **Linting:** golangci-lint with strict configuration covering major linters (govet, staticcheck, errcheck, gosec, etc.)
- **Naming:**
  - Use idiomatic Go naming (MixedCaps for exported, mixedCaps for unexported)
  - Package names: short, lowercase, no underscores (e.g., `spec`, `webhook`, `context`)
  - Interface names: end with `-er` suffix when appropriate (e.g., `SpecParser`, `TaskRunner`)
- **Comments:**
  - All exported functions, types, and packages require doc comments
  - Comments start with the name of the item being described
  - Use `//` style comments (not `/* */` except for package doc blocks)
- **Error Handling:**
  - Always check errors, never ignore with `_`
  - Wrap errors with context using `fmt.Errorf` with `%w` verb
  - Return errors to caller rather than logging when appropriate
- **File Organization:**
  - One primary type per file when possible
  - Group related functionality in same package
  - Keep files focused and under 500 lines when practical

### Architecture Patterns
**Pragmatic Layers Structure:**
```
spekka/
├── cmd/               # CLI entry points and command definitions
│   └── spekka/        # Main CLI application
├── internal/          # Private application code
│   ├── spec/          # Specification management
│   ├── task/          # Task breakdown and execution
│   ├── style/         # Style guide enforcement
│   ├── webhook/       # Webhook server and handlers
│   ├── context/       # 3-layer context system
│   └── config/        # Configuration management
├── pkg/               # Public libraries (if needed for extensibility)
└── docs/              # Documentation and examples
```

**Patterns:**
- **Dependency Injection:** Pass dependencies explicitly via constructors/functions
- **Interface Segregation:** Define small, focused interfaces at point of use
- **Command Pattern:** Use Cobra's command structure for CLI operations
- **Repository Pattern:** Abstract storage operations (filesystem, git) behind interfaces for testability
- **Factory Pattern:** Use factory functions for complex object initialization

### Testing Strategy
- **Unit Tests:**
  - Use table-driven tests for comprehensive scenario coverage
  - Test files named `*_test.go` alongside implementation
  - Use testify/assert for readable assertions
  - Use testify/mock or gomock for mocking dependencies when testify is insufficient
  - Aim for >80% code coverage on core logic
- **Integration Tests:**
  - Test full command execution in `cmd/` packages
  - Use temporary directories/files for filesystem operations
  - Mock external HTTP calls to issue trackers
- **Test Organization:**
  - Table-driven tests with descriptive test case names
  - Use `t.Run()` for subtests to improve output clarity
  - Setup/teardown in test helper functions
- **Benchmarks:**
  - Benchmark performance-critical paths (spec parsing, context building)
  - Use `testing.B` for benchmark tests

### Implementation Verification Requirements
All task implementations MUST pass these checks before being marked complete:

1. **Linting:** Run `golangci-lint run` and resolve all errors
   - Zero tolerance for linting errors
   - Address all warnings when practical
2. **Testing:** Run `go test ./...` with 100% pass rate
   - All existing tests must continue passing
   - New functionality must have test coverage
3. **Build:** Run `go build ./cmd/spekka` successfully
   - Must compile without errors
   - Verify no build warnings
4. **Coverage:** Verify test coverage ≥80% for modified packages
   - Check with `go test -cover ./...`
   - Focus on core logic paths
5. **Manual Verification:** Smoke test the implemented feature
   - Verify feature works as intended
   - Check for obvious regressions

**Important:** These checks apply regardless of which AI coding assistant is being used. Do not mark implementation tasks as complete until all verification checks pass. When using OpenSpec workflow, include these as final tasks in every `tasks.md` file.

### Git Workflow
- **Branching Strategy:** GitHub Flow
  - `main` branch is always deployable
  - Feature branches created from `main` for all changes
  - Branch naming: `feature/description`, `fix/issue-name`, `refactor/area`
  - Merge to `main` via Pull Requests with review
- **Commit Conventions:**
  - Use conventional commits format: `type(scope): description`
  - Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`
  - Keep commits focused and atomic
  - Write imperative mood in subject ("add feature" not "added feature")
  - Reference issues in commit body when applicable
- **Pull Requests:**
  - Require CI checks to pass (tests, linting)
  - Require at least one code review approval
  - Squash merge for clean history
  - Delete feature branches after merge

## Domain Context

### 3-Layer Context System
Spekka organizes project information into three distinct layers to optimize LLM context window usage:
1. **Standards Layer (How):** Coding standards, style guides, architectural patterns - defines how code should be written
2. **Product Layer (Why):** Product vision, user personas, business goals - defines why we're building features
3. **Specifications Layer (What):** Detailed feature specs, task breakdowns - defines what to build

This separation allows AI assistants to load only relevant context for each type of work.

### Key Concepts
- **Spec-Driven Development:** Specifications are written before implementation and stored as markdown in the repository
- **Git-Based Style Guides:** Style guides can be imported and shared using Git URLs (similar to Go package imports)
- **Webhook Integration:** External issue trackers (GitHub Issues, Linear, JIRA) sync bidirectionally via webhooks
- **Task Breakdown:** Specifications are automatically decomposed into implementable tasks with effort estimates
- **Context Injection:** The right context (standards, product vision, specs) is injected at the right time during AI-assisted development

### Workflow Commands
- `install`: Bootstrap repository with Spekka structure and config
- `setup-project`: Initialize project-level configuration and standards
- `create-spec`: Generate structured specification documents
- `implement-tasks`: Break down specs into tasks and execute with AI
- `serve`: Run webhook server for issue tracker integration

## Important Constraints

### Technical Constraints
- **Go Version:** Require Go 1.25+ for modern generics and standard library features
- **Filesystem-Based:** All specs, tasks, and context stored as files in repository (no external database)
- **CLI-First:** Must work seamlessly in terminal without GUI dependencies
- **Cross-Platform:** Support macOS, Linux, and Windows
- **Git Required:** Repository must be a Git repository for version control integration

### Design Constraints
- **LLM Context Optimization:** All features must consider impact on LLM context window size
- **Markdown-Based:** Specifications and documentation use markdown for readability and version control
- **No Lock-In:** Users should be able to use specs without Spekka (they're just markdown files)
- **Minimal Dependencies:** Prefer standard library when possible to reduce complexity

### Performance Constraints
- **CLI Responsiveness:** Commands should complete in <1s for typical operations
- **Large Repositories:** Must handle repositories with hundreds of specs efficiently
- **Webhook Latency:** Webhook responses should complete within reasonable HTTP timeout windows

## External Dependencies

### Required External Tools
- **Git:** For version control and Git-based style guide imports
- **Go Toolchain:** For building and running the CLI

### Optional External Integrations
- **GitHub Issues:** Webhook integration for issue synchronization
- **Linear:** Webhook integration for issue synchronization
- **JIRA:** Webhook integration for issue synchronization
- **Trello:** Webhook integration for board synchronization

### AI Coding Assistants
Spekka is designed to work with:
- Claude Code (Anthropic)
- Cursor
- GitHub Copilot
- Other LLM-based coding assistants

### Libraries and Frameworks
- **CLI:** github.com/spf13/cobra, github.com/spf13/viper
- **HTTP:** gorilla/mux or go-chi/chi for routing
- **Markdown:** yuin/goldmark or russross/blackfriday
- **Testing:** stretchr/testify, golang/mock (if needed)
- **Validation:** go-playground/validator
- **Logging:** log/slog (standard library)
