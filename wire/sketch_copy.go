// SPDX-License-Identifier: Apache-2.0

package wire

// Sketch-to-sketch copy (Oblikovati/Oblikovati#151). Copy geometry from one 2D sketch into
// another — reusing a profile across features and planes (Inventor CopyContentsTo /
// CopyEntitiesTo). The source entities' sketch-local coordinates are re-instantiated in the
// target sketch's plane, optionally offset by Position.

// CopySketchArgs is the request of [MethodSketchCopyTo]. SourceIndex/TargetIndex are the
// sketches' indices. EntityIDs selects the source entities to copy (from sketch.entities);
// empty copies the whole sketch's contents. Position (target-plane [x, y] in cm) offsets the
// copies; absent (nil) copies in place.
type CopySketchArgs struct {
	SourceIndex int       `json:"sourceIndex"`
	TargetIndex int       `json:"targetIndex"`
	EntityIDs   []uint64  `json:"entityIds,omitempty"`
	Position    []float64 `json:"position,omitempty"`
}

// CopySketchResult is the response of [MethodSketchCopyTo]: the session ids of the entities
// created in the target sketch, and their count. Only geometry is copied; constraints and
// dimensions are not carried over (an external-reference-free copy).
type CopySketchResult struct {
	Created []uint64 `json:"created"`
	Count   int      `json:"count"`
}
