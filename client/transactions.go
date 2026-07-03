// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Transactions is the undo/redo operation group for the active document's transaction
// stream. Undo and Redo are navigators over that stream — non-destructive cursor moves,
// not a destructive pop — so an undone step can be redone until a new edit truncates the
// forward branch.
type Transactions struct{ c *Client }

// Transactions returns the undo/redo operation group.
func (c *Client) Transactions() Transactions { return Transactions{c} }

// Undo moves the active document's cursor back one transaction event and returns the
// resulting state (what undo/redo can do next).
//
// mcp:tool undo
// mcp:summary Undo the active document's last change (step the transaction cursor back).
func (t Transactions) Undo() (wire.UndoState, error) {
	return call[wire.UndoState](t.c, wire.MethodTransactionUndo, nil)
}

// Redo moves the cursor forward one transaction event and returns the resulting state.
//
// mcp:tool redo
// mcp:summary Redo the next change (step the transaction cursor forward).
func (t Transactions) Redo() (wire.UndoState, error) {
	return call[wire.UndoState](t.c, wire.MethodTransactionRedo, nil)
}

// State reports what undo/redo can currently do, with the labels of the next steps.
//
// mcp:tool get_undo_state
// mcp:summary Read the active document's undo/redo state (whether undo/redo are available and the cursor position).
func (t Transactions) State() (wire.UndoState, error) {
	return call[wire.UndoState](t.c, wire.MethodTransactionState, nil)
}

// History reads one open document's whole undo stream — every step since the document was
// opened, with the cursor position and the save checkpoints — for a history browser. Pass a
// document id from documents.list, or 0 for the active document. Reading does not activate the
// document, so a browser can show several documents' timelines side by side.
//
// mcp:tool get_history
// mcp:summary Read a document's full undo history (every step since it was opened, the cursor position, and which steps are saved). document=0 means the active document.
func (t Transactions) History(document uint64) (wire.TransactionHistory, error) {
	return call[wire.TransactionHistory](t.c, wire.MethodTransactionHistory, wire.TransactionHistoryArgs{Document: document})
}

// JumpTo moves one document's undo cursor to an absolute position (0 = open state,
// len(entries) = latest), undoing or redoing as many steps as needed in one call. It returns
// the document's resulting history. Pass a document id from documents.list, or 0 for the
// active document.
//
// mcp:tool jump_to_history
// mcp:summary Jump a document's undo cursor to an absolute position (0=open state), undoing/redoing as many steps as needed. document=0 means the active document.
func (t Transactions) JumpTo(document uint64, position int) (wire.TransactionHistory, error) {
	return call[wire.TransactionHistory](t.c, wire.MethodTransactionJumpTo, wire.TransactionJumpToArgs{Document: document, Position: position})
}

// Begin opens a bounded transaction: every edit recorded until the matching End is
// coalesced into a single undo step named label. Begin/End nest; only the outermost End
// commits the group. Use it to make a batch of operations one team-shared undo step.
//
//	client.Transactions().Begin("apply remote batch")
//	// … several edits …
//	client.Transactions().End()
//
// mcp:tool transaction_begin
// mcp:summary Opens a bounded transaction: every edit recorded until the matching End is coalesced into a single undo step named label.
func (t Transactions) Begin(label string) (wire.OKResult, error) {
	return call[wire.OKResult](t.c, wire.MethodTransactionBegin, wire.TransactionBeginArgs{Label: label})
}

// End closes the innermost open transaction and returns the resulting undo/redo state.
//
// mcp:tool transaction_end
// mcp:summary Closes the innermost open transaction and returns the resulting undo/redo state.
func (t Transactions) End() (wire.UndoState, error) {
	return call[wire.UndoState](t.c, wire.MethodTransactionEnd, nil)
}

// Abort discards the innermost open transaction instead of committing it: the model
// reverts to the group's pre-Begin state and no undo step is recorded. Use it when a
// batch fails partway so the document is not left half-edited (M04-F05).
//
// mcp:tool transaction_abort
// mcp:summary Discards the innermost open transaction instead of committing it: the model reverts to the group's pre-Begin state and no undo step is recorded.
func (t Transactions) Abort() (wire.UndoState, error) {
	return call[wire.UndoState](t.c, wire.MethodTransactionAbort, nil)
}
