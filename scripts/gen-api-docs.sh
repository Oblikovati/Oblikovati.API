#!/usr/bin/env bash
# SPDX-License-Identifier: Apache-2.0
#
# gen-api-docs.sh — render the public API contract's Go doc comments into a single
# Markdown page (the wiki's "API Docs" page) using gomarkdoc.
#
# The output covers ONLY the Apache-2.0 contract surface and, as a hard gate, must not
# name any third-party CAD product (a copyright constraint for this public repo).
#
# Usage: scripts/gen-api-docs.sh [output-path]   (default: build/wiki/API-Docs.md)
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

OUT="${1:-build/wiki/API-Docs.md}"
mkdir -p "$(dirname "$OUT")"

# Packages in logical reading order: vocabulary → interfaces → wire DTOs → client.
PKGS=(
  oblikovati.org/api/types
  oblikovati.org/api/contract
  oblikovati.org/api/wire
  oblikovati.org/api/client
)

GOMARKDOC="${GOMARKDOC:-gomarkdoc}"
if ! command -v "$GOMARKDOC" >/dev/null 2>&1; then
  echo "gen-api-docs: '$GOMARKDOC' not found on PATH. Install it with:" >&2
  echo "  go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest" >&2
  exit 1
fi

TMP="$(mktemp)"
trap 'rm -f "$TMP"' EXIT
"$GOMARKDOC" --output "$TMP" "${PKGS[@]}"

# Copyright gate: the public docs must never name a third-party CAD product.
if grep -niE 'inventor|autodesk' "$TMP"; then
  echo "gen-api-docs: generated docs name a forbidden third-party product (lines above)." >&2
  echo "gen-api-docs: scrub the offending doc comment in the source before regenerating." >&2
  exit 1
fi

{
  echo "# API Docs"
  echo
  echo "_Auto-generated from the \`oblikovati.org/api\` Go source and doc comments on every merge"
  echo "to \`develop\`. Do not edit this page by hand._"
  echo
  cat "$TMP"
} > "$OUT"

echo "gen-api-docs: wrote $OUT ($(wc -l < "$OUT") lines)"
