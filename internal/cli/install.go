package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

// spekkaDirectories defines the directory structure to create during installation
var spekkaDirectories = []string{
	".spekka/standards",
	".spekka/product",
	".spekka/specs",
	".spekka/tasks",
	".spekka/skills",
	".spekka/profiles",
	".spekka/verification",
	".spekka/visuals",
}

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Bootstrap repository with Spekka structure and config",
	Long: `The install command bootstraps a repository with the Spekka directory
structure and creates an initial .spekka.yaml configuration file.

This command should be run once per repository to set up Spekka for use.`,
	RunE: runInstall,
}

// runInstall executes the install command logic
func runInstall(cmd *cobra.Command, args []string) error {
	force, _ := cmd.Flags().GetBool("force")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	// Check for existing installation
	if !force {
		if exists, err := configFileExists(); err != nil {
			return fmt.Errorf("failed to check for existing installation: %w", err)
		} else if exists {
			return fmt.Errorf("Spekka is already installed in this directory.\nUse --force to reinstall or choose a different directory")
		}
	}

	if dryRun {
		return runDryRun()
	}

	// Backup existing config if force is used
	if force {
		if err := backupExistingConfig(); err != nil {
			return fmt.Errorf("failed to backup existing configuration: %w", err)
		}
	}

	// Create directory structure
	if err := createDirectoryStructure(); err != nil {
		return fmt.Errorf("failed to create directory structure: %w", err)
	}

	// Generate configuration file
	if err := generateConfigFile(); err != nil {
		return fmt.Errorf("failed to generate configuration file: %w", err)
	}

	// Create template files
	if err := createTemplateFiles(); err != nil {
		return fmt.Errorf("failed to create template files: %w", err)
	}

	// Display success message
	displaySuccessMessage()

	return nil
}

