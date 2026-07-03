// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Styles is the style-manager operation group: it lets an add-in read, edit, and delete the
// document's color styles and import style libraries, so an automation can drive the document's
// visual styling. (Lighting styles use the Lighting group.)
type Styles struct{ c *Client }

// Styles returns the style-manager operation group.
func (c *Client) Styles() Styles { return Styles{c} }

// List returns every color style in the document.
//
//	styles, _ := client.Styles().List()
//
// mcp:tool list_color_styles
// mcp:summary List the document's color styles (name + diffuse/ambient/specular/emissive colors).
func (s Styles) List() (wire.ColorStylesResult, error) {
	return call[wire.ColorStylesResult](s.c, wire.MethodStylesList, nil)
}

// Get returns the named color style.
//
// mcp:tool get_color_style
// mcp:summary Read one color style by name; see list_color_styles.
func (s Styles) Get(name string) (wire.ColorStyleView, error) {
	return call[wire.ColorStyleView](s.c, wire.MethodStylesGet, wire.GetStyleArgs{Name: name})
}

// Set creates or updates a color style, returning the stored style. A style edit propagates to
// every consumer of that style.
//
// mcp:tool set_color_style
// mcp:summary Create or update a color style; see get_color_style for the shape.
func (s Styles) Set(v wire.ColorStyleView) (wire.ColorStyleView, error) {
	return call[wire.ColorStyleView](s.c, wire.MethodStylesSet, v)
}

// Delete removes the named color style.
//
// mcp:tool delete_color_style
// mcp:summary Delete a color style by name.
func (s Styles) Delete(name string) (wire.OKResult, error) {
	return call[wire.OKResult](s.c, wire.MethodStylesDelete, wire.GetStyleArgs{Name: name})
}

// Libraries returns the loaded style libraries, in cascade order.
//
// mcp:tool list_style_libraries
// mcp:summary List the loaded style libraries in cascade order.
func (s Styles) Libraries() (wire.StyleLibrariesResult, error) {
	return call[wire.StyleLibrariesResult](s.c, wire.MethodStylesListLibraries, nil)
}

// ImportLibrary loads a style library file into the cascade, returning the updated library list.
//
// mcp:tool import_style_library
// mcp:summary Load a style library file into the cascade.
func (s Styles) ImportLibrary(path string) (wire.StyleLibrariesResult, error) {
	return call[wire.StyleLibrariesResult](s.c, wire.MethodStylesImportLibrary, wire.ImportStyleLibraryArgs{Path: path})
}
