# SPDX-License-Identifier: Apache-2.0
"""Rust identifiers and type mapping; reject unknown shapes rather than erase types."""

import re

from schema import Shape

PRIMITIVES = {
    "int": "i64",
    "int8": "i8",
    "int16": "i16",
    "int32": "i32",
    "int64": "i64",
    "uint": "u64",
    "uint8": "u8",
    "byte": "u8",
    "uint16": "u16",
    "uint32": "u32",
    "uint64": "u64",
    "uintptr": "usize",
    "float32": "f32",
    "float64": "f64",
    "string": "String",
    "bool": "bool",
    "untyped int": "i64",
    "untyped float": "f64",
    "untyped string": "&str",
    "untyped bool": "bool",
}
KEYWORDS = {
    "type",
    "ref",
    "move",
    "loop",
    "match",
    "self",
    "mod",
    "fn",
    "use",
    "in",
    "as",
    "box",
    "where",
    "impl",
    "trait",
    "async",
    "await",
    "dyn",
    "const",
    "static",
    "enum",
    "struct",
    "super",
    "crate",
    "pub",
    "return",
    "continue",
    "break",
    "yield",
    "gen",
    "virtual",
    "abstract",
    "become",
    "do",
    "final",
    "macro",
    "override",
    "priv",
    "typeof",
    "unsized",
    "try",
}


def snake(name: str) -> str:
    first = re.sub(r"([A-Z]+)([A-Z][a-z])", r"\1_\2", name)
    result = re.sub(r"([a-z0-9])([A-Z])", r"\1_\2", first).lower()
    return result + "_" if result in KEYWORDS else result


def rust_path(package: str, name: str) -> str:
    module = "wire::featureargs" if package == "featureargs" else package
    return f"crate::{module}::{name}"


def child(shape: Shape) -> Shape:
    if shape.element is None:
        raise ValueError(f"{shape}: expected an element type")
    return shape.element


def rust_type(shape: Shape, interfaces: set[tuple[str, str]]) -> str:
    if shape.kind == "basic":
        return PRIMITIVES[shape.name]
    if shape.kind in ("parameter", "empty"):
        return shape.name if shape.kind == "parameter" else "()"
    if shape.kind == "named":
        return named_type(shape, interfaces)
    if shape.kind == "pointer":
        return f"Option<Box<{rust_type(child(shape), interfaces)}>>"
    if shape.kind == "slice":
        if child(shape).underlying in ("byte", "uint8"):
            return f"crate::json::ByteSlice<{rust_type(child(shape), interfaces)}>"
        return (
            "crate::json::Bytes"
            if child(shape).name in ("byte", "uint8")
            else f"Option<Vec<{rust_type(child(shape), interfaces)}>>"
        )
    if shape.kind == "array":
        return f"[{rust_type(child(shape), interfaces)}; {shape.length}]"
    if shape.kind == "map" and shape.key is not None:
        return f"Option<std::collections::BTreeMap<{rust_type(shape.key, interfaces)}, {rust_type(child(shape), interfaces)}>>"
    if shape.kind == "struct" and not shape.fields:
        return "crate::json::EmptyObject"
    raise ValueError(f"{shape}: expected a supported public value type")


def named_type(shape: Shape, interfaces: set[tuple[str, str]]) -> str:
    if shape.package in ("json", "jsontext"):
        return "crate::json::JsonValue"
    if shape.package == "time" and shape.name == "Time":
        return "crate::json::Timestamp"
    if shape.name == "error" and not shape.package:
        return "crate::Error"
    result = rust_path(shape.package, shape.name)
    if shape.parameters:
        result += (
            "<"
            + ", ".join(rust_type(item, interfaces) for item in shape.parameters)
            + ">"
        )
    return (
        f"std::sync::Arc<dyn {result}>"
        if (shape.package, shape.name) in interfaces
        else result
    )
