// SPDX-License-Identifier: Apache-2.0

package wire

// CreateSketchArgs is the request of [MethodSketchCreate]: the origin plane for the
// new sketch (XY | XZ | YZ; empty defaults to XY).
type CreateSketchArgs struct {
	Plane string `json:"plane"`
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
