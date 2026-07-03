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
	return call[wire.PointCloudInfo](p.c, wire.MethodPointCloudsAttach, wire.AttachPointCloudArgs{Name: name, FullFileName: fullFileName})
}

// List enumerates the active part's attached clouds.
//
// mcp:tool point_clouds_list
// mcp:summary Enumerate the active part's attached point clouds (name, point counts, scale, visibility).
func (p PointClouds) List() (wire.ListPointCloudsResult, error) {
	return call[wire.ListPointCloudsResult](p.c, wire.MethodPointCloudsList, struct{}{})
}

// Get returns one attached cloud's state by name.
//
// mcp:tool point_clouds_get
// mcp:summary Return one attached point cloud's state (placement, scale, point counts) by name.
func (p PointClouds) Get(name string) (wire.PointCloudInfo, error) {
	return call[wire.PointCloudInfo](p.c, wire.MethodPointCloudsGet, wire.PointCloudNameArgs{Name: name})
}

// Delete removes the named cloud from the active part.
//
// mcp:tool point_clouds_delete
// mcp:summary Remove an attached point cloud from the active part by name.
func (p PointClouds) Delete(name string) (wire.DeletePointCloudResult, error) {
	return call[wire.DeletePointCloudResult](p.c, wire.MethodPointCloudsDelete, wire.PointCloudNameArgs{Name: name})
}

// SetVisible shows or hides the named cloud.
//
// mcp:tool point_clouds_set_visible
// mcp:summary Show or hide an attached point cloud by name.
func (p PointClouds) SetVisible(name string, visible bool) (wire.PointCloudInfo, error) {
	return call[wire.PointCloudInfo](p.c, wire.MethodPointCloudsSetVisible, wire.SetPointCloudVisibleArgs{Name: name, Visible: visible})
}

// SetTransform sets the named cloud's cloud→model placement.
//
// mcp:tool point_clouds_set_transform
// mcp:summary Set an attached point cloud's placement transform by name.
func (p PointClouds) SetTransform(name string, transform types.Matrix) (wire.PointCloudInfo, error) {
	return call[wire.PointCloudInfo](p.c, wire.MethodPointCloudsSetTransform, wire.SetPointCloudTransformArgs{Name: name, Transform: transform})
}

// SetScale sets the named cloud's uniform cloud→model scale (must be positive).
//
// mcp:tool point_clouds_set_scale
// mcp:summary Set an attached point cloud's uniform scale by name (must be positive).
func (p PointClouds) SetScale(name string, scale float64) (wire.PointCloudInfo, error) {
	return call[wire.PointCloudInfo](p.c, wire.MethodPointCloudsSetScale, wire.SetPointCloudScaleArgs{Name: name, Scale: scale})
}

// SetDensity sets the named cloud's display budget (MaximumPointCount); 0 shows every point.
//
// mcp:tool point_clouds_set_density
// mcp:summary Set an attached point cloud's display point budget by name (0 = show all).
func (p PointClouds) SetDensity(name string, maximumPointCount int) (wire.PointCloudInfo, error) {
	return call[wire.PointCloudInfo](p.c, wire.MethodPointCloudsSetDensity, wire.SetPointCloudDensityArgs{Name: name, MaximumPointCount: maximumPointCount})
}

// ToModelSpace maps a point from the named cloud's local space into model space.
//
// mcp:tool point_clouds_to_model_space
// mcp:summary Map a point from a cloud's local space into model space.
func (p PointClouds) ToModelSpace(name string, point types.Point) (wire.PointCloudSpaceResult, error) {
	return call[wire.PointCloudSpaceResult](p.c, wire.MethodPointCloudsToModelSpace, wire.PointCloudSpaceArgs{Name: name, Point: point})
}

// FromModelSpace maps a model-space point into the named cloud's local space (OK is false when
// the placement is non-invertible).
//
// mcp:tool point_clouds_from_model_space
// mcp:summary Map a model-space point into a cloud's local space.
func (p PointClouds) FromModelSpace(name string, point types.Point) (wire.PointCloudSpaceResult, error) {
	return call[wire.PointCloudSpaceResult](p.c, wire.MethodPointCloudsFromModelSpace, wire.PointCloudSpaceArgs{Name: name, Point: point})
}

// AddCrop adds an active crop over the model-space box [min, max] on the named cloud, limiting its
// display to points inside (the host mints the crop name).
//
// mcp:tool point_clouds_add_crop
// mcp:summary Add a crop volume that limits a point cloud's display to a model-space box.
func (p PointClouds) AddCrop(cloud string, min, max types.Point) (wire.PointCloudCropInfo, error) {
	return call[wire.PointCloudCropInfo](p.c, wire.MethodPointCloudsAddCrop, wire.AddPointCloudCropArgs{Cloud: cloud, Min: min, Max: max})
}

// ListCrops enumerates the named cloud's crop volumes.
//
// mcp:tool point_clouds_list_crops
// mcp:summary Enumerate a point cloud's crop volumes (name, active, box).
func (p PointClouds) ListCrops(cloud string) (wire.ListPointCloudCropsResult, error) {
	return call[wire.ListPointCloudCropsResult](p.c, wire.MethodPointCloudsListCrops, wire.ListPointCloudCropsArgs{Cloud: cloud})
}

// DeleteCrop removes a named crop from a cloud.
//
// mcp:tool point_clouds_delete_crop
// mcp:summary Remove a crop volume from a point cloud by name.
func (p PointClouds) DeleteCrop(cloud, crop string) (wire.DeletePointCloudCropResult, error) {
	return call[wire.DeletePointCloudCropResult](p.c, wire.MethodPointCloudsDeleteCrop, wire.PointCloudCropArgs{Cloud: cloud, Crop: crop})
}

// SetCropActive toggles whether a named crop limits the cloud's display.
//
// mcp:tool point_clouds_set_crop_active
// mcp:summary Toggle whether a point cloud crop volume limits display.
func (p PointClouds) SetCropActive(cloud, crop string, active bool) (wire.PointCloudCropInfo, error) {
	return call[wire.PointCloudCropInfo](p.c, wire.MethodPointCloudsSetCropActive, wire.SetPointCloudCropActiveArgs{Cloud: cloud, Crop: crop, Active: active})
}

// FitPlane fits a least-squares work plane to the named cloud's displayed points (those passing its
// active crops) and returns the new work plane's name with the fitted origin (centroid) and unit
// normal. Crop to a planar region first to control what the plane is fitted to.
//
// mcp:tool point_clouds_fit_plane
// mcp:summary Fit a work plane to a point cloud's displayed points (least-squares).
func (p PointClouds) FitPlane(cloud string) (wire.FitPointCloudPlaneResult, error) {
	return call[wire.FitPointCloudPlaneResult](p.c, wire.MethodPointCloudsFitPlane, wire.FitPointCloudPlaneArgs{Cloud: cloud})
}

// NearestPoint snaps the model-space point onto the named cloud, returning its scan point nearest
// the query and the distance to it (Found is false only for an empty cloud). Compose with
// WorkPoints.Create to anchor a datum on the as-built scan data.
//
// mcp:tool point_clouds_nearest_point
// mcp:summary Find a point cloud's scan point nearest a model-space query (snap).
func (p PointClouds) NearestPoint(cloud string, point types.Point) (wire.NearestPointResult, error) {
	return call[wire.NearestPointResult](p.c, wire.MethodPointCloudsNearestPoint, wire.NearestPointArgs{Cloud: cloud, Point: point})
}
