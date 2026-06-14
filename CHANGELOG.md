# Changelog

All notable changes to `oblikovati.org/api` are documented here. The format is
based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this
project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
The git tag for a release is `v<version>` (the Go module convention; the `v` is
not part of the semver string). See [RELEASING.md](RELEASING.md).

## [Unreleased]

## [0.1.0] - 2026-06-14

First versioned release of the public API contract. While the major version is
`0` the surface is still in initial development and may change in any minor
release (semver §4); there is no backward-compatibility guarantee yet.

### Added

- `types` — value types, enums, and ids for the automation contract.
- `contract` — the in-process Go interfaces the host implements.
- `wire` — method-name constants and JSON request/response DTOs, with `mcp:`
  tool annotations the MCP bridge is generated from.
- `client` — a `Transport` plus typed client groups for add-ins (documents,
  sketches, features, work planes/points, assemblies, BOM, properties, …).
- `api.Version` — the module's semantic-version constant, the single source of
  truth a release is tagged from.

[Unreleased]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/Oblikovati/Oblikovati.API/releases/tag/v0.1.0
