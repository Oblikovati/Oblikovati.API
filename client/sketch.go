// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Sketch is the sketch-authoring operation group for the active part.
type Sketch struct{ c *Client }

// Sketch returns the sketch operation group.
func (c *Client) Sketch() Sketch { return Sketch{c} }

// Create adds a sketch on an origin plane and returns its index.
//
// mcp:tool create_sketch
// mcp:summary Create a sketch and return its sketchIndex. Default is an origin plane (plane: XY|XZ|YZ, default XY); set workPlaneIndex to sketch on a user work plane instead — the way to sketch on a plane built on a feature-created face (see create_work_plane + get_reference_keys) so later features reference earlier geometry.
func (s Sketch) Create(args wire.CreateSketchArgs) (wire.CreateSketchResult, error) {
	var r wire.CreateSketchResult
	return r, s.c.call(wire.MethodSketchCreate, args, &r)
}

// Rectangle adds a closed rectangle (one profile) to a sketch.
//
// mcp:tool sketch_rectangle
// mcp:summary Add a closed rectangle to a sketch (width, height as unit expressions, e.g. "40 mm"), forming a profile to extrude.
func (s Sketch) Rectangle(args wire.SketchRectangleArgs) (wire.SketchRectangleResult, error) {
	var r wire.SketchRectangleResult
	return r, s.c.call(wire.MethodSketchRectangle, args, &r)
}

// List enumerates the active part's sketches with their identity, DOF, and health.
//
// mcp:tool list_sketches
// mcp:summary List the active part's 2D sketches (index, name, plane, entity count, remaining DOF).
// mcp:digest summarizeSketches
func (s Sketch) List() (wire.ListSketchesResult, error) {
	var r wire.ListSketchesResult
	return r, s.c.call(wire.MethodSketchList, nil, &r)
}

// Get returns a single sketch's info by index.
//
// mcp:tool get_sketch
// mcp:summary Get one 2D sketch's properties by sketchIndex (name, plane, visibility, entity count, DOF).
func (s Sketch) Get(index int) (wire.SketchInfo, error) {
	var r wire.SketchInfo
	return r, s.c.call(wire.MethodSketchGet, wire.SketchArgs{SketchIndex: index}, &r)
}

// Edit opens the sketch for geometry editing (enters edit mode).
//
// mcp:tool edit_sketch
// mcp:summary Open a sketch for editing (enter its sketch environment).
func (s Sketch) Edit(index int) (wire.EditSketchResult, error) {
	var r wire.EditSketchResult
	return r, s.c.call(wire.MethodSketchEdit, wire.SketchArgs{SketchIndex: index}, &r)
}

// ExitEdit leaves edit mode, returning to the previous environment.
//
// mcp:tool exit_sketch
// mcp:summary Leave the sketch environment and update the part.
func (s Sketch) ExitEdit(index int) (wire.EditSketchResult, error) {
	var r wire.EditSketchResult
	return r, s.c.call(wire.MethodSketchExitEdit, wire.SketchArgs{SketchIndex: index}, &r)
}

// Solve resolves the sketch from its constraints and reports DOF/status/health.
//
// mcp:tool solve_sketch
// mcp:summary Re-solve a sketch's constraints and report its resulting degrees of freedom.
func (s Sketch) Solve(index int) (wire.SolveSketchResult, error) {
	var r wire.SolveSketchResult
	return r, s.c.call(wire.MethodSketchSolve, wire.SketchArgs{SketchIndex: index}, &r)
}

// Delete removes the sketch (only valid when no feature consumes it).
//
// mcp:tool delete_sketch
// mcp:summary Delete a sketch by sketchIndex.
func (s Sketch) Delete(index int) (wire.OKResult, error) {
	var r wire.OKResult
	return r, s.c.call(wire.MethodSketchDelete, wire.SketchArgs{SketchIndex: index}, &r)
}

// Entities enumerates the sketch's geometry (kind, construction flag, points, radius).
//
// mcp:tool list_sketch_entities
// mcp:summary Enumerate a sketch's geometry (each entity's index, session id, kind, construction flag) — the ids constraints/dimensions/transform reference.
// mcp:digest summarizeEntities
func (s Sketch) Entities(index int) (wire.EnumerateEntitiesResult, error) {
	var r wire.EnumerateEntitiesResult
	return r, s.c.call(wire.MethodSketchEntities, wire.SketchArgs{SketchIndex: index}, &r)
}

