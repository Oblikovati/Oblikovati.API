// SPDX-License-Identifier: Apache-2.0

package wire

// Body point/ray/validity query DTOs (M07-F07, Oblikovati/Oblikovati#630).
// Entity kinds spell "vertex" | "edge" | "face"; containment verdicts use the
// [oblikovati.org/api/types.Containment] spellings.

// LocateUsingPointArgs is the request of [MethodBodyLocateUsingPoint].
type LocateUsingPointArgs struct {
	BodyIndex int       `json:"bodyIndex"`
	Point     []float64 `json:"point"`
	// EntityKind filters the search ("vertex"/"edge"/"face"; empty = any).
	EntityKind string `json:"entityKind,omitempty"`
	// ProximityTolerance bounds the accepted distance (cm).
	ProximityTolerance float64 `json:"proximityTolerance"`
}

// LocatedEntityInfo describes one located/hit entity.
type LocatedEntityInfo struct {
	Kind string `json:"kind"`
	// Key is the persistent reference key; TransientKey the session id.
	Key          string `json:"key"`
	TransientKey uint64 `json:"transientKey"`
	// Point is the closest/hit point on the entity; Distance from the query
	// point (or along the ray for [MethodBodyFindUsingRay]).
	Point    []float64 `json:"point,omitempty"`
	Distance float64   `json:"distance"`
}

// LocateUsingPointResult is the response of [MethodBodyLocateUsingPoint].
type LocateUsingPointResult struct {
	Found  bool              `json:"found"`
	Entity LocatedEntityInfo `json:"entity,omitempty"`
}

// FindUsingRayArgs is the request of [MethodBodyFindUsingRay].
type FindUsingRayArgs struct {
	BodyIndex int       `json:"bodyIndex"`
	Origin    []float64 `json:"origin"`
	Direction []float64 `json:"direction"`
	// Radius widens the ray into a pick cylinder for edges and vertices.
	Radius        float64 `json:"radius,omitempty"`
	FindFirstOnly bool    `json:"findFirstOnly,omitempty"`
}

// FindUsingRayResult is the response of [MethodBodyFindUsingRay], nearest
// first.
type FindUsingRayResult struct {
	Hits []LocatedEntityInfo `json:"hits,omitempty"`
}

// IsPointInsideArgs is the request of [MethodBodyIsPointInside]. ShellIndex
// (when non-nil) classifies against one shell's bounded region instead of the
// whole body's material.
type IsPointInsideArgs struct {
	BodyIndex  int       `json:"bodyIndex"`
	ShellIndex *int      `json:"shellIndex,omitempty"`
	Point      []float64 `json:"point"`
	// OnTolerance is the distance treated as ON the boundary (0 → 1e-6 cm).
	OnTolerance float64 `json:"onTolerance,omitempty"`
}

// IsPointInsideResult is the response of [MethodBodyIsPointInside].
type IsPointInsideResult struct {
	Containment string `json:"containment"`
}

// ConvexityEdgesArgs is the request of [MethodBodyConvexityEdges]; Collection
// is an [oblikovati.org/api/types.EdgeCollectionKind] wire spelling.
type ConvexityEdgesArgs struct {
	BodyIndex  int    `json:"bodyIndex"`
	Collection string `json:"collection"`
}

// ConvexityEdgesResult is the response of [MethodBodyConvexityEdges].
type ConvexityEdgesResult struct {
	Edges []TopologyRef `json:"edges,omitempty"`
}

// MinimumDistanceArgs is the request of [MethodBodyMinimumDistance]: the closest
// approach between the body and a transient probe polyline (e.g. a CAM travel
// path) — the out-of-process projection of Inventor's
// MeasureTools.GetMinimumDistance for a transient operand. Points is a flat
// [x,y,z, x,y,z, ...] list in database units (cm); consecutive pairs are the
// probe's segments, and a lone point measures that point to the body. Radius
// (cm) widens the probe into a swept-cylinder cross-section (the tool),
// subtracted from the raw distance and clamped at 0 — 0 leaves the probe a bare
// polyline.
type MinimumDistanceArgs struct {
	BodyIndex int       `json:"bodyIndex"`
	Points    []float64 `json:"points"`
	Radius    float64   `json:"radius,omitempty"`
}

// MinimumDistanceResult is the response of [MethodBodyMinimumDistance]: the
// minimum distance in database units (cm), 0 when the probe (after Radius)
// touches or enters the body's material.
type MinimumDistanceResult struct {
	Distance float64 `json:"distance"`
}

// ValidateBodyArgs is the request of [MethodBodyValidate]. CheckLevel 1 runs
// the topology checks (manifold/orientation/closure); 2 adds the face
// self-intersection scan. 0 means 1.
type ValidateBodyArgs struct {
	BodyIndex  int `json:"bodyIndex"`
	CheckLevel int `json:"checkLevel,omitempty"`
}

// ProblemEntityInfo is one offending entity of a failed validity check.
type ProblemEntityInfo struct {
	Kind         string `json:"kind"`
	Key          string `json:"key"`
	TransientKey uint64 `json:"transientKey"`
	Issue        string `json:"issue"`
}

// ValidateBodyResult is the response of [MethodBodyValidate].
type ValidateBodyResult struct {
	Valid    bool                `json:"valid"`
	Problems []ProblemEntityInfo `json:"problems,omitempty"`
}

// BodyRangeBoxArgs is the request of [MethodBodyRangeBox]: the topology box
// by default, the tessellation-tight box with Precise, the minimal oriented
// box with Oriented.
type BodyRangeBoxArgs struct {
	BodyIndex int  `json:"bodyIndex"`
	Precise   bool `json:"precise,omitempty"`
	Oriented  bool `json:"oriented,omitempty"`
}

// BodyRangeBoxResult is the response of [MethodBodyRangeBox]: Min/Max for the
// axis-aligned forms; Corner plus the three edge vectors for the oriented one.
type BodyRangeBoxResult struct {
	Min            []float64 `json:"min,omitempty"`
	Max            []float64 `json:"max,omitempty"`
	Corner         []float64 `json:"corner,omitempty"`
	DirectionOne   []float64 `json:"directionOne,omitempty"`
	DirectionTwo   []float64 `json:"directionTwo,omitempty"`
	DirectionThree []float64 `json:"directionThree,omitempty"`
}

// BindTransientKeyArgs is the request of [MethodBodyBindTransientKey].
type BindTransientKeyArgs struct {
	BodyIndex    int    `json:"bodyIndex"`
	TransientKey uint64 `json:"transientKey"`
}

// BindTransientKeyResult is the response of [MethodBodyBindTransientKey].
type BindTransientKeyResult struct {
	Found bool   `json:"found"`
	Kind  string `json:"kind,omitempty"`
	Key   string `json:"key,omitempty"`
}
