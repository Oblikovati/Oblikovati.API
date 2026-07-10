// SPDX-License-Identifier: Apache-2.0

package types

// CoordinateSystemType is how a 3D equation curve's three expressions are interpreted (Inventor
// SketchEquationCurve3D CoordinateSystemTypeEnum, #1846): Cartesian x/y/z, cylindrical radius/
// theta/z, or spherical radius/theta/phi. The reference enum's numeric ids are not in-repo, so this
// is an Oblikovati-owned frozen block — the wire contract is the string spelling; never renumber the
// ids. The zero value is Cartesian, so an omitted selector keeps the pre-#1846 behaviour.
type CoordinateSystemType int32

const (
	// CoordinateSystemCartesian interprets the expressions as x(t), y(t), z(t) — the default.
	CoordinateSystemCartesian CoordinateSystemType = iota
	// CoordinateSystemCylindrical interprets them as radius(t), theta(t), z(t): the point is
	// (r·cosθ, r·sinθ, z). r/z are lengths (cm); theta is an angle (radians).
	CoordinateSystemCylindrical
	// CoordinateSystemSpherical interprets them as radius(t), theta(t), phi(t): the point is
	// (r·sinφ·cosθ, r·sinφ·sinθ, r·cosφ). r is a length (cm); theta/phi are angles (radians).
	CoordinateSystemSpherical
)

// coordinateSystemTypeNames are the frozen wire spellings.
var coordinateSystemTypeNames = map[CoordinateSystemType]string{
	CoordinateSystemCartesian:   "cartesian",
	CoordinateSystemCylindrical: "cylindrical",
	CoordinateSystemSpherical:   "spherical",
}

// String returns the coordinate system's wire spelling.
func (c CoordinateSystemType) String() string {
	return enumName(coordinateSystemTypeNames, c)
}

// ParseCoordinateSystemType resolves a wire spelling back to its coordinate system; the empty
// string maps to Cartesian so an omitted selector is the default.
func ParseCoordinateSystemType(s string) (CoordinateSystemType, bool) {
	if s == "" {
		return CoordinateSystemCartesian, true
	}
	return enumFromName(coordinateSystemTypeNames, s)
}
