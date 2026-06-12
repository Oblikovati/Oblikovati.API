// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestQueryEnumValuesAreFrozen pins the reference enum values.
func TestQueryEnumValuesAreFrozen(t *testing.T) {
	if UnknownSolutionNature != 0 || UniqueSolution != 1 || DistinctlyManySolutions != 2 ||
		InfinitelyManySolutions != 3 || NoSolution != 4 {
		t.Error("SolutionNature values drifted from the frozen block 0…4")
	}
	if UnknownContainment != 30977 || InsideContainment != 30978 ||
		OnContainment != 30979 || OutsideContainment != 30980 {
		t.Error("Containment values drifted from the frozen block 30977…")
	}
}

func TestQueryEnumNames(t *testing.T) {
	if UniqueSolution.String() != "unique" || NoSolution.String() != "none" {
		t.Errorf("SolutionNature names: %s, %s", UniqueSolution, NoSolution)
	}
	if OnContainment.String() != "on" || Containment(1).String() != "containment(?)" {
		t.Errorf("Containment names: %s, %s", OnContainment, Containment(1))
	}
}
