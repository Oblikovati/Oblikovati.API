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

// ReliefShape names the cut placed at the ends of a BEND so the material can fold without
// tearing the adjacent web — Inventor's BendReliefShapeEnum.
type ReliefShape int32

const (
	// ReliefRound cuts a rounded (filleted) notch at the bend end — the gentlest on the
	// material. Inventor's kRound.
	ReliefRound ReliefShape = iota
	// ReliefStraight cuts a plain rectangular notch — simplest to laser/punch, and Inventor's
	// kStraight, which is the shipped default of its Default style. It was spelled "square"
	// before the enum was reconciled with Inventor's (#1960); that spelling still parses and
	// means the same rectangular cut, so an existing style is unchanged.
	ReliefStraight
	// ReliefTear leaves no cut: the material tears along the bend end (no relief geometry).
	ReliefTear
)

var reliefShapeNames = map[ReliefShape]string{
	ReliefRound:    "round",
	ReliefStraight: "straight",
	ReliefTear:     "tear",
}

// String returns the relief shape's wire spelling.
func (r ReliefShape) String() string { return enumName(reliefShapeNames, r) }

// ParseReliefShape resolves a wire spelling back to its relief shape. "square" is the older
// spelling of ReliefStraight and still resolves to it.
func ParseReliefShape(s string) (ReliefShape, bool) {
	if s == "square" {
		return ReliefStraight, true
	}
	return enumFromName(reliefShapeNames, s)
}

// CornerReliefShape names the cut placed where flanges meet at a CORNER — a different set from
// the bend reliefs, and a separate style property (Inventor's CornerReliefShapeEnum, #1960).
type CornerReliefShape int32

const (
	// CornerTrimToBend trims the corner back to the bend tangents — Inventor's kTrimToBend, the
	// shipped default of its Default style, and the zero value here for the same reason.
	CornerTrimToBend CornerReliefShape = iota
	// CornerRound / CornerSquare / CornerTear are the plain notch shapes.
	CornerRound
	CornerSquare
	CornerTear
	// CornerFullRound cuts the corner to a full radius joining both bend reliefs; CornerRoundWithRadius
	// is the same with an explicit radius (Inventor's default THREE-BEND corner relief).
	CornerFullRound
	CornerRoundWithRadius
	// CornerIntersection leaves the two walls running into each other, relieved only where they cross.
	CornerIntersection
)

var cornerReliefShapeNames = map[CornerReliefShape]string{
	CornerTrimToBend:      "trimToBend",
	CornerRound:           "round",
	CornerSquare:          "square",
	CornerTear:            "tear",
	CornerFullRound:       "fullRound",
	CornerRoundWithRadius: "roundWithRadius",
	CornerIntersection:    "intersection",
}

// String returns the corner-relief shape's wire spelling.
func (c CornerReliefShape) String() string { return enumName(cornerReliefShapeNames, c) }

// ParseCornerReliefShape resolves a wire spelling back to its corner-relief shape.
func ParseCornerReliefShape(s string) (CornerReliefShape, bool) {
	return enumFromName(cornerReliefShapeNames, s)
}

// CornerReliefPlacement says where the corner relief sits relative to the bend tangents
// (Inventor's CornerReliefPlacementEnum, #1960).
type CornerReliefPlacement int32

const (
	// CornerReliefAtBendTangent places the relief on the bend tangent lines — the default.
	CornerReliefAtBendTangent CornerReliefPlacement = iota
	// CornerReliefAtBendIntersection places it where the two bends' tangents cross.
	CornerReliefAtBendIntersection
	// CornerReliefAtAlongBend places it along the bend rather than at its end.
	CornerReliefAtAlongBend
)

var cornerReliefPlacementNames = map[CornerReliefPlacement]string{
	CornerReliefAtBendTangent:      "bendTangent",
	CornerReliefAtBendIntersection: "bendIntersection",
	CornerReliefAtAlongBend:        "alongBend",
}

// String returns the placement's wire spelling.
func (c CornerReliefPlacement) String() string { return enumName(cornerReliefPlacementNames, c) }

// ParseCornerReliefPlacement resolves a wire spelling back to its placement.
func ParseCornerReliefPlacement(s string) (CornerReliefPlacement, bool) {
	return enumFromName(cornerReliefPlacementNames, s)
}

// BendTransition is how the material is shaped where a bend zone runs into the face beside it —
// Inventor's BendTransitionEnum (#1959).
//
// Read the shapes carefully before assuming they are all the same kind of thing: three of them
// describe the FLAT PATTERN's outline through the transition region (a straight line across the
// bend zone, a straight line to the bent feature's edge, or an arc tangent to both), one is a CUT
// in the folded model, and the default is neither — the geometry simply runs on as it does.
type BendTransition int32

const (
	// NoBendTransition leaves the material as the geometry makes it, which is Inventor's shipped
	// default and the zero value here for the same reason.
	NoBendTransition BendTransition = iota
	// IntersectionBendTransition runs a straight line from the bend zone's edge to where it meets
	// the bent feature's edge.
	IntersectionBendTransition
	// StraightLineBendTransition runs a straight line from one edge of the bend zone to the other.
	StraightLineBendTransition
	// ArcBendTransition replaces that straight line with an arc of BendTransitionArcRadius,
	// tangent to the bent feature's edge and to the straight transition.
	ArcBendTransition
	// TrimToBendBendTransition cuts the bend zone back perpendicular to the bent feature — the one
	// transition that shows in the FOLDED model rather than only in the flat.
	TrimToBendBendTransition
	// DefaultBendTransition defers to the style, for a per-feature override that does not override.
	DefaultBendTransition
)

var bendTransitionNames = map[BendTransition]string{
	NoBendTransition:           "none",
	IntersectionBendTransition: "intersection",
	StraightLineBendTransition: "straightLine",
	ArcBendTransition:          "arc",
	TrimToBendBendTransition:   "trimToBend",
	DefaultBendTransition:      "default",
}

// String returns the transition's wire spelling.
func (b BendTransition) String() string { return enumName(bendTransitionNames, b) }

// ParseBendTransition resolves a wire spelling back to its transition.
func ParseBendTransition(s string) (BendTransition, bool) {
	return enumFromName(bendTransitionNames, s)
}
