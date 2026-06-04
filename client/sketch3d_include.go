// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

// Include links the referenced part edges/vertices (by reference key) into the 3D sketch
// as associative reference geometry, returning the created entity ids and whether every
// reference resolved.
func (s Sketch3D) Include(index int, refs []string) (wire.IncludeSketch3DResult, error) {
	var r wire.IncludeSketch3DResult
	return r, s.c.call(wire.MethodSketch3DInclude, wire.IncludeSketch3DArgs{SketchIndex: index, Refs: refs}, &r)
}
