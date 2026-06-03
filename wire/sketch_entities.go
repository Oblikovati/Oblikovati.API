// SPDX-License-Identifier: Apache-2.0

package wire

// AddSketchEntityArgs is the request of [MethodSketchAddEntity] — the discriminated
// entity constructor. Kind is the base entity
// ([github.com/Oblikovati/api/types.SketchEntityKind]: line | point | circle | arc | …);
// Variant selects the overload within that kind ("centerRadius" | "threePoint" |
// "centerStartEnd"; empty ⇒ the kind's default). Points are the defining points, each
// [x,y] in sketch-plane database units (cm), in the constructor's expected order. Radius
// is a unit-bearing expression ("10 mm") for the center-radius circle. CCW orients a
// center-start-end arc. Construction marks the entity as reference geometry.
type AddSketchEntityArgs struct {
	SketchIndex  int         `json:"sketchIndex"`
	Kind         string      `json:"kind"`
	Variant      string      `json:"variant,omitempty"`
	Points       [][]float64 `json:"points,omitempty"`
	Radius       string      `json:"radius,omitempty"`
	CCW          bool        `json:"ccw,omitempty"`
	Construction bool        `json:"construction,omitempty"`

	// Conic fields (ellipse / ellipticalArc): Points[0] is the center, Axis is the
	// major-axis direction [x,y], MajorRadius/MinorRadius are unit-bearing lengths, and
	// StartAngle/EndAngle (unit-bearing angles, ellipticalArc only) bound the sweep.
	Axis        []float64 `json:"axis,omitempty"`
	MajorRadius string    `json:"majorRadius,omitempty"`
	MinorRadius string    `json:"minorRadius,omitempty"`
	StartAngle  string    `json:"startAngle,omitempty"`
	EndAngle    string    `json:"endAngle,omitempty"`

	// Closed marks a spline a closed loop (spline kinds only).
	Closed bool `json:"closed,omitempty"`

	// Sides is the edge count for the polygon kind (≥ 3); Width is a unit-bearing slot
	// width. These belong to the composite kinds (rectangle/slot/polygon).
	Sides int    `json:"sides,omitempty"`
	Width string `json:"width,omitempty"`
}

// AddSketchEntityResult is the response of [MethodSketchAddEntity]: the primary entity's
// session id, its base kind, the session ids of its defining points, and — for composite
// kinds (rectangle/slot/polygon) that create several entities — every created entity id.
type AddSketchEntityResult struct {
	EntityID  uint64   `json:"entityId"`
	Kind      string   `json:"kind"`
	PointIDs  []uint64 `json:"pointIds"`
	EntityIDs []uint64 `json:"entityIds,omitempty"`
}
