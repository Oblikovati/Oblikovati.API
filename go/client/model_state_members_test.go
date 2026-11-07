// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"oblikovati.org/api/wire"
	"testing"
)

func TestRenameModelStateUsesCanonicalRequestAndKeepsItsID(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"modelState":{"id":7,"name":"Prototype","active":true}}`)}
	request := wire.ModelStateNameArgs{ID: 7, Name: "Prototype"}
	result, err := New(transport).ModelStates().Rename(request)
	if err != nil {
		t.Fatal(err)
	}
	checkMemberRequest(t, transport, wire.MethodModelStatesRename, request)
	if result.ModelState.ID != 7 || result.ModelState.Name != "Prototype" || !result.ModelState.Active {
		t.Fatalf("renamed member = %+v, want active Prototype with id7", result.ModelState)
	}
}

func TestCopyModelStateUsesCanonicalRequestAndDecodesAnInactiveCopy(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"modelState":{"id":8,"name":"Prototype Copy","primary":false}}`)}
	request := wire.ModelStateNameArgs{ID: 7, Name: "Prototype Copy"}
	result, err := New(transport).ModelStates().Copy(request)
	if err != nil {
		t.Fatal(err)
	}
	checkMemberRequest(t, transport, wire.MethodModelStatesCopy, request)
	if result.ModelState.ID != 8 || result.ModelState.Name != "Prototype Copy" || result.ModelState.Active || result.ModelState.Primary {
		t.Fatalf("copied member = %+v, want inactive non-primary Prototype Copy with id8", result.ModelState)
	}
}

func checkMemberRequest(t *testing.T, transport *fakeTransport, method string, request wire.ModelStateNameArgs) {
	t.Helper()
	if transport.gotMethod != method {
		t.Fatalf("method = %q, want %q", transport.gotMethod, method)
	}
	var decoded wire.ModelStateNameArgs
	if err := json.Unmarshal(transport.gotReq, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded != request {
		t.Fatalf("request = %+v, want %+v", decoded, request)
	}
}
