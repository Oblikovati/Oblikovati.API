# SPDX-License-Identifier: Apache-2.0
"""Check the API duplication gate's threshold and failure handling."""

import json
import os
import subprocess
import tempfile
import unittest
from pathlib import Path

from duplication_gate import (
    REQUIRED_FORMATS,
    SOURCE_FILES,
    SOURCE_TREES,
    check_report,
    check_scope,
    run_scan,
)

SOURCE_EXTENSIONS = (
    ".go",
    ".rs",
    ".py",
    ".sh",
    ".bash",
    ".ps1",
    ".c",
    ".h",
    ".cc",
    ".cpp",
    ".hpp",
    ".js",
    ".mjs",
    ".ts",
    ".tsx",
    ".json",
    ".jsonc",
    ".toml",
    ".yml",
    ".yaml",
    ".java",
    ".kt",
    ".rb",
    ".swift",
)


def scan_report(lines: int, duplicated: int) -> dict[str, object]:
    """Create a scanner report; example: scan_report(100, 2)."""
    return {
        "statistics": {
            "total": {
                "lines": lines,
                "duplicatedLines": duplicated,
                "sources": len(REQUIRED_FORMATS),
            },
            "formats": {name: {"sources": 1} for name in REQUIRED_FORMATS},
        }
    }


class FakeDuplicationScanner:
    """Write a chosen report or return an error without running a program."""

    def __init__(self, report: object, code: int = 0) -> None:
        self.report, self.code = report, code
        self.command: list[str] = []

    def __call__(
        self, command: list[str], **kwargs: object
    ) -> subprocess.CompletedProcess[str]:
        self.command = command
        report_dir = Path(command[command.index("--output") + 1])
        (report_dir / "jscpd-report.json").write_text(json.dumps(self.report))
        return subprocess.CompletedProcess(
            command, self.code, stdout="", stderr="scanner failed"
        )


class FakeMissingReportScanner:
    """Succeed without producing a report to exercise fail-closed behavior."""

    def __call__(
        self, command: list[str], **kwargs: object
    ) -> subprocess.CompletedProcess[str]:
        return subprocess.CompletedProcess(command, 0, stdout="", stderr="")


class FakeMalformedReportScanner:
    """Produce malformed JSON while reporting scanner success."""

    def __call__(
        self, command: list[str], **kwargs: object
    ) -> subprocess.CompletedProcess[str]:
        report_dir = Path(command[command.index("--output") + 1])
        (report_dir / "jscpd-report.json").write_text("not JSON")
        return subprocess.CompletedProcess(command, 0, stdout="", stderr="")


def create_scan_roots(root: Path) -> None:
    """Create a minimal tree; example: create_scan_roots(Path('/tmp/fixture'))."""
    for name in SOURCE_TREES:
        (root / name).mkdir(parents=True, exist_ok=True)
    for name in SOURCE_FILES:
        (root / name).write_text("fixture", encoding="utf-8")


class DuplicationGateTests(unittest.TestCase):
    def test_unrounded_threshold_rejects_three_percent_and_above(self) -> None:
        self.assertEqual(check_report(scan_report(1000, 29)), (1000, 29))
        self.assertEqual(check_report(scan_report(10000, 299)), (10000, 299))
        for count in (30, 31):
            with self.subTest(duplicated=count), self.assertRaisesRegex(
                ValueError, "below 3%"
            ):
                check_report(scan_report(1000, count))

    def test_invalid_or_empty_reports_are_refused(self) -> None:
        for report in (
            {},
            scan_report(0, 0),
            scan_report(1, 2),
            {"statistics": {"total": {}}},
        ):
            with self.subTest(report=report), self.assertRaises(ValueError):
                check_report(report)

    def test_scanner_failure_cannot_pass_with_a_valid_report(self) -> None:
        scanner = FakeDuplicationScanner(scan_report(100, 1), code=1)
        with tempfile.TemporaryDirectory() as folder:
            create_scan_roots(Path(folder))
            with self.assertRaisesRegex(ValueError, "scanner.*exit 1"):
                run_scan(Path(folder), scanner)

    def test_success_without_report_is_rejected(self) -> None:
        with tempfile.TemporaryDirectory() as folder:
            create_scan_roots(Path(folder))
            with self.assertRaisesRegex(ValueError, "expected valid JSON"):
                run_scan(Path(folder), FakeMissingReportScanner())

    def test_malformed_json_report_is_rejected(self) -> None:
        with tempfile.TemporaryDirectory() as folder:
            create_scan_roots(Path(folder))
            with self.assertRaisesRegex(ValueError, "expected valid JSON"):
                run_scan(Path(folder), FakeMalformedReportScanner())

    @unittest.skipUnless(os.name == "posix", "requires unprivileged symbolic links")
    def test_linked_source_tree_entry_is_refused(self) -> None:
        with tempfile.TemporaryDirectory() as folder:
            root = Path(folder)
            create_scan_roots(root)
            (root / "go/link").symlink_to(root / "rust", target_is_directory=True)
            with self.assertRaisesRegex(ValueError, "no symlinks"):
                check_scope(root)

    @unittest.skipUnless(os.name == "posix", "requires unprivileged symbolic links")
    def test_linked_rust_parent_is_refused(self) -> None:
        with tempfile.TemporaryDirectory() as folder:
            root = Path(folder)
            create_scan_roots(root)
            (root / "rust").rename(root / "actual-rust")
            (root / "rust").symlink_to(root / "actual-rust", target_is_directory=True)
            with self.assertRaisesRegex(ValueError, "expected a real directory"):
                check_scope(root)

    def test_run_scans_all_trees_without_gitignore_or_silent_empty_output(
        self,
    ) -> None:
        scanner = FakeDuplicationScanner(scan_report(100, 1))
        with tempfile.TemporaryDirectory() as folder:
            create_scan_roots(Path(folder))
            self.assertEqual(run_scan(Path(folder), scanner), (100, 1))
        self.assertIn("--no-gitignore", scanner.command)
        self.assertIn("--fail-on-empty", scanner.command)
        paths = list(SOURCE_TREES + SOURCE_FILES)
        self.assertEqual(scanner.command[-len(paths) :], paths)

    def test_missing_source_format_is_rejected(self) -> None:
        report = scan_report(100, 1)
        del report["statistics"]["formats"]["python"]
        with self.assertRaisesRegex(ValueError, "expected statistics"):
            check_report(report)

    def test_tracked_source_files_are_inside_the_scan_scope(self) -> None:
        root = Path(__file__).resolve().parent.parent
        result = subprocess.check_output(["git", "ls-files", "-z"], cwd=root)
        for entry in filter(None, result.split(b"\0")):
            name = os.fsdecode(entry)
            if not name.endswith(SOURCE_EXTENSIONS) and Path(name).name != "Makefile":
                continue
            with self.subTest(source=name):
                inside = name in SOURCE_FILES or any(
                    name.startswith(f"{tree}/") for tree in SOURCE_TREES
                )
                self.assertTrue(
                    inside, f"tracked source {name!r} is outside duplication scope"
                )


if __name__ == "__main__":
    unittest.main()
