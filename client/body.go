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
// mcp:summary Enumerates the active part's bodies (name, solid flag, visibility, face/edge counts).
func (b Body) List() (wire.BodyListResult, error) {
	return call[wire.BodyListResult](b.c, wire.MethodBodyList, struct{}{})
}

// SetVisible shows or hides the body at index (from List), for multi-body workflows (#158).
//
// mcp:tool body_set_visible
// mcp:summary Show or hide one body of the active part by index.
func (b Body) SetVisible(index int, visible bool) (wire.BodyInfoResult, error) {
	return call[wire.BodyInfoResult](b.c, wire.MethodBodySetVisible, wire.BodySetVisibleArgs{BodyIndex: index, Visible: visible})
}

// Rename sets the display name of the body at index (from List); an empty name reverts to the
// "Solid{N}" default. The name is stored per body, survives recompute, and round-trips in the
// document (#1078).
//
// mcp:tool body_rename
// mcp:summary Set the display name of one body of the active part by index (empty reverts to the default).
func (b Body) Rename(index int, name string) (wire.BodyInfoResult, error) {
	return call[wire.BodyInfoResult](b.c, wire.MethodBodyRename, wire.BodyRenameArgs{BodyIndex: index, Name: name})
}

// Delete removes the body at index (from List) from the active part, returning the refreshed
// body list (#1078).
//
// mcp:tool body_delete
// mcp:summary Delete one body of the active part by index, returning the refreshed body list.
func (b Body) Delete(index int) (wire.BodyListResult, error) {
	return call[wire.BodyListResult](b.c, wire.MethodBodyDelete, wire.BodyIndexArgs{BodyIndex: index})
}

// PhysicalProperties returns one body's geometry and mass properties — the per-body counterpart
// of Model.PhysicalProperties, which sums all bodies (#1078). DensityGCm3 0 uses the part's
// material density; accuracy is "low"/"medium"/"high" (empty ⇒ medium).
//
// mcp:tool body_physical_properties
// mcp:summary Geometry and mass properties (volume, area, mass, centroid, inertia) of one body of the active part.
func (b Body) PhysicalProperties(index int, densityGCm3 float64, accuracy string) (wire.MassPropertiesResult, error) {
	args := wire.BodyPhysicalPropertiesArgs{BodyIndex: index, DensityGCm3: densityGCm3, Accuracy: accuracy}
	return call[wire.MassPropertiesResult](b.c, wire.MethodBodyPhysicalProps, args)
}

// Shells lists one body's face shells (the outer skin and any cavity skins).
//
// mcp:tool body_shells
// mcp:summary Lists one body's face shells (the outer skin and any cavity skins).
func (b Body) Shells(bodyIndex int) (wire.BodyShellsResult, error) {
	return call[wire.BodyShellsResult](b.c, wire.MethodBodyShells, wire.BodyIndexArgs{BodyIndex: bodyIndex})
}

// Wires lists one body's wires (face-less edge chains).
//
// mcp:tool body_wires
// mcp:summary Lists one body's wires (face-less edge chains).
func (b Body) Wires(bodyIndex int) (wire.BodyWiresResult, error) {
	return call[wire.BodyWiresResult](b.c, wire.MethodBodyWires, wire.BodyIndexArgs{BodyIndex: bodyIndex})
}

// OffsetPlanarWire offsets a planar wire by distance in the plane with the
// given normal, closing gap corners per closure. Set args.Handle to offset a
// transient body's wire (a section/silhouette result).
//
// mcp:tool wire_offset_planar
// mcp:summary Offsets a planar wire by distance in the plane with the given normal, closing gap corners per closure.
func (b Body) OffsetPlanarWire(args wire.OffsetPlanarWireArgs, closure types.OffsetCornerClosureType) (wire.OffsetPlanarWireResult, error) {
	args.CornerClosure = closure.String()
	return call[wire.OffsetPlanarWireResult](b.c, wire.MethodWireOffsetPlanar, args)
}

