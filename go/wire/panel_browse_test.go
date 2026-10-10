// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"
	"testing"
)

func TestTreeNodeRoundTrip(t *testing.T) {
	in := PanelControlSpec{
		Nodes: []TreeNode{{ID: "bearings", Label: "Bearings", Expanded: true, Children: []TreeNode{
			{ID: "iso15-6200", Label: "6200 series"},
		}}},
	}
	var out PanelControlSpec
	mustReJSON(t, in, &out)
	if out.Nodes[0].ID != "bearings" || !out.Nodes[0].Expanded ||
		out.Nodes[0].Children[0].ID != "iso15-6200" {
		t.Fatalf("tree round-trip lost data: %+v", out.Nodes)
	}
}

func TestTableRoundTrip(t *testing.T) {
	in := PanelControlSpec{
		TableColumns: []string{"d", "D", "B"},
		TableRows:    []TableRow{{Key: "d=10,D=30", Cells: []string{"10", "30", "9"}}},
		Value:        "d=10,D=30",
	}
	var out PanelControlSpec
	mustReJSON(t, in, &out)
	if len(out.TableColumns) != 3 || out.TableRows[0].Key != "d=10,D=30" ||
		out.TableRows[0].Cells[1] != "30" || out.Value != "d=10,D=30" {
		t.Fatalf("table round-trip lost data: %+v", out)
	}
}

func mustReJSON(t *testing.T, in, out any) {
	t.Helper()
	b, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if err := json.Unmarshal(b, out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
}
