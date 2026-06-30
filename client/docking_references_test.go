// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestPanelReferenceListBuilder(t *testing.T) {
	c := PanelReferenceList("faces", "Faces", []string{"face"},
		[]wire.PanelReferenceRow{{Ref: "face/abc"}})
	if c.Kind != types.PanelReferenceList || c.ID != "faces" ||
		c.Accepts[0] != "face" || c.Rows[0].Ref != "face/abc" {
		t.Fatalf("builder produced %+v", c)
	}
}

func TestDockableWindowsSetReferences(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	cl := New(ft)
	if _, err := cl.DockableWindows().SetReferences("w", "faces", []string{"face/abc"}); err != nil {
		t.Fatalf("SetReferences: %v", err)
	}
	if ft.gotMethod != wire.MethodDockableWindowsSetReferences {
		t.Fatalf("method = %q, want %q", ft.gotMethod, wire.MethodDockableWindowsSetReferences)
	}
	var sent wire.SetDockableWindowReferencesArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil ||
		sent.WindowId != "w" || sent.ControlId != "faces" || sent.Refs[0] != "face/abc" {
		t.Fatalf("SetReferences sent %s", ft.gotReq)
	}
}
