// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestWindowsFramesAndTabsDecode(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"frames":[{"caption":"bracket.obk — Oblikovati","state":1,"width":2560,"height":1480}]}`)}
	c := New(ft)
	frames, err := c.Windows().Frames()
	if err != nil || len(frames.Frames) != 1 || frames.Frames[0].State != types.WindowMaximized {
		t.Fatalf("Frames = (%+v, %v), want one maximized frame", frames, err)
	}

	ft.reply = []byte(`{"tabs":[{"document":3,"title":"bracket.obk","active":true,"dirty":true}]}`)
	tabs, err := c.Windows().Tabs()
	if err != nil || len(tabs.Tabs) != 1 || !tabs.Tabs[0].Active || !tabs.Tabs[0].Dirty {
		t.Fatalf("Tabs = (%+v, %v), want the active dirty tab", tabs, err)
	}
}

func TestWindowsTabVerbsMarshal(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.Windows().ActivateTab(3); err != nil {
		t.Fatalf("ActivateTab: %v", err)
	}
	if ft.gotMethod != wire.MethodWindowsActivateTab {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodWindowsActivateTab)
	}
	if _, err := c.Windows().CloseTab(3, true); err != nil {
		t.Fatalf("CloseTab: %v", err)
	}
	var sent wire.CloseViewTabArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Document != 3 || !sent.Force {
		t.Errorf("CloseTab sent %s, want document 3 forced", ft.gotReq)
	}
}
