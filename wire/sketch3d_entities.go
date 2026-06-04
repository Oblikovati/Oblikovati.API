// SPDX-License-Identifier: Apache-2.0

package wire

// AddSketch3DEntityArgs is the request of [MethodSketch3DAddEntity] — the discriminated
// 3D entity constructor. Kind is the base entity
// ([github.com/Oblikovati/api/types.Sketch3DEntityKind]: point | line | circle | arc | …).
// Points are the defining points, each [x,y,z] in model database units (cm), in the
// constructor's expected order (line: A,B; circle: center; arc: center,start,end). Radius
// is a unit-bearing expression ("10 mm") for circular kinds. Axis is the circle plane's
// normal [x,y,z] (defaults to +Z when empty). CCW orients an arc. Construction marks the
// entity as reference geometry.
type AddSketch3DEntityArgs struct {
	SketchIndex  int         `json:"sketchIndex"`
	Kind         string      `json:"kind"`
	Points       [][]float64 `json:"points,omitempty"`
	Radius       string      `json:"radius,omitempty"`
	Axis         []float64   `json:"axis,omitempty"`
	CCW          bool        `json:"ccw,omitempty"`
	Construction bool        `json:"construction,omitempty"`

	// Helix-only fields (kind "helical"). Mode selects which two of pitch/height/
	// revolutions define the helix ("pitchHeight" | "pitchRevolution" |
	// "revolutionHeight" | "spiral"); the third is derived. Pitch/Height are
	// unit-bearing lengths, Revolutions a turn count, Taper a per-revolution radial
	// growth length (spiral pitch). Clockwise sets the handedness; Points[0] is the
	// axis-base origin and Axis the winding axis (defaults to +Z).
	Mode        string  `json:"mode,omitempty"`
	Pitch       string  `json:"pitch,omitempty"`
	Height      string  `json:"height,omitempty"`
	Revolutions float64 `json:"revolutions,omitempty"`
	Taper       string  `json:"taper,omitempty"`
	Clockwise   bool    `json:"clockwise,omitempty"`

	// Conic fields (kind "ellipse"/"ellipticalArc"). Points[0] is the center, Axis the
	// plane normal (defaults to +Z), MajorAxis the in-plane major direction (defaults to
	// +X). MajorRadius/MinorRadius are unit-bearing lengths; StartAngle/SweepAngle are
	// unit-bearing angles bounding an elliptical arc.
	MajorAxis   []float64 `json:"majorAxis,omitempty"`
	MajorRadius string    `json:"majorRadius,omitempty"`
	MinorRadius string    `json:"minorRadius,omitempty"`
	StartAngle  string    `json:"startAngle,omitempty"`
	SweepAngle  string    `json:"sweepAngle,omitempty"`

	// Spline fields. For spline/controlPointSpline/fixedSpline, Points are the defining
	// points (each [x,y,z] in cm); Closed marks a closed loop. For equationCurve, XExpr/
	// YExpr/ZExpr are x(t)/y(t)/z(t) over [T0,T1].
	Closed bool    `json:"closed,omitempty"`
	XExpr  string  `json:"xExpr,omitempty"`
	YExpr  string  `json:"yExpr,omitempty"`
	ZExpr  string  `json:"zExpr,omitempty"`
	T0     float64 `json:"t0,omitempty"`
	T1     float64 `json:"t1,omitempty"`
}

// AddSketch3DEntityResult is the response of [MethodSketch3DAddEntity]: the created
// entity's session id, its base kind, and the session ids of its defining points.
type AddSketch3DEntityResult struct {
	EntityID uint64   `json:"entityId"`
	Kind     string   `json:"kind"`
	PointIDs []uint64 `json:"pointIds,omitempty"`
}
