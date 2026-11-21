// SPDX-License-Identifier: Apache-2.0

package client

import (
	"reflect"
	"strings"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

const linkedTextSource = "A🦀\uFFFC\n\uFFFC"
const linkedTextStyle = `{"entityId":42,"style":{"content":"A🦀\ufffc\n\ufffc","height":0.5,"linkedFields":[{"at":2,"kind":"modelParameter","parameter":"d0","precision":0},{"at":4,"kind":"fileName"}]}}`

func TestSketchTextLinkedFieldsRoundTrip(t *testing.T) {
	zero := uint8(0)
	fields := []types.SketchTextField{
		{At: 2, Kind: types.SketchTextFieldModelParameter, Parameter: "d0", Precision: &zero},
		{At: 4, Kind: types.SketchTextFieldFileName},
	}
	transport := &fakeTransport{reply: []byte(`{"entityId":42}`)}
	sketch := New(transport).Sketch()
	if _, err := sketch.AddTextWith(wire.AddTextArgs{Text: linkedTextSource, LinkedFields: fields}); err != nil {
		t.Fatal(err)
	}
	expectLinkedFieldPayload(t, transport.gotReq)
	transport.reply = []byte(linkedTextStyle)
	result, err := sketch.EditText(wire.EditTextArgs{EntityID: 42, LinkedFields: &fields})
	assertResolvedLinkedFields(t, result, err, fields)
	expectLinkedFieldPayload(t, transport.gotReq)
	result, err = sketch.GetText(0, 42)
	assertResolvedLinkedFields(t, result, err, fields)
}

func expectLinkedFieldPayload(t *testing.T, request []byte) {
	t.Helper()
	const expected = `"linkedFields":[{"at":2,"kind":"modelParameter","parameter":"d0","precision":0},{"at":4,"kind":"fileName"}]`
	if !strings.Contains(string(request), expected) {
		t.Fatalf("request %s: expected %s", request, expected)
	}
}

func assertResolvedLinkedFields(t *testing.T, result wire.SketchTextResult, err error, fields []types.SketchTextField) {
	t.Helper()
	if err != nil || result.Style.Content != linkedTextSource || !reflect.DeepEqual(result.Style.LinkedFields, fields) {
		t.Fatalf("result %+v, error %v: expected %d linked fields in %q", result, err, len(fields), linkedTextSource)
	}
}

func TestSketchTextLinkedFieldsClearOrPreserve(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"entityId":42,"style":{"content":"plain","height":0.5}}`)}
	sketch := New(transport).Sketch()
	for _, fields := range []*[]types.SketchTextField{nil, &[]types.SketchTextField{}} {
		result, err := sketch.EditText(wire.EditTextArgs{EntityID: 42, LinkedFields: fields})
		if err != nil || len(result.Style.LinkedFields) != 0 {
			t.Fatalf("legacy result %+v, error %v: expected plain text", result, err)
		}
		request := string(transport.gotReq)
		if strings.Contains(request, `"linkedFields":[]`) != (fields != nil) {
			t.Fatalf("request %s: expected explicit empty linkedFields only when provided", request)
		}
	}
}

func TestSketchTextLinkedFieldsRejectInvalidTransit(t *testing.T) {
	transport := &fakeTransport{reply: []byte(strings.Replace(linkedTextStyle, `"precision":0`, `"precision":9`, 1))}
	sketch := New(transport).Sketch()
	fields := []types.SketchTextField{{At: 1, Kind: types.SketchTextFieldFileName}}
	if _, err := sketch.AddTextWith(wire.AddTextArgs{Text: linkedTextSource, LinkedFields: fields}); err == nil || transport.gotMethod != "" {
		t.Fatalf("method %q, error %v: expected to reject a field pointing at literal text before calling the host", transport.gotMethod, err)
	}
	if _, err := sketch.GetText(0, 42); err == nil || !strings.Contains(err.Error(), "zero through eight") {
		t.Fatalf("error %v: expected to reject precision 9 in the host result", err)
	}
}
