// SPDX-License-Identifier: Apache-2.0

package types

// GeometricCharacteristic is the geometric-tolerance symbol carried by a feature-control
// frame in a model GD&T annotation (parity: GeometricCharacteristicEnum). The values are the
// reference API's frozen bit ids and must never be renumbered.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases it (ADR-0018).
// Used by the model-tolerance carrier (a metadata feature; #866).
type GeometricCharacteristic int32

const (
	CharacteristicStraightness         GeometricCharacteristic = 1
	CharacteristicFlatness             GeometricCharacteristic = 2
	CharacteristicCircularity          GeometricCharacteristic = 4
	CharacteristicProfileOfAnyLine     GeometricCharacteristic = 8
	CharacteristicProfileOfAnySurface  GeometricCharacteristic = 16
	CharacteristicAngularity           GeometricCharacteristic = 32
	CharacteristicPerpendicularity     GeometricCharacteristic = 64
	CharacteristicParallelism          GeometricCharacteristic = 128
	CharacteristicPosition             GeometricCharacteristic = 256
	CharacteristicConcentricity        GeometricCharacteristic = 512
	CharacteristicCircularRunout       GeometricCharacteristic = 1024
	CharacteristicSymmetry             GeometricCharacteristic = 2048
	CharacteristicTotalRunout          GeometricCharacteristic = 4096
	CharacteristicCylindricity         GeometricCharacteristic = 8192
	CharacteristicParallelProfile      GeometricCharacteristic = 16384
	CharacteristicAxisIntersection     GeometricCharacteristic = 32768
	CharacteristicCircularRunoutFilled GeometricCharacteristic = 65536
	CharacteristicTotalRunoutFilled    GeometricCharacteristic = 131072
	CharacteristicProfileOfASection    GeometricCharacteristic = 262144
	CharacteristicAxiality             GeometricCharacteristic = 524288
)

// geometricCharacteristicNames are the wire spellings (model-tolerance feature-control frames).
var geometricCharacteristicNames = map[GeometricCharacteristic]string{
	CharacteristicStraightness:         "straightness",
	CharacteristicFlatness:             "flatness",
	CharacteristicCircularity:          "circularity",
	CharacteristicProfileOfAnyLine:     "profileOfAnyLine",
	CharacteristicProfileOfAnySurface:  "profileOfAnySurface",
	CharacteristicAngularity:           "angularity",
	CharacteristicPerpendicularity:     "perpendicularity",
	CharacteristicParallelism:          "parallelism",
	CharacteristicPosition:             "position",
	CharacteristicConcentricity:        "concentricity",
	CharacteristicCircularRunout:       "circularRunout",
	CharacteristicSymmetry:             "symmetry",
	CharacteristicTotalRunout:          "totalRunout",
	CharacteristicCylindricity:         "cylindricity",
	CharacteristicParallelProfile:      "parallelProfile",
	CharacteristicAxisIntersection:     "axisIntersection",
	CharacteristicCircularRunoutFilled: "circularRunoutFilled",
	CharacteristicTotalRunoutFilled:    "totalRunoutFilled",
	CharacteristicProfileOfASection:    "profileOfASection",
	CharacteristicAxiality:             "axiality",
}

// String returns the characteristic's wire spelling.
func (c GeometricCharacteristic) String() string {
	return enumName(geometricCharacteristicNames, c, "enum(?)")
}

// ParseGeometricCharacteristic resolves a wire spelling back to its characteristic.
func ParseGeometricCharacteristic(s string) (GeometricCharacteristic, bool) {
	return enumFromName(geometricCharacteristicNames, s)
}
