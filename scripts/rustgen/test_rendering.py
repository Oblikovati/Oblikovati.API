# SPDX-License-Identifier: Apache-2.0
"""Check mapping rules and refusal paths independently of the generated snapshots."""

import unittest
from clients import render_group, render_rpc
from declarations import (
    render_declaration,
    render_field,
    render_method,
    omission_predicate,
    needs_finite_check,
    method_result,
    rust_doc,
)
from names import snake, rust_type, rust_path
from schema import (
    Shape,
    Field,
    Method,
    Declaration,
    Rpc,
    declarations_from_json,
    calls_from_json,
)


def declaration(shape: Shape, **options: object) -> Declaration:
    result = Declaration(
        "Example", "types", "example.go", "Example documents intent.", shape, False, ""
    )
    for name, value in options.items():
        setattr(result, name, value)
    return result


class TypeMappingTests(unittest.TestCase):
    def test_names_and_paths(self) -> None:
        for original, expected in [
            ("HTTPServerID", "http_server_id"),
            ("Type", "type_"),
            ("Sketch3dID", "sketch3d_id"),
        ]:
            self.assertEqual(snake(original), expected)
        self.assertEqual(
            rust_path("featureargs", "Extrude"), "crate::wire::featureargs::Extrude"
        )

    def test_container_types_keep_nullability_and_bytes(self) -> None:
        integer = Shape("basic", "uint64")
        for shape, expected in [
            (integer, "u64"),
            (Shape("empty"), "()"),
            (Shape("parameter", "T"), "T"),
            (Shape("pointer", element=integer), "Option<Box<u64>>"),
            (Shape("slice", element=integer), "Option<Vec<u64>>"),
            (Shape("slice", element=Shape("basic", "byte")), "crate::json::Bytes"),
            (Shape("array", length=3, element=integer), "[u64; 3]"),
            (
                Shape("map", key=integer, element=integer),
                "Option<std::collections::BTreeMap<u64, u64>>",
            ),
            (Shape("struct"), "crate::json::EmptyObject"),
        ]:
            self.assertEqual(rust_type(shape, set()), expected)
        enum = Shape("named", "Channel", "types", underlying="uint8")
        self.assertEqual(
            rust_type(Shape("slice", element=enum), set()),
            "crate::json::ByteSlice<crate::types::Channel>",
        )

    def test_named_types_and_traits(self) -> None:
        for shape, expected in [
            (Shape("named", "RawMessage", "json"), "crate::json::JsonValue"),
            (Shape("named", "Time", "time"), "crate::json::Timestamp"),
            (Shape("named", "error"), "crate::Error"),
        ]:
            self.assertEqual(rust_type(shape, set()), expected)
        generic = Shape(
            "named", "Enumerable", "contract", parameters=[Shape("basic", "uint64")]
        )
        self.assertEqual(
            rust_type(generic, {("contract", "Enumerable")}),
            "std::sync::Arc<dyn crate::contract::Enumerable<u64>>",
        )

    def test_unknown_types_fail_instead_of_erasing_the_contract(self) -> None:
        for shape in [
            Shape("channel"),
            Shape("slice"),
            Shape("struct", fields=[Field("X", Shape("basic", "int"))]),
        ]:
            with self.assertRaisesRegex(ValueError, "expected"):
                rust_type(shape, set())


