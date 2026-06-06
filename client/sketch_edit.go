// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati/api/wire"

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

// Trim removes the segment of a line containing the pick point ([x,y] cm), cutting at the
// nearest crossings; returns the surviving line(s).
func (s Sketch) Trim(index int, line uint64, pick []float64) (wire.TransformSketchResult, error) {
	return s.transform(wire.TransformSketchArgs{
		SketchIndex: index, Op: "trim", Entities: []uint64{line}, Vector: pick,
	})
}

// Split splits a line at the pick point ([x,y] cm) into two; returns the resulting lines.
func (s Sketch) Split(index int, line uint64, pick []float64) (wire.TransformSketchResult, error) {
	return s.transform(wire.TransformSketchArgs{
		SketchIndex: index, Op: "split", Entities: []uint64{line}, Vector: pick,
	})
}

// Extend lengthens the end of a line nearest the pick point ([x,y] cm) to the next crossing.
func (s Sketch) Extend(index int, line uint64, pick []float64) (wire.TransformSketchResult, error) {
	return s.transform(wire.TransformSketchArgs{
		SketchIndex: index, Op: "extend", Entities: []uint64{line}, Vector: pick,
	})
}

func (s Sketch) transform(args wire.TransformSketchArgs) (wire.TransformSketchResult, error) {
	var r wire.TransformSketchResult
	return r, s.c.call(wire.MethodSketchTransform, args, &r)
}
