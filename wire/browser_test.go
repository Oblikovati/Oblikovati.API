// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"
	"testing"
)

// TestBrowserNodeBareRoundTrip guards backward-compatibility: a plain node (no icon/menu)
// marshals without the new fields, so existing add-in panes are unchanged.
func TestBrowserNodeBareRoundTrip(t *testing.T) {
	b, _ := json.Marshal(BrowserNodeSpec{ID: "job", Label: "Job"})
	got := string(b)
	for _, leaked := range []string{"iconSVG", "menu"} {
		if contains(got, leaked) {
			t.Errorf("bare node leaked %q: %s", leaked, got)
		}
	}
}

// TestBrowserNodeIconAndMenuRoundTrip checks a node carrying an inline SVG icon and a
// right-click context menu survives a JSON round-trip with structure intact.
func TestBrowserNodeIconAndMenuRoundTrip(t *testing.T) {
	node := BrowserNodeSpec{
		ID: "op1", Label: "Profile", IconSVG: "<svg/>",
		Menu: []BrowserMenuItem{
			{ID: "edit", Label: "Edit"},
			{ID: "del", Label: "Delete", Disabled: true},
		},
	}
	var back BrowserNodeSpec
	b, _ := json.Marshal(node)
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if back.IconSVG != "<svg/>" {
		t.Errorf("iconSVG lost: %q", back.IconSVG)
	}
	if len(back.Menu) != 2 || back.Menu[0].ID != "edit" || !back.Menu[1].Disabled {
		t.Errorf("menu lost: %+v", back.Menu)
	}
}

// TestBrowserNodeEventCarriesMenuItem checks the node event can report which context-menu item
// the user chose (Gesture "menu").
func TestBrowserNodeEventCarriesMenuItem(t *testing.T) {
	ev := BrowserNodeEvent{Type: EventBrowserNode, Pane: "cam", Node: "op1", Gesture: "menu", MenuItem: "edit"}
	var back BrowserNodeEvent
	b, _ := json.Marshal(ev)
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if back.MenuItem != "edit" || back.Gesture != "menu" {
		t.Errorf("menu event lost: %+v", back)
	}
}
