// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/Oblikovati/api/types"
	"github.com/Oblikovati/api/wire"
)

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

// ConstraintStatus reports the sketch's DOF/over-under-constraint state without solving.
func (s Sketch) ConstraintStatus(index int) (wire.ConstraintStatusResult, error) {
	var r wire.ConstraintStatusResult
	return r, s.c.call(wire.MethodSketchConstraintStatus, wire.SketchArgs{SketchIndex: index}, &r)
}

// SetProperty sets one of the sketch's scalar properties and returns the updated info.
// Prefer the typed helpers below; this is the escape hatch.
func (s Sketch) SetProperty(index int, property, value string) (wire.SketchInfo, error) {
	var r wire.SketchInfo
	args := wire.SetSketchPropertyArgs{SketchIndex: index, Property: property, Value: value}
	return r, s.c.call(wire.MethodSketchSetProperty, args, &r)
}

// SetName renames the sketch.
func (s Sketch) SetName(index int, name string) (wire.SketchInfo, error) {
	return s.SetProperty(index, "name", name)
}

// SetVisible shows or hides the sketch.
func (s Sketch) SetVisible(index int, visible bool) (wire.SketchInfo, error) {
	return s.SetProperty(index, "visible", boolText(visible))
}

// SetColor overrides the sketch's color (empty ⇒ inherit the document default).
func (s Sketch) SetColor(index int, color string) (wire.SketchInfo, error) {
	return s.SetProperty(index, "color", color)
}

// SetLineType overrides the sketch's line style (a [github.com/Oblikovati/api/types.SketchLineType]).
func (s Sketch) SetLineType(index int, lineType types.SketchLineType) (wire.SketchInfo, error) {
	return s.SetProperty(index, "lineType", string(lineType))
}

// SetLineWeight overrides the sketch's line weight (a unit-bearing length like "0.5 mm").
func (s Sketch) SetLineWeight(index int, weight string) (wire.SketchInfo, error) {
	return s.SetProperty(index, "lineWeight", weight)
}

// SetDeferUpdates toggles whether the sketch batches edits (solving on resume).
func (s Sketch) SetDeferUpdates(index int, defer_ bool) (wire.SketchInfo, error) {
	return s.SetProperty(index, "deferUpdates", boolText(defer_))
}

// boolText renders a bool as the "true"/"false" the property setter expects.
func boolText(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
