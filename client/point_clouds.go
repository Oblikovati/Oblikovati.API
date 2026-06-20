// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// PointClouds is the attached-scan method group (M17-F06, Oblikovati/Oblikovati#645): attach,
// query, place, and budget laser-scan / photogrammetry clouds on the active part.
type PointClouds struct{ c *Client }

// PointClouds returns the point-cloud method group.
func (c *Client) PointClouds() PointClouds { return PointClouds{c} }

// Attach reads the scan file at fullFileName and attaches it to the active part. An empty name
// lets the host mint one (Cloud1, Cloud2, …).
//
// mcp:tool point_clouds_attach
// mcp:summary Attach a laser-scan / photogrammetry file (.xyz/.pts) to the active part as a referenced point cloud.
func (p PointClouds) Attach(name, fullFileName string) (wire.PointCloudInfo, error) {
	var r wire.PointCloudInfo
	return r, p.c.call(wire.MethodPointCloudsAttach, wire.AttachPointCloudArgs{Name: name, FullFileName: fullFileName}, &r)
}

// List enumerates the active part's attached clouds.
//
// mcp:tool point_clouds_list
// mcp:summary Enumerate the active part's attached point clouds (name, point counts, scale, visibility).
func (p PointClouds) List() (wire.ListPointCloudsResult, error) {
	var r wire.ListPointCloudsResult
	return r, p.c.call(wire.MethodPointCloudsList, struct{}{}, &r)
}

// Get returns one attached cloud's state by name.
//
// mcp:tool point_clouds_get
// mcp:summary Return one attached point cloud's state (placement, scale, point counts) by name.
func (p PointClouds) Get(name string) (wire.PointCloudInfo, error) {
	var r wire.PointCloudInfo
	return r, p.c.call(wire.MethodPointCloudsGet, wire.PointCloudNameArgs{Name: name}, &r)
}

// Delete removes the named cloud from the active part.
//
// mcp:tool point_clouds_delete
// mcp:summary Remove an attached point cloud from the active part by name.
func (p PointClouds) Delete(name string) (wire.DeletePointCloudResult, error) {
	var r wire.DeletePointCloudResult
	return r, p.c.call(wire.MethodPointCloudsDelete, wire.PointCloudNameArgs{Name: name}, &r)
}

// SetVisible shows or hides the named cloud.
//
// mcp:tool point_clouds_set_visible
// mcp:summary Show or hide an attached point cloud by name.
func (p PointClouds) SetVisible(name string, visible bool) (wire.PointCloudInfo, error) {
	var r wire.PointCloudInfo
	return r, p.c.call(wire.MethodPointCloudsSetVisible, wire.SetPointCloudVisibleArgs{Name: name, Visible: visible}, &r)
}

// SetTransform sets the named cloud's cloud→model placement.
//
// mcp:tool point_clouds_set_transform
// mcp:summary Set an attached point cloud's placement transform by name.
func (p PointClouds) SetTransform(name string, transform types.Matrix) (wire.PointCloudInfo, error) {
	var r wire.PointCloudInfo
	return r, p.c.call(wire.MethodPointCloudsSetTransform, wire.SetPointCloudTransformArgs{Name: name, Transform: transform}, &r)
}

// SetScale sets the named cloud's uniform cloud→model scale (must be positive).
//
// mcp:tool point_clouds_set_scale
// mcp:summary Set an attached point cloud's uniform scale by name (must be positive).
func (p PointClouds) SetScale(name string, scale float64) (wire.PointCloudInfo, error) {
	var r wire.PointCloudInfo
	return r, p.c.call(wire.MethodPointCloudsSetScale, wire.SetPointCloudScaleArgs{Name: name, Scale: scale}, &r)
}

// SetDensity sets the named cloud's display budget (MaximumPointCount); 0 shows every point.
//
// mcp:tool point_clouds_set_density
// mcp:summary Set an attached point cloud's display point budget by name (0 = show all).
func (p PointClouds) SetDensity(name string, maximumPointCount int) (wire.PointCloudInfo, error) {
	var r wire.PointCloudInfo
	return r, p.c.call(wire.MethodPointCloudsSetDensity, wire.SetPointCloudDensityArgs{Name: name, MaximumPointCount: maximumPointCount}, &r)
}

// ToModelSpace maps a point from the named cloud's local space into model space.
//
// mcp:tool point_clouds_to_model_space
// mcp:summary Map a point from a cloud's local space into model space.
func (p PointClouds) ToModelSpace(name string, point types.Point) (wire.PointCloudSpaceResult, error) {
	var r wire.PointCloudSpaceResult
	return r, p.c.call(wire.MethodPointCloudsToModelSpace, wire.PointCloudSpaceArgs{Name: name, Point: point}, &r)
}

// FromModelSpace maps a model-space point into the named cloud's local space (OK is false when
// the placement is non-invertible).
//
// mcp:tool point_clouds_from_model_space
// mcp:summary Map a model-space point into a cloud's local space.
func (p PointClouds) FromModelSpace(name string, point types.Point) (wire.PointCloudSpaceResult, error) {
	var r wire.PointCloudSpaceResult
	return r, p.c.call(wire.MethodPointCloudsFromModelSpace, wire.PointCloudSpaceArgs{Name: name, Point: point}, &r)
}
