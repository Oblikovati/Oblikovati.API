# SPDX-License-Identifier: Apache-2.0
"""Stress both compiled C add-ins through a running, foreground MCPBridge host."""

import argparse
import json
from pathlib import Path
import time
from typing import cast
from urllib.request import Request, urlopen


class McpSession:
    def __init__(self, endpoint: str) -> None:
        self.endpoint = endpoint
        self.sequence = 0
        self.headers = {
            "Content-Type": "application/json",
            "Accept": "application/json, text/event-stream",
            "MCP-Protocol-Version": "2025-03-26",
        }
        self.request(
            "initialize",
            {
                "protocolVersion": "2025-03-26",
                "capabilities": {},
                "clientInfo": {"name": "api-live-parity", "version": "1.0.0"},
            },
        )
        self.exchange({"jsonrpc": "2.0", "method": "notifications/initialized"})

    def exchange(self, payload: dict[str, object]) -> str:
        request = Request(self.endpoint, json.dumps(payload).encode(), self.headers)
        with urlopen(request, timeout=20) as response:
            session = response.headers.get("Mcp-Session-Id")
            if session:
                self.headers["Mcp-Session-Id"] = session
            return response.read().decode()

    def request(self, method: str, params: dict[str, object]) -> dict[str, object]:
        self.sequence += 1
        response = self.exchange(
            {"jsonrpc": "2.0", "id": self.sequence, "method": method, "params": params}
        )
        content = (
            response
            if response.startswith("{")
            else next(
                line[6:] for line in response.splitlines() if line.startswith("data: ")
            )
        )
        reply = json.loads(content)
        if "error" in reply:
            raise AssertionError(
                f"{method}: expected a result, received {reply['error']}"
            )
        return cast(dict[str, object], reply["result"])

    def call(self, name: str, arguments: dict[str, object]) -> dict[str, object]:
        result = self.request("tools/call", {"name": name, "arguments": arguments})
        if result.get("isError"):
            raise AssertionError(
                f"{name}({arguments}): expected success, received {result}"
            )
        return result


def text_result(result: dict[str, object]) -> dict[str, object]:
    contents = cast(list[dict[str, object]], result["content"])
    return cast(
        dict[str, object],
        json.loads(
            str(next(item["text"] for item in contents if item["type"] == "text"))
        ),
    )


def probe_counts(log: Path, version: str) -> dict[str, int]:
    counts = {"go": 0, "rust": 0}
    for line in log.read_text().splitlines():
        if not line.startswith("{"):
            continue
        record = json.loads(line)
        if record.get("msg") == "api-version" and record.get("version") == version:
            counts[record["language"]] += 1
    return counts


def wait_for_probe(log: Path, version: str, language: str, expected: int) -> None:
    deadline = time.monotonic() + 15
    while time.monotonic() < deadline:
        if probe_counts(log, version)[language] >= expected:
            return
        time.sleep(0.025)
    raise AssertionError(
        f"{language}: expected {expected} successful {version} probes in {log}"
    )


def exercise_lifecycle(
    session: McpSession, log: Path, version: str, cycles: int
) -> None:
    for language in ("go", "rust"):
        wait_for_probe(log, version, language, 1)
        initial = probe_counts(log, version)[language]
        identity = {"id": f"com.oblikovati.example.{language}"}
        for cycle in range(cycles):
            session.call("addins_deactivate", identity)
            session.call("addins_activate", identity)
            wait_for_probe(log, version, language, initial + cycle + 1)
        session.call("addins_deactivate", identity)


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--endpoint", default="http://127.0.0.1:7815/mcp")
    parser.add_argument("--log", type=Path, required=True)
    parser.add_argument("--version", required=True)
    parser.add_argument("--screenshot", type=Path, required=True)
    parser.add_argument("--cycles", type=int, default=20)
    args = parser.parse_args()
    session = McpSession(args.endpoint)
    version = text_result(session.call("application_api_version", {}))
    assert version["version"] == args.version, version
    exercise_lifecycle(session, args.log, args.version, args.cycles)
    session.call("capture_viewport", {"path": str(args.screenshot.resolve())})
    assert args.screenshot.is_file(), args.screenshot
    print(
        json.dumps(
            {
                "version": args.version,
                "cycles_per_language": args.cycles,
                "probes": probe_counts(args.log, args.version),
                "screenshot": str(args.screenshot),
            }
        )
    )


if __name__ == "__main__":
    main()
