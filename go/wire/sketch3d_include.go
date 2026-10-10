// SPDX-License-Identifier: Apache-2.0

package wire

// IncludeSketch3DArgs is the request of [MethodSketch3DInclude]: include referenced part
// geometry (edges/vertices, by reference key) into a 3D sketch as associative reference
// geometry. Refs are the reference keys of the part edges/vertices to include.
type IncludeSketch3DArgs struct {
	SketchIndex int      `json:"sketchIndex"`
	Refs        []string `json:"refs"`
}

// IncludeSketch3DResult is the response of [MethodSketch3DInclude]: the session ids of the
// created included entities, and whether every reference resolved (Healthy false ⇒ a
// reference was lost, its include skipped).
type IncludeSketch3DResult struct {
	Created []uint64 `json:"created,omitempty"`
	Healthy bool     `json:"healthy"`
}

// IncludeSketch2DArgs is the request of [MethodSketch3DIncludeSketch]: include geometry of
// an existing 2D sketch into a 3D sketch as associative reference geometry, lifted through
// the 2D sketch's host plane. SketchIndex is the target 3D sketch; SourceSketchIndex is
// the source 2D sketch; EntityIDs are the session ids of the 2D points/curves to include.
type IncludeSketch2DArgs struct {
	SketchIndex       int      `json:"sketchIndex"`
	SourceSketchIndex int      `json:"sourceSketchIndex"`
	EntityIDs         []uint64 `json:"entityIDs"`
}
