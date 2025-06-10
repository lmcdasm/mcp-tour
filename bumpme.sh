#!/bin/bash
#set -x

BUMP_TYPE="$1"
shift
MSG="$*"

if [[ -z "$BUMP_TYPE" || -z "$MSG" ]]; then
  echo "Usage: ./bumpme.sh [point|minor|major] \"tag message\""
  exit 1
fi

BRANCH=$(git rev-parse --abbrev-ref HEAD)
LATEST_TAG=$(git tag -l | grep -E '^[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)?$' | sort -V | tail -n 1)

if [[ -z "$LATEST_TAG" ]]; then
  MAJOR=0
  MINOR=0
  PATCH=0
  SUFFIX="-alpha"
else
  VERSION=$(echo "$LATEST_TAG" | cut -d '-' -f1)
  SUFFIX=$(echo "$LATEST_TAG" | grep -oE '\-[a-zA-Z0-9]+$' || echo "-alpha")
  MAJOR=$(echo "$VERSION" | cut -d. -f1 | sed 's/^0*//')
  MINOR=$(echo "$VERSION" | cut -d. -f2 | sed 's/^0*//')
  PATCH=$(echo "$VERSION" | cut -d. -f3 | sed 's/^0*//')

  # fallback to zero if empty
  MAJOR=${MAJOR:-0}
  MINOR=${MINOR:-0}
  PATCH=${PATCH:-0}
fi

echo "📌 Current version: ${MAJOR}.${MINOR}.${PATCH}${SUFFIX}"

case "$BUMP_TYPE" in
  major)
    ((MAJOR++))
    MINOR=0
    PATCH=0
    ;;
  minor)
    ((MINOR++))
    PATCH=0
    ;;
  point)
    ((PATCH++))
    ;;
  *)
    echo "Invalid bump type: $BUMP_TYPE"
    exit 1
    ;;
esac

NEW_TAG="${MAJOR}.${MINOR}.${PATCH}${SUFFIX}"
echo "🚀 New version will be: $NEW_TAG"

git add -A .
if git diff --cached --quiet; then
  echo "ℹ️ No file changes. Creating empty commit for tag..."
  git commit --allow-empty -m "$MSG"
else
  echo "🔧 Committing staged changes to $BRANCH..."
  git commit -m "$MSG"
fi

echo "🏷️  Tagging commit as $NEW_TAG"
git tag -a "$NEW_TAG" -m "$MSG"

echo "📤 Pushing branch '$BRANCH' and tag '$NEW_TAG'..."
git push origin "$BRANCH"
git push origin "$NEW_TAG"

echo "✅ Done: $NEW_TAG pushed with commit to $BRANCH"

