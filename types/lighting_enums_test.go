// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestLightingEnumIdsAreInventorValues pins the frozen Inventor ids for every lighting/shadow
// enum so a rename or renumber is caught (clients and saved automations depend on these exact
// values). The reference is Oblikovati.Contracts.CSharp/Enums.
func TestLightingEnumIdsAreInventorValues(t *testing.T) {
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

// TestLightingEnumNamesAndValidity checks every All* entry is valid with a non-placeholder
// name, and that an undefined value is reported invalid.
func TestLightingEnumNamesAndValidity(t *testing.T) {
	for _, v := range AllLightTypes() {
		if !v.IsValid() || v.String() == "lightType(?)" {
			t.Errorf("LightType %d has no valid name", int32(v))
		}
	}
	for _, v := range AllLightDefinitionTypes() {
		if !v.IsValid() || v.String() == "lightDefinitionType(?)" {
			t.Errorf("LightDefinitionType %d has no valid name", int32(v))
		}
	}
	for _, v := range AllLightingStyleTypes() {
		if !v.IsValid() || v.String() == "lightingStyleType(?)" {
			t.Errorf("LightingStyleType %d has no valid name", int32(v))
		}
	}
	for _, v := range AllShadowDirections() {
		if !v.IsValid() || v.String() == "shadowDirection(?)" {
			t.Errorf("ShadowDirection %d has no valid name", int32(v))
		}
	}
	for _, v := range AllGroundShadows() {
		if !v.IsValid() || v.String() == "groundShadow(?)" {
			t.Errorf("GroundShadow %d has no valid name", int32(v))
		}
	}
	if GroundShadowEnum(0).IsValid() || LightTypeEnum(0).IsValid() {
		t.Errorf("zero value must be invalid for these Inventor-id enums")
	}
}
