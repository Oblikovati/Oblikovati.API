// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The transient-curve contracts (M01-F05, #602): ownerless immutable curve
// values created through [TransientGeometry]. Every curve carries its kind
// discriminator and can evaluate positions/tangents over its parameter domain;
// the member-level evaluator surface (curvature, arc length, strokes…) is
// reached via Evaluator() (M01-F06, #603). Angles are radians; lengths are
// database units (cm).

// Curve is the umbrella every 3D transient curve satisfies.
type Curve interface {
	// CurveType identifies the concrete kind (frozen reference values).
	CurveType() types.CurveType
	// GeometryForm reports whether the curve is exactly NURBS-representable.
	GeometryForm() types.CurveGeometryForm
	// Evaluate returns the position at parameter t.
	Evaluate(t float64) types.Point
	// Tangent returns the (unnormalized) first derivative at t.
	Tangent(t float64) types.Vector
	// Domain returns the parameter range (±Inf for unbounded curves).
	Domain() (lo, hi float64)
	// Evaluator returns the member-level query surface of this curve.
	Evaluator() CurveEvaluator
}

// Curve2d is the umbrella every 2D transient curve satisfies.
type Curve2d interface {
	CurveType() types.Curve2dType
	GeometryForm() types.CurveGeometryForm
	Evaluate(t float64) types.Point2d
	Tangent(t float64) types.Vector2d
	Domain() (lo, hi float64)
	// Evaluator returns the member-level query surface of this curve.
	Evaluator() Curve2dEvaluator
}

// Line is an unbounded 3D line.
type Line interface {
	Curve
	RootPoint() types.Point
	Direction() types.UnitVector
}

// LineSegment is a bounded 3D line span.
type LineSegment interface {
	Curve
	StartPoint() types.Point
	EndPoint() types.Point
}

// Circle is a full 3D circle (parameter 0..1 around the normal).
type Circle interface {
	Curve
	Center() types.Point
	Normal() types.UnitVector
	Radius() float64
}

// Arc3d is a 3D circular arc.
type Arc3d interface {
	Curve
	Center() types.Point
	Normal() types.UnitVector
	ReferenceVector() types.UnitVector
	Radius() float64
	StartAngle() float64
	SweepAngle() float64
}

// EllipseFull is a full 3D ellipse.
type EllipseFull interface {
	Curve
	Center() types.Point
	Normal() types.UnitVector
	MajorAxis() types.UnitVector
	MajorRadius() float64
	MinorRadius() float64
}

// EllipticalArc is a bounded span of a 3D ellipse.
type EllipticalArc interface {
	Curve
	Center() types.Point
	Normal() types.UnitVector
	MajorAxis() types.UnitVector
	MajorRadius() float64
	MinorRadius() float64
	StartAngle() float64
	SweepAngle() float64
}

// Polyline3d is a piecewise-linear 3D path.
type Polyline3d interface {
	Curve
	Points() []types.Point
}

// BSplineCurve is a 3D NURBS curve.
type BSplineCurve interface {
	Curve
	Definition() types.BSplineCurveDef
}

// Helix is a 3D (possibly conical) helix — an Oblikovati extension; the
// reference has no transient helix.
type Helix interface {
	Curve
	BasePoint() types.Point
	Axis() types.UnitVector
	StartRadius() float64
	// Pitch is the axial advance per turn; TaperPerTurn the radial change per
	// turn (0 for a true cylinder helix).
	Pitch() float64
	TaperPerTurn() float64
	Turns() float64
}

// Line2d is an unbounded 2D line.
type Line2d interface {
	Curve2d
	RootPoint() types.Point2d
	Direction() types.UnitVector2d
}

// LineSegment2d is a bounded 2D line span.
type LineSegment2d interface {
	Curve2d
	StartPoint() types.Point2d
	EndPoint() types.Point2d
}

// Circle2d is a full 2D circle.
type Circle2d interface {
	Curve2d
	Center() types.Point2d
	Radius() float64
}

// Arc2d is a 2D circular arc.
type Arc2d interface {
	Curve2d
	Center() types.Point2d
	Radius() float64
	StartAngle() float64
	SweepAngle() float64
}

// EllipseFull2d is a full 2D ellipse.
type EllipseFull2d interface {
	Curve2d
	Center() types.Point2d
	MajorAxis() types.UnitVector2d
	MajorRadius() float64
	MinorRadius() float64
}

// EllipticalArc2d is a bounded span of a 2D ellipse.
type EllipticalArc2d interface {
	Curve2d
	Center() types.Point2d
	MajorAxis() types.UnitVector2d
	MajorRadius() float64
	MinorRadius() float64
	StartAngle() float64
	SweepAngle() float64
}

// Polyline2d is a piecewise-linear 2D path.
type Polyline2d interface {
	Curve2d
	Points() []types.Point2d
}

// BSplineCurve2d is a 2D NURBS curve.
type BSplineCurve2d interface {
	Curve2d
	Definition() types.BSplineCurve2dDef
}
