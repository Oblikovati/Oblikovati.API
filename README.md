# Oblikovati API (`oblikovati.org/api`)

The **public automation contract** for Oblikovati, with Go and Rust bindings under
the **Apache-2.0** license. It is the single source of truth for the API: the GPL
application (the [`Oblikovati`](../Oblikovati) repo) implements it, and add-ins —
including closed-source ones — build against it. Everything an add-in needs ships
here: the Go and Rust contracts and the C ABI header
([`include/oblikovati_addin.h`](include/oblikovati_addin.h)). See
[ADR-0018](../Oblikovati/architecture/decisions/ADR-0018-apache-api-contract-module.md).

**Invariant:** this module must never import the implementation module
`oblikovati`. The dependency only flows the other way; CI
fails the build if it is ever violated.

## Packages & consumption paths

The Go module is in `go/` and the Rust crate is in `rust/`. The C header remains in
`include/` and is shared by both languages.

Four packages: `types` (canonical enums/ids/values), `contract` (in-proc Go
interfaces / Rust traits), `wire` (method constants + JSON DTOs), `client` (a `Caller` +
typed client for add-ins). The full architecture story — why four, the two
consumption paths (in-proc `contract`+`types` vs out-of-runtime `wire`+`client`
over the C ABI), and worked examples — lives canonically in the
[API architecture wiki page](docs/wiki/Oblikovati-API-Architecture.md); this
README stays a pointer plus the quickstart below.

## Adding to the API (the pattern every change follows)

1. **Contract here, in this repo:** enum/value type → `go/types`; Go interface → `go/contract`;
   method-name constant + DTOs → `go/wire`; typed method group → `go/client`.
   Run `make generate` to update Rust, the parity corpus, and the wiki examples.
2. **Implementation in the `Oblikovati` repo:** behavior in `kernel/model/app/head`; a
   compile-time assertion (`var _ contract.X = (*impl.X)(nil)`); the handler wired
   into `addin/router` keyed on the `wire` method constant.

Never re-declare a DTO or method string outside `wire`; never call the host from an
add-in with raw JSON — use `client`.

## Develop

```sh
cd go
GOWORK=off go build ./...    # build
GOWORK=off go vet ./...
GOWORK=off go test ./...     # tests all Go packages and examples
```

Both the `Oblikovati` application and the bridge add-in consume this module via a
`go.work` workspace over sibling checkouts, so changes here are picked up without
publishing. Every `.go` file carries an `SPDX-License-Identifier: Apache-2.0`
header, enforced by `scripts/check-spdx.py`.

The full repository check is `make gate`. It requires Go from `go/go.mod`, Rust
from `rust-toolchain.toml`, Python 3.11+, Node.js 22, a C compiler, and
`gomarkdoc` (`go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@v1.1.0`).
The C shared-library test runs on Linux and macOS. CI also runs Rust tests on Windows.

## Versioning

This module follows [Semantic Versioning 2.0.0](https://semver.org). The version
is the `api.Version` constant in [`go/version.go`](go/version.go); the git tag for a
release is `go/v<version>`. While the major version is `0` the API is in initial
development and may change in any minor release. Changes are recorded in
[CHANGELOG.md](CHANGELOG.md); the release process is in [RELEASING.md](RELEASING.md).

## Rust

The `oblikovati-api` crate provides `types`, `contract`, `wire`, `client`, and
`ffi` modules under the same Apache-2.0 license. It has no application dependency.
Rust clients send typed request structs through an injected transport.

```toml
[dependencies]
oblikovati-api = { path = "../Oblikovati.API/rust" }
```

```sh
cargo test --manifest-path rust/Cargo.toml
make gate
```

Use `make generate` after a Go contract change. The checked-in Rust declarations
are generated from the type-checked Go source. Building the Rust crate does not
require Go. The gate checks generation drift, tests both languages, and compares
Rust JSON against fixtures produced by the Go encoder. Rust client calls take
the canonical request DTO; Go convenience constructors and scalar helpers are language-specific.

For local Go consumers, change the workspace replacement to
`replace oblikovati.org/api => ../Oblikovati.API/go`. The Go import path is unchanged.
