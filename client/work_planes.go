// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// WorkPlanes is the datum-plane construction group for the active part: List enumerates
// the part's planes, Create is the general constructor, and the typed helpers wrap each
// datum-plane constructor. References are work-feature reference strings —
// origin constants (types.WorkRefXYPlane …), refs returned by List, or a face reference
// for the tangent helpers.
type WorkPlanes struct{ c *Client }

// WorkPlanes returns the work-plane construction group.
func (c *Client) WorkPlanes() WorkPlanes { return WorkPlanes{c} }

// List returns the part's datum planes (origin frame first, then user planes).
func (w WorkPlanes) List() (wire.ListWorkPlanesResult, error) {
	var r wire.ListWorkPlanesResult
	return r, w.c.call(wire.MethodWorkPlanesList, nil, &r)
}

// Create adds a datum plane from an explicit request — the escape hatch covering every
// kind; prefer the typed helpers below for the common constructors.
func (w WorkPlanes) Create(args wire.CreateWorkPlaneArgs) (wire.CreateWorkPlaneResult, error) {
	var r wire.CreateWorkPlaneResult
	return r, w.c.call(wire.MethodWorkPlanesCreate, args, &r)
}

// Offset adds a plane parallel to base, offset by a unit-bearing distance ("10 mm").
func (w WorkPlanes) Offset(base, distance string) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlaneOffset), Refs: []string{base}, Offset: distance})
}

// ThreePoints adds a plane through three point references.
func (w WorkPlanes) ThreePoints(p1, p2, p3 string) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlaneThreePoints), Refs: []string{p1, p2, p3}})
}

// Fixed adds a plane fixed at origin [x,y,z] with the given X/Y axis direction components.
func (w WorkPlanes) Fixed(origin, xAxis, yAxis []float64) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlaneFixed), Origin: origin, XAxis: xAxis, YAxis: yAxis})
}

// PlaneAndPoint adds a plane parallel to base passing through point.
func (w WorkPlanes) PlaneAndPoint(base, point string) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlanePlaneAndPoint), Refs: []string{base, point}})
}

// TwoPlanes adds the bisecting plane of plane1 and plane2.
func (w WorkPlanes) TwoPlanes(plane1, plane2 string) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlaneTwoPlanes), Refs: []string{plane1, plane2}})
}

// LinePlaneAndAngle adds a plane through line at a unit-bearing angle ("45 deg") to plane.
func (w WorkPlanes) LinePlaneAndAngle(line, plane, angle string) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlaneLinePlaneAngle), Refs: []string{line, plane}, Angle: angle})
}

// TwoLines adds a plane from two line references (line1 is the X axis).
func (w WorkPlanes) TwoLines(line1, line2 string) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlaneTwoLines), Refs: []string{line1, line2}})
}

// NormalToCurve adds a plane through point, normal to curve (a line/axis reference).
func (w WorkPlanes) NormalToCurve(curve, point string) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlaneNormalToCurve), Refs: []string{curve, point}})
}

// TorusMidPlane adds the mid-plane of a torus face reference.
func (w WorkPlanes) TorusMidPlane(face string) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlaneTorusMidPlane), Refs: []string{face}})
}

// PointAndTangent adds the tangent plane of a surface (face) at a point on it.
func (w WorkPlanes) PointAndTangent(point, face string) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlanePointAndTangent), Refs: []string{point, face}})
}

// PlaneAndTangent adds a plane parallel to base and tangent to a surface (face).
func (w WorkPlanes) PlaneAndTangent(base, face string) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlanePlaneAndTangent), Refs: []string{base, face}})
}

// LineAndTangent adds a plane through line and tangent to a surface (face).
func (w WorkPlanes) LineAndTangent(line, face string) (wire.CreateWorkPlaneResult, error) {
	return w.Create(wire.CreateWorkPlaneArgs{Kind: string(types.WorkPlaneLineAndTangent), Refs: []string{line, face}})
}
