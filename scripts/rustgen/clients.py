# SPDX-License-Identifier: Apache-2.0
"""Generate a typed request method for every Go transport call."""

from names import rust_type, snake
from schema import Rpc


def render_group(
    group: str, calls: list[Rpc], interfaces: set[tuple[str, str]], first: bool = True
) -> str:
    methods = "\n".join(render_rpc(call, interfaces) for call in calls)
    if group == "Client":
        return f"impl<T: super::Caller> super::Client<T> {{\n{methods}\n}}\n"
    if not first:
        return f"impl<T: super::Caller> super::{group}<'_, T> {{\n{methods}\n}}\n"
    getter = snake(group)
    return (
        f"/// Typed `{group}` calls over the shared API transport.\n"
        f"pub struct {group}<'a, T: super::Caller>(pub(super) &'a super::Client<T>);\n"
        f"impl<T: super::Caller> super::Client<T> {{\n"
        f"    /// Access the `{group}` API. Example: `client.{getter}()`.\n"
        f"    pub fn {getter}(&self) -> {group}<'_, T> {{ {group}(self) }}\n}}\n"
        f"impl<T: super::Caller> {group}<'_, T> {{\n{methods}\n}}\n"
    )


def render_rpc(call: Rpc, interfaces: set[tuple[str, str]]) -> str:
    request = rust_type(call.request, interfaces)
    response = rust_type(call.response, interfaces)
    empty = call.request.kind == "empty"
    params = "" if empty else f", args: &{request}"
    receiver = "self" if call.group == "Client" else "self.0"
    invoke = "invoke_void" if call.response.kind == "empty" else "invoke"
    payload = "None::<&()>" if empty else "Some(args)"
    example = "client" if call.group == "Client" else f"client.{snake(call.group)}()"
    return (
        f"    /// Call `{call.method}` with the canonical request and response types.\n"
        f"    /// Example: `{example}.{snake(call.name)}({'&args' if not empty else ''})`.\n"
        f"    pub fn {snake(call.name)}(&self{params}) -> Result<{response}, crate::Error> {{\n"
        f"        {receiver}.{invoke}(crate::wire::{snake(call.method).upper()}, {payload})\n    }}\n"
    )
