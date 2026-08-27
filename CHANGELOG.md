# Changelog

All notable changes to `oblikovati.org/api` are documented here. The format is
based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this
project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
The git tag for a release is `v<version>` (the Go module convention; the `v` is
not part of the semver string). See [RELEASING.md](RELEASING.md).

## [Unreleased]

## [0.153.1] - 2026-08-27

### Fixed

- Generic optionGroupField[V] helper to collapse client/options.go's 5 group getters
- Generic boundsCheckedAt[T]/indexOfFunc[T] helpers for client/ collection types
- Introduce Enumerable[T any] to collapse 16 duplicate Count()/Item(i) collection interfaces
- Widen enumName/enumFromName to comparable, adopt across all hand-rolled enum lookups

## [0.153.0] - 2026-08-27

### Added

- feat: convert Client.call to a generic method

## [0.152.0] - 2026-08-26

### Added

- feat: raise the module go directive to 1.27.0

### Fixed

- Ignore local .worktrees/ directory

## [0.151.1] - 2026-08-24

### Fixed

- Rename OpenPBRAppearance to Appearance (M46-F01)
- Delete legacy Appearance contract/wire/client (M46-F01)

## [0.151.0] - 2026-08-24

### Added

- feat(client): OpenPBRAppearances typed method group
- feat(wire): OpenPBRAppearance DTOs + openpbrAppearances.* methods
- feat(contract): OpenPBRAppearance interface
- feat(types): OpenPBR Surface v1.1.1 parameter/group/color types

### Fixed

- fix(wire): UpdateOpenPBRAppearanceArgs carries DisplayName

## [0.150.0] - 2026-08-22

### Added

