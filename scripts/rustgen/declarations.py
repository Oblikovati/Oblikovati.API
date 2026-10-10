# SPDX-License-Identifier: Apache-2.0
"""Generate Rust values and scalar contracts without altering the JSON protocol."""

import json

from names import PRIMITIVES, rust_path, rust_type, snake
from schema import Declaration, Field, Method, Shape

CUSTOM = {
    "Point",
    "Point2d",
    "Vector",
    "Vector2d",
    "UnitVector",
    "UnitVector2d",
    "Matrix",
    "Matrix2d",
    "Variant",
}
DERIVE = "#[derive(Debug, Clone, PartialEq, Default, serde::Serialize, serde::Deserialize)]\n"


def rust_doc(doc: str) -> str:
    lines = doc.strip().splitlines()
    if not lines:
        return ""
    # Preserve the Go examples as text; rustdoc must not compile them as Rust.
    return (
        "#[doc = "
        + json.dumps("```text\n" + doc.strip() + "\n```", ensure_ascii=False)
        + "]\n"
    )


def render_declaration(decl: Declaration, interfaces: set[tuple[str, str]]) -> str:
    if decl.package == "types" and decl.name in CUSTOM:
        return ""
    if decl.constant:
        return render_constant(decl, interfaces)
    if decl.alias and (decl.package, decl.name) in interfaces:
        target = rust_type(decl.shape, interfaces).removeprefix("std::sync::Arc<dyn ")[
            :-1
        ]
        return (
            rust_doc(decl.doc) + f"pub trait {decl.name}: {target} {{}}\n"
            f"impl<T: {target} + ?Sized> {decl.name} for T {{}}\n"
        )
    if decl.alias and decl.shape.kind == "basic":
        return (
            rust_doc(decl.doc)
            + f"pub type {decl.name} = {rust_type(decl.shape, interfaces)};\n"
        )
    if decl.shape.kind == "interface":
        return render_interface(decl, interfaces)
    if decl.shape.kind == "struct":
        if decl.package == "contract":
            fields = "\n".join(
                f"    pub {snake(item.name)}: {rust_type(item.shape, interfaces)},"
                for item in decl.shape.fields
            )
            return (
                rust_doc(decl.doc)
                + f"#[derive(Debug, Clone)]\npub struct {decl.name} {{\n{fields}\n}}\n"
            )
        body = "\n".join(
            render_field(item, interfaces)
            for item in decl.shape.fields
            if item.tag != "-"
        )
        result = (
            rust_doc(decl.doc)
            + DERIVE
            + f"#[serde(default)]\npub struct {decl.name} {{\n{body}\n}}\n"
        )
        if decl.feature_kind:
            result += f"impl crate::wire::featureargs::Arg for {decl.name} {{ fn kind(&self) -> String {{ {decl.feature_kind}.into() }} }}\n"
        return result
    if decl.shape.kind == "basic":
        return render_scalar(decl, interfaces)
    return (
        rust_doc(decl.doc)
        + f"pub type {decl.name} = {rust_type(decl.shape, interfaces)};\n"
    )


def render_scalar(decl: Declaration, interfaces: set[tuple[str, str]]) -> str:
    underlying = rust_type(decl.shape, interfaces)
    if underlying == "String":
        underlying = "std::borrow::Cow<'static, str>"
    derives = "Debug, Clone, Default, PartialEq, PartialOrd, serde::Serialize, serde::Deserialize"
    if underlying not in ("f32", "f64"):
        derives += ", Eq, Ord, Hash"
    if underlying != "std::borrow::Cow<'static, str>":
        derives += ", Copy"
    result = (
        rust_doc(decl.doc)
        + f"#[derive({derives})]\n#[repr(transparent)]\n#[serde(transparent)]\npub struct {decl.name}(pub {underlying});\n"
    )
    if underlying == "u8":
        result += f"impl From<u8> for {decl.name} {{ fn from(value: u8) -> Self {{ Self(value) }} }}\n"
        result += f"impl From<{decl.name}> for u8 {{ fn from(value: {decl.name}) -> Self {{ value.0 }} }}\n"
    return result


