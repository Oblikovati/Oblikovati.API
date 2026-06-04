// SPDX-License-Identifier: Apache-2.0

package contract

// ClientGraphics is the scalar view of one client-graphics group the host implements
// (model/clientgraphics.Group) and in-process callers consume. It exposes only the
// scalar surface — the geometry itself travels as wire DTOs (see api/wire). The host
// satisfies it with a compile-time assertion.
type ClientGraphics interface {
	// Name is the group's client id (the key it was submitted under).
	Name() string
	// Lane is the display lane the group lives in (a types.GraphicsLane value).
	Lane() string
	// Visible reports whether the group is currently drawn.
	Visible() bool
	// NodeCount is the number of graphics nodes the group owns.
	NodeCount() int
}
