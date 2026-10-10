// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestStatusSetAndGetText(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.Status().SetText("Meshing…"); err != nil {
		t.Fatalf("SetText: %v", err)
	}
	if ft.gotMethod != wire.MethodStatusSetText {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodStatusSetText)
	}
	ft.reply = []byte(`{"text":"Meshing…"}`)
	r, err := c.Status().Text()
	if err != nil || r.Text != "Meshing…" {
		t.Fatalf("Text = (%+v, %v), want the set message", r, err)
	}
}

func TestProgressLifecycle(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":7}`)}
	c := New(ft)
	bar, err := c.Progress().Begin(100, "Meshing…")
	if err != nil || bar.ID != 7 {
		t.Fatalf("Begin = (%+v, %v), want id 7", bar, err)
	}

	ft.reply = []byte(`{"ok":true,"cancelled":true}`)
	upd, err := c.Progress().Update(7, 42, "halfway")
	if err != nil || !upd.Cancelled {
		t.Fatalf("Update = (%+v, %v), want cancelled reported", upd, err)
	}
	var sent wire.UpdateProgressArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.ID != 7 || sent.Step != 42 {
		t.Errorf("Update sent %s, want id 7 step 42", ft.gotReq)
	}

	ft.reply = []byte(`{"ok":true}`)
	if _, err := c.Progress().End(7); err != nil {
		t.Fatalf("End: %v", err)
	}
	if ft.gotMethod != wire.MethodProgressEnd {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodProgressEnd)
	}
}

func TestMessagesBalloonAndPrompt(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.Messages().RegisterBalloonTip(wire.RegisterBalloonTipArgs{
		ID: "sim.done", Title: "Done", Text: "Study finished",
	}); err != nil {
		t.Fatalf("RegisterBalloonTip: %v", err)
	}

	ft.reply = []byte(`{"shown":false}`)
	shown, err := c.Messages().ShowBalloonTip("sim.done")
	if err != nil || shown.Shown {
		t.Fatalf("ShowBalloonTip = (%+v, %v), want suppressed", shown, err)
	}

	ft.reply = []byte(`{"resolved":true,"answer":"Replace"}`)
	res, err := c.Messages().ShowPrompt(wire.ShowPromptArgs{
		ID: "sim.replace", Message: "Replace results?", Buttons: []string{"Replace", "Keep"},
		Restriction: types.PromptAllowRemember,
	})
	if err != nil || !res.Resolved || res.Answer != "Replace" {
		t.Fatalf("ShowPrompt = (%+v, %v), want the remembered Replace", res, err)
	}
	var sentPrompt wire.ShowPromptArgs
	if err := json.Unmarshal(ft.gotReq, &sentPrompt); err != nil ||
		sentPrompt.Restriction != types.PromptAllowRemember || len(sentPrompt.Buttons) != 2 {
		t.Errorf("prompt sent %s, want allow-remember with two buttons", ft.gotReq)
	}
}

func TestMessagesSectionsAndList(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"section":3}`)}
	c := New(ft)
	sec, err := c.Messages().BeginSection("Meshing")
	if err != nil || sec.Section != 3 {
		t.Fatalf("BeginSection = (%+v, %v), want section 3", sec, err)
	}

	ft.reply = []byte(`{"ok":true}`)
	if _, err := c.Messages().AddMessage(wire.AddErrorMessageArgs{
		Text: "degenerate face", Severity: types.SeverityWarning,
	}); err != nil {
		t.Fatalf("AddMessage: %v", err)
	}
	if _, err := c.Messages().EndSection(3); err != nil {
		t.Fatalf("EndSection: %v", err)
	}

	ft.reply = []byte(`{"root":{"sections":[{"title":"Meshing","messages":[{"text":"degenerate face","severity":1}]}]},"hasErrors":false,"hasWarnings":true,"lastMessage":"degenerate face"}`)
	lst, err := c.Messages().List()
	if err != nil || !lst.HasWarnings || lst.HasErrors {
		t.Fatalf("List = (%+v, %v), want warnings only", lst, err)
	}
	if len(lst.Root.Sections) != 1 || lst.Root.Sections[0].Messages[0].Severity != types.SeverityWarning {
		t.Errorf("tree = %+v, want the Meshing section's warning", lst.Root)
	}
}
