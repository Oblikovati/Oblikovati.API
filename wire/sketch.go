// SPDX-License-Identifier: Apache-2.0

package wire

// CreateSketchArgs is the request of [MethodSketchCreate]: where to start the new sketch.
// By default it is an origin plane (Plane: XY | XZ | YZ; empty defaults to XY). Set
// WorkPlaneIndex to sketch on a user work plane instead (its index in list_work_planes) —
// the way to sketch on a plane built on a feature-created face, so later features reference
// earlier geometry. WorkPlaneIndex, when set, takes precedence over Plane.
type CreateSketchArgs struct {
	Plane          string `json:"plane,omitempty"`
	WorkPlaneIndex *int   `json:"workPlaneIndex,omitempty"`
}

// CreateSketchResult is the response of [MethodSketchCreate]: the new sketch's index
// (for sketch.rectangle / features.add) and the normalized plane label.
type CreateSketchResult struct {
	SketchIndex int    `json:"sketchIndex"`
	Plane       string `json:"plane"`
}

// SketchRectangleArgs is the request of [MethodSketchRectangle]: a closed rectangle
// from the sketch origin to (Width, Height), each a unit-bearing expression
// (e.g. "40 mm").
type SketchRectangleArgs struct {
	SketchIndex int    `json:"sketchIndex"`
	Width       string `json:"width"`
	Height      string `json:"height"`
}

// SketchRectangleResult is the response of [MethodSketchRectangle]: the sketch index
// and its resulting profile count.
type SketchRectangleResult struct {
	SketchIndex int `json:"sketchIndex"`
	Profiles    int `json:"profiles"`
}
