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
	// FormatSTEP reserves the ISO 10303 B-rep format (translator ships separately, M17-F02).
	FormatSTEP ExchangeFormat = "step"
)

// IsMesh reports whether the format is a faceted-mesh format (STL/OBJ/3MF) — the set
// the mesh-exchange translator handles. STEP is a B-rep format (a different translator).
func (f ExchangeFormat) IsMesh() bool {
	return f == FormatSTL || f == FormatOBJ || f == Format3MF
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
