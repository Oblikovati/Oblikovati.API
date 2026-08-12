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
		OrdinateDimension:      "ordinate",
		ArcLengthDimension:     "arcLength",
		ForeshortenedDimension: "foreshortened", SymmetricDimension: "symmetric", SumDimension: "sum",
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

// TestDimensionToleranceTypeRoundTrip pins the tolerance methods (#1990).
func TestDimensionToleranceTypeRoundTrip(t *testing.T) {
	if DimensionToleranceType(0) != NoTolerance {
		t.Errorf("zero DimensionToleranceType = %v, want NoTolerance", DimensionToleranceType(0))
	}
	for typ, want := range map[DimensionToleranceType]string{
		NoTolerance: "none", SymmetricTolerance: "symmetric", DeviationTolerance: "deviation",
		LimitsTolerance: "limits", FitsTolerance: "fits",
	} {
		if got := typ.String(); got != want {
			t.Errorf("%v.String() = %q, want %q", typ, got, want)
		}
		if got, ok := ParseDimensionToleranceType(want); !ok || got != typ {
			t.Errorf("ParseDimensionToleranceType(%q) = (%v,%v), want (%v,true)", want, got, ok, typ)
		}
	}
	if _, ok := ParseDimensionToleranceType("gd&t"); ok {
		t.Error("unknown tolerance type should not resolve")
	}
}

// TestInspectionShapeRoundTrip pins the inspection border shapes, the empty-string/zero default,
// and rejection of an unknown spelling (#1996).
func TestInspectionShapeRoundTrip(t *testing.T) {
	if InspectionShape(0) != NoInspectionBorder {
		t.Errorf("zero InspectionShape = %v, want NoInspectionBorder", InspectionShape(0))
	}
	for shape, want := range map[InspectionShape]string{
		NoInspectionBorder: "none", AngularEndsInspectionBorder: "angular", RoundedEndsInspectionBorder: "rounded",
	} {
		if got := shape.String(); got != want {
			t.Errorf("%v.String() = %q, want %q", shape, got, want)
		}
		if got, ok := ParseInspectionShape(want); !ok || got != shape {
			t.Errorf("ParseInspectionShape(%q) = (%v,%v), want (%v,true)", want, got, ok, shape)
		}
	}
	if got, ok := ParseInspectionShape(""); !ok || got != NoInspectionBorder {
		t.Errorf(`ParseInspectionShape("") = (%v,%v), want (NoInspectionBorder,true)`, got, ok)
	}
	if _, ok := ParseInspectionShape("dashed"); ok {
		t.Error("unknown inspection shape should not resolve")
	}
}
