// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

// Transactions is the undo/redo operation group for the active document's transaction
// stream. Undo and Redo are navigators over that stream — non-destructive cursor moves,
// not a destructive pop — so an undone step can be redone until a new edit truncates the
// forward branch.
type Transactions struct{ c *Client }

// Transactions returns the undo/redo operation group.
func (c *Client) Transactions() Transactions { return Transactions{c} }

// Undo moves the active document's cursor back one transaction event and returns the
// resulting state (what undo/redo can do next).
func (t Transactions) Undo() (wire.UndoState, error) {
	var r wire.UndoState
	return r, t.c.call(wire.MethodTransactionUndo, nil, &r)
}

// Redo moves the cursor forward one transaction event and returns the resulting state.
func (t Transactions) Redo() (wire.UndoState, error) {
	var r wire.UndoState
	return r, t.c.call(wire.MethodTransactionRedo, nil, &r)
}

// State reports what undo/redo can currently do, with the labels of the next steps.
func (t Transactions) State() (wire.UndoState, error) {
	var r wire.UndoState
	return r, t.c.call(wire.MethodTransactionState, nil, &r)
}
