# SPDX-License-Identifier: Apache-2.0
"""Dispatch and validate the hosted CI run for one release candidate."""

import argparse
import json
import re
import subprocess
import time
from dataclasses import dataclass
from typing import Callable, Protocol

API_VERSION = "2026-03-10"
REQUEST_TIMEOUT = 30
JOBS = {"gate", "rust-portability (windows-latest)", "rust-portability (macos-latest)"}
WORKFLOW = ".github/workflows/ci.yml"


class GitHubAPI(Protocol):
    def get(self, endpoint: str) -> dict[str, object]: ...
    def post(self, endpoint: str, payload: dict[str, str]) -> dict[str, object]: ...


class GhAPI:
    """Use gh without a shell; example: GhAPI().get('repos/owner/repo')."""

    def get(self, endpoint: str) -> dict[str, object]:
        """Read an API object; example: api.get('repos/owner/repo')."""
        return self._call(endpoint, None)

    def post(self, endpoint: str, payload: dict[str, str]) -> dict[str, object]:
        """Post an API object; example: api.post(endpoint, {'ref': 'topic'})."""
        return self._call(endpoint, payload)

    def _call(self, endpoint: str, payload: dict[str, str] | None) -> dict[str, object]:
        command = ["gh", "api", "-H", f"X-GitHub-Api-Version: {API_VERSION}"]
        if payload is not None:
            command += ["--method", "POST", "--input", "-"]
        result = subprocess.run(
            command + [endpoint], input=json.dumps(payload) if payload is not None else None,
            text=True, capture_output=True, check=True, timeout=REQUEST_TIMEOUT,
        )
        value = json.loads(result.stdout)
        if not isinstance(value, dict):
            raise ValueError(f"GitHub response {value!r}: expected an object")
        return value


@dataclass(frozen=True)
class Candidate:
    """The identity that CI must test; example: Candidate('o/r', 'release/a', 'a'*40)."""

    repository: str
    branch: str
    sha: str

    def validate(self) -> None:
        """Reject malformed candidate identity; example: candidate.validate()."""
        if not re.fullmatch(r"[A-Za-z0-9_.-]+/[A-Za-z0-9_.-]+", self.repository):
            raise ValueError(f"repository {self.repository!r}: expected owner/repository")
        if not re.fullmatch(r"api-release-candidate/[0-9]+-[0-9]+", self.branch):
            raise ValueError(f"branch {self.branch!r}: expected a unique release candidate ref")
        if not re.fullmatch(r"[0-9a-f]{40}", self.sha):
            raise ValueError(f"candidate SHA {self.sha!r}: expected a full lowercase SHA-1")


def require_run(run: dict[str, object], candidate: Candidate, workflow_id: int, run_id: int) -> None:
    """Reject a different CI invocation; example: require_run(run, candidate, workflow_id, run_id)."""
    expected = {
        "id": run_id, "workflow_id": workflow_id, "path": WORKFLOW,
        "event": "workflow_dispatch", "head_sha": candidate.sha,
        "head_branch": candidate.branch, "run_attempt": 1,
    }
    for key, value in expected.items():
        if run.get(key) != value:
            raise ValueError(f"CI {key} {run.get(key)!r}: expected {value!r}")


def require_jobs(result: dict[str, object]) -> None:
    """Require the complete matrix to pass; example: require_jobs(response)."""
    jobs = result.get("jobs")
    if result.get("total_count") != len(JOBS) or not isinstance(jobs, list) or len(jobs) != len(JOBS):
        raise ValueError(f"CI jobs {result!r}: expected exactly {sorted(JOBS)!r}")
    names = [job.get("name") for job in jobs if isinstance(job, dict)]
    if len(names) != len(JOBS) or set(names) != JOBS:
        raise ValueError(f"CI job names {names!r}: expected {sorted(JOBS)!r}")
    for job in jobs:
        if job.get("status") != "completed" or job.get("conclusion") != "success":
            raise ValueError(f"CI job {job!r}: expected completed success")


def dispatch(api: GitHubAPI, candidate: Candidate) -> tuple[int, int]:
    """Dispatch CI and retain its returned identity; example: dispatch(api, candidate)."""
    prefix = f"repos/{candidate.repository}/actions/workflows/ci.yml"
    workflow = api.get(prefix)
    if workflow.get("path") != WORKFLOW or workflow.get("state") != "active":
        raise ValueError(f"CI workflow {workflow!r}: expected active {WORKFLOW}")
    workflow_id = workflow.get("id")
    if type(workflow_id) is not int or workflow_id <= 0:
        raise ValueError(f"CI workflow ID {workflow_id!r}: expected a positive integer")
    response = api.post(f"{prefix}/dispatches", {"ref": candidate.branch})
    run_id = response.get("workflow_run_id")
    if type(run_id) is not int or run_id <= 0:
        raise ValueError(f"CI dispatch {response!r}: expected positive workflow and run IDs")
    return workflow_id, run_id


def await_ci(
    api: GitHubAPI, candidate: Candidate, sleep: Callable[[float], None] = time.sleep,
    monotonic: Callable[[], float] = time.monotonic,
) -> int:
    """Wait for this candidate's complete hosted CI; example: await_ci(GhAPI(), candidate)."""
    candidate.validate()
    workflow_id, run_id = dispatch(api, candidate)
    base = f"repos/{candidate.repository}/actions/runs/{run_id}"
    deadline = monotonic() + 3600
    while monotonic() < deadline:
        run = api.get(base)
        require_run(run, candidate, workflow_id, run_id)
        if run.get("status") == "completed":
            if run.get("conclusion") != "success":
                raise ValueError(f"CI result {run!r}: expected success")
            require_jobs(api.get(f"{base}/attempts/1/jobs?per_page=100"))
            return run_id
        sleep(min(15, max(0, deadline - monotonic())))
    raise TimeoutError(f"CI run {run_id}: expected completion within one hour")


def main() -> None:
    """Dispatch and verify a candidate; run `python3 scripts/release_ci.py --help`."""
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("repository")
    parser.add_argument("branch")
    parser.add_argument("sha")
    args = parser.parse_args()
    print(f"Verified CI run {await_ci(GhAPI(), Candidate(args.repository, args.branch, args.sha))}")


if __name__ == "__main__":
    main()
