// SPDX-License-Identifier: Apache-2.0
package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestDocumentsOpenExpressSendsOptionAndReturnsMode(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":7,"name":"assembly","type":"assembly","dirty":false,"visible":true,"active":true,"express":true}`)}
	info, err := New(ft).Documents().Open(wire.OpenDocumentArgs{FullDocumentName: "/w/a.oad", Visible: true, Express: true})
	if err != nil {
		t.Fatal(err)
	}
	var args wire.OpenDocumentArgs
	if err := json.Unmarshal(ft.gotReq, &args); err != nil {
		t.Fatal(err)
	}
	if ft.gotMethod != wire.MethodDocumentsOpen || !args.Express || !info.Express {
		t.Fatalf("method %q args %+v result %+v: expected express assembly open", ft.gotMethod, args, info)
	}
}

func TestDocumentsLoadFullSendsIdentityAndReturnsFullMode(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":7,"name":"assembly","type":"assembly","dirty":false,"visible":true,"active":true}`)}
	info, err := New(ft).Documents().LoadFull(7)
	if err != nil {
		t.Fatal(err)
	}
	var args wire.LoadFullDocumentArgs
	if err := json.Unmarshal(ft.gotReq, &args); err != nil {
		t.Fatal(err)
	}
	if ft.gotMethod != wire.MethodDocumentsLoadFull || args.ID != 7 || info.Express {
		t.Fatalf("method %q args %+v result %+v: expected full assembly 7", ft.gotMethod, args, info)
	}
}

func TestDocumentsOpenLegacyWireDefaultsToFullMode(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":7,"name":"assembly","type":"assembly","dirty":false,"visible":true,"active":true}`)}
	info, err := New(ft).Documents().Open(wire.OpenDocumentArgs{FullDocumentName: "/w/a.oad", Visible: true})
	if err != nil {
		t.Fatal(err)
	}
	var request map[string]json.RawMessage
	if err := json.Unmarshal(ft.gotReq, &request); err != nil {
		t.Fatal(err)
	}
	if _, present := request["express"]; present || info.Express {
		t.Fatalf("request %s result %+v: expected omitted express and full mode", ft.gotReq, info)
	}
}
