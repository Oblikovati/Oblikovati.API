// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

// Sketch is the sketch-authoring operation group for the active part.
type Sketch struct{ c *Client }

// Sketch returns the sketch operation group.
func (c *Client) Sketch() Sketch { return Sketch{c} }

// Create adds a sketch on an origin plane and returns its index.
func (s Sketch) Create(args wire.CreateSketchArgs) (wire.CreateSketchResult, error) {
	var r wire.CreateSketchResult
	return r, s.c.call(wire.MethodSketchCreate, args, &r)
}

// Rectangle adds a closed rectangle (one profile) to a sketch.
func (s Sketch) Rectangle(args wire.SketchRectangleArgs) (wire.SketchRectangleResult, error) {
	var r wire.SketchRectangleResult
	return r, s.c.call(wire.MethodSketchRectangle, args, &r)
}
