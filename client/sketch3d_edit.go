// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati/api/wire"

// Transform applies an editing operation (move/copy/rotate/delete) to a 3D-sketch
// selection; prefer the typed helpers below.
func (s Sketch3D) Transform(args wire.Transform3DArgs) (wire.Transform3DResult, error) {
	var r wire.Transform3DResult
	return r, s.c.call(wire.MethodSketch3DTransform, args, &r)
}

// Move translates the entities by vector [x,y,z] (cm); Copy duplicates them translated.
func (s Sketch3D) Move(index int, entities []uint64, vector []float64) (wire.Transform3DResult, error) {
	return s.Transform(wire.Transform3DArgs{SketchIndex: index, Op: "move", Entities: entities, Vector: vector})
}

func (s Sketch3D) Copy(index int, entities []uint64, vector []float64) (wire.Transform3DResult, error) {
	return s.Transform(wire.Transform3DArgs{SketchIndex: index, Op: "copy", Entities: entities, Vector: vector})
}

// Rotate rotates the entities by a unit-bearing angle about the axis through center in
// direction axis (empty axis ⇒ +Z).
func (s Sketch3D) Rotate(index int, entities []uint64, center, axis []float64, angle string) (wire.Transform3DResult, error) {
	return s.Transform(wire.Transform3DArgs{SketchIndex: index, Op: "rotate", Entities: entities, Center: center, Axis: axis, Angle: angle})
}

// DeleteEntities removes the given entities from the 3D sketch.
func (s Sketch3D) DeleteEntities(index int, entities []uint64) (wire.Transform3DResult, error) {
	return s.Transform(wire.Transform3DArgs{SketchIndex: index, Op: "delete", Entities: entities})
}
