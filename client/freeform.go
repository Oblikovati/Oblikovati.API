// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Freeform is the freeform (sub-D) cage-editing operation group (M10-F03
// PBI-114, Oblikovati#699). Every edit recomputes the part and returns the
// refreshed feature detail.
type Freeform struct{ c *Client }

// Freeform returns the freeform cage-editing operation group.
func (c *Client) Freeform() Freeform { return Freeform{c} }

// SetLevel changes the subdivision level a placed freeform feature's cage is
// evaluated at, e.g. SetLevel(7, 2).
func (f Freeform) SetLevel(id uint64, level int) (wire.FeatureDetailResult, error) {
	var r wire.FeatureDetailResult
	args := wire.SetFreeformLevelArgs{ID: id, Level: level}
	return r, f.c.call(wire.MethodFreeformSetLevel, args, &r)
}

// MoveVertices translates the selected cage vertices (by cage index) by
// [dx,dy,dz] in document units, e.g.
// MoveVertices(wire.MoveFreeformVerticesArgs{ID: 7, Vertices: []int{0}, Translation: [3]float64{1, 0, 0}}).
func (f Freeform) MoveVertices(args wire.MoveFreeformVerticesArgs) (wire.FeatureDetailResult, error) {
	var r wire.FeatureDetailResult
	return r, f.c.call(wire.MethodFreeformMoveVertices, args, &r)
}

// CreaseEdges sets the crease sharpness (0 smooth … 1 fully sharp) on the
// selected cage edges, each addressed by its two cage vertex indices, e.g.
// CreaseEdges(wire.CreaseFreeformEdgesArgs{ID: 7, Edges: [][2]int{{0, 1}}, Sharpness: 1}).
func (f Freeform) CreaseEdges(args wire.CreaseFreeformEdgesArgs) (wire.FeatureDetailResult, error) {
	var r wire.FeatureDetailResult
	return r, f.c.call(wire.MethodFreeformCreaseEdges, args, &r)
}
