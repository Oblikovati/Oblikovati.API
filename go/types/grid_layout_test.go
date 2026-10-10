// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestGridContainerKindsAreStable pins the container-kind values added for the
// nesting/grid layout extension (ADR-0019). They append after slider(8), so they
// must never be renumbered or existing add-in registrations shift meaning.
func TestGridContainerKindsAreStable(t *testing.T) {
	if PanelGrid != 9 || PanelGroup != 10 || PanelTabs != 11 {
		t.Fatalf("container kinds = %d,%d,%d, want 9,10,11",
			PanelGrid, PanelGroup, PanelTabs)
	}
	want := map[PanelControlKind]string{
		PanelGrid: "grid", PanelGroup: "group", PanelTabs: "tabs",
	}
	for k, name := range want {
		if k.String() != name {
			t.Errorf("%d.String() = %q, want %q", k, k.String(), name)
		}
	}
}

// TestGridTrackKindsAreStable pins the track-sizing enum: auto is the zero value
// (a track with no declared size sizes to content).
func TestGridTrackKindsAreStable(t *testing.T) {
	if GridTrackAuto != 0 || GridTrackFixed != 1 || GridTrackFraction != 2 {
		t.Fatalf("track kinds = %d,%d,%d, want 0,1,2",
			GridTrackAuto, GridTrackFixed, GridTrackFraction)
	}
}

// TestGridTrackCarriesSizeAndClamp checks a track holds its sizing value plus an
// optional minmax clamp, the faithful subset of CSS grid track sizing.
func TestGridTrackCarriesSizeAndClamp(t *testing.T) {
	fr := GridTrack{Kind: GridTrackFraction, Value: 2}
	if fr.Kind != GridTrackFraction || fr.Value != 2 {
		t.Errorf("fraction track = %+v", fr)
	}
	clamped := GridTrack{Kind: GridTrackFraction, Value: 1, MinPx: 80, MaxPx: 240}
	if clamped.MinPx != 80 || clamped.MaxPx != 240 {
		t.Errorf("minmax clamp lost: %+v", clamped)
	}
}

// TestGridCellPlacement checks explicit placement carries column + span; the zero
// value means auto-flow (Col 0 with the convention that a nil *GridCell flows).
func TestGridCellPlacement(t *testing.T) {
	c := GridCell{Col: 1, ColSpan: 2}
	if c.Col != 1 || c.ColSpan != 2 {
		t.Errorf("cell placement = %+v", c)
	}
}
