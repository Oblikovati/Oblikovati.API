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
		ReliefRound:  "round",
		ReliefSquare: "square",
		ReliefTear:   "tear",
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
}
