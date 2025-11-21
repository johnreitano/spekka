# Design: Core CLI Framework & Project Structure

## Context
Spekka is a CLI-based development workflow tool built in Go. This is the foundational implementation that establishes the project architecture, dependency management, and command structure. The design must support future extensibility while maintaining simplicity and following Go best practices as documented in `openspec/project.md`.

**Key Constraints:**
- Must use Go 1.25+ for modern language features
- CLI-first design with no GUI dependencies
- Cross-platform support (macOS, Linux, Windows)
- Minimal dependencies, prefer standard library
- Commands should respond in <1s for typical operations

**Stakeholders:**
- Developers implementing future Spekka features
- End users interacting with the CLI
- Contributors extending Spekka functionality

## Goals / Non-Goals

**Goals:**
- Establish standard Go project structure (cmd/internal/pkg)
- Implement Cobra-based CLI with five primary commands
- Create configuration system for .spekka.yaml files
- Set up proper error handling and logging
- Enable future command implementations to plug in easily
- Provide excellent CLI UX with help text and validation

**Non-Goals:**
- Implementing full functionality of each command (that comes in subsequent changes)
- Adding web UI or webhook server (separate changes)
- Implementing style guide enforcement or verification systems
- Creating actual workflow logic (spec creation, task breakdown, etc.)

## Decisions

### Decision 1: Use Cobra for CLI Framework
**Rationale:** Cobra is the de facto standard for Go CLI applications, used by kubectl, Hugo, and GitHub CLI. It provides:
- Automatic help generation
- Flag parsing and validation
- Subcommand structure
- Shell completion support
- Well-documented and actively maintained

**Alternatives Considered:**
- `urfave/cli`: Simpler but less powerful, lacks some features we'll need for complex command hierarchies
- `kong`: Type-safe but more opinionated, steeper learning curve

### Decision 2: Use Viper for Configuration
**Rationale:** Viper complements Cobra and provides:
- Multiple config format support (YAML, JSON, TOML)
- Environment variable integration
- Config file precedence (project, user, system)
- Flag binding to config values
- Made by same author as Cobra, seamless integration

**Alternatives Considered:**
- Standard library (encoding/yaml): Would require more custom code for precedence and env var handling
- Custom solution: Unnecessary reinvention

### Decision 3: Pragmatic Layers Architecture
**Rationale:** Following project.md conventions:
```
spekka/
├── cmd/spekka/          # Main entry point
├── internal/            # Private application code
│   ├── cli/            # Command implementations
│   └── config/         # Configuration management
└── go.mod              # Dependencies
```

This structure:
- Keeps main.go minimal (just CLI initialization)
- Uses internal/ to prevent external package imports
- Reserves pkg/ for future public libraries (if needed)
- Clear separation of concerns

**Alternatives Considered:**
- Flat structure: Doesn't scale well for larger projects
- Clean Architecture layers: Overcomplicated for CLI tool at this stage

### Decision 4: Command Structure
**Commands to scaffold:**
- `spekka` - Root command with global flags
- `spekka install` - Repository bootstrapping
- `spekka setup-project` - Project configuration
- `spekka create-spec` - Specification creation
- `spekka implement-tasks` - Task implementation
- `spekka serve` - Webhook server

Each command will have:
- Help text and usage examples
- Flag definitions (even if not yet implemented)
- Placeholder implementation returning helpful messages
- RunE function for error handling

### Decision 5: Configuration File Structure
**.spekka.yaml format:**
```yaml
project:
  name: string
  version: string
standards:
  path: string        # Path to standards directory
  remote: []string    # Git URLs for remote style guides
product:
  path: string        # Path to product vision files
specs:
  path: string        # Path to specifications
tasks:
  path: string        # Path to task files
```

Start minimal, extend in future changes.

### Decision 6: Error Handling Pattern
Use Go 1.13+ error wrapping:
```go
if err != nil {
    return fmt.Errorf("failed to load config: %w", err)
}
```

Benefits:
- Maintains error context through call stack
- Enables error inspection with errors.Is/As
- Clear error messages for users

### Decision 7: Logging with log/slog
Use standard library log/slog (Go 1.25+):
- Structured logging built-in
- Multiple log levels (debug, info, warn, error)
- Configurable output format
- No external dependencies

### Decision 8: Flag Style - Both Short and Long Forms
**Rationale:** Support both short and long flags for common operations to balance usability and clarity.
- Common flags get both forms: `-v`/`--verbose`, `-h`/`--help`, `-q`/`--quiet`, `-d`/`--debug`
- Less common flags: long form only to avoid namespace pollution
- Follows Unix/GNU conventions familiar to CLI users

**Benefits:**
- Better UX for frequent users (less typing)
- Clear self-documenting commands with long flags
- Familiar patterns from other CLI tools

### Decision 9: Configuration File Requirements
**Rationale:** Make configuration optional for bootstrap commands but required for context-dependent operations.

**Policy:**
- **Optional:** `install` command (it creates the config file)
- **Optional:** `--help`, `--version` flags (informational only)
- **Required:** Commands needing project context:
  - `setup-project` - configures project settings
  - `create-spec` - needs standards and product paths
  - `implement-tasks` - needs specs and tasks paths
  - `serve` - needs webhook configuration

**Error Handling:**
- When config is missing but required, display clear error message
- Include suggestion to run `spekka install` first
- Provide example .spekka.yaml structure in error output

### Decision 10: Runtime Go Version Validation
**Rationale:** Validate Go version at runtime to provide clear, immediate feedback if requirements aren't met.

**Implementation:**
- Check runtime Go version in main.go before CLI initialization
- Fail fast with clear error if version < 1.25
- Error message includes:
  - Current Go version detected
  - Minimum required version (1.25)
  - Instructions for upgrading Go

**Trade-offs:**
- Small startup overhead (~1ms) - acceptable for CLI tool
- Better UX than cryptic build errors
- Catches version mismatches in pre-built binaries

## Risks / Trade-offs

**Risk: Over-engineering the initial framework**
- **Mitigation:** Keep commands as minimal scaffolds, implement behavior in subsequent changes
- **Trade-off:** May need minor refactoring as we add actual functionality

**Risk: Configuration file format may need changes**
- **Mitigation:** Start with minimal required fields, use Viper's flexibility for evolution
- **Trade-off:** Early adopters may need config updates (acceptable for alpha/beta)

**Risk: Cobra's API may change**
- **Mitigation:** Cobra is mature and stable, breaking changes unlikely
- **Trade-off:** Dependency on external package (acceptable given benefits)

## Migration Plan

**Initial Setup:**
1. Initialize Go module: `go mod init github.com/jreitano/spekka`
2. Add dependencies: `go get github.com/spf13/cobra@latest github.com/spf13/viper@latest`
3. Create directory structure
4. Implement main.go and root command
5. Scaffold subcommands with placeholders

**Testing:**
- Unit tests for configuration parsing
- Integration tests for command execution
- Manual testing on macOS, Linux, Windows

**Rollback:**
- N/A (greenfield implementation, no existing system to rollback)
