# Implementation Tasks

## 1. Project Initialization
- [x] 1.1 Initialize Go module with `go mod init github.com/jreitano/spekka`
- [x] 1.2 Create directory structure: cmd/spekka/, internal/cli/, internal/config/
- [x] 1.3 Add Cobra dependency: `go get github.com/spf13/cobra@latest`
- [x] 1.4 Add Viper dependency: `go get github.com/spf13/viper@latest`
- [x] 1.5 Run `go mod tidy` to clean up dependencies

## 2. Root Command Implementation
- [x] 2.1 Create cmd/spekka/main.go with minimal main function
- [x] 2.2 Create internal/cli/root.go with root command definition
- [x] 2.3 Add version flag (--version, -v) with placeholder version string
- [x] 2.4 Add verbose/debug flag (--verbose, --debug) for log level control
- [x] 2.5 Add quiet flag (--quiet) for minimal output
- [x] 2.6 Configure root command help text and usage information
- [x] 2.7 Set up log/slog with configurable log level based on flags

## 3. Configuration System
- [x] 3.1 Create internal/config/config.go with Config struct
- [x] 3.2 Define .spekka.yaml structure (project, standards, product, specs, tasks paths)
- [x] 3.3 Implement LoadConfig function using Viper
- [x] 3.4 Add configuration file discovery (current dir, home dir, etc.)
- [x] 3.5 Implement configuration validation for required fields
- [x] 3.6 Add error handling for missing or invalid config files
- [x] 3.7 Write unit tests for config loading and validation

## 4. Install Command
- [x] 4.1 Create internal/cli/install.go with install command
- [x] 4.2 Add command description and help text
- [x] 4.3 Define flags (if any needed for future use)
- [x] 4.4 Implement placeholder RunE function with "not yet implemented" message
- [x] 4.5 Register command with root command

## 5. Setup Project Command
- [x] 5.1 Create internal/cli/setup.go with setup-project command
- [x] 5.2 Add command description and help text
- [x] 5.3 Define flags for future configuration options
- [x] 5.4 Implement placeholder RunE function with "not yet implemented" message
- [x] 5.5 Register command with root command

## 6. Create Spec Command
- [x] 6.1 Create internal/cli/createspec.go with create-spec command
- [x] 6.2 Add command description and help text
- [x] 6.3 Define flags for spec creation options
- [x] 6.4 Implement placeholder RunE function with "not yet implemented" message
- [x] 6.5 Register command with root command

## 7. Implement Tasks Command
- [x] 7.1 Create internal/cli/implement.go with implement-tasks command
- [x] 7.2 Add command description and help text
- [x] 7.3 Define flags for task implementation options
- [x] 7.4 Implement placeholder RunE function with "not yet implemented" message
- [x] 7.5 Register command with root command

## 8. Serve Command
- [x] 8.1 Create internal/cli/serve.go with serve command
- [x] 8.2 Add command description and help text
- [x] 8.3 Define flags (--port, --host, etc.)
- [x] 8.4 Implement placeholder RunE function with "not yet implemented" message
- [x] 8.5 Register command with root command

## 9. Error Handling
- [x] 9.1 Implement error wrapping pattern using fmt.Errorf with %w
- [x] 9.2 Create helper function for formatting user-facing errors
- [x] 9.3 Ensure all errors write to stderr
- [x] 9.4 Set proper exit codes for different error types
- [x] 9.5 Add error handling tests

## 10. Testing
- [x] 10.1 Write unit tests for root command initialization
- [x] 10.2 Write unit tests for configuration loading
- [x] 10.3 Write integration tests for each command help output
- [x] 10.4 Write tests for flag parsing
- [x] 10.5 Write tests for error scenarios (missing config, invalid YAML, etc.)
- [x] 10.6 Ensure test coverage ≥80% for core packages

## 11. Documentation
- [x] 11.1 Add package documentation comments to all packages
- [x] 11.2 Add doc comments to all exported functions and types
- [x] 11.3 Create example .spekka.yaml file with comments
- [x] 11.4 Add usage examples to command help text

## 12. Build Configuration
- [x] 12.1 Test build with `go build ./cmd/spekka`
- [x] 12.2 Test cross-platform builds (linux, darwin, windows)
- [x] 12.3 Verify binary executes on target platforms
- [x] 12.4 Add .gitignore for build artifacts

## Verification (Required)
- [x] Run `golangci-lint run` - all errors resolved
- [x] Run `go test ./...` - all tests passing
- [x] Run `go build ./cmd/spekka` - builds successfully
- [x] Check coverage `go test -cover ./...` - ≥80% for modified packages
- [x] Manual smoke test - feature works as intended, no obvious regressions
  - [x] Run `spekka --help` and verify output
  - [x] Run `spekka --version` and verify output
  - [x] Run each subcommand with `--help` flag
  - [x] Test with valid .spekka.yaml file
  - [x] Test with missing .spekka.yaml file
  - [x] Test with invalid .spekka.yaml file
