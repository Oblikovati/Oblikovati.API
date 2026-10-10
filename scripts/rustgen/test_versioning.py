# SPDX-License-Identifier: Apache-2.0
"""Version changes must reach both the Cargo package and the C ABI handshake."""

from pathlib import Path
from tempfile import TemporaryDirectory
import unittest
from versioning import synchronize_version


class VersionWorkspace:
    def __init__(self, directory: str) -> None:
        self.root = Path(directory)
        (self.root / "go").mkdir()
        (self.root / "rust/src").mkdir(parents=True)
        (self.root / "rust/Cargo.toml").write_text(
            '[package]\nname = "oblikovati-api"\nversion = "0.0.0"\n'
        )

    def version(self, value: str) -> None:
        (self.root / "go/version.go").write_text(
            f'package api\nconst Version = "{value}"\n'
        )


class VersionSynchronizationTests(unittest.TestCase):
    def test_every_release_component_reaches_rust(self) -> None:
        with TemporaryDirectory() as directory:
            workspace = VersionWorkspace(directory)
            for version in ["0.155.0", "0.155.1", "1.0.0", "2.7.13"]:
                workspace.version(version)
                synchronize_version(workspace.root)
                self.assertIn(
                    f'version = "{version}"',
                    (workspace.root / "rust/Cargo.toml").read_text(),
                )
                constants = (workspace.root / "rust/src/version.rs").read_text()
                major, minor, _ = version.split(".")
                self.assertIn(f'VERSION: &str = "{version}"', constants)
                self.assertIn(f"API_MAJOR: i32 = {major};", constants)
                self.assertIn(f"API_MINOR: i32 = {minor};", constants)

    def test_invalid_version_fails_before_writing_rust(self) -> None:
        with TemporaryDirectory() as directory:
            workspace = VersionWorkspace(directory)
            workspace.version("invalid")
            with self.assertRaisesRegex(ValueError, "invalid.*expected const Version"):
                synchronize_version(workspace.root)
            self.assertIn(
                'version = "0.0.0"', (workspace.root / "rust/Cargo.toml").read_text()
            )


if __name__ == "__main__":
    unittest.main()
