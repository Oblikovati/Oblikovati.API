// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Transaction lifecycle push events (M04-F05, Oblikovati/Oblikovati#613): the
// host announces every move of a document's transaction stream so an add-in can
// keep external state (caches, overlays, exported data) consistent with
// undo/redo. Like [EditCommittedEvent], these are push-only — the host delivers
// them to an add-in's Notify entry point and the add-in matches on Type.

// TransactionEventPayload is the JSON shape of the five transaction push events
// ([EventTransactionCommitted], [EventTransactionUndone],
// [EventTransactionRedone], [EventTransactionAborted],
// [EventTransactionDeleted]). Document is the affected document's id, Label the
// display name of the undo step acted on ("" when the event has no single step,
// e.g. a deleted stream), and Point locates that step relative to the stream's
// cursor: a commit acts on the current point, undo on the previous, redo on the
// next, a deleted stream on no point at all.
type TransactionEventPayload struct {
	Type     string                 `json:"type"`
	Document uint64                 `json:"document"`
	Label    string                 `json:"label,omitempty"`
	Point    types.TransactionPoint `json:"point,omitempty"`
}
