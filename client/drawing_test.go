// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestAddSheetDecodesAndRoutes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"sheet":{"name":"Sheet:2","size":"a3","orientation":"landscape","widthMm":420,"heightMm":297,"active":true,"hasBorder":true,"hasTitleBlock":true}}`)}
	c := New(ft)

	r, err := c.Drawing().AddSheet(wire.AddSheetArgs{Size: "a3", Orientation: "landscape"})
	if err != nil {
		t.Fatalf("AddSheet: %v", err)
	}
	if ft.gotMethod != wire.MethodDrawingAddSheet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDrawingAddSheet)
	}
	if r.Sheet.Name != "Sheet:2" || r.Sheet.WidthMM != 420 || !r.Sheet.Active {
		t.Fatalf("decoded = %+v, want active A3 landscape 420 wide", r.Sheet)
	}
	var sent wire.AddSheetArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Size != "a3" || sent.Orientation != "landscape" {
		t.Errorf("sent = %+v, want a3 landscape", sent)
	}
}

// TestDrawingMethodsRoute guards that every typed drawing method targets its own wire
// constant (the wire-coverage invariant for the whole F01 surface).
func TestDrawingMethodsRoute(t *testing.T) {
	for _, tc := range []struct {
		method string
		call   func(c *Client) error
	}{
		{wire.MethodDrawingListSheets, func(c *Client) error { _, e := c.Drawing().ListSheets(); return e }},
		{wire.MethodDrawingAddSheet, func(c *Client) error { _, e := c.Drawing().AddSheet(wire.AddSheetArgs{}); return e }},
		{wire.MethodDrawingRemoveSheet, func(c *Client) error {
			_, e := c.Drawing().RemoveSheet(wire.RemoveSheetArgs{Name: "Sheet:1"})
			return e
		}},
		{wire.MethodDrawingSetActiveSheet, func(c *Client) error {
			_, e := c.Drawing().SetActiveSheet(wire.SetActiveSheetArgs{Name: "Sheet:1"})
			return e
		}},
		{wire.MethodDrawingSetModelReference, func(c *Client) error {
			_, e := c.Drawing().SetModelReference(wire.SetModelReferenceArgs{FullDocumentName: "p.opd"})
			return e
		}},
		{wire.MethodDrawingTitleBlockFields, func(c *Client) error {
			_, e := c.Drawing().TitleBlockFields(wire.TitleBlockFieldsArgs{})
			return e
		}},
		{wire.MethodDrawingExportDXF, func(c *Client) error {
			_, e := c.Drawing().ExportDXF(wire.ExportDrawingDXFArgs{Path: "sheet.dxf"})
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
