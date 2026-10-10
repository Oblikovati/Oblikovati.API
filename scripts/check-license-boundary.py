# SPDX-License-Identifier: Apache-2.0
"""Check the Go and Rust runtime graphs for dependencies on the GPL application."""

import json
import os
import subprocess


def main() -> None:
    modules = subprocess.check_output(
        ["go", "list", "-deps", "./..."],
        cwd="go",
        env={**os.environ, "GOWORK": "off"},
        text=True,
    ).splitlines()
    forbidden = [
        name
        for name in modules
        if (name == "oblikovati.org" or name.startswith("oblikovati.org/"))
        and not (name == "oblikovati.org/api" or name.startswith("oblikovati.org/api/"))
    ]
    graph = json.loads(
        subprocess.check_output(
            [
                "cargo",
                "metadata",
                "--manifest-path",
                "rust/Cargo.toml",
                "--format-version",
                "1",
                "--locked",
            ],
            text=True,
        )
    )
    forbidden += [
        package["name"]
        for package in graph["packages"]
        if "GPL" in (package.get("license") or "").replace("LGPL", "")
    ]
    if forbidden:
        raise SystemExit(
            f"API dependencies {forbidden}: expected an Apache-compatible graph without the application"
        )
    print("Go and Rust API dependencies respect the application boundary")


if __name__ == "__main__":
    main()
