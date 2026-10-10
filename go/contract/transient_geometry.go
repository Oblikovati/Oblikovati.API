// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// TransientGeometry is the single discoverable construction point for the
// transient-geometry vocabulary (M01-F05, #602) — the reference factory of the
// same name. Value types (points, vectors, matrices, boxes) are pure data with
// their math implemented in [oblikovati.org/api/types]; this factory exists for
// the curve/surface objects, whose construction validates inputs and whose
// evaluation runs against the host kernel. Constructors error on degenerate
// input (zero directions, non-positive radii, malformed knot vectors) — they
// never return a half-valid curve.
//
// Example:
//
//	circle, err := tg.CreateCircle(types.NewPoint(0, 0, 0), zAxis, 2.5)
//	p := circle.Evaluate(0.25) // a quarter of the way around
//
// The vocabulary is segregated into three embedded capability families
// ([TransientCurves3d], [TransientCurves2d], [TransientSurfaces]) so a consumer or
// fake that needs only one axis depends only on that family (Interface Segregation,
// audit I9). TransientGeometry stays their union, so every existing implementer and
// caller is unaffected — new host code and add-in helpers should accept the narrowest
// family they use, not the whole union.
type TransientGeometry interface {
	TransientCurves3d
	TransientCurves2d
	TransientSurfaces
}

// TransientCurves3d constructs the 3D transient curves. Constructors error on degenerate
// input (zero directions, non-positive radii, malformed knot vectors).
type TransientCurves3d interface {
	CreateLine(root types.Point, direction types.UnitVector) (Line, error)
	CreateLineSegment(start, end types.Point) (LineSegment, error)
	CreateCircle(center types.Point, normal types.UnitVector, radius float64) (Circle, error)
	CreateCircleByThreePoints(a, b, c types.Point) (Circle, error)
	CreateArc(center types.Point, normal, reference types.UnitVector, radius, startAngle, sweepAngle float64) (Arc3d, error)
	CreateArcByThreePoints(start, on, end types.Point) (Arc3d, error)
	CreateEllipseFull(center types.Point, normal, majorAxis types.UnitVector, majorRadius, minorRadius float64) (EllipseFull, error)
	CreateEllipticalArc(center types.Point, normal, majorAxis types.UnitVector, majorRadius, minorRadius, startAngle, sweepAngle float64) (EllipticalArc, error)
	CreatePolyline(points []types.Point) (Polyline3d, error)
	CreateBSplineCurve(def types.BSplineCurveDef) (BSplineCurve, error)
	CreateFittedBSplineCurve(through []types.Point) (BSplineCurve, error)
	CreateHelix(base types.Point, axis, reference types.UnitVector, startRadius, pitch, taperPerTurn, turns float64, clockwise bool) (Helix, error)
}

// TransientCurves2d constructs the 2D (sketch-space) transient curves.
type TransientCurves2d interface {
	CreateLine2d(root types.Point2d, direction types.UnitVector2d) (Line2d, error)
	CreateLineSegment2d(start, end types.Point2d) (LineSegment2d, error)
	CreateCircle2d(center types.Point2d, radius float64) (Circle2d, error)
	CreateArc2d(center types.Point2d, radius, startAngle, sweepAngle float64) (Arc2d, error)
	CreateEllipseFull2d(center types.Point2d, majorAxis types.UnitVector2d, majorRadius, minorRadius float64) (EllipseFull2d, error)
	CreateEllipticalArc2d(center types.Point2d, majorAxis types.UnitVector2d, majorRadius, minorRadius, startAngle, sweepAngle float64) (EllipticalArc2d, error)
	CreatePolyline2d(points []types.Point2d) (Polyline2d, error)
	CreateBSplineCurve2d(def types.BSplineCurve2dDef) (BSplineCurve2d, error)
	CreateFittedBSplineCurve2d(through []types.Point2d) (BSplineCurve2d, error)
}

// TransientSurfaces constructs the transient surfaces (analytic and NURBS).
type TransientSurfaces interface {
	CreatePlane(root types.Point, normal types.UnitVector) (Plane, error)
	CreatePlaneByThreePoints(a, b, c types.Point) (Plane, error)
	CreateCylinder(base types.Point, axis types.UnitVector, radius float64) (Cylinder, error)
	CreateCone(apex types.Point, axis types.UnitVector, halfAngle float64) (Cone, error)
	CreateSphere(center types.Point, radius float64) (Sphere, error)
	CreateTorus(center types.Point, axis types.UnitVector, majorRadius, minorRadius float64) (Torus, error)
	CreateEllipticalCylinder(base types.Point, axis, majorAxis types.UnitVector, majorRadius, minorRadius float64) (EllipticalCylinder, error)
	CreateEllipticalCone(apex types.Point, axis, majorAxis types.UnitVector, majorHalfAngle, minorHalfAngle float64) (EllipticalCone, error)
	CreateBSplineSurface(def types.BSplineSurfaceDef) (BSplineSurface, error)
}
