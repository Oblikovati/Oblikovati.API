# SPDX-License-Identifier: Apache-2.0
"""Keep workflow executables pinned without changing the requested tool versions."""
from pathlib import Path
import re
import tomllib

USES = re.compile(r"^\s*(?:-\s+)?uses:\s*(\S+)")
INSTALL = re.compile(r"\bgo\s+install\s+(\S+)")
IMMUTABLE = re.compile(r"[^@\s]+@[0-9a-f]{40}")


def step_configuration(lines: list[str], position: int) -> str:
    """Read only the current action's settings; example: step_configuration(lines, 2)."""
    indentation = len(lines[position]) - len(lines[position].lstrip())
    boundary = indentation if lines[position].lstrip().startswith("-") else indentation - 2
    following: list[str] = []
    for line in lines[position + 1:]:
        if line.strip() and len(line) - len(line.lstrip()) <= boundary:
            break
        following.append(line)
    return "\n".join(following)


def action_errors(reference: str, configuration: str, toolchain: str) -> list[str]:
    """Reject mutable actions or implicit Rust versions; example: action_errors(ref, block, version)."""
    if reference.startswith("./"):
        return []
    if not IMMUTABLE.fullmatch(reference):
        return [f"action {reference!r}: expected an immutable 40-digit commit SHA"]
    if not reference.startswith("dtolnay/rust-toolchain@"):
        return []
    version = re.search(r"(?m)^\s*toolchain:\s*[\"']?([^\s\"']+)[\"']?\s*$", configuration)
    if version is None or version.group(1) != toolchain:
        return [f"action {reference!r}: expected explicit toolchain {toolchain!r}"]
    return []


def go_tool_errors(line: str) -> list[str]:
    """Reject mutable Go tool source refs; example: go_tool_errors('go install tool@latest')."""
    return [f"Go tool {installed!r}: expected an immutable 40-digit source commit"
            for installed in INSTALL.findall(line)
            if not installed.startswith("./") and not IMMUTABLE.fullmatch(installed)]


def workflow_errors(contents: str, toolchain: str) -> list[str]:
    """Inspect workflow dependency refs; example: workflow_errors(text, '1.95.0')."""
    lines = contents.splitlines()
    failures: list[str] = []
    for index, line in enumerate(lines):
        if line.lstrip().startswith("#"):
            continue
        action = USES.match(line)
        if action:
            failures.extend(action_errors(action.group(1), step_configuration(lines, index), toolchain))
        failures.extend(go_tool_errors(line))
    return failures


def repository_errors(root: Path) -> list[str]:
    """Check every workflow against the pinned Rust toolchain; example: repository_errors(root)."""
    toolchain = tomllib.loads((root / "rust-toolchain.toml").read_text())["toolchain"]["channel"]
    failures: list[str] = []
    directory = root / ".github/workflows"
    paths = sorted(path for path in directory.iterdir() if path.suffix in (".yml", ".yaml"))
    for path in paths:
        failures.extend(f"{path.relative_to(root)}: {error}" for error in workflow_errors(path.read_text(), toolchain))
    return failures


def main() -> None:
    """Fail the local and hosted gate on mutable refs; example: python3 scripts/workflow_dependencies.py."""
    failures = repository_errors(Path(__file__).resolve().parent.parent)
    if failures:
        raise SystemExit("\n".join(failures))
    print("Workflow actions and Go tools use immutable source revisions")


if __name__ == "__main__":
    main()
