// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The assembly replication operations (M11-F04, Oblikovati/Oblikovati#729) extend the
// Assembly group: pattern, mirror, copy, and substitute placed components by session id.
// The additive ops return the occurrences they created; substitute returns the substitute
// occurrence.

// PatternCreate replicates the seed occurrence across an arrangement into a PERSISTENT pattern,
// e.g. a 4-up circular pattern: PatternCreate(wire.CreatePatternArgs{Seed: id, Kind: "circular", Axis: [3]float64{0, 0, 1}, Angle: math.Pi / 2, Count: 4}).
// The reply carries the pattern's id so it can be re-read and edited (PatternList, PatternSetSuppressed, …).
//
// mcp:tool assembly_pattern_create
// mcp:summary Replicates the seed occurrence across an arrangement into a persistent, editable pattern, e.g.
func (a Assembly) PatternCreate(args wire.CreatePatternArgs) (wire.CreatePatternResult, error) {
	return call[wire.CreatePatternResult](a.c, wire.MethodAssemblyPatternCreate, args)
}

// PatternList returns every persistent occurrence pattern in the active assembly, e.g. PatternList().
//
// mcp:tool assembly_pattern_list
// mcp:summary Returns every persistent occurrence pattern in the active assembly.
func (a Assembly) PatternList() (wire.PatternListResult, error) {
	return call[wire.PatternListResult](a.c, wire.MethodAssemblyPatternList, struct{}{})
}

// PatternSetSuppressed suppresses or unsuppresses a whole pattern by id, moving every element
// together, e.g. PatternSetSuppressed(wire.SetPatternSuppressedArgs{Pattern: id, Suppressed: true}).
//
// mcp:tool assembly_pattern_set_suppressed
// mcp:summary Suppresses or unsuppresses a whole pattern by id, moving every element together, e.g.
func (a Assembly) PatternSetSuppressed(args wire.SetPatternSuppressedArgs) (wire.PatternInfo, error) {
	return call[wire.PatternInfo](a.c, wire.MethodAssemblyPatternSetSuppressed, args)
}

// PatternElementSetSuppressed suppresses or unsuppresses one element (by index) of a pattern,
// e.g. PatternElementSetSuppressed(wire.SetPatternElementSuppressedArgs{Pattern: id, Element: 2, Suppressed: true}).
//
// mcp:tool assembly_pattern_element_set_suppressed
// mcp:summary Suppresses or unsuppresses one element (by index) of a pattern, e.g.
func (a Assembly) PatternElementSetSuppressed(args wire.SetPatternElementSuppressedArgs) (wire.PatternInfo, error) {
	return call[wire.PatternInfo](a.c, wire.MethodAssemblyPatternElementSuppress, args)
}

// PatternElementReposition moves one element (by index) of a pattern to an explicit placement,
// e.g. PatternElementReposition(wire.RepositionPatternElementArgs{Pattern: id, Element: 2, Transform: t}).
//
// mcp:tool assembly_pattern_element_reposition
// mcp:summary Moves one element (by index) of a pattern to an explicit placement, off the regular grid, e.g.
func (a Assembly) PatternElementReposition(args wire.RepositionPatternElementArgs) (wire.PatternInfo, error) {
	return call[wire.PatternInfo](a.c, wire.MethodAssemblyPatternElementReposition, args)
}

// PatternDelete deletes a whole pattern by id, removing the occurrences it generated (the seed
// stays), e.g. PatternDelete(id).
//
// mcp:tool assembly_pattern_delete
// mcp:summary Deletes a whole pattern by id, removing the occurrences it generated (the seed stays), e.g.
func (a Assembly) PatternDelete(pattern uint64) (wire.DeletePatternResult, error) {
	return call[wire.DeletePatternResult](a.c, wire.MethodAssemblyPatternDelete, wire.DeletePatternArgs{Pattern: pattern})
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
