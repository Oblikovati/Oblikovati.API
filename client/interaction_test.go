// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati/api/wire"
)

func TestInteractionStateSendsNilBodyAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"busy":true,"activeTool":"ExtrudeTool","inTransaction":false}`)}
	c := New(ft)

	st, err := c.Interaction().State()
	if err != nil {
		t.Fatalf("State: %v", err)
	}
	if ft.gotMethod != wire.MethodInteractionState {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodInteractionState)
	}
	if ft.gotReq != nil {
		t.Errorf("no-arg query should send nil body, got %q", ft.gotReq)
	}
	if !st.Busy || st.ActiveTool != "ExtrudeTool" {
		t.Errorf("decoded = %+v, want busy=true activeTool=ExtrudeTool", st)
	}
}

func TestInteractionSetNoticeMarshalsMessage(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	if _, err := c.Interaction().SetNotice("Meeting: connected"); err != nil {
		t.Fatalf("SetNotice: %v", err)
	}
	if ft.gotMethod != wire.MethodInteractionSetNotice {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodInteractionSetNotice)
	}
	var sent wire.SetNoticeArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Message != "Meeting: connected" {
		t.Errorf("sent message = %q, want %q", sent.Message, "Meeting: connected")
	}
}
