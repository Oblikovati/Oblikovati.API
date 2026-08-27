// SPDX-License-Identifier: Apache-2.0

package types

// ViewLayout is how a document's open views are tiled in the viewport — how many views
// are shown at once and in what arrangement. A document always has at least one view;
// the layout decides how many of them render simultaneously (many MCAD apps show views
// in separate windows; we tile them in one viewport). This is the canonical Apache-2.0
// definition; the GPL implementation aliases it and maps it onto the tiled renderer.
type ViewLayout int32

const (
	// LayoutSingle — one view fills the viewport (the default).
	LayoutSingle ViewLayout = 0
	// LayoutTwoH — two views side by side (a vertical split: left | right).
	LayoutTwoH ViewLayout = 1
	// LayoutTwoV — two views stacked (a horizontal split: top / bottom).
	LayoutTwoV ViewLayout = 2
	// LayoutThree — three views (one large + two stacked beside it).
	LayoutThree ViewLayout = 3
	// LayoutFour — four views in a 2×2 grid (quad view).
	LayoutFour ViewLayout = 4
)

var viewLayoutNames = map[ViewLayout]string{
	LayoutSingle: "Single View",
	LayoutTwoH:   "Two Views (Horizontal)",
	LayoutTwoV:   "Two Views (Vertical)",
	LayoutThree:  "Three Views",
	LayoutFour:   "Four Views",
}

// String returns the layout's stable, user-facing name.
func (l ViewLayout) String() string {
	return enumName(viewLayoutNames, l, "viewLayout(?)")
}

// IsValid reports whether l is a defined layout.
func (l ViewLayout) IsValid() bool {
	return enumValid(viewLayoutNames, l)
}

// Tiles is how many views the layout renders at once (1–4).
func (l ViewLayout) Tiles() int {
	switch l {
	case LayoutTwoH, LayoutTwoV:
		return 2
	case LayoutThree:
		return 3
	case LayoutFour:
		return 4
	default:
		return 1
	}
}

// AllViewLayouts returns every layout in menu order — the source list for a layout picker.
func AllViewLayouts() []ViewLayout {
	return []ViewLayout{LayoutSingle, LayoutTwoH, LayoutTwoV, LayoutThree, LayoutFour}
}
