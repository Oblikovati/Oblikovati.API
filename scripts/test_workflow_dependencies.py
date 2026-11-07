# SPDX-License-Identifier: Apache-2.0
import unittest
from pathlib import Path
import tempfile
from unittest.mock import patch
from workflow_dependencies import action_errors, step_configuration, workflow_errors, repository_errors, main

PIN = "a" * 40


class FakeWorkflowTree:
    def __init__(self, root: Path, contents: str, name: str = "ci.yml") -> None:
        self.root = root
        (root / "rust-toolchain.toml").write_text('[toolchain]\nchannel = "1.95.0"\n')
        directory = root / ".github/workflows"
        directory.mkdir(parents=True)
        (directory / name).write_text(contents)


class RefusedWorkflowRepository:
    def __call__(self, root: Path) -> list[str]:
        return ["workflow ref 'mutable': expected an immutable commit"]


class WorkflowDependencyTests(unittest.TestCase):
    def test_immutable_and_local_actions_are_accepted(self) -> None:
        self.assertEqual(action_errors(f"actions/checkout@{PIN}", "", "1.95.0"), [])
        self.assertEqual(action_errors("./.github/actions/local", "", "1.95.0"), [])

    def test_mutable_and_abbreviated_action_refs_are_refused(self) -> None:
        for reference in ("actions/checkout@v5", "actions/checkout@abcdef", "actions/checkout"):
            self.assertIn("expected an immutable", action_errors(reference, "", "1.95.0")[0])

    def test_rust_action_requires_the_repository_toolchain(self) -> None:
        reference = f"dtolnay/rust-toolchain@{PIN}"
        self.assertEqual(action_errors(reference, 'with:\n  toolchain: "1.95.0"', "1.95.0"), [])
        for configuration in ("", "with:\n  toolchain: stable", "with:\n  toolchain: 1.94.0"):
            self.assertIn("expected explicit toolchain", action_errors(reference, configuration, "1.95.0")[0])

    def test_later_steps_cannot_supply_the_missing_toolchain(self) -> None:
        lines = [f"      - uses: dtolnay/rust-toolchain@{PIN}", "      - uses: ./other", "        with:", "          toolchain: 1.95.0"]
        self.assertEqual(step_configuration(lines, 0), "")
        self.assertEqual(len(workflow_errors("\n".join(lines), "1.95.0")), 1)

    def test_named_step_keeps_its_own_toolchain_settings(self) -> None:
        lines = ["      - name: Rust", f"        uses: dtolnay/rust-toolchain@{PIN}", "        with:", "          toolchain: 1.95.0", "      - run: cargo test"]
        self.assertEqual(workflow_errors("\n".join(lines), "1.95.0"), [])

    def test_go_tools_require_exact_source_commits(self) -> None:
        self.assertEqual(workflow_errors(f"- run: go install example.org/tool@{PIN}", "1.95.0"), [])
        self.assertEqual(workflow_errors("- run: go install ./cmd/local", "1.95.0"), [])
        for reference in ("example.org/tool@latest", "example.org/tool@v1.1.0", "example.org/tool"):
            self.assertIn("expected an immutable", workflow_errors(f"- run: go install {reference}", "1.95.0")[0])

    def test_commented_examples_are_not_executed_dependencies(self) -> None:
        self.assertEqual(workflow_errors("# - uses: actions/checkout@v5\n# go install example.org/tool@latest", "1.95.0"), [])

    def test_repository_paths_identify_the_rejected_workflow(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            tree = FakeWorkflowTree(Path(directory), "- uses: actions/checkout@v5")
            failures = repository_errors(tree.root)
            self.assertTrue(failures[0].startswith(".github/workflows/ci.yml: action"))

    def test_repository_accepts_a_pinned_workflow(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            tree = FakeWorkflowTree(Path(directory), f"- uses: actions/checkout@{PIN}")
            self.assertEqual(repository_errors(tree.root), [])

    def test_command_refuses_a_reported_mutable_dependency(self) -> None:
        with patch("workflow_dependencies.repository_errors", RefusedWorkflowRepository()):
            with self.assertRaisesRegex(SystemExit, "expected an immutable commit"):
                main()

    def test_yaml_extension_cannot_bypass_the_dependency_guard(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            tree = FakeWorkflowTree(Path(directory), "- uses: actions/checkout@v5", "publish.yaml")
            failures = repository_errors(tree.root)
            self.assertTrue(failures[0].startswith(".github/workflows/publish.yaml: action"))
