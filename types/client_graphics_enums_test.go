// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestClientGraphicsEnumIdsAreStable pins the frozen ids for every client-graphics object-model
// enum so a rename or renumber is caught (add-ins and cached graphics depend on these values).
func TestClientGraphicsEnumIdsAreStable(t *testing.T) {
	cg := map[ClientGraphicsTypeEnum]int32{TransientClientGraphics: 45313, PreviewClientGraphics: 45314}
	for v, id := range cg {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	cb := map[ColorBindingEnum]int32{
		PerVertexColors: 19457, PerStripColors: 19458, PerItemColors: 19459, OverallColor: 19460,
	}
	for v, id := range cb {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	nb := map[NormalBindingEnum]int32{
		PerVertexNormals: 19713, PerStripNormals: 19714, PerItemNormals: 19715, OverallNormal: 19716,
	}
	for v, id := range nb {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	sel := map[GraphicsSelectabilityEnum]int32{
		NoGraphicsSelectable: 25345, SomeGraphicsSelectable: 25346, AllGraphicsSelectable: 25347,
	}
	for v, id := range sel {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	vis := map[GraphicsVisibilityEnum]int32{
		NoGraphicsVisible: 25089, SomeGraphicsVisible: 25090, AllGraphicsVisible: 25091,
	}
	for v, id := range vis {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	pt := map[PointRenderStyleEnum]int32{
		XPointStyle: 20225, CirclePointStyle: 20226, FilledCirclePointStyle: 20228,
		CrossPointStyle: 20230, CustomImagePointStyle: 20235,
	}
	for v, id := range pt {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	cs := map[CachedGraphicsStatusEnum]int32{
		NoneCachedGraphics: 103169, OutOfDateCachedGraphics: 103170, UpToDateCachedGraphics: 103171,
	}
	for v, id := range cs {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	ld := map[LineDefinitionSpaceEnum]int32{ScreenSpace: 49921, ModelSpace: 49922, HybridSpace: 49923}
	for v, id := range ld {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
}

// TestClientGraphicsEnumNamesAndValidity checks every All* entry is valid with a non-placeholder
// name, and the zero value is reported invalid.
func TestClientGraphicsEnumNamesAndValidity(t *testing.T) {
	assertNamed(t, "ClientGraphicsType", "clientGraphicsType(?)", AllClientGraphicsTypes())
	assertNamed(t, "ColorBinding", "colorBinding(?)", AllColorBindings())
	assertNamed(t, "NormalBinding", "normalBinding(?)", AllNormalBindings())
	assertNamed(t, "GraphicsSelectability", "graphicsSelectability(?)", AllGraphicsSelectabilities())
	assertNamed(t, "GraphicsVisibility", "graphicsVisibility(?)", AllGraphicsVisibilities())
	assertNamed(t, "PointRenderStyle", "pointRenderStyle(?)", AllPointRenderStyles())
	assertNamed(t, "CachedGraphicsStatus", "cachedGraphicsStatus(?)", AllCachedGraphicsStatuses())
	assertNamed(t, "LineDefinitionSpace", "lineDefinitionSpace(?)", AllLineDefinitionSpaces())
	if ColorBindingEnum(0).IsValid() || PointRenderStyleEnum(0).IsValid() || LineDefinitionSpaceEnum(0).IsValid() {
		t.Errorf("zero value must be invalid for these id-based enums")
	}
}
