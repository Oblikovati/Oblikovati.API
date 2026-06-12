// SPDX-License-Identifier: Apache-2.0

package wire

// Freeform cage editing (M10-F03 PBI-114, Oblikovati#699): edit a placed
// freeform (sub-D) feature's control cage after placement. The feature is
// addressed by the stable id from [FeatureInfo] (model.tree), like the
// features.* lifecycle methods; every edit recomputes the part and returns the
// refreshed [FeatureDetailResult].

// SetFreeformLevelArgs is the request of [MethodFreeformSetLevel]: change the
// Catmull–Clark subdivision level the cage is evaluated at (clamped at 0 by the
// host — higher levels refine the limit-surface approximation).
type SetFreeformLevelArgs struct {
	ID    uint64 `json:"id"`
	Level int    `json:"level"`
}

// MoveFreeformVerticesArgs is the request of [MethodFreeformMoveVertices]:
// translate the selected cage vertices (by cage index) by [dx,dy,dz] in
// document units. Indices outside the cage are rejected with the offending
// index and the cage size.
type MoveFreeformVerticesArgs struct {
	ID          uint64     `json:"id"`
	Vertices    []int      `json:"vertices"`
	Translation [3]float64 `json:"translation"`
}

// CreaseFreeformEdgesArgs is the request of [MethodFreeformCreaseEdges]: set
// the crease sharpness on the selected cage edges, each addressed by its two
// cage vertex indices. Sharpness runs 0 (smooth) to 1 (fully sharp).
type CreaseFreeformEdgesArgs struct {
	ID        uint64   `json:"id"`
	Edges     [][2]int `json:"edges"`
	Sharpness float64  `json:"sharpness"`
}
