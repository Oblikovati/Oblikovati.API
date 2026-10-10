# SPDX-License-Identifier: Apache-2.0
"""Exercise real output assembly in isolated directories, including stale-file removal."""

from contextlib import chdir
from pathlib import Path
from tempfile import TemporaryDirectory
import unittest
from main import generate_bindings, validate_calls, write_chunks
from schema import Declaration, Shape


class GenerationTests(unittest.TestCase):
    def test_missing_typed_client_is_a_hard_error(self) -> None:
        method = Declaration(
            "MethodMissing",
            "wire",
            "methods.go",
            "",
            Shape("basic", "untyped string"),
            True,
            '"missing"',
        )
        with self.assertRaisesRegex(
            ValueError, "MethodMissing.*expected a typed client"
        ):
            validate_calls([method], [])

    def test_output_is_bounded_and_stale_parts_are_deleted(self) -> None:
        with TemporaryDirectory() as directory:
            root = Path(directory)
            chunks = [
                f"pub const VALUE_{index}: u64 = {index};\n" for index in range(600)
            ]
            modules = write_chunks(root, chunks)
            self.assertGreater(len(modules), 1)
            for path in root.glob("*.rs"):
                self.assertLess(len(path.read_text().splitlines()), 500)
            write_chunks(root, chunks[:1])
            self.assertEqual(len(list(root.glob("*.rs"))), 1)

    def test_complete_generation_preserves_shapes_and_client_coverage(self) -> None:
        with TemporaryDirectory() as directory, chdir(directory):
            Path("go/client").mkdir(parents=True)
            generate_bindings(sample_schema(), Path("rust/src"))
            text = Path("rust/src/wire/bindings_000.rs").read_text()
            self.assertIn('rename = "id"', text)
            self.assertIn("pub id: u64", text)
            self.assertIn("METHOD_GET", text)
            self.assertIn(
                "sample_get",
                Path("rust/tests/client_generated/calls_000.rs").read_text(),
            )
            self.assertIn(
                "TestRustParitySampleGet",
                Path("go/client/rust_parity_000_test.go").read_text(),
            )
            self.assertIn(
                "reflect.TypeFor[wire.Reply]",
                Path("go/scripts/apifixtures/registry_000.go").read_text(),
            )
            self.assertIn(
                "kind()", Path("rust/tests/feature_kinds/kinds_000.rs").read_text()
            )
            self.assertIn(
                "point_matches_go_fields",
                Path("rust/tests/geometry_contract.rs").read_text(),
            )


def sample_schema() -> dict[str, object]:
    integer = {"Kind": "basic", "Name": "uint64"}
    reply = {
        "Name": "Reply",
        "Package": "wire",
        "File": "reply.go",
        "Doc": "Reply keeps integer IDs.",
        "Constant": False,
        "Shape": {
            "Kind": "struct",
            "Fields": [{"Name": "ID", "Type": integer, "Tag": "id"}],
        },
    }
    constant = {
        "Name": "MethodGet",
        "Package": "wire",
        "File": "reply.go",
        "Doc": "",
        "Constant": True,
        "Value": '"sample.get"',
        "Shape": {"Kind": "basic", "Name": "untyped string"},
    }
    feature = {
        "Name": "Sample",
        "Package": "featureargs",
        "File": "sample.go",
        "Doc": "",
        "Constant": False,
        "Kind": '"sample"',
        "Shape": {"Kind": "struct"},
    }
    rpc = {
        "Group": "Sample",
        "Name": "Get",
        "Method": "MethodGet",
        "Doc": "",
        "Request": {"Kind": "empty"},
        "Response": {"Kind": "named", "Name": "Reply", "Package": "wire"},
    }
    point = {
        "Name": "Point",
        "Package": "types",
        "File": "point.go",
        "Doc": "",
        "Constant": False,
        "Shape": {
            "Kind": "struct",
            "Fields": [
                {"Name": name, "Type": {"Kind": "basic", "Name": "float64"}}
                for name in ("X", "Y", "Z")
            ],
        },
    }
    return {"Declarations": [reply, constant, feature, point], "Calls": [rpc]}


if __name__ == "__main__":
    unittest.main()
