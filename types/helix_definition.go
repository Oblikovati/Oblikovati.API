// SPDX-License-Identifier: Apache-2.0

package types

// HelicalShapeDefinitionKind discriminates how a helical curve's shape is
// specified (M06-F09, Oblikovati/Oblikovati#624). The wire spellings double as
// the helix Mode strings already accepted by sketch3d.addEntity since M22-F04.
//
// The values are a frozen block matching the reference API's helical-shape
// definition enum; never renumber them.
type HelicalShapeDefinitionKind int32

const (
	// HelixShapePitchRevolution fixes pitch and turn count (height follows).
	HelixShapePitchRevolution HelicalShapeDefinitionKind = 115713
	// HelixShapePitchHeight fixes pitch and total height (turns follow).
	HelixShapePitchHeight HelicalShapeDefinitionKind = 115714
	// HelixShapeRevolutionHeight fixes turn count and total height (pitch follows).
	HelixShapeRevolutionHeight HelicalShapeDefinitionKind = 115715
	// HelixShapeSpiral is a flat spiral: no axial advance, radial growth per turn.
	HelixShapeSpiral HelicalShapeDefinitionKind = 115716
)

// helicalShapeDefinitionNames are the frozen wire spellings.
var helicalShapeDefinitionNames = map[HelicalShapeDefinitionKind]string{
	HelixShapePitchRevolution:  "pitchRevolution",
	HelixShapePitchHeight:      "pitchHeight",
	HelixShapeRevolutionHeight: "revolutionHeight",
	HelixShapeSpiral:           "spiral",
}

// String returns the shape kind's wire spelling.
func (k HelicalShapeDefinitionKind) String() string {
	return enumName(helicalShapeDefinitionNames, k, "enum(?)")
}

// ParseHelicalShapeDefinitionKind resolves a wire spelling back to its kind.
func ParseHelicalShapeDefinitionKind(s string) (HelicalShapeDefinitionKind, bool) {
	return enumFromName(helicalShapeDefinitionNames, s)
}

// HelixEndKind is the transition condition at one end of a helical curve
// (M06-F09). A natural end stops on the helix; a flat end appends a planar
// transition (governed by the definition's transition/flat angles) so the
// coil seats flat — the standard spring end treatment.
//
// The values are a frozen block matching the reference API's helix-end enum;
// never renumber them.
type HelixEndKind int32

const (
	HelixEndNatural HelixEndKind = 115969
	HelixEndFlat    HelixEndKind = 115970
)

// helixEndNames are the frozen wire spellings.
var helixEndNames = map[HelixEndKind]string{
	HelixEndNatural: "natural",
	HelixEndFlat:    "flat",
}

// String returns the end kind's wire spelling.
func (k HelixEndKind) String() string { return enumName(helixEndNames, k, "enum(?)") }

// ParseHelixEndKind resolves a wire spelling back to its kind.
func ParseHelixEndKind(s string) (HelixEndKind, bool) { return enumFromName(helixEndNames, s) }
