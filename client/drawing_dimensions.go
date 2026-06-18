// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The drawing dimensions operation group (M14-F03 PBI-141, #388): associative linear dimensions
// placed on a drawing view by two pick points.

// DrawingDimensions is the drawing-dimensions operation group.
type DrawingDimensions struct{ c *Client }

// DrawingDimensions returns the drawing-dimensions operation group.
func (c *Client) DrawingDimensions() DrawingDimensions { return DrawingDimensions{c} }

// List returns the active sheet's dimensions.
//
// mcp:tool drawing_list_dimensions
// mcp:summary List the active sheet's drawing dimensions (name, type = aligned|horizontal|vertical, the view, measured value in mm, displayed text, curve count).
func (d DrawingDimensions) List() (wire.ListDrawingDimensionsResult, error) {
	var r wire.ListDrawingDimensionsResult
	return r, d.c.call(wire.MethodDrawingDimensionsList, struct{}{}, &r)
}

// AddLinear adds a linear dimension on a view between two pick points.
//
// mcp:tool drawing_add_linear_dimension
// mcp:summary Add a linear dimension on a drawing view between two pick points (x1,y1,x2,y2 in sheet mm, each snapped to the nearest projected model vertex). type = aligned (true distance, default) | horizontal | vertical; offsetMm stands the dimension line off the points. The measured value is the true model size and updates with the model.
func (d DrawingDimensions) AddLinear(args wire.AddLinearDimensionArgs) (wire.DimensionResult, error) {
	var r wire.DimensionResult
	return r, d.c.call(wire.MethodDrawingDimensionsAddLinear, args, &r)
}

// AddRadial adds a radius or diameter dimension on the circular edge nearest a pick point.
//
// mcp:tool drawing_add_radial_dimension
// mcp:summary Add a radius or diameter dimension on a drawing view, attached to the circular model edge nearest the pick point (pickXmm/pickYmm sheet mm). type = radius (default) | diameter. The value is the true model size and updates with the model.
func (d DrawingDimensions) AddRadial(args wire.AddRadialDimensionArgs) (wire.DimensionResult, error) {
	var r wire.DimensionResult
	return r, d.c.call(wire.MethodDrawingDimensionsAddRadial, args, &r)
}

// Delete removes the named dimension.
//
// mcp:tool drawing_delete_dimension
// mcp:summary Delete a drawing dimension by name.
func (d DrawingDimensions) Delete(args wire.DeleteDimensionArgs) (wire.ListDrawingDimensionsResult, error) {
	var r wire.ListDrawingDimensionsResult
	return r, d.c.call(wire.MethodDrawingDimensionsDelete, args, &r)
}
