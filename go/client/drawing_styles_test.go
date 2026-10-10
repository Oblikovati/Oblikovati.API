// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestSetStandardDecodesAndRoutes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"style":{"standard":"ansi","dimension":{"name":"Default (ANSI)","textHeightMm":3,"arrowSizeMm":3,"decimalPlaces":3,"unit":"in","lineWeightMm":0.3},"text":{"name":"Note","fontName":"Arial","heightMm":3},"line":{"name":"Visible","weightMm":0.5}}}`)}
	c := New(ft)

	r, err := c.DrawingStyles().SetStandard(wire.SetStandardArgs{Standard: "ansi"})
	if err != nil {
		t.Fatalf("SetStandard: %v", err)
	}
	if ft.gotMethod != wire.MethodDrawingStylesSetStandard {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDrawingStylesSetStandard)
	}
	if r.Style.Standard != "ansi" || r.Style.Dimension.Unit != "in" || r.Style.Dimension.DecimalPlaces != 3 {
		t.Fatalf("decoded = %+v, want ANSI/in/3dp", r.Style)
	}
	var sent wire.SetStandardArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Standard != "ansi" {
		t.Errorf("sent standard = %q, want ansi", sent.Standard)
	}
}

// TestDrawingStylesMethodsRoute guards typed coverage of every drawingStyles wire method.
func TestDrawingStylesMethodsRoute(t *testing.T) {
	for _, tc := range []struct {
		method string
		call   func(c *Client) error
	}{
		{wire.MethodDrawingStylesListStandards, func(c *Client) error { _, e := c.DrawingStyles().ListStandards(); return e }},
		{wire.MethodDrawingStylesGetActiveStyle, func(c *Client) error { _, e := c.DrawingStyles().GetActiveStyle(); return e }},
		{wire.MethodDrawingStylesSetStandard, func(c *Client) error {
			_, e := c.DrawingStyles().SetStandard(wire.SetStandardArgs{Standard: "iso"})
			return e
		}},
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
