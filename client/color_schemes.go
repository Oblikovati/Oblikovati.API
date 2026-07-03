// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// ColorSchemes is the application color-scheme operation group: it lets an add-in read the
// available named palettes and switch the active one, so an automation can drive the
// viewport background and the highlight/selection colors.
type ColorSchemes struct{ c *Client }

// ColorSchemes returns the application color-scheme operation group.
func (c *Client) ColorSchemes() ColorSchemes { return ColorSchemes{c} }

// List returns every color scheme, flagging the active one.
//
//	schemes, _ := client.ColorSchemes().List()
//
// mcp:tool list_color_schemes
// mcp:summary List the application color schemes (the names set_color_scheme accepts).
func (s ColorSchemes) List() (wire.ColorSchemesResult, error) {
	return call[wire.ColorSchemesResult](s.c, wire.MethodColorSchemesList, nil)
}

// Active returns the currently active color scheme.
//
// mcp:tool get_active_color_scheme
// mcp:summary Read the active color scheme (background, highlight, and selection colors).
func (s ColorSchemes) Active() (wire.ColorSchemeView, error) {
	return call[wire.ColorSchemeView](s.c, wire.MethodColorSchemesGetActive, nil)
}

// SetActive activates the named color scheme, returning the now-active scheme.
//
//	client.ColorSchemes().SetActive("Presentation")
//
// mcp:tool set_active_color_scheme
// mcp:summary Switch the active color scheme by name; see list_color_schemes.
func (s ColorSchemes) SetActive(name string) (wire.ColorSchemeView, error) {
	return call[wire.ColorSchemeView](s.c, wire.MethodColorSchemesSetActive, wire.SetColorSchemeArgs{Name: name})
}
