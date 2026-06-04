// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

// Sketch3D is the 3D-sketch-authoring operation group for the active part. A 3D sketch
// has no host plane — its geometry lives directly in model space (sweep/loft paths,
// helices, on-surface curves).
type Sketch3D struct{ c *Client }

// Sketch3D returns the 3D-sketch operation group.
func (c *Client) Sketch3D() Sketch3D { return Sketch3D{c} }

// Create adds an empty 3D sketch and returns its index.
func (s Sketch3D) Create(args wire.CreateSketch3DArgs) (wire.CreateSketch3DResult, error) {
	var r wire.CreateSketch3DResult
	return r, s.c.call(wire.MethodSketch3DCreate, args, &r)
}

// List enumerates the active part's 3D sketches with their identity, DOF, and health.
func (s Sketch3D) List() (wire.ListSketches3DResult, error) {
	var r wire.ListSketches3DResult
	return r, s.c.call(wire.MethodSketch3DList, nil, &r)
}

// Get returns a single 3D sketch's info by index.
func (s Sketch3D) Get(index int) (wire.Sketch3DInfo, error) {
	var r wire.Sketch3DInfo
	return r, s.c.call(wire.MethodSketch3DGet, wire.Sketch3DArgs{SketchIndex: index}, &r)
}

// Edit opens the 3D sketch for geometry editing (enters edit mode).
func (s Sketch3D) Edit(index int) (wire.EditSketch3DResult, error) {
	var r wire.EditSketch3DResult
	return r, s.c.call(wire.MethodSketch3DEdit, wire.Sketch3DArgs{SketchIndex: index}, &r)
}

// ExitEdit leaves edit mode, returning to the previous environment.
func (s Sketch3D) ExitEdit(index int) (wire.EditSketch3DResult, error) {
	var r wire.EditSketch3DResult
	return r, s.c.call(wire.MethodSketch3DExitEdit, wire.Sketch3DArgs{SketchIndex: index}, &r)
}

// Solve resolves the 3D sketch from its constraints and reports DOF/status/health.
func (s Sketch3D) Solve(index int) (wire.SolveSketch3DResult, error) {
	var r wire.SolveSketch3DResult
	return r, s.c.call(wire.MethodSketch3DSolve, wire.Sketch3DArgs{SketchIndex: index}, &r)
}

// Delete removes the 3D sketch (only valid when no feature consumes it).
func (s Sketch3D) Delete(index int) (wire.OKResult, error) {
	var r wire.OKResult
	return r, s.c.call(wire.MethodSketch3DDelete, wire.Sketch3DArgs{SketchIndex: index}, &r)
}

// Entities enumerates the 3D sketch's geometry (kind, construction flag, points, radius).
func (s Sketch3D) Entities(index int) (wire.EnumerateEntities3DResult, error) {
	var r wire.EnumerateEntities3DResult
	return r, s.c.call(wire.MethodSketch3DEntities, wire.Sketch3DArgs{SketchIndex: index}, &r)
}

// Constraints enumerates the 3D sketch's geometric constraints.
func (s Sketch3D) Constraints(index int) (wire.ListConstraints3DResult, error) {
	var r wire.ListConstraints3DResult
	return r, s.c.call(wire.MethodSketch3DConstraints, wire.Sketch3DArgs{SketchIndex: index}, &r)
}

// Dimensions enumerates the 3D sketch's dimensional constraints.
func (s Sketch3D) Dimensions(index int) (wire.ListDimensions3DResult, error) {
	var r wire.ListDimensions3DResult
	return r, s.c.call(wire.MethodSketch3DDimensions, wire.Sketch3DArgs{SketchIndex: index}, &r)
}

// ConstraintStatus reports the 3D sketch's DOF/over-under-constraint state without solving.
func (s Sketch3D) ConstraintStatus(index int) (wire.ConstraintStatusResult, error) {
	var r wire.ConstraintStatusResult
	return r, s.c.call(wire.MethodSketch3DConstraintStatus, wire.Sketch3DArgs{SketchIndex: index}, &r)
}

// SetProperty sets one of the 3D sketch's scalar properties and returns the updated info.
// Prefer the typed helpers below; this is the escape hatch.
func (s Sketch3D) SetProperty(index int, property, value string) (wire.Sketch3DInfo, error) {
	var r wire.Sketch3DInfo
	args := wire.SetSketch3DPropertyArgs{SketchIndex: index, Property: property, Value: value}
	return r, s.c.call(wire.MethodSketch3DSetProperty, args, &r)
}

// SetName renames the 3D sketch.
func (s Sketch3D) SetName(index int, name string) (wire.Sketch3DInfo, error) {
	return s.SetProperty(index, "name", name)
}

// SetVisible shows or hides the 3D sketch.
func (s Sketch3D) SetVisible(index int, visible bool) (wire.Sketch3DInfo, error) {
	return s.SetProperty(index, "visible", boolText(visible))
}

// SetDimensionsVisible shows or hides the 3D sketch's dimensions.
func (s Sketch3D) SetDimensionsVisible(index int, visible bool) (wire.Sketch3DInfo, error) {
	return s.SetProperty(index, "dimensionsVisible", boolText(visible))
}

// SetColor overrides the 3D sketch's color (empty ⇒ inherit the document default).
func (s Sketch3D) SetColor(index int, color string) (wire.Sketch3DInfo, error) {
	return s.SetProperty(index, "color", color)
}

// SetDeferUpdates toggles whether the 3D sketch batches edits (solving on resume).
func (s Sketch3D) SetDeferUpdates(index int, defer_ bool) (wire.Sketch3DInfo, error) {
	return s.SetProperty(index, "deferUpdates", boolText(defer_))
}
