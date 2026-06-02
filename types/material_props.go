// SPDX-License-Identifier: Apache-2.0

package types

// Mechanical groups a material's structural properties (Inventor Material parity). Units
// follow Inventor's conventions so values transfer 1:1 from existing material libraries.
type Mechanical struct {
	YoungsModulus           float64 `json:"youngsModulus"`           // GPa
	PoissonsRatio           float64 `json:"poissonsRatio"`           // dimensionless
	YieldStrength           float64 `json:"yieldStrength"`           // MPa
	UltimateTensileStrength float64 `json:"ultimateTensileStrength"` // MPa
}

// Thermal groups a material's heat-related properties.
type Thermal struct {
	Conductivity   float64 `json:"conductivity"`   // W/(m·K)
	ExpansionCoeff float64 `json:"expansionCoeff"` // 1/K (linear)
	SpecificHeat   float64 `json:"specificHeat"`   // J/(kg·K)
}

// Electrical groups a material's electrical properties. Inventor's Material stops at
// mechanical/thermal; electrical is added here because the user models it explicitly.
type Electrical struct {
	Resistivity          float64 `json:"resistivity"`          // Ω·m
	RelativePermittivity float64 `json:"relativePermittivity"` // dimensionless (εr)
}

// PhysicalProperties is the computed mass/geometry summary of a body or part given its
// material's density. Volume/area/centroid come from geometry; mass = density × volume.
// Lengths are in database units (cm), so volume is cm³, area cm², and — with density in
// g/cm³ — mass is grams.
type PhysicalProperties struct {
	Mass     float64    `json:"mass"`     // g
	Volume   float64    `json:"volume"`   // cm³
	Area     float64    `json:"area"`     // cm²
	Density  float64    `json:"density"`  // g/cm³
	Centroid [3]float64 `json:"centroid"` // center of mass (cm)
}
