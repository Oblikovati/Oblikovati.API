// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestDisplayEnumIdsAreStable pins the frozen ids for every display-options enum so a rename
// or renumber is caught (clients and saved display settings depend on these exact values).
func TestDisplayEnumIdsAreStable(t *testing.T) {
	quality := map[DisplayQualityEnum]int32{
		SmoothDisplayQuality: 58881, MediumDisplayQuality: 58882,
		RoughDisplayQuality: 58883, SmootherDisplayQuality: 58884,
	}
	for v, id := range quality {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	source := map[DisplayModeSourceTypeEnum]int32{
		DefaultDisplayModeSource: 54273, OverrideDisplayModeSource: 54274,
	}
	for v, id := range source {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	sep := map[DisplaySeparateColorsTypeEnum]int32{
		NoneDisplaySeparateColors: 127489, EachUniqueTopLevelComponentColors: 127490,
		EachTopLevelComponentColors: 127491, EachUniquePartColors: 127492, EachPartColors: 127493,
	}
	for v, id := range sep {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	transform := map[DisplayTransformBehaviorEnum]int32{
		FrontFacingBehavior: 26113, PixelScalingBehavior: 26114,
		FrontFacingAndPixelScalingBehavior: 26115, NoTransformBehaviors: 26116,
	}
	for v, id := range transform {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	cull := map[BackFaceCullingEnum]int32{
		CullNone: 96769, CullClockwise: 96770, CullCounterClockwise: 96771,
	}
	for v, id := range cull {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	transp := map[TransparencyTypeEnum]int32{
		BlendingTransparency: 58625, ScreenDoorTransparency: 58626,
	}
	for v, id := range transp {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	rt := map[RayTracingQualityEnum]int32{
		InteractiveRayTracingQuality: 95745, GoodRayTracingQuality: 95746, BestRayTracingQuality: 95747,
		LowRayTracingQuality: 95748, DraftRayTracingQuality: 95749, HighRayTracingQuality: 95750,
	}
	for v, id := range rt {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	proj := map[ProjectionTypeEnum]int32{
		OrthographicProjection: 86273, PerspectiveProjection: 86274, PerspectiveWithOrthoFacesProjection: 86275,
	}
	for v, id := range proj {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
}

// TestDisplayEnumNamesAndValidity checks every All* entry is valid with a non-placeholder
// name, and that the zero value is reported invalid for these id-based enums.
func TestDisplayEnumNamesAndValidity(t *testing.T) {
	assertNamed(t, "DisplayQuality", "displayQuality(?)", AllDisplayQualities())
	assertNamed(t, "DisplayModeSource", "displayModeSource(?)", AllDisplayModeSources())
	assertNamed(t, "DisplaySeparateColors", "displaySeparateColors(?)", AllDisplaySeparateColors())
	assertNamed(t, "DisplayTransformBehavior", "displayTransformBehavior(?)", AllDisplayTransformBehaviors())
	assertNamed(t, "BackFaceCulling", "backFaceCulling(?)", AllBackFaceCullings())
	assertNamed(t, "TransparencyType", "transparencyType(?)", AllTransparencyTypes())
	assertNamed(t, "RayTracingQuality", "rayTracingQuality(?)", AllRayTracingQualities())
	assertNamed(t, "ProjectionType", "projectionType(?)", AllProjectionTypes())
	if DisplayQualityEnum(0).IsValid() || ProjectionTypeEnum(0).IsValid() || RayTracingQualityEnum(0).IsValid() {
		t.Errorf("zero value must be invalid for these id-based enums")
	}
}
