// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"oblikovati.org/api/wire"
)

func TestTransactionsUndoSendsNoBodyAndDecodesState(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"canUndo":false,"canRedo":true,"nextRedo":"Extrude"}`)}
	c := New(ft)

	st, err := c.Transactions().Undo()
	if err != nil {
		t.Fatalf("Undo: %v", err)
	}
	if ft.gotMethod != wire.MethodTransactionUndo {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodTransactionUndo)
	}
	if ft.gotReq != nil {
		t.Errorf("request body = %q, want nil for a no-arg method", ft.gotReq)
	}
	if st.CanUndo || !st.CanRedo || st.NextRedo != "Extrude" {
		t.Errorf("decoded state = %+v, want canUndo=false canRedo=true nextRedo=Extrude", st)
	}
}

func TestTransactionsStateDecodesLabels(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"canUndo":true,"canRedo":false,"nextUndo":"Fillet"}`)}
	c := New(ft)

	st, err := c.Transactions().State()
	if err != nil {
		t.Fatalf("State: %v", err)
	}
	if ft.gotMethod != wire.MethodTransactionState {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodTransactionState)
	}
	if !st.CanUndo || st.NextUndo != "Fillet" {
		t.Errorf("decoded state = %+v, want canUndo=true nextUndo=Fillet", st)
	}
}

// TestTransactionsAbortSendsNoBodyAndDecodesState covers the abort path: the
// discard of an open bounded transaction (M04-F05, Oblikovati#613).
func TestTransactionsAbortSendsNoBodyAndDecodesState(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"canUndo":true,"canRedo":false,"nextUndo":"Sketch"}`)}
	c := New(ft)

	st, err := c.Transactions().Abort()
	if err != nil {
		t.Fatalf("Abort: %v", err)
	}
	if ft.gotMethod != wire.MethodTransactionAbort {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodTransactionAbort)
	}
	if ft.gotReq != nil {
		t.Errorf("request body = %q, want nil for a no-arg method", ft.gotReq)
	}
	if !st.CanUndo || st.NextUndo != "Sketch" {
		t.Errorf("decoded state = %+v, want canUndo=true nextUndo=Sketch", st)
	}
}
