#!/usr/bin/env bash
#
# Replaces the SDK version string in common/version.go and README.md.
# Used by semantic-release (@semantic-release/exec prepareCmd). It only edits
# files; it never commits. @semantic-release/git makes the single release commit.
#
# Usage: scripts/bump-version.sh [--dry-run] <current-version> <new-version>
#   --dry-run  validate and report what would change, without writing any file.

set -euo pipefail

DRY_RUN=false
if [ "${1:-}" = "--dry-run" ]; then
  DRY_RUN=true
  shift
fi

if [ "$#" -ne 2 ]; then
  echo "Usage: $0 [--dry-run] <current-version> <new-version>" >&2
  exit 2
fi

CURRENT="$1"
NEW="$2"
SEMVER='^[0-9]+\.[0-9]+\.[0-9]+([-+][0-9A-Za-z.-]+)?$'

for v in "$CURRENT" "$NEW"; do
  if ! [[ "$v" =~ $SEMVER ]]; then
    echo "ERROR: '$v' is not a valid semantic version" >&2
    exit 2
  fi
done

if [ "$CURRENT" = "$NEW" ]; then
  echo "ERROR: current and new version are identical ($CURRENT)" >&2
  exit 2
fi

cd "$(dirname "$0")/.."

# file:expected number of occurrences of the current version
TARGETS=("common/version.go:1" "README.md:3")

ESCAPED_CURRENT="${CURRENT//./\\.}"

# Pre-check: every target must contain the current version the expected number of times.
for target in "${TARGETS[@]}"; do
  file="${target%%:*}"
  expected="${target##*:}"
  actual=$(grep -c -F -- "$CURRENT" "$file" || true)
  if [ "$actual" -ne "$expected" ]; then
    echo "ERROR: $file has $actual line(s) containing $CURRENT, expected $expected" >&2
    exit 1
  fi
done

if [ "$DRY_RUN" = true ]; then
  for target in "${TARGETS[@]}"; do
    echo "[dry-run] ${target%%:*}: $CURRENT -> $NEW (${target##*:} line(s))"
  done
  echo "[dry-run] OK, no files were modified"
  exit 0
fi

for target in "${TARGETS[@]}"; do
  file="${target%%:*}"
  tmp=$(mktemp)
  sed "s/${ESCAPED_CURRENT}/${NEW}/g" "$file" > "$tmp"
  cat "$tmp" > "$file"
  rm -f "$tmp"
done

# Post-check: old version gone, new version present.
for target in "${TARGETS[@]}"; do
  file="${target%%:*}"
  expected="${target##*:}"
  if grep -q -F -- "$CURRENT" "$file"; then
    echo "ERROR: $file still contains $CURRENT after bump" >&2
    exit 1
  fi
  actual=$(grep -c -F -- "$NEW" "$file" || true)
  if [ "$actual" -ne "$expected" ]; then
    echo "ERROR: $file has $actual line(s) containing $NEW, expected $expected" >&2
    exit 1
  fi
  echo "Updated $file: $CURRENT -> $NEW"
done