def render_constant(decl: Declaration, interfaces: set[tuple[str, str]]) -> str:
    typ = rust_type(decl.shape, interfaces)
    value = decl.value
    if decl.shape.kind == "named":
        if value.startswith('"'):
            return f"pub const {snake(decl.name).upper()}: {typ} = {typ}(std::borrow::Cow::Borrowed({value}));\n"
        value = f"{typ}({value})"
    return f"pub const {snake(decl.name).upper()}: {typ} = {value};\n"


def render_field(field: Field, interfaces: set[tuple[str, str]]) -> str:
    name, *options = field.tag.split(",")
    name = name or field.name
    attrs = ["flatten"] if field.embedded else ["rename = " + json.dumps(name)]
    predicate = omission_predicate(field.shape)
    if "omitempty" in options and predicate:
        attrs.append(f'skip_serializing_if = "{predicate}"')
    if needs_finite_check(field.shape):
        attrs.append('serialize_with = "crate::json::finite::serialize"')
    # Go accepts null for scalar/array/struct fields as their zero value.
    custom = field.shape.kind == "named" and (
        field.shape.name in CUSTOM
        or field.shape.package in ("json", "jsontext", "time")
    )
    if not custom:
        attrs.append('deserialize_with = "crate::json::null_default"')
    return (
        rust_doc(field.doc)
        + f"    #[serde({', '.join(attrs)})]\n    pub {snake(field.name)}: {rust_type(field.shape, interfaces)},"
    )


def omission_predicate(shape: Shape) -> str:
    if shape.kind == "pointer":
        return "Option::is_none"
    if shape.kind == "map":
        return "crate::json::map_is_empty"
    if shape.kind == "slice" or shape.underlying == "slice":
        if shape.package in ("json", "jsontext"):
            return "crate::json::JsonValue::is_empty"
        if shape.element and shape.element.name in ("byte", "uint8"):
            return "crate::json::Bytes::is_empty"
        if shape.element and shape.element.underlying in ("byte", "uint8"):
            return "crate::json::ByteSlice::is_empty"
        return "crate::json::slice_is_empty"
    return (
        "crate::json::is_zero"
        if shape.kind == "basic" or shape.underlying in PRIMITIVES
        else ""
    )


def needs_finite_check(shape: Shape) -> bool:
    if shape.kind == "basic":
        return shape.name in ("float64", "float32")
    return bool(shape.element and needs_finite_check(shape.element))


def render_interface(decl: Declaration, interfaces: set[tuple[str, str]]) -> str:
    params = (
        "<" + ", ".join(item.name for item in decl.shape.parameters) + ">"
        if decl.shape.parameters
        else ""
    )
    supers = [
        rust_type(item, interfaces)
        .removeprefix("std::sync::Arc<dyn ")
        .removesuffix(">")
        for item in decl.shape.embedded
    ]
    bounds = ": " + " + ".join(supers) if supers else ""
    methods = "\n".join(render_method(item, interfaces) for item in decl.shape.methods)
    return (
        rust_doc(decl.doc)
        + f"pub trait {decl.name}{params}{bounds} {{\n{methods}\n}}\n"
    )


def render_method(method: Method, interfaces: set[tuple[str, str]]) -> str:
    params = [
        f"{snake(item.name) if item.name else f'arg_{index}'}: {rust_type(item.shape, interfaces)}"
        for index, item in enumerate(method.parameters)
    ]
    result = method_result(method.results, interfaces)
    return (
        rust_doc(method.doc)
        + f"    fn {snake(method.name)}(&self{', ' if params else ''}{', '.join(params)}) -> {result};"
    )


def method_result(results: list[Field], interfaces: set[tuple[str, str]]) -> str:
    has_error = bool(results and results[-1].shape.name == "error")
    values = results[:-1] if has_error else results
    types = [rust_type(item.shape, interfaces) for item in values]
    result = types[0] if len(types) == 1 else "(" + ", ".join(types) + ")"
    return f"Result<{result}, crate::Error>" if has_error else result
