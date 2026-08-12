// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The assembly occurrence operations (M11-F01/F02, Oblikovati/Oblikovati#728) extend the
// Assembly group: read the active assembly's occurrence tree and place/transform/ground/
// suppress/replace/remove components, addressed by session id. Each single-occurrence
// operation returns the affected occurrence's refreshed info; remove returns the tree.

// Occurrences returns the active assembly's occurrence tree.
//
// mcp:tool list_occurrences
// mcp:summary Read the active assembly's occurrence tree: each placed component with its id, name, 4×4 placement transform, state flags (suppressed/grounded/adaptive/substitute), and nested children. The ids address the other assembly tools.
func (a Assembly) Occurrences() (wire.OccurrencesResult, error) {
	return call[wire.OccurrencesResult](a.c, wire.MethodAssemblyOccurrences, struct{}{})
}

// Place places the component held by the open document (by id) into the active assembly,
// e.g. Place(wire.PlaceOccurrenceArgs{Document: docID, Name: "pin:1", Transform: t}).
//
// mcp:tool place_component
// mcp:summary Place an open document (document: the id from list_documents — an open part or assembly) as a component in the active assembly, under name, at a row-major 4×4 transform (16 cells in assembly space; the identity [1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1] drops it at the origin). Returns the new occurrence.
// mcp:input placeComponentArg
func (a Assembly) Place(args wire.PlaceOccurrenceArgs) (wire.OccurrenceResult, error) {
	return call[wire.OccurrenceResult](a.c, wire.MethodAssemblyPlace, args)
}

// PlaceByDefinition places another instance of the component that the source occurrence
// already instances, e.g. PlaceByDefinition(wire.PlaceByDefinitionArgs{Source: occID, Name: "pin:2", Transform: t}).
//
// mcp:tool place_component_copy
// mcp:summary Place another instance of the component an existing occurrence (source: its id) already instances, under name at a 16-cell row-major transform — reuses the shared component definition without re-resolving a document.
// mcp:input placeComponentCopyArg
func (a Assembly) PlaceByDefinition(args wire.PlaceByDefinitionArgs) (wire.OccurrenceResult, error) {
	return call[wire.OccurrenceResult](a.c, wire.MethodAssemblyPlaceByDefinition, args)
}

// PlaceByDefinitionBatch places many instances of the component an existing occurrence already
// instances, in ONE call, e.g. PlaceByDefinitionBatch(wire.PlaceByDefinitionBatchArgs{Source: occID,
// Placements: []wire.BatchPlacement{{Name: "pin:2", Transform: t2}, {Name: "pin:3", Transform: t3}}}).
// For a large assembly this is far faster than PlaceByDefinition in a loop: a live host recomputes
// once for the whole batch instead of once per copy.
//
// mcp:tool place_component_copies
// mcp:summary Place MANY instances of the component an existing occurrence (source: its id) already instances, in one call — placements is a list of {name, transform (16-cell row-major)}. Far faster than place_component_copy per copy for a large assembly (one recompute, not one per placement). Returns the new occurrences in order.
// mcp:input placeComponentCopiesArg
func (a Assembly) PlaceByDefinitionBatch(args wire.PlaceByDefinitionBatchArgs) (wire.PlaceByDefinitionBatchResult, error) {
	return call[wire.PlaceByDefinitionBatchResult](a.c, wire.MethodAssemblyPlaceByDefinitionBatch, args)
}

// Transform repositions the occurrence, e.g. Transform(wire.TransformOccurrenceArgs{ID: id, Transform: t}).
//
// mcp:tool transform_occurrence
// mcp:summary Reposition an occurrence (id) to a new row-major 4×4 transform (16 cells) in the assembly's space. Returns the occurrence's refreshed info.
// mcp:input transformOccurrenceArg
func (a Assembly) Transform(args wire.TransformOccurrenceArgs) (wire.OccurrenceResult, error) {
	return call[wire.OccurrenceResult](a.c, wire.MethodAssemblyTransform, args)
}

// Ground fixes or releases the occurrence in space, e.g. Ground(id, true).
//
// mcp:tool ground_occurrence
// mcp:summary Fix (grounded:true) or release (grounded:false) an occurrence (id) in the assembly's space.
func (a Assembly) Ground(id uint64, grounded bool) (wire.OccurrenceResult, error) {
	return call[wire.OccurrenceResult](a.c, wire.MethodAssemblyGround, wire.GroundOccurrenceArgs{ID: id, Grounded: grounded})
}

