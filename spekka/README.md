# Spekka

An simple-to-learn Spec-Driven Development system built on Claude Code and Agent OS.

## Acknowledgement

This project is based on Agent OS, with some simplifcations and some enhancements. See the [Agent OS docs](https://buildermethods.com/agent-os/workflow) for more info on the workflow and commands that make up this project.

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
    - Generates `spekka/project/{mission,roadmap,tech-stack}.md` (review these carefully)

- Customize skills for your tech stack
    - Note: This is a new command added by Spekka, and not part of Agent OS.
    - In Claude Code, run: `/spekka:customize-skills` 
    - Updates `.claude/skills/*/*.md` (review these carefully)

## Build a feature

- Create the spec
    - In Claude Code, run: `/shape-os:create-spec`
    - Asks clarifying questions
    - Records questions and your answers in `spec/my-new-feature/planning/requirements.md` 
    - Generates `spec/my-new-feature/spec.md` (review this carefully)
- Create the tasks
    - In Claude Code, run: `/spekka:create-tasks`
    - Generates `spec/my-new-feature/tasks.md`
- Implement the tasks
    - In Claude Code, run: `/spekka:implement-tasks`
    - Generates code for your feature (review this code carefully)
    - Checks off items in `spec/my-new-feature/tasks.md`

    