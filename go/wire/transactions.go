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

// TransactionHistoryArgs is the request of [MethodTransactionHistory]: read the full undo
// stream of one open document for a history browser. Document is the document's id (from
// documents.list); 0 (or omitted) means the active document. Reading another document's
// history does not activate it — a history browser can show several documents' timelines
// side by side.
type TransactionHistoryArgs struct {
	Document uint64 `json:"document,omitempty"`
}

// TransactionHistoryEntry is one committed step in a document's undo stream, oldest first.
// Saved marks the steps after which the document was written to disk (a save checkpoint), so
// a browser can show which edits are persisted versus only in memory.
type TransactionHistoryEntry struct {
	Label string `json:"label"`
	Saved bool   `json:"saved,omitempty"`
}

// TransactionHistory is the response of [MethodTransactionHistory] and [MethodTransactionJumpTo]:
// one open document's whole undo stream for a history browser. Entries lists every committed
// step oldest-first; Position is the cursor — how many steps are currently applied, so
// Entries[:Position] are undoable past and Entries[Position:] are redoable future, and
// Position 0 is the document's open/baseline state. Saving does not truncate the stream, so a
// browser shows every event since the document was opened, with the save points flagged.
type TransactionHistory struct {
	Document uint64                    `json:"document"`
	Name     string                    `json:"name"`
	Position int                       `json:"position"`
	Entries  []TransactionHistoryEntry `json:"entries"`
}

// TransactionJumpToArgs is the request of [MethodTransactionJumpTo]: move one document's undo
// cursor to an absolute Position (0 = open state, len(Entries) = latest), undoing or redoing
// as many steps as needed in one call — the click-to-jump a history browser needs for a long
// stream. Document is the document's id; 0 (or omitted) means the active document.
type TransactionJumpToArgs struct {
	Document uint64 `json:"document,omitempty"`
	Position int    `json:"position"`
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
