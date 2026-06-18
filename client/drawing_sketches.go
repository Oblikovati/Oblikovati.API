// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The drawing sketches operation group (M14-F08 #638): 2D geometry drawn directly in sheet space
// (millimetres) on a sheet — linework and boundaries that hatch regions can fill.

// DrawingSketches is the drawing-sketches operation group.
type DrawingSketches struct{ c *Client }

// DrawingSketches returns the drawing-sketches operation group.
func (c *Client) DrawingSketches() DrawingSketches { return DrawingSketches{c} }

// List returns the active sheet's drawing sketches.
//
// mcp:tool drawing_list_sketches
// mcp:summary List the active sheet's drawing sketches (name, entity count, curve count).
func (d DrawingSketches) List() (wire.ListDrawingSketchesResult, error) {
	var r wire.ListDrawingSketchesResult
	return r, d.c.call(wire.MethodDrawingSketchesList, struct{}{}, &r)
}

// Add creates a new empty drawing sketch on the active sheet.
//
// mcp:tool drawing_add_sketch
// mcp:summary Add a new empty 2D sketch to the active sheet (sheet-millimetre space); add geometry to it with drawing_add_sketch_entity.
func (d DrawingSketches) Add(args wire.AddDrawingSketchArgs) (wire.DrawingSketchResult, error) {
	var r wire.DrawingSketchResult
	return r, d.c.call(wire.MethodDrawingSketchesAdd, args, &r)
}

// AddEntity adds one entity (line, circle or rectangle) to a drawing sketch.
//
// mcp:tool drawing_add_sketch_entity
// mcp:summary Add an entity to a drawing sketch (sketchName): kind line|circle|rectangle, points = sheet-mm [x,y] pairs (2 for line endpoints / rectangle corners, 1 for a circle centre with radiusMm).
func (d DrawingSketches) AddEntity(args wire.AddDrawingSketchEntityArgs) (wire.DrawingSketchResult, error) {
	var r wire.DrawingSketchResult
	return r, d.c.call(wire.MethodDrawingSketchesAddEntity, args, &r)
}
