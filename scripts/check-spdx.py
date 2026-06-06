#!/usr/bin/env python3
"""Verify (or insert) the Apache-2.0 SPDX header on every Go file in this module.

This is the standalone public API contract module (oblikovati/api),
Apache-2.0 throughout — see ADR-0018 in the Oblikovati app repo. Unlike the
monorepo predecessor, there is a single license tree rooted at the repo root, so
the mapping is just "every *.go -> Apache-2.0".

Placement rules (so Go semantics are preserved):
  * The SPDX comment is its own block followed by a blank line, so it never merges
    into a following `// Package ...` doc comment.
  * In files beginning with a build constraint (`//go:build` / `// +build`), the
    header goes AFTER the constraint block (the constraint must stay first).
  * Files that already carry an SPDX-License-Identifier are left untouched.

Usage: python3 scripts/check-spdx.py [--check]
  --check exits non-zero if any file would change (for CI), without writing.
"""
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
IDENTIFIER = "Apache-2.0"
# Directories that are not part of the module's licensed Go surface.
SKIP_DIRS = {".git", "scripts"}


def header(identifier: str) -> str:
    return f"// SPDX-License-Identifier: {identifier}\n"


def is_constraint(line: str) -> bool:
    s = line.lstrip()
    return s.startswith("//go:build") or s.startswith("// +build")


def insert_index(lines: list[str]) -> int:
    if lines and is_constraint(lines[0]):
        i = 0
        while i < len(lines) and (is_constraint(lines[i]) or lines[i].strip() == ""):
            i += 1
        return i
    return 0


def patched(text: str, identifier: str) -> str | None:
    if "SPDX-License-Identifier" in text:
        return None
    lines = text.splitlines(keepends=True)
    at = insert_index(lines)
    return "".join(lines[:at] + [header(identifier), "\n"] + lines[at:])


def go_files() -> list[Path]:
    return [
        p
        for p in sorted(ROOT.rglob("*.go"))
        if not any(part in SKIP_DIRS for part in p.relative_to(ROOT).parts)
    ]


def main() -> int:
    check = "--check" in sys.argv[1:]
    changed = []
    for path in go_files():
        out = patched(path.read_text(), IDENTIFIER)
        if out is None:
            continue
        changed.append(path.relative_to(ROOT))
        if not check:
            path.write_text(out)
    if check and changed:
        print("missing SPDX header:")
        for p in changed:
            print(f"  {p}")
        return 1
    if not check:
        print(f"added SPDX headers to {len(changed)} files")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
