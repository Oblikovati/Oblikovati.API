// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"math"
)

// Matrix is a 4×4 affine 3D transform, row-major, with the bottom row expected
// to stay [0,0,0,1] (the builders guarantee it; a hand-built matrix is the
// caller's responsibility). JSON form: the 16 cells as one array.
type Matrix struct {
	Cells [16]float64
}

// IdentityMatrix returns the do-nothing transform.
func IdentityMatrix() Matrix {
	return Matrix{Cells: [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}}
}

// TranslationMatrix returns the transform that displaces by v.
func TranslationMatrix(v Vector) Matrix {
	m := IdentityMatrix()
	m.Cells[3], m.Cells[7], m.Cells[11] = v.X, v.Y, v.Z
	return m
}

// RotationMatrix returns the rotation by angle (radians) about the axis through
// center (Rodrigues, composed T·R·T⁻¹).
func RotationMatrix(angle float64, axis UnitVector, center Point) Matrix {
	c, s := math.Cos(angle), math.Sin(angle)
	t := 1 - c
	x, y, z := axis.X, axis.Y, axis.Z
	r := [9]float64{
		t*x*x + c, t*x*y - s*z, t*x*z + s*y,
		t*x*y + s*z, t*y*y + c, t*y*z - s*x,
		t*x*z - s*y, t*y*z + s*x, t*z*z + c,
	}
	tx := center.X - (r[0]*center.X + r[1]*center.Y + r[2]*center.Z)
	ty := center.Y - (r[3]*center.X + r[4]*center.Y + r[5]*center.Z)
	tz := center.Z - (r[6]*center.X + r[7]*center.Y + r[8]*center.Z)
	return Matrix{Cells: [16]float64{
		r[0], r[1], r[2], tx,
		r[3], r[4], r[5], ty,
		r[6], r[7], r[8], tz,
		0, 0, 0, 1,
	}}
}

// CoordinateSystemMatrix returns the transform mapping the standard frame onto
// the frame (origin, xAxis, yAxis, zAxis) — column-axes form.
func CoordinateSystemMatrix(origin Point, xAxis, yAxis, zAxis UnitVector) Matrix {
	return Matrix{Cells: [16]float64{
		xAxis.X, yAxis.X, zAxis.X, origin.X,
		xAxis.Y, yAxis.Y, zAxis.Y, origin.Y,
		xAxis.Z, yAxis.Z, zAxis.Z, origin.Z,
		0, 0, 0, 1,
	}}
}

// At returns the cell at (row, col), both 0-based.
func (m Matrix) At(row, col int) float64 { return m.Cells[row*4+col] }

// Translation returns the transform's displacement column.
func (m Matrix) Translation() Vector { return Vector{m.Cells[3], m.Cells[7], m.Cells[11]} }

// TransformPoint applies the full affine transform to a position.
func (m Matrix) TransformPoint(p Point) Point {
	c := m.Cells
	return Point{
		c[0]*p.X + c[1]*p.Y + c[2]*p.Z + c[3],
		c[4]*p.X + c[5]*p.Y + c[6]*p.Z + c[7],
		c[8]*p.X + c[9]*p.Y + c[10]*p.Z + c[11],
	}
}

// TransformVector applies the linear part only (displacements ignore translation).
func (m Matrix) TransformVector(v Vector) Vector {
	c := m.Cells
	return Vector{
		c[0]*v.X + c[1]*v.Y + c[2]*v.Z,
		c[4]*v.X + c[5]*v.Y + c[6]*v.Z,
		c[8]*v.X + c[9]*v.Y + c[10]*v.Z,
	}
}

// Mul returns m·o (apply o first, then m).
func (m Matrix) Mul(o Matrix) Matrix {
	var out Matrix
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			sum := 0.0
			for k := 0; k < 4; k++ {
				sum += m.Cells[r*4+k] * o.Cells[k*4+c]
			}
			out.Cells[r*4+c] = sum
		}
	}
	return out
}

// Determinant returns the determinant of the linear (3×3) part — affine
// matrices' bottom row contributes nothing.
func (m Matrix) Determinant() float64 {
	c := m.Cells
	return c[0]*(c[5]*c[10]-c[6]*c[9]) - c[1]*(c[4]*c[10]-c[6]*c[8]) + c[2]*(c[4]*c[9]-c[5]*c[8])
}

