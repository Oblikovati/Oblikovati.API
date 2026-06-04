// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestDisplayModeIdsAreInventorValues pins the frozen Inventor ids so a rename or renumber is
// caught (clients and saved automations depend on these exact values).
func TestDisplayModeIdsAreInventorValues(t *testing.T) {
	cases := map[DisplayModeEnum]int32{
		WireframeRendering: 8706, ShadedRendering: 8708, RealisticRendering: 8709,
		ShadedWithEdgesRendering: 8710, WireframeNoHiddenEdges: 8711,
		WireframeWithHiddenEdgesRendering: 8712, MonochromeRendering: 8713,
		WatercolorRendering: 8714, IllustrationRendering: 8715, TechnicalIllustrationRendering: 8716,
	}
	for mode, id := range cases {
		if int32(mode) != id {
			t.Errorf("%s id = %d, want %d", mode, int32(mode), id)
		}
	}
}

// TestHiddenEdgeAliasShares8707 documents that the two Inventor names for 8707 are aliases.
func TestHiddenEdgeAliasShares8707(t *testing.T) {
	if HiddenEdgeRendering != ShadedWithHiddenEdgesRendering || int32(HiddenEdgeRendering) != 8707 {
		t.Errorf("8707 alias broken: HiddenEdge=%d ShadedWithHiddenEdges=%d",
			int32(HiddenEdgeRendering), int32(ShadedWithHiddenEdgesRendering))
	}
}

// TestAllDisplayModesCoverGallery checks the gallery lists all 11 distinct modes and that each
// is valid with a non-placeholder name.
func TestAllDisplayModesCoverGallery(t *testing.T) {
	modes := AllDisplayModes()
	if len(modes) != 11 {
		t.Fatalf("AllDisplayModes = %d, want 11", len(modes))
	}
	seen := map[DisplayModeEnum]bool{}
	for _, m := range modes {
		if !m.IsValid() || m.String() == "displayMode(?)" {
			t.Errorf("mode %d has no valid name", int32(m))
		}
		if seen[m] {
			t.Errorf("duplicate mode %s in gallery", m)
		}
		seen[m] = true
	}
}

// TestUnknownDisplayModeIsInvalid checks an out-of-range id is reported invalid.
func TestUnknownDisplayModeIsInvalid(t *testing.T) {
	if DisplayModeEnum(9999).IsValid() {
		t.Error("9999 should be an invalid display mode")
	}
}
