# SPDX-License-Identifier: Apache-2.0
"""Exercise create-only release candidate branch ownership."""

import subprocess
import tempfile
import unittest
from pathlib import Path

from promote_release import GitRepository
from push_release_candidate import push_candidate
from test_promote_release import ReleaseGitFixture

BRANCH = "api-release-candidate/17-1"


class RacingCandidateGit(GitRepository):
    """Create the candidate ref from another push just before this push."""

    def __init__(self, fixture: ReleaseGitFixture, competing_sha: str) -> None:
        super().__init__(fixture.first)
        self.fixture, self.competing_sha = fixture, competing_sha

    def output(self, *args: str) -> str:
        if args[0] == "push":
            self.fixture.git(
                self.fixture.first, "push", "-q", "origin",
                f"{self.competing_sha}:refs/heads/{BRANCH}",
            )
        return super().output(*args)


class CandidatePushTests(unittest.TestCase):
    def test_creates_the_candidate_ref_at_the_expected_commit(self) -> None:
        with tempfile.TemporaryDirectory() as folder:
            fixture = ReleaseGitFixture(Path(folder))
            push_candidate(GitRepository(fixture.first), fixture.base, fixture.candidate, BRANCH)
            remote = fixture.git(fixture.first, "ls-remote", "--heads", "origin", BRANCH)
            self.assertEqual(remote, f"{fixture.candidate}\trefs/heads/{BRANCH}")

    def test_ref_created_after_precheck_is_not_overwritten_or_claimed(self) -> None:
        for same_sha in (False, True):
            with self.subTest(same_sha=same_sha), tempfile.TemporaryDirectory() as folder:
                fixture = ReleaseGitFixture(Path(folder))
                competing = fixture.candidate if same_sha else fixture.base
                with self.assertRaises((subprocess.CalledProcessError, ValueError)):
                    push_candidate(RacingCandidateGit(fixture, competing), fixture.base, fixture.candidate, BRANCH)
                remote = fixture.git(fixture.first, "ls-remote", "--heads", "origin", BRANCH)
                self.assertEqual(remote, f"{competing}\trefs/heads/{BRANCH}")

    def test_existing_ref_is_never_claimed(self) -> None:
        with tempfile.TemporaryDirectory() as folder:
            fixture = ReleaseGitFixture(Path(folder))
            fixture.git(fixture.first, "push", "-q", "origin", f"HEAD:refs/heads/{BRANCH}")
            with self.assertRaisesRegex(ValueError, "expected an absent"):
                push_candidate(GitRepository(fixture.first), fixture.base, fixture.candidate, BRANCH)


if __name__ == "__main__":
    unittest.main()
