// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The member-level evaluator contracts (M01-F06, #603): the workhorse query
// surface of the geometry library — curvature, higher derivatives, arc length,
// tolerance-driven stroking, closest-point classification, range boxes and
// iso-curves. Reach an evaluator from its curve/surface via Evaluator().
//
// Conventions shared by every member: parameters are the owning geometry's
// native parameterization; lengths are database units (cm); tangent and
// derivative vectors are unnormalized d/dt magnitudes; querying outside the
// domain clamps to it.

// CurveEvaluator is the member-level query surface of a 3D transient curve.
type CurveEvaluator interface {
	// RangeBox returns a box enclosing the curve (±Inf faces when unbounded;
	// a NURBS box is the control-hull box, enclosing but not minimal).
	RangeBox() types.Box
	// Continuity returns the largest maintained continuity order: 0 for a
	// polyline, degree−1 across NURBS knots, [types.ContinuityInfinite] for
	// analytic curves.
	Continuity() int
	// EndPoints returns the curve's bounds, bounded=false for an infinite curve.
	EndPoints() (start, end types.Point, bounded bool)
	// ParamExtents returns the parameter domain [lo, hi] (±Inf when unbounded).
	ParamExtents() (lo, hi float64)
	// Derivatives returns the first three parametric derivatives at t.
	Derivatives(t float64) (d1, d2, d3 types.Vector)
	// Curvature returns the unit principal-normal direction and curvature
	// magnitude at t (zero direction and magnitude where the curve is straight).
	Curvature(t float64) (direction types.Vector, magnitude float64)
	// Length returns the arc length between the two parameters (exact for
	// analytic curves, adaptive quadrature otherwise).
	Length(from, to float64) float64
	// ParamAtLength returns the parameter at signed arc length from the given
	// parameter — the inverse of Length.
	ParamAtLength(from, length float64) float64
	// Strokes tessellates [from, to] into a polyline whose chordal deviation
	// from the curve stays within tolerance.
	Strokes(from, to, tolerance float64) []types.Point
	// ParamAtPoint returns the parameter of the point on the curve closest to
	// p, classifying how many equally close answers exist.
	ParamAtPoint(p types.Point) (t float64, nature types.SolutionNature)
	// ParamAnomaly reports periodicity, singular points and unboundedness.
	ParamAnomaly() types.ParamAnomaly
}

// Curve2dEvaluator is the sketch-space analogue of [CurveEvaluator].
type Curve2dEvaluator interface {
	// RangeBox returns a 2D box enclosing the curve (±Inf when unbounded;
	// a NURBS box is the control-hull box, enclosing but not minimal).
	RangeBox() types.Box2d
	// Continuity returns the largest maintained continuity order.
	Continuity() int
	// EndPoints returns the curve's bounds, bounded=false for an infinite curve.
	EndPoints() (start, end types.Point2d, bounded bool)
	// ParamExtents returns the parameter domain [lo, hi] (±Inf when unbounded).
	ParamExtents() (lo, hi float64)
	// Derivatives returns the first three parametric derivatives at t.
	Derivatives(t float64) (d1, d2, d3 types.Vector2d)
	// Curvature returns the signed curvature at t (positive turning left along
	// increasing parameter, zero where straight).
	Curvature(t float64) float64
	// Length returns the arc length between the two parameters.
	Length(from, to float64) float64
	// ParamAtLength returns the parameter at signed arc length from the given
	// parameter — the inverse of Length.
	ParamAtLength(from, length float64) float64
	// Strokes tessellates [from, to] within the chordal tolerance.
	Strokes(from, to, tolerance float64) []types.Point2d
	// ParamAtPoint returns the parameter of the closest point to p, classified.
	ParamAtPoint(p types.Point2d) (t float64, nature types.SolutionNature)
	// ParamAnomaly reports periodicity, singular points and unboundedness.
	ParamAnomaly() types.ParamAnomaly
}

// SurfaceEvaluator is the member-level query surface of a transient surface. It is
// segregated into four embedded capability families ([SurfaceExtents],
// [SurfaceDifferential], [SurfaceProjection], [SurfaceIso]) so a consumer that only needs,
// say, the projection queries does not depend on the differential-geometry width (audit
// I9). SurfaceEvaluator stays their union — every existing implementer and caller is
// unaffected.
type SurfaceEvaluator interface {
	SurfaceExtents
	SurfaceDifferential
	SurfaceProjection
	SurfaceIso
}

// SurfaceExtents is a surface's bounding, parameter domain, area and continuity.
type SurfaceExtents interface {
	// RangeBox returns a box enclosing the surface (±Inf faces when unbounded;
	// a NURBS box is the control-net box, enclosing but not minimal).
	RangeBox() types.Box
	// ParamRangeRect returns the (u, v) parameter domain as a 2D box.
	ParamRangeRect() types.Box2d
	// Area returns the total surface area: exact for closed analytic surfaces
	// (sphere, torus), +Inf for unbounded ones, numeric quadrature for NURBS.
	Area() float64
	// Continuity returns the largest maintained continuity order.
	Continuity() int
}

// SurfaceDifferential is a surface's differential geometry — tangents, partials, curvatures.
type SurfaceDifferential interface {
	// Tangents returns the unit tangents along u and v at (u, v) (zero at
	// degeneracies such as a sphere pole).
	Tangents(u, v float64) (uTangent, vTangent types.Vector)
	// FirstPartials returns ∂P/∂u and ∂P/∂v at (u, v).
	FirstPartials(u, v float64) (pu, pv types.Vector)
	// SecondPartials returns ∂²P/∂u², ∂²P/∂u∂v and ∂²P/∂v² at (u, v).
	SecondPartials(u, v float64) (puu, puv, pvv types.Vector)
	// ThirdPartials returns the corner partials ∂³P/∂u³ and ∂³P/∂v³ at (u, v).
	ThirdPartials(u, v float64) (puuu, pvvv types.Vector)
	// Curvatures returns the principal curvatures at (u, v) and the unit
	// direction of the maximum one.
	Curvatures(u, v float64) (maxDirection types.Vector, maxCurvature, minCurvature float64)
}

// SurfaceProjection is a surface's closest-point queries — inversion and normal at a point.
type SurfaceProjection interface {
	// ParamAtPoint returns the (u, v) of the point on the surface closest to
	// p, classifying how many equally close answers exist.
	ParamAtPoint(p types.Point) (u, v float64, nature types.SolutionNature)
	// NormalAtPoint returns the unit surface normal at the point on the
	// surface closest to p.
	NormalAtPoint(p types.Point) types.Vector
}

// SurfaceIso is a surface's iso-curve extraction and parameter-domain anomalies.
type SurfaceIso interface {
	// IsoCurve extracts the curve of constant parameter: u = param when
	// uDirection (the curve runs along v), else v = param.
	IsoCurve(uDirection bool, param float64) (Curve, error)
	// ParamAnomaly reports periodicity, singularities and unboundedness of
	// each parameter direction.
	ParamAnomaly() (u, v types.ParamAnomaly)
}
