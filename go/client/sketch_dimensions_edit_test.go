// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

// Deleting and moving a sketch dimension had no wire method at all — a dimension could be added
// and driven, but never removed or repositioned, unlike a drawing dimension (#2017).

func TestDeleteSketchDimensionSendsIndices(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"dof":1}`)}
	c := New(ft)

	got, err := c.Sketch().Dimension(2).Delete(3)
	if err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchDeleteDimension {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchDeleteDimension)
	}
	var sent wire.DeleteSketchDimensionArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.SketchIndex != 2 || sent.DimensionIndex != 3 {
		t.Errorf("sent=%+v, want sketch 2 dimension 3", sent)
	}
	if got.DOF != 1 {
		t.Errorf("DOF = %d, want the 1 the host reported", got.DOF)
	}
}

func TestMoveSketchDimensionSendsTextPoint(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	if _, err := c.Sketch().Dimension(0).Move(1, 2.5, -4); err != nil {
		t.Fatalf("Move: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchMoveDimension {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchMoveDimension)
	}
	var sent wire.MoveSketchDimensionArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.DimensionIndex != 1 || len(sent.TextPoint) != 2 {
		t.Fatalf("sent=%+v, want dimension 1 with an [x,y] text point", sent)
	}
	if sent.TextPoint[0] != 2.5 || sent.TextPoint[1] != -4 {
		t.Errorf("textPoint = %v, want [2.5 -4] in the order given", sent.TextPoint)
	}
}
