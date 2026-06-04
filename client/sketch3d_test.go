// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"github.com/Oblikovati/api/wire"
)

func TestSketch3DCreateSendsNameAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"sketchIndex":0}`)}
	c := New(ft)

	got, err := c.Sketch3D().Create(wire.CreateSketch3DArgs{Name: "Path"})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DCreate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DCreate)
	}
	var sent wire.CreateSketch3DArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Name != "Path" {
		t.Errorf("sent name = %q, want Path", sent.Name)
	}
	if got.SketchIndex != 0 {
		t.Errorf("decoded index = %d, want 0", got.SketchIndex)
	}
}

func TestSketch3DGetSendsIndexAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":1,"name":"3D Sketch2","visible":true,"dimensionsVisible":true,"entityCount":3,"dof":9,"healthy":true}`)}
	c := New(ft)

	got, err := c.Sketch3D().Get(1)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DGet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DGet)
	}
	var sent wire.Sketch3DArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.SketchIndex != 1 {
		t.Errorf("sent index = %d, want 1", sent.SketchIndex)
	}
	if got.Name != "3D Sketch2" || got.EntityCount != 3 || got.DOF != 9 {
		t.Errorf("decoded = %+v, want 3D Sketch2 / 3 entities / DOF 9", got)
	}
}

func TestSketch3DSetDimensionsVisibleSendsProperty(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"name":"3D Sketch1","dimensionsVisible":false}`)}
	c := New(ft)

	got, err := c.Sketch3D().SetDimensionsVisible(0, false)
	if err != nil {
		t.Fatalf("SetDimensionsVisible: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DSetProperty {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DSetProperty)
	}
	var sent wire.SetSketch3DPropertyArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Property != "dimensionsVisible" || sent.Value != "false" {
		t.Errorf("sent = %+v, want property dimensionsVisible / value false", sent)
	}
	if got.DimensionsVisible {
		t.Error("decoded dimensionsVisible should be false")
	}
}

func TestSketch3DListSendsNilBodyAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"sketches":[{"index":0,"name":"3D Sketch1","visible":true,"dof":0,"healthy":true}]}`)}
	c := New(ft)

	res, err := c.Sketch3D().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DList)
	}
	if ft.gotReq != nil {
		t.Errorf("List should send a nil body, got %q", ft.gotReq)
	}
	if len(res.Sketches) != 1 || res.Sketches[0].Name != "3D Sketch1" {
		t.Errorf("decoded = %+v, want one 3D Sketch1", res.Sketches)
	}
}
