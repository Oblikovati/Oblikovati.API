// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

// TestDocumentsGetEndOfPart asserts GetEndOfPart calls the method and decodes the marker (#141).
func TestDocumentsGetEndOfPart(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"position":2,"rolledBack":true}`)}
	c := New(ft)

	got, err := c.Documents().GetEndOfPart()
	if err != nil {
		t.Fatalf("GetEndOfPart: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentGetEndOfPart {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentGetEndOfPart)
	}
	if got.Position != 2 || !got.RolledBack {
		t.Errorf("result = %+v, want position 2 / rolled back", got)
	}
}

// TestDocumentsSetEndOfPartSendsPosition asserts SetEndOfPart marshals the target position (#141).
func TestDocumentsSetEndOfPartSendsPosition(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"position":-1,"rolledBack":false}`)}
	c := New(ft)

	got, err := c.Documents().SetEndOfPart(-1)
	if err != nil {
		t.Fatalf("SetEndOfPart: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentSetEndOfPart {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentSetEndOfPart)
	}
	var sent wire.SetEndOfPartArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Position != -1 {
		t.Errorf("sent position = %d, want -1", sent.Position)
	}
	if got.RolledBack {
		t.Errorf("result = %+v, want not rolled back after restore-to-end", got)
	}
}
