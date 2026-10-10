#!/usr/bin/env bash
# SPDX-License-Identifier: Apache-2.0
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MODULE="$(mktemp -d)"
trap 'rm -rf "$MODULE"' EXIT
# A published module has no parent repository: examples must use its packaged header.
cp -R "$ROOT/go/." "$MODULE/"
cd "$MODULE"
GOWORK=off CGO_ENABLED=1 go test ./...
