# Spec Implementation Process

Now that we have a spec and tasks list ready for implementation, we will proceed with implementation of this spec by following this multi-phase process:

PHASE 1: Determine which task group(s) from tasks.md should be implemented
PHASE 2: Delegate implementation to the implementer subagent
PHASE 3: After ALL task groups have been implemented, delegate to implementation-verifier to produce the final verification report.

Follow each of these phases and their individual workflows IN SEQUENCE:

## Multi-Phase Process

### PHASE 1: Determine which task group(s) to implement

First, check if the user has already provided instructions about which task group(s) to implement.

**If the user HAS provided instructions:** Proceed to PHASE 2 to delegate implementation of those specified task group(s) to the **implementer** subagent.

**If the user has NOT provided instructions:**

Read `spekka/specs/[this-spec]/tasks.md` to review the available task groups, then output the following message to the user and WAIT for their response:

```
Should we proceed with implementation of all task groups in tasks.md?

If not, then please specify which task(s) to implement.
```

### PHASE 2: Delegate implementation to the implementer subagent

Delegate to the **implementer** subagent to implement the specified task group(s):

Provide to the subagent:
- The specific task group(s) from `spekka/specs/[this-spec]/tasks.md` including the parent task, all sub-tasks, and any sub-bullet points
- The path to this spec's documentation: `spekka/specs/[this-spec]/spec.md`
- The path to this spec's requirements: `spekka/specs/[this-spec]/planning/requirements.md`
- The path to this spec's visuals (if any): `spekka/specs/[this-spec]/planning/visuals`

Instruct the subagent to:
1. Analyze the provided spec.md, requirements.md, and visuals (if any)
2. Analyze patterns in the codebase according to its built-in workflow
3. IF your task involves user-facing UI, and IF any of the views and UI elements that you will be modifying already exist, and IF you have access to a browser testing tool (such as the Playwright MCP server), THEN
  - Using this browser testing tool, use all of these pre-existing views and/or UI elements as if you are a user to capture their visual appearance BEFORE your change.
  - Take before-change screenshots of these views and UI elements you will be modifying and store those in `spekka/specs/[this-spec]/verification/screenshots/before-change/`. Do not store before-change screenshots anywhere in the codebase other than this location.
4. Implement the assigned task group according to requirements
5. Update `spekka/specs/[this-spec]/tasks.md` to update the tasks you've implemented to mark that as done by updating their checkbox to checked state: `- [x]`

### PHASE 3: Produce the final verification report

IF ALL task groups in tasks.md are marked complete with `- [x]`, then proceed with this step. Otherwise, return to PHASE 1.

Assuming all tasks are marked complete, then delegate to the **implementation-verifier** subagent to do its implementation verification and produce its final verification report.

Provide to the subagent the following:
- The path to this spec: `spekka/specs/[this-spec]`
Instruct the subagent to do the following:
  1. Run all of its final verifications according to its built-in workflow
  2. Produce the final verification report in `spekka/specs/[this-spec]/verifications/final-verification.md`.
