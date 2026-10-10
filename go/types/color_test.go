// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestColorEnumIdsAreStable pins the frozen ids for the color/background enums so a rename or
// renumber is caught (clients and saved schemes depend on these exact values).
func TestColorEnumIdsAreStable(t *testing.T) {
	source := map[ColorSourceTypeEnum]int32{
		OverrideColorSource: 79105, AutomaticColorSource: 79106,
		LayerColorSource: 79107, SheetColorSource: 79108,
	}
	for v, id := range source {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
	bg := map[BackgroundTypeEnum]int32{
		OneColorBackground: 52737, GradientBackground: 52738, ImageBackground: 52739,
	}
	for v, id := range bg {
		if int32(v) != id {
			t.Errorf("%s id = %d, want %d", v, int32(v), id)
		}
	}
}

// TestColorEnumNamesAndValidity checks every All* entry is valid with a non-placeholder name,
// and that an undefined value is reported invalid.
func TestColorEnumNamesAndValidity(t *testing.T) {
	assertNamed(t, "ColorSourceType", "colorSource(?)", AllColorSources())
	assertNamed(t, "BackgroundType", "backgroundType(?)", AllBackgroundTypes())
	if ColorSourceTypeEnum(0).IsValid() || BackgroundTypeEnum(0).IsValid() {
		t.Errorf("zero value must be invalid for these id-based enums")
	}
}

// TestNewColorIsOpaqueOverride checks the common-case constructor sets opaque opacity and an
// override source.
func TestNewColorIsOpaqueOverride(t *testing.T) {
	c := NewColor(10, 20, 30)
	if c.R != 10 || c.G != 20 || c.B != 30 {
		t.Errorf("NewColor components = (%d,%d,%d), want (10,20,30)", c.R, c.G, c.B)
	}
	if c.Opacity != 1 {
		t.Errorf("NewColor opacity = %v, want 1", c.Opacity)
	}
	if c.Source != OverrideColorSource {
		t.Errorf("NewColor source = %s, want Override", c.Source)
	}
}

// TestColorRgbaConversion checks 8-bit components and opacity map onto the renderer's [0,1]
// float channels, white round-tripping to all-ones.
func TestColorRgbaConversion(t *testing.T) {
	got := NewColor(255, 255, 255).Rgba()
	if got.R != 1 || got.G != 1 || got.B != 1 || got.A != 1 {
		t.Errorf("white Rgba = %+v, want all ones", got)
	}
	half := Color{R: 0, G: 0, B: 0, Opacity: 0.5}.Rgba()
	if half.A != 0.5 {
		t.Errorf("alpha = %v, want 0.5", half.A)
	}
}

// TestColorHex formats opaque components as "#rrggbb", ignoring opacity.
func TestColorHex(t *testing.T) {
	if got := NewColor(255, 0, 128).Hex(); got != "#ff0080" {
		t.Errorf("Hex = %q, want #ff0080", got)
	}
}
