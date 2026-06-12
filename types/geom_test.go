// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"
	"math"
	"testing"
)

// TestGeometryJSONIsArrayCompatible pins the wire encoding: the typed geometry
// must serialize EXACTLY like the ad-hoc [3]float64 / [16]float64 fields it
// replaced, so frozen payloads keep parsing on both sides.
func TestGeometryJSONIsArrayCompatible(t *testing.T) {
	cases := []struct {
		name  string
		value interface{ MarshalJSON() ([]byte, error) }
		want  string
	}{
		{"Point", NewPoint(1, 2.5, -3), "[1,2.5,-3]"},
		{"Point2d", NewPoint2d(4, -5), "[4,-5]"},
		{"Vector", NewVector(0, 1, 0), "[0,1,0]"},
		{"Vector2d", NewVector2d(1, 1), "[1,1]"},
		{"UnitVector", UnitVector{0, 0, 1}, "[0,0,1]"},
		{"Matrix", IdentityMatrix(), "[1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1]"},
		{"Matrix2d", IdentityMatrix2d(), "[1,0,0,0,1,0,0,0,1]"},
	}
	for _, tc := range cases {
		got, err := json.Marshal(tc.value)
		if err != nil {
			t.Fatalf("%s marshal: %v", tc.name, err)
		}
		if string(got) != tc.want {
			t.Errorf("%s JSON = %s, want %s", tc.name, got, tc.want)
		}
	}
}

func TestGeometryJSONRoundTrip(t *testing.T) {
	var p Point
	if err := json.Unmarshal([]byte("[1,2,3]"), &p); err != nil || p != NewPoint(1, 2, 3) {
		t.Fatalf("Point round trip = (%+v, %v)", p, err)
	}
	var m Matrix
	if err := json.Unmarshal([]byte("[1,0,0,4, 0,1,0,5, 0,0,1,6, 0,0,0,1]"), &m); err != nil {
		t.Fatalf("Matrix unmarshal: %v", err)
	}
	if m.Translation() != NewVector(4, 5, 6) {
		t.Errorf("translation = %+v, want (4,5,6)", m.Translation())
	}
	var u UnitVector
	if err := json.Unmarshal([]byte("[0,3,0]"), &u); err != nil || u != (UnitVector{0, 1, 0}) {
		t.Fatalf("UnitVector should renormalize: (%+v, %v)", u, err)
	}
	if err := json.Unmarshal([]byte("[0,0,0]"), &u); err == nil {
		t.Error("a zero unit vector must be rejected")
	}
	if err := json.Unmarshal([]byte("[1,2]"), &p); err == nil {
		t.Error("a wrong-length array must be rejected with the offending payload named")
	}
}

func TestVectorAndPointAlgebra(t *testing.T) {
	a, b := NewPoint(1, 1, 1), NewPoint(4, 5, 1)
	if a.VectorTo(b) != NewVector(3, 4, 0) {
		t.Errorf("VectorTo = %+v", a.VectorTo(b))
	}
	if a.DistanceTo(b) != 5 {
		t.Errorf("DistanceTo = %v, want 5", a.DistanceTo(b))
	}
	if a.Midpoint(b) != NewPoint(2.5, 3, 1) {
		t.Errorf("Midpoint = %+v", a.Midpoint(b))
	}
	v := NewVector(1, 0, 0).Cross(NewVector(0, 1, 0))
	if v != NewVector(0, 0, 1) {
		t.Errorf("Cross = %+v, want +Z", v)
	}
	if NewVector2d(1, 0).Cross(NewVector2d(0, 1)) != 1 {
		t.Error("2D cross should be the signed area term")
	}
	if _, err := NewVector(0, 0, 0).AsUnit(); err == nil {
		t.Error("normalizing a zero vector must error")
	}
	u, _ := NewVector(0, 0, 9).AsUnit()
	if u != (UnitVector{0, 0, 1}) {
		t.Errorf("AsUnit = %+v", u)
	}
}

