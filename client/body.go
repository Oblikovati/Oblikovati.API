// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Body is the solid-body topology and query method group (M07-F03/F06/F07,
// Oblikovati/Oblikovati#293/#629/#630).
type Body struct{ c *Client }

// Body returns the body method group.
func (c *Client) Body() Body { return Body{c} }

// List enumerates the active part's bodies.
//
// mcp:tool body_list
// mcp:summary Enumerates the active part's bodies.
func (b Body) List() (wire.BodyListResult, error) {
	var r wire.BodyListResult
	return r, b.c.call(wire.MethodBodyList, struct{}{}, &r)
}

// Shells lists one body's face shells (the outer skin and any cavity skins).
//
// mcp:tool body_shells
// mcp:summary Lists one body's face shells (the outer skin and any cavity skins).
func (b Body) Shells(bodyIndex int) (wire.BodyShellsResult, error) {
	var r wire.BodyShellsResult
	return r, b.c.call(wire.MethodBodyShells, wire.BodyIndexArgs{BodyIndex: bodyIndex}, &r)
}

// Wires lists one body's wires (face-less edge chains).
//
// mcp:tool body_wires
// mcp:summary Lists one body's wires (face-less edge chains).
func (b Body) Wires(bodyIndex int) (wire.BodyWiresResult, error) {
	var r wire.BodyWiresResult
	return r, b.c.call(wire.MethodBodyWires, wire.BodyIndexArgs{BodyIndex: bodyIndex}, &r)
}

// OffsetPlanarWire offsets a planar wire by distance in the plane with the
// given normal, closing gap corners per closure. Set args.Handle to offset a
// transient body's wire (a section/silhouette result).
//
// mcp:tool wire_offset_planar
// mcp:summary Offsets a planar wire by distance in the plane with the given normal, closing gap corners per closure.
func (b Body) OffsetPlanarWire(args wire.OffsetPlanarWireArgs, closure types.OffsetCornerClosureType) (wire.OffsetPlanarWireResult, error) {
	args.CornerClosure = closure.String()
	var r wire.OffsetPlanarWireResult
	return r, b.c.call(wire.MethodWireOffsetPlanar, args, &r)
}

// LocateUsingPoint finds the topology entity nearest the point within the
// proximity tolerance (kind empty = any of vertex/edge/face).
//
// mcp:tool body_locate_using_point
// mcp:summary Finds the topology entity nearest the point within the proximity tolerance (kind empty = any of vertex/edge/face).
func (b Body) LocateUsingPoint(bodyIndex int, point []float64, entityKind string, proximityTolerance float64) (wire.LocateUsingPointResult, error) {
	var r wire.LocateUsingPointResult
	return r, b.c.call(wire.MethodBodyLocateUsingPoint, wire.LocateUsingPointArgs{
		BodyIndex: bodyIndex, Point: point, EntityKind: entityKind, ProximityTolerance: proximityTolerance,
	}, &r)
}

// FindUsingRay fires a pick ray into the body, returning hits nearest first.
//
// mcp:tool body_find_using_ray
// mcp:summary Fires a pick ray into the body, returning hits nearest first.
func (b Body) FindUsingRay(args wire.FindUsingRayArgs) (wire.FindUsingRayResult, error) {
	var r wire.FindUsingRayResult
	return r, b.c.call(wire.MethodBodyFindUsingRay, args, &r)
}

// IsPointInside classifies a point against the body's material (or one
// shell's bounded region).
//
// mcp:tool body_is_point_inside
// mcp:summary Classifies a point against the body's material (or one shell's bounded region).
func (b Body) IsPointInside(args wire.IsPointInsideArgs) (types.Containment, error) {
	var r wire.IsPointInsideResult
	if err := b.c.call(wire.MethodBodyIsPointInside, args, &r); err != nil {
		return types.UnknownContainment, err
	}
	if c, ok := types.ParseContainment(r.Containment); ok {
		return c, nil
	}
	return types.UnknownContainment, nil
}

// ConvexityEdges returns the body's edges of the requested dihedral class.
//
// mcp:tool body_convexity_edges
// mcp:summary Returns the body's edges of the requested dihedral class.
func (b Body) ConvexityEdges(bodyIndex int, collection types.EdgeCollectionKind) (wire.ConvexityEdgesResult, error) {
	var r wire.ConvexityEdgesResult
	return r, b.c.call(wire.MethodBodyConvexityEdges, wire.ConvexityEdgesArgs{
		BodyIndex: bodyIndex, Collection: collection.String(),
	}, &r)
}

// Validate checks the body (checkLevel 1 = topology, 2 = + self-intersection)
// and reports any offending entities.
//
// mcp:tool body_validate
// mcp:summary Checks the body (checkLevel 1 = topology, 2 = + self-intersection) and reports any offending entities.
func (b Body) Validate(bodyIndex, checkLevel int) (wire.ValidateBodyResult, error) {
	var r wire.ValidateBodyResult
	return r, b.c.call(wire.MethodBodyValidate, wire.ValidateBodyArgs{BodyIndex: bodyIndex, CheckLevel: checkLevel}, &r)
}

