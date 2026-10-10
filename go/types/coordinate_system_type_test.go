// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestCoordinateSystemTypeFrozenBlock pins the Oblikovati-owned ids and wire spellings (#1846).
func TestCoordinateSystemTypeFrozenBlock(t *testing.T) {
	want := map[CoordinateSystemType]string{
		0: "cartesian", 1: "cylindrical", 2: "spherical",
	}
	assertFrozenBlock(t, "CoordinateSystemType", want, coordinateSystemTypeNames, ParseCoordinateSystemType)
}

// TestParseCoordinateSystemTypeEmptyIsCartesian checks the omitted-selector default (#1846).
func TestParseCoordinateSystemTypeEmptyIsCartesian(t *testing.T) {
	if got, ok := ParseCoordinateSystemType(""); !ok || got != CoordinateSystemCartesian {
		t.Errorf(`ParseCoordinateSystemType("") = (%v, %v), want cartesian, true`, got, ok)
	}
}
