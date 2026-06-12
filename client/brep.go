// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// TransientBRep is the transient B-rep factory method group (M07-F05,
// Oblikovati/Oblikovati#628): ownerless bodies created and manipulated by
// session handle.
type TransientBRep struct{ c *Client }

// TransientBRep returns the transient B-rep method group.
func (c *Client) TransientBRep() TransientBRep { return TransientBRep{c} }

// CreatePrimitive creates a solid block/cylinderCone/sphere/torus.
func (t TransientBRep) CreatePrimitive(args wire.CreatePrimitiveArgs) (wire.BrepHandleResult, error) {
	var r wire.BrepHandleResult
	return r, t.c.call(wire.MethodBrepCreatePrimitive, args, &r)
}

// DoBoolean combines the blank (modified in place) with the tool.
func (t TransientBRep) DoBoolean(blankHandle int, tool wire.BrepBodyRef, op types.BooleanType) (wire.BrepHandleResult, error) {
	var r wire.BrepHandleResult
	return r, t.c.call(wire.MethodBrepBoolean, wire.BrepBooleanArgs{
		BlankHandle: blankHandle, Tool: tool, Operation: op.String(),
	}, &r)
}

// Transform maps the body by a 4×4 row-major rigid/similarity matrix.
func (t TransientBRep) Transform(handle int, matrix []float64) (wire.BrepHandleResult, error) {
	var r wire.BrepHandleResult
	return r, t.c.call(wire.MethodBrepTransform, wire.BrepTransformArgs{Handle: handle, Matrix: matrix}, &r)
}

// Copy clones a transient or document body into a new transient body.
func (t TransientBRep) Copy(source wire.BrepBodyRef) (wire.BrepHandleResult, error) {
	var r wire.BrepHandleResult
	return r, t.c.call(wire.MethodBrepCopy, wire.BrepCopyArgs{Source: source}, &r)
}

// CreateIntersectionWithPlane sections a body with a plane; the section
// curves come back as wires on a new transient body.
func (t TransientBRep) CreateIntersectionWithPlane(source wire.BrepBodyRef, planeOrigin, planeNormal []float64) (wire.BrepWiresResult, error) {
	var r wire.BrepWiresResult
	return r, t.c.call(wire.MethodBrepSectionWithPlane, wire.BrepSectionArgs{
		Source: source, PlaneOrigin: planeOrigin, PlaneNormal: planeNormal,
	}, &r)
}

// DeleteFaces removes the named faces (or with keepInstead all others)
// without healing.
func (t TransientBRep) DeleteFaces(handle int, faceKeys []string, keepInstead bool) (wire.BrepHandleResult, error) {
	var r wire.BrepHandleResult
	return r, t.c.call(wire.MethodBrepDeleteFaces, wire.BrepDeleteFacesArgs{
		Handle: handle, FaceKeys: faceKeys, KeepInstead: keepInstead,
	}, &r)
}

// CreateSilhouetteCurve traces one face's silhouette from a view direction.
func (t TransientBRep) CreateSilhouetteCurve(args wire.BrepSilhouetteArgs) (wire.BrepWiresResult, error) {
	var r wire.BrepWiresResult
	return r, t.c.call(wire.MethodBrepSilhouette, args, &r)
}

// CreateRuledSurface builds the ruled surface between two wire sections.
func (t TransientBRep) CreateRuledSurface(sectionOne, sectionTwo wire.BrepWireRef) (wire.BrepHandleResult, error) {
	var r wire.BrepHandleResult
	return r, t.c.call(wire.MethodBrepRuledSurface, wire.BrepRuledSurfaceArgs{
		SectionOne: sectionOne, SectionTwo: sectionTwo,
	}, &r)
}

// ImprintBodies face-splits two bodies along their intersections without
// removing material.
func (t TransientBRep) ImprintBodies(args wire.BrepImprintArgs) (wire.BrepImprintResult, error) {
	var r wire.BrepImprintResult
	return r, t.c.call(wire.MethodBrepImprint, args, &r)
}

// GetIdenticalBodies groups congruent bodies.
func (t TransientBRep) GetIdenticalBodies(args wire.BrepIdenticalBodiesArgs) (wire.BrepIdenticalBodiesResult, error) {
	var r wire.BrepIdenticalBodiesResult
	return r, t.c.call(wire.MethodBrepIdenticalBodies, args, &r)
}

// CreateFromDefinition compiles a bottom-up definition graph into a body,
// returning per-definition issues instead when the graph is unsound.
func (t TransientBRep) CreateFromDefinition(def types.BrepBodyDefinition) (wire.BrepCreateFromDefinitionResult, error) {
	var r wire.BrepCreateFromDefinitionResult
	return r, t.c.call(wire.MethodBrepCreateFromDefinition, wire.BrepCreateFromDefinitionArgs{Definition: def}, &r)
}

// Describe returns a transient body's current stats.
func (t TransientBRep) Describe(handle int) (wire.BrepHandleResult, error) {
	var r wire.BrepHandleResult
	return r, t.c.call(wire.MethodBrepDescribe, wire.BrepHandleArgs{Handle: handle}, &r)
}

// List enumerates the live transient handles.
func (t TransientBRep) List() (wire.BrepListResult, error) {
	var r wire.BrepListResult
	return r, t.c.call(wire.MethodBrepList, struct{}{}, &r)
}

// Delete frees a transient body.
func (t TransientBRep) Delete(handle int) error {
	var r wire.OKResult
	return t.c.call(wire.MethodBrepDelete, wire.BrepHandleArgs{Handle: handle}, &r)
}
