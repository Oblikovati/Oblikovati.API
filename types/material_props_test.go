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

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
