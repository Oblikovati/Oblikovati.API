// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestInterferenceAnalyzeDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"results":[{"occurrenceA":1,"occurrenceB":2,"volume":0.5,"center":[0,0,1]}],"totalVolume":0.5}`)}
	c := New(ft)

	r, err := c.Interference().Analyze(wire.AnalyzeInterferenceArgs{})
	if err != nil {
		t.Fatalf("Analyze: %v", err)
	}
	if ft.gotMethod != wire.MethodInterferenceAnalyze {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodInterferenceAnalyze)
	}
	if len(r.Results) != 1 || r.Results[0].Volume != 0.5 || r.TotalVolume != 0.5 {
		t.Fatalf("decoded = %+v, want one 0.5-volume overlap", r)
	}
}

// TestContactAndFlexibleMethodsRoute guards the wire-coverage invariant for the M12-F05 contact
// surface and the M12-F06 flexible flag.
func TestContactAndFlexibleMethodsRoute(t *testing.T) {
	for _, tc := range []struct {
		method string
		call   func(c *Client) error
	}{
		{wire.MethodContactSetsCreate, func(c *Client) error { _, e := c.ContactSets().Create("s"); return e }},
		{wire.MethodContactSetsList, func(c *Client) error { _, e := c.ContactSets().List(); return e }},
		{wire.MethodContactSetsDelete, func(c *Client) error { _, e := c.ContactSets().Delete(1); return e }},
		{wire.MethodContactSetsAddMember, func(c *Client) error {
			_, e := c.ContactSets().AddMember(wire.ContactMemberArgs{Set: 1, Occurrence: 2})
			return e
		}},
		{wire.MethodContactSetsRemoveMember, func(c *Client) error {
			_, e := c.ContactSets().RemoveMember(wire.ContactMemberArgs{Set: 1, Occurrence: 2})
			return e
		}},
		{wire.MethodContactSolverSetEnabled, func(c *Client) error { _, e := c.ContactSolver().SetEnabled(true); return e }},
		{wire.MethodContactSolverStatus, func(c *Client) error { _, e := c.ContactSolver().Status(); return e }},
		{wire.MethodInterferenceAnalyze, func(c *Client) error {
			_, e := c.Interference().Analyze(wire.AnalyzeInterferenceArgs{})
			return e
		}},
		{wire.MethodAssemblySetFlexible, func(c *Client) error { _, e := c.Assembly().SetFlexible(1, true); return e }},
	} {
		ft := &fakeTransport{reply: []byte(`{}`)}
		if err := tc.call(New(ft)); err != nil {
			t.Fatalf("%s: %v", tc.method, err)
		}
		if ft.gotMethod != tc.method {
			t.Errorf("method = %q, want %q", ft.gotMethod, tc.method)
		}
	}
}

// TestSetFlexibleMarshalsArgs checks the flexible flag round-trips.
func TestSetFlexibleMarshalsArgs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"occurrence":{"id":3,"name":"sub:1","flexible":true}}`)}
	c := New(ft)

	r, err := c.Assembly().SetFlexible(3, true)
	if err != nil {
		t.Fatalf("SetFlexible: %v", err)
	}
	if !r.Occurrence.Flexible {
		t.Errorf("decoded occurrence not flexible: %+v", r.Occurrence)
	}
	var sent wire.SetFlexibleOccurrenceArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ID != 3 || !sent.Flexible {
		t.Errorf("sent = %+v, want id 3 flexible", sent)
	}
}
