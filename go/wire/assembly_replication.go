// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The assembly replication surface (M11-F04, Oblikovati/Oblikovati#729): replicate placed
// components by session id — pattern (circular/rectangular), mirror across a plane,
// independent copy, and substitute a set with one simplified representation. Each op adds
// occurrences to the active assembly and replies with the occurrences it created.

// NewOccurrencesResult is the reply of the additive replication ops (pattern, mirror,
// copy): the occurrences they added, in element/source order.
type NewOccurrencesResult struct {
	Created []OccurrenceInfo `json:"created"`
}

// CreatePatternArgs is the request of [MethodAssemblyPatternCreate]: replicate the Seed
// occurrence (by session id) across an arrangement, adding one occurrence per generated
// element beyond the seed. Kind selects the arrangement and the fields it reads:
//   - "circular": Origin, Axis, Angle (radians between adjacent elements), Count (total,
//     including the seed);
//   - "rectangular": Direction1/Spacing1/Count1 and Direction2/Spacing2/Count2 (a
//     Count1×Count2 grid; the seed is element 0,0).
type CreatePatternArgs struct {
	Seed uint64 `json:"seed"`
	Kind string `json:"kind"`

	// Circular arrangement.
	Origin [3]float64 `json:"origin,omitempty"`
	Axis   [3]float64 `json:"axis,omitempty"`
	Angle  float64    `json:"angle,omitempty"`
	Count  int        `json:"count,omitempty"`

	// Rectangular arrangement.
	Direction1 [3]float64 `json:"direction1,omitempty"`
	Spacing1   float64    `json:"spacing1,omitempty"`
	Count1     int        `json:"count1,omitempty"`
	Direction2 [3]float64 `json:"direction2,omitempty"`
	Spacing2   float64    `json:"spacing2,omitempty"`
	Count2     int        `json:"count2,omitempty"`
}

// PatternElementInfo is one element of a persistent occurrence pattern (#1976): its position
// in arrangement order, whether it is suppressed, and whether its placement was individually
// repositioned off the regular grid. Element 0 is the seed.
type PatternElementInfo struct {
	Index        int  `json:"index"`
	Suppressed   bool `json:"suppressed"`
	Repositioned bool `json:"repositioned"`
}

// PatternInfo is a persistent occurrence pattern (#1976): its session id, name, arrangement
// kind ("circular"/"rectangular"), the whole-pattern suppression state ("all"/"none"/"some"),
// and its elements. It is returned by patternCreate and the edit ops, and listed by patternList.
type PatternInfo struct {
	ID          uint64               `json:"id"`
	Name        string               `json:"name"`
	Kind        string               `json:"kind"`
	Suppression string               `json:"suppression"`
	Elements    []PatternElementInfo `json:"elements"`
}

// CreatePatternResult is the reply of [MethodAssemblyPatternCreate]: the persistent pattern
// (by id, so it can be re-read and edited) and the occurrences it added, in element order.
type CreatePatternResult struct {
	Pattern PatternInfo      `json:"pattern"`
	Created []OccurrenceInfo `json:"created"`
}

// PatternListResult is the reply of [MethodAssemblyPatternList]: every persistent pattern in
// the active assembly.
type PatternListResult struct {
	Patterns []PatternInfo `json:"patterns"`
}

// SetPatternSuppressedArgs is the request of [MethodAssemblyPatternSetSuppressed]: suppress or
// unsuppress the whole pattern (by id), moving every element together.
type SetPatternSuppressedArgs struct {
	Pattern    uint64 `json:"pattern"`
	Suppressed bool   `json:"suppressed"`
}

// SetPatternElementSuppressedArgs is the request of [MethodAssemblyPatternElementSetSuppressed]:
// suppress or unsuppress one element (by index) of the pattern (by id).
type SetPatternElementSuppressedArgs struct {
	Pattern    uint64 `json:"pattern"`
	Element    int    `json:"element"`
	Suppressed bool   `json:"suppressed"`
}

// RepositionPatternElementArgs is the request of [MethodAssemblyPatternElementReposition]: move
// one element (by index) of the pattern (by id) to an explicit placement, off the regular grid.
type RepositionPatternElementArgs struct {
	Pattern   uint64       `json:"pattern"`
	Element   int          `json:"element"`
	Transform types.Matrix `json:"transform"`
}

// DeletePatternArgs is the request of [MethodAssemblyPatternDelete]: delete the whole pattern
// (by id) and remove the occurrences it generated (the seed component stays).
type DeletePatternArgs struct {
	Pattern uint64 `json:"pattern"`
}

// DeletePatternResult is the reply of [MethodAssemblyPatternDelete]: the id of the deleted pattern.
type DeletePatternResult struct {
	Deleted uint64 `json:"deleted"`
}

// MirrorComponentsArgs is the request of [MethodAssemblyMirror]: add a mirror of each
// Source occurrence (by session id), reflected across the plane through Origin with unit
// Normal. Each mirror shares its source's component, handed by the reflection transform.
type MirrorComponentsArgs struct {
	Sources []uint64   `json:"sources"`
	Origin  [3]float64 `json:"origin"`
	Normal  [3]float64 `json:"normal"`
}

// MirrorIntoPartArgs is the request of [MethodAssemblyMirrorIntoPart]: like
// [MethodAssemblyMirror], but for each chiral Source it derives a NEW opposite-hand PART
// document (the source geometry reflected across the plane through Origin with unit Normal,
// baked into its own editable, separately-openable part) and places THAT as the mirrored
// occurrence — rather than sharing the source definition handed only by the placement. The
// reply is the created occurrences (each instancing a freshly derived mirror part). A
// Source that does not reference a saved component document is rejected (there is no source
// document to derive from).
type MirrorIntoPartArgs struct {
	Sources []uint64   `json:"sources"`
	Origin  [3]float64 `json:"origin"`
	Normal  [3]float64 `json:"normal"`
}

// CopyComponentsArgs is the request of [MethodAssemblyCopy]: add an independent copy of
// each Source occurrence (by session id) — same component and placement, a new instance.
type CopyComponentsArgs struct {
	Sources []uint64 `json:"sources"`
}

// SubstituteComponentsArgs is the request of [MethodAssemblySubstitute]: suppress the
// Source occurrences (by session id) and add one occurrence, named Name at Transform, that
// instances the simplified component held by the open Document (by document id) — the
// "substitute representation" of the sources.
type SubstituteComponentsArgs struct {
	Sources   []uint64     `json:"sources"`
	Document  uint64       `json:"document"`
	Name      string       `json:"name"`
	Transform types.Matrix `json:"transform"`
}
