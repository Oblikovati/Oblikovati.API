// SPDX-License-Identifier: Apache-2.0

package wire

// Feature-pattern kinds for [MethodFeaturesAdd] (#189): the part-feature replication
// operations, addressed by these kind strings. They replicate one or more existing
// features (referenced by tree name) rather than bodies — the canonical parametric
// workflow "model one slot, cut it, pattern it N = `slots`".
const (
	FeatureKindPatternCircular    = "patternCircular"
	FeatureKindPatternRectangular = "patternRectangular"
	FeatureKindMirror             = "mirror"
)

// CircularPatternFeatureArgs is the args object for a [FeatureKindPatternCircular]
// [MethodFeaturesAdd] call: replicate SourceFeatures (tree names) in a circular array
// about an axis. Count is the occurrence count; CountExpr is its parameter-expression
// form ("slots", "poles/2") evaluated through the document's parameter engine and, when
// set, supersedes Count — so the pattern count is parameter-driven (#189). Angle is the
// unit-bearing total sweep ("360 deg"). AxisPoint/AxisDir default to the origin and +Z.
type CircularPatternFeatureArgs struct {
	SourceFeatures []string  `json:"sourceFeatures"`
	Count          int       `json:"count,omitempty"`
	CountExpr      string    `json:"countExpr,omitempty"`
	Angle          string    `json:"angle,omitempty"`
	AxisPoint      []float64 `json:"axisPoint,omitempty"`
	AxisDir        []float64 `json:"axisDir,omitempty"`
}

// RectangularPatternFeatureArgs is the args object for a [FeatureKindPatternRectangular]
// [MethodFeaturesAdd] call: replicate SourceFeatures on a grid. CountX/CountY are the per-
// direction counts; CountXExpr/CountYExpr are their parameter-expression forms (when set,
// they supersede the numeric counts) (#189). StepX/StepY are the per-direction spacing
// vectors [x,y,z] in cm (StepY optional for a 1-D pattern).
type RectangularPatternFeatureArgs struct {
	SourceFeatures []string  `json:"sourceFeatures"`
	CountX         int       `json:"countX,omitempty"`
	CountY         int       `json:"countY,omitempty"`
	CountXExpr     string    `json:"countXExpr,omitempty"`
	CountYExpr     string    `json:"countYExpr,omitempty"`
	StepX          []float64 `json:"stepX,omitempty"`
	StepY          []float64 `json:"stepY,omitempty"`
}

// MirrorFeatureArgs is the args object for a [FeatureKindMirror] [MethodFeaturesAdd] call:
// mirror SourceFeatures across the plane through Origin (default origin) with the given
// Normal ([x,y,z], e.g. [1,0,0] for the YZ plane) (#189).
type MirrorFeatureArgs struct {
	SourceFeatures []string  `json:"sourceFeatures"`
	Origin         []float64 `json:"origin,omitempty"`
	Normal         []float64 `json:"normal"`
}
