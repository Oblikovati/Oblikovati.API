// SPDX-License-Identifier: Apache-2.0

package wire

// ThemeView is the JSON shape of one theme: its name, kind (light/dark/custom), and
// every semantic color as "#RRGGBBAA" hex keyed by the token string. Hex (not floats)
// keeps the payload compact and human-readable, matching the on-disk theme files.
type ThemeView struct {
	Name   string            `json:"name"`
	Kind   string            `json:"kind"`
	Colors map[string]string `json:"colors"`
}

// ThemeSummary is the lightweight entry of [MethodThemeList]: enough to populate a
// theme picker without sending every color.
type ThemeSummary struct {
	Name   string `json:"name"`
	Kind   string `json:"kind"`
	Active bool   `json:"active"`
}

// ListThemesResult is the response of [MethodThemeList].
type ListThemesResult struct {
	Themes []ThemeSummary `json:"themes"`
}
