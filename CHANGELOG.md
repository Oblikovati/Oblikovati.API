# Changelog

All notable changes to `oblikovati.org/api` are documented here. The format is
based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this
project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
The git tag for a release is `v<version>` (the Go module convention; the `v` is
not part of the semver string). See [RELEASING.md](RELEASING.md).

## [Unreleased]

## [0.2.0] - 2026-06-15

### Added

- feat(assembly): batch place-by-definition (place many copies in one call)
- feat(wire): add MaxHoleDiameter to ShrinkwrapCreateArgs
- feat(types): add GeometricCharacteristic enum for model GD&T tolerances
- feat(api): set_flexible_child for the M12-F06 independent solve
- feat(api): M12-F05 contact/interference + F06 flexible contracts
- feat(api): assembly representation contract (M12-F04)
- feat(api): assembly drive contract (M12-F03)
- feat(api): assembly joint event constants (M12-F02)
- feat(api): assembly joint contract surface (M12-F02)
- feat(api): assembly relationship event constants (M12-F01)
- feat(api): assembly constraint contract surface (M12-F01)
- feat(api): add-in/host API version handshake (major+minor)
- feat(keymap): contract for command alias & shortcut customization (M05-F17, #831)

### Fixed

- types: add ChamferType (M20-F03 #474)
- fix(api): mcp:input override for add_design_view_section (M12-F04)
- Add M20 F16–F20 API parity contracts
- fix(api): inline c.call in joint add methods for tool codegen (M12-F02)
- M26: add commandLine.submit to the public API
- fix(api): set_camera takes vector arrays over MCP (mcp:input override)

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

[Unreleased]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/Oblikovati/Oblikovati.API/releases/tag/v0.1.0
