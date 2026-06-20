// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// TestDocumentsGetSketchSettings asserts GetSketchSettings addresses the document and decodes the
// returned settings (#147).
func TestDocumentsGetSketchSettings(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"settings":{"inferConstraints":true,"autoApplyConstraints":false,"constraintPriority":50434}}`)}
	c := New(ft)

	got, err := c.Documents().GetSketchSettings(7)
	if err != nil {
		t.Fatalf("GetSketchSettings: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentGetSketchSettings {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentGetSketchSettings)
	}
	var sent wire.GetSketchSettingsArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Document != 7 {
		t.Errorf("document = %d, want 7", sent.Document)
	}
	if !got.Settings.InferConstraints || got.Settings.AutoApplyConstraints {
		t.Errorf("settings = %+v, want infer on / auto-apply off", got.Settings)
	}
}

// TestDocumentsSetSketchSettingsSendsValues asserts SetSketchSettings marshals the document and the
// settings struct (#147).
func TestDocumentsSetSketchSettingsSendsValues(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"settings":{"inferConstraints":false,"autoApplyConstraints":true,"constraintPriority":50433}}`)}
	c := New(ft)

	in := types.SketchSettings{InferConstraints: false, AutoApplyConstraints: true, ConstraintPriority: types.PriorityParallelPerpendicular}
	got, err := c.Documents().SetSketchSettings(9, in)
	if err != nil {
		t.Fatalf("SetSketchSettings: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentSetSketchSettings {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentSetSketchSettings)
	}
	var sent wire.SetSketchSettingsArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Document != 9 || sent.Settings != in {
		t.Errorf("sent = %+v, want document 9 with %+v", sent, in)
	}
	if got.Settings.ConstraintPriority != types.PriorityParallelPerpendicular {
		t.Errorf("decoded priority = %v, want parallelPerpendicular", got.Settings.ConstraintPriority)
	}
}
