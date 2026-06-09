// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestRibbonListSendsMethodAndDecodesReply(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"key":"Part","tabs":[{"name":"3D Model","panels":[{"name":"Create","controls":[{"commandId":"Create.Extrude","displayName":"Extrude"}]}]}]}`)}
	c := New(ft)

	got, err := c.Ribbon().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodRibbonList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodRibbonList)
	}
	if got.Key != types.PartRibbon {
		t.Errorf("key = %q, want Part", got.Key)
	}
	if len(got.Tabs) != 1 || got.Tabs[0].Name != "3D Model" {
		t.Fatalf("tabs = %+v, want one 3D Model tab", got.Tabs)
	}
	ctl := got.Tabs[0].Panels[0].Controls[0]
	if ctl.CommandID != "Create.Extrude" {
		t.Errorf("control = %+v, want Create.Extrude", ctl)
	}
}

func TestCommandsCreateCarriesRibbonAndEnvironment(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	if _, err := c.Commands().Create(wire.CreateCommandArgs{
		ID: "acme.slot", DisplayName: "Slot", Ribbon: types.PartRibbon,
		Tab: "Sketch", Category: "Draw", Environment: types.SketchEnvironment,
	}); err != nil {
		t.Fatalf("Create: %v", err)
	}
	if ft.gotMethod != wire.MethodCommandsCreate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodCommandsCreate)
	}
	// The ribbon/environment placement must reach the wire unchanged.
	if !contains(string(ft.gotReq), `"ribbon":"Part"`) || !contains(string(ft.gotReq), `"environment":1`) {
		t.Errorf("request did not carry ribbon/environment: %s", ft.gotReq)
	}
}
