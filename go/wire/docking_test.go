// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
)

// TestPanelControlSpecFlatOmitsNestingFields guards backward-compatibility: a plain
// leaf control must marshal exactly as before — none of the new nesting fields appear.
func TestPanelControlSpecFlatOmitsNestingFields(t *testing.T) {
	b, err := json.Marshal(PanelControlSpec{Kind: types.PanelButton, ID: "ok", Text: "OK"})
	if err != nil {
		t.Fatal(err)
	}
	got := string(b)
	for _, leaked := range []string{"children", "columns", "columnGap", "rowGap", "cell", "title"} {
		if contains(got, leaked) {
			t.Errorf("flat control leaked nesting field %q: %s", leaked, got)
		}
	}
}

// TestPanelGridRoundTrip checks a nested grid (columns + gaps + placed children)
// survives a JSON round-trip with structure intact — the wire shape add-ins depend on.
func TestPanelGridRoundTrip(t *testing.T) {
	grid := PanelControlSpec{
		Kind:      types.PanelGrid,
		ID:        "form",
		Columns:   []types.GridTrack{{Kind: types.GridTrackAuto}, {Kind: types.GridTrackFraction, Value: 1}},
		ColumnGap: 6, RowGap: 4,
		Children: []PanelControlSpec{
			{Kind: types.PanelLabel, ID: "l1", Text: "Start depth"},
			{Kind: types.PanelTextBox, ID: "start", Value: "0"},
			{Kind: types.PanelButton, ID: "full", Text: "Full width",
				Cell: &types.GridCell{Col: 0, ColSpan: 2}},
		},
	}
	var back PanelControlSpec
	b, err := json.Marshal(grid)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if back.Kind != types.PanelGrid || len(back.Columns) != 2 || back.ColumnGap != 6 {
		t.Fatalf("grid header lost: %+v", back)
	}
	if len(back.Children) != 3 {
		t.Fatalf("children lost: %d", len(back.Children))
	}
	spanned := back.Children[2]
	if spanned.Cell == nil || spanned.Cell.ColSpan != 2 {
		t.Errorf("cell placement lost: %+v", spanned.Cell)
	}
}

// TestPanelTabsNest checks a tabs container holding grid panes round-trips two levels deep.
func TestPanelTabsNest(t *testing.T) {
	tabs := PanelControlSpec{
		Kind: types.PanelTabs, ID: "jobedit",
		Children: []PanelControlSpec{
			{Kind: types.PanelGroup, Title: "Setup", Children: []PanelControlSpec{
				{Kind: types.PanelLabel, Text: "Stock"},
			}},
		},
	}
	b, _ := json.Marshal(tabs)
	var back PanelControlSpec
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if len(back.Children) != 1 || back.Children[0].Title != "Setup" ||
		len(back.Children[0].Children) != 1 {
		t.Errorf("nested tabs/group lost: %+v", back)
	}
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
