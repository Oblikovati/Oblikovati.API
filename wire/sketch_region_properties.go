// SPDX-License-Identifier: Apache-2.0

package wire

// Region properties (M06-F08, Oblikovati/Oblikovati#623): the full section
// property set of a closed profile region — beyond the area already carried by
// [ProfileInfo] — computed by integrating over the region's loops.

// RegionPropertiesArgs is the request of [MethodSketchRegionProperties] and
// [MethodSketch3DRegionProperties]: which sketch, which profile (the index
// from [MethodSketchProfiles] / [MethodSketch3DProfiles]; the profile must be
// closed — and planar, for the 3D method), and the desired computational
// accuracy ([oblikovati.org/api/types.Accuracy] wire spelling; empty ⇒ "high").
type RegionPropertiesArgs struct {
	SketchIndex  int    `json:"sketchIndex"`
	ProfileIndex int    `json:"profileIndex"`
	Accuracy     string `json:"accuracy,omitempty"`
}

// RegionPropertiesResult is the response of [MethodSketchRegionProperties] and
// [MethodSketch3DRegionProperties]. All values are in database units (cm) on
// the profile's plane, about the sketch origin unless stated otherwise:
//   - Area (cm², holes subtracted) and Perimeter (cm, holes' rims included);
//   - Centroid [x,y] (cm, in sketch-plane coordinates);
//   - MomentsOfInertia [Ixx, Iyy, Ixy] (cm⁴) about the centroid;
//   - PrincipalMoments [I1, I2] (cm⁴) and RotationAngle (radians CCW from the
//     sketch X axis to the first principal axis);
//   - PrincipalAxes the two unit axes [[x1,y1],[x2,y2]];
//   - Accuracy the wire spelling actually used.
type RegionPropertiesResult struct {
	Area             float64     `json:"area"`
	Perimeter        float64     `json:"perimeter"`
	Centroid         []float64   `json:"centroid"`
	MomentsOfInertia []float64   `json:"momentsOfInertia"`
	PrincipalMoments []float64   `json:"principalMoments"`
	RotationAngle    float64     `json:"rotationAngle"`
	PrincipalAxes    [][]float64 `json:"principalAxes"`
	Accuracy         string      `json:"accuracy"`
}
