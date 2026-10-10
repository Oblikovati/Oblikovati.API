// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestUnfoldMethodTypeRoundTrip pins each unfold method's wire spelling and that Parse
// inverts String.
func TestUnfoldMethodTypeRoundTrip(t *testing.T) {
	cases := map[UnfoldMethodType]string{
		KFactorUnfold:   "kFactor",
		BendTableUnfold: "bendTable",
		EquationUnfold:  "equation",
	}
	for m, want := range cases {
		if got := m.String(); got != want {
			t.Errorf("UnfoldMethodType(%d).String() = %q, want %q", m, got, want)
		}
		parsed, ok := ParseUnfoldMethodType(want)
		if !ok || parsed != m {
			t.Errorf("ParseUnfoldMethodType(%q) = (%d, %v), want (%d, true)", want, parsed, ok, m)
		}
	}
	if _, ok := ParseUnfoldMethodType("nonsense"); ok {
		t.Error("ParseUnfoldMethodType(\"nonsense\") = ok, want not ok")
	}
}

// TestReliefShapeRoundTrip pins each relief shape's wire spelling and that the zero value is
// ReliefRound (an unstyled rule defaults to the gentlest relief).
func TestReliefShapeRoundTrip(t *testing.T) {
	if ReliefShape(0) != ReliefRound {
		t.Errorf("zero ReliefShape = %v, want ReliefRound", ReliefShape(0))
	}
	cases := map[ReliefShape]string{
		ReliefRound:    "round",
		ReliefStraight: "straight",
		ReliefTear:     "tear",
	}
	for r, want := range cases {
		if got := r.String(); got != want {
			t.Errorf("ReliefShape(%d).String() = %q, want %q", r, got, want)
		}
		parsed, ok := ParseReliefShape(want)
		if !ok || parsed != r {
			t.Errorf("ParseReliefShape(%q) = (%d, %v), want (%d, true)", want, parsed, ok, r)
		}
	}
	// "square" was this enum's spelling for the rectangular cut before it was reconciled with
	// Inventor's BendReliefShapeEnum (#1960); a style written with it must still read back as the
	// same relief, not fail to parse.
	if got, ok := ParseReliefShape("square"); !ok || got != ReliefStraight {
		t.Errorf(`ParseReliefShape("square") = (%d, %v), want ReliefStraight`, got, ok)
	}
}

// TestCornerReliefEnums pins the corner-relief vocabularies (#1960) and their defaults: Inventor's
// Default style trims the corner to the bend, and places the relief on the bend tangents.
func TestCornerReliefEnums(t *testing.T) {
	if CornerReliefShape(0) != CornerTrimToBend {
		t.Errorf("zero CornerReliefShape = %v, want CornerTrimToBend", CornerReliefShape(0))
	}
	if CornerReliefPlacement(0) != CornerReliefAtBendTangent {
		t.Errorf("zero CornerReliefPlacement = %v, want CornerReliefAtBendTangent", CornerReliefPlacement(0))
	}
	for shape, want := range map[CornerReliefShape]string{
		CornerTrimToBend: "trimToBend", CornerRound: "round", CornerSquare: "square",
		CornerTear: "tear", CornerFullRound: "fullRound",
		CornerRoundWithRadius: "roundWithRadius", CornerIntersection: "intersection",
	} {
		if got := shape.String(); got != want {
			t.Errorf("CornerReliefShape(%d).String() = %q, want %q", shape, got, want)
		}
		if got, ok := ParseCornerReliefShape(want); !ok || got != shape {
			t.Errorf("ParseCornerReliefShape(%q) = (%d, %v), want (%d, true)", want, got, ok, shape)
		}
	}
	for placement, want := range map[CornerReliefPlacement]string{
		CornerReliefAtBendTangent: "bendTangent", CornerReliefAtBendIntersection: "bendIntersection",
		CornerReliefAtAlongBend: "alongBend",
	} {
		if got := placement.String(); got != want {
			t.Errorf("CornerReliefPlacement(%d).String() = %q, want %q", placement, got, want)
		}
		if got, ok := ParseCornerReliefPlacement(want); !ok || got != placement {
			t.Errorf("ParseCornerReliefPlacement(%q) = (%d, %v), want (%d, true)", want, got, ok, placement)
		}
	}
	if _, ok := ParseCornerReliefShape("laserWeld"); ok {
		t.Error("a weld corner-relief shape should not resolve — those are not implemented")
	}
}

