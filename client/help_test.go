// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestHelpRegisterAndDisplay(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.Help().RegisterContext("com.x.sim", "https://docs.example.org/sim/"); err != nil {
		t.Fatalf("RegisterContext: %v", err)
	}
	var sent wire.RegisterHelpContextArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Source != "com.x.sim" {
		t.Errorf("sent = %s, want the sim source", ft.gotReq)
	}
	if _, err := c.Help().Display("com.x.sim", "mesh-setup"); err != nil {
		t.Fatalf("Display: %v", err)
	}
	if ft.gotMethod != wire.MethodHelpDisplay {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodHelpDisplay)
	}

	ft.reply = []byte(`{"source":"com.x.sim","base":"https://docs.example.org/sim/"}`)
	p, err := c.Help().Path("com.x.sim")
	if err != nil || p.Base == "" {
		t.Fatalf("Path = (%+v, %v), want the base", p, err)
	}

	ft.reply = []byte(`{"locale":"en-US"}`)
	li, err := c.Help().LanguageInfo()
	if err != nil || li.Locale != "en-US" {
		t.Fatalf("LanguageInfo = (%+v, %v), want en-US", li, err)
	}
}
