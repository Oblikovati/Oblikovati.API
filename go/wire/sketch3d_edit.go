// SPDX-License-Identifier: Apache-2.0

package wire

// Transform3DArgs is the request of [MethodSketch3DTransform]: an editing operation over a
// selection of 3D entities. Op is "move" | "copy" (translate by Vector [x,y,z] cm),
// "rotate" (by the unit-bearing Angle about the axis through Center [x,y,z] in direction
// Axis [x,y,z], default +Z), or "delete" (remove the entities). Entities are the session
// ids of the selection.
type Transform3DArgs struct {
	SketchIndex int       `json:"sketchIndex"`
	Op          string    `json:"op"`
	Entities    []uint64  `json:"entities"`
	Vector      []float64 `json:"vector,omitempty"`
	Center      []float64 `json:"center,omitempty"`
	Axis        []float64 `json:"axis,omitempty"`
	Angle       string    `json:"angle,omitempty"`
}

// Transform3DResult is the response of [MethodSketch3DTransform]: the session ids created
// by a copy (empty otherwise) and the sketch's resulting entity count.
type Transform3DResult struct {
	Created     []uint64 `json:"created,omitempty"`
	EntityCount int      `json:"entityCount"`
}
