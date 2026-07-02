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

## Packages & consumption paths

Four packages: `types` (canonical enums/ids/values), `contract` (in-proc Go
interfaces), `wire` (method constants + JSON DTOs), `client` (a `Transport` +
typed client for add-ins). The full architecture story — why four, the two
consumption paths (in-proc `contract`+`types` vs out-of-runtime `wire`+`client`
over the C ABI), and worked examples — lives canonically in the
[API architecture wiki page](docs/wiki/Oblikovati-API-Architecture.md); this
README stays a pointer plus the quickstart below.

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
