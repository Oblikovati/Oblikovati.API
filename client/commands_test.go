// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati/api/wire"
)

func TestCommandsSetStateMarshalsArgs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	if _, err := c.Commands().SetState(wire.SetCommandStateArgs{
		ID: "x.toggle", Active: true, DisplayName: "Presenting",
	}); err != nil {
		t.Fatalf("SetState: %v", err)
	}
	if ft.gotMethod != wire.MethodCommandsSetState {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodCommandsSetState)
	}
	var sent wire.SetCommandStateArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ID != "x.toggle" || !sent.Active || sent.DisplayName != "Presenting" {
		t.Errorf("sent = %+v, want id=x.toggle active=true displayName=Presenting", sent)
	}
}
