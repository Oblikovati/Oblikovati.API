// SPDX-License-Identifier: Apache-2.0

package contract

import "github.com/Oblikovati/api/types"

// Theme is the in-process contract for one UI color theme: a name, a kind (light/dark/
// custom), and a color for every semantic token. The GPL implementation's theme.Theme
// satisfies it (compile-time asserted there); the head reads it to style the shell and
// add-ins read it (via the wire view) to match the host's look.
//
// Color must return a defined value for every token in [types.AllThemeTokens] — a
// complete theme never leaves a slot undefined, so callers need not handle a missing
// color.
type Theme interface {
	// Name is the display label, unique within the user's theme library.
	Name() string
	// Kind classifies the theme as a built-in (light/dark) or a user copy (custom).
	Kind() types.ThemeKind
	// Color returns the color bound to token in this theme.
	Color(token types.ThemeToken) types.Rgba
}
