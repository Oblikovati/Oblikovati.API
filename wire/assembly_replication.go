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
