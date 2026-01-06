## Customize standards for the project's tech stack

Update the standards under `spekka/standards/**/*.md` based on well-known best practices for the technologies that comprise this project's tech stack.

CRITICAL EXECUTION REQUIREMENTS:
- Create a TodoWrite list with the 4 phases below BEFORE starting
- Execute phases SEQUENTIALLY - mark each as in_progress ONLY when starting
- **IMMEDIATELY after completing ANY phase (including skill work), call TodoWrite to mark it completed**
- **IMMEDIATELY after marking a phase completed, start the next phase**
- If ANY phase fails, STOP and report the failure to the user
- NEVER wait for user input between phases

### PHASE 1: Customize standards

The content in the files `spekka/standards/**/*.md` can be divided into three categories:
1. Content that can usefully be updated or extended to be specific to the technologies in this project's tech stack (`spekka/product/tech-stack.md`). Update or extend this content in such a way that when subsequently used in a prompt to generate code, the code will comply with the well-known best practices for the relevant technologies. For example, the generic statement to "Use consistent indentation (spaces or tabs) and configure your editor/linter to enforce it" in `standards/global/coding-style.md` should be changed to make more specific recommendations for the language(s) used in the project's tech stack.
2. Content that doesn't need to be customized because it is already applicable, as is, to the tech stack. Leave this content unchanged. For example, this category would include the all of the content in the version control standards (`standards/global/version-control.md`) for most projects.
3. Content that should be removed because it is not applicable to the tech stack. Remove this content. For example, the recommendation to "Use Transactions for Related Changes" in `standards/backend/queries.md` should be removed for an app that has a first-class design principle of using a read-only database.

**After completing:** Mark Phase 1 as completed, then immediately start Phase 2.

### PHASE 2: Refine standards

Update the standards (in `spekka/standards/**/*.md`) to comply with the guidelines below.

#### Keep standards concise

Ensure the new standards are written in such a way that they do not use any more of the token context window than necessary.

#### Keep standards modular

Use modular files that are specific and focused files rather than monolithic.
- Bad: A monolithic file such as `spekka/standards/all-standards.md`.
- Good: Modules such as the following:
    - `spekka/standards/frontend/react-hooks.md`
    - `spekka/standards/frontend/component-props.md`
    - `spekka/standards/frontend/state-management.md`
    - `spekka/standards/frontend/styling.md`
    - `spekka/standards/backend/api.md`
    - `spekka/standards/backend/migrations.md`

#### Be specific and actionable

Keep standards clear and implementable. Bad: Vague standards such as "Write clean code with good naming". Good: Specific standards such as: "Use camelCase for variables. Use descriptive names (getUserById not getUser). Prefix booleans with is/has/should."

#### Provide Examples

Always provide concise examples of how to follow the standard.

**After completing:** Mark Phase 2 as completed, then immediately start Phase 3.

### PHASE 3: Improve associated skills

**Do this work directly - do NOT invoke the /spekka:improve-skills command.**

Improve all skills in `.claude/skills/` by performing these steps for each SKILL.md file:

1. Read the SKILL.md file
2. Read its linked standards file (found in the markdown link)
3. Rewrite the `description` frontmatter to be specific and actionable:
   - First sentence: clearly describe what this skill is
   - Following sentences: describe when this skill should be used with specific examples
   - Include file types, directories, and coding situations where the skill applies
4. Add a "When to use this skill:" section below the frontmatter with a bullet list of use cases

**After completing:** Mark Phase 3 as completed, then immediately start Phase 4.

### PHASE 4: Inform the user

Output the following to inform the user:

```
Your skill customization is complete!

✅ Customized standards: `spekka/standards/**/*.md`
✅ Customized skills: `.claude/skills/**/*.md`

NEXT STEP 👉 Run `/spekka:shape-spec` or `/spekka:write-spec` to start work on a feature!
```

Mark Phase 4 as completed.
