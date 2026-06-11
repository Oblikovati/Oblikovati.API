// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestTriadShowForcesVisible(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.Triad().Show(wire.TriadSpec{
		Position: [3]float64{1, 2, 3},
		Allowed:  []types.TriadSegment{types.TriadXAxis, types.TriadZRing},
	}); err != nil {
		t.Fatalf("Show: %v", err)
	}
	var sent wire.ShowTriadArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if !sent.Triad.Visible || sent.Triad.Position[2] != 3 || len(sent.Triad.Allowed) != 2 {
		t.Errorf("sent = %+v, want a visible constrained triad", sent.Triad)
	}

	ft.reply = []byte(`{"position":[1,2,3],"visible":true}`)
	got, err := c.Triad().Get()
	if err != nil || !got.Visible || got.Position[0] != 1 {
		t.Fatalf("Get = (%+v, %v), want the placed triad", got, err)
	}
}

func TestManipulatorsSetAndRemove(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.Manipulators().Set("sim.handles", []wire.ManipulatorHandleSpec{
		{ID: "tip", Position: [3]float64{0, 0, 5}, RadiusPx: 10},
	}); err != nil {
		t.Fatalf("Set: %v", err)
	}
	var sent wire.SetManipulatorsArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil ||
		sent.ID != "sim.handles" || sent.Handles[0].ID != "tip" {
		t.Errorf("sent = %s, want the tip handle", ft.gotReq)
	}
	if _, err := c.Manipulators().Remove("sim.handles"); err != nil {
		t.Fatalf("Remove: %v", err)
	}
	if ft.gotMethod != wire.MethodManipulatorsRemove {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodManipulatorsRemove)
	}
}
