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
	return call[wire.ListDrawingDimensionsResult](d.c, wire.MethodDrawingDimensionsList, struct{}{})
}

// AddLinear adds a linear dimension on a view between two pick points.
//
// mcp:tool drawing_add_linear_dimension
// mcp:summary Add a linear dimension on a drawing view between two pick points (x1,y1,x2,y2 in sheet mm, each snapped to the nearest projected model vertex). type = aligned (true distance, default) | horizontal | vertical; offsetMm stands the dimension line off the points. The measured value is the true model size and updates with the model.
func (d DrawingDimensions) AddLinear(args wire.AddLinearDimensionArgs) (wire.DimensionResult, error) {
	return call[wire.DimensionResult](d.c, wire.MethodDrawingDimensionsAddLinear, args)
}

// AddRadial adds a radius or diameter dimension on the circular edge nearest a pick point.
//
// mcp:tool drawing_add_radial_dimension
// mcp:summary Add a radius or diameter dimension on a drawing view, attached to the circular model edge nearest the pick point (pickXmm/pickYmm sheet mm). type = radius (default) | diameter. The value is the true model size and updates with the model.
func (d DrawingDimensions) AddRadial(args wire.AddRadialDimensionArgs) (wire.DimensionResult, error) {
	return call[wire.DimensionResult](d.c, wire.MethodDrawingDimensionsAddRadial, args)
}

// AddAngular adds an angular dimension between the two straight edges nearest two pick points.
//
// mcp:tool drawing_add_angular_dimension
// mcp:summary Add an angular dimension on a drawing view between the two straight model edges nearest the pick points (x1,y1,x2,y2 sheet mm). The measured angle (degrees) is reported in valueDeg and updates with the model.
func (d DrawingDimensions) AddAngular(args wire.AddAngularDimensionArgs) (wire.DimensionResult, error) {
	return call[wire.DimensionResult](d.c, wire.MethodDrawingDimensionsAddAngular, args)
}

// AddBaseline adds a baseline set: linear dimensions from the first pick point to each of the
// others, stacked.
//
// mcp:tool drawing_add_baseline_dimensions
// mcp:summary Add a baseline dimension set on a drawing view: linear dimensions from the first pick point to each of the other points (each [x,y] sheet mm, snapped to model vertices), stacked. type = aligned|horizontal|vertical. The values update with the model.
func (d DrawingDimensions) AddBaseline(args wire.AddDimensionSetArgs) (wire.DimensionSetResult, error) {
	return call[wire.DimensionSetResult](d.c, wire.MethodDrawingDimensionsAddBaseline, args)
}

// AddChain adds a chain set: linear dimensions between consecutive pick points, in a line.
//
// mcp:tool drawing_add_chain_dimensions
// mcp:summary Add a chain dimension set on a drawing view: linear dimensions between consecutive pick points (each [x,y] sheet mm, snapped to model vertices), running in a line. type = aligned|horizontal|vertical. The values update with the model.
func (d DrawingDimensions) AddChain(args wire.AddDimensionSetArgs) (wire.DimensionSetResult, error) {
	return call[wire.DimensionSetResult](d.c, wire.MethodDrawingDimensionsAddChain, args)
}

// AddOrdinate adds an ordinate set: one leader-to-value dimension per point, each measuring that
// point's offset from a common datum along one axis.
//
// mcp:tool drawing_add_ordinate_dimensions
// mcp:summary Add an ordinate dimension set on a drawing view: one dimension per point measuring its offset from a common datum ([x,y] sheet mm, snapped to model vertices) along axis = horizontal (view-X, default) | vertical (view-Y). Each is drawn as a leader to its value with no dimension line; the values update with the model.
func (d DrawingDimensions) AddOrdinate(args wire.AddOrdinateDimensionsArgs) (wire.DimensionSetResult, error) {
	return call[wire.DimensionSetResult](d.c, wire.MethodDrawingDimensionsAddOrdinate, args)
}

// AddArcLength adds an arc-length dimension on the circular/arc edge nearest a pick point.
//
// mcp:tool drawing_add_arc_length_dimension
// mcp:summary Add an arc-length dimension on a drawing view, attached to the circular/arc model edge nearest the pick point (pickXmm/pickYmm sheet mm). It measures the edge's swept length (a full circle's circumference) with the dimension line following the arc. The value is the true model size and updates with the model.
func (d DrawingDimensions) AddArcLength(args wire.AddArcLengthDimensionArgs) (wire.DimensionResult, error) {
	return call[wire.DimensionResult](d.c, wire.MethodDrawingDimensionsAddArcLength, args)
}

// Delete removes the named dimension.
//
// mcp:tool drawing_delete_dimension
// mcp:summary Delete a drawing dimension by name.
// SetTextStyle changes any subset of a dimension's text overrides — prefix, suffix, free-text
// override, hide-value, dual-unit — leaving the unset ones alone (#1992/#1993).
//
// mcp:tool set_dimension_text_style
// mcp:summary Change a drawing dimension's text overrides (prefix/suffix/overrideText/hideValue/dualUnit); unset fields unchanged.
func (d DrawingDimensions) SetTextStyle(args wire.SetDimensionTextStyleArgs) (wire.ListDrawingDimensionsResult, error) {
	return call[wire.ListDrawingDimensionsResult](d.c, wire.MethodDrawingDimensionsSetTextStyle, args)
}

// SetTolerance sets a dimension's engineering tolerance — symmetric, deviation, limits or a fit
// class (#1990). A none tolerance clears it.
//
// mcp:tool set_dimension_tolerance
// mcp:summary Set a drawing dimension's engineering tolerance (symmetric/deviation/limits/fits).
func (d DrawingDimensions) SetTolerance(args wire.SetDimensionToleranceArgs) (wire.ListDrawingDimensionsResult, error) {
	return call[wire.ListDrawingDimensionsResult](d.c, wire.MethodDrawingDimensionsSetTolerance, args)
}

// Delete removes a dimension by name.
func (d DrawingDimensions) Delete(args wire.DeleteDimensionArgs) (wire.ListDrawingDimensionsResult, error) {
	return call[wire.ListDrawingDimensionsResult](d.c, wire.MethodDrawingDimensionsDelete, args)
}
