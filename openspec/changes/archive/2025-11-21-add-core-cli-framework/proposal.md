# Change: Add Core CLI Framework & Project Structure

## Why
Spekka needs a foundational CLI framework to serve as the entry point for all user interactions. This establishes the base architecture and command structure that all future features will build upon. Without this foundation, we cannot implement any of the planned workflow commands (install, setup-project, create-spec, implement-tasks, serve) or provide a cohesive user experience.

## What Changes
- Create standard Go project layout with cmd/, internal/, and pkg/ directories following Go best practices
- Implement base CLI application using Cobra framework with command structure and flag parsing
- Add five primary commands: install, setup-project, create-spec, implement-tasks, and serve (initial scaffolding only)
- Create configuration system for loading and parsing .spekka.yaml configuration files
- Set up Go module with required dependencies (Cobra, Viper, etc.)
- Implement basic error handling, logging (using log/slog), and help documentation
- Create main entry point and build configuration

This is the **foundation change** - it establishes the technical architecture but does not implement the full functionality of each command. Subsequent changes will build upon this structure to add actual command behaviors.

## Impact
- Affected specs: `cli-framework` (new capability)
- Affected code: Creates new Go codebase structure
  - `cmd/spekka/main.go` - Entry point
  - `internal/cli/` - Command implementations
  - `internal/config/` - Configuration management
  - `go.mod` - Dependency management
- Dependencies: None (this is the first implementation)
- Breaking changes: None (greenfield implementation)
