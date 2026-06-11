// SPDX-License-Identifier: Apache-2.0

package wire

// BrowserNodeSpec is one node of an add-in browser pane: a label with an optional
// icon key and children. ID must be unique within its pane — it names the node in
// the [BrowserNodeEvent] the add-in receives when the user interacts with it.
// Expanded is the node's initial state; the user's expansion is kept across
// re-sets of the pane (the spec declares content, not view state).
type BrowserNodeSpec struct {
	ID       string            `json:"id"`
	Label    string            `json:"label"`
	Icon     string            `json:"icon,omitempty"`
	Expanded bool              `json:"expanded,omitempty"`
	Children []BrowserNodeSpec `json:"children,omitempty"`
}

// BrowserPaneSpec is one add-in browser pane: a named tree shown alongside the
// built-in Model pane (the ClientBrowserNodeDefinition equivalent, M05-F03 #256).
// Setting a pane replaces its whole tree — the declared-bulk-state model the
// clientGraphics groups established — so an add-in never patches node-by-node.
type BrowserPaneSpec struct {
	ID    string            `json:"id"`
	Title string            `json:"title"`
	Nodes []BrowserNodeSpec `json:"nodes,omitempty"`
}

// SetBrowserPaneArgs is the request of [MethodBrowserSetPane]: create the pane or
// replace its content if it exists.
type SetBrowserPaneArgs struct {
	Pane BrowserPaneSpec `json:"pane"`
}

// DeleteBrowserPaneArgs is the request of [MethodBrowserDeletePane].
type DeleteBrowserPaneArgs struct {
	ID string `json:"id"`
}

// ListBrowserPanesResult is the response of [MethodBrowserListPanes]: every add-in
// pane in creation order (the built-in Model pane is not listed — it is the
// document tree, served by model.tree).
type ListBrowserPanesResult struct {
	Panes []BrowserPaneSpec `json:"panes"`
}

// BrowserNodeEvent is the push event (type [EventBrowserNode]) an add-in receives
// when the user interacts with one of its pane nodes. Gesture is "select",
// "double", "expand" or "collapse".
type BrowserNodeEvent struct {
	Type    string `json:"type"` // always EventBrowserNode
	Pane    string `json:"pane"`
	Node    string `json:"node"`
	Gesture string `json:"gesture"`
}
