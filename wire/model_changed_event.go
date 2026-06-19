// SPDX-License-Identifier: Apache-2.0

package wire

// ModelChangedEvent is the [EventModelChanged] push payload (#148): a committed batch of
// model changes on a document — feature/sketch/parameter mutations the engine just applied.
// Document is the affected document's display name; Changes is how many changes were in the
// batch. An add-in matches on Type and re-queries the document's state.
type ModelChangedEvent struct {
	Type     string `json:"type"`
	Document string `json:"document,omitempty"`
	Changes  int    `json:"changes,omitempty"`
}