// ReferenceKey returns the sketch's persistent reference key (#153): a document-scoped UUID
// stable across save/load and edits. Store it to refer to the sketch durably; rebind it
// with [Sketch.ResolveReference].
//
// mcp:tool sketch_reference_key
// mcp:summary Get a sketch's persistent reference key — a document-scoped UUID stable across save/load, for durable references.
func (s Sketch) ReferenceKey(index int) (wire.SketchReferenceKeyResult, error) {
	var r wire.SketchReferenceKeyResult
	return r, s.c.call(wire.MethodSketchReferenceKey, wire.SketchArgs{SketchIndex: index}, &r)
}

// ResolveReference rebinds a previously stored persistent key (a sketch's or an entity's,
// #153) to its current location. Found is false when the referent was deleted.
//
// mcp:tool resolve_sketch_reference
// mcp:summary Rebind a stored persistent sketch/entity reference key to its current sketch index and entity id.
func (s Sketch) ResolveReference(key string) (wire.ResolveSketchReferenceResult, error) {
	var r wire.ResolveSketchReferenceResult
	return r, s.c.call(wire.MethodSketchResolveReference, wire.ResolveSketchReferenceArgs{ReferenceKey: key}, &r)
}

// Constraints enumerates the sketch's geometric constraints.
//
// mcp:tool list_sketch_constraints
// mcp:summary Enumerate a sketch's geometric constraints (index, kind, related entity ids).
// mcp:digest summarizeConstraints
func (s Sketch) Constraints(index int) (wire.ListConstraintsResult, error) {
	var r wire.ListConstraintsResult
	return r, s.c.call(wire.MethodSketchConstraints, wire.SketchArgs{SketchIndex: index}, &r)
}

// Dimensions enumerates the sketch's dimensional constraints.
//
// mcp:tool list_sketch_dimensions
// mcp:summary Enumerate a sketch's dimensional constraints (index, kind, backing parameter, expression, value, driven flag).
// mcp:digest summarizeDimensions
func (s Sketch) Dimensions(index int) (wire.ListDimensionsResult, error) {
	var r wire.ListDimensionsResult
	return r, s.c.call(wire.MethodSketchDimensions, wire.SketchArgs{SketchIndex: index}, &r)
}

// ConstraintStatus reports the sketch's DOF/over-under-constraint state without solving.
//
// mcp:tool get_sketch_constraint_status
// mcp:summary Report a sketch's constraint state and remaining degrees of freedom WITHOUT moving geometry (non-mutating DOF analysis).
func (s Sketch) ConstraintStatus(index int) (wire.ConstraintStatusResult, error) {
	var r wire.ConstraintStatusResult
	return r, s.c.call(wire.MethodSketchConstraintStatus, wire.SketchArgs{SketchIndex: index}, &r)
}

// Profiles enumerates the closed regions the sketch yields (area + hole count); the
// Index of each feeds features.add's profileIndex.
//
// mcp:tool list_sketch_profiles
// mcp:summary Enumerate a sketch's closed profiles (index, area, closed, hole count) — the profileIndex an extrude/revolve consumes.
// mcp:digest summarizeProfiles
func (s Sketch) Profiles(index int) (wire.ListProfilesResult, error) {
	var r wire.ListProfilesResult
	return r, s.c.call(wire.MethodSketchProfiles, wire.SketchArgs{SketchIndex: index}, &r)
}

// SetProperty sets one of the sketch's scalar properties and returns the updated info.
// Prefer the typed helpers below; this is the escape hatch.
//
// mcp:tool set_sketch_property
// mcp:summary Set a sketch property by name (e.g. property="name"|"visible"|"color"|"lineType"|"lineWeight").
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

// SetLineType overrides the sketch's line style (a [oblikovati.org/api/types.SketchLineType]).
func (s Sketch) SetLineType(index int, lineType types.SketchLineType) (wire.SketchInfo, error) {
	return s.SetProperty(index, "lineType", string(lineType))
}

// SetLineWeight overrides the sketch's line weight (a unit-bearing length like "0.5 mm").
func (s Sketch) SetLineWeight(index int, weight string) (wire.SketchInfo, error) {
	return s.SetProperty(index, "lineWeight", weight)
}

// SetDeferUpdates toggles whether the sketch batches edits (solving on resume).
func (s Sketch) SetDeferUpdates(index int, deferred bool) (wire.SketchInfo, error) {
	return s.SetProperty(index, "deferUpdates", boolText(deferred))
}

// boolText renders a bool as the "true"/"false" the property setter expects.
func boolText(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
