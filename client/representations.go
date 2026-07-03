// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The assembly representation operation groups (M12-F04, Oblikovati/Oblikovati#361/#367):
// capture, activate, and edit the three representation families — design-view, positional,
// level-of-detail — plus model states selecting one of each. Representations are named
// override layers over an immutable base; capture snapshots the current scene, activate
// applies a representation, and the set* methods edit individual overrides.

// DesignReps is the design-view representation group (visibility/appearance/section/camera).
type DesignReps struct{ c *Client }

// DesignReps returns the design-view representation group.
func (c *Client) DesignReps() DesignReps { return DesignReps{c} }

// Capture snapshots the current visibility, appearance, sections, and camera into a new
// design-view representation.
//
// mcp:tool capture_design_view
// mcp:summary Capture the current scene's visibility/appearance/section/camera state into a new named design-view representation. Returns the representation.
func (d DesignReps) Capture(name string) (wire.DesignViewResult, error) {
	return call[wire.DesignViewResult](d.c, wire.MethodDesignRepsCapture, wire.CaptureRepArgs{Name: name})
}

// Activate applies a design-view representation's overrides to the scene.
//
// mcp:tool activate_design_view
// mcp:summary Activate a design-view representation (id) — apply its visibility/appearance/section/camera overrides to the scene. Returns the activated representation.
func (d DesignReps) Activate(id uint64) (wire.DesignViewResult, error) {
	return call[wire.DesignViewResult](d.c, wire.MethodDesignRepsActivate, wire.RepRef{ID: id})
}

// List returns the assembly's design-view representations in creation order.
//
// mcp:tool list_design_views
// mcp:summary List the assembly's design-view representations: each with id, name, active flag, hidden/appearance-override counts, section planes, and captured camera.
func (d DesignReps) List() (wire.DesignViewsResult, error) {
	return call[wire.DesignViewsResult](d.c, wire.MethodDesignRepsList, struct{}{})
}

// Delete removes a design-view representation and returns the remaining set.
//
// mcp:tool delete_design_view
// mcp:summary Delete a design-view representation (id). Returns the remaining set.
func (d DesignReps) Delete(id uint64) (wire.DesignViewsResult, error) {
	return call[wire.DesignViewsResult](d.c, wire.MethodDesignRepsDelete, wire.RepRef{ID: id})
}

// SetVisibility hides or shows an occurrence within a design-view representation.
//
// mcp:tool set_design_view_visibility
// mcp:summary Hide or show an occurrence (by id) within a design-view representation (rep id). Returns the updated representation.
func (d DesignReps) SetVisibility(args wire.SetVisibilityArgs) (wire.DesignViewResult, error) {
	return call[wire.DesignViewResult](d.c, wire.MethodDesignRepsSetVisibility, args)
}

// SetAppearance overrides (or clears) an occurrence's appearance within a design-view
// representation.
//
// mcp:tool set_design_view_appearance
// mcp:summary Override an occurrence's appearance (appearanceId; empty clears) within a design-view representation. Returns the updated representation.
func (d DesignReps) SetAppearance(args wire.SetAppearanceArgs) (wire.DesignViewResult, error) {
	return call[wire.DesignViewResult](d.c, wire.MethodDesignRepsSetAppearance, args)
}

// AddSection adds a section/clipping plane to a design-view representation.
//
// mcp:tool add_design_view_section
// mcp:input addSectionArg
// mcp:summary Add a section/clipping plane (origin + normal as [x,y,z], optional flipped) to a design-view representation. Returns the updated representation.
func (d DesignReps) AddSection(args wire.AddSectionArgs) (wire.DesignViewResult, error) {
	return call[wire.DesignViewResult](d.c, wire.MethodDesignRepsAddSection, args)
}

// PositionalReps is the positional representation group (constraint/joint value overrides).
type PositionalReps struct{ c *Client }

// PositionalReps returns the positional representation group.
func (c *Client) PositionalReps() PositionalReps { return PositionalReps{c} }

// Capture snapshots the current constraint/joint values into a new positional representation.
//
// mcp:tool capture_positional
// mcp:summary Capture the current constraint/joint values into a new named positional representation. Returns the representation.
func (p PositionalReps) Capture(name string) (wire.PositionalResult, error) {
	return call[wire.PositionalResult](p.c, wire.MethodPositionalRepsCapture, wire.CaptureRepArgs{Name: name})
}

// Activate applies a positional representation's value overrides and re-solves the assembly.
//
// mcp:tool activate_positional
// mcp:summary Activate a positional representation (id) — apply its constraint/joint value overrides and re-solve the assembly to that position. Returns the representation.
func (p PositionalReps) Activate(id uint64) (wire.PositionalResult, error) {
	return call[wire.PositionalResult](p.c, wire.MethodPositionalRepsActivate, wire.RepRef{ID: id})
}

// List returns the assembly's positional representations in creation order.
//
// mcp:tool list_positionals
// mcp:summary List the assembly's positional representations: each with id, name, active flag, and override count.
func (p PositionalReps) List() (wire.PositionalsResult, error) {
	return call[wire.PositionalsResult](p.c, wire.MethodPositionalRepsList, struct{}{})
}

