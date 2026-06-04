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
