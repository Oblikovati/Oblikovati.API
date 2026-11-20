# SPDX-License-Identifier: Apache-2.0
"""Promote one tested release commit without overwriting a changed develop branch."""

import argparse
import re
import subprocess
from pathlib import Path


class GitRepository:
    """Run Git in one checkout; example: GitRepository(Path('.')).output('rev-parse', 'HEAD')."""

    def __init__(self, root: Path) -> None:
        self.root = root

    def output(self, *args: str) -> str:
        """Read Git output; example: repository.output('rev-parse', 'HEAD')."""
        return subprocess.check_output(["git", "-C", str(self.root), *args], text=True).strip()

    def run(self, *args: str) -> None:
        """Run a Git update; example: repository.run('tag', '-a', 'v0.1.0', '-m', 'Release')."""
        subprocess.run(["git", "-C", str(self.root), *args], check=True)


def validate_release_identity(base: str, candidate: str, tag: str) -> None:
    """Require unambiguous Git identities; example: validate_release_identity('a'*40, 'b'*40, 'v0.1.0')."""
    for name, value in (("base", base), ("candidate", candidate)):
        if not re.fullmatch(r"[0-9a-f]{40}", value):
            raise ValueError(f"{name} SHA {value!r}: expected a full lowercase SHA-1")
    if not re.fullmatch(r"v(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)", tag):
        raise ValueError(f"release tag {tag!r}: expected vMAJOR.MINOR.PATCH")


def require_release_base(git: GitRepository, base: str, candidate: str) -> None:
    """Require the exact tested commit and unchanged base; example: require_release_base(git, base, candidate)."""
    head, parent = git.output("rev-parse", "HEAD"), git.output("rev-parse", "HEAD^")
    if head != candidate or parent != base:
        raise ValueError(f"candidate {head!r} parent {parent!r}: expected {candidate!r} on {base!r}")
    if git.output("status", "--porcelain", "--untracked-files=all"):
        raise ValueError("release checkout: expected no uncommitted changes")
    remote = git.output("ls-remote", "--exit-code", "--heads", "origin", "refs/heads/develop")
    if remote != f"{base}\trefs/heads/develop":
        raise ValueError(f"remote develop {remote!r}: expected {base!r}")


def promote_release(git: GitRepository, base: str, candidate: str, tag: str) -> None:
    """Tag and atomically promote a tested direct child; example: promote_release(git, base, sha, 'v0.1.0')."""
    validate_release_identity(base, candidate, tag)
    require_release_base(git, base, candidate)
    git.run("tag", "-a", tag, "-m", f"Release {tag}")
    git.run("tag", "-a", f"go/{tag}", "-m", f"Go module {tag}")
    git.run(
        "push", "--atomic", f"--force-with-lease=refs/heads/develop:{base}", "origin",
        "HEAD:refs/heads/develop", f"refs/tags/{tag}", f"refs/tags/go/{tag}",
    )


def main() -> None:
    """Promote a tested candidate; run `python3 scripts/promote_release.py --help`."""
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("base")
    parser.add_argument("candidate")
    parser.add_argument("tag")
    args = parser.parse_args()
    promote_release(GitRepository(Path.cwd()), args.base, args.candidate, args.tag)


if __name__ == "__main__":
    main()
