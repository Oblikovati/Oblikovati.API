// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestGeometricCharacteristicFrozenBlock pins the reference bit ids and wire spellings — the
// block is frozen and a renumber or respell is a breaking wire change.
func TestGeometricCharacteristicFrozenBlock(t *testing.T) {
	want := map[GeometricCharacteristic]string{
		1: "straightness", 2: "flatness", 4: "circularity", 8: "profileOfAnyLine",
		16: "profileOfAnySurface", 32: "angularity", 64: "perpendicularity",
		128: "parallelism", 256: "position", 512: "concentricity", 1024: "circularRunout",
		2048: "symmetry", 4096: "totalRunout", 8192: "cylindricity", 16384: "parallelProfile",
		32768: "axisIntersection", 65536: "circularRunoutFilled", 131072: "totalRunoutFilled",
		262144: "profileOfASection", 524288: "axiality",
	}
	if len(want) != len(geometricCharacteristicNames) {
		t.Fatalf("characteristic count = %d, want %d", len(geometricCharacteristicNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("GeometricCharacteristic(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseGeometricCharacteristic(name); !ok || parsed != v {
			t.Errorf("ParseGeometricCharacteristic(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	if _, ok := ParseGeometricCharacteristic("noSuchCharacteristic"); ok {
		t.Error("ParseGeometricCharacteristic must reject unknown spellings")
	}
}
