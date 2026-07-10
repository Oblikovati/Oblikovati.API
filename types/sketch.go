// SPDX-License-Identifier: Apache-2.0

package types

// SketchEntityKind discriminates the kind of a 2D sketch entity in the wire
// protocol — the value of [oblikovati.org/api/wire.AddSketchEntityArgs.Kind]
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
	// SketchEntitySpline interpolates its points (a fit spline);
	// SketchEntityControlPointSpline approximates them (a control-point spline, the 2D
	// analog of Sketch3DEntityControlPointSpline). sketch.addEntity accepts either kind;
	// enumeration reports which mode a spline is in (Oblikovati/Oblikovati#150).
	SketchEntitySpline             SketchEntityKind = "spline"
	SketchEntityControlPointSpline SketchEntityKind = "controlPointSpline"
	SketchEntityRectangle          SketchEntityKind = "rectangle"
	SketchEntityPolygon            SketchEntityKind = "polygon"
	SketchEntitySlot               SketchEntityKind = "slot"
	SketchEntityFillet             SketchEntityKind = "fillet"
	SketchEntityChamfer            SketchEntityKind = "chamfer"
	SketchEntityImage              SketchEntityKind = "image"
	SketchEntityFillRegion         SketchEntityKind = "fillRegion"
	SketchEntityText               SketchEntityKind = "text"
	SketchEntityEquationCurve      SketchEntityKind = "equationCurve"
	SketchEntityFixedSpline        SketchEntityKind = "fixedSpline"
	SketchEntityOffsetSpline       SketchEntityKind = "offsetSpline"
	SketchEntityProjectedPoint     SketchEntityKind = "projectedPoint"
	SketchEntityProjectedCurve     SketchEntityKind = "projectedCurve"
	// SketchEntitySplineHandle is the tangency handle attached to one fit
	// point of an interpolation spline (M06-F11, Oblikovati/Oblikovati#626).
	SketchEntitySplineHandle SketchEntityKind = "splineHandle"
	SketchEntityUnknown      SketchEntityKind = "unknown"
)

// SketchLineType is a sketch's line-style override. The empty value means "inherit the
// document default". String values are frozen.
type SketchLineType string

const (
	SketchLineContinuous SketchLineType = "continuous"
	SketchLineDashed     SketchLineType = "dashed"
	SketchLineHidden     SketchLineType = "hidden"
	SketchLineCenter     SketchLineType = "center"
	SketchLinePhantom    SketchLineType = "phantom"
	// SketchLineCustom is a definition loaded from an industry-standard .lin
	// line-type file via sketch.setCustomLineType; it is set by that method,
	// not directly through sketch.setProperty.
	SketchLineCustom SketchLineType = "custom"
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
// constraint. Used by [oblikovati.org/api/wire.AddConstraintArgs.Kind] and by
// the enumerated constraint's Kind. String values are frozen.
type GeometricConstraintKind string

const (
	GeoConstraintCoincident    GeometricConstraintKind = "coincident"
	GeoConstraintPointOnLine   GeometricConstraintKind = "pointOnLine"
	GeoConstraintMidpoint      GeometricConstraintKind = "midpoint"
	GeoConstraintPointOnCircle GeometricConstraintKind = "pointOnCircle"
	GeoConstraintHorizontal    GeometricConstraintKind = "horizontal"
	GeoConstraintVertical      GeometricConstraintKind = "vertical"
	// GeoConstraintHorizontalAlign / VerticalAlign level TWO points (Inventor's
	// HorizontalAlign/VerticalAlign), distinct from the single-entity horizontal/
	// vertical that makes one line (or an ellipse axis) horizontal/vertical
	// (#1871): a single line ref under horizontal/vertical is the single-entity
	// form; two point refs are the align form.
	GeoConstraintHorizontalAlign GeometricConstraintKind = "horizontalAlign"
	GeoConstraintVerticalAlign   GeometricConstraintKind = "verticalAlign"
	GeoConstraintParallel        GeometricConstraintKind = "parallel"
	GeoConstraintPerpendicular   GeometricConstraintKind = "perpendicular"
	GeoConstraintCollinear       GeometricConstraintKind = "collinear"
	GeoConstraintConcentric      GeometricConstraintKind = "concentric"
	GeoConstraintEqualLength     GeometricConstraintKind = "equalLength"
	GeoConstraintEqualRadius     GeometricConstraintKind = "equalRadius"
	GeoConstraintTangent         GeometricConstraintKind = "tangent"
	GeoConstraintSymmetry        GeometricConstraintKind = "symmetry"
	GeoConstraintFix             GeometricConstraintKind = "fix"
	GeoConstraintSmooth          GeometricConstraintKind = "smooth"
	GeoConstraintGround          GeometricConstraintKind = "ground"
	GeoConstraintOffset          GeometricConstraintKind = "offset"
	GeoConstraintPattern         GeometricConstraintKind = "patternLink"
	// GeoConstraintTextBox is the auto-created anchor tying a text box to its
	// anchor geometry; it is never deletable on its own (M06-F11,
	// Oblikovati/Oblikovati#626).
	GeoConstraintTextBox GeometricConstraintKind = "textBox"
	// GeoConstraintCustom is an add-in-owned tag constraint: a named,
	// attribute-carrying record on sketch entities, not a solver callback
	// (M06-F11).
	GeoConstraintCustom  GeometricConstraintKind = "custom"
	GeoConstraintUnknown GeometricConstraintKind = "unknown"
)

// DimensionConstraintKind discriminates a sketch dimensional (driving/driven)
// constraint. Used by [oblikovati.org/api/wire.AddDimensionArgs.Kind] and by
// the enumerated dimension's Kind. String values are frozen.
type DimensionConstraintKind string

const (
	DimConstraintDistance        DimensionConstraintKind = "distance"
	DimConstraintAngle           DimensionConstraintKind = "angle"
	DimConstraintRadius          DimensionConstraintKind = "radius"
	DimConstraintDiameter        DimensionConstraintKind = "diameter"
	DimConstraintArcLength       DimensionConstraintKind = "arcLength"
	DimConstraintOffset          DimensionConstraintKind = "offsetDim"
	DimConstraintThreePointAngle DimensionConstraintKind = "threePointAngle"
	DimConstraintEllipseRadius   DimensionConstraintKind = "ellipseRadius"
	// DimConstraintTangentDistance dimensions the distance from a line to a circle/arc
	// measured to its tangent point — the near side by default, the far side with
	// AddDimensionArgs.FarSide (Oblikovati/Oblikovati#152).
	DimConstraintTangentDistance DimensionConstraintKind = "tangentDistance"
	DimConstraintUnknown         DimensionConstraintKind = "unknown"
)