- feat(wire): viewport.scroll + extended mouse buttons (#1822)

### Added

- feat(wire): viewport.scroll method + extended mouse-button spellings (back/forward/button5-7) for the full GLFW button set and a horizontal wheel axis (#1822)

## [0.149.0] - 2026-08-22

### Added

- feat(wire): sketch/profile reference for the split path tool (#2068)

## [0.148.0] - 2026-08-21

### Added

- feat(wire): flange edge-set collection for multi-edge flanges (#2071)

## [0.147.0] - 2026-08-19

### Added

- feat(api): persistent occurrence-pattern editing surface (#1976)

## [0.146.0] - 2026-08-18

### Added

- feat(client): expose drawing_delete_dimension as an MCP tool
- feat(drawing): retrieve model dimensions onto a view (#1991)
- feat(assembly): AssemblyOptions surface — get/set editing defaults (#1981)
- feat(assembly): occurrence DOF split — translation/rotation + axes (#1980)
- feat(assembly): virtual component API — geometry-free BOM lines (#1979)
- feat(assembly): BOM structure — Default/Varies values + per-occurrence set (#1978)
- feat(assembly): joint origin definition modes — infer/offset/betweenTwoFaces (#1973)
- feat(drawing): sheet authoring API — zoned borders, title-block corner, revision, formats (#1989)
- feat(drawing): view rotation + alignment API (#1988)
- feat(drawing): DrawingEdgeType + per-view tangent-edge display (#1984)
- feat(drawing): chamfer-note and bend-note annotation kinds (#1995)
- feat(drawing): report a hole note's tapped-hole count (#1995)
- feat(assembly): reference-vector axis on the angle constraint (#1972)
- feat(drawing): view crop — rectangular/circular fence + break-mark (#1987)
- feat(drawing): inspection dimensions — border shape, label, rate (#1996)
- feat(drawing): section view cut depth / reverse / partial-cut type (#1982)
- feat(types): dimension engineering tolerance surface (#1990)
- feat(drawing): view label surface — show-label/scale/name + position (#1983)
- feat(drawing): dimension text overrides + dual-unit (#1992, #1993)
- feat(types): drawing overlay view + foreshortened/symmetric/sum dimensions (#1986, #1994)
- feat(types): drawing view styles — hidden-line-removed, from-base, shaded-hidden (#1985)
- feat(assembly): joint gap/position + locked/protected surface (#1970, #1974)
- feat(assembly): occurrence visibility, opacity and state overrides (#1975, #1977)
- feat(types): mate solution undirected + no-solution (#1971)
- feat(types,wire): punch representation type and die-tool settings (#1968)
- feat(featureargs): corner chamfer variants and multi-radius round sets (#1967)
- feat(types,featureargs): lofted-flange output type and facet tolerance (#1966)
- feat(types,featureargs): the three rip types and the gap side (#1965)
- feat(types,featureargs): corner-seam finish styles and their measures (#1964)
- feat(featureargs): contour flange operation and bend radius (#1961)
- feat(featureargs): auto-miter the corner between two walls (#1961)
- feat(types,wire): bend transitions and per-feature bend options (#1959)
- feat(wire): flat-pattern punch results (#1963)
- feat(types,wire): corner relief, and the bend relief Inventor actually defaults to (#1960)
- feat(featureargs): partial-width flanges (#1958)
- feat(featureargs): where a flange's wall lands (#1957)
- feat(featureargs): the four hem shapes (#1956)
- feat(featureargs): the modify ops' missing options (#1864, #1891, #1892, #1894)
- feat(featureargs): emboss flavour, wrap-to-face and the wall taper (#1893)
- feat(featureargs): rib wall options, coil handedness and the flat spiral (#1882, #1883)
- feat(featureargs): mirror a whole body, and keep only the reflection (#1890)
- feat(featureargs): mid-plane patterns and per-element suppression (#1889)
- feat(featureargs): chamfer reference face and partial span (#1888)
- feat(featureargs): chordal width on the face fillet (#1887)
- feat(featureargs): hole placements, the seat/tap split, clearance and terminations
- feat(featureargs): revolve extents — to-face, from-to and to-next (#1860)

## [0.145.1] - 2026-08-05

### Fixed

- Fail the release when the version output cannot be written
- Add a theme token for the sketch dimension annotation
- Add synthesised viewport input to the API (click and key)

## [0.145.0] - 2026-08-04

### Added

- feat(wire): a view's projection on the camera frame (#2019 follow-up)
- feat(wire): revolve sweep direction (#2019)
- feat(wire): delete and move a sketch dimension (#2017)
- feat(wire): auto-project origin option in the sketch option group (#2016)
- feat(types): sketch entity format and Format-panel modes (#2015)
- feat(types): SelectionListButton ribbon style (#2015)
- feat(types): Color.IsOverride for optional colour overrides (#2015)
- feat(types): heads-up display options for in-canvas sketch input (#2014)

## [0.144.0] - 2026-07-11

### Added

- feat(surface): boundary-patch 3D edge loop, curvature, guide rails (#1867) (#274)

## [0.143.0] - 2026-07-11

### Added

- feat(surface): ruled-surface sweep direction, draft, flip (#1868) (#273)

## [0.142.5] - 2026-07-11

### Fixed

- featureargs: Sculpt per-surface direction / affected body (#1881) (#272)

## [0.142.4] - 2026-07-11

### Fixed

- featureargs: Trim cutting tools — work plane / surface body / sketch line (#1880) (#271)

## [0.142.3] - 2026-07-11

### Fixed

- featureargs: MidSurface min/max range, body selection, manual pairs (#1885) (#270)

## [0.142.2] - 2026-07-11

### Fixed

- featureargs: Extend multi-edge / to-plane / extension type (#1878) (#269)

## [0.142.1] - 2026-07-11

### Fixed

- featureargs: surface options for deleteFace / thicken / replaceFace (#1884, #1876, #1886) (#268)

## [0.142.0] - 2026-07-11

### Added

- feat(parameters): ISO limits-and-fits tolerance contract (#1848) (#267)

## [0.141.0] - 2026-07-11

### Added

- feat(parameters): introspection members BuiltIn/Renamed/DisabledActionTypes (#1853) (#266)

## [0.140.0] - 2026-07-11

### Added

- feat(sketch3d): wrap projection flattening frame (#1841 part 2) (#265)

## [0.139.0] - 2026-07-11

### Added

- feat(sketch3d): project-to-surface projection type + direction (#1841 part 1) (#264)

## [0.138.0] - 2026-07-10

### Added

- feat(sketch3d): OnFace constraint surface (#1839) (#263)

## [0.137.0] - 2026-07-10

### Added

- feat(sketch3d): intersection-curve work-plane operand (#1854) (#262)

## [0.136.0] - 2026-07-10

### Added

- feat(sketch3d): equation-curve coordinate systems (#1846) (#261)

## [0.135.0] - 2026-07-10

### Added

- feat(sketch): sketch-settings parity — grid/snap + constraint-display/relax (#1877) (#260)

## [0.134.0] - 2026-07-10

### Added

- feat(sketch): project cut edges + silhouette methods (#1873) (#259)

## [0.133.0] - 2026-07-10

### Added

- feat(client): add PanelTree/PanelTable builders (Part of #48)
- feat(wire): add TreeNode/TableRow DTOs + PanelTree/PanelTable spec fields (Part of #48)
- feat(types): add PanelTree/PanelTable control kinds (Part of #48)

## [0.132.0] - 2026-07-10

### Added

- feat(sketch): Batch B dimension surface — offsetSplineDim + driven/textPoint/linearDiameter (#257)

## [0.131.0] - 2026-07-10

### Added

- feat(sketch): align constraint kinds + ellipse-axis operand flags (#1871, #1879) (#256)

## [0.130.0] - 2026-07-10

### Added

- feat(sketch): pin a new sketch's in-plane frame to a reference axis

## [0.129.0] - 2026-07-10

### Added

- feat(sketch): revolve about the sketch centerline

## [0.128.0] - 2026-07-10

### Added

- feat(work-features): wire surface for the final Work Features gaps (#1842, #1849, #1857) (#253)

## [0.127.0] - 2026-07-09

### Added

- feat(work-features): ADR-0040 geometric edge ref + analytic-edge/line-by-entity axis + edge-midpoint point (#1840, #1842) (#252)

## [0.126.0] - 2026-07-09

### Added

- feat(work-features): surface-derived datums — revolved-face axis, sphere/torus centre point (#1840, #1842) (#251)

## [0.125.0] - 2026-07-09

### Added

- feat(work-planes): FlipNormal + AutoResize/Grounded/Size + tangent/bisector solution point (#1851, #1844) (#250)

## [0.124.0] - 2026-07-09

### Added

- feat(work-features): delete (tombstone + cascade/retain) and construction flag (#1855, #1849) (#249)

## [0.123.0] - 2026-07-09

### Added

- feat(work-features): relational axis/point constructors + workPoints.list (#1840, #1842) (#248)

## [0.122.0] - 2026-07-09

### Added

- feat(work-features): line-point plane + datum visibility toggle (#1843, #1856) (#247)

## [0.121.0] - 2026-07-09

### Added

- feat(parameters): type conversion (ConvertTo user/model/reference) (#1850) (#246)

## [0.120.0] - 2026-07-09

### Added

- feat(sketch): distance-dimension orientation (aligned/horizontal/vertical) (#1869) (#245)

## [0.119.0] - 2026-07-09

### Added

- feat(featureargs): shell direction (inside/outside/both) (#1864) (#244)

## [0.118.0] - 2026-07-09

### Added

- feat(featureargs): extrude from-to + distance-from-face targets (#1859) (#243)

## [0.117.0] - 2026-07-09

### Added

- feat(work-points): plane-axis-intersection constructor (#1842) (#242)

## [0.116.0] - 2026-07-09

### Added

- feat(featureargs): draft neutral plane, coil end conditions, hole drill point (#1866, #1883, #1863) (#241)

## [0.115.0] - 2026-07-09

### Added

- feat(parameters): non-numeric + model-kind creation and rename (#1845, #1847) (#240)

## [0.114.0] - 2026-07-08

### Added

- feat(featureargs): add Extrude.ToFaceGeom geometric to-face target (#239)

## [0.113.0] - 2026-07-08

### Added

- feat(featureargs): partial-length cosmetic thread (Length/Offset)

## [0.112.0] - 2026-07-08

### Added

- feat(featureargs): partial-length cosmetic thread — `Thread.Length` / `Thread.Offset` (Inventor's ThreadDepth/ThreadOffset), so a double-ended stud threads only its ends

## [0.111.0] - 2026-07-08

### Added

- feat: add visible flag to workPlanes.create for hidden construction datums

## [0.110.0] - 2026-07-08

### Added

- feat(featureargs): select extrude/revolve regions by interior seed point (#236)

## [0.109.0] - 2026-07-08

### Added

- feat(featureargs): geometric selectors for hole placement and dress-up edges/faces (#235)

## [0.108.0] - 2026-07-08

### Added

- feat: API gaps for the Inventor exporter (work axes, hole center, draft pull, sweep path) (#234)

## [0.107.0] - 2026-07-07

### Added

- feat(types/wire/client): point-cloud viewport display mode (#645)

### Fixed

- fix(client): remove duplicate generic call[Resp] left by the G2 merge

### Added

- feat(types/wire/client): point-cloud viewport display mode (#645). Adds
  `types.PointCloudDisplayMode` (the lowercase tokens `default` / `rgb` /
  `intensity`, with `IsValid`, `String`, and `AllPointCloudDisplayModes`), the
  `DisplayMode` field on `wire.PointCloudInfo`, the `wire.SetPointCloudDisplayModeArgs`
  DTO with `wire.MethodPointCloudsSetDisplayMode` (`pointClouds.setDisplayMode`),
  and the typed `client.PointClouds.SetDisplayMode` method. This lets an add-in
  read and switch how an attached scan is coloured (neutral marker, per-point scan
  RGB, or intensity greyscale) — the contract the host's RGB/intensity render work
  implements.
- docs(wire/model): document the `body/<url-base64(key)>` selection-reference form
  on `SelectionResult` and `SelectArgs` (#1492). A whole-body viewport pick now
  reports a non-empty, recompute-stable reference (the same key as `BodyInfo.Key`)
  in the `Refs` slot instead of an empty string, and it round-trips through
  `model.select`, so an add-in (e.g. an FEA study picking which bodies to analyse)
  can read a directly-selected body unambiguously.

## [0.106.0] - 2026-07-03

### Added

- feat(wire/featureargs): promote the remaining 63 feature-arg kinds (#1709)

### Added

- feat(wire/featureargs): promote the remaining 63 feature-arg kinds to typed
  structs (#1709) — the composite kinds (loft, sweep, the fillet/chamfer/draft/
  shell/lip dress-up family, patterns, model tolerance, move-body), the multi-kind
  freeform primitives, the args-less kinds (hull, sheet-metal unfold/refold as
  zero-field types), and the mechanical remainder (booleans, surfacing, direct
  edits, the sheet-metal family). Every registered feature kind now has a
  compile-checked `featureargs` argument struct with a `Kind()` method, so add-ins
  build feature arguments against a typed shape instead of hand-assembled JSON and
  the host decodes into the same type (the wire<->host parity guard's exception
  list is now empty).

## [0.105.0] - 2026-07-03

### Added

- feat(wire,client): object.renamed / property.changed metadata events (#1644)

### Fixed

- Restore S10 [Unreleased] changelog entry after develop merge

### Added

- feat(wire,client): metadata-mutation events so add-ins can observe rename and
  property changes — `types.ObjectKind`, `wire.ObjectRenamedEvent` /
  `PropertyChangedEvent` DTOs, their event-name constants, and
  `EventDispatcher.OnObjectRenamed` / `OnPropertyChanged` (audit S10, #1644)

## [0.104.1] - 2026-07-03

### Fixed

- I9: split fat contract interfaces into embedded capability families (semver-safe)

### Added

- feat(contract): split the fat `Document`, `DisplayOptions`, `FileDescriptor`,
  `TransientGeometry`, and `SurfaceEvaluator` interfaces into embedded capability
  families (e.g. `DocumentIdentity`/`DirtyState`/`DocumentLifecycle`). Each fat
  interface is now the union of its families, so existing consumers are unaffected
  while new host/add-in signatures can accept the narrowest capability they need —
  semver-safe, additive only (audit I9, #1632)

## [0.103.1] - 2026-07-03

### Fixed

- Correct the graphics object model's false "compile-time asserted" claim
- Promote per-kind feature-arg DTOs into api/wire/featureargs (+typed client ctor)

## [0.103.0] - 2026-07-03

### Added

- feat(wire): featureargs package — typed per-kind feature-creation arg structs
  (Extrude/Revolve/Rib/Emboss/Coil/Hole/Boss/Thread/Grill/Mesh/DirectEdit), each
  carrying its own Kind(), so add-ins build features with compile-checked types
  instead of raw JSON (ADR-0018; audit B5, #1616)
- feat(client): generic AddFeature[A featureargs.Arg] constructor that tags the
  wire envelope from the arg's Kind() — one typed constructor for every kind
- feat(wire,client): `object.renamed` / `property.changed` metadata-mutation events
  so add-ins observe body/sketch/feature/occurrence renames and suppression /
  sketch-setting changes without polling (#1644); adds `types.ObjectKind`,
  `wire.ObjectRenamedEvent` / `wire.PropertyChangedEvent`, and
  `EventDispatcher.OnObjectRenamed` / `OnPropertyChanged`

## [0.102.1] - 2026-07-02

### Fixed

- fix(client): move attribute mcp annotations to the wire-calling On variants

## [0.102.0] - 2026-06-30

### Added

- feat(client): TaskPanels group (Show/Close) for modal task panels
- feat(wire): modal TaskPanelSpec DTOs + show/close methods + closed event
- feat(client): PanelReferenceList builder + DockableWindows.SetReferences
- feat(wire): referenceList DTOs + setReferences method + referencesChanged event
- feat(types): add referenceList panel control kind (12)

### Fixed

- release: v0.101.0 — referenceList control + modal TaskPanelSpec
- fix(wire): drop out-of-scope client SetReferences (A3 owns it); gofmt alignment

## [0.100.1] - 2026-06-30

### Fixed

- Add dockableWindows.setValue (drive an add-in panel control)

## [0.99.0] - 2026-06-30

### Added

- feat(client): make sketch Symmetry constraint creatable (#1574)

## [0.98.0] - 2026-06-30

### Added

- feat(client): `Constrain.Symmetric(a, b, mirrorLine)` — make a 2D sketch
  Symmetry constraint creatable over the API (was enumerable-only) (#1574)

## [0.97.0] - 2026-06-30

### Added

- feat(browser): node icons and right-click context menus for add-in panes

## [0.96.0] - 2026-06-30

### Added

- feat(panel): CSS-grid-like nesting for dockable-window layouts

## [0.95.0] - 2026-06-29

### Added

- feat(contract): complete the Parameter interface — value/unit/tolerance/health (#1562)
- feat(wire): derived-parameter-table fidelity — references & provenance (#1561)

### Added

- feat(contract): `contract.Parameter` completeness — adds `NominalValue`, `UnitName`,
  `Tolerance`, `IsHealthy`, and `HealthReason` so an add-in reads a parameter's value,
  unit category, tolerance band, and evaluation health in-process, not just over the
  wire. New canonical `types.Tolerance` value type (aliased by the host's model/param);
  health is projected leanly (flag + reason) because the host's health enum is a
  source-internal concept (#1501) (M39-F06, #1562).
- feat(wire): derived-parameter-table fidelity — `DerivedParameterTableInfo` gains
  `References` (each derived parameter's link back to its source document + source
  parameter, the reference API's `DerivedParameter.ReferencedEntity`),
  `HasReferenceComponent`, and `ReferenceComponent` provenance; new
  `DerivedParameterReference` DTO (M39-F05, #1561).

### Changed

- docs(client): the Parameters, parameter-groups, and derived-parameter-table
  operation groups document their scope as the active document — a part OR an
  assembly — not just the active part (assemblies are first-class parameter
  holders; M39-F03, #1559). The wire surface is unchanged; only the documented
  scope broadened.

## [0.94.0] - 2026-06-29

### Added

- feat(exchange): add PDF as a sketch-import format

### Added

- feat(exchange): FormatPDF + IsSketch — import vector PDFs (CAD drawings plotted to PDF) into a sketch
- feat(wire): MethodImportPDF constant and ImportPDFArgs/ImportPDFResult DTOs
- feat(client): Client.ImportPDF typed method (mcp:tool import_pdf)

## [0.93.0] - 2026-06-28

### Added

- feat(wire): add EventCommandEnded constant

### Added

- feat(wire): EventCommandEnded constant for the command-lifecycle event pair (de-hardcodes "command.ended")

## [0.92.1] - 2026-06-27

### Fixed

- attributes: target selector — anchor attributes to a body/face by reference key

## [0.92.0] - 2026-06-27

### Added

- feat: add MaterialID to BodyInfo for per-body material read-back

## [0.91.0] - 2026-06-23

### Added

- feat(types): SurfaceContinuity (G0/G1/G2/G3) for surfacing operations

## [0.90.0] - 2026-06-23

### Added

- feat(types): FilletCrossSection (G2/conic) and LoftG3 continuity

## [0.89.0] - 2026-06-23

### Added

- feat(wire): expose document working scale in DocumentUnitsInfo

## [0.88.0] - 2026-06-22

### Added

- feat: add Sketch3DEnvironment ribbon environment

## [0.87.0] - 2026-06-21

### Added

- feat(wire): add brep.offsetFaces for 3D surface offset

### Added

- `brep.offsetFaces` (`wire.BrepOffsetFacesArgs` → `BrepHandleResult`,
  `client.TransientBRep().OffsetFaces`): offset a body's named faces by a distance along
  their surface normals, returning a transient body of the offset faces to sample. The
  out-of-process surface-offset primitive CAM needs for 3D surfacing tool compensation;
  parameters mirror Inventor's `FaceOffsetDefinition` (Distance / Reverse / Tolerance).

## [0.86.0] - 2026-06-21

### Added

- feat(wire): add-in-provided command icon SVG

### Added

- `wire.CreateCommandArgs.IconSVG` (and `wire.CommandInfo.IconSVG`): an add-in can
  now ship its own ribbon-button glyph as inline SVG markup instead of referencing a
  host-bundled icon key, so its buttons are not limited to the icons the host embeds.
  When set it takes precedence over `Icon`. (Oblikovati#671)

## [0.85.0] - 2026-06-21

### Added

- feat(wire): add body.minimumDistance for transient-probe distance

### Added

- feat(wire): add body.minimumDistance — closest approach between a body and a
  transient probe polyline (optional swept-tool radius), the out-of-process
  projection of MeasureTools.GetMinimumDistance for a transient operand

## [0.84.0] - 2026-06-20

### Added

- feat(wire): add pointClouds.nearestPoint to snap onto scan data

## [0.83.0] - 2026-06-20

### Added

- feat(wire): add pointClouds.fitPlane to derive a work plane from a scan

## [0.82.0] - 2026-06-20

### Added

- feat(types): add FormatLAS as a point-cloud exchange format

## [0.81.0] - 2026-06-20

### Added

- feat(types): add FormatE57 as a point-cloud exchange format

## [0.80.1] - 2026-06-20

### Fixed

- fix(api): classify FormatPLY as a point-cloud format, not a mesh (#645)

## [0.80.0] - 2026-06-20

### Added

- feat(api): add FormatPLY mesh exchange format (#645)

## [0.79.0] - 2026-06-20

### Added

- feat(api): point cloud crop-volume methods (#645)

## [0.78.0] - 2026-06-20

### Added

- feat(api): point cloud attach/query/place wire + client surface (#645)

## [0.77.0] - 2026-06-20

### Added

- feat(body): batched face surface evaluation over the wire

## [0.76.1] - 2026-06-20

### Fixed

- Add body rename/delete/physical-properties contract (#1078)

## [0.76.0] - 2026-06-20

### Added

- feat(documents): part end-of-part rollback marker (#141)

## [0.75.0] - 2026-06-20

### Added

- feat(documents): part end-of-part rollback marker — `document.getEndOfPart` /
  `document.setEndOfPart` over `contract.EndOfPart` (#141)

## [0.74.0] - 2026-06-20

### Added

- feat(documents): per-document sketch settings (#147)

## [0.73.0] - 2026-06-20

### Added

- feat(documents): per-document sketch settings — `document.getSketchSettings` /
  `document.setSketchSettings` over `types.SketchSettings` (constraint-inference toggles + family
  priority) (#147)

## [0.72.0] - 2026-06-19

### Added

- feat(features): expression-driven sketch coords + typed pattern features (#189)
- feat(events): feature lifecycle + sketch-edit push events (#148)

## [0.71.0] - 2026-06-19

### Added

- feat(ui): editable dockable-window panel controls
- feat(graphics): AddFloodPlot for FEA flood plots

### Added

- feat(ui): editable dockable-window panel controls — `types.PanelControlKind` gains
  `TextBox`, `ValueEditor`, `CheckBox`, `Dropdown`, `ComboBox`, `Slider` (modeled on
  Inventor's `MiniToolbarControlTypeEnum`); `wire.PanelControlSpec` gains `Value`,
  `Options`, `Min`/`Max`/`Step`; `wire.PanelValueChangedEvent` (`panel.valueChanged`)
  delivers a user's edit back to the add-in; `client` adds the `PanelTextBox`/
  `PanelValueEditor`/`PanelCheckBox`/`PanelDropdown`/`PanelComboBox`/`PanelSlider`
  constructors. Add-ins can now build editable forms in a dockable window.
- feat(graphics): `client.Graphics().AddFloodPlot` — an on-top, translucent scalar
  heatmap for drawing an FEA flood plot over the analyzed geometry.

## [0.70.0] - 2026-06-19

### Added

- feat(materials): add Magnetic property group to the material contract

### Added

- feat(materials): `types.Magnetic` group on `MaterialInfo` / `contract.Material` —
  magnetostatics constitutive data (μr, remanence Br, coercivity Hc, saturation Bsat,
  core loss) for soft-magnetic cores and permanent magnets; the zero value is a
  non-magnetic material (no migration). Unblocks the FEMM bridge reading magnetic
  block materials off host geometry (closes the bridge's documented GAP #3).

## [0.69.0] - 2026-06-19

### Added

- feat(model): highlight sets contract (#157)

## [0.68.0] - 2026-06-19

### Added

- feat(events): document lifecycle + model-changed push events (#148)

## [0.67.0] - 2026-06-19

### Added

- feat(sketch): sketch-to-sketch copy contract (#151)

## [0.66.0] - 2026-06-19

### Added

- feat(sketch): control-point spline kind + tangent-distance dimension (#150, #152)

## [0.65.0] - 2026-06-19

### Added

- feat(documents): update/rebuild + requiresUpdate (#139)

## [0.64.0] - 2026-06-19

### Added

- feat(body): name + visibility in BodyInfo and body.setVisible (#158)

## [0.63.0] - 2026-06-18

### Added

- feat(sketch): consumed/owned-by state + dependents enumeration (#154)

## [0.62.0] - 2026-06-18

### Added

- feat(features): re-pick geometric references in features.edit (#163)

## [0.61.0] - 2026-06-18

### Added

- feat(sketch3d): expose persistent reference keys for 3D sketches and entities (#153)

## [0.60.0] - 2026-06-18

### Added

- feat(sketch): persistent reference keys for sketches and entities (#153)

## [0.59.0] - 2026-06-18

### Added

- feat: analysis.modelHealth — aggregate feature health

## [0.58.0] - 2026-06-18

### Added

- feat(model): selection-mutation contract — select/deselect/clear (#157)

## [0.57.0] - 2026-06-18

### Added

- feat(events): parameters.changed push event contract (#148)

## [0.56.0] - 2026-06-18

### Added

- feat: loopLength measure type (face perimeter)

## [0.55.0] - 2026-06-18

### Added

- feat(attributes): add-in attribute-set contract (#155)

## [0.54.0] - 2026-06-18

### Added

- feat: angle measure type between entities

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

[Unreleased]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.153.1...HEAD
[0.153.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.153.0...v0.153.1
[0.153.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.152.0...v0.153.0
[0.152.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.151.1...v0.152.0
[0.151.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.151.0...v0.151.1
[0.151.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.150.0...v0.151.0
[0.150.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.149.0...v0.150.0
[0.149.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.148.0...v0.149.0
[0.148.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.147.0...v0.148.0
[0.147.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.146.0...v0.147.0
[0.146.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.145.1...v0.146.0
[0.145.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.145.0...v0.145.1
[0.145.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.144.0...v0.145.0
[0.144.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.143.0...v0.144.0
[0.143.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.142.5...v0.143.0
[0.142.5]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.142.4...v0.142.5
[0.142.4]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.142.3...v0.142.4
[0.142.3]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.142.2...v0.142.3
[0.142.2]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.142.1...v0.142.2
[0.142.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.142.0...v0.142.1
[0.142.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.141.0...v0.142.0
[0.141.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.140.0...v0.141.0
[0.140.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.139.0...v0.140.0
[0.139.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.138.0...v0.139.0
[0.138.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.137.0...v0.138.0
[0.137.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.136.0...v0.137.0
[0.136.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.135.0...v0.136.0
[0.135.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.134.0...v0.135.0
[0.134.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.133.0...v0.134.0
[0.133.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.132.0...v0.133.0
[0.132.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.131.0...v0.132.0
[0.131.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.130.0...v0.131.0
[0.130.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.129.0...v0.130.0
[0.129.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.128.0...v0.129.0
[0.128.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.127.0...v0.128.0
[0.127.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.126.0...v0.127.0
[0.126.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.125.0...v0.126.0
[0.125.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.124.0...v0.125.0
[0.124.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.123.0...v0.124.0
[0.123.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.122.0...v0.123.0
[0.122.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.121.0...v0.122.0
[0.121.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.120.0...v0.121.0
[0.120.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.119.0...v0.120.0
[0.119.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.118.0...v0.119.0
[0.118.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.117.0...v0.118.0
[0.117.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.116.0...v0.117.0
[0.116.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.115.0...v0.116.0
[0.115.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.114.0...v0.115.0
[0.114.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.113.0...v0.114.0
[0.113.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.112.0...v0.113.0
[0.111.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.110.0...v0.111.0
[0.110.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.109.0...v0.110.0
[0.109.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.108.0...v0.109.0
[0.108.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.107.0...v0.108.0
[0.107.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.106.0...v0.107.0
[0.106.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.105.0...v0.106.0
[0.105.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.104.1...v0.105.0
[0.104.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.104.0...v0.104.1
[0.103.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.103.0...v0.103.1
[0.102.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.102.0...v0.102.1
[0.102.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.101.0...v0.102.0
[0.100.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.100.0...v0.100.1
[0.99.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.98.0...v0.99.0
[0.97.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.96.0...v0.97.0
[0.96.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.95.0...v0.96.0
[0.95.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.94.0...v0.95.0
[0.94.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.93.0...v0.94.0
[0.93.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.92.1...v0.93.0
[0.92.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.92.0...v0.92.1
[0.92.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.91.0...v0.92.0
[0.91.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.90.0...v0.91.0
[0.90.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.89.0...v0.90.0
[0.89.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.88.0...v0.89.0
[0.88.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.87.0...v0.88.0
[0.87.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.86.0...v0.87.0
[0.86.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.85.0...v0.86.0
[0.85.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.84.0...v0.85.0
[0.84.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.83.0...v0.84.0
[0.83.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.82.0...v0.83.0
[0.82.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.81.0...v0.82.0
[0.81.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.80.1...v0.81.0
[0.80.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.80.0...v0.80.1
[0.80.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.79.0...v0.80.0
[0.79.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.78.0...v0.79.0
[0.78.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.77.0...v0.78.0
[0.77.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.76.1...v0.77.0
[0.76.1]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.76.0...v0.76.1
[0.76.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.75.0...v0.76.0
[0.74.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.73.0...v0.74.0
[0.72.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.71.0...v0.72.0
[0.71.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.70.0...v0.71.0
[0.70.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.69.0...v0.70.0
[0.69.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.68.0...v0.69.0
[0.68.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.67.0...v0.68.0
[0.67.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.66.0...v0.67.0
[0.66.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.65.0...v0.66.0
[0.65.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.64.0...v0.65.0
[0.64.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.63.0...v0.64.0
[0.63.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.62.0...v0.63.0
[0.62.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.61.0...v0.62.0
[0.61.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.60.0...v0.61.0
[0.60.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.59.0...v0.60.0
[0.59.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.58.0...v0.59.0
[0.58.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.57.0...v0.58.0
[0.57.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.56.0...v0.57.0
[0.56.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.55.0...v0.56.0
[0.55.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.54.0...v0.55.0
[0.54.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.53.0...v0.54.0
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