// Invert returns the inverse transform, erroring on a singular linear part.
func (m Matrix) Invert() (Matrix, error) {
	det := m.Determinant()
	if math.Abs(det) < 1e-300 {
		return Matrix{}, fmt.Errorf("types: matrix is singular (determinant %g)", det)
	}
	c := m.Cells
	inv := [9]float64{ // adjugate of the linear part / det
		(c[5]*c[10] - c[6]*c[9]) / det, (c[2]*c[9] - c[1]*c[10]) / det, (c[1]*c[6] - c[2]*c[5]) / det,
		(c[6]*c[8] - c[4]*c[10]) / det, (c[0]*c[10] - c[2]*c[8]) / det, (c[2]*c[4] - c[0]*c[6]) / det,
		(c[4]*c[9] - c[5]*c[8]) / det, (c[1]*c[8] - c[0]*c[9]) / det, (c[0]*c[5] - c[1]*c[4]) / det,
	}
	tx := -(inv[0]*c[3] + inv[1]*c[7] + inv[2]*c[11])
	ty := -(inv[3]*c[3] + inv[4]*c[7] + inv[5]*c[11])
	tz := -(inv[6]*c[3] + inv[7]*c[7] + inv[8]*c[11])
	return Matrix{Cells: [16]float64{
		inv[0], inv[1], inv[2], tx,
		inv[3], inv[4], inv[5], ty,
		inv[6], inv[7], inv[8], tz,
		0, 0, 0, 1,
	}}, nil
}

// MarshalJSON encodes the 16 cells as one array.
func (m Matrix) MarshalJSON() ([]byte, error) { return marshalFloats(m.Cells[:]...) }

// UnmarshalJSON decodes a 16-cell array.
func (m *Matrix) UnmarshalJSON(b []byte) error {
	ptrs := make([]*float64, 16)
	for i := range m.Cells {
		ptrs[i] = &m.Cells[i]
	}
	return unmarshalFloats(b, "Matrix", ptrs...)
}

// Matrix2d is a 3×3 affine 2D transform, row-major, bottom row [0,0,1].
type Matrix2d struct {
	Cells [9]float64
}

// IdentityMatrix2d returns the do-nothing 2D transform.
func IdentityMatrix2d() Matrix2d {
	return Matrix2d{Cells: [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1}}
}

// TranslationMatrix2d returns the transform that displaces by v.
func TranslationMatrix2d(v Vector2d) Matrix2d {
	m := IdentityMatrix2d()
	m.Cells[2], m.Cells[5] = v.X, v.Y
	return m
}

// RotationMatrix2d returns the rotation by angle (radians) about center.
func RotationMatrix2d(angle float64, center Point2d) Matrix2d {
	c, s := math.Cos(angle), math.Sin(angle)
	tx := center.X - (c*center.X - s*center.Y)
	ty := center.Y - (s*center.X + c*center.Y)
	return Matrix2d{Cells: [9]float64{c, -s, tx, s, c, ty, 0, 0, 1}}
}

// At returns the cell at (row, col), both 0-based.
func (m Matrix2d) At(row, col int) float64 { return m.Cells[row*3+col] }

// Translation returns the transform's displacement column.
func (m Matrix2d) Translation() Vector2d { return Vector2d{m.Cells[2], m.Cells[5]} }

// TransformPoint applies the full affine transform to a position.
func (m Matrix2d) TransformPoint(p Point2d) Point2d {
	c := m.Cells
	return Point2d{c[0]*p.X + c[1]*p.Y + c[2], c[3]*p.X + c[4]*p.Y + c[5]}
}

// TransformVector applies the linear part only.
func (m Matrix2d) TransformVector(v Vector2d) Vector2d {
	c := m.Cells
	return Vector2d{c[0]*v.X + c[1]*v.Y, c[3]*v.X + c[4]*v.Y}
}

// Mul returns m·o (apply o first, then m).
func (m Matrix2d) Mul(o Matrix2d) Matrix2d {
	var out Matrix2d
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			sum := 0.0
			for k := 0; k < 3; k++ {
				sum += m.Cells[r*3+k] * o.Cells[k*3+c]
			}
			out.Cells[r*3+c] = sum
		}
	}
	return out
}

// MarshalJSON encodes the 9 cells as one array.
func (m Matrix2d) MarshalJSON() ([]byte, error) { return marshalFloats(m.Cells[:]...) }

// UnmarshalJSON decodes a 9-cell array.
func (m *Matrix2d) UnmarshalJSON(b []byte) error {
	ptrs := make([]*float64, 9)
	for i := range m.Cells {
		ptrs[i] = &m.Cells[i]
	}
	return unmarshalFloats(b, "Matrix2d", ptrs...)
}
