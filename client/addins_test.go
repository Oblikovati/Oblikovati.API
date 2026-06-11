// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestAddInsListDecodesRegistry(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"addIns":[{"id":"com.x.a","activated":true,"loadBehavior":1,"hasAutomation":true}]}`)}
	c := New(ft)

	r, err := c.AddIns().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodAddInsList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAddInsList)
	}
	if len(r.AddIns) != 1 || r.AddIns[0].ID != "com.x.a" || !r.AddIns[0].Activated {
		t.Fatalf("decoded = %+v, want one activated com.x.a entry", r.AddIns)
	}
	if r.AddIns[0].LoadBehavior != types.LoadOnDemand || !r.AddIns[0].HasAutomation {
		t.Errorf("entry = %+v, want loadBehavior=demand hasAutomation=true", r.AddIns[0])
	}
}

func TestAddInsSetLoadBehaviorMarshalsArgs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	if _, err := c.AddIns().SetLoadBehavior("com.x.a", types.LoadDisabled); err != nil {
		t.Fatalf("SetLoadBehavior: %v", err)
	}
	if ft.gotMethod != wire.MethodAddInsSetLoadBehavior {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAddInsSetLoadBehavior)
	}
	var sent wire.SetAddInLoadBehaviorArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ID != "com.x.a" || sent.LoadBehavior != types.LoadDisabled {
		t.Errorf("sent = %+v, want id=com.x.a loadBehavior=disabled", sent)
	}
}

func TestAddInsLifecycleAndGetUseRefArgs(t *testing.T) {
	for _, tc := range []struct {
		method string
		call   func(c *Client) error
	}{
		{wire.MethodAddInsGet, func(c *Client) error { _, err := c.AddIns().Get("com.x.a"); return err }},
		{wire.MethodAddInsActivate, func(c *Client) error { _, err := c.AddIns().Activate("com.x.a"); return err }},
		{wire.MethodAddInsDeactivate, func(c *Client) error { _, err := c.AddIns().Deactivate("com.x.a"); return err }},
	} {
		ft := &fakeTransport{reply: []byte(`{"ok":true,"id":"com.x.a"}`)}
		if err := tc.call(New(ft)); err != nil {
			t.Fatalf("%s: %v", tc.method, err)
		}
		if ft.gotMethod != tc.method {
			t.Errorf("method = %q, want %q", ft.gotMethod, tc.method)
		}
		var sent wire.AddInRefArgs
		if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.ID != "com.x.a" {
			t.Errorf("%s sent %s, want {\"id\":\"com.x.a\"}", tc.method, ft.gotReq)
		}
	}
}

func TestAddInsCallAutomationPassesPayloadThrough(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"result":{"sum":7}}`)}
	c := New(ft)

	out, err := c.AddIns().CallAutomation("com.x.calc", "add", json.RawMessage(`{"a":3,"b":4}`))
	if err != nil {
		t.Fatalf("CallAutomation: %v", err)
	}
	if ft.gotMethod != wire.MethodAddInsCallAutomation {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAddInsCallAutomation)
	}
	var sent wire.CallAddInAutomationArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ID != "com.x.calc" || sent.Method != "add" || string(sent.Args) != `{"a":3,"b":4}` {
		t.Errorf("sent = %+v, want target com.x.calc method add args {\"a\":3,\"b\":4}", sent)
	}
	if string(out) != `{"sum":7}` {
		t.Errorf("result = %s, want {\"sum\":7}", out)
	}
}

func TestClientApplicationsRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":2}`)}
	c := New(ft)

	reg, err := c.ClientApplications().Register("acme-pipeline")
	if err != nil {
		t.Fatalf("Register: %v", err)
	}
	if ft.gotMethod != wire.MethodClientAppsRegister || reg.ID != 2 {
		t.Errorf("method=%q id=%d, want %q id=2", ft.gotMethod, reg.ID, wire.MethodClientAppsRegister)
	}

	ft.reply = []byte(`{"clients":[{"id":2,"name":"acme-pipeline"}]}`)
	lst, err := c.ClientApplications().List()
	if err != nil || len(lst.Clients) != 1 || lst.Clients[0].Name != "acme-pipeline" {
		t.Fatalf("List = (%+v, %v), want one acme-pipeline client", lst, err)
	}

	ft.reply = []byte(`{"ok":true}`)
	if _, err := c.ClientApplications().Unregister(2); err != nil {
		t.Fatalf("Unregister: %v", err)
	}
	var sent wire.UnregisterClientApplicationArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.ID != 2 {
		t.Errorf("Unregister sent %s, want {\"id\":2}", ft.gotReq)
	}
}
