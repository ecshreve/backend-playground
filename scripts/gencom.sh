#!/bin/sh

# Navigate to the root of the Git repository
repo_root=$(git rev-parse --show-toplevel)
cd "$repo_root"

# Check if there are too many changes to commit
DIFF_LINES=$(git diff --cached | wc -l)
if [ "$DIFF_LINES" -gt 500 ]; then
  echo "Too many changes to send to mods. Please commit fewer changes at a time."
  exit 1
fi

# Check if env vars are set
if [ -z "$OPENAI_API_KEY" ]; then
    echo "OPENAI_API_KEY is not set, not generating commit message."
    gum confirm "continue without generating commit message?" || exit 1
fi

# Check if mods config matches the file in the repo
if ! diff ~/.config/mods/mods.yml mods-config.yml > /dev/null; then
    echo "mods.yml does not match the file in the repo"
    cp ~/.config/mods/mods.yml ~/.config/mods/mods.yml.bak
    cp mods-config.yml ~/.config/mods/mods.yml
    echo "mods.yml updated"
else
    echo "mods.yml up to date"
fi

OUT=$(gum spin --spinner line --show-output --title "generating commit..." -- sh -c 'git diff --cached | mods --role gencom')

gum style --border double --margin "1" --padding "1 1" --border-foreground 35 --width 80 "$OUT"

gum confirm "commit?" && git commit -m "$OUT" && exit 0

gum confirm "edit?" && git commit -m "$OUT" --edit && exit 0

echo "exiting without commit"


