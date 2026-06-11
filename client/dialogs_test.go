// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestDialogsShowFileDialogMarshalsArgs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.Dialogs().ShowFileDialog(wire.ShowFileDialogArgs{
		ID: "sim.report", Save: true, Title: "Save report",
		Filter: "HTML (*.html)|*.html", MultiSelect: true,
	}); err != nil {
		t.Fatalf("ShowFileDialog: %v", err)
	}
	if ft.gotMethod != wire.MethodDialogsShowFileDialog {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDialogsShowFileDialog)
	}
	var sent wire.ShowFileDialogArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil ||
		!sent.Save || sent.ID != "sim.report" || !sent.MultiSelect {
		t.Errorf("sent = %s, want the save multi-select request", ft.gotReq)
	}
}

func TestDialogsWebViewLifecycle(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.Dialogs().ShowWebDialog(wire.WebDialogSpec{
		ID: "docs", Title: "Help", URL: "https://example.org", Dock: types.DockRight, Visible: true,
	}); err != nil {
		t.Fatalf("ShowWebDialog: %v", err)
	}
	var sent wire.ShowWebDialogArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Dialog.Dock != types.DockRight {
		t.Errorf("sent = %s, want a dock-right web view", ft.gotReq)
	}

	ft.reply = []byte(`{"views":[{"id":"docs","title":"Help","url":"https://example.org","visible":true}]}`)
	lst, err := c.Dialogs().ListWebViews()
	if err != nil || len(lst.Views) != 1 || lst.Views[0].URL != "https://example.org" {
		t.Fatalf("ListWebViews = (%+v, %v), want the docs view", lst, err)
	}

	ft.reply = []byte(`{"ok":true}`)
	if _, err := c.Dialogs().CloseWebDialog("docs"); err != nil {
		t.Fatalf("CloseWebDialog: %v", err)
	}
	if ft.gotMethod != wire.MethodDialogsCloseWebDialog {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDialogsCloseWebDialog)
	}
}
