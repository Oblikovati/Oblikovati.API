// SPDX-License-Identifier: Apache-2.0
package wire

import (
	"encoding/json"
	"testing"
)

func TestPartSketchViewDistinguishesOmissionFalseAndTrue(t *testing.T) {
	disabled, enabled := false, true
	for _, expected := range []*bool{nil, &disabled, &enabled} {
		raw, err := json.Marshal(PartOptionsView{LookAtSketchPlane: expected})
		if err != nil {
			t.Fatal(err)
		}
		var decoded PartOptionsView
		if err := json.Unmarshal(raw, &decoded); err != nil {
			t.Fatal(err)
		}
		if (decoded.LookAtSketchPlane == nil) != (expected == nil) {
			t.Fatalf("%s: expected optional value %v", raw, expected)
		}
		if expected != nil && *decoded.LookAtSketchPlane != *expected {
			t.Fatalf("%s: expected %v", raw, *expected)
		}
	}
}
