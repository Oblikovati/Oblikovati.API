# SPDX-License-Identifier: Apache-2.0
"""Refresh wiki code blocks from examples compiled by make gate."""

from pathlib import Path
import re
import sys

EXAMPLE = re.compile(r"<!-- example: ([\w/.-]+) -->.*?<!-- /example -->", re.DOTALL)


def render_example(match: re.Match[str]) -> str:
    path = Path(match[1])
    if path.parts[0] not in ("go", "rust") or ".." in path.parts:
        raise ValueError(f"{path}: expected an example inside go/ or rust/")
    language = "go" if path.suffix == ".go" else "rust"
    return f"<!-- example: {path} -->\n```{language}\n{path.read_text().rstrip()}\n```\n<!-- /example -->"


def main() -> None:
    changed: list[str] = []
    for page in sorted(Path("docs/wiki").glob("*.md")):
        original = page.read_text()
        updated = EXAMPLE.sub(render_example, original)
        if original == updated:
            continue
        changed.append(str(page))
        if "--check" not in sys.argv:
            page.write_text(updated)
    if changed and "--check" in sys.argv:
        raise SystemExit(f"Wiki examples are stale in {changed}; run make generate")


if __name__ == "__main__":
    main()
