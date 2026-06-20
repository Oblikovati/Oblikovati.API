// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Attached point clouds (M17-F06, Oblikovati/Oblikovati#645): laser-scan / photogrammetry files
// referenced by the active part as transformable display objects modeled against — not B-rep
// geometry. A cloud has a placement transform + uniform scale into model space, a display point
// budget, and visibility. These methods operate on the active part's cloud collection, keyed by
// the cloud's unique name.

// PointCloudInfo is one attached cloud's state. Transform is the cloud→model placement; Scale is
// the uniform cloud→model factor. TotalPointCount is the scan's size; DisplayedPointCount is how
// many render after MaximumPointCount is applied (0 = unbounded).
type PointCloudInfo struct {
	Name                string       `json:"name"`
	Source              string       `json:"source,omitempty"`
	Visible             bool         `json:"visible"`
	Scale               float64      `json:"scale"`
	Transform           types.Matrix `json:"transform"`
	TotalPointCount     int          `json:"totalPointCount"`
	DisplayedPointCount int          `json:"displayedPointCount"`
	MaximumPointCount   int          `json:"maximumPointCount,omitempty"`
}

// AttachPointCloudArgs is the request of [MethodPointCloudsAttach]: read the scan file at
// FullFileName and attach it to the active part. Name is the unique cloud name; when empty the
// host mints one (Cloud1, Cloud2, …).
type AttachPointCloudArgs struct {
	Name         string `json:"name,omitempty"`
	FullFileName string `json:"fullFileName"`
}

// ListPointCloudsResult is the response of [MethodPointCloudsList].
type ListPointCloudsResult struct {
	PointClouds []PointCloudInfo `json:"pointClouds"`
}

// PointCloudNameArgs names a single cloud — the request of [MethodPointCloudsGet] and
// [MethodPointCloudsDelete].
type PointCloudNameArgs struct {
	Name string `json:"name"`
}

// DeletePointCloudResult is the response of [MethodPointCloudsDelete].
type DeletePointCloudResult struct {
	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`
}

// SetPointCloudVisibleArgs is the request of [MethodPointCloudsSetVisible].
type SetPointCloudVisibleArgs struct {
	Name    string `json:"name"`
	Visible bool   `json:"visible"`
}

// SetPointCloudTransformArgs is the request of [MethodPointCloudsSetTransform]: the new
// cloud→model placement.
type SetPointCloudTransformArgs struct {
	Name      string       `json:"name"`
	Transform types.Matrix `json:"transform"`
}

// SetPointCloudScaleArgs is the request of [MethodPointCloudsSetScale]: the new uniform
// cloud→model scale, which must be positive.
type SetPointCloudScaleArgs struct {
	Name  string  `json:"name"`
	Scale float64 `json:"scale"`
}

// SetPointCloudDensityArgs is the request of [MethodPointCloudsSetDensity]: the display budget
// (MaximumPointCount); 0 shows every point.
type SetPointCloudDensityArgs struct {
	Name              string `json:"name"`
	MaximumPointCount int    `json:"maximumPointCount"`
}

// PointCloudSpaceArgs is the request of [MethodPointCloudsToModelSpace] /
// [MethodPointCloudsFromModelSpace]: a point to map between cloud and model space.
type PointCloudSpaceArgs struct {
	Name  string      `json:"name"`
	Point types.Point `json:"point"`
}

// PointCloudSpaceResult is the response of the space-conversion methods: the mapped point. OK is
// false from fromModelSpace when the cloud's placement is non-invertible.
type PointCloudSpaceResult struct {
	Point types.Point `json:"point"`
	OK    bool        `json:"ok"`
}

// Crop volumes (M17-F06, #645): a model-space box on a cloud that, while active, limits the
// cloud's display to points inside it. A cloud's active crops union; with no active crop every
// point shows. Crops are addressed by the owning cloud's name plus the crop's name.

// PointCloudCropInfo is one crop volume: the owning cloud, the crop name, whether it is active,
// and its model-space box corners.
type PointCloudCropInfo struct {
	Cloud  string      `json:"cloud"`
	Crop   string      `json:"crop"`
	Active bool        `json:"active"`
	Min    types.Point `json:"min"`
	Max    types.Point `json:"max"`
}

// AddPointCloudCropArgs is the request of [MethodPointCloudsAddCrop]: add an active crop over the
// model-space box [Min, Max] on the named cloud. The host mints the crop name.
type AddPointCloudCropArgs struct {
	Cloud string      `json:"cloud"`
	Min   types.Point `json:"min"`
	Max   types.Point `json:"max"`
}

// ListPointCloudCropsArgs is the request of [MethodPointCloudsListCrops]: the owning cloud's name.
type ListPointCloudCropsArgs struct {
	Cloud string `json:"cloud"`
}

// ListPointCloudCropsResult is the response of [MethodPointCloudsListCrops].
type ListPointCloudCropsResult struct {
	Crops []PointCloudCropInfo `json:"crops"`
}

// PointCloudCropArgs names one crop on a cloud — the request of [MethodPointCloudsDeleteCrop].
type PointCloudCropArgs struct {
	Cloud string `json:"cloud"`
	Crop  string `json:"crop"`
}

// DeletePointCloudCropResult is the response of [MethodPointCloudsDeleteCrop].
type DeletePointCloudCropResult struct {
	Crop    string `json:"crop"`
	Deleted bool   `json:"deleted"`
}

// SetPointCloudCropActiveArgs is the request of [MethodPointCloudsSetCropActive]: toggle whether a
// named crop limits the cloud's display.
type SetPointCloudCropActiveArgs struct {
	Cloud  string `json:"cloud"`
	Crop   string `json:"crop"`
	Active bool   `json:"active"`
}

// FitPointCloudPlaneArgs is the request of [MethodPointCloudsFitPlane]: fit a work plane to the
// named cloud's currently displayed points — those in model space passing the cloud's active crops
// — so cropping to a planar region first selects what the plane is fitted to.
type FitPointCloudPlaneArgs struct {
	Cloud string `json:"cloud"`
}

// FitPointCloudPlaneResult is the response of [MethodPointCloudsFitPlane]: the created work plane's
// browser name, and the fitted plane's Origin (the points' centroid) and unit Normal (the
// least-variance direction). Normal is a direction, not a position.
type FitPointCloudPlaneResult struct {
	WorkPlane string      `json:"workPlane"`
	Origin    types.Point `json:"origin"`
	Normal    types.Point `json:"normal"`
}
