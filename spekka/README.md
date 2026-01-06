# Spekka

An simple-to-learn Spec-Driven Development system built on Claude Code and Agent OS.

## Acknowledgement

This project is based on Agent OS, with some simplifcations and some enhancements. See the [Agent OS docs](https://buildermethods.com/agent-os/workflow) for more info on the design workflow and commands that make up this project.

## Install (or re-install)

In your project directory, run:
```
curl -sSL "https://raw.githubusercontent.com/johnreitano/spekka/refs/heads/main/scripts/install.sh" | bash
```

The will install files into your `.claude/` directory similar to the following:

```
.claude/
├── agents/spekka/
│   ├── implementation-verifier.md
│   ├── implementer.md
│   └── ...
├── commands/spekka/
│   ├── create-tasks.md
│   ├── implement-tasks.md
│   └── ...
└── skills/
    ├── backend-api/
    │   └── SKILL.md
    ├── backend-migrations/
    │   └── SKILL.md
    └── ...
```

## Set up project and standards

- Configure core project info
    - In Claude Code, run: `/spekka:plan-product`
    - Generates `spekka/project/{mission,roadmap,tech-stack}.md` (review these carefully!)

- Customize standards for your tech stack
    - In Claude Code, run: `/spekka:customize-standards`
    - Generates `spekka/standards/**/*.md` (review these carefully!)

## Build a feature

- Create the spec
    - In Claude Code, run: `/shape-os:shape-spec`
        - Asks clarifying questions
        - Generates
            - `spec/my-new-feature/requirements.md`
    - In Claude Code, run: `/spekka:write-spec`
        - Generates
        - `spec/my-new-feature/spec.md` (review this carefully!)
- Create the tasks
    - In Claude Code, run: `/spekka:create-tasks`
    - Generates `spec/my-new-feature/tasks.md`
- Implement the tasks
    - In Claude Code, run: `/spekka:implement-tasks`
    - Generates code for your feature (review this code extra carefully!)
    