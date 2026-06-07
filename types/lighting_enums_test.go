// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestLightingEnumIdsAreStable pins the frozen ids for every lighting/shadow
// enum so a rename or renumber is caught (clients and saved automations depend on these exact
// values).
func TestLightingEnumIdsAreStable(t *testing.T) {
	light := map[LightTypeEnum]int32{
		ModelSpaceLight: 52993, ViewSpaceLight: 52994, GroundPlaneSpaceLight: 52995,
	}
	for v, id := range light {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	def := map[LightDefinitionTypeEnum]int32{
		DirectionalLight: 53249, PointLight: 53250, SpotLight: 53251,
	}
	for v, id := range def {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	style := map[LightingStyleTypeEnum]int32{
		StandardLightingStyle: 50750977, ImageBasedLightingStyle: 50750978,
	}
	for v, id := range style {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	ground := map[GroundShadowEnum]int32{
		NoGroundShadow: 69121, GroundShadow: 69122, XRayGroundShadow: 69123,
	}
	for v, id := range ground {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	dir := map[ShadowDirectionEnum]int32{
		FortyFiveDegreesLeftShadow: 92161, FortyFiveDegreesRightShadow: 92162,
		AboveShadow: 92163, LightOneShadow: 92164, LightTwoShadow: 92165,
		LightThreeShadow: 92166, LightFourShadow: 92167, EnvironmentShadow: 92168,
	}
	for v, id := range dir {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
}

// namedEnum is the shared shape of the lighting/shadow id-enums: a validity check plus a
// String() that returns a placeholder ("lightType(?)") for values with no defined name.
type namedEnum interface {
	IsValid() bool
	String() string
}

// assertNamed fails t for any value that is invalid or stringifies to placeholder, so the
// per-enum loops below collapse to one call each (keeps the test's complexity low).
func assertNamed[T namedEnum](t *testing.T, label, placeholder string, values []T) {
	t.Helper()
	for _, v := range values {
		if !v.IsValid() || v.String() == placeholder {
			t.Errorf("%s %v has no valid name", label, v)
		}
	}
}

// TestLightingEnumNamesAndValidity checks every All* entry is valid with a non-placeholder
// name, and that an undefined value is reported invalid.
func TestLightingEnumNamesAndValidity(t *testing.T) {
	assertNamed(t, "LightType", "lightType(?)", AllLightTypes())
	assertNamed(t, "LightDefinitionType", "lightDefinitionType(?)", AllLightDefinitionTypes())
	assertNamed(t, "LightingStyleType", "lightingStyleType(?)", AllLightingStyleTypes())
	assertNamed(t, "ShadowDirection", "shadowDirection(?)", AllShadowDirections())
	assertNamed(t, "GroundShadow", "groundShadow(?)", AllGroundShadows())
	if GroundShadowEnum(0).IsValid() || LightTypeEnum(0).IsValid() {
		t.Errorf("zero value must be invalid for these id-based enums")
	}
}