// configFileExists checks if .spekka/config.yaml exists in the current directory
func configFileExists() (bool, error) {
	_, err := os.Stat(".spekka/config.yaml")
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// runDryRun shows what would be done without actually doing it
func runDryRun() error {
	fmt.Println("Dry run mode - showing planned actions:")
	fmt.Println()
	fmt.Println("Directories to be created:")
	for _, dir := range spekkaDirectories {
		fmt.Printf("  - %s/\n", dir)
	}
	fmt.Println()
	fmt.Println("Files to be created:")
	fmt.Println("  - .spekka/config.yaml")
	for _, dir := range spekkaDirectories {
		fmt.Printf("  - %s/README.md\n", dir)
	}
	fmt.Println()
	fmt.Println("Configuration file preview:")
	fmt.Println(getConfigTemplate())
	return nil
}

// backupExistingConfig creates a backup of existing .spekka/config.yaml
func backupExistingConfig() error {
	if exists, _ := configFileExists(); !exists {
		return nil
	}

	timestamp := time.Now().Format("20060102-150405")
	backupName := fmt.Sprintf(".spekka/config.yaml.backup-%s", timestamp)
	
	if err := os.Rename(".spekka/config.yaml", backupName); err != nil {
		return err
	}
	
	fmt.Printf("Existing configuration backed up to: %s\n", backupName)
	return nil
}

// createDirectoryStructure creates all required directories
func createDirectoryStructure() error {
	for _, dir := range spekkaDirectories {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	return nil
}

// generateConfigFile creates the .spekka/config.yaml configuration file
func generateConfigFile() error {
	config := getConfigTemplate()
	return os.WriteFile(".spekka/config.yaml", []byte(config), 0644)
}

// getConfigTemplate returns the default configuration template
func getConfigTemplate() string {
	return `# Spekka Configuration File
# This file configures your Spekka project structure and settings

project:
  # Project name - update this to match your project
  name: "my-project"
  # Project version - semantic versioning recommended
  version: "0.1.0"

# Standards layer - coding standards and style guides
standards:
  # Local path to standards directory (relative to .spekka/)
  path: "standards"
  # Remote style guides (Git URLs) - examples:
  # remote:
  #   - "https://github.com/your-org/go-standards.git"
  #   - "https://github.com/your-org/frontend-standards.git"

# Product layer - product vision and user documentation
product:
  # Local path to product directory (relative to .spekka/)
  path: "product"

# Specifications layer - detailed feature specifications
specs:
  # Local path to specifications directory (relative to .spekka/)
  path: "specs"

# Tasks layer - implementation task breakdown
tasks:
  # Local path to tasks directory (relative to .spekka/)
  path: "tasks"
`
}

// createTemplateFiles creates README.md files in each directory
func createTemplateFiles() error {
	templates := map[string]string{
		".spekka/standards/README.md": getStandardsTemplate(),
		".spekka/product/README.md":   getProductTemplate(),
		".spekka/specs/README.md":     getSpecsTemplate(),
		".spekka/tasks/README.md":     getTasksTemplate(),
		".spekka/skills/README.md":    getSkillsTemplate(),
		".spekka/profiles/README.md":  getProfilesTemplate(),
		".spekka/verification/README.md": getVerificationTemplate(),
		".spekka/visuals/README.md":   getVisualsTemplate(),
	}

	for filePath, content := range templates {
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create template file %s: %w", filePath, err)
		}
	}

	// Create .gitkeep files to preserve empty directories in Git
	for _, dir := range spekkaDirectories {
		gitkeepPath := filepath.Join(dir, ".gitkeep")
		if err := os.WriteFile(gitkeepPath, []byte(""), 0644); err != nil {
			return fmt.Errorf("failed to create .gitkeep in %s: %w", dir, err)
		}
	}

	return nil
}

// displaySuccessMessage shows installation success and next steps
func displaySuccessMessage() {
	fmt.Println("✅ Spekka installation completed successfully!")
	fmt.Println()
	fmt.Println("Created directories:")
	for _, dir := range spekkaDirectories {
		fmt.Printf("  - %s/\n", dir)
	}
	fmt.Println()
	fmt.Println("Created files:")
	fmt.Println("  - .spekka/config.yaml (configuration)")
	for _, dir := range spekkaDirectories {
		fmt.Printf("  - %s/README.md (template)\n", dir)
	}
	fmt.Println()
	fmt.Println("🚀 Next steps:")
	fmt.Println("  1. Review and customize .spekka/config.yaml for your project")
	fmt.Println("  2. Run 'spekka setup-project' to configure project-specific settings")
	fmt.Println("  3. Add your coding standards to the .spekka/standards/ directory")
	fmt.Println("  4. Document your product vision in the .spekka/product/ directory")
	fmt.Println("  5. Start creating specifications with 'spekka create-spec'")
	fmt.Println()
	fmt.Println("📚 Documentation: https://github.com/jreitano/spekka#getting-started")
}

// Template functions for directory README files

func getStandardsTemplate() string {
	return `# Standards Directory

This directory contains coding standards, style guides, and architectural patterns that define **how** code should be written in your project.

## Purpose
The standards layer is part of Spekka's 3-layer context system and provides:
- Coding style guides and conventions
- Architectural patterns and best practices  
- Code review guidelines
- Quality standards and metrics

## Organization
- Create subdirectories for different languages/frameworks (e.g., go/, frontend/, backend/)
- Use markdown files for documentation
- Include example code snippets and templates
- Reference external style guides via Git URLs in .spekka.yaml

## Git URL-based Style Guides
You can import and share style guides using Git URLs (similar to Go package imports):

` + "```yaml" + `
standards:
  remote:
    - "https://github.com/your-org/go-standards.git"
    - "https://github.com/your-org/frontend-standards.git"
` + "```" + `

## Examples
- go/style-guide.md - Go coding conventions
- frontend/react-patterns.md - React component patterns
- backend/api-design.md - REST API design guidelines
- security/guidelines.md - Security best practices
`
}

func getProductTemplate() string {
	return `# Product Directory

This directory contains product vision, user personas, and business context that defines **why** you're building features.

## Purpose
The product layer is part of Spekka's 3-layer context system and provides:
- Product mission and vision
- User personas and journey maps
- Business goals and success metrics
- Market research and competitive analysis

## Recommended Files
- mission.md - Product mission statement and core values
- roadmap.md - Product roadmap and feature priorities
- personas.md - User personas and target audience
- metrics.md - Key performance indicators and success metrics

## Template Structure
Create a mission.md file with sections like:
- Vision Statement
- Target Users
- Core Problems Solved
- Success Metrics
- Competitive Landscape

This context helps AI assistants understand the business rationale behind technical decisions.
`
}

func getSpecsTemplate() string {
	return `# Specifications Directory

This directory contains detailed feature specifications that define **what** to build.

## Purpose
The specifications layer is part of Spekka's 3-layer context system and provides:
- Detailed feature requirements
- User stories and acceptance criteria
- API specifications and data models
- Integration requirements

## Specification Format
Use this structure for specification files:

` + "```markdown" + `
# Feature Name

## Purpose
Brief description of what this feature does and why it's needed.

## Requirements
### Requirement: Feature Name
The system SHALL provide [specific capability].

#### Scenario: Success case
- **WHEN** user performs action
- **THEN** expected result occurs
- **AND** additional conditions are met

#### Scenario: Error case
- **WHEN** invalid input is provided
- **THEN** appropriate error message is shown
` + "```" + `

## Organization
- One specification file per major feature
- Use clear, testable requirements
- Include both positive and negative scenarios
- Reference related specifications when needed

## Integration with Tasks
Specifications are automatically parsed by Spekka to generate implementation tasks with effort estimates and dependencies.
`
}

func getTasksTemplate() string {
	return `# Tasks Directory

This directory contains implementation task breakdowns with effort estimates, dependencies, and acceptance criteria.

## Purpose
Tasks are generated from specifications and provide:
- Actionable implementation steps
- Effort estimates and time tracking
- Dependency management
- Progress tracking and status updates

## Task File Format
Tasks use markdown with YAML frontmatter:

` + "```yaml" + `---
id: "feature-login-ui"
title: "Implement login UI component"
spec: "user-authentication"
effort: "4h"
status: "todo"
assignee: "developer@example.com"
dependencies: ["feature-auth-api"]
---

# Task: Implement Login UI Component

## Description
Create a reusable login component with form validation and error handling.

## Acceptance Criteria
- [ ] Form validates email format
- [ ] Form validates password requirements
- [ ] Error messages display clearly
- [ ] Component is responsive
- [ ] Unit tests achieve >80% coverage

## Implementation Notes
- Use existing design system components
- Follow accessibility guidelines
- Integrate with authentication API
` + "```" + `

## Status Values
- todo: Not started
- in-progress: Currently being worked on
- review: Ready for code review
- done: Completed and merged

## Integration
Tasks sync bidirectionally with external issue trackers (GitHub Issues, Linear, JIRA) via webhooks.
`
}

func getSkillsTemplate() string {
	return `# Skills Directory

This directory contains AI agent skill definitions and specialized capabilities for different types of development work.

## Purpose
Skills define specialized AI agent capabilities for:
- Code generation and refactoring
- Testing and quality assurance
- Documentation and technical writing
- Code review and analysis
- Deployment and DevOps tasks

## Skill Definition Format
Skills are defined using structured templates that specify:
- Input requirements and context
- Expected outputs and deliverables
- Quality criteria and validation steps
- Tool integrations and dependencies

## Examples
- code-generation.md - Code writing and implementation skills
- testing.md - Test creation and quality assurance skills
- documentation.md - Technical writing and documentation skills
- review.md - Code review and analysis skills

## Integration
Skills are used by the multi-agent orchestration system to assign appropriate AI agents to different types of tasks based on their capabilities.
`
}

func getProfilesTemplate() string {
	return `# Profiles Directory

This directory contains AI agent profile configurations for specialized development roles.

## Purpose
Profiles configure AI agents for different development roles:
- Implementation agents (focused on code writing)
- Review agents (focused on code quality and standards)
- Testing agents (focused on test creation and validation)
- Documentation agents (focused on technical writing)

## Profile Configuration
Each profile defines:
- Agent specialization and focus areas
- Context preferences and priorities
- Quality standards and validation criteria
- Communication style and output format

## Examples
- implementation.yaml - Configuration for implementation-focused agents
- review.yaml - Configuration for code review agents
- testing.yaml - Configuration for testing-focused agents
- docs.yaml - Configuration for documentation agents

## Multi-Agent Orchestration
Profiles are used by the orchestration system to:
- Route tasks to appropriate specialized agents
- Configure agent behavior and priorities
- Coordinate collaboration between agents
- Aggregate results from multiple agents
`
}

func getVerificationTemplate() string {
	return `# Verification Directory

This directory contains verification and quality assurance configurations, test results, and compliance documentation.

## Purpose
Verification provides:
- Automated quality checks and validation rules
- Test coverage reports and metrics
- Compliance documentation and audit trails
- Verification status tracking across the project lifecycle

## Verification Types
- spec-completeness.md - Specification completeness checks
- implementation-coverage.md - Implementation coverage validation
- test-coverage.md - Test coverage requirements and reports
- acceptance-criteria.md - Acceptance criteria fulfillment tracking

## Quality Reports
- Generate automated reports on specification quality
- Track implementation coverage against specifications
- Monitor test coverage and quality metrics
- Validate acceptance criteria fulfillment

## Integration
Verification integrates with:
- CI/CD pipelines for automated quality gates
- Issue trackers for verification status updates
- Code review tools for quality enforcement
- Deployment systems for release validation
`
}

func getVisualsTemplate() string {
	return `# Visuals Directory

This directory contains diagrams, architecture visuals, UI mockups, and other visual documentation that supports specifications and design decisions.

## Purpose
Visual documentation provides:
- Architecture diagrams and system overviews
- UI mockups and design specifications
- Flow charts and process diagrams
- Data model and relationship diagrams

## Supported Formats
- Mermaid diagrams (embedded in markdown)
- PlantUML diagrams
- SVG graphics and illustrations
- PNG/JPG images and screenshots

## Organization
- architecture/ - System architecture diagrams
- ui/ - User interface mockups and designs
- flows/ - Process and workflow diagrams
- data/ - Data model and database diagrams

## Integration
Visual documentation integrates with:
- Specifications (embedded diagrams and mockups)
- Design systems and style guides
- Documentation generation tools
- Collaborative design platforms

## Examples
` + "```mermaid" + `
graph TD
    A[User Request] --> B[Authentication]
    B --> C[Business Logic]
    C --> D[Database]
    D --> E[Response]
` + "```" + `
`
}

func init() {
	rootCmd.AddCommand(installCmd)
	
	// Add flags
	installCmd.Flags().Bool("force", false, "Force reinstallation, backing up existing files")
	installCmd.Flags().Bool("dry-run", false, "Show what would be done without executing")
}


