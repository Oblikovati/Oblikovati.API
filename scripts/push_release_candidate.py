# SPDX-License-Identifier: Apache-2.0
"""Create one release candidate branch without overwriting a remote ref."""

import argparse
import re
from pathlib import Path

from promote_release import GitRepository, require_release_base, validate_release_identity


def push_candidate(git: GitRepository, base: str, candidate: str, branch: str) -> None:
    """Push a new candidate ref; example: push_candidate(git, base, sha, 'api-release-candidate/1-1')."""
    validate_release_identity(base, candidate, "v0.0.0")
    if not re.fullmatch(r"api-release-candidate/[0-9]+-[0-9]+", branch):
        raise ValueError(f"candidate branch {branch!r}: expected a unique release candidate ref")
    require_release_base(git, base, candidate)
    remote_ref = f"refs/heads/{branch}"
    existing = git.output("ls-remote", "--heads", "origin", remote_ref)
    if existing:
        raise ValueError(f"candidate branch {existing!r}: expected an absent remote ref")
    result = git.output(
        "push", "--porcelain", f"--force-with-lease={remote_ref}:", "origin",
        f"HEAD:{remote_ref}",
    )
    created = f"*\tHEAD:{remote_ref}\t[new branch]"
    if created not in result.splitlines():
        raise ValueError(f"candidate push {result!r}: expected {created!r}")


def main() -> None:
    """Create one candidate ref; run `python3 scripts/push_release_candidate.py --help`."""
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("base")
    parser.add_argument("candidate")
    parser.add_argument("branch")
    args = parser.parse_args()
    push_candidate(GitRepository(Path.cwd()), args.base, args.candidate, args.branch)


if __name__ == "__main__":
    main()
