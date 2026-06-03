// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

// Offset offsets a single line/circle/arc (by entity id) by a signed unit-bearing
// distance, returning the new entity's id and kind.
func (s Sketch) Offset(index int, entity uint64, distance string) (wire.OffsetSketchResult, error) {
	var r wire.OffsetSketchResult
	args := wire.OffsetSketchArgs{SketchIndex: index, Entity: entity, Distance: distance}
	return r, s.c.call(wire.MethodSketchOffset, args, &r)
}

// AddImage places a raster image (ref is a package-store reference) anchored at [x,y] cm
// with unit-bearing width/height; rotation and opacity are optional ("" / 0).
func (s Sketch) AddImage(index int, ref string, anchor []float64, width, height, rotation string, opacity float64) (wire.AddSketchImageResult, error) {
	var r wire.AddSketchImageResult
	args := wire.AddSketchImageArgs{
		SketchIndex: index, Ref: ref, Anchor: anchor,
		Width: width, Height: height, Rotation: rotation, Opacity: opacity,
	}
	return r, s.c.call(wire.MethodSketchAddImage, args, &r)
}
