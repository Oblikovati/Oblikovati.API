// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The transient-surface contracts (M01-F05, #602): each surface carries its
// kind discriminator and evaluates positions/normals over (u, v); the
// member-level evaluator surface (area, curvatures, iso-curves…) is reached
// via Evaluator() (M01-F06, #603).

// Surface is the umbrella every transient surface satisfies.
type Surface interface {
	// SurfaceType identifies the concrete kind (frozen reference values).
	SurfaceType() types.SurfaceType
	// GeometryForm classifies the representation (NURBS flags, frozen values).
	GeometryForm() types.SurfaceGeometryForm
	// Evaluate returns the position at (u, v).
	Evaluate(u, v float64) types.Point
	// Normal returns the unit surface normal at (u, v) (zero at degeneracies).
	Normal(u, v float64) types.Vector
	// Domains returns the (u, v) parameter ranges (±Inf where unbounded).
	Domains() (uLo, uHi, vLo, vHi float64)
	// Parameter inverts Evaluate for a point on (or near) the surface.
	Parameter(p types.Point) (u, v float64)
	// Evaluator returns the member-level query surface of this surface.
	Evaluator() SurfaceEvaluator
}

// Plane is an unbounded planar surface with an orthonormal in-plane basis.
type Plane interface {
	Surface
	RootPoint() types.Point
	PlaneNormal() types.UnitVector
	UAxis() types.UnitVector
	VAxis() types.UnitVector
}

// Cylinder is an unbounded circular cylinder.
type Cylinder interface {
	Surface
	BasePoint() types.Point
	Axis() types.UnitVector
	Radius() float64
}

// Cone is an unbounded circular cone (apex at v = 0, opening along the axis).
type Cone interface {
	Surface
	ApexPoint() types.Point
	Axis() types.UnitVector
	HalfAngle() float64
}

// Sphere is a full sphere (u longitude, v latitude).
type Sphere interface {
	Surface
	Center() types.Point
	Radius() float64
}

// Torus is a full torus (u around the axis, v around the tube).
type Torus interface {
	Surface
	Center() types.Point
	Axis() types.UnitVector
	MajorRadius() float64
	MinorRadius() float64
}

// EllipticalCylinder is an unbounded cylinder with an elliptical section.
type EllipticalCylinder interface {
	Surface
	BasePoint() types.Point
	Axis() types.UnitVector
	MajorAxis() types.UnitVector
	MajorRadius() float64
	MinorRadius() float64
}

// EllipticalCone is an unbounded cone with an elliptical section.
type EllipticalCone interface {
	Surface
	ApexPoint() types.Point
	Axis() types.UnitVector
	MajorAxis() types.UnitVector
	MajorHalfAngle() float64
	MinorHalfAngle() float64
}

// BSplineSurface is a NURBS surface.
type BSplineSurface interface {
	Surface
	Definition() types.BSplineSurfaceDef
}
