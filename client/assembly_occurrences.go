// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The assembly occurrence operations (M11-F01/F02, Oblikovati/Oblikovati#728) extend the
// Assembly group: read the active assembly's occurrence tree and place/transform/ground/
// suppress/replace/remove components, addressed by session id. Each single-occurrence
// operation returns the affected occurrence's refreshed info; remove returns the tree.

// Occurrences returns the active assembly's occurrence tree.
func (a Assembly) Occurrences() (wire.OccurrencesResult, error) {
	var r wire.OccurrencesResult
	return r, a.c.call(wire.MethodAssemblyOccurrences, struct{}{}, &r)
}

// Place places the component held by the open document (by id) into the active assembly,
// e.g. Place(wire.PlaceOccurrenceArgs{Document: docID, Name: "pin:1", Transform: t}).
func (a Assembly) Place(args wire.PlaceOccurrenceArgs) (wire.OccurrenceResult, error) {
	var r wire.OccurrenceResult
	return r, a.c.call(wire.MethodAssemblyPlace, args, &r)
}

// PlaceByDefinition places another instance of the component that the source occurrence
// already instances, e.g. PlaceByDefinition(wire.PlaceByDefinitionArgs{Source: occID, Name: "pin:2", Transform: t}).
func (a Assembly) PlaceByDefinition(args wire.PlaceByDefinitionArgs) (wire.OccurrenceResult, error) {
	var r wire.OccurrenceResult
	return r, a.c.call(wire.MethodAssemblyPlaceByDefinition, args, &r)
}

// Transform repositions the occurrence, e.g. Transform(wire.TransformOccurrenceArgs{ID: id, Transform: t}).
func (a Assembly) Transform(args wire.TransformOccurrenceArgs) (wire.OccurrenceResult, error) {
	var r wire.OccurrenceResult
	return r, a.c.call(wire.MethodAssemblyTransform, args, &r)
}

// Ground fixes or releases the occurrence in space, e.g. Ground(id, true).
func (a Assembly) Ground(id uint64, grounded bool) (wire.OccurrenceResult, error) {
	var r wire.OccurrenceResult
	return r, a.c.call(wire.MethodAssemblyGround, wire.GroundOccurrenceArgs{ID: id, Grounded: grounded}, &r)
}

// Suppress excludes or restores the occurrence from the model, e.g. Suppress(id, true).
func (a Assembly) Suppress(id uint64, suppressed bool) (wire.OccurrenceResult, error) {
	var r wire.OccurrenceResult
	return r, a.c.call(wire.MethodAssemblySuppress, wire.SuppressOccurrenceArgs{ID: id, Suppressed: suppressed}, &r)
}

// Replace swaps the occurrence's component for the one held by the open document (by id),
// keeping the occurrence's id/name/transform/state, e.g. Replace(occID, docID).
func (a Assembly) Replace(id, document uint64) (wire.OccurrenceResult, error) {
	var r wire.OccurrenceResult
	return r, a.c.call(wire.MethodAssemblyReplace, wire.ReplaceOccurrenceArgs{ID: id, Document: document}, &r)
}

// Remove deletes the occurrence and returns the refreshed tree, e.g. Remove(id).
func (a Assembly) Remove(id uint64) (wire.OccurrencesResult, error) {
	var r wire.OccurrencesResult
	return r, a.c.call(wire.MethodAssemblyRemove, wire.RemoveOccurrenceArgs{ID: id}, &r)
}
