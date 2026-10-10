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

// TestOptionsMissingPayloadIsAnError guards all 5 group getters, which share
// optionGroupField: a reply missing its group's payload must error with the
// exact "no <group> payload" wording, not return a zero value silently.
func TestOptionsMissingPayloadIsAnError(t *testing.T) {
	cases := []struct {
		group string
		get   func(Options) error
	}{
		{"general", func(o Options) error { _, err := o.General(); return err }},
		{"display", func(o Options) error { _, err := o.Display(); return err }},
		{"sketch", func(o Options) error { _, err := o.Sketch(); return err }},
		{"part", func(o Options) error { _, err := o.Part(); return err }},
		{"save", func(o Options) error { _, err := o.Save(); return err }},
	}
	for _, c := range cases {
		ft := &fakeTransport{reply: []byte(`{"group":"` + c.group + `"}`)}
		err := c.get(New(ft).Options())
		if err == nil {
			t.Errorf("%s: a group reply without its payload should error, not zero-value", c.group)
			continue
		}
		want := `client: options.getGroup("` + c.group + `") reply carries no ` + c.group + ` payload`
		if err.Error() != want {
			t.Errorf("%s: error = %q, want %q", c.group, err, want)
		}
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
