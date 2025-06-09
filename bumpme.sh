#!/bin/bash

set -e

BUMP_TYPE=$1
MSG=$2

if [[ -z "$BUMP_TYPE" || -z "$MSG" ]]; then
  echo "Usage: ./bumpme.sh [point|minor|major] \"tag message\""
  exit 1
fi

# Get current branch
BRANCH=$(git rev-parse --abbrev-ref HEAD)

# Get latest semver-style tag
LATEST_TAG=$(git tag -l | grep -E '^[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)?$' | sort -V | tail -n 1)

if [[ -z "$LATEST_TAG" ]]; then
  echo "No existing tag found. Starting at 0.0.0-alpha"
  MAJOR=0
  MINOR=0
  PATCH=0
  SUFFIX="-alpha"
else
  VERSION=$(echo "$LATEST_TAG" | cut -d '-' -f1)
  SUFFIX=$(echo "$LATEST_TAG" | grep -oE '\-[a-zA-Z0-9]+$' || echo "-alpha")
  MAJOR=$(echo "$VERSION" | cut -d. -f1)
  MINOR=$(echo "$VERSION" | cut -d. -f2)
  PATCH=$(echo "$VERSION" | cut -d. -f3)
fi

case $BUMP_TYPE in
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
    echo "Invalid bump type: $BUMP_TYPE (use point, minor, or major)"
    exit 1
    ;;
esac

NEW_TAG="${MAJOR}.${MINOR}.${PATCH}${SUFFIX}"

# Stage + commit
echo "🔧 Committing changes to $BRANCH..."
git add -A .
git commit -m "$MSG"

# Tag
echo "🏷️  Tagging commit as $NEW_TAG"
git tag -a "$NEW_TAG" -m "$MSG"

# Push commit and tag
echo "📤 Pushing branch '$BRANCH' and tag '$NEW_TAG'..."
git push origin "$BRANCH"
git push origin "$NEW_TAG"

echo "✅ Done: $NEW_TAG pushed with commit to $BRANCH"


