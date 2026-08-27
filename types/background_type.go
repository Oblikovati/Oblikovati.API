// SPDX-License-Identifier: Apache-2.0

package types

// BackgroundTypeEnum is how the viewport background is painted: a single solid color, a
// vertical two-color gradient, or a background image. The numeric ids are stable, frozen
// values (52737–52739).
//
// First consumed by the color-scheme palette (which carries the background colors); the
// display-settings surface reuses it. This is the canonical Apache-2.0 definition; the GPL
// implementation aliases it (app.BackgroundTypeEnum).
type BackgroundTypeEnum int32

const (
	// OneColorBackground paints a single solid background color (52737).
	OneColorBackground BackgroundTypeEnum = 52737
	// GradientBackground paints a top-to-bottom two-color gradient (52738).
	GradientBackground BackgroundTypeEnum = 52738
	// ImageBackground paints a background image (52739).
	ImageBackground BackgroundTypeEnum = 52739
)

var backgroundTypeNames = map[BackgroundTypeEnum]string{
	OneColorBackground: "One Color",
	GradientBackground: "Gradient",
	ImageBackground:    "Image",
}

// String returns the background type's user-facing name.
func (b BackgroundTypeEnum) String() string {
	return enumName(backgroundTypeNames, b, "backgroundType(?)")
}

// IsValid reports whether b is a defined background type.
func (b BackgroundTypeEnum) IsValid() bool {
	return enumValid(backgroundTypeNames, b)
}

// AllBackgroundTypes returns every defined background type, in picker order.
func AllBackgroundTypes() []BackgroundTypeEnum {
	return []BackgroundTypeEnum{OneColorBackground, GradientBackground, ImageBackground}
}
