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
		HiddenLineRemovedViewStyle: "hiddenLineRemoved", FromBaseViewStyle: "fromBase",
		ShadedHiddenLineViewStyle: "shadedHiddenLine",
	}
	for s, want := range cases {
		if got := s.String(); got != want {
			t.Errorf("DrawingViewStyle(%d).String() = %q, want %q", s, got, want)
		}
		if parsed, ok := ParseDrawingViewStyle(want); !ok || parsed != s {
			t.Errorf("ParseDrawingViewStyle(%q) = (%d, %v), want (%d, true)", want, parsed, ok, s)
		}
	}
	// Only the removed style drops hidden edges; the rest keep them (dashed or, for wireframe, visible).
	if !HiddenLineRemovedViewStyle.RemovesHiddenEdges() || HiddenLineViewStyle.RemovesHiddenEdges() || WireframeViewStyle.RemovesHiddenEdges() {
		t.Error("RemovesHiddenEdges should be true only for HiddenLineRemoved")
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

// TestSectionViewTypeRoundTrip checks the section partial-cut kinds spell and parse back, the
// zero value and empty string both resolve to NoSectionView, and an unknown spelling is rejected
// (#1982).
func TestSectionViewTypeRoundTrip(t *testing.T) {
	if SectionViewType(0) != NoSectionView {
		t.Errorf("zero SectionViewType = %v, want NoSectionView", SectionViewType(0))
	}
	cases := map[SectionViewType]string{
		NoSectionView: "none", QuarterSectionView: "quarter",
		HalfSectionView: "half", ThreeQuarterSectionView: "threeQuarter",
	}
	for typ, want := range cases {
		if got := typ.String(); got != want {
			t.Errorf("SectionViewType(%d).String() = %q, want %q", typ, got, want)
		}
		if parsed, ok := ParseSectionViewType(want); !ok || parsed != typ {
			t.Errorf("ParseSectionViewType(%q) = (%d, %v), want (%d, true)", want, parsed, ok, typ)
		}
	}
	if parsed, ok := ParseSectionViewType(""); !ok || parsed != NoSectionView {
		t.Errorf(`ParseSectionViewType("") = (%d, %v), want (NoSectionView, true)`, parsed, ok)
	}
	if _, ok := ParseSectionViewType("octant"); ok {
		t.Error("ParseSectionViewType(octant) = ok, want rejected")
	}
}

// TestDrawingViewAlignmentRoundTrip pins the alignment spellings, the empty/zero default (free) and
// rejection of an unknown spelling (#1988).
func TestDrawingViewAlignmentRoundTrip(t *testing.T) {
	cases := map[DrawingViewAlignment]string{
		InPositionViewAlignment: "inPosition", HorizontalViewAlignment: "horizontal", VerticalViewAlignment: "vertical",
	}
	for a, want := range cases {
		if got := a.String(); got != want {
			t.Errorf("DrawingViewAlignment(%d).String() = %q, want %q", a, got, want)
		}
		if parsed, ok := ParseDrawingViewAlignment(want); !ok || parsed != a {
			t.Errorf("ParseDrawingViewAlignment(%q) = (%d, %v), want (%d, true)", want, parsed, ok, a)
		}
	}
	if parsed, ok := ParseDrawingViewAlignment(""); !ok || parsed != InPositionViewAlignment {
		t.Errorf(`ParseDrawingViewAlignment("") = (%d, %v), want (InPositionViewAlignment, true)`, parsed, ok)
	}
	if _, ok := ParseDrawingViewAlignment("diagonal"); ok {
		t.Error("ParseDrawingViewAlignment(diagonal) = ok, want rejected")
	}
}

// TestViewJustificationRoundTrip pins the justification spellings, the empty/zero default (centered)
// and rejection of an unknown spelling (#1988).
func TestViewJustificationRoundTrip(t *testing.T) {
	cases := map[ViewJustification]string{
		CenteredViewJustification: "centered", FixedViewJustification: "fixed",
	}
	for j, want := range cases {
		if got := j.String(); got != want {
			t.Errorf("ViewJustification(%d).String() = %q, want %q", j, got, want)
		}
		if parsed, ok := ParseViewJustification(want); !ok || parsed != j {
			t.Errorf("ParseViewJustification(%q) = (%d, %v), want (%d, true)", want, parsed, ok, j)
		}
	}
	if parsed, ok := ParseViewJustification(""); !ok || parsed != CenteredViewJustification {
		t.Errorf(`ParseViewJustification("") = (%d, %v), want (CenteredViewJustification, true)`, parsed, ok)
	}
	if _, ok := ParseViewJustification("left"); ok {
		t.Error("ParseViewJustification(left) = ok, want rejected")
	}
}

// TestDrawingEdgeTypeRoundTrip pins every edge-type spelling, the empty/zero default and rejection of
// an unknown spelling (#1984).
func TestDrawingEdgeTypeRoundTrip(t *testing.T) {
	cases := map[DrawingEdgeType]string{
		UnknownDrawingEdge: "unknown", TangentDrawingEdge: "tangent", ThreadDrawingEdge: "thread",
		BendUpDrawingEdge: "bendUp", BendDownDrawingEdge: "bendDown", BendExtentDrawingEdge: "bendExtent",
		PunchDrawingEdge: "punch", ContourRollDrawingEdge: "contourRoll",
	}
	for typ, want := range cases {
		if got := typ.String(); got != want {
			t.Errorf("DrawingEdgeType(%d).String() = %q, want %q", typ, got, want)
		}
		if parsed, ok := ParseDrawingEdgeType(want); !ok || parsed != typ {
			t.Errorf("ParseDrawingEdgeType(%q) = (%d, %v), want (%d, true)", want, parsed, ok, typ)
		}
	}
	if parsed, ok := ParseDrawingEdgeType(""); !ok || parsed != UnknownDrawingEdge {
		t.Errorf(`ParseDrawingEdgeType("") = (%d, %v), want (UnknownDrawingEdge, true)`, parsed, ok)
	}
	if _, ok := ParseDrawingEdgeType("interference"); ok {
		t.Error("ParseDrawingEdgeType(interference) = ok, want rejected")
	}
}

// TestChamferBendNoteKindsRoundTrip pins the model-derived note kinds' wire spellings and their
// round-trip, and rejection of an unknown spelling (#1995).
func TestChamferBendNoteKindsRoundTrip(t *testing.T) {
	cases := map[DrawingAnnotationKind]string{
		ChamferNoteAnnotation: "chamferNote",
		BendNoteAnnotation:    "bendNote",
		HoleNoteAnnotation:    "holeNote",
	}
	for kind, want := range cases {
		if got := kind.String(); got != want {
			t.Errorf("DrawingAnnotationKind(%d).String() = %q, want %q", kind, got, want)
		}
		if parsed, ok := ParseDrawingAnnotationKind(want); !ok || parsed != kind {
			t.Errorf("ParseDrawingAnnotationKind(%q) = (%d, %v), want (%d, true)", want, parsed, ok, kind)
		}
	}
	if _, ok := ParseDrawingAnnotationKind("filletNote"); ok {
		t.Error("ParseDrawingAnnotationKind(filletNote) = ok, want rejected")
	}
}

// TestCropBreakMarkLineTypeRoundTrip pins the crop break-mark types, the empty/zero default and
// rejection of an unknown spelling (#1987).
func TestCropBreakMarkLineTypeRoundTrip(t *testing.T) {
	if CropBreakMarkLineType(0) != NoCropBreakMark {
		t.Errorf("zero CropBreakMarkLineType = %v, want NoCropBreakMark", CropBreakMarkLineType(0))
	}
	for typ, want := range map[CropBreakMarkLineType]string{
		NoCropBreakMark: "none", ContinuousCropBreakMark: "continuous", ZigzagCropBreakMark: "zigzag",
	} {
		if got := typ.String(); got != want {
			t.Errorf("%v.String() = %q, want %q", typ, got, want)
		}
		if got, ok := ParseCropBreakMarkLineType(want); !ok || got != typ {
			t.Errorf("ParseCropBreakMarkLineType(%q) = (%v,%v), want (%v,true)", want, got, ok, typ)
		}
	}
	if got, ok := ParseCropBreakMarkLineType(""); !ok || got != NoCropBreakMark {
		t.Errorf(`ParseCropBreakMarkLineType("") = (%v,%v), want (NoCropBreakMark,true)`, got, ok)
	}
	if _, ok := ParseCropBreakMarkLineType("dotted"); ok {
		t.Error("unknown crop break-mark type should not resolve")
	}
}

// TestDrawingViewOverlayRoundTrip the overlay view type round-trips (#1986).
func TestDrawingViewOverlayRoundTrip(t *testing.T) {
	if got := DrawingViewOverlay.String(); got != "overlay" {
		t.Errorf("DrawingViewOverlay.String() = %q, want overlay", got)
	}
	if got, ok := ParseDrawingViewType("overlay"); !ok || got != DrawingViewOverlay {
		t.Errorf("ParseDrawingViewType(overlay) = (%v,%v), want (overlay,true)", got, ok)
	}
}
