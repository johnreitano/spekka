#!/usr/bin/env bash
set -euo pipefail

REPO_URL="https://github.com/johnreitano/spekka"
BRANCH="main"

# Check dependencies
command -v git >/dev/null 2>&1 || { echo "Error: git is required"; exit 1; }

TEMP_DIR=$(mktemp -d)

cleanup() {
    rm -rf "$TEMP_DIR"
}
trap cleanup EXIT

echo "Installing Spekka..."

# Clone the repository to temp directory
if ! git clone --depth 1 --branch "$BRANCH" "$REPO_URL" "$TEMP_DIR" >/dev/null 2>&1; then
    echo "Error: Failed to clone repository from $REPO_URL"
    exit 1
fi

SKIP_AGENTS=
SKIP_COMMANDS=
SKIP_SKILLS=
SKIP_STANDARDS=
ASKED_ABOUT_SKILLS=

# Check and copy .claude/agents/spekka files
if [ -d ".claude/agents/spekka" ]; then
    read -p ".claude/agents/spekka already exists. Overwrite? [y/N] " -n 1 -r < /dev/tty || true
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        SKIP_AGENTS=true
    fi
fi

# Check and copy .claude/commands/spekka files
if [ -d ".claude/commands/spekka" ]; then
    read -p ".claude/commands/spekka already exists. Overwrite? [y/N] " -n 1 -r < /dev/tty || true
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        SKIP_COMMANDS=true
    fi
fi

# Check if any .claude/skills/**/*.md files would be overwritten
while IFS= read -r -d '' src_file; do
    # Extract path relative to .claude/skills/ (e.g., "backend-api/SKILL.md")
    partial_path="${src_file#"$TEMP_DIR"/.claude/skills/}"
    if [ -f ".claude/skills/$partial_path" ]; then
        if [[ ! $ASKED_ABOUT_SKILLS ]]; then
            read -p "At least one skill file already exists and would be overwritten. Overwrite? [y/N] " -n 1 -r < /dev/tty || true
            echo
            if [[ ! $REPLY =~ ^[Yy]$ ]]; then
                SKIP_SKILLS=true
            fi
            ASKED_ABOUT_SKILLS=true
        fi
    fi
done < <(find "$TEMP_DIR"/.claude/skills -name "*.md" -type f -print0 2>/dev/null)

# Check and copy spekka/standards files
if [ -d "spekka/standards" ]; then
    read -p "spekka/standards already exists. Overwrite? [y/N] " -n 1 -r < /dev/tty || true
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        SKIP_STANDARDS=true
    fi
fi

if [[ ! $SKIP_AGENTS ]]; then
    mkdir -p .claude/agents
    rm -rf .claude/agents/spekka
    command cp -r "$TEMP_DIR"/.claude/agents/spekka .claude/agents/
fi
if [[ ! $SKIP_COMMANDS ]]; then
    mkdir -p .claude/commands
    rm -rf .claude/commands/spekka
    command cp -r "$TEMP_DIR"/.claude/commands/spekka .claude/commands/
fi
if [[ ! $SKIP_SKILLS ]]; then
    # Copy skills with subdirectory structure preserved
    mkdir -p .claude/skills
    while IFS= read -r -d '' src_file; do
        partial_path="${src_file#"$TEMP_DIR"/.claude/skills/}"
        mkdir -p ".claude/skills/$(dirname "$partial_path")"
        command cp -f "$src_file" ".claude/skills/$partial_path"
    done < <(find "$TEMP_DIR"/.claude/skills -name "*.md" -type f -print0 2>/dev/null)
fi
if [[ ! $SKIP_STANDARDS ]]; then
    mkdir -p spekka
    rm -rf spekka/standards
    command cp -r "$TEMP_DIR"/spekka/standards spekka/
fi

# if all of the variables are true, then exit
if [[ $SKIP_AGENTS && $SKIP_COMMANDS && $SKIP_SKILLS && $SKIP_STANDARDS ]]; then
    echo "No components were installed. Exiting..."
    exit 0
fi

mkdir -p spekka
command cp -f "$TEMP_DIR"/spekka/README.md spekka/
command cp -f "$TEMP_DIR"/LICENSE spekka/
command cp -f "$TEMP_DIR"/THIRD_PARTY_LICENSES spekka/

echo ""
echo "All Spekka-related files are installed!"
echo ""
if [[ ! $SKIP_AGENTS ]]; then
    echo "  Claude-code agents:       .claude/agents/spekka/"
fi
if [[ ! $SKIP_COMMANDS ]]; then
    echo "  Claude-code commands:     .claude/commands/spekka/"
fi
if [[ ! $SKIP_SKILLS ]]; then
    echo "  Claude-code skills:       .claude/skills/"
fi
if [[ ! $SKIP_STANDARDS ]]; then
    echo "  General coding standards: spekka/standards/"
fi
echo "  Instructions for use:     spekka/README.md"

echo ""
echo "NEXT STEP: Restart Claude Code and run /spekka:plan-product to plan your mission, roadmap and tech-stack!"
