// SPDX-License-Identifier: Apache-2.0

package featureargs

// The loft feature and its nested wire helpers (#1709 promotes the composite loft kind from the
// host's opaque json.RawMessage path to a typed struct an add-in can build with compile checks).

const KindLoft = "loft"

// LoftSection is one cross-section the loft blends through: a sketch profile, an apex point
// ([x,y] on the sketch plane), or a body-face section (FaceRef).
type LoftSection struct {
	SketchIndex  int       `json:"sketchIndex"`
	ProfileIndex int       `json:"profileIndex"`
	Point        []float64 `json:"point,omitempty"`   // [x,y] on the sketch plane → an apex (point) section
	FaceRef      string    `json:"faceRef,omitempty"` // a body-face reference key → a face section (Tangent/Smooth)
}

// LoftEnd is a loft end-section condition: how the surface leaves the first (or arrives at the
// last) section. "angle"/"direction" tilt the takeoff by Angle, weighted by Impact.
type LoftEnd struct {
	Condition string  `json:"condition,omitempty"`
	Angle     string  `json:"angle,omitempty"`
	Impact    float64 `json:"impact,omitempty"`
	Reversed  bool    `json:"reversed,omitempty"`
}

// LoftRail identifies a loft guide polyline (rail, centerline, or map curve): either an open path
// in a sketch, or an explicit model-space [x,y,z] point list (Points overrides the path).
type LoftRail struct {
	PathSketchIndex int         `json:"pathSketchIndex"`
	PathIndex       int         `json:"pathIndex"`
	Points          [][]float64 `json:"points,omitempty"`
}

// LoftAreaStop is one area-graph control point: at fraction t (0..1 along the loft) the
// cross-section area is scaled by Scale.
type LoftAreaStop struct {
	T     float64 `json:"t"`
	Scale float64 `json:"scale"`
}

// Loft blends two or more cross-sections into a solid, optionally guided by rails, a centerline,
// map curves, an area graph, and end conditions (KindLoft).
type Loft struct {
	Sections   []LoftSection  `json:"sections"`
	Closed     bool           `json:"closed,omitempty"`
	Operation  string         `json:"operation,omitempty"`
	First      *LoftEnd       `json:"first,omitempty"`
	Last       *LoftEnd       `json:"last,omitempty"`
	Rails      []LoftRail     `json:"rails,omitempty"`
	Centerline *LoftRail      `json:"centerline,omitempty"`
	AreaGraph  []LoftAreaStop `json:"areaGraph,omitempty"`
	MapCurves  []LoftRail     `json:"mapCurves,omitempty"`
}

// Kind reports the feature kind Loft creates.
func (Loft) Kind() string { return KindLoft }

// loftArgs is the loft family's contribution to [All].
var loftArgs = []Arg{Loft{}}
