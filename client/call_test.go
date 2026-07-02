// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"errors"
	"testing"

	"oblikovati.org/api/wire"
)

// TestCallGenericMarshalsRequestAndDecodesTypedReply is the fake-Transport guard for
// the generic call[Resp] helper (Oblikovati/Oblikovati#1650): the emitted method
// constant + request JSON and the decoded response must match the untyped path.
func TestCallGenericMarshalsRequestAndDecodesTypedReply(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"name":"height","kind":"user","expression":"3 cm","value":"3 cm"}`)}
	c := New(ft)

	got, err := call[wire.ParameterInfo](c, wire.MethodParametersAdd,
		wire.ParameterSetArgs{Name: "height", Expression: "3 cm"})
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	if ft.gotMethod != wire.MethodParametersAdd {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodParametersAdd)
	}
	var sent wire.ParameterSetArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Name != "height" || sent.Expression != "3 cm" {
		t.Errorf("sent = %+v, want height=3 cm", sent)
	}
	if got.Name != "height" || got.Expression != "3 cm" {
		t.Errorf("decoded = %+v, want height=3 cm", got)
	}
}

// TestCallGenericNilRequestSendsNilBody mirrors the untyped nil-request contract:
// req=nil must reach the transport as a nil body, not "null".
func TestCallGenericNilRequestSendsNilBody(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"parameters":[]}`)}
	c := New(ft)

	if _, err := call[wire.ListParametersResult](c, wire.MethodParametersList, nil); err != nil {
		t.Fatalf("call: %v", err)
	}
	if ft.gotReq != nil {
		t.Errorf("request body = %q, want nil for a nil request", ft.gotReq)
	}
}

// TestCallGenericTransportErrorReturnsZeroValue asserts the (zero, err) contract of
// the old three-line bodies: on transport failure the typed result is the zero DTO.
func TestCallGenericTransportErrorReturnsZeroValue(t *testing.T) {
	wantErr := errors.New("host gone")
	c := New(&fakeTransport{err: wantErr})

	got, err := call[wire.ParameterInfo](c, wire.MethodParametersGet,
		wire.ParameterNameArgs{Name: "height"})
	if !errors.Is(err, wantErr) {
		t.Fatalf("err = %v, want %v", err, wantErr)
	}
	if got != (wire.ParameterInfo{}) {
		t.Errorf("result = %+v, want zero value on error", got)
	}
}

// TestCallGenericMatchesUntypedCallGoldenJSON is the before/after equivalence gate
// from #1650: both paths must emit byte-identical requests and decode equal replies.
func TestCallGenericMatchesUntypedCallGoldenJSON(t *testing.T) {
	reply := []byte(`{"name":"width","kind":"user","expression":"7 mm","value":"7 mm"}`)
	args := wire.ParameterSetArgs{Name: "width", Expression: "7 mm"}

	oldFT := &fakeTransport{reply: reply}
	var oldRes wire.ParameterInfo
	if err := New(oldFT).call(wire.MethodParametersSet, args, &oldRes); err != nil {
		t.Fatalf("untyped call: %v", err)
	}

	newFT := &fakeTransport{reply: reply}
	newRes, err := call[wire.ParameterInfo](New(newFT), wire.MethodParametersSet, args)
	if err != nil {
		t.Fatalf("generic call: %v", err)
	}

	if oldFT.gotMethod != newFT.gotMethod {
		t.Errorf("method: untyped %q vs generic %q", oldFT.gotMethod, newFT.gotMethod)
	}
	if string(oldFT.gotReq) != string(newFT.gotReq) {
		t.Errorf("request JSON: untyped %q vs generic %q", oldFT.gotReq, newFT.gotReq)
	}
	if oldRes != newRes {
		t.Errorf("decoded: untyped %+v vs generic %+v", oldRes, newRes)
	}
}
