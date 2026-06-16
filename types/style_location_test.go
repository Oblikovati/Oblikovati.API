// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestStyleLocationIdsAreStable pins the frozen ids for the style-location cascade enum.
func TestStyleLocationIdsAreStable(t *testing.T) {
	want := map[StyleLocationEnum]int32{
		BothStyleLocation: 51201, LocalStyleLocation: 51202, LibraryStyleLocation: 51203,
	}
	for v, id := range want {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
}

// TestStyleLocationNamesAndValidity checks every All* entry is valid with a non-placeholder
// name, and the zero value is reported invalid.
func TestStyleLocationNamesAndValidity(t *testing.T) {
	assertNamed(t, "StyleLocation", "styleLocation(?)", AllStyleLocations())
	if StyleLocationEnum(0).IsValid() {
		t.Errorf("zero value must be invalid for this id-based enum")
	}
}
