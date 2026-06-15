// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestDesignViewCaptureDecodesAndRoutes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"representation":{"id":3,"name":"Exploded","kind":"designView","active":true,"hiddenCount":2}}`)}
	c := New(ft)

	r, err := c.DesignReps().Capture("Exploded")
	if err != nil {
		t.Fatalf("Capture: %v", err)
	}
	if ft.gotMethod != wire.MethodDesignRepsCapture {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDesignRepsCapture)
	}
	if r.Representation.Name != "Exploded" || r.Representation.HiddenCount != 2 {
		t.Fatalf("decoded = %+v, want Exploded with 2 hidden", r.Representation)
	}
	var sent wire.CaptureRepArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Name != "Exploded" {
		t.Errorf("sent name = %q, want Exploded", sent.Name)
	}
}

// TestRepresentationMethodsRoute guards that every typed representation method targets its own
// wire constant (the wire-coverage invariant for the whole F04 surface).
func TestRepresentationMethodsRoute(t *testing.T) {
	for _, tc := range []struct {
		method string
		call   func(c *Client) error
	}{
		{wire.MethodDesignRepsActivate, func(c *Client) error { _, e := c.DesignReps().Activate(1); return e }},
		{wire.MethodDesignRepsList, func(c *Client) error { _, e := c.DesignReps().List(); return e }},
		{wire.MethodDesignRepsDelete, func(c *Client) error { _, e := c.DesignReps().Delete(1); return e }},
		{wire.MethodDesignRepsSetVisibility, func(c *Client) error {
			_, e := c.DesignReps().SetVisibility(wire.SetVisibilityArgs{Rep: 1, Occurrence: 2})
			return e
		}},
		{wire.MethodDesignRepsSetAppearance, func(c *Client) error {
			_, e := c.DesignReps().SetAppearance(wire.SetAppearanceArgs{Rep: 1, Occurrence: 2})
			return e
		}},
		{wire.MethodDesignRepsAddSection, func(c *Client) error {
			_, e := c.DesignReps().AddSection(wire.AddSectionArgs{Rep: 1})
			return e
		}},
		{wire.MethodPositionalRepsCapture, func(c *Client) error { _, e := c.PositionalReps().Capture("p"); return e }},
		{wire.MethodPositionalRepsActivate, func(c *Client) error { _, e := c.PositionalReps().Activate(1); return e }},
		{wire.MethodPositionalRepsList, func(c *Client) error { _, e := c.PositionalReps().List(); return e }},
		{wire.MethodPositionalRepsDelete, func(c *Client) error { _, e := c.PositionalReps().Delete(1); return e }},
		{wire.MethodPositionalRepsSetOverride, func(c *Client) error {
			_, e := c.PositionalReps().SetOverride(wire.SetPositionalOverrideArgs{Rep: 1, Relationship: 2, Value: 0.5})
			return e
		}},
		{wire.MethodPositionalRepsSetFlexible, func(c *Client) error {
			_, e := c.PositionalReps().SetFlexible(wire.SetFlexibleArgs{Rep: 1, Occurrence: 2, Flexible: true})
			return e
		}},
		{wire.MethodLODRepsCapture, func(c *Client) error { _, e := c.LODReps().Capture("l"); return e }},
		{wire.MethodLODRepsActivate, func(c *Client) error { _, e := c.LODReps().Activate(1); return e }},
		{wire.MethodLODRepsList, func(c *Client) error { _, e := c.LODReps().List(); return e }},
		{wire.MethodLODRepsDelete, func(c *Client) error { _, e := c.LODReps().Delete(1); return e }},
		{wire.MethodLODRepsSetSuppressed, func(c *Client) error {
			_, e := c.LODReps().SetSuppressed(wire.SetSuppressedArgs{Rep: 1, Occurrence: 2, Suppressed: true})
			return e
		}},
		{wire.MethodModelStatesCreate, func(c *Client) error {
			_, e := c.ModelStates().Create(wire.CreateModelStateArgs{Name: "m"})
			return e
		}},
		{wire.MethodModelStatesActivate, func(c *Client) error { _, e := c.ModelStates().Activate(1); return e }},
		{wire.MethodModelStatesList, func(c *Client) error { _, e := c.ModelStates().List(); return e }},
		{wire.MethodModelStatesDelete, func(c *Client) error { _, e := c.ModelStates().Delete(1); return e }},
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
