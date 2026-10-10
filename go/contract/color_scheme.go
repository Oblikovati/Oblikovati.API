// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// ColorScheme is the in-process contract for one named application palette — the colors the
// viewport and selection pipeline traffic in. It is read-mostly here; activation and edits go
// through the wire methods / the app. The GPL app satisfies it (compile-time asserted there).
//
// Background colors honor the owning [ColorSchemes].BackgroundType: ScreenColor for a solid
// background, TopScreenColor/BottomScreenColor for a gradient. Highlight/select colors feed the
// selection pipeline and HighlightSet.
type ColorScheme interface {
	// Name is the scheme's user-facing label (unique within the collection).
	Name() string
	// ScreenColor is the solid viewport background color.
	ScreenColor() types.Color
	// TopScreenColor is the gradient background's top color.
	TopScreenColor() types.Color
	// BottomScreenColor is the gradient background's bottom color.
	BottomScreenColor() types.Color
	// HighlightColor is the pre-highlight (hover) color.
	HighlightColor() types.Color
	// PrimarySelectColor is the first-selection color.
	PrimarySelectColor() types.Color
	// SecondarySelectColor is the secondary-selection color.
	SecondarySelectColor() types.Color
}

// ColorSchemes is the in-process contract for the application's set of color schemes — the
// equivalent of the Color tab of the application options. It enumerates the schemes, reports
// and switches the active one, and carries the application-wide background type. The GPL app
// satisfies it (compile-time asserted there).
type ColorSchemes interface {
	Enumerable[ColorScheme]
	// Active is the currently active scheme.
	Active() ColorScheme
	// SetActive makes the named scheme active, returning an error naming the scheme if absent.
	SetActive(name string) error
	// BackgroundType is the application-wide viewport background type.
	BackgroundType() types.BackgroundTypeEnum
	// SetBackgroundType sets the application-wide viewport background type.
	SetBackgroundType(t types.BackgroundTypeEnum) error
}
