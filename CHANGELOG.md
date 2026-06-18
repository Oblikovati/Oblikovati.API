# Changelog

All notable changes to `oblikovati.org/api` are documented here. The format is
based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this
project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
The git tag for a release is `v<version>` (the Go module convention; the `v` is
not part of the semver string). See [RELEASING.md](RELEASING.md).

## [Unreleased]

## [0.53.0] - 2026-06-18

### Added

- feat: minDistance measure type between two entities

## [0.52.1] - 2026-06-18

### Fixed

- fix: measure resolves raw reference keys, not hex

## [0.52.0] - 2026-06-18

### Added

- feat: measurement contract (M18-F01 PBI-164, #428)

## [0.51.0] - 2026-06-18

### Added

- feat: mass-properties inertia, principal axes & accuracy contract (M18-F01, #429)

## [0.50.0] - 2026-06-18

### Added

- feat: mass-properties analysis contract (M18-F01, #423)

## [0.49.0] - 2026-06-18

### Added

- feat: drawing hatch region contract (M14-F08, #638)

## [0.48.0] - 2026-06-18

### Added

- feat: drawing sketch contract (M14-F08, #638)

## [0.47.0] - 2026-06-18

### Added

- feat: hole-note format-string override contract (M14-F07, #637)

## [0.46.0] - 2026-06-18

### Added

- feat: hole-note quantity grouping contract (M14-F07, #637)

## [0.45.0] - 2026-06-18

### Added

- feat: hole feature-note annotation contract (M14-F07, #637)

## [0.44.0] - 2026-06-18

### Added

- feat: drawing note & custom table annotation contract (M14-F04 PBI-144, #391)

## [0.43.0] - 2026-06-18

### Added

- feat: revision table & revision tag annotation contract (M14-F04 PBI-144, #391)

## [0.42.0] - 2026-06-18

### Added

- feat: hole table annotation contract (M14-F04 PBI-144, #391)

## [0.41.0] - 2026-06-18

### Added

- feat: balloon annotation contract (M14-F04 PBI-143, #390)

## [0.40.0] - 2026-06-18

### Added

- feat: parts list annotation contract (M14-F04 PBI-143, #390)

## [0.39.0] - 2026-06-18

### Added

- feat: surface texture symbol annotation contract (M14-F03 PBI-142, #389)

## [0.38.0] - 2026-06-18

### Added

- feat: GD&T datum feature symbol annotation contract (M14-F03 PBI-142, #389)

## [0.37.0] - 2026-06-18

### Added

- feat: GD&T feature control frame annotation contract (M14-F03 PBI-142, #389)

## [0.36.0] - 2026-06-18

### Added

- feat: centerline drawing annotation contract (M14-F03 PBI-142, #389)

## [0.35.0] - 2026-06-18

### Added

- feat: centre-mark drawing annotation contract (M14-F03 PBI-142, #389)

## [0.34.0] - 2026-06-18

### Added

- feat: arc-length drawing dimensions contract (M14-F03 PBI-141, #388)

## [0.33.0] - 2026-06-18

### Added

- feat: ordinate drawing dimensions contract (M14-F03 PBI-141, #388)

## [0.32.0] - 2026-06-18

### Added

- feat(drawing): baseline & chain dimension-set contract (M14-F03 PBI-141, #388)

## [0.31.0] - 2026-06-18

### Added

- feat(transactions): history-browser contract — transaction.history + transaction.jumpTo

## [0.30.0] - 2026-06-18

### Added

- feat(drawing): angular dimension contract (M14-F03 PBI-141, #388)

## [0.29.0] - 2026-06-18

### Added

- feat(drawing): radius & diameter dimension contract (M14-F03 PBI-141, #388)

## [0.28.0] - 2026-06-17

### Added

- feat(drawing): linear dimension contract (M14-F03 PBI-141, #388)

## [0.27.0] - 2026-06-17

### Added

- feat(drawing): CoG marker & revision-cloud annotation contracts (M14-F02 #813)

## [0.26.0] - 2026-06-17

### Added

- feat(drawing): slice, breakout & draft view contracts (M14-F02 #812)

## [0.25.0] - 2026-06-17

### Added

- feat(drawing): break view contract (M14-F02 PBI-140)

## [0.24.0] - 2026-06-17

### Added

- feat(drawing): detail view contract (M14-F02 PBI-140)

## [0.23.0] - 2026-06-17

### Added

- feat(drawing): section view contract (M14-F02 PBI-140)

## [0.22.0] - 2026-06-17

### Added

- feat(units): public API for document units & unit/expression service (#146)

## [0.21.0] - 2026-06-17

### Added

- feat(drawing): auxiliary view contract + view-type/curve-kind discriminators (M14-F02 PBI-140)

## [0.20.0] - 2026-06-17

### Added

- feat(wire,client): drawing sheet DXF export contract (M14-F05)

## [0.19.0] - 2026-06-17

### Added

- feat(types,wire,client): drawing view contract (M14-F02)

## [0.18.0] - 2026-06-17

### Added

- feat(types,wire,client): drafting standard + drawing style contract (M14-F01)

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

[Unreleased]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.53.0...HEAD
[0.53.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.52.1...v0.53.0
[0.52.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.52.0...v0.52.1
[0.52.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.51.0...v0.52.0
[0.51.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.50.0...v0.51.0
[0.50.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.49.0...v0.50.0
[0.49.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.48.0...v0.49.0
[0.48.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.47.0...v0.48.0
[0.47.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.46.0...v0.47.0
[0.46.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.45.0...v0.46.0
[0.45.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.44.0...v0.45.0
[0.44.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.43.0...v0.44.0
[0.43.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.42.0...v0.43.0
[0.42.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.41.0...v0.42.0
[0.41.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.40.0...v0.41.0
[0.40.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.39.0...v0.40.0
[0.39.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.38.0...v0.39.0
[0.38.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.37.0...v0.38.0
[0.37.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.36.0...v0.37.0
[0.36.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.35.0...v0.36.0
[0.35.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.34.0...v0.35.0
[0.34.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.33.0...v0.34.0
[0.33.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.32.0...v0.33.0
[0.32.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.31.0...v0.32.0
[0.31.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.30.0...v0.31.0
[0.30.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.29.0...v0.30.0
[0.29.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.28.0...v0.29.0
[0.28.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.27.0...v0.28.0
[0.27.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.26.0...v0.27.0
[0.26.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.25.0...v0.26.0
[0.25.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.24.0...v0.25.0
[0.24.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.23.0...v0.24.0
[0.23.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.22.0...v0.23.0
[0.22.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.21.0...v0.22.0
[0.21.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.20.0...v0.21.0
[0.20.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.19.0...v0.20.0
[0.19.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.18.0...v0.19.0
[0.18.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.17.0...v0.18.0
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
