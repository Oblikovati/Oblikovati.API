// SPDX-License-Identifier: Apache-2.0

package wire

// CreateSketchArgs is the request of [MethodSketchCreate]: where to start the new sketch.
// By default it is an origin plane (Plane: XY | XZ | YZ; empty defaults to XY). Set
// WorkPlaneIndex to sketch on a user work plane instead (its index in list_work_planes) —
// the way to sketch on a plane built on a feature-created face, so later features reference
// earlier geometry. WorkPlaneIndex, when set, takes precedence over Plane.
//
// Orientation, when set, pins the sketch's in-plane axes deterministically instead of
// letting the host pick them — the reference CAD API's PlanarSketches.AddWithOrientation. Without it a
// sketch on a non-origin plane (e.g. one built through an axis at an angle) gets a
// host-chosen frame, so an add-in cannot know which sketch direction is "up" and cannot
// place parametric geometry reliably. See [SketchOrientation].
type CreateSketchArgs struct {
	Plane          string             `json:"plane,omitempty"`
	WorkPlaneIndex *int               `json:"workPlaneIndex,omitempty"`
	Orientation    *SketchOrientation `json:"orientation,omitempty"`
}

// SketchOrientation fixes a new sketch's in-plane coordinate frame to a reference axis —
// the reference CAD API's PlanarSketches.AddWithOrientation(AxisEntity, NaturalAxisDirection, AxisIsX,
// Origin). Axis is the reference whose direction, projected into the sketch plane, becomes
// one of the sketch axes; the other is the plane normal crossed with it (right-handed). The
// projection must be non-degenerate — the axis may not be perpendicular to the plane.
//
// Example: to sketch a meridian on a plane built through the Z axis at an angle, set
// Axis:"origin/axis/z", AxisIsX:false — the sketch's Y then runs along +Z (axial) and X
// runs radially, so the same (radius, axial) profile drawn on XZ works unchanged.
type SketchOrientation struct {
	// Axis is the reference axis whose in-plane projection fixes a sketch axis: an origin
	// axis constant (types.WorkRefZAxis …), a work-axis ref, or a linear-edge ref.
	Axis string `json:"axis"`
	// AxisIsX selects which sketch axis the projected reference becomes: X (true) or Y
	// (false). The reference CAD API's AxisIsX.
	AxisIsX bool `json:"axisIsX,omitempty"`
	// Reverse uses the reference direction reversed (the reference CAD API's NaturalAxisDirection=false).
	Reverse bool `json:"reverse,omitempty"`
	// Origin optionally sets the sketch origin [x,y,z] in cm; empty keeps the plane origin.
	Origin []float64 `json:"origin,omitempty"`
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
