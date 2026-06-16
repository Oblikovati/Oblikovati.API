// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The sheet-metal rule/style operation group (M13-F01, Oblikovati/Oblikovati#373/#369): read
// and edit the active sheet-metal part's rule (thickness/bend-radius/relief/gap + unfold
// method), and preview the developed flat length (bend allowance) of a single bend.

// SheetMetal is the sheet-metal rule operation group.
type SheetMetal struct{ c *Client }

// SheetMetal returns the sheet-metal rule operation group.
func (c *Client) SheetMetal() SheetMetal { return SheetMetal{c} }

// GetStyle returns the active sheet-metal rule.
//
// mcp:tool get_sheet_metal_style
// mcp:summary Report the active sheet-metal part's rule: thickness, bend radius, relief shape/size, minimum gap, unfold method and K-factor.
func (s SheetMetal) GetStyle() (wire.SheetMetalStyleResult, error) {
	var r wire.SheetMetalStyleResult
	return r, s.c.call(wire.MethodSheetMetalGetStyle, struct{}{}, &r)
}

// SetStyle edits the active sheet-metal rule and recomputes. Each field is optional;
// an empty value leaves that property unchanged.
//
// mcp:tool set_sheet_metal_style
// mcp:summary Edit the active sheet-metal rule (any of thickness, bendRadius, reliefShape, reliefWidth, reliefDepth, minimumGap, unfoldMethod, kFactor) and recompute. Returns the updated rule.
func (s SheetMetal) SetStyle(args wire.SetSheetMetalStyleArgs) (wire.SheetMetalStyleResult, error) {
	var r wire.SheetMetalStyleResult
	return r, s.c.call(wire.MethodSheetMetalSetStyle, args, &r)
}

// BendAllowance previews the developed flat length of one bend under the active rule.
//
// mcp:tool sheet_metal_bend_allowance
// mcp:summary Preview the developed flat length (bend allowance) and bend deduction of a single bend at a given angle (and optional radius) under the active sheet-metal unfold method.
func (s SheetMetal) BendAllowance(args wire.BendAllowanceArgs) (wire.BendAllowanceResult, error) {
	var r wire.BendAllowanceResult
	return r, s.c.call(wire.MethodSheetMetalBendAllowance, args, &r)
}
