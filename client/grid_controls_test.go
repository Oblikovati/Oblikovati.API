// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"oblikovati.org/api/types"
)

func TestTrackConstructors(t *testing.T) {
	if got := TrackAuto(); got.Kind != types.GridTrackAuto {
		t.Errorf("TrackAuto = %+v", got)
	}
	if got := TrackFixed(120); got.Kind != types.GridTrackFixed || got.Value != 120 {
		t.Errorf("TrackFixed = %+v", got)
	}
	if got := TrackFr(2); got.Kind != types.GridTrackFraction || got.Value != 2 {
		t.Errorf("TrackFr = %+v", got)
	}
	mm := TrackMinMax(TrackFr(1), 80, 240)
	if mm.Kind != types.GridTrackFraction || mm.MinPx != 80 || mm.MaxPx != 240 {
		t.Errorf("TrackMinMax = %+v", mm)
	}
}

func TestPanelGridBuildsContainer(t *testing.T) {
	g := PanelGrid("form", []types.GridTrack{TrackAuto(), TrackFr(1)}, 6, 4,
		PanelLabel("l", "Start depth"),
		PanelTextBox("start", "", "0"),
	)
	if g.Kind != types.PanelGrid || g.ID != "form" {
		t.Fatalf("grid header = %+v", g)
	}
	if len(g.Columns) != 2 || g.ColumnGap != 6 || g.RowGap != 4 {
		t.Errorf("grid tracks/gaps = %+v", g)
	}
	if len(g.Children) != 2 {
		t.Errorf("children = %d, want 2", len(g.Children))
	}
}

func TestPlaceAtStampsCell(t *testing.T) {
	c := PlaceAt(PanelButton("ok", "OK", "cmd"), 0, 2)
	if c.Cell == nil || c.Cell.Col != 0 || c.Cell.ColSpan != 2 {
		t.Errorf("PlaceAt cell = %+v", c.Cell)
	}
}

func TestPanelGroupAndTabs(t *testing.T) {
	grp := PanelGroup("stockgrp", "Stock", PanelLabel("x", "hi"))
	if grp.Kind != types.PanelGroup || grp.Title != "Stock" || len(grp.Children) != 1 {
		t.Errorf("group = %+v", grp)
	}
	tabs := PanelTabs("jobedit", PanelTab("Setup", PanelLabel("s", "stock")))
	if tabs.Kind != types.PanelTabs || len(tabs.Children) != 1 {
		t.Fatalf("tabs = %+v", tabs)
	}
	if tabs.Children[0].Title != "Setup" {
		t.Errorf("tab caption lost: %+v", tabs.Children[0])
	}
}
