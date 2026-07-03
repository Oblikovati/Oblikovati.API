// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The assembly replication operations (M11-F04, Oblikovati/Oblikovati#729) extend the
// Assembly group: pattern, mirror, copy, and substitute placed components by session id.
// The additive ops return the occurrences they created; substitute returns the substitute
// occurrence.

// PatternCreate replicates the seed occurrence across an arrangement, e.g. a 4-up circular
// pattern: PatternCreate(wire.CreatePatternArgs{Seed: id, Kind: "circular", Axis: [3]float64{0, 0, 1}, Angle: math.Pi / 2, Count: 4}).
//
// mcp:tool assembly_pattern_create
// mcp:summary Replicates the seed occurrence across an arrangement, e.g.
func (a Assembly) PatternCreate(args wire.CreatePatternArgs) (wire.NewOccurrencesResult, error) {
	return call[wire.NewOccurrencesResult](a.c, wire.MethodAssemblyPatternCreate, args)
}

// Mirror adds a mirror of each source occurrence across the plane (origin, normal), e.g.
// Mirror(wire.MirrorComponentsArgs{Sources: []uint64{id}, Normal: [3]float64{1, 0, 0}}).
//
// mcp:tool assembly_mirror
// mcp:summary Adds a mirror of each source occurrence across the plane (origin, normal), e.g.
func (a Assembly) Mirror(args wire.MirrorComponentsArgs) (wire.NewOccurrencesResult, error) {
	return call[wire.NewOccurrencesResult](a.c, wire.MethodAssemblyMirror, args)
}

// MirrorIntoPart mirrors each source occurrence into a NEW opposite-hand part document and
// places it, e.g. MirrorIntoPart(wire.MirrorIntoPartArgs{Sources: []uint64{id}, Normal: [3]float64{1, 0, 0}}).
//
// mcp:tool assembly_mirror_into_part
// mcp:summary Mirrors each source occurrence into a NEW opposite-hand part document and places it, e.g.
func (a Assembly) MirrorIntoPart(args wire.MirrorIntoPartArgs) (wire.NewOccurrencesResult, error) {
	return call[wire.NewOccurrencesResult](a.c, wire.MethodAssemblyMirrorIntoPart, args)
}

// Copy adds an independent copy of each source occurrence, e.g. Copy(id1, id2).
//
// mcp:tool assembly_copy
// mcp:summary Adds an independent copy of each source occurrence, e.g.
func (a Assembly) Copy(sources ...uint64) (wire.NewOccurrencesResult, error) {
	return call[wire.NewOccurrencesResult](a.c, wire.MethodAssemblyCopy, wire.CopyComponentsArgs{Sources: sources})
}

// Substitute suppresses the source occurrences and adds one occurrence instancing the
// simplified component held by the open document, e.g.
// Substitute(wire.SubstituteComponentsArgs{Sources: ids, Document: docID, Name: "lod:1", Transform: t}).
//
// mcp:tool assembly_substitute
// mcp:summary Suppresses the source occurrences and adds one occurrence instancing the simplified component held by the open document, e.g.
func (a Assembly) Substitute(args wire.SubstituteComponentsArgs) (wire.OccurrenceResult, error) {
	return call[wire.OccurrenceResult](a.c, wire.MethodAssemblySubstitute, args)
}
