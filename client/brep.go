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
//
// mcp:tool brep_create_primitive
// mcp:summary Creates a solid block/cylinderCone/sphere/torus.
func (t TransientBRep) CreatePrimitive(args wire.CreatePrimitiveArgs) (wire.BrepHandleResult, error) {
	return call[wire.BrepHandleResult](t.c, wire.MethodBrepCreatePrimitive, args)
}

// DoBoolean combines the blank (modified in place) with the tool.
//
// mcp:tool brep_boolean
// mcp:summary Combines the blank (modified in place) with the tool.
func (t TransientBRep) DoBoolean(blankHandle int, tool wire.BrepBodyRef, op types.BooleanType) (wire.BrepHandleResult, error) {
	return call[wire.BrepHandleResult](t.c, wire.MethodBrepBoolean, wire.BrepBooleanArgs{
		BlankHandle: blankHandle, Tool: tool, Operation: op.String(),
	})
}

// Transform maps the body by a 4×4 row-major rigid/similarity matrix.
//
// mcp:tool brep_transform
// mcp:summary Maps the body by a 4×4 row-major rigid/similarity matrix.
func (t TransientBRep) Transform(handle int, matrix []float64) (wire.BrepHandleResult, error) {
	return call[wire.BrepHandleResult](t.c, wire.MethodBrepTransform, wire.BrepTransformArgs{Handle: handle, Matrix: matrix})
}

// Copy clones a transient or document body into a new transient body.
//
// mcp:tool brep_copy
// mcp:summary Clones a transient or document body into a new transient body.
func (t TransientBRep) Copy(source wire.BrepBodyRef) (wire.BrepHandleResult, error) {
	return call[wire.BrepHandleResult](t.c, wire.MethodBrepCopy, wire.BrepCopyArgs{Source: source})
}

// CreateIntersectionWithPlane sections a body with a plane; the section
// curves come back as wires on a new transient body.
//
// mcp:tool brep_section_with_plane
// mcp:summary Sections a body with a plane; the section curves come back as wires on a new transient body.
func (t TransientBRep) CreateIntersectionWithPlane(source wire.BrepBodyRef, planeOrigin, planeNormal []float64) (wire.BrepWiresResult, error) {
	return call[wire.BrepWiresResult](t.c, wire.MethodBrepSectionWithPlane, wire.BrepSectionArgs{
		Source: source, PlaneOrigin: planeOrigin, PlaneNormal: planeNormal,
	})
}

// DeleteFaces removes the named faces (or with keepInstead all others)
// without healing.
//
// mcp:tool brep_delete_faces
// mcp:summary Removes the named faces (or with keepInstead all others) without healing.
func (t TransientBRep) DeleteFaces(handle int, faceKeys []string, keepInstead bool) (wire.BrepHandleResult, error) {
	return call[wire.BrepHandleResult](t.c, wire.MethodBrepDeleteFaces, wire.BrepDeleteFacesArgs{
		Handle: handle, FaceKeys: faceKeys, KeepInstead: keepInstead,
	})
}

// CreateSilhouetteCurve traces one face's silhouette from a view direction.
//
// mcp:tool brep_silhouette
// mcp:summary Traces one face's silhouette from a view direction.
func (t TransientBRep) CreateSilhouetteCurve(args wire.BrepSilhouetteArgs) (wire.BrepWiresResult, error) {
	return call[wire.BrepWiresResult](t.c, wire.MethodBrepSilhouette, args)
}

// CreateRuledSurface builds the ruled surface between two wire sections.
//
// mcp:tool brep_ruled_surface
// mcp:summary Builds the ruled surface between two wire sections.
func (t TransientBRep) CreateRuledSurface(sectionOne, sectionTwo wire.BrepWireRef) (wire.BrepHandleResult, error) {
	return call[wire.BrepHandleResult](t.c, wire.MethodBrepRuledSurface, wire.BrepRuledSurfaceArgs{
		SectionOne: sectionOne, SectionTwo: sectionTwo,
	})
}

// OffsetFaces offsets the named faces of source by distance along their surface normals, returning a
// transient body of the offset faces to sample (e.g. CAM surfacing tool compensation). Parameters:
// reverse to offset into the solid; tolerance for freeform.
//
// mcp:tool brep_offset_faces
// mcp:summary Offsets the named faces of a body by a distance along their normals; the offset faces come back on a new transient body.
func (t TransientBRep) OffsetFaces(args wire.BrepOffsetFacesArgs) (wire.BrepHandleResult, error) {
	return call[wire.BrepHandleResult](t.c, wire.MethodBrepOffsetFaces, args)
}

// ImprintBodies face-splits two bodies along their intersections without
// removing material.
//
// mcp:tool brep_imprint
// mcp:summary Face-splits two bodies along their intersections without removing material.
func (t TransientBRep) ImprintBodies(args wire.BrepImprintArgs) (wire.BrepImprintResult, error) {
	return call[wire.BrepImprintResult](t.c, wire.MethodBrepImprint, args)
}

// GetIdenticalBodies groups congruent bodies.
//
// mcp:tool brep_identical_bodies
// mcp:summary Groups congruent bodies.
func (t TransientBRep) GetIdenticalBodies(args wire.BrepIdenticalBodiesArgs) (wire.BrepIdenticalBodiesResult, error) {
	return call[wire.BrepIdenticalBodiesResult](t.c, wire.MethodBrepIdenticalBodies, args)
}

// CreateFromDefinition compiles a bottom-up definition graph into a body,
// returning per-definition issues instead when the graph is unsound.
//
// mcp:tool brep_create_from_definition
// mcp:summary Compiles a bottom-up definition graph into a body, returning per-definition issues instead when the graph is unsound.
func (t TransientBRep) CreateFromDefinition(def types.BrepBodyDefinition) (wire.BrepCreateFromDefinitionResult, error) {
	return call[wire.BrepCreateFromDefinitionResult](t.c, wire.MethodBrepCreateFromDefinition, wire.BrepCreateFromDefinitionArgs{Definition: def})
}

// Describe returns a transient body's current stats.
//
// mcp:tool brep_describe
// mcp:summary Returns a transient body's current stats.
func (t TransientBRep) Describe(handle int) (wire.BrepHandleResult, error) {
	return call[wire.BrepHandleResult](t.c, wire.MethodBrepDescribe, wire.BrepHandleArgs{Handle: handle})
}

// List enumerates the live transient handles.
//
// mcp:tool brep_list
// mcp:summary Enumerates the live transient handles.
func (t TransientBRep) List() (wire.BrepListResult, error) {
	return call[wire.BrepListResult](t.c, wire.MethodBrepList, struct{}{})
}

// Delete frees a transient body.
//
// mcp:tool brep_delete
// mcp:summary Frees a transient body.
func (t TransientBRep) Delete(handle int) error {
	var r wire.OKResult
	return t.c.call(wire.MethodBrepDelete, wire.BrepHandleArgs{Handle: handle}, &r)
}
