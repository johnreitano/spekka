# Change: Implement Install Command & Directory Scaffolding

## Why
The install command currently exists as a placeholder that only prints a message. Users need a functional install command that bootstraps repositories with the complete Spekka directory structure and configuration files to enable spec-driven development workflows. Without this implementation, users cannot properly initialize Spekka in their repositories.

## What Changes
- Implement full install command functionality in `internal/cli/install.go`
- Create directory scaffolding for: standards/, product/, specs/, tasks/, skills/, profiles/, verification/, and visuals/
- Generate initial .spekka.yaml configuration file with sensible defaults
- Add template files and README documentation in appropriate directories
- Add validation to prevent overwriting existing Spekka installations
- Include helpful success messages and next steps guidance

## Impact
- Affected specs: cli-framework (modify existing install command requirements)
- Affected code: internal/cli/install.go (primary implementation)
- New functionality: Directory creation, file generation, configuration setup
- User experience: Transforms placeholder into fully functional bootstrapping tool

