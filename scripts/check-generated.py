# SPDX-License-Identifier: Apache-2.0
"""Fail when regenerating either language or its examples changes tracked artifacts."""

from hashlib import sha256
from pathlib import Path
import subprocess

ROOTS = ("rust/src", "rust/tests", "go/scripts/apifixtures", "go/client", "docs/wiki")


def snapshot() -> dict[str, str]:
    paths = [path for root in ROOTS for path in Path(root).rglob("*") if path.is_file()]
    paths += [Path("rust/Cargo.toml"), Path("rust/Cargo.lock")]
    return {str(path): sha256(path.read_bytes()).hexdigest() for path in paths}


def main() -> None:
    before = snapshot()
    subprocess.run(["make", "generate"], check=True)
    after = snapshot()
    changed = sorted(
        path
        for path in before.keys() | after.keys()
        if before.get(path) != after.get(path)
    )
    if changed:
        raise SystemExit(
            "Generated API files are out of sync; commit make generate output:\n"
            + "\n".join(changed)
        )
    print("Go/Rust definitions, version, fixtures and wiki examples are in sync")


if __name__ == "__main__":
    main()
