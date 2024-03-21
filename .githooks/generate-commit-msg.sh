#!/bin/sh

COMMIT_MSG_FILE=$1

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
    exit 0
fi

# OUT=$(git diff --cached | mods --fanciness 4 --temp .4 --role comm)
OUT=$(gum spin --spinner line --show-output --title "generating commit..." -- sh -c 'git diff --cached | mods --fanciness 4 --temp .4 --quiet --role gencomm')
echo "$OUT" > "$COMMIT_MSG_FILE.tmp"
cat" $COMMIT_MSG_FILE" >> "$COMMIT_MSG_FILE.tmp"
mv "$COMMIT_MSG_FILE.tmp" "$COMMIT_MSG_FILE"

