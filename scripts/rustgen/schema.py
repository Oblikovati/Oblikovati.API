# SPDX-License-Identifier: Apache-2.0
"""Typed input model for the Go compiler's public API description."""

from __future__ import annotations

from dataclasses import dataclass, field
from typing import cast


@dataclass
class Shape:
    kind: str
    name: str = ""
    package: str = ""
    length: int = 0
    element: Shape | None = None
    key: Shape | None = None
    fields: list[Field] = field(default_factory=list)
    methods: list[Method] = field(default_factory=list)
    embedded: list[Shape] = field(default_factory=list)
    parameters: list[Shape] = field(default_factory=list)
    underlying: str = ""


@dataclass
class Field:
    name: str
    shape: Shape
    tag: str = ""
    embedded: bool = False
    doc: str = ""


@dataclass
class Method:
    name: str
    parameters: list[Field]
    results: list[Field]
    doc: str = ""


@dataclass
class Declaration:
    name: str
    package: str
    filename: str
    doc: str
    shape: Shape
    constant: bool
    value: str
    alias: bool = False
    feature_kind: str = ""


@dataclass
class Rpc:
    group: str
    name: str
    method: str
    request: Shape
    response: Shape
    doc: str


def shape_from_json(record: dict[str, object]) -> Shape:
    return Shape(
        str(record["Kind"]),
        str(record.get("Name", "")),
        str(record.get("Package", "")),
        int(cast(int, record.get("Length", 0))),
        (
            shape_from_json(cast(dict[str, object], record["Element"]))
            if "Element" in record
            else None
        ),
        (
            shape_from_json(cast(dict[str, object], record["Key"]))
            if "Key" in record
            else None
        ),
        [field_from_json(item) for item in records(record, "Fields")],
        [method_from_json(item) for item in records(record, "Methods")],
        [shape_from_json(item) for item in records(record, "Embedded")],
        [shape_from_json(item) for item in records(record, "Parameters")],
        str(record.get("Underlying", "")),
    )


def records(record: dict[str, object], key: str) -> list[dict[str, object]]:
    return cast(list[dict[str, object]], record.get(key) or [])


def field_from_json(record: dict[str, object]) -> Field:
    return Field(
        str(record["Name"]),
        shape_from_json(cast(dict[str, object], record["Type"])),
        str(record.get("Tag", "")),
        bool(record.get("Embedded", False)),
        str(record.get("Doc", "")),
    )


def method_from_json(record: dict[str, object]) -> Method:
    return Method(
        str(record["Name"]),
        [field_from_json(item) for item in records(record, "Parameters")],
        [field_from_json(item) for item in records(record, "Results")],
        str(record.get("Doc", "")),
    )


def declarations_from_json(record: dict[str, object]) -> list[Declaration]:
    return [
        Declaration(
            str(item["Name"]),
            str(item["Package"]),
            str(item["File"]),
            str(item["Doc"]),
            shape_from_json(cast(dict[str, object], item["Shape"])),
            bool(item["Constant"]),
            str(item.get("Value", "")),
            bool(item.get("Alias", False)),
            str(item.get("Kind", "")),
        )
        for item in records(record, "Declarations")
    ]


def calls_from_json(record: dict[str, object]) -> list[Rpc]:
    return [
        Rpc(
            str(item["Group"]),
            str(item["Name"]),
            str(item["Method"]),
            shape_from_json(cast(dict[str, object], item["Request"])),
            shape_from_json(cast(dict[str, object], item["Response"])),
            str(item["Doc"]),
        )
        for item in records(record, "Calls")
    ]
