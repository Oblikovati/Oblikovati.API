// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestBaseViewOrientationRoundTrip(t *testing.T) {
	if BaseViewOrientation(0) != BaseViewFront {
		t.Errorf("zero BaseViewOrientation = %v, want BaseViewFront", BaseViewOrientation(0))
	}
	cases := map[BaseViewOrientation]string{
		BaseViewFront: "front", BaseViewTop: "top", BaseViewRight: "right", BaseViewBack: "back",
		BaseViewLeft: "left", BaseViewBottom: "bottom", BaseViewIso: "iso",
	}
	for o, want := range cases {
		if got := o.String(); got != want {
			t.Errorf("BaseViewOrientation(%d).String() = %q, want %q", o, got, want)
		}
		if parsed, ok := ParseBaseViewOrientation(want); !ok || parsed != o {
			t.Errorf("ParseBaseViewOrientation(%q) = (%d, %v), want (%d, true)", want, parsed, ok, o)
		}
	}
	if _, ok := ParseBaseViewOrientation("sideways"); ok {
		t.Error("ParseBaseViewOrientation(\"sideways\") = ok, want not ok")
	}
}

func TestDrawingViewStyleRoundTrip(t *testing.T) {
	if DrawingViewStyle(0) != HiddenLineViewStyle {
		t.Errorf("zero DrawingViewStyle = %v, want HiddenLineViewStyle", DrawingViewStyle(0))
	}
	cases := map[DrawingViewStyle]string{
		HiddenLineViewStyle: "hiddenLine", WireframeViewStyle: "wireframe", ShadedViewStyle: "shaded",
	}
	for s, want := range cases {
		if got := s.String(); got != want {
			t.Errorf("DrawingViewStyle(%d).String() = %q, want %q", s, got, want)
		}
		if parsed, ok := ParseDrawingViewStyle(want); !ok || parsed != s {
			t.Errorf("ParseDrawingViewStyle(%q) = (%d, %v), want (%d, true)", want, parsed, ok, s)
		}
	}
}

func TestProjectionDirectionRoundTrip(t *testing.T) {
	cases := map[ProjectionDirection]string{
		ProjectRight: "right", ProjectLeft: "left", ProjectUp: "up", ProjectDown: "down",
	}
	for d, want := range cases {
		if got := d.String(); got != want {
			t.Errorf("ProjectionDirection(%d).String() = %q, want %q", d, got, want)
		}
		if parsed, ok := ParseProjectionDirection(want); !ok || parsed != d {
			t.Errorf("ParseProjectionDirection(%q) = (%d, %v), want (%d, true)", want, parsed, ok, d)
		}
	}
}
