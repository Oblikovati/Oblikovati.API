// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"math"
)

// Vector is a 3D displacement: direction and magnitude, unaffected by
// translation (unlike a [Point]).
type Vector struct {
	X, Y, Z float64
}

// NewVector builds a vector from its components.
func NewVector(x, y, z float64) Vector { return Vector{X: x, Y: y, Z: z} }

// Add / Sub / Scale / Negate are the affine combinators.
func (v Vector) Add(o Vector) Vector    { return Vector{v.X + o.X, v.Y + o.Y, v.Z + o.Z} }
func (v Vector) Sub(o Vector) Vector    { return Vector{v.X - o.X, v.Y - o.Y, v.Z - o.Z} }
func (v Vector) Scale(s float64) Vector { return Vector{v.X * s, v.Y * s, v.Z * s} }
func (v Vector) Negate() Vector         { return v.Scale(-1) }
func (v Vector) Dot(o Vector) float64   { return v.X*o.X + v.Y*o.Y + v.Z*o.Z }
func (v Vector) Cross(o Vector) Vector {
	return Vector{v.Y*o.Z - v.Z*o.Y, v.Z*o.X - v.X*o.Z, v.X*o.Y - v.Y*o.X}
}

// Length / LengthSquared are the Euclidean norms.
func (v Vector) Length() float64        { return math.Sqrt(v.Dot(v)) }
func (v Vector) LengthSquared() float64 { return v.Dot(v) }

// AsPoint reinterprets the vector as a position from the origin.
func (v Vector) AsPoint() Point { return Point(v) }

// AsUnit normalizes the vector, erroring on (near-)zero length rather than
// returning garbage.
func (v Vector) AsUnit() (UnitVector, error) { return NewUnitVector(v.X, v.Y, v.Z) }

// AngleTo returns the unsigned angle to o in radians (0 for a zero vector pair).
func (v Vector) AngleTo(o Vector) float64 {
	denom := v.Length() * o.Length()
	if denom == 0 {
		return 0
	}
	return math.Acos(clamp(v.Dot(o)/denom, -1, 1))
}

// IsEqualTo reports component-wise equality within tol.
func (v Vector) IsEqualTo(o Vector, tol float64) bool {
	return math.Abs(v.X-o.X) <= tol && math.Abs(v.Y-o.Y) <= tol && math.Abs(v.Z-o.Z) <= tol
}

// MarshalJSON encodes the vector as [x,y,z].
func (v Vector) MarshalJSON() ([]byte, error) { return marshalFloats(v.X, v.Y, v.Z) }

// UnmarshalJSON decodes [x,y,z].
func (v *Vector) UnmarshalJSON(b []byte) error {
	return unmarshalFloats(b, "Vector", &v.X, &v.Y, &v.Z)
}

// Vector2d is a 2D displacement.
type Vector2d struct {
	X, Y float64
}

// NewVector2d builds a 2D vector from its components.
func NewVector2d(x, y float64) Vector2d { return Vector2d{X: x, Y: y} }

func (v Vector2d) Add(o Vector2d) Vector2d  { return Vector2d{v.X + o.X, v.Y + o.Y} }
func (v Vector2d) Sub(o Vector2d) Vector2d  { return Vector2d{v.X - o.X, v.Y - o.Y} }
func (v Vector2d) Scale(s float64) Vector2d { return Vector2d{v.X * s, v.Y * s} }
func (v Vector2d) Negate() Vector2d         { return v.Scale(-1) }
func (v Vector2d) Dot(o Vector2d) float64   { return v.X*o.X + v.Y*o.Y }

// Cross returns the scalar (z) cross product — the signed area term.
func (v Vector2d) Cross(o Vector2d) float64 { return v.X*o.Y - v.Y*o.X }

func (v Vector2d) Length() float64        { return math.Sqrt(v.Dot(v)) }
func (v Vector2d) LengthSquared() float64 { return v.Dot(v) }

// AsPoint reinterprets the vector as a position from the origin.
func (v Vector2d) AsPoint() Point2d { return Point2d(v) }

// AsUnit normalizes the vector, erroring on (near-)zero length.
func (v Vector2d) AsUnit() (UnitVector2d, error) { return NewUnitVector2d(v.X, v.Y) }

// AngleTo returns the unsigned angle to o in radians.
func (v Vector2d) AngleTo(o Vector2d) float64 {
	denom := v.Length() * o.Length()
	if denom == 0 {
		return 0
	}
	return math.Acos(clamp(v.Dot(o)/denom, -1, 1))
}

