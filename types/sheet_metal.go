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

// CornerSeamType names how the seam is finished where two flange walls meet at a corner —
// Inventor's CornerTypeEnum plus the ripped (gap) corner it models with IsRippedCorner (#1964).
// The four are not interchangeable relief styles: gap LEAVES a controlled gap, no-overlap butts
// the two walls at a miter, and the two overlaps lap one wall over the other (differing only in
// WHICH wall is on top), so the choice changes the manufactured corner, not merely its size.
type CornerSeamType int32

const (
	// CornerSeamGap leaves a gap between the two walls (Inventor's ripped corner) — the default,
	// and the zero value so an existing seam record (which stored only a gap) reads back unchanged.
	CornerSeamGap CornerSeamType = iota
	// CornerSeamOverlap laps one wall OVER the other by PercentOverlap — Inventor's kCornerOverlap.
	CornerSeamOverlap
	// CornerSeamReverseOverlap is the same lap with the walls' roles swapped (the other wall on
	// top) — Inventor's kCornerReverseOverlap.
	CornerSeamReverseOverlap
	// CornerSeamNoOverlap butts the two walls with neither gap nor lap — Inventor's kCornerNoOverlap.
	CornerSeamNoOverlap
)

var cornerSeamTypeNames = map[CornerSeamType]string{
	CornerSeamGap:            "gap",
	CornerSeamOverlap:        "overlap",
	CornerSeamReverseOverlap: "reverseOverlap",
	CornerSeamNoOverlap:      "noOverlap",
}

// String returns the corner-seam type's wire spelling.
func (c CornerSeamType) String() string { return enumName(cornerSeamTypeNames, c) }

// ParseCornerSeamType resolves a wire spelling back to its corner-seam type. The empty string
// resolves to the gap default so an omitted type keeps its long-standing meaning.
func ParseCornerSeamType(s string) (CornerSeamType, bool) {
	if s == "" {
		return CornerSeamGap, true
	}
	return enumFromName(cornerSeamTypeNames, s)
}

// CornerSeamDefinitionType says how the seam gap is MEASURED — Inventor's CornerDefinitionTypeEnum
// (#1964). The two give the same corner only on a square miter: max-distance measures the widest
// clear span across the corner, while face-edge measures perpendicular from one wall's face to the
// other's edge, so on an oblique corner they place the relief differently.
type CornerSeamDefinitionType int32

const (
	// CornerSeamMaxDistance measures the gap as the maximum clear distance across the corner —
	// Inventor's kCornerMaxDistance, the default and the zero value.
	CornerSeamMaxDistance CornerSeamDefinitionType = iota
	// CornerSeamFaceEdgeDistance measures it from a wall's face to the neighbour's edge — kCornerFaceEdgeDistance.
	CornerSeamFaceEdgeDistance
)

var cornerSeamDefinitionTypeNames = map[CornerSeamDefinitionType]string{
	CornerSeamMaxDistance:      "maxDistance",
	CornerSeamFaceEdgeDistance: "faceEdgeDistance",
}

// String returns the definition type's wire spelling.
func (c CornerSeamDefinitionType) String() string { return enumName(cornerSeamDefinitionTypeNames, c) }

// ParseCornerSeamDefinitionType resolves a wire spelling back to its definition type. The empty
// string resolves to the max-distance default.
func ParseCornerSeamDefinitionType(s string) (CornerSeamDefinitionType, bool) {
	if s == "" {
		return CornerSeamMaxDistance, true
	}
	return enumFromName(cornerSeamDefinitionTypeNames, s)
}

// RipType names how a rip's cut is defined on its face — Inventor's RipTypeEnum (#1965). Every rip
// acts on a RipFace; the type says what draws the cut across it: two points, one point, or the
// whole face. The three are not degrees of the same cut — a single-point rip runs the face's full
// ruling through the picked point, while a point-to-point rip is bounded by the two points.
type RipType int32

const (
	// PointToPointRip cuts between two points on the face — the default and the zero value, so the
	// long-standing two-point (sketch-line) rip keeps its meaning when no type is given.
	PointToPointRip RipType = iota
	// SinglePointRip cuts the face's full extent through one point — Inventor's kSinglePointRipType,
	// the usual way to split a rolled tube open along a generator.
	SinglePointRip
	// FaceExtentsRip cuts the face's whole extent with no picked point — kFaceExtentsRipType.
	FaceExtentsRip
)

var ripTypeNames = map[RipType]string{
	PointToPointRip: "pointToPoint",
	SinglePointRip:  "singlePoint",
	FaceExtentsRip:  "faceExtents",
}

// String returns the rip type's wire spelling.
func (r RipType) String() string { return enumName(ripTypeNames, r) }

// ParseRipType resolves a wire spelling back to its rip type. The empty string resolves to the
// point-to-point default so an omitted type keeps the existing line rip.
func ParseRipType(s string) (RipType, bool) {
	if s == "" {
		return PointToPointRip, true
	}
	return enumFromName(ripTypeNames, s)
}
