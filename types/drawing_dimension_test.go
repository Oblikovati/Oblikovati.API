// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestDrawingDimensionTypeRoundTrip: every dimension type has a stable wire spelling that parses
// back, the zero value is AlignedDimension, and an unknown spelling is rejected.
func TestDrawingDimensionTypeRoundTrip(t *testing.T) {
	if DrawingDimensionType(0) != AlignedDimension {
		t.Errorf("zero DrawingDimensionType = %v, want AlignedDimension", DrawingDimensionType(0))
	}
	want := map[DrawingDimensionType]string{
		AlignedDimension: "aligned", HorizontalDimension: "horizontal", VerticalDimension: "vertical",
		RadiusDimension: "radius", DiameterDimension: "diameter", AngularDimension: "angular",
		OrdinateDimension:  "ordinate",
		ArcLengthDimension: "arcLength",
	}
	for typ, spelling := range want {
		if got := typ.String(); got != spelling {
			t.Errorf("%v.String() = %q, want %q", typ, got, spelling)
		}
		parsed, ok := ParseDrawingDimensionType(spelling)
		if !ok || parsed != typ {
			t.Errorf("ParseDrawingDimensionType(%q) = (%v,%v), want (%v,true)", spelling, parsed, ok, typ)
		}
	}
	if _, ok := ParseDrawingDimensionType("radial"); ok {
		t.Error("ParseDrawingDimensionType should reject an unknown spelling")
	}
}
