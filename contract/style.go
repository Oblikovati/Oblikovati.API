// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// ColorStyle is the in-process contract for one named color style — the classic visual style
// carrying the ambient/diffuse/specular/emissive color components, shininess, and opacity that
// a body or face references. It is read-mostly here; edits go through the wire methods. The
// GPL app satisfies it (compile-time asserted there).
type ColorStyle interface {
	// Name is the style's user-facing label (unique within the collection).
	Name() string
	// DiffuseColor is the main surface color under direct light.
	DiffuseColor() types.Color
	// AmbientColor is the surface color under ambient fill.
	AmbientColor() types.Color
	// SpecularColor is the highlight color.
	SpecularColor() types.Color
	// EmissiveColor is the self-emitted (glow) color.
	EmissiveColor() types.Color
	// Shininess is the specular exponent in [0,1] (higher is tighter highlights).
	Shininess() float64
	// Opacity is the surface opacity in [0,1].
	Opacity() float64
	// Location is where the style lives in the cascade (local overrides library).
	Location() types.StyleLocationEnum
}

// ColorStyles is the in-process contract for a document's set of color styles. The GPL app
// satisfies it (compile-time asserted there).
type ColorStyles interface {
	// Count is the number of color styles.
	Count() int
	// Item returns the color style at index i (0-based), or nil if out of range.
	Item(i int) ColorStyle
	// ByName returns the named color style, or nil if absent.
	ByName(name string) ColorStyle
}

// StyleManager is the in-process contract for the style/standard system: the registry of the
// document's color and lighting styles, the style-library cascade, and library load/save. The
// GPL app satisfies it (compile-time asserted there).
type StyleManager interface {
	// ColorStyles returns the document's color styles.
	ColorStyles() ColorStyles
	// LightingStyles returns the document's lighting styles.
	LightingStyles() []LightingStyle
	// LibraryNames returns the names of the loaded style libraries, in cascade order.
	LibraryNames() []string
	// ImportLibrary loads the style library at the given path into the cascade, returning an
	// error naming the path on a missing or malformed file.
	ImportLibrary(path string) error
}
