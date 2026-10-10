// SPDX-License-Identifier: Apache-2.0

package types

// Box is an axis-aligned 3D bounding range. An empty box has Min > Max in some
// component (the zero value, with both at the origin, is the degenerate
// single-point box; use NewEmptyBox for the absorbing identity of Union).
type Box struct {
	Min Point `json:"min"`
	Max Point `json:"max"`
}

// NewEmptyBox returns the Union identity: extending it with anything yields
// exactly that thing.
func NewEmptyBox() Box {
	const inf = 1.797693134862315708145274237317043567981e+308
	return Box{Min: Point{inf, inf, inf}, Max: Point{-inf, -inf, -inf}}
}

// IsEmpty reports whether the box contains no points.
func (b Box) IsEmpty() bool { return b.Min.X > b.Max.X || b.Min.Y > b.Max.Y || b.Min.Z > b.Max.Z }

// Center returns the box midpoint (meaningless for an empty box).
func (b Box) Center() Point { return b.Min.Midpoint(b.Max) }

// Size returns the per-axis extents.
func (b Box) Size() Vector { return b.Min.VectorTo(b.Max) }

// Contains reports whether p lies inside (inclusive).
func (b Box) Contains(p Point) bool {
	return p.X >= b.Min.X && p.X <= b.Max.X &&
		p.Y >= b.Min.Y && p.Y <= b.Max.Y &&
		p.Z >= b.Min.Z && p.Z <= b.Max.Z
}

// Extend returns the box grown to include p.
func (b Box) Extend(p Point) Box {
	return Box{
		Min: Point{min2(b.Min.X, p.X), min2(b.Min.Y, p.Y), min2(b.Min.Z, p.Z)},
		Max: Point{max2(b.Max.X, p.X), max2(b.Max.Y, p.Y), max2(b.Max.Z, p.Z)},
	}
}

// Union returns the smallest box containing both.
func (b Box) Union(o Box) Box { return b.Extend(o.Min).Extend(o.Max) }

// Box2d is an axis-aligned 2D bounding range.
type Box2d struct {
	Min Point2d `json:"min"`
	Max Point2d `json:"max"`
}

// NewEmptyBox2d returns the Union identity.
func NewEmptyBox2d() Box2d {
	const inf = 1.797693134862315708145274237317043567981e+308
	return Box2d{Min: Point2d{inf, inf}, Max: Point2d{-inf, -inf}}
}

// IsEmpty reports whether the box contains no points.
func (b Box2d) IsEmpty() bool { return b.Min.X > b.Max.X || b.Min.Y > b.Max.Y }

// Center returns the box midpoint.
func (b Box2d) Center() Point2d { return b.Min.Midpoint(b.Max) }

// Contains reports whether p lies inside (inclusive).
func (b Box2d) Contains(p Point2d) bool {
	return p.X >= b.Min.X && p.X <= b.Max.X && p.Y >= b.Min.Y && p.Y <= b.Max.Y
}

// Extend returns the box grown to include p.
func (b Box2d) Extend(p Point2d) Box2d {
	return Box2d{
		Min: Point2d{min2(b.Min.X, p.X), min2(b.Min.Y, p.Y)},
		Max: Point2d{max2(b.Max.X, p.X), max2(b.Max.Y, p.Y)},
	}
}

// Union returns the smallest box containing both.
func (b Box2d) Union(o Box2d) Box2d { return b.Extend(o.Min).Extend(o.Max) }

// OrientedBox is a box aligned to its own orthonormal frame: corner point plus
// three edge DIRECTIONS with per-axis extents (the reference's corner+vectors
// form, expressed with the unit/length split kept explicit).
type OrientedBox struct {
	Corner  Point      `json:"corner"`
	XAxis   UnitVector `json:"xAxis"`
	YAxis   UnitVector `json:"yAxis"`
	ZAxis   UnitVector `json:"zAxis"`
	Extents Vector     `json:"extents"` // edge lengths along the three axes
}

// Corners returns the eight corner points.
func (b OrientedBox) Corners() [8]Point {
	var out [8]Point
	i := 0
	for _, dz := range []float64{0, b.Extents.Z} {
		for _, dy := range []float64{0, b.Extents.Y} {
			for _, dx := range []float64{0, b.Extents.X} {
				out[i] = b.Corner.
					TranslateBy(b.XAxis.AsVector().Scale(dx)).
					TranslateBy(b.YAxis.AsVector().Scale(dy)).
					TranslateBy(b.ZAxis.AsVector().Scale(dz))
				i++
			}
		}
	}
	return out
}

func min2(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max2(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
