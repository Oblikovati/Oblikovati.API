# SPDX-License-Identifier: Apache-2.0
"""Test release CI identity and complete-matrix checks."""

import unittest
import subprocess
from unittest.mock import patch

from release_ci import Candidate, GhAPI, JOBS, WORKFLOW, await_ci, require_jobs, require_run

CANDIDATE = Candidate("Oblikovati/Oblikovati.API", "api-release-candidate/17-1", "a" * 40)


def successful_run() -> dict[str, object]:
    """Describe the dispatched CI run; example: successful_run()."""
    return {
        "id": 25, "workflow_id": 12, "path": WORKFLOW, "event": "workflow_dispatch",
        "head_sha": CANDIDATE.sha, "head_branch": CANDIDATE.branch,
        "run_attempt": 1, "status": "completed", "conclusion": "success",
    }


def successful_jobs() -> dict[str, object]:
    """Describe the entire CI matrix; example: successful_jobs()."""
    return {"total_count": 3, "jobs": [
        {"name": name, "status": "completed", "conclusion": "success"}
        for name in sorted(JOBS)
    ]}


class FakeReleaseGitHub:
    """A named fake for dispatch and the returned run."""

    def __init__(self) -> None:
        self.run = successful_run()
        self.jobs = successful_jobs()
        self.posts: list[tuple[str, dict[str, str]]] = []

    def get(self, endpoint: str) -> dict[str, object]:
        if endpoint.endswith("/workflows/ci.yml"):
            return {"id": 12, "path": WORKFLOW, "state": "active"}
        if endpoint.endswith("/attempts/1/jobs?per_page=100"):
            return self.jobs
        if endpoint.endswith("/runs/25"):
            return self.run
        raise AssertionError(f"unexpected endpoint {endpoint!r}")

    def post(self, endpoint: str, payload: dict[str, str]) -> dict[str, object]:
        self.posts.append((endpoint, payload))
        return {"workflow_run_id": 25}


class FakeDispatchWithoutRun(FakeReleaseGitHub):
    """Represent an accepted dispatch that returned no unique run identity."""

    def post(self, endpoint: str, payload: dict[str, str]) -> dict[str, object]:
        self.posts.append((endpoint, payload))
        return {}


class FakeInactiveWorkflow(FakeReleaseGitHub):
    """Represent an unavailable CI workflow."""

    def get(self, endpoint: str) -> dict[str, object]:
        if endpoint.endswith("/workflows/ci.yml"):
            return {"id": 12, "path": WORKFLOW, "state": "disabled_manually"}
        return super().get(endpoint)


class FakeReleaseClock:
    """Record waits without delaying a test."""

    def __init__(self) -> None:
        self.waits = 0
        self.elapsed = 0.0

    def __call__(self, seconds: float) -> None:
        assert seconds == 15
        self.waits += 1
        self.elapsed += seconds

    def now(self) -> float:
        """Return the simulated time; example: clock.now()."""
        return self.elapsed


class FakeTimedGhProcess:
    """Record the GH subprocess timeout and optionally simulate expiry."""

    def __init__(self, expired: bool = False) -> None:
        self.expired = expired
        self.timeout: object = None

    def __call__(self, command: list[str], **kwargs: object) -> subprocess.CompletedProcess[str]:
        self.timeout = kwargs.get("timeout")
        if self.expired:
            raise subprocess.TimeoutExpired(command, self.timeout)
        return subprocess.CompletedProcess(command, 0, stdout="{}", stderr="")


class ReleaseCITests(unittest.TestCase):
    def test_dispatch_uses_returned_run_and_exact_candidate(self) -> None:
        api = FakeReleaseGitHub()
        self.assertEqual(await_ci(api, CANDIDATE), 25)
        self.assertEqual(api.posts[0][1], {"ref": CANDIDATE.branch})

    def test_prior_or_wrong_workflow_and_candidate_are_refused(self) -> None:
        for key, wrong in [
            ("id", 24), ("workflow_id", 13), ("path", "wrong.yml"),
            ("event", "push"), ("head_sha", "b" * 40),
            ("head_branch", "develop"), ("run_attempt", 2),
        ]:
            with self.subTest(key=key):
                run = successful_run()
                run[key] = wrong
                with self.assertRaisesRegex(ValueError, "expected"):
                    require_run(run, CANDIDATE, 12, 25)

    def test_missing_duplicate_skipped_and_incomplete_jobs_are_refused(self) -> None:
        good = successful_jobs()["jobs"]
        assert isinstance(good, list)
        for jobs in [good[:2], [good[0], good[0], good[2]],
                     [dict(good[0], conclusion="skipped"), *good[1:]],
                     [dict(good[0], status="in_progress"), *good[1:]]]:
            with self.subTest(jobs=jobs):
                with self.assertRaisesRegex(ValueError, "expected"):
                    require_jobs({"total_count": 3, "jobs": jobs})

    def test_failed_run_never_checks_jobs(self) -> None:
        api = FakeReleaseGitHub()
        api.run["conclusion"] = "failure"
        with self.assertRaisesRegex(ValueError, "expected success"):
            await_ci(api, CANDIDATE)

    def test_invalid_candidate_ref_is_refused_before_dispatch(self) -> None:
        api = FakeReleaseGitHub()
        with self.assertRaisesRegex(ValueError, "candidate ref"):
            await_ci(api, Candidate(CANDIDATE.repository, "develop", CANDIDATE.sha))
        self.assertEqual(api.posts, [])

    def test_dispatch_without_new_run_identity_is_refused(self) -> None:
        with self.assertRaisesRegex(ValueError, "expected positive workflow and run IDs"):
            await_ci(FakeDispatchWithoutRun(), CANDIDATE)

    def test_inactive_workflow_is_not_dispatched(self) -> None:
        api = FakeInactiveWorkflow()
        with self.assertRaisesRegex(ValueError, "expected active"):
            await_ci(api, CANDIDATE)
        self.assertEqual(api.posts, [])

    def test_pending_run_does_not_authorize_publication(self) -> None:
        api = FakeReleaseGitHub()
        api.run["status"] = "in_progress"
        clock = FakeReleaseClock()
        with self.assertRaisesRegex(TimeoutError, "expected completion"):
            await_ci(api, CANDIDATE, sleep=clock, monotonic=clock.now)
        self.assertEqual(clock.waits, 240)

    def test_gh_requests_have_a_bounded_timeout(self) -> None:
        for expired in (False, True):
            with self.subTest(expired=expired):
                process = FakeTimedGhProcess(expired)
                with patch("release_ci.subprocess.run", process):
                    if expired:
                        with self.assertRaises(subprocess.TimeoutExpired):
                            GhAPI().get("repos/owner/repo")
                    else:
                        self.assertEqual(GhAPI().get("repos/owner/repo"), {})
                self.assertEqual(process.timeout, 30)


if __name__ == "__main__":
    unittest.main()
