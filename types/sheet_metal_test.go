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
