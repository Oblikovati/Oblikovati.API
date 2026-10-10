// SPDX-License-Identifier: Apache-2.0

package wire

// TransformSketchArgs is the request of [MethodSketchTransform] — the discriminated edit
// operation on a selection of entities. Op is "move" | "rotate" | "copy" | "mirror":
//
//   - move:   translate the selection in place by Vector ([dx,dy] cm).
//   - copy:   duplicate the selection, offset by Vector; the copies are returned.
//   - rotate: rotate the selection in place about Center ([x,y] cm) by Angle (a
//     unit-bearing expression like "90 deg").
//   - mirror: reflect the selection across the line MirrorLine (an entity id); the
//     mirrored copies are returned.
type TransformSketchArgs struct {
	SketchIndex int       `json:"sketchIndex"`
	Op          string    `json:"op"`
	Entities    []uint64  `json:"entities"`
	Vector      []float64 `json:"vector,omitempty"`
	Center      []float64 `json:"center,omitempty"`
	Angle       string    `json:"angle,omitempty"`
	MirrorLine  uint64    `json:"mirrorLine,omitempty"`
}

// TransformSketchResult is the response of [MethodSketchTransform]: the ids of any newly
// created entities (the copies for "copy"/"mirror"; empty for the in-place "move"/"rotate").
type TransformSketchResult struct {
	Created []uint64 `json:"created,omitempty"`
}
