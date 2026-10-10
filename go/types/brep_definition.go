// SPDX-License-Identifier: Apache-2.0

package types

// The bottom-up B-rep definition graph (M07-F05, Oblikovati/Oblikovati#628):
// the value-struct form of the reference SurfaceBodyDefinition family
// (lump → shell → face → loop → edge-use → edge → vertex, plus wire edges),
// shared by the in-proc contract and the wire DTOs. Geometry is carried as
// tagged analytic definitions; coordinates are [x, y, z] in database units.

// BrepCurveDef is one edge's model-space curve, tagged by Kind.
type BrepCurveDef struct {
	// Kind is "lineSegment", "arc" or "polyline".
	Kind string `json:"kind"`
	// Points are flattened xyz triplets: the two ends of a lineSegment, the
	// vertices of a polyline (unused for an arc).
	Points []float64 `json:"points,omitempty"`
	// Arc geometry: center, plane normal, the radius direction at angle 0.
	Center []float64 `json:"center,omitempty"`
	Normal []float64 `json:"normal,omitempty"`
	RefDir []float64 `json:"refDir,omitempty"`
	Radius float64   `json:"radius,omitempty"`
	// StartAngle/SweepAngle in radians; positive sweep is counterclockwise
	// about Normal.
	StartAngle float64 `json:"startAngle,omitempty"`
	SweepAngle float64 `json:"sweepAngle,omitempty"`
}

// BrepSurfaceDef is one face's surface, tagged by Kind.
type BrepSurfaceDef struct {
	// Kind is "plane", "cylinder", "cone", "sphere" or "torus".
	Kind string `json:"kind"`
	// Origin is the plane origin / cylinder base / cone apex / sphere or
	// torus center.
	Origin []float64 `json:"origin,omitempty"`
	// Normal is the plane normal; Axis the cylinder/cone/torus axis.
	Normal []float64 `json:"normal,omitempty"`
	Axis   []float64 `json:"axis,omitempty"`
	Radius float64   `json:"radius,omitempty"`
	// HalfAngle is the cone half-angle in radians.
	HalfAngle   float64 `json:"halfAngle,omitempty"`
	MajorRadius float64 `json:"majorRadius,omitempty"`
	MinorRadius float64 `json:"minorRadius,omitempty"`
}

// BrepVertexDef is one definition-graph vertex.
type BrepVertexDef struct {
	Position []float64 `json:"position"`
	// AssociativeID (non-zero) keys the vertex's reference identity so the
	// caller's picks survive recompute; 0 falls back to the definition index.
	AssociativeID int `json:"associativeId,omitempty"`
}

// BrepEdgeDef is one edge: curve plus endpoint vertex indices.
type BrepEdgeDef struct {
	Curve         BrepCurveDef `json:"curve"`
	StartVertex   int          `json:"startVertex"`
	EndVertex     int          `json:"endVertex"`
	AssociativeID int          `json:"associativeId,omitempty"`
}

// BrepEdgeUseDef is one oriented edge use within a loop.
type BrepEdgeUseDef struct {
	Edge int `json:"edge"`
	// Opposed traverses the edge against its natural start→end direction.
	Opposed bool `json:"opposed,omitempty"`
}

// BrepLoopDef is one boundary loop; a face's FIRST loop is its outer.
type BrepLoopDef struct {
	Uses []BrepEdgeUseDef `json:"uses"`
}

// BrepFaceDef is one face: surface, loops, material sense.
type BrepFaceDef struct {
	Surface BrepSurfaceDef `json:"surface"`
	// ParamReversed marks the material side opposing the surface normal.
	ParamReversed bool          `json:"paramReversed,omitempty"`
	Loops         []BrepLoopDef `json:"loops,omitempty"`
	AssociativeID int           `json:"associativeId,omitempty"`
}

// BrepShellDef groups face indices into one shell.
type BrepShellDef struct {
	Faces []int `json:"faces"`
}

// BrepWireDef is an ordered face-less edge chain.
type BrepWireDef struct {
	Edges []int `json:"edges"`
}

// BrepLumpDef groups shells and wires (one connected region).
type BrepLumpDef struct {
	Shells []BrepShellDef `json:"shells,omitempty"`
	Wires  []BrepWireDef  `json:"wires,omitempty"`
}

// BrepBodyDefinition is the whole bottom-up graph.
type BrepBodyDefinition struct {
	Solid    bool            `json:"solid,omitempty"`
	Vertices []BrepVertexDef `json:"vertices,omitempty"`
	Edges    []BrepEdgeDef   `json:"edges,omitempty"`
	Faces    []BrepFaceDef   `json:"faces,omitempty"`
	Lumps    []BrepLumpDef   `json:"lumps,omitempty"`
}

// BrepDefinitionIssue is one construction problem, addressed by its graph
// path (e.g. "edges[3]").
type BrepDefinitionIssue struct {
	Path    string `json:"path"`
	Problem string `json:"problem"`
}
