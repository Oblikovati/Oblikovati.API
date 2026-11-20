# SPDX-License-Identifier: Apache-2.0
"""Exercise release promotion against a local Git remote."""

import subprocess
import tempfile
import unittest
from pathlib import Path

from promote_release import GitRepository, promote_release, validate_release_identity


class ReleaseGitFixture:
    """A bare remote with a tested candidate and a second writer."""

    def __init__(self, root: Path) -> None:
        self.remote, self.first, self.other = root / "remote.git", root / "first", root / "other"
        subprocess.run(["git", "init", "--bare", "-q", str(self.remote)], check=True)
        subprocess.run(["git", "init", "-q", "-b", "develop", str(self.first)], check=True)
        self.configure(self.first)
        self.git(self.first, "remote", "add", "origin", str(self.remote))
        (self.first / "file").write_text("base")
        self.commit(self.first, "base")
        self.git(self.first, "push", "-q", "origin", "HEAD:develop")
        self.base = self.git(self.first, "rev-parse", "HEAD")
        subprocess.run(["git", "clone", "-q", "-b", "develop", str(self.remote), str(self.other)], check=True)
        self.configure(self.other)
        (self.first / "file").write_text("candidate")
        self.commit(self.first, "candidate")
        self.candidate = self.git(self.first, "rev-parse", "HEAD")

    def git(self, root: Path, *args: str) -> str:
        """Read a Git result; example: fixture.git(root, 'rev-parse', 'HEAD')."""
        return subprocess.check_output(["git", "-C", str(root), *args], text=True).strip()

    def configure(self, root: Path) -> None:
        """Give a checkout a test identity; example: fixture.configure(root)."""
        self.git(root, "config", "user.name", "Fixture")
        self.git(root, "config", "user.email", "fixture@example.invalid")

    def commit(self, root: Path, message: str) -> None:
        """Commit one synthetic file; example: fixture.commit(root, 'change')."""
        self.git(root, "add", "file")
        self.git(root, "commit", "-qm", message)

    def advance_remote(self) -> str:
        """Advance develop independently; example: fixture.advance_remote()."""
        (self.other / "file").write_text("different")
        self.commit(self.other, "concurrent change")
        self.git(self.other, "push", "-q", "origin", "HEAD:develop")
        return self.git(self.other, "rev-parse", "HEAD")


class RacingReleaseGit(GitRepository):
    """Advance the remote after the precheck, just before the atomic push."""

    def __init__(self, fixture: ReleaseGitFixture) -> None:
        super().__init__(fixture.first)
        self.fixture = fixture

    def run(self, *args: str) -> None:
        if args[0] == "push":
            self.fixture.advance_remote()
        super().run(*args)


class ReleasePromotionTests(unittest.TestCase):
    def test_promotion_pushes_annotated_tags_on_the_tested_commit(self) -> None:
        with tempfile.TemporaryDirectory() as folder:
            fixture = ReleaseGitFixture(Path(folder))
            promote_release(GitRepository(fixture.first), fixture.base, fixture.candidate, "v0.2.0")
            for tag in ("v0.2.0", "go/v0.2.0"):
                self.assertEqual(fixture.git(fixture.first, "cat-file", "-t", tag), "tag")
                self.assertEqual(fixture.git(fixture.first, "rev-parse", f"{tag}^{{commit}}"), fixture.candidate)
            remote = fixture.git(fixture.first, "ls-remote", "--heads", "origin", "develop")
            self.assertTrue(remote.startswith(fixture.candidate))

    def test_stale_develop_is_rejected_before_tags(self) -> None:
        with tempfile.TemporaryDirectory() as folder:
            fixture = ReleaseGitFixture(Path(folder))
            fixture.advance_remote()
            with self.assertRaisesRegex(ValueError, "remote develop"):
                promote_release(GitRepository(fixture.first), fixture.base, fixture.candidate, "v0.2.0")
            self.assertEqual(fixture.git(fixture.first, "ls-remote", "--tags", "origin"), "")

    def test_remote_race_cannot_publish_the_commit_or_either_tag(self) -> None:
        with tempfile.TemporaryDirectory() as folder:
            fixture = ReleaseGitFixture(Path(folder))
            with self.assertRaises(subprocess.CalledProcessError):
                promote_release(RacingReleaseGit(fixture), fixture.base, fixture.candidate, "v0.2.0")
            self.assertEqual(fixture.git(fixture.first, "ls-remote", "--tags", "origin"), "")

    def test_invalid_tag_and_commit_are_rejected(self) -> None:
        for base, candidate, tag in [("bad", "b" * 40, "v0.1.0"), ("a" * 40, "b" * 40, "v1;echo")]:
            with self.assertRaisesRegex(ValueError, "expected"):
                validate_release_identity(base, candidate, tag)


if __name__ == "__main__":
    unittest.main()
