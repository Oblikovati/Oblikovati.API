// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

// Sketch is the sketch-authoring operation group for the active part.
type Sketch struct{ c *Client }

// Sketch returns the sketch operation group.
func (c *Client) Sketch() Sketch { return Sketch{c} }

// Create adds a sketch on an origin plane and returns its index.
func (s Sketch) Create(args wire.CreateSketchArgs) (wire.CreateSketchResult, error) {
	var r wire.CreateSketchResult
	return r, s.c.call(wire.MethodSketchCreate, args, &r)
}

// Rectangle adds a closed rectangle (one profile) to a sketch.
func (s Sketch) Rectangle(args wire.SketchRectangleArgs) (wire.SketchRectangleResult, error) {
	var r wire.SketchRectangleResult
	return r, s.c.call(wire.MethodSketchRectangle, args, &r)
}

// List enumerates the active part's sketches with their identity, DOF, and health.
func (s Sketch) List() (wire.ListSketchesResult, error) {
	var r wire.ListSketchesResult
	return r, s.c.call(wire.MethodSketchList, nil, &r)
}

// Get returns a single sketch's info by index.
func (s Sketch) Get(index int) (wire.SketchInfo, error) {
	var r wire.SketchInfo
	return r, s.c.call(wire.MethodSketchGet, wire.SketchArgs{SketchIndex: index}, &r)
}

// Edit opens the sketch for geometry editing (enters edit mode).
func (s Sketch) Edit(index int) (wire.EditSketchResult, error) {
	var r wire.EditSketchResult
	return r, s.c.call(wire.MethodSketchEdit, wire.SketchArgs{SketchIndex: index}, &r)
}

// ExitEdit leaves edit mode, returning to the previous environment.
func (s Sketch) ExitEdit(index int) (wire.EditSketchResult, error) {
	var r wire.EditSketchResult
	return r, s.c.call(wire.MethodSketchExitEdit, wire.SketchArgs{SketchIndex: index}, &r)
}

// Solve resolves the sketch from its constraints and reports DOF/status/health.
func (s Sketch) Solve(index int) (wire.SolveSketchResult, error) {
	var r wire.SolveSketchResult
	return r, s.c.call(wire.MethodSketchSolve, wire.SketchArgs{SketchIndex: index}, &r)
}

// Delete removes the sketch (only valid when no feature consumes it).
func (s Sketch) Delete(index int) (wire.OKResult, error) {
	var r wire.OKResult
	return r, s.c.call(wire.MethodSketchDelete, wire.SketchArgs{SketchIndex: index}, &r)
}

// Entities enumerates the sketch's geometry (kind, construction flag, points, radius).
func (s Sketch) Entities(index int) (wire.EnumerateEntitiesResult, error) {
	var r wire.EnumerateEntitiesResult
	return r, s.c.call(wire.MethodSketchEntities, wire.SketchArgs{SketchIndex: index}, &r)
}

// Constraints enumerates the sketch's geometric constraints.
func (s Sketch) Constraints(index int) (wire.ListConstraintsResult, error) {
	var r wire.ListConstraintsResult
	return r, s.c.call(wire.MethodSketchConstraints, wire.SketchArgs{SketchIndex: index}, &r)
}

// Dimensions enumerates the sketch's dimensional constraints.
func (s Sketch) Dimensions(index int) (wire.ListDimensionsResult, error) {
	var r wire.ListDimensionsResult
	return r, s.c.call(wire.MethodSketchDimensions, wire.SketchArgs{SketchIndex: index}, &r)
}
