// SPDX-License-Identifier: Apache-2.0

package types

import "math"

// The transient-geometry value types (M01-F05, #602). They are implemented HERE,
// in the Apache-2.0 contract module, because they are ownerless, immutable pure
// math created in huge volume: an add-in computes with them locally at zero wire
// cost, and the GPL host converts to its kernel types at the boundary (the
// in-proc/over-the-wire split recorded in ADR-0018's geometry addendum).
//
// JSON form: fixed-length arrays ([x,y,z] / [x,y] / 16 cells), byte-compatible
// with the ad-hoc [3]float64 fields the wire used before these types existed.

// Point is a 3D position. Lengths are in database units (cm), like every
// geometric quantity crossing the contract.
type Point struct {
	X, Y, Z float64
}

// NewPoint builds a point from its coordinates (the TransientGeometry.CreatePoint
// equivalent for local computation).
func NewPoint(x, y, z float64) Point { return Point{X: x, Y: y, Z: z} }

// VectorTo returns the displacement from p to o.
func (p Point) VectorTo(o Point) Vector { return Vector{o.X - p.X, o.Y - p.Y, o.Z - p.Z} }

// TranslateBy returns p displaced by v.
func (p Point) TranslateBy(v Vector) Point { return Point{p.X + v.X, p.Y + v.Y, p.Z + v.Z} }

// DistanceTo returns the Euclidean distance to o.
func (p Point) DistanceTo(o Point) float64 { return p.VectorTo(o).Length() }

// Midpoint returns the point halfway to o.
func (p Point) Midpoint(o Point) Point {
	return Point{(p.X + o.X) / 2, (p.Y + o.Y) / 2, (p.Z + o.Z) / 2}
}

// IsEqualTo reports coordinate-wise equality within tol.
func (p Point) IsEqualTo(o Point, tol float64) bool {
	return math.Abs(p.X-o.X) <= tol && math.Abs(p.Y-o.Y) <= tol && math.Abs(p.Z-o.Z) <= tol
}

// MarshalJSON encodes the point as [x,y,z].
func (p Point) MarshalJSON() ([]byte, error) { return marshalFloats(p.X, p.Y, p.Z) }

// UnmarshalJSON decodes [x,y,z].
func (p *Point) UnmarshalJSON(b []byte) error {
	return unmarshalFloats(b, "Point", &p.X, &p.Y, &p.Z)
}

// Point2d is a 2D position (sketch space).
type Point2d struct {
	X, Y float64
}

// NewPoint2d builds a 2D point from its coordinates.
func NewPoint2d(x, y float64) Point2d { return Point2d{X: x, Y: y} }

// VectorTo returns the displacement from p to o.
func (p Point2d) VectorTo(o Point2d) Vector2d { return Vector2d{o.X - p.X, o.Y - p.Y} }

// TranslateBy returns p displaced by v.
func (p Point2d) TranslateBy(v Vector2d) Point2d { return Point2d{p.X + v.X, p.Y + v.Y} }

// DistanceTo returns the Euclidean distance to o.
func (p Point2d) DistanceTo(o Point2d) float64 { return p.VectorTo(o).Length() }

// Midpoint returns the point halfway to o.
func (p Point2d) Midpoint(o Point2d) Point2d { return Point2d{(p.X + o.X) / 2, (p.Y + o.Y) / 2} }

// IsEqualTo reports coordinate-wise equality within tol.
func (p Point2d) IsEqualTo(o Point2d, tol float64) bool {
	return math.Abs(p.X-o.X) <= tol && math.Abs(p.Y-o.Y) <= tol
}

// MarshalJSON encodes the point as [x,y].
func (p Point2d) MarshalJSON() ([]byte, error) { return marshalFloats(p.X, p.Y) }

// UnmarshalJSON decodes [x,y].
func (p *Point2d) UnmarshalJSON(b []byte) error {
	return unmarshalFloats(b, "Point2d", &p.X, &p.Y)
}
