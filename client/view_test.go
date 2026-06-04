// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"github.com/Oblikovati/api/types"
	"github.com/Oblikovati/api/wire"
)

// TestViewSetDisplayModeMarshalsRequestAndDecodesReply checks the View client group sends the
// right method + body and decodes the reply into typed values.
func TestViewSetDisplayModeMarshalsRequestAndDecodesReply(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"mode":8709,"name":"Realistic"}`)}
	c := New(ft)

	got, err := c.View().SetDisplayMode(types.RealisticRendering)
	if err != nil {
		t.Fatalf("SetDisplayMode: %v", err)
	}
	if ft.gotMethod != wire.MethodViewSetDisplayMode {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodViewSetDisplayMode)
	}
	var sent wire.SetDisplayModeArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Mode != types.RealisticRendering {
		t.Errorf("sent mode = %d, want %d", sent.Mode, types.RealisticRendering)
	}
	if got.Mode != types.RealisticRendering || got.Name != "Realistic" {
		t.Errorf("decoded = %+v, want Realistic", got)
	}
}

// TestViewDisplayModeSendsNilBody checks the no-arg getter sends the right method with no body.
func TestViewDisplayModeSendsNilBody(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"mode":8710,"name":"Shaded with Edges"}`)}
	c := New(ft)

	if _, err := c.View().DisplayMode(); err != nil {
		t.Fatalf("DisplayMode: %v", err)
	}
	if ft.gotMethod != wire.MethodViewGetDisplayMode {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodViewGetDisplayMode)
	}
	if ft.gotReq != nil {
		t.Errorf("no-arg getter sent body %q, want nil", ft.gotReq)
	}
}
