// SPDX-License-Identifier: Apache-2.0

package types

// SketchEntityKind discriminates the kind of a 2D sketch entity in the wire
// protocol — the value of [github.com/Oblikovati/api/wire.AddSketchEntityArgs.Kind]
// and of each enumerated entity's Kind. The set grows as M21 adds entity families
// (conics/splines, slots, polygons, …); the string values are frozen.
type SketchEntityKind string

const (
	SketchEntityLine          SketchEntityKind = "line"
	SketchEntityPoint         SketchEntityKind = "point"
	SketchEntityCircle        SketchEntityKind = "circle"
	SketchEntityArc           SketchEntityKind = "arc"
	SketchEntityEllipse       SketchEntityKind = "ellipse"
	SketchEntityEllipticalArc SketchEntityKind = "ellipticalArc"
	SketchEntitySpline        SketchEntityKind = "spline"
	SketchEntityRectangle     SketchEntityKind = "rectangle"
	SketchEntityPolygon       SketchEntityKind = "polygon"
	SketchEntitySlot          SketchEntityKind = "slot"
	SketchEntityFillet        SketchEntityKind = "fillet"
	SketchEntityChamfer       SketchEntityKind = "chamfer"
	SketchEntityImage         SketchEntityKind = "image"
	SketchEntityFillRegion    SketchEntityKind = "fillRegion"
	SketchEntityText          SketchEntityKind = "text"
	SketchEntityUnknown       SketchEntityKind = "unknown"
)

// SketchLineType is a sketch's line-style override (Inventor LineTypeEnum). The empty
// value means "inherit the document default". String values are frozen.
type SketchLineType string

const (
	SketchLineContinuous SketchLineType = "continuous"
	SketchLineDashed     SketchLineType = "dashed"
	SketchLineHidden     SketchLineType = "hidden"
	SketchLineCenter     SketchLineType = "center"
	SketchLinePhantom    SketchLineType = "phantom"
)

// SketchPatternKind discriminates a sketch pattern. String values are frozen.
type SketchPatternKind string

const (
	SketchPatternRectangular SketchPatternKind = "rectangular"
	SketchPatternCircular    SketchPatternKind = "circular"
)

// ConstraintStatus is a sketch's (or entity's) constraint state, derived from the
// solver's DOF analysis. String values are frozen.
type ConstraintStatus string

const (
	ConstraintWell  ConstraintStatus = "well"  // fully constrained: 0 DOF, no redundancy
	ConstraintUnder ConstraintStatus = "under" // free degrees of freedom remain
	ConstraintOver  ConstraintStatus = "over"  // redundant or conflicting constraints
)

// GeometricConstraintKind discriminates a sketch geometric (non-dimensional)
// constraint. Used by [github.com/Oblikovati/api/wire.AddConstraintArgs.Kind] and by
// the enumerated constraint's Kind. String values are frozen.
type GeometricConstraintKind string

const (
	GeoConstraintCoincident    GeometricConstraintKind = "coincident"
	GeoConstraintPointOnLine   GeometricConstraintKind = "pointOnLine"
	GeoConstraintMidpoint      GeometricConstraintKind = "midpoint"
	GeoConstraintPointOnCircle GeometricConstraintKind = "pointOnCircle"
	GeoConstraintHorizontal    GeometricConstraintKind = "horizontal"
	GeoConstraintVertical      GeometricConstraintKind = "vertical"
	GeoConstraintParallel      GeometricConstraintKind = "parallel"
	GeoConstraintPerpendicular GeometricConstraintKind = "perpendicular"
	GeoConstraintCollinear     GeometricConstraintKind = "collinear"
	GeoConstraintConcentric    GeometricConstraintKind = "concentric"
	GeoConstraintEqualLength   GeometricConstraintKind = "equalLength"
	GeoConstraintEqualRadius   GeometricConstraintKind = "equalRadius"
	GeoConstraintTangent       GeometricConstraintKind = "tangent"
	GeoConstraintSymmetry      GeometricConstraintKind = "symmetry"
	GeoConstraintFix           GeometricConstraintKind = "fix"
	GeoConstraintSmooth        GeometricConstraintKind = "smooth"
	GeoConstraintGround        GeometricConstraintKind = "ground"
	GeoConstraintOffset        GeometricConstraintKind = "offset"
	GeoConstraintPattern       GeometricConstraintKind = "patternLink"
	GeoConstraintUnknown       GeometricConstraintKind = "unknown"
)

// DimensionConstraintKind discriminates a sketch dimensional (driving/driven)
// constraint. Used by [github.com/Oblikovati/api/wire.AddDimensionArgs.Kind] and by
// the enumerated dimension's Kind. String values are frozen.
type DimensionConstraintKind string

const (
	DimConstraintDistance  DimensionConstraintKind = "distance"
	DimConstraintAngle     DimensionConstraintKind = "angle"
	DimConstraintRadius    DimensionConstraintKind = "radius"
	DimConstraintDiameter  DimensionConstraintKind = "diameter"
	DimConstraintArcLength DimensionConstraintKind = "arcLength"
	DimConstraintUnknown   DimensionConstraintKind = "unknown"
)
