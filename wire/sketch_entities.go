// SPDX-License-Identifier: Apache-2.0

package wire

// AddSketchEntityArgs is the request of [MethodSketchAddEntity] — the discriminated
// entity constructor. Kind is the base entity
// ([oblikovati.org/api/types.SketchEntityKind]: line | point | circle | arc | …);
// Variant selects the overload within that kind ("centerRadius" | "threePoint" |
// "centerStartEnd"; empty ⇒ the kind's default). Points are the defining points, each
// [x,y] in sketch-plane database units (cm), in the constructor's expected order. Radius
// is a unit-bearing expression ("10 mm") for the center-radius circle. CCW orients a
// center-start-end arc. Construction marks the entity as reference geometry.
type AddSketchEntityArgs struct {
	SketchIndex int         `json:"sketchIndex"`
	Kind        string      `json:"kind"`
	Variant     string      `json:"variant,omitempty"`
	Points      [][]float64 `json:"points,omitempty"`
	// PointExprs is the parameter-expression form of Points (#189): each entry is the
	// defining point as ["x-expr","y-expr"], evaluated through the document's parameter
	// engine ("bore_r + 2 mm") and yielding a length in cm — so generated line/arc/point
	// geometry is parametric at construction, consistent with Radius/Width. When set it
	// supersedes Points; a unitless literal ("12.5") is read as cm. The two forms may be
	// mixed across a kind's points only via PointExprs (set every point as expressions).
	PointExprs   [][]string `json:"pointExprs,omitempty"`
	Radius       string     `json:"radius,omitempty"`
	CCW          bool       `json:"ccw,omitempty"`
	Construction bool       `json:"construction,omitempty"`

	// Conic fields (ellipse / ellipticalArc): Points[0] is the center, Axis is the
	// major-axis direction [x,y], MajorRadius/MinorRadius are unit-bearing lengths, and
	// StartAngle/EndAngle (unit-bearing angles, ellipticalArc only) bound the sweep.
	Axis        []float64 `json:"axis,omitempty"`
	MajorRadius string    `json:"majorRadius,omitempty"`
	MinorRadius string    `json:"minorRadius,omitempty"`
	StartAngle  string    `json:"startAngle,omitempty"`
	EndAngle    string    `json:"endAngle,omitempty"`

	// Closed marks a spline or a polyline a closed loop (spline and polyline kinds). A
	// closed polyline joins its last point back to its first, yielding one closed profile.
	Closed bool `json:"closed,omitempty"`

	// FitMethod is the interpolation parameterization for the spline kind
	// ([oblikovati.org/api/types.SplineFitMethod] wire spelling; empty ⇒
	// "smooth", the pre-field behavior — M06-F11, Oblikovati/Oblikovati#626).
	FitMethod string `json:"fitMethod,omitempty"`

	// Sides is the edge count for the polygon kind (≥ 3); Width is a unit-bearing slot
	// width. These belong to the composite kinds (rectangle/slot/polygon/polyline). The
	// polyline kind connects arbitrary Points with shared-endpoint lines (Closed ⇒ a
	// closed profile) — the way to author a non-regular outline (an L-bracket, a custom
	// extrusion section) over the API.
	Sides int    `json:"sides,omitempty"`
	Width string `json:"width,omitempty"`

	// EntityRefs are existing entity ids the kind operates on (the two line ids for the
	// fillet/chamfer corner blends; the parent spline id for offsetSpline). Radius is the
	// fillet radius / chamfer first distance / offset-spline distance; Distance2 is the
	// chamfer second distance (defaults to Radius when empty).
	EntityRefs []uint64 `json:"entityRefs,omitempty"`
	Distance2  string   `json:"distance2,omitempty"`

	// Equation-curve fields: x(t)/y(t) expressions over t ∈ [T0, T1].
	XExpr string  `json:"xExpr,omitempty"`
	YExpr string  `json:"yExpr,omitempty"`
	T0    float64 `json:"t0,omitempty"`
	T1    float64 `json:"t1,omitempty"`
}

// AddSketchEntityResult is the response of [MethodSketchAddEntity]: the primary entity's
// session id, its base kind, the session ids of its defining points, and — for composite
// kinds (rectangle/slot/polygon) that create several entities — every created entity id.
// When sketch inference is enabled (M06-F10, Oblikovati/Oblikovati#625),
// InferredConstraints reports the geometric constraints the engine auto-applied during
// creation and InferredPoints how defining points were snapped onto existing geometry.
type AddSketchEntityResult struct {
	EntityID  uint64   `json:"entityId"`
	Kind      string   `json:"kind"`
	PointIDs  []uint64 `json:"pointIds"`
	EntityIDs []uint64 `json:"entityIds,omitempty"`

	InferredConstraints []AppliedConstraintInference `json:"inferredConstraints,omitempty"`
	InferredPoints      []AppliedPointInference      `json:"inferredPoints,omitempty"`
}