// IsEqualTo reports component-wise equality within tol.
func (v Vector2d) IsEqualTo(o Vector2d, tol float64) bool {
	return math.Abs(v.X-o.X) <= tol && math.Abs(v.Y-o.Y) <= tol
}

// MarshalJSON encodes the vector as [x,y].
func (v Vector2d) MarshalJSON() ([]byte, error) { return marshalFloats(v.X, v.Y) }

// UnmarshalJSON decodes [x,y].
func (v *Vector2d) UnmarshalJSON(b []byte) error {
	return unmarshalFloats(b, "Vector2d", &v.X, &v.Y)
}

// UnitVector is a 3D direction with the length-one invariant enforced at
// construction (NewUnitVector errors on a zero vector).
type UnitVector struct {
	X, Y, Z float64
}

// unitTol is how far from length 1 a stored unit vector may drift before
// construction renormalizes it.
const unitTol = 1e-12

// NewUnitVector normalizes (x,y,z), erroring when the length is (near-)zero.
func NewUnitVector(x, y, z float64) (UnitVector, error) {
	l := math.Sqrt(x*x + y*y + z*z)
	if l < unitTol {
		return UnitVector{}, fmt.Errorf("types: cannot normalize the near-zero vector (%g, %g, %g)", x, y, z)
	}
	return UnitVector{x / l, y / l, z / l}, nil
}

// AsVector widens the direction to a plain vector.
func (u UnitVector) AsVector() Vector { return Vector(u) }

// Negate flips the direction.
func (u UnitVector) Negate() UnitVector { return UnitVector{-u.X, -u.Y, -u.Z} }

// Dot / Cross / AngleTo mirror the vector forms.
func (u UnitVector) Dot(o UnitVector) float64  { return u.AsVector().Dot(o.AsVector()) }
func (u UnitVector) Cross(o UnitVector) Vector { return u.AsVector().Cross(o.AsVector()) }
func (u UnitVector) AngleTo(o UnitVector) float64 {
	return math.Acos(clamp(u.Dot(o), -1, 1))
}

// MarshalJSON encodes the direction as [x,y,z].
func (u UnitVector) MarshalJSON() ([]byte, error) { return marshalFloats(u.X, u.Y, u.Z) }

// UnmarshalJSON decodes [x,y,z], renormalizing so a hand-written payload still
// satisfies the invariant (a zero direction is rejected).
func (u *UnitVector) UnmarshalJSON(b []byte) error {
	var x, y, z float64
	if err := unmarshalFloats(b, "UnitVector", &x, &y, &z); err != nil {
		return err
	}
	unit, err := NewUnitVector(x, y, z)
	if err != nil {
		return err
	}
	*u = unit
	return nil
}

// UnitVector2d is a 2D direction with the length-one invariant.
type UnitVector2d struct {
	X, Y float64
}

// NewUnitVector2d normalizes (x,y), erroring when the length is (near-)zero.
func NewUnitVector2d(x, y float64) (UnitVector2d, error) {
	l := math.Sqrt(x*x + y*y)
	if l < unitTol {
		return UnitVector2d{}, fmt.Errorf("types: cannot normalize the near-zero vector (%g, %g)", x, y)
	}
	return UnitVector2d{x / l, y / l}, nil
}

// AsVector widens the direction to a plain vector.
func (u UnitVector2d) AsVector() Vector2d { return Vector2d(u) }

// Negate flips the direction.
func (u UnitVector2d) Negate() UnitVector2d { return UnitVector2d{-u.X, -u.Y} }

func (u UnitVector2d) Dot(o UnitVector2d) float64 { return u.AsVector().Dot(o.AsVector()) }
func (u UnitVector2d) AngleTo(o UnitVector2d) float64 {
	return math.Acos(clamp(u.Dot(o), -1, 1))
}

// MarshalJSON encodes the direction as [x,y].
func (u UnitVector2d) MarshalJSON() ([]byte, error) { return marshalFloats(u.X, u.Y) }

// UnmarshalJSON decodes [x,y] with renormalization.
func (u *UnitVector2d) UnmarshalJSON(b []byte) error {
	var x, y float64
	if err := unmarshalFloats(b, "UnitVector2d", &x, &y); err != nil {
		return err
	}
	unit, err := NewUnitVector2d(x, y)
	if err != nil {
		return err
	}
	*u = unit
	return nil
}

// clamp bounds v to [lo, hi] (acos domain safety).
func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
