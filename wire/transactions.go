// SPDX-License-Identifier: Apache-2.0

package wire

// UndoState is the JSON shape of the active document's transaction-stream cursor: what
// undo and redo can do next. It is the result of [MethodTransactionState] and of the
// mutating [MethodTransactionUndo] / [MethodTransactionRedo] (which return the state the
// move produced, so a caller learns the new cursor position in one round trip). The
// Next* labels are the names of the steps undo/redo would act on, for a menu/tooltip;
// they are empty when there is nothing to act on.
type UndoState struct {
	CanUndo  bool   `json:"canUndo"`
	CanRedo  bool   `json:"canRedo"`
	NextUndo string `json:"nextUndo,omitempty"`
	NextRedo string `json:"nextRedo,omitempty"`
}

// TransactionBeginArgs is the request of [MethodTransactionBegin]: open a bounded
// transaction that coalesces every edit recorded until the matching [MethodTransactionEnd]
// into a single undo step. Label names that step (for the undo menu/tooltip). Begin/End
// nest; only the outermost End commits the group. Begin returns [OKResult]; End returns
// the resulting [UndoState].
//
// A collaboration add-in drains its buffer of remote operations inside one Begin/End so
// the whole batch is one team-shared undo step (oblikovati-meeting ADR-0005).
type TransactionBeginArgs struct {
	Label string `json:"label,omitempty"`
}
