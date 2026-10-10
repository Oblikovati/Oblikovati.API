// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

// detailReply is a minimal FeatureDetailResult payload shared by the freeform
// client tests.
const detailReply = `{"feature":{"id":7,"name":"Freeform1","index":0}}`

func TestFreeformSetLevelMarshalsIDAndLevel(t *testing.T) {
	ft := &fakeTransport{reply: []byte(detailReply)}
	c := New(ft)

	got, err := c.Freeform().SetLevel(7, 2)
	if err != nil {
		t.Fatalf("SetLevel: %v", err)
	}
	if ft.gotMethod != wire.MethodFreeformSetLevel {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodFreeformSetLevel)
	}
	var sent wire.SetFreeformLevelArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ID != 7 || sent.Level != 2 {
		t.Errorf("sent = %+v, want id 7 level 2", sent)
	}
	if got.Feature.ID != 7 {
		t.Errorf("decoded feature id = %d, want 7", got.Feature.ID)
	}
}

func TestFreeformMoveVerticesMarshalsSelection(t *testing.T) {
	ft := &fakeTransport{reply: []byte(detailReply)}
	c := New(ft)

	_, err := c.Freeform().MoveVertices(wire.MoveFreeformVerticesArgs{
		ID: 7, Vertices: []int{0, 3}, Translation: [3]float64{1, 0, -2},
	})
	if err != nil {
		t.Fatalf("MoveVertices: %v", err)
	}
	if ft.gotMethod != wire.MethodFreeformMoveVertices {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodFreeformMoveVertices)
	}
	var sent wire.MoveFreeformVerticesArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if len(sent.Vertices) != 2 || sent.Vertices[1] != 3 || sent.Translation[2] != -2 {
		t.Errorf("sent = %+v, want vertices [0 3] translation [1 0 -2]", sent)
	}
}

func TestFreeformCreaseEdgesMarshalsSharpness(t *testing.T) {
	ft := &fakeTransport{reply: []byte(detailReply)}
	c := New(ft)

	_, err := c.Freeform().CreaseEdges(wire.CreaseFreeformEdgesArgs{
		ID: 7, Edges: [][2]int{{0, 1}, {1, 2}}, Sharpness: 1,
	})
	if err != nil {
		t.Fatalf("CreaseEdges: %v", err)
	}
	if ft.gotMethod != wire.MethodFreeformCreaseEdges {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodFreeformCreaseEdges)
	}
	var sent wire.CreaseFreeformEdgesArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if len(sent.Edges) != 2 || sent.Edges[0] != [2]int{0, 1} || sent.Sharpness != 1 {
		t.Errorf("sent = %+v, want 2 edges sharpness 1", sent)
	}
}