// Suppress excludes or restores the occurrence from the model, e.g. Suppress(id, true).
//
// mcp:tool suppress_occurrence
// mcp:summary Exclude (suppressed:true) or restore (suppressed:false) an occurrence (id) from/to the model.
func (a Assembly) Suppress(id uint64, suppressed bool) (wire.OccurrenceResult, error) {
	return call[wire.OccurrenceResult](a.c, wire.MethodAssemblySuppress, wire.SuppressOccurrenceArgs{ID: id, Suppressed: suppressed})
}

// SetVisible shows or hides one occurrence — a display override independent of any representation,
// e.g. SetVisible(id, false) (#1975).
//
// mcp:tool set_occurrence_visible
// mcp:summary Show (visible:true) or hide an occurrence (id) — a display override.
func (a Assembly) SetVisible(id uint64, visible bool) (wire.OccurrenceResult, error) {
	return call[wire.OccurrenceResult](a.c, wire.MethodAssemblySetVisible, wire.SetVisibleOccurrenceArgs{ID: id, Visible: visible})
}

// SetOccurrenceState changes any subset of an occurrence's display/state overrides — transparency,
// opacity, enabled, excluded, reference, contact-set — leaving the unset ones alone (#1975/#1977).
//
// mcp:tool set_occurrence_state
// mcp:summary Change an occurrence's state overrides (transparent/opacity/enabled/excluded/reference/contactSet); unset fields are unchanged.
func (a Assembly) SetOccurrenceState(args wire.SetOccurrenceStateArgs) (wire.OccurrenceResult, error) {
	return call[wire.OccurrenceResult](a.c, wire.MethodAssemblySetOccurrenceState, args)
}

// SetFlexible marks a subassembly occurrence flexible (it solves independently per placement)
// or rigid, e.g. SetFlexible(id, true).
//
// mcp:tool set_flexible_occurrence
// mcp:summary Mark a subassembly occurrence (id) flexible (flexible:true — its components solve independently per placement of the shared definition) or rigid. Mutually exclusive with adaptive; only a subassembly occurrence can be flexible (M12-F06).
func (a Assembly) SetFlexible(id uint64, flexible bool) (wire.OccurrenceResult, error) {
	return call[wire.OccurrenceResult](a.c, wire.MethodAssemblySetFlexible, wire.SetFlexibleOccurrenceArgs{ID: id, Flexible: flexible})
}

// SetFlexibleChild positions a child component within a flexible subassembly occurrence
// independently of the subassembly's other placements (M12-F06 independent solve), e.g.
// SetFlexibleChild(wire.SetFlexibleChildArgs{Occurrence: id, Child: "arm:1", Transform: m}).
//
// mcp:tool set_flexible_child
// mcp:summary Position a child component (child: its instance name) within a flexible subassembly occurrence (occurrence id) to a row-major 4×4 transform (16 cells) — independently of the subassembly's other placements. The occurrence must be flexible. Returns the occurrence's refreshed info.
// mcp:input setFlexibleChildArg
func (a Assembly) SetFlexibleChild(args wire.SetFlexibleChildArgs) (wire.OccurrenceResult, error) {
	return call[wire.OccurrenceResult](a.c, wire.MethodAssemblySetFlexibleChild, args)
}

// Replace swaps the occurrence's component for the one held by the open document (by id),
// keeping the occurrence's id/name/transform/state, e.g. Replace(occID, docID).
//
// mcp:tool replace_occurrence
// mcp:summary Swap the component of an occurrence (id) for the one held by an open document (document: its id), keeping the occurrence's id, name, transform, and state — the replace-component operation.
func (a Assembly) Replace(id, document uint64) (wire.OccurrenceResult, error) {
	return call[wire.OccurrenceResult](a.c, wire.MethodAssemblyReplace, wire.ReplaceOccurrenceArgs{ID: id, Document: document})
}

// Remove deletes the occurrence and returns the refreshed tree, e.g. Remove(id).
//
// mcp:tool remove_occurrence
// mcp:summary Delete an occurrence (id) from the active assembly. Returns the remaining occurrence tree.
func (a Assembly) Remove(id uint64) (wire.OccurrencesResult, error) {
	return call[wire.OccurrencesResult](a.c, wire.MethodAssemblyRemove, wire.RemoveOccurrenceArgs{ID: id})
}
