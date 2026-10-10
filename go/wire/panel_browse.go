// SPDX-License-Identifier: Apache-2.0

package wire

// TreeNode is one node of a PanelTree (a category-browser hierarchy). ID is the stable key the
// host echoes in a PanelValueChangedEvent when the node's label is clicked; Label is the display
// text; Children nests sub-nodes (empty = a leaf). Expanded is a FIRST-RENDER hint only —
// afterwards the host owns expand/collapse (twirling a node does not notify the add-in), so a
// re-sent spec does not fight the user's open/closed state.
type TreeNode struct {
	ID       string     `json:"id"`
	Label    string     `json:"label"`
	Children []TreeNode `json:"children,omitempty"`
	Expanded bool       `json:"expanded,omitempty"`
}

// TableRow is one row of a PanelTable. Key is the stable identifier the host echoes in a
// PanelValueChangedEvent when the row is clicked (kept distinct from the display Cells so the
// selection survives re-sends and filtering); Cells are the per-column display strings, in the
// column order declared by the control's TableColumns.
type TableRow struct {
	Key   string   `json:"key"`
	Cells []string `json:"cells"`
}
