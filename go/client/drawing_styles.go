// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The drawing styles operation group (M14-F01 PBI-138, Oblikovati/Oblikovati#385): read the
// active drawing's drafting standard and its dimension/text/line style preset, and switch the
// standard (which re-points the preset, changing the drawing's appearance).

// DrawingStyles is the drawing-styles operation group.
type DrawingStyles struct{ c *Client }

// DrawingStyles returns the drawing-styles operation group.
func (c *Client) DrawingStyles() DrawingStyles { return DrawingStyles{c} }

// ListStandards returns the available drafting standards and the active one.
//
// mcp:tool drawing_list_standards
// mcp:summary List the active drawing's available drafting standards (iso, ansi) and which is active.
func (d DrawingStyles) ListStandards() (wire.ListStandardsResult, error) {
	return call[wire.ListStandardsResult](d.c, wire.MethodDrawingStylesListStandards, struct{}{})
}

// GetActiveStyle returns the active standard's resolved dimension/text/line style preset.
//
// mcp:tool drawing_get_active_style
// mcp:summary Read the active drawing standard's style preset — dimension (text/arrow size, decimals, unit, line weight), text (font, height) and line (weight) styles.
func (d DrawingStyles) GetActiveStyle() (wire.StandardStyleResult, error) {
	return call[wire.StandardStyleResult](d.c, wire.MethodDrawingStylesGetActiveStyle, struct{}{})
}

// SetStandard makes the named drafting standard active and returns its style preset; every
// annotation re-renders to it.
//
// mcp:tool drawing_set_standard
// mcp:summary Switch the active drawing's drafting standard (iso|ansi); returns the new active style preset (dimension/text/line), so the appearance change is visible in one call.
func (d DrawingStyles) SetStandard(args wire.SetStandardArgs) (wire.StandardStyleResult, error) {
	return call[wire.StandardStyleResult](d.c, wire.MethodDrawingStylesSetStandard, args)
}
