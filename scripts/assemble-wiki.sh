#!/usr/bin/env bash
# SPDX-License-Identifier: Apache-2.0
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"
mkdir -p build/wiki
find build/wiki -maxdepth 1 -type f -name '*.md' -delete
python3 scripts/wiki-examples.py --check
for page in docs/wiki/*.md; do
  [[ "$(basename "$page")" == README.md ]] && continue
  cp "$page" build/wiki/
done
(cd go && GOWORK=off go run ./scripts/apischema) > build/api-schema.json
bash scripts/gen-api-docs.sh build/wiki/Go-API-Reference.md
python3 scripts/gen-rust-docs.py
cat > build/wiki/API-Docs.md <<'DOCS'
# API Docs

The Go and Rust references describe the same version of the public API.

- [[Go API Reference]]: types, interfaces, wire DTOs and typed Go client.
- [[Rust API Reference]]: types, traits, wire DTOs and typed Rust client.
- [[First Steps]]: complete examples in both languages.
- [[C ABI Add-ins]]: the shared C header and callback ownership rules.
- [[Typed Operations]]: matching requests, feature arguments and events.

Both references are generated from the same checked API source. Run `make gate`
before publishing changes.
DOCS