class DeclarationTests(unittest.TestCase):
    def test_custom_values_are_owned_by_the_runtime(self) -> None:
        self.assertEqual(
            render_declaration(declaration(Shape("struct"), name="Point"), set()), ""
        )
        self.assertEqual(rust_doc(""), "")
        self.assertIn("```text", rust_doc("original Go example"))

    def test_scalars_constants_and_aliases(self) -> None:
        for name in ["uint8", "string", "float64"]:
            text = render_declaration(declaration(Shape("basic", name)), set())
            self.assertIn("pub struct Example", text)
        text = render_declaration(
            declaration(Shape("basic", "string"), alias=True), set()
        )
        self.assertIn("pub type Example = String", text)
        for shape, value, expected in [
            (Shape("basic", "untyped string"), '""', 'pub const EXAMPLE: &str = "";'),
            (Shape("named", "Mode", "types"), "7", "crate::types::Mode(7)"),
            (Shape("named", "Name", "types"), '"a"', 'Cow::Borrowed("a")'),
        ]:
            self.assertIn(
                expected,
                render_declaration(
                    declaration(shape, constant=True, value=value), set()
                ),
            )
        self.assertIn(
            "pub type Example",
            render_declaration(
                declaration(Shape("slice", element=Shape("basic", "int"))), set()
            ),
        )

    def test_struct_fields_preserve_wire_names_and_omission_rules(self) -> None:
        field = Field("Type", Shape("basic", "float64"), "type,omitempty")
        output = render_field(field, set())
        self.assertIn('rename = "type"', output)
        self.assertIn("pub type_: f64", output)
        self.assertIn("finite::serialize", output)
        self.assertIn("null_default", output)
        shape = Shape(
            "struct", fields=[field, Field("Hidden", Shape("basic", "int"), "-")]
        )
        self.assertNotIn("hidden:", render_declaration(declaration(shape), set()))
        self.assertIn(
            "impl crate::wire::featureargs::Arg",
            render_declaration(declaration(shape, feature_kind='"extrude"'), set()),
        )
        self.assertIn(
            "derive(Debug, Clone)",
            render_declaration(declaration(shape, package="contract"), set()),
        )

    def test_omission_is_not_applied_to_structs_or_fixed_arrays(self) -> None:
        for shape in [
            Shape("struct"),
            Shape("array", element=Shape("basic", "float64"), length=3),
        ]:
            self.assertEqual(omission_predicate(shape), "")
        cases = [
            (Shape("pointer"), "Option::is_none"),
            (Shape("map"), "crate::json::map_is_empty"),
            (
                Shape("slice", element=Shape("basic", "byte")),
                "crate::json::Bytes::is_empty",
            ),
            (
                Shape("slice", element=Shape("named", underlying="uint8")),
                "crate::json::ByteSlice::is_empty",
            ),
            (
                Shape("slice", element=Shape("basic", "int")),
                "crate::json::slice_is_empty",
            ),
            (
                Shape("named", package="json", underlying="slice"),
                "crate::json::JsonValue::is_empty",
            ),
        ]
        for shape, expected in cases:
            self.assertEqual(omission_predicate(shape), expected)
        self.assertTrue(
            needs_finite_check(Shape("slice", element=Shape("basic", "float64")))
        )
        self.assertFalse(needs_finite_check(Shape("basic", "int")))

    def test_embedded_fields_and_custom_json(self) -> None:
        embedded = Field("Info", Shape("named", "Info", "wire"), embedded=True)
        self.assertIn("flatten", render_field(embedded, set()))
        self.assertNotIn(
            "null_default",
            render_field(Field("Origin", Shape("named", "Point", "types")), set()),
        )

    def test_contract_methods_errors_generics_and_aliases(self) -> None:
        method = Method(
            "Find",
            [Field("", Shape("basic", "uint64"))],
            [Field("", Shape("basic", "string")), Field("", Shape("named", "error"))],
        )
        self.assertIn("arg_0: u64", render_method(method, set()))
        self.assertEqual(
            method_result(method.results, set()), "Result<String, crate::Error>"
        )
        self.assertEqual(method_result([], set()), "()")
        interfaces = {("contract", "Enumerable"), ("contract", "Example")}
        generic = Shape(
            "interface", methods=[method], parameters=[Shape("parameter", "T")]
        )
        self.assertIn(
            "trait Example<T>",
            render_declaration(declaration(generic, package="contract"), interfaces),
        )
        alias = declaration(
            Shape(
                "named", "Enumerable", "contract", parameters=[Shape("basic", "uint64")]
            ),
            package="contract",
            alias=True,
        )
        self.assertIn(
            "impl<T: crate::contract::Enumerable<u64> + ?Sized>",
            render_declaration(alias, interfaces),
        )


class ClientRenderingTests(unittest.TestCase):
    def test_empty_typed_and_void_calls(self) -> None:
        rpc = Rpc(
            "Application",
            "ApiVersion",
            "MethodApiVersion",
            Shape("empty"),
            Shape("basic", "int"),
            "",
        )
        self.assertIn("None::<&()>", render_rpc(rpc, set()))
        self.assertIn("pub fn application", render_group("Application", [rpc], set()))
        self.assertNotIn("pub struct", render_group("Application", [rpc], set(), False))
        rpc.group, rpc.request, rpc.response = (
            "Client",
            Shape("basic", "int"),
            Shape("empty"),
        )
        self.assertIn("self.invoke_void", render_rpc(rpc, set()))
        self.assertIn("super::Client<T>", render_group("Client", [rpc], set()))

    def test_schema_ingestion_retains_empty_constants_and_aliases(self) -> None:
        schema: dict[str, object] = {
            "Declarations": [
                {
                    "Name": "Empty",
                    "Package": "types",
                    "File": "a.go",
                    "Doc": "",
                    "Shape": {"Kind": "basic", "Name": "string"},
                    "Constant": True,
                    "Value": "",
                    "Alias": False,
                }
            ],
            "Calls": [],
        }
        parsed = declarations_from_json(schema)
        self.assertTrue(parsed[0].constant)
        self.assertEqual(parsed[0].value, "")
        self.assertEqual(calls_from_json(schema), [])


if __name__ == "__main__":
    unittest.main()
