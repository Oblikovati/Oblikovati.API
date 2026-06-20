// SPDX-License-Identifier: Apache-2.0

package types

// ExchangeFormat names a foreign file format the host can import from or export to.
// It is a pure value enum (a stable wire string); the parser/encoder lives in the GPL
// host, never here. STEP is reserved (its translator ships separately, M17-F02).
//
// Example:
//
//	req := wire.ExportRequest{Path: "part.stl", Format: string(types.FormatSTL)}
type ExchangeFormat string

const (
	// FormatSTL is the STL triangle-soup format (binary on export; binary or ASCII on import).
	FormatSTL ExchangeFormat = "stl"
	// FormatOBJ is the Wavefront OBJ format (v/f records).
	FormatOBJ ExchangeFormat = "obj"
	// Format3MF is the 3D Manufacturing Format (a ZIP container around a 3D-model XML part).
	Format3MF ExchangeFormat = "3mf"
	// FormatPLY is the Stanford PLY format (ASCII or binary), the common export of 3D scanners
	// (structured-light / photogrammetry). It carries a vertex list (and, for a mesh, faces); the
	// host imports it as a POINT CLOUD — as-built reference scan data — not as a solid/mesh body,
	// so it is a point-cloud format (IsPointCloud), not a mesh format (#645).
	FormatPLY ExchangeFormat = "ply"
	// FormatE57 is the ASTM E2807 (E57) format, the vendor-neutral, structured export of most
	// laser/structured-light scanners (its points live in a bit-packed CompressedVector inside a
	// checksummed-page container). Like PLY the host imports it as a POINT CLOUD — as-built scan
	// data — so it is a point-cloud format (IsPointCloud), not a mesh format (#645).
	FormatE57 ExchangeFormat = "e57"
	// FormatSTEP reserves the ISO 10303 B-rep format (translator ships separately, M17-F02).
	FormatSTEP ExchangeFormat = "step"
	// FormatDWG is the AutoCAD DWG drawing format. Unlike the mesh/B-rep formats it
	// carries 2D/3D curve geometry, so it imports into a sketch (2D Sketch on a chosen
	// plane, or Sketch3D) rather than into surface bodies.
	FormatDWG ExchangeFormat = "dwg"
	// FormatDXF is the ASCII DXF drawing-exchange format — DWG's open, text sibling. Like
	// DWG it carries curve geometry and imports into a sketch; on export the version is
	// selectable (see DXFVersion).
	FormatDXF ExchangeFormat = "dxf"
)

// IsMesh reports whether the format is a faceted-mesh format (STL/OBJ/3MF) — the set
// the mesh-exchange translator handles. STEP is a B-rep format (a different translator).
func (f ExchangeFormat) IsMesh() bool {
	return f == FormatSTL || f == FormatOBJ || f == Format3MF
}

// IsPointCloud reports whether the format imports as point-cloud scan data (a referenced display
// object the design is modeled against) rather than as a body or sketch — the 3D-scanner formats
// (PLY and E57 now; LAS later). Such an import attaches a point cloud, not a solid (#645).
func (f ExchangeFormat) IsPointCloud() bool {
	return f == FormatPLY || f == FormatE57
}

// IsSketch reports whether the format imports as sketch curve geometry (DWG/DXF) rather
// than as surface bodies (mesh/STEP). Such an import targets a sketch and, when 2D,
// a chosen work plane.
func (f ExchangeFormat) IsSketch() bool {
	return f == FormatDWG || f == FormatDXF
}

// DXFVersion selects the generation an exported DXF targets. The geometry is identical
// across versions; the version sets $ACADVER and the surrounding section scaffolding. The
// zero value "" is treated as R2000.
//
// Example:
//
//	req := wire.ExportDXFArgs{Path: "part.dxf", Version: string(types.DXFR2018)}
type DXFVersion string

const (
	// DXFR2000 is the AutoCAD 2000 (AC1015) generation — broadest compatibility.
	DXFR2000 DXFVersion = "r2000"
	// DXFR2018 is the AutoCAD 2018 (AC1032) generation.
	DXFR2018 DXFVersion = "r2018"
)

// Normalized maps the zero value to R2000, leaving any explicit value unchanged.
func (v DXFVersion) Normalized() DXFVersion {
	if v == "" {
		return DXFR2000
	}
	return v
}

// MeshResolution selects the tessellation density of an exported mesh: coarser (low)
// to finer (high). It maps in the host to chord/angle tolerances — higher resolution
// yields more triangles for curved bodies (planar bodies are unaffected; they
// triangulate exactly). The zero value "" is treated as Medium.
//
// Example:
//
//	req := wire.ExportRequest{Path: "p.stl", Format: "stl", Resolution: string(types.ResolutionHigh)}
type MeshResolution string

const (
	// ResolutionLow is a coarse preview density.
	ResolutionLow MeshResolution = "low"
	// ResolutionMedium is the default display density.
	ResolutionMedium MeshResolution = "medium"
	// ResolutionHigh is a fine print/CAM density.
	ResolutionHigh MeshResolution = "high"
)

// Normalized maps the zero value to Medium, leaving any explicit value unchanged, so
// callers can treat an unset resolution as the sensible default.
func (r MeshResolution) Normalized() MeshResolution {
	if r == "" {
		return ResolutionMedium
	}
	return r
}
