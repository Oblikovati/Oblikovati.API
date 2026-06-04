// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"github.com/Oblikovati/api/types"
	"github.com/Oblikovati/api/wire"
)

func TestGraphicsAddHeatmapMarshalsPerVertexRequest(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"clientId":"fea","nodeCount":1,"primitiveCount":1}`)}
	c := New(ft)

	mapper := wire.GraphicsColorMapper{Values: []float64{0, 1}, Colors: []float32{0, 0, 1, 1, 1, 0, 0, 1}}
	got, err := c.Graphics().AddHeatmap("fea",
		[]float64{0, 0, 0, 1, 0, 0, 0, 1, 0}, []int{0, 1, 2}, []float64{0, 0.5, 1}, mapper)
	if err != nil {
		t.Fatalf("AddHeatmap: %v", err)
	}
	if ft.gotMethod != wire.MethodClientGraphicsSet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodClientGraphicsSet)
	}
	var sent wire.SetClientGraphicsArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ClientId != "fea" || sent.Lane != string(types.GraphicsLanePersistent) {
		t.Errorf("sent = %+v, want clientId=fea lane=persistent", sent)
	}
	if len(sent.Nodes) != 1 || len(sent.Nodes[0].Primitives) != 1 {
		t.Fatalf("sent nodes = %+v, want one node with one primitive", sent.Nodes)
	}
	p := sent.Nodes[0].Primitives[0]
	if p.Kind != string(types.GraphicsTriangles) || p.ColorBinding != string(types.GraphicsColorPerVertex) {
		t.Errorf("primitive = %+v, want triangles+perVertex", p)
	}
	if p.ColorMapper == nil || len(p.Scalars) != 3 {
		t.Errorf("primitive = %+v, want mapper + 3 scalars", p)
	}
	if got.PrimitiveCount != 1 {
		t.Errorf("decoded primitiveCount = %d, want 1", got.PrimitiveCount)
	}
}

func TestGraphicsDeleteSendsClientId(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	if err := c.Graphics().Delete("fea"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if ft.gotMethod != wire.MethodClientGraphicsDelete {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodClientGraphicsDelete)
	}
	var sent wire.DeleteClientGraphicsArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ClientId != "fea" {
		t.Errorf("sent clientId = %q, want fea", sent.ClientId)
	}
}

func TestInteractionUpdateSendsLane(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	nodes := []wire.GraphicsNode{{Primitives: []wire.GraphicsPrimitive{{
		Kind: string(types.GraphicsLines), Coordinates: []float64{0, 0, 0, 1, 1, 1}, Indices: []int{0, 1},
	}}}}
	if err := c.Graphics().Interaction().Update(types.GraphicsLaneOverlay, nodes); err != nil {
		t.Fatalf("Update: %v", err)
	}
	if ft.gotMethod != wire.MethodInteractionGraphicsUpdate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodInteractionGraphicsUpdate)
	}
	var sent wire.UpdateInteractionGraphicsArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Lane != string(types.GraphicsLaneOverlay) {
		t.Errorf("sent lane = %q, want overlay", sent.Lane)
	}
}
