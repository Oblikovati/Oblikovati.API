// SPDX-License-Identifier: Apache-2.0

package wire

// BrowserNodeSpec is one node of an add-in browser pane: a label with an optional
// icon and children, drawn like a built-in document-tree node. ID must be unique
// within its pane — it names the node in the [BrowserNodeEvent] the add-in receives
// when the user interacts with it. Expanded is the node's initial state; the user's
// expansion is kept across re-sets of the pane (the spec declares content, not view
// state). IconSVG is an inline themed glyph (the ribbon's sentinel-colour SVG
// convention) the host rasterises beside the label, so an add-in node looks like a
// document node. Menu is the node's right-click context menu; choosing an item
// reports a [BrowserNodeEvent] with Gesture "menu" and MenuItem set.
type BrowserNodeSpec struct {
	ID       string            `json:"id"`
	Label    string            `json:"label"`
	Icon     string            `json:"icon,omitempty"`    // built-in icon key (host nodes)
	IconSVG  string            `json:"iconSVG,omitempty"` // inline themed glyph (add-in nodes)
	Expanded bool              `json:"expanded,omitempty"`
	Menu     []BrowserMenuItem `json:"menu,omitempty"`
	Children []BrowserNodeSpec `json:"children,omitempty"`
}

// BrowserMenuItem is one entry of a node's right-click context menu. ID names the item
// in the [BrowserNodeEvent] when chosen; Label is what's shown; Disabled greys it out
// (the zero value is enabled, the common case).
type BrowserMenuItem struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Disabled bool   `json:"disabled,omitempty"`
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
// when the user interacts with one of its pane nodes. Gesture is "select", "double",
// "expand", "collapse", or "menu". For "menu", MenuItem is the chosen item's ID.
type BrowserNodeEvent struct {
	Type     string `json:"type"` // always EventBrowserNode
	Pane     string `json:"pane"`
	Node     string `json:"node"`
	Gesture  string `json:"gesture"`
	MenuItem string `json:"menuItem,omitempty"` // chosen context-menu item id (Gesture "menu")
}
