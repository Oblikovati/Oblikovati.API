// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestWorkSurfacesListSendsNilBodyAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"surfaces":[{"index":0,"name":"Surface1","ref":"surface/0","visible":true,"bodies":1,"source":"BoundaryPatch1"}]}`)}
	c := New(ft)

	res, err := c.WorkSurfaces().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodWorkSurfacesList || ft.gotReq != nil {
		t.Errorf("List sent method=%q body=%q, want %q / nil", ft.gotMethod, ft.gotReq, wire.MethodWorkSurfacesList)
	}
	if len(res.Surfaces) != 1 || res.Surfaces[0].Ref != "surface/0" || res.Surfaces[0].Bodies != 1 {
		t.Errorf("decoded = %+v, want one surface ref surface/0 with 1 body", res.Surfaces)
	}
}

func TestWorkSurfacesGetMarshalsIndex(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"surface":{"index":2,"name":"Surface3","ref":"surface/2","visible":true}}`)}
	c := New(ft)

	got, err := c.WorkSurfaces().Get(2)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if ft.gotMethod != wire.MethodWorkSurfacesGet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodWorkSurfacesGet)
	}
	var sent wire.WorkSurfaceRefArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Index != 2 {
		t.Errorf("sent = %+v, want index 2", sent)
	}
	if got.Surface.Index != 2 || got.Surface.Ref != "surface/2" {
		t.Errorf("decoded = %+v, want surface index 2", got.Surface)
	}
}

func TestWorkSurfacesSetVisibleMarshalsArgs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"surface":{"index":1,"name":"Surface2","ref":"surface/1","visible":false}}`)}
	c := New(ft)

	got, err := c.WorkSurfaces().SetVisible(1, false)
	if err != nil {
		t.Fatalf("SetVisible: %v", err)
	}
	if ft.gotMethod != wire.MethodWorkSurfacesSetVisible {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodWorkSurfacesSetVisible)
	}
	var sent wire.SetWorkSurfaceVisibleArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Index != 1 || sent.Visible {
		t.Errorf("sent = %+v, want index 1 visible=false", sent)
	}
	if got.Surface.Visible {
		t.Errorf("decoded = %+v, want hidden surface", got.Surface)
	}
}

func TestWorkSurfacesRenameMarshalsArgs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"surface":{"index":0,"name":"Parting Surface","ref":"surface/0","visible":true}}`)}
	c := New(ft)

	if _, err := c.WorkSurfaces().Rename(0, "Parting Surface"); err != nil {
		t.Fatalf("Rename: %v", err)
	}
	if ft.gotMethod != wire.MethodWorkSurfacesRename {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodWorkSurfacesRename)
	}
	var sent wire.RenameWorkSurfaceArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Index != 0 || sent.Name != "Parting Surface" {
		t.Errorf("sent = %+v, want index 0 name 'Parting Surface'", sent)
	}
}
