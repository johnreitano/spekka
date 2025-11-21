# Spekka

## Pitch
Spekka is a CLI-based development workflow tool that helps software engineering teams build better products through spec-driven development by providing a structured, AI-assisted approach to planning, specification writing, and implementation that maximizes the value of LLM context windows. It also integrates via webhooks with popular issue tracking systems.

## Users

### Primary Customers
- **Software Engineers**: Individual developers who want to implement spec-driven development workflows with AI assistance
- **Engineering Teams**: Small to medium-sized teams building products who need better coordination between planning, specification, and implementation

### User Personas

**Senior Software Engineer**
- **Role:** Tech Lead or Senior IC on product teams
- **Context:** Works on complex features requiring detailed specs, coordinates with designers and PMs, mentors junior developers
- **Pain Points:**
  - Specifications get out of sync with implementation
  - LLM context windows fill up quickly with unnecessary information
  - Difficult to share and enforce team standards when using AI coding tools
- **Goals:**
  - Streamline spec-driven development workflow
  - Ensure AI tools respect team coding standards
  - Maintain high code quality while moving fast

**Engineering Manager**
- **Role:** Manager of 5-15 person engineering team
- **Context:** Oversees multiple projects, ensures team velocity and code quality, coordinates with product and design
- **Pain Points:**
  - Inconsistent development practices across team members
  - Difficulty tracking work from spec to implementation
  - Code reviews take too long due to style inconsistencies
  - Hard to onboard new team members to team standards
- **Goals:**
  - Standardize team development workflows
  - Improve code consistency and quality
  - Make team standards easily shareable and enforceable

**UX/UI Designer**
- **Role:** Product Designer working with engineering teams
- **Context:** Creates designs and specifications, needs to communicate design requirements to engineers
- **Pain Points:**
  - Designs don't translate well to implementation
  - Hard to track what's been implemented vs. spec
  - Disconnect between design specs and code
- **Goals:**
  - Better handoff from design to engineering
  - Track implementation against design specs
  - Collaborate more effectively with engineers

## The Problem

### Spec-Implementation Disconnect in AI-Assisted Development
Teams using AI coding assistants like Claude Code, Cursor, and Codex struggle to maintain spec-driven development workflows. LLMs have limited context windows that quickly fill with code, leaving little room for specifications, standards, and product vision. This leads to implementations that drift from specs, inconsistent coding styles, and repeated manual intervention to re-establish context.

**Our Solution:** Spekka implements a 3-layer context system (standards, product vision, specifications) and provides CLI commands that inject the right context at the right time. Style guides and standards are enforced automatically, can be version-controlled, and shared via Git URLs like Go packages.

### Fragmented Development Workflow
Traditional development tools separate planning (JIRA, Linear), specification (Confluence, Notion), and implementation (IDE, Git). This fragmentation causes context switching, outdated specs, and difficulty tracking work from concept to completion.

**Our Solution:** Spekka provides an integrated workflow from `create-spec` through `implement-tasks` with webhook integration to existing issue trackers. All specifications and tasks are stored as markdown files in the repository, keeping everything version-controlled and close to the code.

### Coding Standards Enforcement at Scale
Teams struggle to maintain consistent coding standards when using AI tools. Each developer's AI assistant might generate code differently, and manual code reviews become bottlenecks for enforcing standards.

**Our Solution:** Spekka's style guide system allows teams to define standards once and share them across the organization. Style guides are enforced during both code generation and review phases, ensuring consistency without manual intervention.

## Differentiators

### AI-Optimized Context Management
Unlike traditional project management tools, Spekka is designed specifically for AI-assisted development. We optimize what information goes into an LLM's context window and when, ensuring maximum value from limited context space.

### Git-Based Style Guide Sharing
Unlike style guide tools that require custom registries or databases, Spekka uses Git URLs for sharing style guides, similar to Go's package system. This makes it trivial to share standards across teams and organizations while leveraging existing Git infrastructure.

### File-Based, Version-Controlled Specifications
Unlike cloud-based specification tools (Confluence, Notion), Spekka stores all specs and tasks as markdown files in your repository. This keeps specifications version-controlled alongside code and eliminates sync issues between docs and implementation.

### Webhook Integration with Existing Tools
Unlike standalone tools that require full migration, Spekka integrates with existing issue trackers (GitHub Issues, Linear, Trello, JIRA, etc) via webhooks. Teams can continue using their current tools while benefiting from Spekka's spec-driven workflow.

### CLI-First Design
Unlike web-based tools that require context switching, Spekka is designed as a CLI tool that integrates directly into developer workflows. Works seamlessly with Claude Code, Cursor, Codex, and other AI coding assistants.

## Key Features

### Core Workflow Commands
- **install:** Bootstrap a repository with Spekka's directory structure and configuration files, setting up the 3-layer context system (standards, product, specs)
- **setup-project:** Initialize project-level configuration including style guides, team standards, and agent profiles for multi-agent orchestration
- **create-spec:** Generate structured specification documents from user input, enforcing template standards and organizing specs by feature area
- **implement-tasks:** Break down specifications into actionable tasks and execute them with appropriate context injection for AI assistants
- **serve:** Run webhook server to process events from issue tracking systems, synchronizing external tasks with local spec-driven workflow

### Style Guide & Standards Management
- **Style Guide Enforcement:** Automatically validate code against team style guides during creation and review phases
- **Git-Based Sharing:** Import and share style guides using Git URLs (e.g., `spekka import github.com/yourorg/go-standards`)
- **Multi-Language Support:** Define and enforce standards for multiple languages in the same project
- **Automated Validation:** Pre-commit hooks and CI integration to ensure standards compliance

### Specification & Task Management
- **3-Layer Context System:** Organize project information as standards (how), product vision (why), and specifications (what)
- **Markdown-Based Storage:** All specs and tasks stored as markdown files, version-controlled in the repository
- **Task Breakdown:** Automatically decompose specs into implementable tasks with effort estimates and dependencies
- **Progress Tracking:** Track task completion and sync status with external issue trackers

### Integration & Collaboration
- **Webhook Server:** Process events from GitHub Issues, Linear, and other issue trackers to keep specifications in sync
- **Web Interface:** Lightweight web UI for managing issues, tasks, and viewing project status
- **Multi-Agent Orchestration:** Coordinate multiple AI agents for parallel task execution with the orchestration system
- **CI/CD Integration:** Run spec validation and standards checking in continuous integration pipelines

### Advanced Features
- **Agent Profiles:** Configure specialized AI agent behaviors for different types of tasks (implementation, review, testing)
- **Verification System:** Automated quality checks for specs, code, and implementation completeness
- **Workflow Templates:** Reusable workflow patterns for common development scenarios (feature development, bug fixes, refactoring)
- **Visual Documentation:** Support for diagrams and visual aids in specifications (architecture diagrams, UI mockups)
