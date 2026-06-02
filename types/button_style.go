// SPDX-License-Identifier: Apache-2.0

package types

// ButtonStyle is how a command renders in the ribbon — Inventor's ButtonDisplayType,
// narrowed to the three styles we support. The control's behavior is identical; only
// its size and whether it shows an icon differ.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases it
// (app.ButtonStyle) so existing call sites are unaffected.
type ButtonStyle uint8

const (
	// TextOnlyButton shows the display name as a plain button (the zero value).
	TextOnlyButton ButtonStyle = 0
	// SmallIconButton shows a small (16px) icon only — for dense tool grids.
	SmallIconButton ButtonStyle = 1
	// LargeIconButton shows a large (32px) icon with the name as a caption beneath.
	LargeIconButton ButtonStyle = 2
)

var buttonStyleNames = map[ButtonStyle]string{
	TextOnlyButton: "text", SmallIconButton: "small-icon", LargeIconButton: "large-icon",
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
	return s == SmallIconButton || s == LargeIconButton
}
