# Model-Agnostic Skills System

A guide to implementing Claude Code's progressive disclosure skill system for any LLM.

## Overview

Claude Code uses a **skills system** that dynamically loads specialized instructions only when needed. This prevents context bloat while maintaining discoverability. The pattern is model-agnostic and can be replicated for GPT 5.2, Gemini, or other models.

---

## How Claude Code's Skill System Works

### Progressive Disclosure Pattern

Skills implement a three-tier information release strategy:

| Tier | When Loaded | Token Cost | Content |
|------|-------------|------------|---------|
| **1. Metadata** | Always in context | ~100 tokens/skill | Name + description only |
| **2. Full Instructions** | On skill invocation | <5000 tokens | Complete SKILL.md content |
| **3. Supporting Files** | On demand | Variable | Scripts, references, assets |

### Tier 1: Skill Registry (Always Present)

```xml
<available_skills>
<skill>
  <name>pdf</name>
  <description>Create and manipulate PDF documents</description>
  <location>managed</location>
</skill>
<skill>
  <name>commit</name>
  <description>Generate git commits with proper formatting</description>
  <location>user</location>
</skill>
</available_skills>
```

### Tier 2: Full Skill Content (Loaded on Invocation)

When the model calls the `Skill` tool, the full `SKILL.md` content is injected into the conversation.

### Tier 3: Supporting Files (Loaded on Demand)

Files from `scripts/`, `references/`, `assets/` directories only load when explicitly referenced during skill execution.

---

## The Two-Message Injection Pattern

When a skill activates, Claude Code injects **two separate messages**:

1. **Visible message** (`isMeta: false`): Status indicator shown to users
   - Example: "The 'pdf' skill is loading"

2. **Hidden message** (`isMeta: true`): Full skill prompt sent to API but hidden from UI
   - Contains the complete SKILL.md content

This dual-channel approach provides transparency without overwhelming users.

---

## Key System Prompt Instructions

From Claude Code's extracted system prompts:

```markdown
When users ask you to perform tasks, check if any of the available
skills below can help complete the task more effectively.

When a skill is relevant, you must invoke this tool IMMEDIATELY
as your first action. NEVER just announce or mention a skill in
your text response without actually calling this tool.

Only use skills listed in <available_skills> below.
Do not invoke a skill that is already running.
```

### Critical Behaviors Enforced

1. **Immediate invocation** - Don't describe, just call the tool
2. **Skill-first priority** - Check skills before responding
3. **Strict allowlist** - Only use registered skills
4. **No re-invocation** - Don't reload already-active skills

---

## SKILL.md Format

### Required Frontmatter

```yaml
---
name: skill-name
description: What this skill does and when to use it.
---
```

**Field Constraints:**
- `name`: 1-64 chars, lowercase alphanumeric + hyphens, no leading/trailing/consecutive hyphens
- `description`: 1-1024 chars, include keywords for discoverability

### Optional Frontmatter

```yaml
---
name: pdf
description: Create and manipulate PDF documents
license: Apache-2.0
compatibility: Requires Python 3.9+, network access for external resources
metadata:
  author: anthropic
  version: 1.2.0
allowed-tools: Bash Write Read
---
```

### Body Content

```markdown
# PDF Creation Skill

When creating PDFs, follow these guidelines:

## Supported Operations
- Create new documents
- Merge existing PDFs
- Extract text and images

## Libraries
Use `reportlab` for simple documents, `PyMuPDF` for manipulation.

## Examples
[Include concrete examples here]
```

### Directory Structure

```
skills/
└── pdf/
    ├── SKILL.md           # Required: main instructions
    ├── scripts/           # Optional: executable code
    │   └── merge_pdfs.py
    ├── references/        # Optional: additional docs
    │   └── api-guide.md
    └── assets/            # Optional: templates, data
        └── template.pdf
```

---

## Implementing for Other Models

### 1. Define the Skill Tool

```json
{
  "name": "invoke_skill",
  "description": "Load specialized instructions for a task. Check available skills and call this IMMEDIATELY when one matches. Available: pdf, commit, review-pr, etc.",
  "parameters": {
    "type": "object",
    "properties": {
      "skill_name": {
        "type": "string",
        "description": "Name of the skill to invoke"
      },
      "args": {
        "type": "string",
        "description": "Optional arguments for the skill"
      }
    },
    "required": ["skill_name"]
  }
}
```

### 2. System Prompt Template

