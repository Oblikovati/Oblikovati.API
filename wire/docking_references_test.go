// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
)

func TestReferenceListSpecMarshals(t *testing.T) {
	spec := PanelControlSpec{
		Kind:    types.PanelReferenceList,
		ID:      "faces",
		Text:    "Faces",
		Accepts: []string{"face"},
		Rows:    []PanelReferenceRow{{Ref: "face/abc", Label: "Face3"}},
	}
	b, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var back PanelControlSpec
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(back.Rows) != 1 || back.Rows[0].Ref != "face/abc" || back.Accepts[0] != "face" {
		t.Fatalf("round-trip lost data: %+v", back)
	}
}

func TestReferenceMethodAndEventConstants(t *testing.T) {
	if MethodDockableWindowsSetReferences != "dockableWindows.setReferences" {
		t.Fatalf("method = %q", MethodDockableWindowsSetReferences)
	}
	if EventPanelReferencesChanged != "panel.referencesChanged" {
		t.Fatalf("event = %q", EventPanelReferencesChanged)
	}
	var a SetDockableWindowReferencesArgs = SetDockableWindowReferencesArgs{
		WindowId: "w", ControlId: "faces", Refs: []string{"face/abc"},
	}
	if a.Refs[0] != "face/abc" {
		t.Fatalf("args lost refs: %+v", a)
	}
}
