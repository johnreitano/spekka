## Product Planning Process

You are helping to plan and document the mission, roadmap and tech stack for the current product.  This will include:

- **Gathering Information**: The user's product vision, user personas, problems and key features
- **Mission Document**: Take what you've gathered and create a concise mission document
- **Roadmap**: Create a phased development plan with prioritized features
- **Tech stack**: Establish the technical stack used for all aspects of this product's codebase

This process will create these files in `spekka/product/` directory.

### PHASE 1: Gather Product Requirements

Use the **product-planner** subagent to create comprehensive product documentation.

IF the user has provided any details in regards to the product idea, its purpose, features list, target users and any other details then provide those to the **product-planner** subagent.

The product-planner will:
- Confirm (or gather) product idea, features, target users, confirm the tech stack and gather other details
- Create `spekka/product/mission.md` with product vision and strategy
- Create `spekka/product/roadmap.md` with phased development plan
- Create `spekka/product/tech-stack.md` documenting all of this product's tech stack choices

### PHASE 2: Inform the user

After all steps are complete, output the following to inform the user:

```
Your product planning is all set!

✅ Product mission: `spekka/product/mission.md`
✅ Product roadmap: `spekka/product/roadmap.md`
✅ Product tech stack: `spekka/product/tech-stack.md`

NEXT STEPS:
 👉 First, review and edit the files above, especially `spekka/product/tech-stack.md`
 👉 Then, run `/spekka:customize-skills` to customize skills for your tech-stack!
```
