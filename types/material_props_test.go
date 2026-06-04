// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"
	"testing"
)

func TestAssetSourceEditable(t *testing.T) {
	if AssetBuiltin.Editable() {
		t.Error("built-in assets must be read-only")
	}
	for _, s := range []AssetSource{AssetProject, AssetDocument} {
		if !s.Editable() {
			t.Errorf("%s should be editable", s)
		}
	}
}

// Property groups must serialize with stable JSON keys (the wire DTOs embed them).
func TestMechanicalJSONKeys(t *testing.T) {
	b, err := json.Marshal(Mechanical{YoungsModulus: 69, PoissonsRatio: 0.33})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	got := string(b)
	for _, key := range []string{"youngsModulus", "poissonsRatio", "yieldStrength", "ultimateTensileStrength"} {
		if !contains(got, key) {
			t.Errorf("Mechanical JSON %q missing key %q", got, key)
		}
	}
}

// TestIsotropyClassAnisotropic pins the symmetry helper a solver branches on: only the two
// direction-dependent classes need the AnisotropicElastic group; the empty class is
// isotropic so existing materials need no migration.
func TestIsotropyClassAnisotropic(t *testing.T) {
	cases := map[IsotropyClass]bool{
		"":                    false, // unset == isotropic
		Isotropic:             false,
		Orthotropic:           true,
		TransverselyIsotropic: true,
	}
	for class, want := range cases {
		if got := class.Anisotropic(); got != want {
			t.Errorf("IsotropyClass(%q).Anisotropic() = %v, want %v", class, got, want)
		}
	}
}

// TestAnisotropicElasticJSONKeys guards the wire keys an FEA add-in reads off MaterialInfo.
func TestAnisotropicElasticJSONKeys(t *testing.T) {
	b, err := json.Marshal(AnisotropicElastic{E1: 135, E2: 10})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	got := string(b)
	for _, key := range []string{"e1", "e2", "e3", "g12", "g23", "g13", "nu12", "nu23", "nu13", "alpha1", "alpha2", "alpha3"} {
		if !contains(got, key) {
			t.Errorf("AnisotropicElastic JSON %q missing key %q", got, key)
		}
	}
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
