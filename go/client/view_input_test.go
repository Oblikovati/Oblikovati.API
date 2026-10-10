// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Synthesised input is the only way a client reaches behaviour that lives in the INPUT path —
// which constraints a tool infers from where a click landed, what it previews between clicks, how
// a multi-click chain builds. These pin that each method sends the method constant and request the
// host expects.

// TestClickSendsAModelPoint: the common form — a model-space point the host projects, so the
// caller needs no knowledge of the camera.
func TestClickSendsAModelPoint(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"x":120,"y":80,"activeTool":"Rectangle"}`)}
	c := New(ft)
	at := types.NewPoint(6, 0, 0)

	got, err := c.View().Click(wire.ClickViewportArgs{Point: &at})
	if err != nil {
		t.Fatalf("Click: %v", err)
	}

	if ft.gotMethod != wire.MethodViewportClick {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodViewportClick)
	}
	var sent wire.ClickViewportArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Point == nil || *sent.Point != at {
		t.Errorf("sent point %v, want %v", sent.Point, at)
	}
	if got.ActiveTool != "Rectangle" || got.X != 120 {
		t.Errorf("decoded %+v, want the pixel clicked and the running command", got)
	}
}

// TestClickSendsPixelsAndModifiers: the pixel form with held modifiers, which is how a caller
// extends a selection.
func TestClickSendsPixelsAndModifiers(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"x":40,"y":40,"activeTool":""}`)}
	c := New(ft)

	if _, err := c.View().Click(wire.ClickViewportArgs{X: 40, Y: 40, Button: "right", Shift: true}); err != nil {
		t.Fatalf("Click: %v", err)
	}

	var sent wire.ClickViewportArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.X != 40 || sent.Y != 40 || sent.Button != "right" || !sent.Shift {
		t.Errorf("sent %+v, want the pixel, button and modifier as given", sent)
	}
	if sent.Point != nil {
		t.Error("a pixel click must not carry a model point")
	}
}

// TestPressKeySendsTheKeyName: Escape and Enter are how a variable-length command (a continuous
// line chain, a spline) is finished, so the key name has to reach the host verbatim.
func TestPressKeySendsTheKeyName(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"activeTool":""}`)}
	c := New(ft)

	got, err := c.View().PressKey(wire.PressKeyArgs{Key: "Escape"})
	if err != nil {
		t.Fatalf("PressKey: %v", err)
	}

	if ft.gotMethod != wire.MethodViewportKey {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodViewportKey)
	}
	var sent wire.PressKeyArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Key != "Escape" {
		t.Errorf("sent key %q, want %q", sent.Key, "Escape")
	}
	if got.ActiveTool != "" {
		t.Errorf("ActiveTool = %q, want empty once the key ended the command", got.ActiveTool)
	}
}
