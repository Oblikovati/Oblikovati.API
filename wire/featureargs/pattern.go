// SPDX-License-Identifier: Apache-2.0

package featureargs

// The pattern feature kinds (#1709): rectangular grid, circular array, mirror, and sketch-driven.
// Source features are named as shown in model.tree; the layout is numeric (counts, steps, axis,
// plane normal). Counts have both a numeric and a parameter-expression form (the Expr form wins).

const (
	KindPatternRectangular  = "patternRectangular"
	KindPatternCircular     = "patternCircular"
	KindMirror              = "mirror"
	KindPatternSketchDriven = "patternSketchDriven"
)

// PatternBoundary is the optional pattern-clipping boundary (M20-F18): a closed loop of 3D points
// (cm) projected into the plane through PlaneOrigin with PlaneNormal; an occurrence is dropped when
// its Inclusion point falls outside the loop.
type PatternBoundary struct {
	PlaneOrigin []float64   `json:"planeOrigin,omitempty"`
	PlaneNormal []float64   `json:"planeNormal"`
	Polygon     [][]float64 `json:"polygon"`
	Inclusion   string      `json:"inclusion,omitempty"`
}

// PatternPlacement is the shared occurrence-placement option set of the rectangular and circular
// patterns (spacing/compute/orientation/positioning + an optional clipping boundary). It is
// embedded so its fields stay flat on the wire.
type PatternPlacement struct {
	SpacingType       string           `json:"spacingType,omitempty"`
	ComputeType       string           `json:"computeType,omitempty"`
	Orientation       string           `json:"orientation,omitempty"`
	PositioningMethod string           `json:"positioningMethod,omitempty"`
	Boundary          *PatternBoundary `json:"boundary,omitempty"`
	// SuppressedElements drops individual occurrences from the pattern by element index
	// (#1889). Element 0 is the seed — the source features' own material, which the recipe
	// already applied before the pattern ran — so it cannot be suppressed here; suppress the
	// source feature instead. Indices survive a count change, so an occurrence stays dropped
	// while the pattern is resized.
	SuppressedElements []int `json:"suppressedElements,omitempty"`
}

// PatternRectangular replicates features on a rectangular grid (KindPatternRectangular).
type PatternRectangular struct {
	SourceFeatures []string  `json:"sourceFeatures"`
	CountX         int       `json:"countX,omitempty"`
	CountY         int       `json:"countY,omitempty"`
	CountXExpr     string    `json:"countXExpr,omitempty"`
	CountYExpr     string    `json:"countYExpr,omitempty"`
	StepX          []float64 `json:"stepX,omitempty"`
	StepY          []float64 `json:"stepY,omitempty"`
	// MidPlaneX/MidPlaneY spread that direction's occurrences to BOTH sides of the seed
	// instead of running one way from it (#1889). The seed does not move. With an even
	// count the two sides cannot match, and the extra occurrence goes on the step's own
	// side — reverse the step to put it on the other.
	MidPlaneX bool `json:"midPlaneX,omitempty"`
	MidPlaneY bool `json:"midPlaneY,omitempty"`
	PatternPlacement
}

// Kind reports the feature kind PatternRectangular creates.
func (PatternRectangular) Kind() string { return KindPatternRectangular }

// PatternCircular replicates features in a circular array about an axis (KindPatternCircular).
type PatternCircular struct {
	SourceFeatures []string  `json:"sourceFeatures"`
	Count          int       `json:"count,omitempty"`
	CountExpr      string    `json:"countExpr,omitempty"`
	Angle          string    `json:"angle,omitempty"`
	AxisPoint      []float64 `json:"axisPoint,omitempty"`
	AxisDir        []float64 `json:"axisDir,omitempty"`
	// MidPlane sweeps the occurrences to both sides of the seed rather than all one way
	// round the axis (#1889); see [PatternRectangular.MidPlaneX] for the even-count rule.
	MidPlane bool `json:"midPlane,omitempty"`
	PatternPlacement
}

// Kind reports the feature kind PatternCircular creates.
func (PatternCircular) Kind() string { return KindPatternCircular }

// Mirror mirrors features across a plane (KindMirror).
type Mirror struct {
	SourceFeatures []string  `json:"sourceFeatures"`
	Origin         []float64 `json:"origin,omitempty"`
	Normal         []float64 `json:"normal,omitempty"`
}

// Kind reports the feature kind Mirror creates.
func (Mirror) Kind() string { return KindMirror }

// PatternSketchDriven replicates features at a set of placement points (KindPatternSketchDriven).
type PatternSketchDriven struct {
	SourceFeatures []string    `json:"sourceFeatures"`
	Points         [][]float64 `json:"points"`
}

// Kind reports the feature kind PatternSketchDriven creates.
func (PatternSketchDriven) Kind() string { return KindPatternSketchDriven }

// patternArgs is the pattern family's contribution to [All].
var patternArgs = []Arg{
	PatternRectangular{}, PatternCircular{}, Mirror{}, PatternSketchDriven{},
}
