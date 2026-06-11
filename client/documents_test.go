// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

// TestDocumentsCloseMarshalsIDAndForce asserts the typed Close call sends the
// documented CloseDocumentArgs shape and decodes the closed count.
func TestDocumentsCloseMarshalsIDAndForce(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"closed":1}`)}
	c := New(ft)

	got, err := c.Documents().Close(7, true)
	if err != nil {
		t.Fatalf("Close: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentsClose {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentsClose)
	}
	var sent wire.CloseDocumentArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ID != 7 || !sent.Force {
		t.Errorf("sent = %+v, want id=7 force=true", sent)
	}
	if got.Closed != 1 {
		t.Errorf("decoded closed = %d, want 1", got.Closed)
	}
}

// TestDocumentsCloseAllMarshalsForce asserts the typed CloseAll call hits the
// closeAll method with the force flag.
func TestDocumentsCloseAllMarshalsForce(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"closed":3}`)}
	c := New(ft)

	got, err := c.Documents().CloseAll(true)
	if err != nil {
		t.Fatalf("CloseAll: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentsCloseAll {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentsCloseAll)
	}
	var sent wire.CloseAllDocumentsArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if !sent.Force {
		t.Errorf("sent = %+v, want force=true", sent)
	}
	if got.Closed != 3 {
		t.Errorf("decoded closed = %d, want 3", got.Closed)
	}
}
