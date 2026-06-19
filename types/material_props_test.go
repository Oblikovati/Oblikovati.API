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

// TestMagneticJSONKeys guards the wire keys a magnetostatics add-in reads off MaterialInfo.
func TestMagneticJSONKeys(t *testing.T) {
	b, err := json.Marshal(Magnetic{Class: HardMagnetic, Remanence: 1.30, Coercivity: 915, RelativePermeability: 1.05})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	got := string(b)
	for _, key := range []string{"class", "relativePermeability", "remanence", "coercivity"} {
		if !contains(got, key) {
			t.Errorf("Magnetic JSON %q missing key %q", got, key)
		}
	}
}

// TestMagneticZeroValueOmitted pins that a non-magnetic material (the zero value) emits no
// magnetic fields, so the overwhelming majority of materials add no YAML/JSON noise.
func TestMagneticZeroValueOmitted(t *testing.T) {
	b, err := json.Marshal(Magnetic{})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if got := string(b); got != "{}" {
		t.Errorf("zero Magnetic JSON = %q, want {} (all fields omitempty)", got)
	}
}

// TestMagneticIsMagnetic pins the helper a solver branches on: only soft/hard classes read
// the group; the non-magnetic default (incl. the empty class) is treated as free space.
func TestMagneticIsMagnetic(t *testing.T) {
	cases := map[MagneticClass]bool{
		"":           false,
		NonMagnetic:  false,
		SoftMagnetic: true,
		HardMagnetic: true,
	}
	for class, want := range cases {
		if got := (Magnetic{Class: class}).IsMagnetic(); got != want {
			t.Errorf("Magnetic{Class:%q}.IsMagnetic() = %v, want %v", class, got, want)
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
