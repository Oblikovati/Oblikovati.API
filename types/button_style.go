// SPDX-License-Identifier: Apache-2.0

package types

// ButtonStyle is how a command renders in the ribbon, narrowed to the four styles we
// support. The control's behavior is identical; only its size and how the icon and
// display name combine differ.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases it
// (app.ButtonStyle) so existing call sites are unaffected.
type ButtonStyle uint8

const (
	// TextOnlyButton shows the display name as a plain button (the zero value).
	TextOnlyButton ButtonStyle = 0
	// SmallIconButton shows a small (16px) icon with the display name beside it —
	// the stacked rows of a ribbon panel (Move / Copy / Rotate).
	SmallIconButton ButtonStyle = 1
	// LargeIconButton shows a large (32px) icon with the name as a caption beneath.
	LargeIconButton ButtonStyle = 2
	// CompactIconButton shows a small (16px) icon only, no label — dense tool grids
	// like the sketch constraint palette, where the glyph is the whole affordance.
	CompactIconButton ButtonStyle = 3
	// SelectionListButton renders as a dropdown showing the current value with a preview — a
	// dash pattern, a colour swatch, a stroke sample — rather than a menu of commands. It is
	// the sketch Format panel's line type, colour and thickness lists, where a value is picked
	// by seeing it (Oblikovati/Oblikovati#2015).
	SelectionListButton ButtonStyle = 4
)

var buttonStyleNames = map[ButtonStyle]string{
	TextOnlyButton: "text", SmallIconButton: "small-icon", LargeIconButton: "large-icon",
	CompactIconButton: "compact-icon", SelectionListButton: "selection-list",
}

// String returns the style's stable name.
func (s ButtonStyle) String() string {
	if name, ok := buttonStyleNames[s]; ok {
		return name
	}
	return "buttonStyle(?)"
}

// ShowsIcon reports whether the style renders an icon (small or large), so a renderer
// knows to resolve the command's icon key.
func (s ButtonStyle) ShowsIcon() bool {
	return s == SmallIconButton || s == LargeIconButton || s == CompactIconButton
}
