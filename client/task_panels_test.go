// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestTaskPanelsShowAndClose(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	cl := New(ft)
	if _, err := cl.TaskPanels().Show(wire.TaskPanelSpec{ID: "fix", Title: "Fixed"}); err != nil {
		t.Fatalf("Show: %v", err)
	}
	if ft.gotMethod != wire.MethodTaskPanelShow {
		t.Fatalf("method = %q, want %q", ft.gotMethod, wire.MethodTaskPanelShow)
	}
	var shown wire.ShowTaskPanelArgs
	if err := json.Unmarshal(ft.gotReq, &shown); err != nil || shown.Panel.ID != "fix" {
		t.Fatalf("Show sent %s", ft.gotReq)
	}
	if _, err := cl.TaskPanels().Close("fix"); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if ft.gotMethod != wire.MethodTaskPanelClose {
		t.Fatalf("method = %q, want %q", ft.gotMethod, wire.MethodTaskPanelClose)
	}
}