// Delete removes a positional representation and returns the remaining set.
//
// mcp:tool delete_positional
// mcp:summary Delete a positional representation (id). Returns the remaining set.
func (p PositionalReps) Delete(id uint64) (wire.PositionalsResult, error) {
	return call[wire.PositionalsResult](p.c, wire.MethodPositionalRepsDelete, wire.RepRef{ID: id})
}

// SetOverride overrides a constraint's or joint's value within a positional representation.
//
// mcp:tool set_positional_override
// mcp:summary Override a constraint's (or joint's, isJoint:true) value within a positional representation. Returns the updated representation.
func (p PositionalReps) SetOverride(args wire.SetPositionalOverrideArgs) (wire.PositionalResult, error) {
	return call[wire.PositionalResult](p.c, wire.MethodPositionalRepsSetOverride, args)
}

// SetFlexible sets a subassembly occurrence's flexibility within a positional representation.
//
// mcp:tool set_positional_flexible
// mcp:summary Set a subassembly occurrence's flexibility within a positional representation. Returns the updated representation.
func (p PositionalReps) SetFlexible(args wire.SetFlexibleArgs) (wire.PositionalResult, error) {
	return call[wire.PositionalResult](p.c, wire.MethodPositionalRepsSetFlexible, args)
}

// LODReps is the level-of-detail representation group (occurrence suppression).
type LODReps struct{ c *Client }

// LODReps returns the level-of-detail representation group.
func (c *Client) LODReps() LODReps { return LODReps{c} }

// Capture snapshots the current suppression state into a new level-of-detail representation.
//
// mcp:tool capture_lod
// mcp:summary Capture the current occurrence-suppression state into a new named level-of-detail representation. Returns the representation.
func (l LODReps) Capture(name string) (wire.LODResult, error) {
	return call[wire.LODResult](l.c, wire.MethodLODRepsCapture, wire.CaptureRepArgs{Name: name})
}

// Activate applies a level-of-detail representation's suppression set.
//
// mcp:tool activate_lod
// mcp:summary Activate a level-of-detail representation (id) — apply its occurrence-suppression set. Returns the representation.
func (l LODReps) Activate(id uint64) (wire.LODResult, error) {
	return call[wire.LODResult](l.c, wire.MethodLODRepsActivate, wire.RepRef{ID: id})
}

// List returns the assembly's level-of-detail representations in creation order.
//
// mcp:tool list_lods
// mcp:summary List the assembly's level-of-detail representations: each with id, name, active flag, and suppressed-occurrence count.
func (l LODReps) List() (wire.LODsResult, error) {
	return call[wire.LODsResult](l.c, wire.MethodLODRepsList, struct{}{})
}

// Delete removes a level-of-detail representation and returns the remaining set.
//
// mcp:tool delete_lod
// mcp:summary Delete a level-of-detail representation (id). Returns the remaining set.
func (l LODReps) Delete(id uint64) (wire.LODsResult, error) {
	return call[wire.LODsResult](l.c, wire.MethodLODRepsDelete, wire.RepRef{ID: id})
}

// SetSuppressed suppresses or restores an occurrence within a level-of-detail representation.
//
// mcp:tool set_lod_suppressed
// mcp:summary Suppress or restore an occurrence (by id) within a level-of-detail representation. Returns the updated representation.
func (l LODReps) SetSuppressed(args wire.SetSuppressedArgs) (wire.LODResult, error) {
	return call[wire.LODResult](l.c, wire.MethodLODRepsSetSuppressed, args)
}

// ModelStates is the model-state group — a model state selects one representation of each
// family and activating it switches all three together.
type ModelStates struct{ c *Client }

// ModelStates returns the model-state group.
func (c *Client) ModelStates() ModelStates { return ModelStates{c} }

// Create makes a model state selecting one representation of each family by name.
//
// mcp:tool create_model_state
// mcp:summary Create a model state selecting one representation of each family by name (designView/positional/levelOfDetail; empty leaves a family unchanged). Returns the model state.
func (m ModelStates) Create(args wire.CreateModelStateArgs) (wire.ModelStateResult, error) {
	return call[wire.ModelStateResult](m.c, wire.MethodModelStatesCreate, args)
}

// Activate switches the assembly to a model state — activating its three representations
// together.
//
// mcp:tool activate_model_state
// mcp:summary Activate a model state (id) — switch the assembly's design-view, positional, and level-of-detail representations together. Returns the model state.
func (m ModelStates) Activate(id uint64) (wire.ModelStateResult, error) {
	return call[wire.ModelStateResult](m.c, wire.MethodModelStatesActivate, wire.RepRef{ID: id})
}

// List returns the assembly's model states in creation order.
//
// mcp:tool list_model_states
// mcp:summary List the assembly's model states: each with id, name, the selected representation names, and active flag.
func (m ModelStates) List() (wire.ModelStatesResult, error) {
	return call[wire.ModelStatesResult](m.c, wire.MethodModelStatesList, struct{}{})
}

// Delete removes a model state and returns the remaining set.
//
// mcp:tool delete_model_state
// mcp:summary Delete a model state (id). Returns the remaining set.
func (m ModelStates) Delete(id uint64) (wire.ModelStatesResult, error) {
	return call[wire.ModelStatesResult](m.c, wire.MethodModelStatesDelete, wire.RepRef{ID: id})
}
