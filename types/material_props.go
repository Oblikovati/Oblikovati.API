// SPDX-License-Identifier: Apache-2.0

package types

// Mechanical groups a material's structural properties. Units follow common
// mechanical-CAD conventions so values transfer 1:1 from existing material libraries.
// The yaml tags keep the on-disk document/library form readable (the tags are plain
// strings, so types still has no yaml dependency).
type Mechanical struct {
	YoungsModulus           float64 `json:"youngsModulus" yaml:"youngsModulus"`                     // GPa
	PoissonsRatio           float64 `json:"poissonsRatio" yaml:"poissonsRatio"`                     // dimensionless
	YieldStrength           float64 `json:"yieldStrength" yaml:"yieldStrength"`                     // MPa
	UltimateTensileStrength float64 `json:"ultimateTensileStrength" yaml:"ultimateTensileStrength"` // MPa
}

// Thermal groups a material's heat-related properties.
type Thermal struct {
	Conductivity   float64 `json:"conductivity" yaml:"conductivity"`     // W/(m·K)
	ExpansionCoeff float64 `json:"expansionCoeff" yaml:"expansionCoeff"` // 1/K (linear)
	SpecificHeat   float64 `json:"specificHeat" yaml:"specificHeat"`     // J/(kg·K)
}

// Electrical groups a material's electrical properties. Common mechanical-CAD material
// models stop at mechanical/thermal; electrical is added here because the user models it
// explicitly.
type Electrical struct {
	Resistivity          float64 `json:"resistivity" yaml:"resistivity"`                   // Ω·m
	RelativePermittivity float64 `json:"relativePermittivity" yaml:"relativePermittivity"` // dimensionless (εr)
}

// IsotropyClass declares a material's elastic symmetry, telling a structural (FEA) solver
// how to read its stiffness. An isotropic material is fully described by the scalar
// [Mechanical] group (E, ν); orthotropic and transversely-isotropic materials additionally
// carry an [AnisotropicElastic] group of direction-dependent constants. The zero value ("")
// is treated as Isotropic, so existing materials need no migration.
type IsotropyClass string

const (
	Isotropic             IsotropyClass = "isotropic"
	Orthotropic           IsotropyClass = "orthotropic"
	TransverselyIsotropic IsotropyClass = "transversely-isotropic"
)

// Anisotropic reports whether the class needs an [AnisotropicElastic] group — i.e. it is
// not isotropic. The empty class counts as isotropic.
func (c IsotropyClass) Anisotropic() bool {
	return c == Orthotropic || c == TransverselyIsotropic
}

// AnisotropicElastic holds the direction-dependent elastic constants of an orthotropic
// material (a transversely-isotropic material is the constrained special case E2=E3,
// G12=G13, ν12=ν13), expressed in the material's principal axes. Axis convention: for wood
// 1 = longitudinal (along grain), 2 = radial, 3 = tangential; for a unidirectional fibre
// lamina 1 = fibre, 2 = 3 = transverse. A solver uses these in place of the scalar
// [Mechanical] E/ν when [IsotropyClass] is not isotropic; yield and ultimate strengths
// still come from [Mechanical] as an equivalent scalar until directional strength
// allowables are added (ADR-0025).
//
// The nine constants are the independent entries of the orthotropic compliance matrix; νIJ
// is the Poisson contraction along J from a load along I, with the symmetry νIJ/Ei = νJI/Ej.
// Alpha1..3 are the linear thermal expansion coefficients along each axis, for orthotropic
// thermal-stress analysis (these can be negative, e.g. the fibre direction of carbon/aramid
// laminates).
type AnisotropicElastic struct {
	E1 float64 `json:"e1" yaml:"e1"` // GPa, axis-1 Young's modulus
	E2 float64 `json:"e2" yaml:"e2"` // GPa, axis-2 Young's modulus
	E3 float64 `json:"e3" yaml:"e3"` // GPa, axis-3 Young's modulus

	G12 float64 `json:"g12" yaml:"g12"` // GPa, shear modulus in the 1-2 plane
	G23 float64 `json:"g23" yaml:"g23"` // GPa, shear modulus in the 2-3 plane
	G13 float64 `json:"g13" yaml:"g13"` // GPa, shear modulus in the 1-3 plane

	Nu12 float64 `json:"nu12" yaml:"nu12"` // major Poisson's ratio, load on 1 → strain on 2
	Nu23 float64 `json:"nu23" yaml:"nu23"` // major Poisson's ratio, load on 2 → strain on 3
	Nu13 float64 `json:"nu13" yaml:"nu13"` // major Poisson's ratio, load on 1 → strain on 3

	Alpha1 float64 `json:"alpha1" yaml:"alpha1"` // 1/K, axis-1 linear thermal expansion
	Alpha2 float64 `json:"alpha2" yaml:"alpha2"` // 1/K, axis-2 linear thermal expansion
	Alpha3 float64 `json:"alpha3" yaml:"alpha3"` // 1/K, axis-3 linear thermal expansion
}

// MassPropertiesAccuracy selects the tessellation fidelity a mass-properties computation uses —
// trading speed for accuracy on curved geometry (planar bodies are exact at any level). The zero
// value is MassPropertiesMedium.
type MassPropertiesAccuracy int32

const (
	// MassPropertiesMedium is the default tessellation fidelity.
	MassPropertiesMedium MassPropertiesAccuracy = iota
	// MassPropertiesLow is a coarse, faster tessellation.
	MassPropertiesLow
	// MassPropertiesHigh is a fine, slower tessellation for tighter accuracy on curved bodies.
	MassPropertiesHigh
)

var massPropertiesAccuracyNames = map[MassPropertiesAccuracy]string{
	MassPropertiesMedium: "medium",
	MassPropertiesLow:    "low",
	MassPropertiesHigh:   "high",
}

// String returns the accuracy level's wire spelling ("medium", "low", "high").
func (a MassPropertiesAccuracy) String() string { return enumName(massPropertiesAccuracyNames, a) }

// ParseMassPropertiesAccuracy resolves a wire spelling back to its accuracy level.
func ParseMassPropertiesAccuracy(s string) (MassPropertiesAccuracy, bool) {
	return enumFromName(massPropertiesAccuracyNames, s)
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
