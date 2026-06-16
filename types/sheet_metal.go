// SPDX-License-Identifier: Apache-2.0

package types

// Sheet-metal rule value types (M13-F01). A sheet-metal part carries a rule (style)
// that fixes the constant material thickness, the default bend radius, the relief
// shape cut at bend ends, and the unfold method that governs how a bend develops into
// the flat pattern. These are the canonical, Apache-2.0 definitions; the GPL
// implementation aliases them (ADR-0018) and drives the bend-development kernel ops.

// UnfoldMethodType names how a bend's flat length (its bend allowance) is computed when
// the part is developed into its flat pattern. The three methods are the industry
// standard set: a single K-factor, a per-angle bend table, or a custom equation.
type UnfoldMethodType int32

const (
	// KFactorUnfold computes bend allowance from a single neutral-axis ratio (K-factor):
	// BA = angle·(radius + K·thickness). The default.
	KFactorUnfold UnfoldMethodType = iota
	// BendTableUnfold looks the bend allowance (or deduction) up in a thickness/radius/angle
	// table, interpolating between rows — used when shop tests have characterised a material.
	BendTableUnfold
	// EquationUnfold evaluates a custom user equation in {thickness, radius, angle} for the
	// bend allowance, for materials whose behaviour neither a K-factor nor a table captures.
	EquationUnfold
)

var unfoldMethodTypeNames = map[UnfoldMethodType]string{
	KFactorUnfold:   "kFactor",
	BendTableUnfold: "bendTable",
	EquationUnfold:  "equation",
}

// String returns the unfold method's wire spelling.
func (m UnfoldMethodType) String() string { return enumName(unfoldMethodTypeNames, m) }

// ParseUnfoldMethodType resolves a wire spelling back to its unfold method.
func ParseUnfoldMethodType(s string) (UnfoldMethodType, bool) {
	return enumFromName(unfoldMethodTypeNames, s)
}

// ReliefShape names the cut placed at the ends of a bend (and at corners) so the
// material can fold without tearing the adjacent web. The zero value is ReliefRound.
type ReliefShape int32

const (
	// ReliefRound cuts a rounded (filleted) notch at the bend end — the gentlest on the
	// material and the default.
	ReliefRound ReliefShape = iota
	// ReliefSquare cuts a square notch — simplest to laser/punch, slightly more stress-prone.
	ReliefSquare
	// ReliefTear leaves no cut: the material tears along the bend end (no relief geometry).
	ReliefTear
)

var reliefShapeNames = map[ReliefShape]string{
	ReliefRound:  "round",
	ReliefSquare: "square",
	ReliefTear:   "tear",
}

// String returns the relief shape's wire spelling.
func (r ReliefShape) String() string { return enumName(reliefShapeNames, r) }

// ParseReliefShape resolves a wire spelling back to its relief shape.
func ParseReliefShape(s string) (ReliefShape, bool) { return enumFromName(reliefShapeNames, s) }