// TestCornerSeamTypeRoundTrip pins each seam type's wire spelling, that the zero value and the
// empty string both mean gap (so an older seam record reads back unchanged), and that an unknown
// type is rejected (#1964).
func TestCornerSeamTypeRoundTrip(t *testing.T) {
	if CornerSeamType(0) != CornerSeamGap {
		t.Errorf("zero CornerSeamType = %v, want CornerSeamGap", CornerSeamType(0))
	}
	for seam, want := range map[CornerSeamType]string{
		CornerSeamGap: "gap", CornerSeamOverlap: "overlap",
		CornerSeamReverseOverlap: "reverseOverlap", CornerSeamNoOverlap: "noOverlap",
	} {
		if got := seam.String(); got != want {
			t.Errorf("CornerSeamType(%d).String() = %q, want %q", seam, got, want)
		}
		if got, ok := ParseCornerSeamType(want); !ok || got != seam {
			t.Errorf("ParseCornerSeamType(%q) = (%d, %v), want (%d, true)", want, got, ok, seam)
		}
	}
	if got, ok := ParseCornerSeamType(""); !ok || got != CornerSeamGap {
		t.Errorf(`ParseCornerSeamType("") = (%d, %v), want (gap, true)`, got, ok)
	}
	if _, ok := ParseCornerSeamType("weldedCorner"); ok {
		t.Error("an unknown corner-seam type should not resolve")
	}
}

// TestCornerSeamDefinitionTypeRoundTrip pins the gap-measurement vocabulary and its max-distance
// default (#1964).
func TestCornerSeamDefinitionTypeRoundTrip(t *testing.T) {
	if CornerSeamDefinitionType(0) != CornerSeamMaxDistance {
		t.Errorf("zero CornerSeamDefinitionType = %v, want CornerSeamMaxDistance", CornerSeamDefinitionType(0))
	}
	for def, want := range map[CornerSeamDefinitionType]string{
		CornerSeamMaxDistance: "maxDistance", CornerSeamFaceEdgeDistance: "faceEdgeDistance",
	} {
		if got := def.String(); got != want {
			t.Errorf("CornerSeamDefinitionType(%d).String() = %q, want %q", def, got, want)
		}
		if got, ok := ParseCornerSeamDefinitionType(want); !ok || got != def {
			t.Errorf("ParseCornerSeamDefinitionType(%q) = (%d, %v), want (%d, true)", want, got, ok, def)
		}
	}
	if got, ok := ParseCornerSeamDefinitionType(""); !ok || got != CornerSeamMaxDistance {
		t.Errorf(`ParseCornerSeamDefinitionType("") = (%d, %v), want (maxDistance, true)`, got, ok)
	}
	if _, ok := ParseCornerSeamDefinitionType("byArea"); ok {
		t.Error("an unknown corner-seam definition type should not resolve")
	}
}

// TestRipTypeRoundTrip pins each rip type's wire spelling, that the zero value and the empty
// string both mean point-to-point (so an older line rip reads back unchanged), and that an
// unknown type is rejected (#1965).
func TestRipTypeRoundTrip(t *testing.T) {
	if RipType(0) != PointToPointRip {
		t.Errorf("zero RipType = %v, want PointToPointRip", RipType(0))
	}
	for rip, want := range map[RipType]string{
		PointToPointRip: "pointToPoint", SinglePointRip: "singlePoint", FaceExtentsRip: "faceExtents",
	} {
		if got := rip.String(); got != want {
			t.Errorf("RipType(%d).String() = %q, want %q", rip, got, want)
		}
		if got, ok := ParseRipType(want); !ok || got != rip {
			t.Errorf("ParseRipType(%q) = (%d, %v), want (%d, true)", want, got, ok, rip)
		}
	}
	if got, ok := ParseRipType(""); !ok || got != PointToPointRip {
		t.Errorf(`ParseRipType("") = (%d, %v), want (pointToPoint, true)`, got, ok)
	}
	if _, ok := ParseRipType("zigzag"); ok {
		t.Error("an unknown rip type should not resolve")
	}
}

