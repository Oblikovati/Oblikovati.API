// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati/api/types"
	"oblikovati/api/wire"
)

func TestViewsListSendsDocumentAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"views":[{"index":0,"name":"View 1","active":true,"camera":{"eye":[0,0,10],"target":[0,0,0],"up":[0,1,0],"fov":0.8}}],"activeIndex":0,"layout":4}`)}
	c := New(ft)

	got, err := c.Views().List(7)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodViewsList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodViewsList)
	}
	var sent wire.ListViewsArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Document != 7 {
		t.Errorf("sent document = %d, want 7", sent.Document)
	}
	if len(got.Views) != 1 || got.Views[0].Name != "View 1" || got.Layout != types.LayoutFour {
		t.Errorf("decoded = %+v, want one view 'View 1' layout=four", got)
	}
}

func TestViewsAddMarshalsArgsAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":1,"name":"Iso","active":true,"camera":{"eye":[1,1,1],"target":[0,0,0],"up":[0,1,0],"fov":0.8}}`)}
	c := New(ft)

	got, err := c.Views().Add(wire.AddViewArgs{Name: "Iso", CopyActiveCamera: true})
	if err != nil {
		t.Fatalf("Add: %v", err)
	}
	if ft.gotMethod != wire.MethodViewsAdd {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodViewsAdd)
	}
	var sent wire.AddViewArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Name != "Iso" || !sent.CopyActiveCamera {
		t.Errorf("sent = %+v, want name=Iso copyActiveCamera=true", sent)
	}
	if got.Index != 1 || got.Name != "Iso" {
		t.Errorf("decoded = %+v, want index=1 name=Iso", got)
	}
}

func TestViewsSetLayoutMarshalsLayout(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"layout":4}`)}
	c := New(ft)

	got, err := c.Views().SetLayout(wire.SetLayoutArgs{Layout: types.LayoutFour})
	if err != nil {
		t.Fatalf("SetLayout: %v", err)
	}
	if ft.gotMethod != wire.MethodViewsSetLayout {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodViewsSetLayout)
	}
	var sent wire.SetLayoutArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Layout != types.LayoutFour {
		t.Errorf("sent layout = %v, want LayoutFour", sent.Layout)
	}
	if got.Layout != types.LayoutFour {
		t.Errorf("decoded layout = %v, want LayoutFour", got.Layout)
	}
}

func TestViewSetCameraCarriesViewAddressing(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"eye":[2,2,2],"target":[0,0,0],"up":[0,1,0],"fov":0.8}`)}
	c := New(ft)

	if _, err := c.View().SetCamera(wire.SetCameraArgs{
		Document: 3, View: 2, Eye: [3]float64{2, 2, 2}, Target: [3]float64{0, 0, 0}, Up: [3]float64{0, 1, 0}, FOV: 0.8,
	}); err != nil {
		t.Fatalf("SetCamera: %v", err)
	}
	var sent wire.SetCameraArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Document != 3 || sent.View != 2 {
		t.Errorf("sent addressing = doc %d view %d, want doc 3 view 2", sent.Document, sent.View)
	}
}
