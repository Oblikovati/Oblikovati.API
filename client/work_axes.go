// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// WorkAxes is the datum-axis construction group for the active part or assembly: List
// enumerates the model's axes, Create is the general constructor, and the typed helpers wrap
// each datum-axis constructor. References are work-feature reference strings — origin
// constants (types.WorkRefXAxis …), point refs from create_work_point, or plane refs
// returned by list_work_planes. A created axis's reference can be a revolve/sweep axis or an
// input to a further datum (a normal-to-curve work plane).
type WorkAxes struct{ c *Client }

// WorkAxes returns the work-axis construction group.
func (c *Client) WorkAxes() WorkAxes { return WorkAxes{c} }

// List returns the active model's datum axes (origin axes first, then user axes).
//
// mcp:tool list_work_axes
// mcp:summary List the work axes of the active part or assembly (origin + user). Each axis reports its kind and current geometry: its ref, origin point and unit direction (database units, cm), whether it is an origin axis, and its health.
func (w WorkAxes) List() (wire.ListWorkAxesResult, error) {
	return call[wire.ListWorkAxesResult](w.c, wire.MethodWorkAxesList, nil)
}

// Create adds a datum axis from an explicit request — the escape hatch covering every
// kind; prefer the typed helpers below for the common constructors.
//
// mcp:tool create_work_axis
// mcp:summary Create a work axis (line: a grounded axis from origin+direction; two-points; plane-intersection); see the args schema. Returns its ref (e.g. "axis/1") to use as a revolve/sweep axis or a datum input. Pair with capture_viewport to SEE the datum axis.
func (w WorkAxes) Create(args wire.CreateWorkAxisArgs) (wire.CreateWorkAxisResult, error) {
	return call[wire.CreateWorkAxisResult](w.c, wire.MethodWorkAxesCreate, args)
}

// Line adds a raw grounded axis through origin [x,y,z] along direction [dx,dy,dz]
// (database units, cm) — the axis a revolve spins about when it is not an origin axis.
func (w WorkAxes) Line(origin, direction []float64) (wire.CreateWorkAxisResult, error) {
	return w.Create(wire.CreateWorkAxisArgs{Kind: string(types.WorkAxisLine), Origin: origin, Direction: direction})
}

// TwoPoints adds an axis through two point references.
func (w WorkAxes) TwoPoints(p1, p2 string) (wire.CreateWorkAxisResult, error) {
	return w.Create(wire.CreateWorkAxisArgs{Kind: string(types.WorkAxisTwoPoints), Refs: []string{p1, p2}})
}

// PlaneIntersection adds the axis where two plane references meet.
func (w WorkAxes) PlaneIntersection(plane1, plane2 string) (wire.CreateWorkAxisResult, error) {
	return w.Create(wire.CreateWorkAxisArgs{Kind: string(types.WorkAxisPlaneIntersection), Refs: []string{plane1, plane2}})
}

// PointAndPlane adds the axis through a point reference, normal to a plane reference (#1840).
func (w WorkAxes) PointAndPlane(point, plane string) (wire.CreateWorkAxisResult, error) {
	return w.Create(wire.CreateWorkAxisArgs{Kind: string(types.WorkAxisPointAndPlane), Refs: []string{point, plane}})
}

// LineAndPoint adds the axis through a point reference, parallel to a line reference (#1840).
func (w WorkAxes) LineAndPoint(line, point string) (wire.CreateWorkAxisResult, error) {
	return w.Create(wire.CreateWorkAxisArgs{Kind: string(types.WorkAxisLineAndPoint), Refs: []string{line, point}})
}

// LineAndPlane adds the axis of a line reference projected onto a plane reference (#1840).
func (w WorkAxes) LineAndPlane(line, plane string) (wire.CreateWorkAxisResult, error) {
	return w.Create(wire.CreateWorkAxisArgs{Kind: string(types.WorkAxisLineAndPlane), Refs: []string{line, plane}})
}

// RevolvedFace adds the axis of revolution of a cylindrical, conical, or toroidal face reference —
// the axis a round hole/boss or a revolve sits on. Reports healthy=false for a face with no axis of
// revolution (#1840).
func (w WorkAxes) RevolvedFace(face string) (wire.CreateWorkAxisResult, error) {
	return w.Create(wire.CreateWorkAxisArgs{Kind: string(types.WorkAxisRevolvedFace), Refs: []string{face}})
}
