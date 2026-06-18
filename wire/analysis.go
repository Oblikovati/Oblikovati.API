// SPDX-License-Identifier: Apache-2.0

package wire

// Analysis DTOs (M18-F01 #423): engineering analysis on the model. Mass properties of the active
// part — volume, surface area, centre of mass and mass.

// MeasureArgs is the request of [MethodAnalysisMeasure]: measure an entity (or pair) of the active
// part's body BodyIndex, identified by reference key(s). Type selects the quantity: "length"
// (edge KeyA), "area" (face KeyA), "distance" (straight line between vertices KeyA and KeyB), or
// "minDistance" (closest approach between the two entities KeyA and KeyB, each a vertex/edge/face).
type MeasureArgs struct {
	BodyIndex int    `json:"bodyIndex,omitempty"`
	Type      string `json:"type"`
	KeyA      string `json:"keyA"`
	KeyB      string `json:"keyB,omitempty"`
}

// MeasureResult is the response of [MethodAnalysisMeasure]: the measured value and its unit
// ("mm" for length/distance, "mm²" for area).
type MeasureResult struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// MassPropertiesArgs is the request of [MethodAnalysisMassProperties]: compute the active part's
// mass properties. DensityGCm3 is the material density in g/cm³ used for the mass (0 ⇒ 1.0, so the
// mass equals the volume in cm³).
type MassPropertiesArgs struct {
	// DensityGCm3 overrides the material density (g/cm³). 0 ⇒ the part's assigned material density,
	// falling back to 1.0 when no material is assigned.
	DensityGCm3 float64 `json:"densityGCm3,omitempty"`
	// Accuracy is the tessellation fidelity ("low"/"medium"/"high"; empty ⇒ medium).
	Accuracy string `json:"accuracy,omitempty"`
}

// MassPropertiesResult is the response of [MethodAnalysisMassProperties]: the part's combined
// geometry properties (over all its solid bodies), mass, and mass moment of inertia about the
// centroid. Lengths are millimetres; inertia is g·mm².
type MassPropertiesResult struct {
	VolumeMm3      float64 `json:"volumeMm3"`
	SurfaceAreaMm2 float64 `json:"surfaceAreaMm2"`
	MassG          float64 `json:"massG"`
	DensityGCm3    float64 `json:"densityGCm3"`
	CentroidXMm    float64 `json:"centroidXMm"`
	CentroidYMm    float64 `json:"centroidYMm"`
	CentroidZMm    float64 `json:"centroidZMm"`
	// Mass moment of inertia about the centroid (g·mm²); Ixy/Iyz/Izx are products of inertia.
	InertiaXxGmm2 float64 `json:"inertiaXxGmm2"`
	InertiaYyGmm2 float64 `json:"inertiaYyGmm2"`
	InertiaZzGmm2 float64 `json:"inertiaZzGmm2"`
	InertiaXyGmm2 float64 `json:"inertiaXyGmm2"`
	InertiaYzGmm2 float64 `json:"inertiaYzGmm2"`
	InertiaZxGmm2 float64 `json:"inertiaZxGmm2"`
	// Principal moments of inertia (g·mm²) and their axes (unit vectors, rows aligned to the
	// moments) about the centroid.
	PrincipalMomentsGmm2 [3]float64    `json:"principalMomentsGmm2"`
	PrincipalAxes        [3][3]float64 `json:"principalAxes"`
}
