// SPDX-License-Identifier: Apache-2.0

package wire

// CreateWorkPointArgs is the request of [MethodWorkPointsCreate]: which constructor (Kind, a
// [oblikovati.org/api/types.WorkPointKind] value) and its inputs. An empty Kind means the
// position constructor, so the original position-only request (just At) is unchanged.
//
//   - position (default): At is the point [x, y, z] in model units.
//   - plane-axis-intersection: Refs = [plane, axis] — the point where the axis pierces the
//     plane. At is ignored.
type CreateWorkPointArgs struct {
	At   []float64 `json:"at,omitempty"`   // position kind: [x, y, z] in model units
	Kind string    `json:"kind,omitempty"` // a types.WorkPointKind value (empty = position)
	Refs []string  `json:"refs,omitempty"` // reference-model kinds: plane-axis-intersection = [plane, axis]
}

// CreateWorkPointResult is the response of [MethodWorkPointsCreate]: the new point's index in
// the work-points collection, its stable reference (usable as a point input to a work plane —
// e.g. a three-point plane — or to [MethodWorkPlanesRedefine]'s Repick), and its name. Healthy
// is false (with Reason) for a well-formed but unsatisfiable definition — e.g. a
// plane-axis-intersection whose axis is parallel to the plane; a position point is always
// healthy.
type CreateWorkPointResult struct {
	Index   int    `json:"index"`
	Ref     string `json:"ref"`
	Name    string `json:"name"`
	Healthy bool   `json:"healthy"`
	Reason  string `json:"reason,omitempty"` // empty when healthy
}
