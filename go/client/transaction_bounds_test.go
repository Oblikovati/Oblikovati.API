// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestTransactionsBeginMarshalsLabel(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	res, err := c.Transactions().Begin("apply remote batch")
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	if ft.gotMethod != wire.MethodTransactionBegin {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodTransactionBegin)
	}
	var sent wire.TransactionBeginArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Label != "apply remote batch" {
		t.Errorf("sent label = %q, want %q", sent.Label, "apply remote batch")
	}
	if !res.OK {
		t.Errorf("decoded = %+v, want ok=true", res)
	}
}

func TestTransactionsEndSendsNilBodyAndDecodesState(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"canUndo":true,"canRedo":false,"nextUndo":"apply remote batch"}`)}
	c := New(ft)

	st, err := c.Transactions().End()
	if err != nil {
		t.Fatalf("End: %v", err)
	}
	if ft.gotMethod != wire.MethodTransactionEnd {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodTransactionEnd)
	}
	if ft.gotReq != nil {
		t.Errorf("End should send nil body, got %q", ft.gotReq)
	}
	if !st.CanUndo || st.NextUndo != "apply remote batch" {
		t.Errorf("decoded = %+v, want canUndo=true nextUndo='apply remote batch'", st)
	}
}