func TestMatrixTransforms(t *testing.T) {
	zAxis := UnitVector{0, 0, 1}
	rot := RotationMatrix(math.Pi/2, zAxis, NewPoint(0, 0, 0))
	got := rot.TransformPoint(NewPoint(1, 0, 0))
	if !got.IsEqualTo(NewPoint(0, 1, 0), 1e-12) {
		t.Errorf("90° about Z: %+v, want (0,1,0)", got)
	}

	tr := TranslationMatrix(NewVector(10, 0, 0))
	combined := tr.Mul(rot) // rotate, then translate
	got = combined.TransformPoint(NewPoint(1, 0, 0))
	if !got.IsEqualTo(NewPoint(10, 1, 0), 1e-12) {
		t.Errorf("T·R: %+v, want (10,1,0)", got)
	}
	// Vectors ignore translation.
	if v := tr.TransformVector(NewVector(1, 2, 3)); v != NewVector(1, 2, 3) {
		t.Errorf("TransformVector under translation = %+v", v)
	}

	inv, err := combined.Invert()
	if err != nil {
		t.Fatalf("Invert: %v", err)
	}
	round := inv.Mul(combined).TransformPoint(NewPoint(7, -2, 3))
	if !round.IsEqualTo(NewPoint(7, -2, 3), 1e-9) {
		t.Errorf("M⁻¹·M ≠ I: %+v", round)
	}

	cs := CoordinateSystemMatrix(NewPoint(1, 2, 3), UnitVector{0, 1, 0}, UnitVector{0, 0, 1}, UnitVector{1, 0, 0})
	if got := cs.TransformPoint(NewPoint(1, 0, 0)); !got.IsEqualTo(NewPoint(1, 3, 3), 1e-12) {
		t.Errorf("frame map: %+v, want origin + yAxis", got)
	}
}

func TestMatrix2dTransforms(t *testing.T) {
	rot := RotationMatrix2d(math.Pi/2, NewPoint2d(0, 0))
	if got := rot.TransformPoint(NewPoint2d(1, 0)); !got.IsEqualTo(NewPoint2d(0, 1), 1e-12) {
		t.Errorf("90°: %+v, want (0,1)", got)
	}
	tr := TranslationMatrix2d(NewVector2d(3, 4))
	if got := tr.TransformPoint(NewPoint2d(1, 1)); got != NewPoint2d(4, 5) {
		t.Errorf("translate: %+v", got)
	}
}

func TestBoxes(t *testing.T) {
	b := NewEmptyBox()
	if !b.IsEmpty() {
		t.Fatal("NewEmptyBox should be empty")
	}
	b = b.Extend(NewPoint(1, 1, 1)).Extend(NewPoint(-1, 2, 0))
	if b.IsEmpty() || !b.Contains(NewPoint(0, 1.5, 0.5)) || b.Contains(NewPoint(2, 0, 0)) {
		t.Errorf("box bounds wrong: %+v", b)
	}
	if b.Center() != NewPoint(0, 1.5, 0.5) {
		t.Errorf("center = %+v", b.Center())
	}

	ob := OrientedBox{
		Corner: NewPoint(0, 0, 0),
		XAxis:  UnitVector{1, 0, 0}, YAxis: UnitVector{0, 1, 0}, ZAxis: UnitVector{0, 0, 1},
		Extents: NewVector(2, 3, 4),
	}
	corners := ob.Corners()
	if corners[0] != NewPoint(0, 0, 0) || corners[7] != NewPoint(2, 3, 4) {
		t.Errorf("oriented corners = %+v … %+v", corners[0], corners[7])
	}
}

// TestGeometryEnumValuesAreFrozen pins the reference enum values.
func TestGeometryEnumValuesAreFrozen(t *testing.T) {
	if LineCurve != 5122 || PolylineCurve != 5129 || HelixCurve != 5130 {
		t.Error("CurveType values drifted from the frozen block 5121…")
	}
	if LineCurve2d != 5250 || PolylineCurve2d != 5257 {
		t.Error("Curve2dType values drifted from the frozen block 5249…")
	}
	if PlaneSurface != 5890 || BSplineSurfaceKind != 5897 {
		t.Error("SurfaceType values drifted from the frozen block 5889…")
	}
	if CurveFormNURBS != 1 || SurfaceFormProceduralToNURBS != 16 {
		t.Error("geometry-form values drifted")
	}
}