// RangeBox returns the body's range box (topology, precise or oriented).
//
// mcp:tool body_range_box
// mcp:summary Returns the body's range box (topology, precise or oriented).
func (b Body) RangeBox(args wire.BodyRangeBoxArgs) (wire.BodyRangeBoxResult, error) {
	var r wire.BodyRangeBoxResult
	return r, b.c.call(wire.MethodBodyRangeBox, args, &r)
}

// BindTransientKey resolves a session transient key back to its entity.
//
// mcp:tool body_bind_transient_key
// mcp:summary Resolves a session transient key back to its entity.
func (b Body) BindTransientKey(bodyIndex int, transientKey uint64) (wire.BindTransientKeyResult, error) {
	var r wire.BindTransientKeyResult
	return r, b.c.call(wire.MethodBodyBindTransientKey, wire.BindTransientKeyArgs{
		BodyIndex: bodyIndex, TransientKey: transientKey,
	}, &r)
}

// CalculateFacets facets the body at the tolerance (cached under it).
//
// mcp:tool body_calculate_facets
// mcp:summary Facets the body at the tolerance (cached under it).
func (b Body) CalculateFacets(args wire.CalculateFacetsArgs) (wire.FacetSetResult, error) {
	var r wire.FacetSetResult
	return r, b.c.call(wire.MethodBodyCalculateFacets, args, &r)
}

// ExistingFacets retrieves a previously calculated facet set without
// re-faceting (errors when no set exists at the tolerance).
//
// mcp:tool body_existing_facets
// mcp:summary Retrieves a previously calculated facet set without re-faceting (errors when no set exists at the tolerance).
func (b Body) ExistingFacets(bodyIndex int, tolerance float64) (wire.FacetSetResult, error) {
	var r wire.FacetSetResult
	return r, b.c.call(wire.MethodBodyExistingFacets, wire.CalculateFacetsArgs{BodyIndex: bodyIndex, Tolerance: tolerance}, &r)
}

// FacetTolerances lists the tolerances facet sets are cached at, ascending.
//
// mcp:tool body_facet_tolerances
// mcp:summary Lists the tolerances facet sets are cached at, ascending.
func (b Body) FacetTolerances(bodyIndex int) (wire.FacetTolerancesResult, error) {
	var r wire.FacetTolerancesResult
	return r, b.c.call(wire.MethodBodyFacetTolerances, wire.BodyIndexArgs{BodyIndex: bodyIndex}, &r)
}

// CalculateStrokes samples the body's edges at the tolerance (cached).
//
// mcp:tool body_calculate_strokes
// mcp:summary Samples the body's edges at the tolerance (cached).
func (b Body) CalculateStrokes(bodyIndex int, tolerance float64) (wire.StrokeSetResult, error) {
	var r wire.StrokeSetResult
	return r, b.c.call(wire.MethodBodyCalculateStrokes, wire.CalculateStrokesArgs{BodyIndex: bodyIndex, Tolerance: tolerance}, &r)
}

// ExistingStrokes retrieves a previously calculated stroke set.
//
// mcp:tool body_existing_strokes
// mcp:summary Retrieves a previously calculated stroke set.
func (b Body) ExistingStrokes(bodyIndex int, tolerance float64) (wire.StrokeSetResult, error) {
	var r wire.StrokeSetResult
	return r, b.c.call(wire.MethodBodyExistingStrokes, wire.CalculateStrokesArgs{BodyIndex: bodyIndex, Tolerance: tolerance}, &r)
}

// StrokeTolerances lists the tolerances stroke sets are cached at, ascending.
//
// mcp:tool body_stroke_tolerances
// mcp:summary Lists the tolerances stroke sets are cached at, ascending.
func (b Body) StrokeTolerances(bodyIndex int) (wire.FacetTolerancesResult, error) {
	var r wire.FacetTolerancesResult
	return r, b.c.call(wire.MethodBodyStrokeTolerances, wire.BodyIndexArgs{BodyIndex: bodyIndex}, &r)
}

// FaceCalculateFacets facets one face (riding the body's tolerance cache).
//
// mcp:tool face_calculate_facets
// mcp:summary Facets one face (riding the body's tolerance cache).
func (b Body) FaceCalculateFacets(args wire.FaceFacetsArgs) (wire.FacetSetResult, error) {
	var r wire.FacetSetResult
	return r, b.c.call(wire.MethodFaceCalculateFacets, args, &r)
}

// FaceCalculateStrokes samples one face's boundary edges.
//
// mcp:tool face_calculate_strokes
// mcp:summary Samples one face's boundary edges.
func (b Body) FaceCalculateStrokes(args wire.FaceFacetsArgs) (wire.StrokeSetResult, error) {
	var r wire.StrokeSetResult
	return r, b.c.call(wire.MethodFaceCalculateStrokes, args, &r)
}