// LocateUsingPoint finds the topology entity nearest the point within the
// proximity tolerance (kind empty = any of vertex/edge/face).
//
// mcp:tool body_locate_using_point
// mcp:summary Finds the topology entity nearest the point within the proximity tolerance (kind empty = any of vertex/edge/face).
func (b Body) LocateUsingPoint(bodyIndex int, point []float64, entityKind string, proximityTolerance float64) (wire.LocateUsingPointResult, error) {
	return call[wire.LocateUsingPointResult](b.c, wire.MethodBodyLocateUsingPoint, wire.LocateUsingPointArgs{
		BodyIndex: bodyIndex, Point: point, EntityKind: entityKind, ProximityTolerance: proximityTolerance,
	})
}

// FindUsingRay fires a pick ray into the body, returning hits nearest first.
//
// mcp:tool body_find_using_ray
// mcp:summary Fires a pick ray into the body, returning hits nearest first.
func (b Body) FindUsingRay(args wire.FindUsingRayArgs) (wire.FindUsingRayResult, error) {
	return call[wire.FindUsingRayResult](b.c, wire.MethodBodyFindUsingRay, args)
}

// IsPointInside classifies a point against the body's material (or one
// shell's bounded region).
//
// mcp:tool body_is_point_inside
// mcp:summary Classifies a point against the body's material (or one shell's bounded region).
func (b Body) IsPointInside(args wire.IsPointInsideArgs) (types.Containment, error) {
	r, err := call[wire.IsPointInsideResult](b.c, wire.MethodBodyIsPointInside, args)
	if err != nil {
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
	return call[wire.ConvexityEdgesResult](b.c, wire.MethodBodyConvexityEdges, wire.ConvexityEdgesArgs{
		BodyIndex: bodyIndex, Collection: collection.String(),
	})
}

// MinimumDistance returns the closest approach between the body and a transient
// probe polyline (e.g. a CAM travel path), optionally widened by Radius (the
// tool cross-section). It is the out-of-process projection of Inventor's
// MeasureTools.GetMinimumDistance for a transient operand; points, Radius and
// the result are all in database units (cm). The result is 0 when the probe
// (after Radius) touches or enters the body's material.
//
//	d, _ := c.Body().MinimumDistance(wire.MinimumDistanceArgs{
//		BodyIndex: 0, Points: []float64{0, 0, 5, 4, 0, 5}, Radius: 0.3,
//	})
//
// mcp:tool body_minimum_distance
// mcp:summary Minimum distance between the body and a transient probe polyline (flat x,y,z list in cm); radius widens the probe into a swept-tool cylinder; returns the distance (cm), 0 when it touches or enters the body — the out-of-process projection of MeasureTools.GetMinimumDistance for a transient operand.
func (b Body) MinimumDistance(args wire.MinimumDistanceArgs) (wire.MinimumDistanceResult, error) {
	return call[wire.MinimumDistanceResult](b.c, wire.MethodBodyMinimumDistance, args)
}

// Validate checks the body (checkLevel 1 = topology, 2 = + self-intersection)
// and reports any offending entities.
//
// mcp:tool body_validate
// mcp:summary Checks the body (checkLevel 1 = topology, 2 = + self-intersection) and reports any offending entities.
func (b Body) Validate(bodyIndex, checkLevel int) (wire.ValidateBodyResult, error) {
	return call[wire.ValidateBodyResult](b.c, wire.MethodBodyValidate, wire.ValidateBodyArgs{BodyIndex: bodyIndex, CheckLevel: checkLevel})
}

// RangeBox returns the body's range box (topology, precise or oriented).
//
// mcp:tool body_range_box
// mcp:summary Returns the body's range box (topology, precise or oriented).
func (b Body) RangeBox(args wire.BodyRangeBoxArgs) (wire.BodyRangeBoxResult, error) {
	return call[wire.BodyRangeBoxResult](b.c, wire.MethodBodyRangeBox, args)
}

// BindTransientKey resolves a session transient key back to its entity.
//
// mcp:tool body_bind_transient_key
// mcp:summary Resolves a session transient key back to its entity.
func (b Body) BindTransientKey(bodyIndex int, transientKey uint64) (wire.BindTransientKeyResult, error) {
	return call[wire.BindTransientKeyResult](b.c, wire.MethodBodyBindTransientKey, wire.BindTransientKeyArgs{
		BodyIndex: bodyIndex, TransientKey: transientKey,
	})
}

// CalculateFacets facets the body at the tolerance (cached under it).
//
// mcp:tool body_calculate_facets
// mcp:summary Facets the body at the tolerance (cached under it).
func (b Body) CalculateFacets(args wire.CalculateFacetsArgs) (wire.FacetSetResult, error) {
	return call[wire.FacetSetResult](b.c, wire.MethodBodyCalculateFacets, args)
}

// ExistingFacets retrieves a previously calculated facet set without
// re-faceting (errors when no set exists at the tolerance).
//
// mcp:tool body_existing_facets
// mcp:summary Retrieves a previously calculated facet set without re-faceting (errors when no set exists at the tolerance).
func (b Body) ExistingFacets(bodyIndex int, tolerance float64) (wire.FacetSetResult, error) {
	return call[wire.FacetSetResult](b.c, wire.MethodBodyExistingFacets, wire.CalculateFacetsArgs{BodyIndex: bodyIndex, Tolerance: tolerance})
}

// FacetTolerances lists the tolerances facet sets are cached at, ascending.
//
// mcp:tool body_facet_tolerances
// mcp:summary Lists the tolerances facet sets are cached at, ascending.
func (b Body) FacetTolerances(bodyIndex int) (wire.FacetTolerancesResult, error) {
	return call[wire.FacetTolerancesResult](b.c, wire.MethodBodyFacetTolerances, wire.BodyIndexArgs{BodyIndex: bodyIndex})
}

// CalculateStrokes samples the body's edges at the tolerance (cached).
//
// mcp:tool body_calculate_strokes
// mcp:summary Samples the body's edges at the tolerance (cached).
func (b Body) CalculateStrokes(bodyIndex int, tolerance float64) (wire.StrokeSetResult, error) {
	return call[wire.StrokeSetResult](b.c, wire.MethodBodyCalculateStrokes, wire.CalculateStrokesArgs{BodyIndex: bodyIndex, Tolerance: tolerance})
}

// ExistingStrokes retrieves a previously calculated stroke set.
//
// mcp:tool body_existing_strokes
// mcp:summary Retrieves a previously calculated stroke set.
func (b Body) ExistingStrokes(bodyIndex int, tolerance float64) (wire.StrokeSetResult, error) {
	return call[wire.StrokeSetResult](b.c, wire.MethodBodyExistingStrokes, wire.CalculateStrokesArgs{BodyIndex: bodyIndex, Tolerance: tolerance})
}

// StrokeTolerances lists the tolerances stroke sets are cached at, ascending.
//
// mcp:tool body_stroke_tolerances
// mcp:summary Lists the tolerances stroke sets are cached at, ascending.
func (b Body) StrokeTolerances(bodyIndex int) (wire.FacetTolerancesResult, error) {
	return call[wire.FacetTolerancesResult](b.c, wire.MethodBodyStrokeTolerances, wire.BodyIndexArgs{BodyIndex: bodyIndex})
}

// FaceCalculateFacets facets one face (riding the body's tolerance cache).
//
// mcp:tool face_calculate_facets
// mcp:summary Facets one face (riding the body's tolerance cache).
func (b Body) FaceCalculateFacets(args wire.FaceFacetsArgs) (wire.FacetSetResult, error) {
	return call[wire.FacetSetResult](b.c, wire.MethodFaceCalculateFacets, args)
}

// FaceCalculateStrokes samples one face's boundary edges.
//
// mcp:tool face_calculate_strokes
// mcp:summary Samples one face's boundary edges.
func (b Body) FaceCalculateStrokes(args wire.FaceFacetsArgs) (wire.StrokeSetResult, error) {
	return call[wire.StrokeSetResult](b.c, wire.MethodFaceCalculateStrokes, args)
}

// FaceEvaluate batch-evaluates one face's surface (point/normal/tangents at given (u,v)
// parameters, or the projection of given points onto the face), addressed by reference key.
// The mode selects the query (see the wire.FaceEval* constants); inputs and the populated
// result arrays are flat and parallel. Lengths are database units (cm); normals are unit
// vectors. It is the out-of-process query a surface-following or point-projection toolpath
// uses to sample a face densely in one call.
//
// mcp:tool body_face_evaluate
// mcp:summary Batch-evaluates a face's surface (point/normal/tangents by param, or point projection).
func (b Body) FaceEvaluate(args wire.FaceEvaluateArgs) (wire.FaceEvaluateResult, error) {
	return call[wire.FaceEvaluateResult](b.c, wire.MethodBodyFaceEvaluate, args)
}
