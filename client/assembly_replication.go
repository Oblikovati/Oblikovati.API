// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The assembly replication operations (M11-F04, Oblikovati/Oblikovati#729) extend the
// Assembly group: pattern, mirror, copy, and substitute placed components by session id.
// The additive ops return the occurrences they created; substitute returns the substitute
// occurrence.

// PatternCreate replicates the seed occurrence across an arrangement, e.g. a 4-up circular
// pattern: PatternCreate(wire.CreatePatternArgs{Seed: id, Kind: "circular", Axis: [3]float64{0, 0, 1}, Angle: math.Pi / 2, Count: 4}).
func (a Assembly) PatternCreate(args wire.CreatePatternArgs) (wire.NewOccurrencesResult, error) {
	var r wire.NewOccurrencesResult
	return r, a.c.call(wire.MethodAssemblyPatternCreate, args, &r)
}

// Mirror adds a mirror of each source occurrence across the plane (origin, normal), e.g.
// Mirror(wire.MirrorComponentsArgs{Sources: []uint64{id}, Normal: [3]float64{1, 0, 0}}).
func (a Assembly) Mirror(args wire.MirrorComponentsArgs) (wire.NewOccurrencesResult, error) {
	var r wire.NewOccurrencesResult
	return r, a.c.call(wire.MethodAssemblyMirror, args, &r)
}

// Copy adds an independent copy of each source occurrence, e.g. Copy(id1, id2).
func (a Assembly) Copy(sources ...uint64) (wire.NewOccurrencesResult, error) {
	var r wire.NewOccurrencesResult
	return r, a.c.call(wire.MethodAssemblyCopy, wire.CopyComponentsArgs{Sources: sources}, &r)
}

// Substitute suppresses the source occurrences and adds one occurrence instancing the
// simplified component held by the open document, e.g.
// Substitute(wire.SubstituteComponentsArgs{Sources: ids, Document: docID, Name: "lod:1", Transform: t}).
func (a Assembly) Substitute(args wire.SubstituteComponentsArgs) (wire.OccurrenceResult, error) {
	var r wire.OccurrenceResult
	return r, a.c.call(wire.MethodAssemblySubstitute, args, &r)
}
