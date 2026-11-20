# SPDX-License-Identifier: Apache-2.0
"""Publish or safely finish one release whose tags were already promoted."""

import argparse
import hashlib
import json
import subprocess
import tempfile
import time
from pathlib import Path
from typing import Callable, Protocol

from promote_release import GitRepository, validate_release_identity

REQUEST_TIMEOUT = 120


class ReleaseService(Protocol):
    def view(self, tag: str) -> dict[str, object] | None: ...
    def create(self, tag: str, artifact: Path, notes: Path) -> bool: ...
    def upload(self, tag: str, artifact: Path) -> bool: ...
    def download(self, tag: str, name: str) -> bytes: ...


class GitHubReleases:
    """Use gh without a shell; example: GitHubReleases('owner/repo').view('v0.1.0')."""

    def __init__(self, repository: str) -> None:
        self.repository = repository

    def view(self, tag: str) -> dict[str, object] | None:
        """Read a published release; example: releases.view('v0.1.0')."""
        try:
            result = subprocess.run(
                ["gh", "api", f"repos/{self.repository}/releases/tags/{tag}"],
                capture_output=True, text=True, timeout=REQUEST_TIMEOUT,
            )
        except subprocess.TimeoutExpired:
            return None
        if result.returncode:
            return None
        value = json.loads(result.stdout)
        if not isinstance(value, dict):
            raise ValueError(f"release {value!r}: expected an object")
        return value

    def create(self, tag: str, artifact: Path, notes: Path) -> bool:
        """Create a release with an archive; example: releases.create(tag, artifact, notes)."""
        return self._try_command([
            "gh", "release", "create", tag, str(artifact), "--repo", self.repository,
            "--verify-tag", "--title", tag, "--notes-file", str(notes),
        ])

    def upload(self, tag: str, artifact: Path) -> bool:
        """Upload a missing archive; example: releases.upload(tag, artifact)."""
        return self._try_command([
            "gh", "release", "upload", tag, str(artifact), "--repo", self.repository,
        ])

    def _try_command(self, command: list[str]) -> bool:
        try:
            return subprocess.run(command, timeout=REQUEST_TIMEOUT).returncode == 0
        except subprocess.TimeoutExpired:
            return False

    def download(self, tag: str, name: str) -> bytes:
        """Read an existing archive; example: releases.download(tag, name)."""
        with tempfile.TemporaryDirectory(prefix="oblikovati-release-asset-") as folder:
            subprocess.run([
                "gh", "release", "download", tag, "--repo", self.repository,
                "--pattern", name, "--dir", folder,
            ], check=True, timeout=REQUEST_TIMEOUT)
            return (Path(folder) / name).read_bytes()


def require_promoted_tags(git: GitRepository, tag: str, candidate: str) -> None:
    """Check both annotated tags resolve to the candidate; example: require_promoted_tags(git, 'v0.1.0', sha)."""
    refs = [f"refs/tags/{tag}^{{}}", f"refs/tags/go/{tag}^{{}}"]
    actual = set(git.output("ls-remote", "--tags", "origin", *refs).splitlines())
    expected = {f"{candidate}\t{ref}" for ref in refs}
    if actual != expected:
        raise ValueError(f"remote tags {sorted(actual)!r}: expected {sorted(expected)!r}")


def require_matching_release(release: dict[str, object], tag: str, notes: str) -> list[object]:
    """Refuse a release not created by this workflow; example: require_matching_release(release, tag, notes)."""
    author = release.get("author")
    valid = (
        release.get("tag_name") == tag and release.get("name") == tag
        and release.get("draft") is False and release.get("prerelease") is False
        and isinstance(author, dict) and author.get("login") == "github-actions[bot]"
        and isinstance(release.get("body"), str)
        and release["body"].rstrip("\r\n") == notes.rstrip("\r\n")
    )
    assets = release.get("assets")
    if not valid or not isinstance(assets, list):
        raise ValueError(f"release {tag!r}: expected matching published notes, author and assets")
    return assets


def asset_present(assets: list[object], name: str) -> bool:
    """Require at most one uploaded asset by name; example: asset_present(assets, name)."""
    matches = [asset for asset in assets if isinstance(asset, dict) and asset.get("name") == name]
    if len(matches) > 1 or any(asset.get("state") != "uploaded" for asset in matches):
        raise ValueError(f"release asset {name!r}: expected at most one uploaded asset")
    return bool(matches)


def publication_attempt(service: ReleaseService, tag: str, artifact: Path, notes: Path) -> bool:
    """Create or inspect one release; example: publication_attempt(service, tag, artifact, notes)."""
    release = service.view(tag)
    if release is None:
        service.create(tag, artifact, notes)
        return False
    assets = require_matching_release(release, tag, notes.read_text(encoding="utf-8"))
    if not asset_present(assets, artifact.name):
        service.upload(tag, artifact)
        return False
    actual = service.download(tag, artifact.name)
    if hashlib.sha256(actual).digest() != hashlib.sha256(artifact.read_bytes()).digest():
        raise ValueError(f"release asset {artifact.name!r}: expected the verified crate bytes")
    return True


def publish_release(
    service: ReleaseService, git: GitRepository, candidate: str, tag: str,
    artifact: Path, notes: Path, sleep: Callable[[float], None] = time.sleep,
) -> None:
    """Publish or recover without replacing any existing asset; example: publish_release(service, git, sha, tag, artifact, notes)."""
    validate_release_identity(candidate, candidate, tag)
    if artifact.is_symlink() or not artifact.is_file() or artifact.name != f"oblikovati-api-{tag[1:]}.crate":
        raise ValueError(f"release artifact {artifact}: expected the named regular crate archive")
    require_promoted_tags(git, tag, candidate)
    for attempt in range(3):
        if publication_attempt(service, tag, artifact, notes):
            return
        if attempt < 2:
            sleep(5)
    raise ValueError(f"release {tag!r}: expected matching published notes and crate after retries")


def main() -> None:
    """Publish a promoted release; run `python3 scripts/publish_release.py --help`."""
    parser = argparse.ArgumentParser(description=__doc__)
    for name in ("repository", "candidate", "tag"):
        parser.add_argument(name)
    for name in ("artifact", "notes"):
        parser.add_argument(name, type=Path)
    args = parser.parse_args()
    publish_release(GitHubReleases(args.repository), GitRepository(Path.cwd()), args.candidate,
                    args.tag, args.artifact, args.notes)


if __name__ == "__main__":
    main()
