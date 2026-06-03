// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"github.com/Oblikovati/api/types"
	"github.com/Oblikovati/api/wire"
)

func TestSketchSetLineTypeSendsPropertyAndValue(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"name":"Sketch1","plane":"XY","lineType":"center"}`)}
	c := New(ft)

	got, err := c.Sketch().SetLineType(0, types.SketchLineCenter)
	if err != nil {
		t.Fatalf("SetLineType: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchSetProperty {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchSetProperty)
	}
	var sent wire.SetSketchPropertyArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Property != "lineType" || sent.Value != "center" {
		t.Errorf("sent = %+v, want property lineType / value center", sent)
	}
	if got.LineType != "center" {
		t.Errorf("decoded lineType = %q, want center", got.LineType)
	}
}

func TestSketchGetSendsIndexAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":2,"name":"Sketch3","plane":"XY","visible":true,"entityCount":8,"dof":8,"editing":false,"healthy":true}`)}
	c := New(ft)

	got, err := c.Sketch().Get(2)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchGet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchGet)
	}
	var sent wire.SketchArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.SketchIndex != 2 {
		t.Errorf("sent index = %d, want 2", sent.SketchIndex)
	}
	if got.Name != "Sketch3" || got.EntityCount != 8 || got.DOF != 8 {
		t.Errorf("decoded = %+v, want Sketch3 / 8 entities / DOF 8", got)
	}
}

func TestSketchListSendsNilBodyAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"sketches":[{"index":0,"name":"Sketch1","plane":"XY","dof":0,"healthy":true}]}`)}
	c := New(ft)

	res, err := c.Sketch().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchList)
	}
	if ft.gotReq != nil {
		t.Errorf("List sent a body %q, want nil", string(ft.gotReq))
	}
	if len(res.Sketches) != 1 || res.Sketches[0].Name != "Sketch1" {
		t.Errorf("decoded = %+v, want one sketch named Sketch1", res.Sketches)
	}
}

func TestSketchAddCircleByCenterRadiusMarshalsKindVariant(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":7,"kind":"circle","pointIds":[6]}`)}
	c := New(ft)

	got, err := c.Sketch().AddCircleByCenterRadius(0, []float64{0, 0}, "10 mm", false)
	if err != nil {
		t.Fatalf("AddCircleByCenterRadius: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchAddEntity {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchAddEntity)
	}
	var sent wire.AddSketchEntityArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "circle" || sent.Variant != "centerRadius" || sent.Radius != "10 mm" || len(sent.Points) != 1 {
		t.Errorf("sent = %+v, want circle/centerRadius/10 mm/1 point", sent)
	}
	if got.EntityID != 7 || len(got.PointIDs) != 1 {
		t.Errorf("decoded = %+v, want entity 7 / 1 point id", got)
	}
}

func TestSketchAddArcByThreePointsSetsVariant(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":9,"kind":"arc","pointIds":[6,7,8]}`)}
	c := New(ft)

	if _, err := c.Sketch().AddArcByThreePoints(0, []float64{2, 0}, []float64{0, 2}, []float64{-2, 0}, false); err != nil {
		t.Fatalf("AddArcByThreePoints: %v", err)
	}
	var sent wire.AddSketchEntityArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "arc" || sent.Variant != "threePoint" || len(sent.Points) != 3 {
		t.Errorf("sent = %+v, want arc/threePoint/3 points", sent)
	}
}

func TestSketchSolveDecodesStatus(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"sketchIndex":0,"dof":0,"status":"well","converged":true,"healthy":true}`)}
	c := New(ft)

	got, err := c.Sketch().Solve(0)
	if err != nil {
		t.Fatalf("Solve: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchSolve {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchSolve)
	}
	if got.DOF != 0 || got.Status != "well" || !got.Converged {
		t.Errorf("decoded = %+v, want DOF 0 / well / converged", got)
	}
}
