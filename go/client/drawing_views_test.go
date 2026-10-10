// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestAddBaseViewDecodesAndRoutes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"view":{"name":"VIEW1","projected":false,"orientation":"front","scale":1,"style":"hiddenLine","centerXmm":150,"centerYmm":100,"visibleCount":12,"hiddenCount":4}}`)}
	c := New(ft)

	r, err := c.DrawingViews().AddBase(wire.AddBaseViewArgs{Orientation: "front", Scale: 1})
	if err != nil {
		t.Fatalf("AddBase: %v", err)
	}
	if ft.gotMethod != wire.MethodDrawingViewsAddBase {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDrawingViewsAddBase)
	}
	if r.View.Name != "VIEW1" || r.View.VisibleCount != 12 || r.View.HiddenCount != 4 {
		t.Fatalf("decoded = %+v, want VIEW1 12/4", r.View)
	}
	var sent wire.AddBaseViewArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Orientation != "front" {
		t.Errorf("sent orientation = %q, want front", sent.Orientation)
	}
}

func TestDrawingViewsMethodsRoute(t *testing.T) {
	for _, tc := range []struct {
		method string
		call   func(c *Client) error
	}{
		{wire.MethodDrawingViewsList, func(c *Client) error { _, e := c.DrawingViews().List(); return e }},
		{wire.MethodDrawingViewsAddBase, func(c *Client) error {
			_, e := c.DrawingViews().AddBase(wire.AddBaseViewArgs{})
			return e
		}},
		{wire.MethodDrawingViewsAddProjected, func(c *Client) error {
			_, e := c.DrawingViews().AddProjected(wire.AddProjectedViewArgs{BaseView: "VIEW1", Direction: "right"})
			return e
		}},
		{wire.MethodDrawingViewsDelete, func(c *Client) error {
			_, e := c.DrawingViews().Delete(wire.DeleteViewArgs{Name: "VIEW1"})
			return e
		}},
		{wire.MethodDrawingViewsCurves, func(c *Client) error {
			_, e := c.DrawingViews().Curves(wire.ViewCurvesArgs{View: "VIEW1"})
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
