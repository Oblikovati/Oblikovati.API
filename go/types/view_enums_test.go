// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestViewEnumIdsAreStable pins the frozen ids for every view-collection / orientation enum.
func TestViewEnumIdsAreStable(t *testing.T) {
	vt := map[ViewTypeEnum]int32{
		UnknownViewType: 9217, GraphicsViewType: 9218, BrowserViewType: 9219, NoteBookViewType: 9220,
	}
	for v, id := range vt {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	op := map[ViewOperationTypeEnum]int32{
		RotateViewOperation: 30209, PanViewOperation: 30210, ZoomViewOperation: 30211,
	}
	for v, id := range op {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	orbit := map[OrbitTypeEnum]int32{FreeOrbit: 86017, ConstrainedOrbit: 86018}
	for v, id := range orbit {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	tile := map[ViewTileTypeEnum]int32{
		UnknownViewTileType: 117761, ArrangeViewTileType: 117762,
		HorizontalViewTileType: 117763, VerticalViewTileType: 117764,
	}
	for v, id := range tile {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	// Spot-check the orientation range endpoints (the full set is large).
	orient := map[ViewOrientationTypeEnum]int32{
		DefaultViewOrientation: 10753, FrontViewOrientation: 10764,
		IsoTopRightViewOrientation: 10759, FlatBacksidePivot180ViewOrientation: 10773,
	}
	for v, id := range orient {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
}

// TestViewEnumNamesAndValidity checks every All* entry is valid with a non-placeholder name,
// and the zero value is reported invalid.
func TestViewEnumNamesAndValidity(t *testing.T) {
	assertNamed(t, "ViewType", "viewType(?)", AllViewTypes())
	assertNamed(t, "ViewOperation", "viewOperation(?)", AllViewOperations())
	assertNamed(t, "OrbitType", "orbitType(?)", AllOrbitTypes())
	assertNamed(t, "ViewTileType", "viewTileType(?)", AllViewTileTypes())
	assertNamed(t, "ViewOrientation", "viewOrientation(?)", AllViewOrientations())
	if ViewTypeEnum(0).IsValid() || OrbitTypeEnum(0).IsValid() || ViewOrientationTypeEnum(0).IsValid() {
		t.Errorf("zero value must be invalid for these id-based enums")
	}
}
