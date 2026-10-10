#!/usr/bin/env bash
# SPDX-License-Identifier: Apache-2.0
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"
mkdir -p build/abi
case "$(uname -s)" in
  Linux) EXT=so; LINK=(-ldl) ;;
  Darwin) EXT=dylib; LINK=() ;;
  *) echo "C ABI live test requires Linux or macOS, got $(uname -s)" >&2; exit 1 ;;
esac
cc -std=c11 -Wall -Wextra -Werror -Iinclude tests/c_abi_host.c "${LINK[@]}" -o build/abi/host
(cd go && GOWORK=off CGO_ENABLED=1 go build -buildmode=c-shared -o "../build/abi/go_addin.$EXT" ./examples/c_addin)
cargo build --manifest-path rust/Cargo.toml --locked --example c_addin
VERSION="$(cd go && GOWORK=off go run ./examples/version)"
build/abi/host "$ROOT/build/abi/go_addin.$EXT" "$VERSION"
build/abi/host "$ROOT/rust/target/debug/examples/libc_addin.$EXT" "$VERSION"
