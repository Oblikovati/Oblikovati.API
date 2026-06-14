# Oblikovati API (`oblikovati.org/api`)

The **public automation contract** for Oblikovati, as a standalone Go module under
the **Apache-2.0** license. It is the single source of truth for the API: the GPL
application (the [`Oblikovati`](../Oblikovati) repo) implements it, and add-ins —
including closed-source ones — build against it. Everything an add-in needs ships
here: the Go contract and the C ABI header
([`include/oblikovati_addin.h`](include/oblikovati_addin.h)). See
[ADR-0018](../Oblikovati/architecture/decisions/ADR-0018-apache-api-contract-module.md).

**Invariant:** this module must never import the implementation module
`oblikovati`. The dependency only flows the other way; CI
fails the build if it is ever violated.

## Packages

| Package | What it holds | Notes |
|---|---|---|
| `types` | enums, stable ids, value/option structs — pure data | the **canonical** definitions; `/source` aliases them (`type X = types.X`) |
| `contract` | in-proc Go interfaces (`Document`, `Parameter`, …) | `/source` types satisfy these via compile-time assertions |
| `wire` | method-name constants + JSON request/response DTOs | the host↔add-in contract; the router serves it |
| `client` | a `Transport` interface + a typed client over it | how out-of-runtime add-ins drive the host |

## The two consumption paths

A c-shared add-in runs its **own Go runtime** (ADR-0016), so a live Go interface
value can't cross the boundary. The contract therefore serves two audiences:

- **In-process / first-party** code uses `contract` interfaces + `types` directly.
- **Out-of-runtime / add-ins** use `wire` + `client` over a `Transport` (the add-in
  backs it with the host's C-ABI `ObkHostCall`; a gRPC transport behind the same
  `wire` surface is a deferred future, ADR-0003/0016).

## Adding to the API (the pattern every change follows)

1. **Contract here, in this repo:** enum/value type → `types`; Go interface → `contract`;
   method-name constant + DTOs → `wire`; typed method group → `client`.
2. **Implementation in the `Oblikovati` repo:** behavior in `kernel/model/app/head`; a
   compile-time assertion (`var _ contract.X = (*impl.X)(nil)`); the handler wired
   into `addin/router` keyed on the `wire` method constant.

Never re-declare a DTO or method string outside `wire`; never call the host from an
add-in with raw JSON — use `client`.

## Develop

```sh
go build ./...    # build
go vet ./...
go test ./...     # client has a fake-Transport test; types/wire/contract are data
```

Both the `Oblikovati` application and the bridge add-in consume this module via a
`go.work` workspace over sibling checkouts, so changes here are picked up without
publishing. Every `.go` file carries an `SPDX-License-Identifier: Apache-2.0`
header, enforced by `scripts/check-spdx.py`.

## Versioning

This module follows [Semantic Versioning 2.0.0](https://semver.org). The version
is the `api.Version` constant in [`version.go`](version.go); the git tag for a
release is `v<version>`. While the major version is `0` the API is in initial
development and may change in any minor release. Changes are recorded in
[CHANGELOG.md](CHANGELOG.md); the release process is in [RELEASING.md](RELEASING.md).
