// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestProjectCurveToSurfaceTypeFrozenBlock pins the reference ids and wire spellings (#1841).
func TestProjectCurveToSurfaceTypeFrozenBlock(t *testing.T) {
	want := map[ProjectCurveToSurfaceType]string{
		117505: "alongVector", 117506: "closestPoint", 117507: "wrap",
	}
	assertFrozenBlock(t, "ProjectCurveToSurfaceType", want, projectCurveToSurfaceTypeNames,
		ParseProjectCurveToSurfaceType)
}

// TestParseProjectCurveToSurfaceTypeEmptyIsClosest checks the omitted-selector default (#1841).
func TestParseProjectCurveToSurfaceTypeEmptyIsClosest(t *testing.T) {
	if got, ok := ParseProjectCurveToSurfaceType(""); !ok || got != ProjectToClosestPoint {
		t.Errorf(`ParseProjectCurveToSurfaceType("") = (%v, %v), want closestPoint, true`, got, ok)
	}
}
