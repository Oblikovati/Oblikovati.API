// SPDX-License-Identifier: Apache-2.0

package wire

// Analysis DTOs (M18-F01 #423): engineering analysis on the model. Mass properties of the active
// part — volume, surface area, centre of mass and mass.

// MassPropertiesArgs is the request of [MethodAnalysisMassProperties]: compute the active part's
// mass properties. DensityGCm3 is the material density in g/cm³ used for the mass (0 ⇒ 1.0, so the
// mass equals the volume in cm³).
type MassPropertiesArgs struct {
	DensityGCm3 float64 `json:"densityGCm3,omitempty"`
}

// MassPropertiesResult is the response of [MethodAnalysisMassProperties]: the part's combined
// geometry properties (over all its solid bodies) and mass. Lengths are millimetres.
type MassPropertiesResult struct {
	VolumeMm3      float64 `json:"volumeMm3"`
	SurfaceAreaMm2 float64 `json:"surfaceAreaMm2"`
	MassG          float64 `json:"massG"`
	DensityGCm3    float64 `json:"densityGCm3"`
	CentroidXMm    float64 `json:"centroidXMm"`
	CentroidYMm    float64 `json:"centroidYMm"`
	CentroidZMm    float64 `json:"centroidZMm"`
}
