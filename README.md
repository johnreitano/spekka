# Spekka

An simple-to-learn Spec-Driven Development system built on Claude Code and Agent OS.

## Acknowledgement

This project is based on Agent OS, with some simplifcations and some enhancements. See the [Agent OS workflow docs](https://buildermethods.com/agent-os/workflow) for more info on the workflows and commands that make up this project.

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
