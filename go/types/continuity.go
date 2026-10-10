// SPDX-License-Identifier: Apache-2.0

package types

// SurfaceContinuity is the geometric-continuity order a Class-A surfacing operation holds to its
// neighbours across a seam: G0 (position), G1 (tangent), G2 (curvature), G3 (curvature-rate). It is
// the public vocabulary for operations that take a continuity choice — the boundary fill (M36-F07),
// and reusable by match/extend. It is a string enum so it reads self-describingly on the wire; the
// kernel maps it to its internal derivative order via [SurfaceContinuity.Order].
type SurfaceContinuity string

const (
	// ContinuityG0 holds position only (the surfaces meet at the seam; tangents may break).
	ContinuityG0 SurfaceContinuity = "g0"
	// ContinuityG1 holds tangency (a shared tangent plane along the seam).
	ContinuityG1 SurfaceContinuity = "g1"
	// ContinuityG2 holds curvature (no curvature jump across the seam) — the common Class-A choice.
	ContinuityG2 SurfaceContinuity = "g2"
	// ContinuityG3 holds curvature-rate (a level beyond G2).
	ContinuityG3 SurfaceContinuity = "g3"
)

// Order returns the derivative order the continuity imposes: G0→0, G1→1, G2→2, G3→3. An empty or
// unrecognized value returns -1 (use [ParseSurfaceContinuity] to validate and default first).
func (c SurfaceContinuity) Order() int {
	switch c {
	case ContinuityG0:
		return 0
	case ContinuityG1:
		return 1
	case ContinuityG2:
		return 2
	case ContinuityG3:
		return 3
	default:
		return -1
	}
}

// String returns the continuity's wire spelling.
func (c SurfaceContinuity) String() string { return string(c) }

// ParseSurfaceContinuity resolves a wire spelling to a continuity; the empty string maps to the
// supplied fallback so a caller can pick its own default (e.g. the fill op defaults to G2).
func ParseSurfaceContinuity(s string, fallback SurfaceContinuity) (SurfaceContinuity, bool) {
	if s == "" {
		return fallback, true
	}
	c := SurfaceContinuity(s)
	if c.Order() < 0 {
		return "", false
	}
	return c, true
}