```markdown
# Skill System

You have access to specialized skills that provide detailed instructions
for specific tasks. Skills are listed below with name and description only.

## Available Skills
- **pdf**: Create and manipulate PDF documents
- **commit**: Generate git commits with proper formatting
- **review-pr**: Review pull requests for code quality
- **refactor**: Refactor code following best practices

## Skill Invocation Rules
1. When a task matches a skill, call invoke_skill() IMMEDIATELY
2. Do NOT describe what you'll do first - just call the tool
3. After skill content loads, follow its instructions exactly
4. Only use skills from this list - never guess skill names
5. Do not re-invoke a skill that's already loaded

## After Skill Loads
When you see <skill-content>...</skill-content> tags, those instructions
override your default behavior. Follow them precisely until the task
is complete.
```

### 3. Backend Implementation

```python
import yaml
from pathlib import Path

def invoke_skill(skill_name: str, args: str = None) -> dict:
    """Load and inject skill content into conversation."""
    skill_path = Path(f"skills/{skill_name}/SKILL.md")

    if not skill_path.exists():
        return {"error": f"Skill '{skill_name}' not found"}

    content = skill_path.read_text()

    # Parse frontmatter
    if content.startswith("---"):
        _, frontmatter, body = content.split("---", 2)
        metadata = yaml.safe_load(frontmatter)
    else:
        body = content
        metadata = {}

    # Return as a message to inject
    return {
        "role": "user",
        "content": f"""<skill-content name="{skill_name}">
{body.strip()}
</skill-content>

Follow the instructions above to complete the user's task."""
    }

def get_skill_registry() -> str:
    """Generate the available skills list for the system prompt."""
    skills_dir = Path("skills")
    registry = []

    for skill_path in skills_dir.iterdir():
        if skill_path.is_dir():
            skill_md = skill_path / "SKILL.md"
            if skill_md.exists():
                content = skill_md.read_text()
                if content.startswith("---"):
                    _, frontmatter, _ = content.split("---", 2)
                    meta = yaml.safe_load(frontmatter)
                    registry.append({
                        "name": meta.get("name", skill_path.name),
                        "description": meta.get("description", "No description")
                    })

    # Format as XML for the system prompt
    xml = "<available_skills>\n"
    for skill in registry:
        xml += f"""<skill>
  <name>{skill['name']}</name>
  <description>{skill['description']}</description>
</skill>\n"""
    xml += "</available_skills>"

    return xml
```

### 4. Conversation Flow

```
1. User sends message
2. System prompt includes skill registry (Tier 1)
3. Model recognizes task matches a skill
4. Model calls invoke_skill("pdf")
5. Backend loads SKILL.md, injects as message (Tier 2)
6. Model follows skill instructions
7. If needed, model requests supporting files (Tier 3)
8. Task completes, skill context naturally expires
```

---

## Best Practices

### Skill Design

1. **Keep SKILL.md under 500 lines** - Use references/ for lengthy docs
2. **Front-load critical instructions** - Most important info first
3. **Include concrete examples** - Show don't tell
4. **Use clear section headers** - Helps model navigate

### Description Writing

Good descriptions enable accurate skill selection:

```yaml
# Good - specific, includes keywords
description: Create, merge, split, and manipulate PDF documents.
  Supports text extraction, image insertion, and form filling.

# Bad - vague, missing keywords
description: Work with PDF files.
```

### Token Budget Guidelines

| Component | Target | Maximum |
|-----------|--------|---------|
| Frontmatter | 50 tokens | 100 tokens |
| Main instructions | 2000 tokens | 5000 tokens |
| Per reference file | 1000 tokens | 3000 tokens |

---

## Resources

### Official Anthropic Resources
- [anthropics/skills](https://github.com/anthropics/skills) - Official skill examples and templates
- [Agent Skills Specification](https://agentskills.io/specification) - Formal specification document

### Community Resources
- [Piebald-AI/claude-code-system-prompts](https://github.com/Piebald-AI/claude-code-system-prompts) - Extracted Claude Code prompts, updated each release
- [Claude Skills Deep Dive](https://leehanchung.github.io/blogs/2025/10/26/claude-skills-deep-dive/) - Technical analysis of the architecture
- [awesome-claude-skills](https://github.com/travisvn/awesome-claude-skills) - Curated list of community skills

### Key Insights from Deep Dive

> "Skills are specialized prompt templates that inject domain-specific instructions into the conversation context."

> "Progressive Disclosure - showing just enough information to help agents decide what to do next, then reveal more details."

> "Skill selection happens through Claude's native language understanding, not algorithmic matching."

---

## Summary

The skill system pattern is fundamentally simple:

1. **Registry in system prompt** - Names + descriptions only
2. **Tool for loading** - Injects full content on-demand
3. **Clear instructions** - "Invoke immediately, don't describe"
4. **Structured format** - YAML frontmatter + markdown body

This works across models because it relies on:
- Standard tool/function calling
- Clear behavioral instructions
- Structured content injection

The model-specific tuning is minimal - mainly adjusting the system prompt language to match each model's instruction-following style.
