# SPDX-License-Identifier: Apache-2.0
"""Check the API duplication gate's threshold and failure handling."""

import json
import os
import subprocess
import tempfile
import unittest
from pathlib import Path

from duplication_gate import check_report, check_scope, run_scan


def scan_report(lines: int, duplicated: int) -> dict[str, object]:
    """Create a scanner report; example: scan_report(100, 2)."""
    return {
        "statistics": {
            "total": {"lines": lines, "duplicatedLines": duplicated, "sources": 2},
            "formats": {"go": {"sources": 1}, "rust": {"sources": 1}},
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
    (root / "go").mkdir()
    (root / "rust/src").mkdir(parents=True)


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

    def test_run_scans_both_trees_without_gitignore_or_silent_empty_output(
        self,
    ) -> None:
        scanner = FakeDuplicationScanner(scan_report(100, 1))
        with tempfile.TemporaryDirectory() as folder:
            create_scan_roots(Path(folder))
            self.assertEqual(run_scan(Path(folder), scanner), (100, 1))
        self.assertIn("--no-gitignore", scanner.command)
        self.assertIn("--fail-on-empty", scanner.command)
        self.assertEqual(scanner.command[-2:], ["go", "rust/src"])


if __name__ == "__main__":
    unittest.main()
