// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestOptionsGeneralRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"group":"general","general":{"startupAction":1}}`)}
	c := New(ft)

	g, err := c.Options().General()
	if err != nil {
		t.Fatalf("General: %v", err)
	}
	if ft.gotMethod != wire.MethodOptionsGetGroup {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodOptionsGetGroup)
	}
	if g.StartupAction != types.StartupEmptyWorkspace {
		t.Errorf("startupAction = %v, want empty", g.StartupAction)
	}

	ft.reply = []byte(`{"ok":true}`)
	if _, err := c.Options().SetGeneral(wire.GeneralOptionsView{StartupAction: types.StartupNewPart}); err != nil {
		t.Fatalf("SetGeneral: %v", err)
	}
	var sent wire.OptionGroupView
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Group != wire.OptionGroupGeneral || sent.General == nil {
		t.Errorf("sent = %+v, want a general-group payload", sent)
	}
}

func TestOptionsMissingPayloadIsAnError(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"group":"sketch"}`)}
	if _, err := New(ft).Options().Sketch(); err == nil {
		t.Fatal("a group reply without its payload should error, not zero-value")
	}
}

func TestOptionsSketchAndPartCarryValues(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"group":"sketch","sketch":{"gridSpacingCm":0.5,"gridVisible":true,"gridMajorEvery":4,"snapToPoints":true,"snapToGrid":false}}`)}
	c := New(ft)
	sk, err := c.Options().Sketch()
	if err != nil || sk.GridSpacingCm != 0.5 || sk.GridMajorEvery != 4 || sk.SnapToGrid {
		t.Fatalf("Sketch = (%+v, %v), want 0.5cm grid, major 4, no grid snap", sk, err)
	}

	ft.reply = []byte(`{"ok":true}`)
	if _, err := c.Options().SetPart(wire.PartOptionsView{ChamferFlatCorners: false}); err != nil {
		t.Fatalf("SetPart: %v", err)
	}
	var sent wire.OptionGroupView
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Part == nil || sent.Part.ChamferFlatCorners {
		t.Errorf("sent = %+v, want part payload with flat corners off", sent)
	}
}
