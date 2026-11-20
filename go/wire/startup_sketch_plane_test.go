// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
)

func TestStartupSketchPlaneWire(t *testing.T) {
	cases := []types.StartupSketchPlane{types.StartupSketchNone, types.StartupSketchXY, types.StartupSketchYZ, types.StartupSketchXZ}
	for _, plane := range cases {
		encoded, err := json.Marshal(PartOptionsView{StartupSketchPlane: &plane})
		if err != nil {
			t.Fatal(err)
		}
		var decoded PartOptionsView
		if err := json.Unmarshal(encoded, &decoded); err != nil {
			t.Fatal(err)
		}
		if decoded.StartupSketchPlane == nil || *decoded.StartupSketchPlane != plane {
			t.Fatalf("plane %q: got %s", plane, encoded)
		}
	}
	var old PartOptionsView
	if err := json.Unmarshal([]byte(`{"chamferFlatCorners":true}`), &old); err != nil {
		t.Fatal(err)
	}
	if old.StartupSketchPlane != nil {
		t.Fatalf("legacy plane %v: expected nil", old.StartupSketchPlane)
	}
	encoded, err := json.Marshal(old)
	if err != nil || string(encoded) != `{"chamferFlatCorners":true}` {
		t.Fatalf("legacy round trip %s, error %v: expected unchanged fields", encoded, err)
	}
	for _, raw := range []string{`{"startupSketchPlane":null}`, `{"startupSketchPlane":"bogus"}`} {
		var view PartOptionsView
		if err := json.Unmarshal([]byte(raw), &view); err != nil {
			t.Fatal(err)
		}
		if view.StartupSketchPlane != nil && view.StartupSketchPlane.IsValid() {
			t.Fatalf("plane in %s: expected nil or invalid", raw)
		}
	}
}