// TestLoftedFlangeOutputTypeRoundTrip pins each output type's wire spelling, the die-formed default
// (zero value and empty string), the press-brake predicate, and rejection of an unknown (#1966).
func TestLoftedFlangeOutputTypeRoundTrip(t *testing.T) {
	if LoftedFlangeOutputType(0) != DieFormedLoftedFlange {
		t.Errorf("zero LoftedFlangeOutputType = %v, want DieFormedLoftedFlange", LoftedFlangeOutputType(0))
	}
	for out, want := range map[LoftedFlangeOutputType]string{
		DieFormedLoftedFlange: "dieFormed", PressBrakeChordToleranceLoftedFlange: "pressBrakeChordTolerance",
		PressBrakeFacetAngleLoftedFlange: "pressBrakeFacetAngle", PressBrakeFacetDistanceLoftedFlange: "pressBrakeFacetDistance",
	} {
		if got := out.String(); got != want {
			t.Errorf("LoftedFlangeOutputType(%d).String() = %q, want %q", out, got, want)
		}
		if got, ok := ParseLoftedFlangeOutputType(want); !ok || got != out {
			t.Errorf("ParseLoftedFlangeOutputType(%q) = (%d, %v), want (%d, true)", want, got, ok, out)
		}
		if want != "dieFormed" && !out.IsPressBrake() {
			t.Errorf("%s should report IsPressBrake", want)
		}
	}
	if DieFormedLoftedFlange.IsPressBrake() {
		t.Error("dieFormed must not report IsPressBrake")
	}
	if got, ok := ParseLoftedFlangeOutputType(""); !ok || got != DieFormedLoftedFlange {
		t.Errorf(`ParseLoftedFlangeOutputType("") = (%d, %v), want (dieFormed, true)`, got, ok)
	}
	if _, ok := ParseLoftedFlangeOutputType("handHammered"); ok {
		t.Error("an unknown lofted-flange output type should not resolve")
	}
}

// TestPunchRepresentationTypeRoundTrip pins each punch representation's wire spelling, the default
// (zero value and empty string), and rejection of an unknown (#1968).
func TestPunchRepresentationTypeRoundTrip(t *testing.T) {
	if PunchRepresentationType(0) != DefaultPunchRepresentation {
		t.Errorf("zero PunchRepresentationType = %v, want DefaultPunchRepresentation", PunchRepresentationType(0))
	}
	for rep, want := range map[PunchRepresentationType]string{
		DefaultPunchRepresentation: "default", FormedFeaturePunchRepresentation: "formedFeature",
		Sketch2DPunchRepresentation: "sketch2D", CentermarkPunchRepresentation: "centermark",
		Sketch2DAndCentermarkPunchRepresentation: "sketch2DAndCentermark",
	} {
		if got := rep.String(); got != want {
			t.Errorf("PunchRepresentationType(%d).String() = %q, want %q", rep, got, want)
		}
		if got, ok := ParsePunchRepresentationType(want); !ok || got != rep {
			t.Errorf("ParsePunchRepresentationType(%q) = (%d, %v), want (%d, true)", want, got, ok, rep)
		}
	}
	if got, ok := ParsePunchRepresentationType(""); !ok || got != DefaultPunchRepresentation {
		t.Errorf(`ParsePunchRepresentationType("") = (%d, %v), want (default, true)`, got, ok)
	}
	if _, ok := ParsePunchRepresentationType("hologram"); ok {
		t.Error("an unknown punch representation type should not resolve")
	}
}
