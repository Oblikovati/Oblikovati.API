# Changelog

All notable changes to `oblikovati.org/api` are documented here. The format is
based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this
project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
The git tag for a release is `v<version>` (the Go module convention; the `v` is
not part of the semver string). See [RELEASING.md](RELEASING.md).

## [Unreleased]

## [0.17.0] - 2026-06-17

### Added

- feat(types,wire,client): drawing sheet contract (M14-F01)

## [0.16.0] - 2026-06-17

### Added

- feat(types): add FilletConcaveStrategy (outward fill / inward recess)

## [0.15.1] - 2026-06-17

### Fixed

- exchange: add DXF format, import/export methods and typed client

## [0.15.0] - 2026-06-17

### Added

- feat(types): add ChamferConcaveStrategy (outward fill / inward relief)

## [0.14.1] - 2026-06-16

### Fixed

- client: add import_dwg method (DWG → sketch import)
- types: add FormatDWG sketch exchange format

### Added

- feat(exchange): add DWG import surface — `FormatDWG`, the `import.dwg` wire method
  with `ImportDWGArgs`/`ImportDWGResult`, and the `Client.ImportDWG` typed call

## [0.14.0] - 2026-06-16

### Added

- feat(sheet-metal): add flat-pattern cosmetic-centerline contract

## [0.13.0] - 2026-06-16

### Added

- feat(sheet-metal): add flat-pattern bend-order annotation contract

## [0.12.0] - 2026-06-16

### Added

- feat(sheet-metal): add flat-pattern plates + settings contract

## [0.11.0] - 2026-06-16

### Added

- feat(sheet-metal): add flat-pattern folded<->flat entity-map contract

## [0.10.0] - 2026-06-16

### Added

- feat(sheet-metal): add flat-pattern edge/face classification contract

## [0.9.0] - 2026-06-16

### Added

- feat(sheet-metal): add flat-pattern orientation contract

## [0.8.0] - 2026-06-16

### Added

- feat(sheet-metal): add unfold (flat pattern) query

## [0.7.0] - 2026-06-16

### Added

- feat(sheet-metal): add bend-lineage query for the flat pattern

## [0.6.0] - 2026-06-16

### Added

- feat: sheet-metal rule/style contract (M13-F01)

## [0.5.0] - 2026-06-16

### Added

- feat(wire): add assemblyConstraints.snap (grip-snap constraint inference)

## [0.4.0] - 2026-06-16

### Added

- feat(graphics): client-graphics object model, full reference parity (M16-F05, #641)
- feat(views): add view enums, named views & CameraEvents API (M16-F03, #404/#409/#410)
- feat(styles): add style manager, color styles & libraries API (M16-F02, #403/#408)
- feat(display): add display options & document display settings API (M16-F07, #643)
- feat(types): add Color value object + ColorScheme API (M16-F06, #642)

## [0.3.0] - 2026-06-16

### Added

- feat(types): add FilletCornerType (miter/setback/round) corner-treatment enum

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

[Unreleased]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.17.0...HEAD
[0.17.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.16.0...v0.17.0
[0.16.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.15.1...v0.16.0
[0.15.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.15.0...v0.15.1
[0.15.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.14.1...v0.15.0
[0.14.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.14.0...v0.14.1
[0.14.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.13.0...v0.14.0
[0.13.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.12.0...v0.13.0
[0.12.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.11.0...v0.12.0
[0.11.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.10.0...v0.11.0
[0.10.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/Oblikovati/Oblikovati.API/releases/tag/v0.1.0
