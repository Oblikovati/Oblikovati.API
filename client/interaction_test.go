// SPDX-License-Identifier: Apache-2.0

package client

import (
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
