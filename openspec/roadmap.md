# Product Roadmap

1. [x] Core CLI Framework & Project Structure — Implement the base CLI application using Cobra with commands for install, setup-project, create-spec, implement-tasks, and serve. Set up standard Go project layout (cmd/, internal/, pkg/) and create configuration file parsing for .spekka.yaml files. `M`

2. [-] Install Command & Directory Scaffolding — Create the install command that bootstraps a repository with the spekka directory structure including standards/, product/, specs/, tasks/, skills/, profiles/, verification/, and visuals/ directories. Generate initial template files and README documentation. `S`

3. [ ] Standards Management System — Build the system for loading, validating, and enforcing coding standards from markdown files. Implement Git URL-based style guide importing (similar to Go package imports) with caching and version pinning. Create validation engine that checks code against defined standards. `L`

4. [ ] Spec Creation Workflow (Plan Product Phase) — Implement create-spec command with interactive prompts to gather product information (idea, features, users, tech stack). Generate mission.md, roadmap.md, and tech-stack.md files following defined templates. Include validation against project standards. `M`

5. [ ] Specification Shaping & Structuring (Shape Spec Phase) — Build tooling to help users refine and structure specifications with proper formatting, completeness checks, and template enforcement. Implement spec validation that checks for required sections, clarity, and alignment with product mission. `M`

6. [ ] Task Breakdown & Management (Create Tasks Phase) — Create the implement-tasks command that parses specifications and breaks them into actionable tasks with effort estimates, dependencies, and acceptance criteria. Store tasks as markdown files with YAML frontmatter. Implement task status tracking and progress reporting. `L`

7. [ ] Context Injection System — Build the 3-layer context system that intelligently selects and injects relevant context (standards, product vision, specs) into LLM prompts based on current task. Create context optimization to maximize value within token limits. Implement templates for different task types. `L`

8. [ ] Multi-Agent Orchestration (Orchestrate Tasks Phase) — Implement agent profile system for configuring specialized AI agents (implementation, review, testing). Build task queue and coordination system for parallel task execution. Create inter-agent communication protocol and result aggregation. `XL`

9. [ ] Webhook Server & Issue Tracker Integration — Build the serve command that runs an HTTP server processing webhooks from GitHub Issues, Linear, and other trackers. Implement event parsing, task synchronization, and bidirectional updates between external systems and local markdown files. Include authentication and rate limiting. `L`

10. [ ] Web Interface for Task Management — Create a lightweight web UI using a Go-based framework (Echo or Chi with Go templates) for viewing project status, managing issues and tasks, and monitoring agent activity. Implement real-time updates using WebSockets. No database required - read directly from markdown files. `M`

11. [ ] Style Guide Enforcement & Code Review — Implement automated code validation against team style guides during implementation and review. Create pre-commit hooks and CI integration. Build reporting system for standards violations with actionable fix suggestions. `M`

12. [ ] Verification & Quality Assurance System — Build automated verification system that checks spec completeness, implementation coverage, test coverage, and acceptance criteria fulfillment. Generate quality reports and track verification status across the project lifecycle. `M`

13. [ ] Workflow Templates & Reusable Patterns — Create library of reusable workflow templates for common scenarios (feature development, bug fixes, refactoring, technical debt). Implement template discovery, customization, and sharing via Git URLs. `S`

14. [ ] Visual Documentation Support — Add support for embedding and managing diagrams, architecture visuals, UI mockups, and other visual aids in specifications. Integrate with common diagramming formats (Mermaid, PlantUML, SVG). `S`

15. [ ] CLI Experience Enhancements — Improve CLI user experience with better error messages, progress indicators, interactive prompts with validation, autocomplete support, and comprehensive help documentation. Add verbose/debug logging modes. `S`

16. [ ] CI/CD Integration & Automation — Create GitHub Actions, GitLab CI, and CircleCI integrations for running Spekka validation in pipelines. Implement automated spec checks, standards enforcement, and task status updates on commits and PRs. `M`

17. [ ] Configuration & Customization — Build comprehensive configuration system supporting project-level, user-level, and global settings. Implement config inheritance and environment-specific overrides. Support multiple config formats (YAML, JSON, TOML). `S`

18. [ ] Documentation & Examples — Create comprehensive user documentation including getting started guide, command references, workflow tutorials, and integration examples. Build sample projects demonstrating common use cases and best practices. `M`

> Notes
> - Order items by technical dependencies and product architecture
> - Each item should represent an end-to-end (frontend + backend) functional and testable feature
