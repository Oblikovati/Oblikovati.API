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

// TestTransactionsHistorySendsDocumentAndDecodesEntries covers the history-browser read:
// the document id goes out in the request and the entries (with save flags) decode back.
func TestTransactionsHistorySendsDocumentAndDecodesEntries(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"document":7,"name":"Part1","position":1,"entries":[{"label":"Extrude","saved":true}]}`)}
	c := New(ft)

	h, err := c.Transactions().History(7)
	if err != nil {
		t.Fatalf("History: %v", err)
	}
	if ft.gotMethod != wire.MethodTransactionHistory {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodTransactionHistory)
	}
	if string(ft.gotReq) != `{"document":7}` {
		t.Errorf("request body = %q, want the document id", ft.gotReq)
	}
	if h.Document != 7 || h.Name != "Part1" || h.Position != 1 {
		t.Errorf("decoded history = %+v, want document 7 Part1 at position 1", h)
	}
	if len(h.Entries) != 1 || h.Entries[0].Label != "Extrude" || !h.Entries[0].Saved {
		t.Errorf("decoded entries = %+v, want one saved Extrude", h.Entries)
	}
}

// TestTransactionsJumpToSendsPosition covers the click-to-jump call.
func TestTransactionsJumpToSendsPosition(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"document":7,"name":"Part1","position":0,"entries":[]}`)}
	c := New(ft)

	h, err := c.Transactions().JumpTo(7, 0)
	if err != nil {
		t.Fatalf("JumpTo: %v", err)
	}
	if ft.gotMethod != wire.MethodTransactionJumpTo {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodTransactionJumpTo)
	}
	if string(ft.gotReq) != `{"document":7,"position":0}` {
		t.Errorf("request body = %q, want document+position", ft.gotReq)
	}
	if h.Position != 0 {
		t.Errorf("decoded position = %d, want 0", h.Position)
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
