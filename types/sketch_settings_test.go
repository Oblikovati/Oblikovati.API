// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"
	"testing"
)

// TestDefaultSketchSettings pins the out-of-the-box configuration: inference and auto-apply on,
// horizontal/vertical priority — the behaviour a new document had before this surface existed.
func TestDefaultSketchSettings(t *testing.T) {
	s := DefaultSketchSettings()
	if !s.InferConstraints || !s.AutoApplyConstraints {
		t.Errorf("defaults = %+v, want inference and auto-apply on", s)
	}
	if s.ConstraintPriority != PriorityHorizontalVertical {
		t.Errorf("default priority = %v, want horizontalVertical", s.ConstraintPriority)
	}
}

// TestSketchSettingsJSONRoundTrips asserts the wire shape is stable: the field keys and the integer
// priority encoding survive a marshal/unmarshal unchanged.
func TestSketchSettingsJSONRoundTrips(t *testing.T) {
	in := SketchSettings{InferConstraints: true, AutoApplyConstraints: false, ConstraintPriority: PriorityNone}
	b, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	const want = `{"inferConstraints":true,"autoApplyConstraints":false,"constraintPriority":50435}`
	if string(b) != want {
		t.Errorf("JSON = %s, want %s", b, want)
	}
	var out SketchSettings
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out != in {
		t.Errorf("round-trip = %+v, want %+v", out, in)
	}
}
