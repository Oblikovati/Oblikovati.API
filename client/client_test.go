// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"errors"
	"testing"

	"oblikovati/api/types"
	"oblikovati/api/wire"
)

// fakeTransport records the last method+request and replies with canned bytes, so a
// client call can be asserted end-to-end without a live host.
type fakeTransport struct {
	gotMethod string
	gotReq    []byte
	reply     []byte
	err       error
}

func (f *fakeTransport) Call(method string, req []byte) ([]byte, error) {
	f.gotMethod, f.gotReq = method, req
	return f.reply, f.err
}

func TestParametersAddMarshalsRequestAndDecodesReply(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"name":"height","kind":"user","expression":"3 cm","value":"3 cm"}`)}
	c := New(ft)

	got, err := c.Parameters().Add(wire.ParameterSetArgs{Name: "height", Expression: "3 cm"})
	if err != nil {
		t.Fatalf("Add: %v", err)
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

func TestNoArgMethodSendsNilBody(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"documents":[]}`)}
	c := New(ft)

	if _, err := c.Documents().List(); err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentsList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentsList)
	}
	if ft.gotReq != nil {
		t.Errorf("request body = %q, want nil for a no-arg method", ft.gotReq)
	}
}

func TestCommandsExecuteWrapsIDArg(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	res, err := c.Commands().Execute("sketch.finish")
	if err != nil {
		t.Fatalf("Execute: %v", err)
	}
	if !res.OK {
		t.Error("OK = false, want true")
	}
	var sent wire.ExecuteCommandArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.ID != "sketch.finish" {
		t.Errorf("sent = %+v (err %v), want id=sketch.finish", sent, err)
	}
}

func TestCommandsCreateMarshalsButtonMetadata(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	res, err := c.Commands().Create(wire.CreateCommandArgs{
		ID: "AddIn.Ping", DisplayName: "Ping", Tab: "AddInTab",
		Category: "Demo", Icon: "extrude", ButtonStyle: types.LargeIconButton,
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if !res.OK {
		t.Error("OK = false, want true")
	}
	if ft.gotMethod != wire.MethodCommandsCreate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodCommandsCreate)
	}
	var sent wire.CreateCommandArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ID != "AddIn.Ping" || sent.DisplayName != "Ping" || sent.ButtonStyle != types.LargeIconButton {
		t.Errorf("sent = %+v, want id=AddIn.Ping displayName=Ping style=large-icon", sent)
	}
}

func TestTransportErrorPropagates(t *testing.T) {
	ft := &fakeTransport{err: errors.New("host: no active document")}
	c := New(ft)
	if _, err := c.Model().Tree(); err == nil || err.Error() != "host: no active document" {
		t.Errorf("err = %v, want the host error verbatim", err)
	}
}

func TestNilTransportFailsClearly(t *testing.T) {
	c := New(nil)
	_, err := c.Parameters().List()
	if err == nil || !contains(err.Error(), wire.MethodParametersList) {
		t.Errorf("err = %v, want a nil-transport error naming the method", err)
	}
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
