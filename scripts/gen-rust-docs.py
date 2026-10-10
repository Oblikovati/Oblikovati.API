# SPDX-License-Identifier: Apache-2.0
"""Build a wiki reference from the checked Go schema and actual Rust source paths."""

import json
from pathlib import Path
import re

BASE = "https://github.com/Oblikovati/Oblikovati.API/blob/develop/"


def locations() -> dict[tuple[str, str], str]:
    found: dict[tuple[str, str], str] = {}
    for path in sorted(Path("rust/src").rglob("*.rs")):
        package = "featureargs" if "featureargs" in path.parts else path.parts[2]
        for number, line in enumerate(path.read_text().splitlines(), 1):
            match = re.match(
                r"(?:pub (?:struct|enum|trait|type|const) |(?:coordinate_value|unit_value)!\()(\w+)",
                line,
            )
            if match:
                found[package, match[1]] = f"{BASE}{path}#L{number}"
    return found


def main() -> None:
    schema = json.loads(Path("build/api-schema.json").read_text())
    source = locations()
    lines = [
        "# Rust API Reference",
        "",
        "Generated from the checked API definitions. The links open the exact Rust declarations.",
        "",
        "Build local Rustdoc with `cargo doc --manifest-path rust/Cargo.toml --no-deps --open`.",
        "",
        "See [[First Steps]], [[Typed Operations]] and [[C ABI Add-ins]] for compiled examples.",
        "",
    ]
    for package in ("types", "contract", "wire", "featureargs"):
        lines += [f"## {package}", "", "| Type or trait | Source |", "| --- | --- |"]
        for item in schema["Declarations"]:
            if item["Package"] == package and not item["Constant"]:
                lines.append(
                    f"| `{item['Name']}` | [Rust source]({source[package, item['Name']]}) |"
                )
        lines.append("")
    lines += [
        "## Client methods",
        "",
        "Rust methods take a reference to the canonical request DTO. Methods without a request body take no argument.",
        "",
        "| Group | Rust method | Wire constant |",
        "| --- | --- | --- |",
    ]
    from sys import path

    path.insert(0, str(Path("scripts/rustgen").resolve()))
    from names import snake

    lines += [
        f"| `{item['Group']}` | `{snake(item['Name'])}` | `{snake(item['Method']).upper()}` |"
        for item in schema["Calls"]
    ]
    Path("build/wiki/Rust-API-Reference.md").write_text("\n".join(lines) + "\n")


if __name__ == "__main__":
    main()
