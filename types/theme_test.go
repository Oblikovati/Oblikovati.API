// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestParseHexRoundTrip(t *testing.T) {
	for _, hex := range []string{"#000000ff", "#ffffffff", "#3fb6ffff", "#1e2127cc"} {
		c, err := ParseHex(hex)
		if err != nil {
			t.Fatalf("ParseHex(%q) error: %v", hex, err)
		}
		if got := c.Hex(); got != hex {
			t.Errorf("ParseHex(%q).Hex() = %q, want %q", hex, got, hex)
		}
	}
}

func TestParseHexAlphaDefaultsOpaque(t *testing.T) {
	c, err := ParseHex("#3fb6ff")
	if err != nil {
		t.Fatalf("ParseHex error: %v", err)
	}
	if c.A != 1 {
		t.Errorf("ParseHex(\"#3fb6ff\").A = %v, want 1 (opaque)", c.A)
	}
}

func TestParseHexRejectsBad(t *testing.T) {
	for _, bad := range []string{"3fb6ff", "#fff", "#xyzxyz", "#3fb6ff0", ""} {
		if _, err := ParseHex(bad); err == nil {
			t.Errorf("ParseHex(%q) = nil error, want failure", bad)
		}
	}
}

func TestRgbaArray(t *testing.T) {
	c := Rgba{R: 0.25, G: 0.5, B: 0.75, A: 1}
	want := [4]float32{0.25, 0.5, 0.75, 1}
	if c.Array() != want {
		t.Errorf("Array() = %v, want %v", c.Array(), want)
	}
}

// AllThemeTokens must be free of duplicates — the editor and palette both key on it,
// so a repeated token would double a row and shadow a color.
func TestAllThemeTokensUnique(t *testing.T) {
	seen := map[ThemeToken]bool{}
	for _, tk := range AllThemeTokens() {
		if seen[tk] {
			t.Errorf("AllThemeTokens has duplicate %q", tk)
		}
		seen[tk] = true
	}
}

// Regression for Oblikovati#656: TokenViewportActiveBorder was declared but left out of
// AllThemeTokens, so the editor never showed it and completeness tests never checked it.
func TestAllThemeTokensIncludesViewportActiveBorder(t *testing.T) {
	for _, tk := range AllThemeTokens() {
		if tk == TokenViewportActiveBorder {
			return
		}
	}
	t.Errorf("AllThemeTokens() is missing %q", TokenViewportActiveBorder)
}
