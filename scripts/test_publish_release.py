# SPDX-License-Identifier: Apache-2.0
"""Check release publication recovery without replacing tags or assets."""

import tempfile
import unittest
import subprocess
from pathlib import Path
from unittest.mock import patch

from promote_release import GitRepository, promote_release
from publish_release import GitHubReleases, publish_release
from test_promote_release import ReleaseGitFixture


class FakeReleaseService:
    """Simulate a release created before its asset upload failed."""

    def __init__(self, notes: str, content: bytes) -> None:
        self.release: dict[str, object] | None = None
        self.notes, self.content = notes, content
        self.created = self.uploaded = 0

    def view(self, tag: str) -> dict[str, object] | None:
        return self.release

    def create(self, tag: str, artifact: Path, notes: Path) -> bool:
        self.created += 1
        self.release = {
            "tag_name": tag, "name": tag, "draft": False, "prerelease": False,
            "author": {"login": "github-actions[bot]"}, "body": self.notes,
            "assets": [],
        }
        return False

    def upload(self, tag: str, artifact: Path) -> bool:
        self.uploaded += 1
        assert self.release is not None
        self.release["assets"] = [{"name": artifact.name, "state": "uploaded"}]
        return True

    def download(self, tag: str, name: str) -> bytes:
        return self.content


class FakePublicationClock:
    """Avoid waiting during publication retry tests."""

    def __call__(self, seconds: float) -> None:
        assert seconds == 5


class FakeTimedReleaseCommand:
    """Record publication request timeouts and simulate a downloaded file."""

    def __init__(self, expired: bool = False) -> None:
        self.expired = expired
        self.timeouts: list[object] = []

    def __call__(self, command: list[str], **kwargs: object) -> subprocess.CompletedProcess[str]:
        self.timeouts.append(kwargs.get("timeout"))
        if self.expired:
            raise subprocess.TimeoutExpired(command, 120)
        if "download" in command:
            Path(command[command.index("--dir") + 1], "package.crate").write_bytes(b"crate")
        return subprocess.CompletedProcess(command, 0, stdout="{}", stderr="")


class ReleasePublicationTests(unittest.TestCase):
    def test_github_requests_are_bounded_and_expiry_is_safe(self) -> None:
        for expired in (False, True):
            with self.subTest(expired=expired):
                process = FakeTimedReleaseCommand(expired)
                releases = GitHubReleases("owner/repo")
                with patch("publish_release.subprocess.run", process):
                    self.assertEqual(releases.view("v0.1.0"), None if expired else {})
                    self.assertEqual(releases.create("v0.1.0", Path("package.crate"), Path("notes")), not expired)
                    self.assertEqual(releases.upload("v0.1.0", Path("package.crate")), not expired)
                    if expired:
                        with self.assertRaises(subprocess.TimeoutExpired):
                            releases.download("v0.1.0", "package.crate")
                    else:
                        self.assertEqual(releases.download("v0.1.0", "package.crate"), b"crate")
                self.assertEqual(process.timeouts, [120] * 4)

    def test_failed_create_recovers_the_missing_asset_without_replacing_it(self) -> None:
        with tempfile.TemporaryDirectory() as folder:
            root = Path(folder)
            fixture = ReleaseGitFixture(root)
            promote_release(GitRepository(fixture.first), fixture.base, fixture.candidate, "v0.2.0")
            artifact, notes = root / "oblikovati-api-0.2.0.crate", root / "notes.md"
            artifact.write_bytes(b"verified crate")
            notes.write_text("release notes\n")
            service = FakeReleaseService("release notes", b"verified crate")
            publish_release(service, GitRepository(fixture.first), fixture.candidate,
                            "v0.2.0", artifact, notes, FakePublicationClock())
            self.assertEqual((service.created, service.uploaded), (1, 1))

    def test_existing_wrong_asset_or_notes_are_not_overwritten(self) -> None:
        for bad_notes in (False, True):
            with self.subTest(bad_notes=bad_notes), tempfile.TemporaryDirectory() as folder:
                root = Path(folder)
                fixture = ReleaseGitFixture(root)
                promote_release(GitRepository(fixture.first), fixture.base, fixture.candidate, "v0.2.0")
                artifact, notes = root / "oblikovati-api-0.2.0.crate", root / "notes.md"
                artifact.write_bytes(b"verified crate")
                notes.write_text("release notes")
                service = FakeReleaseService("other notes" if bad_notes else "release notes", b"wrong")
                service.create("v0.2.0", artifact, notes)
                service.upload("v0.2.0", artifact)
                with self.assertRaisesRegex(ValueError, "expected"):
                    publish_release(service, GitRepository(fixture.first), fixture.candidate,
                                    "v0.2.0", artifact, notes, FakePublicationClock())
                self.assertEqual(service.uploaded, 1)

    def test_missing_promoted_tags_prevent_release_creation(self) -> None:
        with tempfile.TemporaryDirectory() as folder:
            root = Path(folder)
            fixture = ReleaseGitFixture(root)
            artifact, notes = root / "oblikovati-api-0.2.0.crate", root / "notes.md"
            artifact.write_bytes(b"verified crate")
            notes.write_text("release notes")
            service = FakeReleaseService("release notes", b"verified crate")
            with self.assertRaisesRegex(ValueError, "remote tags"):
                publish_release(service, GitRepository(fixture.first), fixture.candidate,
                                "v0.2.0", artifact, notes, FakePublicationClock())
            self.assertEqual(service.created, 0)


if __name__ == "__main__":
    unittest.main()
