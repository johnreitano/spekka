# Spec Creation Process

You are helping me create a complete specification for a new feature. This combines the spec shaping and spec writing processes into one seamless workflow.

This process will follow 4 main phases, each with their own workflow steps:

Process overview (details to follow)

PHASE 1. Initialize spec
PHASE 2. Research requirements for this spec
PHASE 3. Write the specification document
PHASE 4. Inform the user that the spec is complete

Follow each of these phases and their individual workflows IN SEQUENCE:

## Multi-Phase Process:

### PHASE 1: Initialize Spec

Use the **spec-initializer** subagent to initialize a new spec.

IF the user has provided a description, provide that to the spec-initializer.

The spec-initializer will provide the path to the dated spec folder (YYYY-MM-DD-spec-name) they've created.

### PHASE 2: Research Requirements

After spec-initializer completes, immediately use the **spec-shaper** subagent:

Provide the spec-shaper with:
- The spec folder path from spec-initializer

The spec-shaper will give you several separate responses that you MUST show to the user. These include:
1. Numbered clarifying questions along with a request for visual assets (show these to user, wait for user's response)
2. Follow-up questions if needed (based on user's answers and provided visuals)

**IMPORTANT**:
- Display these questions to the user and wait for their response
- The spec-shaper may ask you to relay follow-up questions that you must present to user

### PHASE 3: Write Specification Document

After requirements gathering is complete, use the **spec-writer** subagent to create the specification document:

Provide the spec-writer with:
- The spec folder path from Phase 1
- The requirements from `planning/requirements.md`
- Any visual assets in `planning/visuals/`

The spec-writer will create `spec.md` inside the spec folder.

### PHASE 4: Inform the User

After all steps complete, inform the user:

```
Spec creation is complete!

✅ Spec folder created: `[spec-path]`
✅ Requirements gathered
✅ Visual assets: [Found X files / No files provided]
✅ Spec document created: `[spec-path]/spec.md`

NEXT STEP 👉 Run `/spekka:create-tasks` to generate your tasks list for this spec.
```
