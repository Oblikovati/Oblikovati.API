# SPDX-License-Identifier: Apache-2.0
"""Keep the C ABI runner independent from Cargo's output-directory choice."""
import json
import os
from pathlib import Path
import platform
import shutil
import subprocess
import tempfile
import unittest

FAKE_COMPILER = '''#!/usr/bin/env python3
import os, pathlib, sys
output = pathlib.Path(sys.argv[sys.argv.index("-o") + 1])
output.write_text(os.environ["FAKE_ABI_HOST"])
output.chmod(0o755)
'''
FAKE_GO = '''#!/usr/bin/env python3
import pathlib, sys
if sys.argv[1] == "run":
    print("test-api")
elif sys.argv[1] == "build":
    pathlib.Path(sys.argv[sys.argv.index("-o") + 1]).touch()
else:
    raise SystemExit("unexpected Go arguments: " + repr(sys.argv))
'''
FAKE_CARGO = '''#!/usr/bin/env python3
import json, os, pathlib, sys
root = pathlib.Path.cwd()
target = pathlib.Path(os.environ.get("CARGO_TARGET_DIR", root / "rust/target"))
if sys.argv[1] == "metadata":
    print(json.dumps({"target_directory": str(target)}))
elif sys.argv[1] == "build":
    library = target / "debug/examples/libc_addin.so"
    library.parent.mkdir(parents=True, exist_ok=True)
    library.touch()
else:
    raise SystemExit("unexpected Cargo arguments: " + repr(sys.argv))
'''
FAKE_HOST = '''#!/usr/bin/env python3
import json, os, pathlib, sys
if not pathlib.Path(sys.argv[1]).is_file():
    raise SystemExit("missing ABI library: " + sys.argv[1])
with open(os.environ["FAKE_ABI_LOG"], "a") as stream:
    stream.write(json.dumps(sys.argv[1:]) + "\\n")
'''


class FakeAbiToolchain:
    """Named fake commands exercise the real shell runner without compiling an SDK."""

    def __init__(self, root: Path) -> None:
        self.root = root
        self.binary = root / "bin"
        self.binary.mkdir()
        self.log = root / "calls.jsonl"
        (root / "scripts").mkdir()
        (root / "go").mkdir()
        shutil.copy(Path(__file__).with_name("test-c-abi.sh"), root / "scripts/test-c-abi.sh")
        for name, source in {"cc": FAKE_COMPILER, "go": FAKE_GO, "cargo": FAKE_CARGO}.items():
            executable = self.binary / name
            executable.write_text(source)
            executable.chmod(0o755)

    def run(self, target: Path | None) -> list[list[str]]:
        environment = os.environ.copy()
        environment.update(PATH=str(self.binary) + os.pathsep + environment["PATH"],
                           FAKE_ABI_LOG=str(self.log), FAKE_ABI_HOST=FAKE_HOST)
        environment.pop("CARGO_TARGET_DIR", None)
        if target is not None:
            environment["CARGO_TARGET_DIR"] = str(target)
        subprocess.run(["bash", "scripts/test-c-abi.sh"], cwd=self.root, env=environment,
                       check=True, capture_output=True, text=True)
        return [json.loads(line) for line in self.log.read_text().splitlines()]


@unittest.skipUnless(platform.system() == "Linux", "the fake ABI artifact uses the Linux suffix")
class CargoAbiTargetTests(unittest.TestCase):
    def test_custom_target_with_spaces_is_used_to_load_the_library(self) -> None:
        with tempfile.TemporaryDirectory(prefix="api ABI ") as directory:
            root = Path(directory)
            target = root / "custom Cargo target"
            calls = FakeAbiToolchain(root).run(target)
            self.assertEqual(len(calls), 2)
            self.assertEqual(calls[1], [str(target / "debug/examples/libc_addin.so"), "test-api"])

    def test_default_manifest_target_still_loads_the_library(self) -> None:
        with tempfile.TemporaryDirectory(prefix="api ABI ") as directory:
            root = Path(directory)
            calls = FakeAbiToolchain(root).run(None)
            self.assertEqual(len(calls), 2)
            self.assertEqual(calls[1][0], str(root / "rust/target/debug/examples/libc_addin.so"))
