// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati/api/types"
	"oblikovati/api/wire"
)

// AddDimension is the general 3D dimensional-constraint constructor; prefer the typed
// helpers below for the common kinds.
func (s Sketch3D) AddDimension(args wire.AddSketch3DDimensionArgs) (wire.AddSketch3DDimensionResult, error) {
	var r wire.AddSketch3DDimensionResult
	return r, s.c.call(wire.MethodSketch3DAddDimension, args, &r)
}

// Distance dimensions the distance between two 3D points to a unit-bearing value.
func (s Sketch3D) Distance(index int, a, b uint64, value string) (wire.AddSketch3DDimensionResult, error) {
	return s.AddDimension(wire.AddSketch3DDimensionArgs{
		SketchIndex: index, Kind: string(types.Dim3DDistance), Entities: []uint64{a, b}, Expression: value,
	})
}

// LineLength dimensions a 3D line's length; Radius dimensions a 3D circle's radius.
func (s Sketch3D) LineLength(index int, line uint64, value string) (wire.AddSketch3DDimensionResult, error) {
	return s.AddDimension(wire.AddSketch3DDimensionArgs{
		SketchIndex: index, Kind: string(types.Dim3DLineLength), Entities: []uint64{line}, Expression: value,
	})
}

func (s Sketch3D) Radius(index int, circle uint64, value string) (wire.AddSketch3DDimensionResult, error) {
	return s.AddDimension(wire.AddSketch3DDimensionArgs{
		SketchIndex: index, Kind: string(types.Dim3DRadius), Entities: []uint64{circle}, Expression: value,
	})
}

// PointPlaneDistance dimensions the distance from a point to an origin plane ("XY"|"XZ"|"YZ").
func (s Sketch3D) PointPlaneDistance(index int, point uint64, plane, value string) (wire.AddSketch3DDimensionResult, error) {
	return s.AddDimension(wire.AddSketch3DDimensionArgs{
		SketchIndex: index, Kind: string(types.Dim3DPointPlaneDistance), Entities: []uint64{point}, Plane: plane, Expression: value,
	})
}

// TwoLineAngle dimensions the angle between two 3D lines to a unit-bearing angle.
func (s Sketch3D) TwoLineAngle(index int, l1, l2 uint64, value string) (wire.AddSketch3DDimensionResult, error) {
	return s.AddDimension(wire.AddSketch3DDimensionArgs{
		SketchIndex: index, Kind: string(types.Dim3DTwoLineAngle), Entities: []uint64{l1, l2}, Expression: value,
	})
}

// DriveDimension edits a 3D dimension's value and/or driving state.
func (s Sketch3D) DriveDimension(args wire.DriveSketch3DDimensionArgs) (wire.OKResult, error) {
	var r wire.OKResult
	return r, s.c.call(wire.MethodSketch3DDriveDimension, args, &r)
}
