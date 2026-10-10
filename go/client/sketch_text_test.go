// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestAddTextWithMarshalsFullStyle(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":42}`)}
	c := New(ft)

	in := wire.AddTextArgs{
		SketchIndex: 0, Anchor: []float64{1, 1}, Text: "PART A",
		Height: "5 mm", Justify: "center", VJustify: "middle",
		Font: "Liberation Sans", FontSize: "5 mm",
	}
	got, err := c.Sketch().AddTextWith(in)
	if err != nil {
		t.Fatalf("AddTextWith: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchAddText {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchAddText)
	}
	var sent wire.AddTextArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Text != "PART A" || sent.VJustify != "middle" || sent.Font != "Liberation Sans" {
		t.Errorf("sent = %+v, want text/valign/font preserved", sent)
	}
	if got.EntityID != 42 {
		t.Errorf("entityId = %d, want 42", got.EntityID)
	}
}

func TestEditTextSendsPartialFields(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":42,"style":{"content":"NEW","height":0.5,"hAlign":"center","vAlign":"baseline"}}`)}
	c := New(ft)

	newText := "NEW"
	got, err := c.Sketch().EditText(wire.EditTextArgs{SketchIndex: 0, EntityID: 42, Text: &newText})
	if err != nil {
		t.Fatalf("EditText: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchEditText {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchEditText)
	}
	if got.Style.Content != "NEW" || got.Style.HAlign != types.TextAlignCenter {
		t.Errorf("decoded style = %+v, want content NEW / center", got.Style)
	}
}

func TestGetTextDecodesStyle(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":42,"style":{"content":"X","family":"Liberation Sans","height":0.5,"vAlign":"upper"}}`)}
	c := New(ft)

	got, err := c.Sketch().GetText(0, 42)
	if err != nil {
		t.Fatalf("GetText: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchGetText {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchGetText)
	}
	if got.Style.Family != "Liberation Sans" || got.Style.VAlign != types.TextAlignUpper {
		t.Errorf("decoded = %+v, want family/upper", got.Style)
	}
}
