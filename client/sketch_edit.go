// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

// Move translates a selection of entities in place by [dx,dy] (cm).
func (s Sketch) Move(index int, entities []uint64, dx, dy float64) (wire.TransformSketchResult, error) {
	return s.transform(wire.TransformSketchArgs{
		SketchIndex: index, Op: "move", Entities: entities, Vector: []float64{dx, dy},
	})
}

// Copy duplicates a selection offset by [dx,dy] (cm), returning the new entity ids.
func (s Sketch) Copy(index int, entities []uint64, dx, dy float64) (wire.TransformSketchResult, error) {
	return s.transform(wire.TransformSketchArgs{
		SketchIndex: index, Op: "copy", Entities: entities, Vector: []float64{dx, dy},
	})
}

// Rotate rotates a selection in place about center [x,y] (cm) by a unit-bearing angle
// ("90 deg").
func (s Sketch) Rotate(index int, entities []uint64, center []float64, angle string) (wire.TransformSketchResult, error) {
	return s.transform(wire.TransformSketchArgs{
		SketchIndex: index, Op: "rotate", Entities: entities, Center: center, Angle: angle,
	})
}

// Mirror reflects a selection across the line entity mirrorLine, returning the copies.
func (s Sketch) Mirror(index int, entities []uint64, mirrorLine uint64) (wire.TransformSketchResult, error) {
	return s.transform(wire.TransformSketchArgs{
		SketchIndex: index, Op: "mirror", Entities: entities, MirrorLine: mirrorLine,
	})
}

func (s Sketch) transform(args wire.TransformSketchArgs) (wire.TransformSketchResult, error) {
	var r wire.TransformSketchResult
	return r, s.c.call(wire.MethodSketchTransform, args, &r)
}
