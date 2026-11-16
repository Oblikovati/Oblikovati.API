# SPDX-License-Identifier: Apache-2.0
"""Enforce the API's duplication limit over all Go and Rust source trees."""

import json
import os
import stat
import subprocess
import sys
import tempfile
from pathlib import Path
from typing import Callable, NoReturn

MAX_BYTES = 1_000_000_000
Scan = Callable[..., subprocess.CompletedProcess[str]]
SCAN_FLAGS = (
    "--no-gitignore",
    "--max-size",
    str(MAX_BYTES),
    "--fail-on-empty",
    "--min-lines",
    "5",
    "--min-tokens",
    "50",
    "--mode",
    "mild",
    "--reporters",
    "json",
)


def scan_command(config: Path, output: Path) -> list[str]:
    """Fix the scanner scope and rules; example: scan_command(config, output)."""
    return [
        "npx",
        "--yes",
        "jscpd@5.4.1",
        "--config",
        str(config),
        *SCAN_FLAGS,
        "--output",
        str(output),
        "go",
        "rust/src",
    ]


def raise_scan_error(error: OSError) -> NoReturn:
    """Refuse unreadable source trees; example: raise_scan_error(PermissionError('go'))."""
    raise error


def check_scope(root: Path) -> None:
    """Refuse inputs the scanner would silently skip; example: check_scope(repo)."""
    for name in ("go", "rust/src"):
        directory = root / name
        if directory.is_symlink() or not directory.is_dir():
            raise ValueError(f"scan directory {directory}: expected a real directory")
        for parent, dirs, files in os.walk(
            directory, followlinks=False, onerror=raise_scan_error
        ):
            check_entries(Path(parent), dirs, files)


def check_entries(parent: Path, directories: list[str], files: list[str]) -> None:
    """Reject links, special files and oversized inputs; example: check_entries(root, dirs, files)."""
    for name in directories + files:
        path = parent / name
        info = path.lstat()
        if stat.S_ISDIR(info.st_mode) and name in directories:
            continue
        if not stat.S_ISREG(info.st_mode) or info.st_size > MAX_BYTES:
            raise ValueError(
                f"scan input {path}: expected regular files below {MAX_BYTES} bytes and no symlinks"
            )


def report_counts(report: object) -> tuple[int, int]:
    """Read unrounded scanner counts; example: report_counts(report)."""
    try:
        statistics = report["statistics"]
        total, formats = statistics["total"], statistics["formats"]
        counts = (total["lines"], total["duplicatedLines"], total["sources"])
        counts += (formats["go"]["sources"], formats["rust"]["sources"])
        return require_counts(counts)
    except (KeyError, TypeError) as error:
        raise ValueError(
            f"scanner report {report!r}: expected statistics for Go and Rust"
        ) from error


def require_counts(counts: tuple[int, ...]) -> tuple[int, int]:
    """Require a nonempty valid scan; example: require_counts((100, 1, 2, 1, 1))."""
    if any(type(value) is not int for value in counts):
        raise ValueError(f"scanner counts {counts!r}: expected integers")
    lines, duplicated = counts[:2]
    if lines <= 0 or duplicated < 0 or duplicated > lines or min(counts[2:]) <= 0:
        raise ValueError(
            f"scanner counts {counts!r}: expected nonempty Go and Rust scans"
        )
    return lines, duplicated


def check_report(report: object) -> tuple[int, int]:
    """Enforce strictly less than three percent; example: check_report(report)."""
    lines, duplicated = report_counts(report)
    if duplicated * 100 >= lines * 3:
        raise ValueError(f"duplication {duplicated}/{lines}: expected below 3%")
    return lines, duplicated


def run_scan(root: Path, run: Scan = subprocess.run) -> tuple[int, int]:
    """Run one isolated scanner and validate its report; example: run_scan(repo)."""
    check_scope(root)
    with tempfile.TemporaryDirectory(prefix="oblikovati-duplication-") as folder:
        output = Path(folder)
        config = output / "config.json"
        config.write_text("{}\n", encoding="utf-8")
        command = scan_command(config, output)
        result = run(command, cwd=root, capture_output=True, text=True, check=False)
        if result.returncode:
            raise ValueError(
                f"scanner {command[2]}: expected exit 0, got exit {result.returncode}: {result.stderr}"
            )
        return check_report(read_report(output))


def read_report(output: Path) -> object:
    """Read the new scanner report; example: read_report(output)."""
    try:
        return json.loads((output / "jscpd-report.json").read_text(encoding="utf-8"))
    except (OSError, ValueError) as error:
        raise ValueError(
            f"scanner report in {output}: expected valid JSON: {error}"
        ) from error


def main() -> int:
    """Check the repository tree; run `python3 scripts/duplication_gate.py`."""
    root = Path(__file__).resolve().parent.parent
    try:
        lines, duplicated = run_scan(root)
    except (OSError, ValueError) as error:
        print(f"duplication gate: {error}", file=sys.stderr)
        return 1
    print(
        f"duplication: {duplicated}/{lines} lines ({duplicated * 100 / lines:.4f}%), below 3%"
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
