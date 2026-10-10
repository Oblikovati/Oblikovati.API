// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Assembly occurrence-lifecycle push events (M11-F07, Oblikovati/Oblikovati#723):
// the host announces every change to an assembly's component occurrences so an
// add-in can keep external state (overlays, BOM mirrors, analysis) consistent
// without polling. Like [TransactionEventPayload], these are push-only — the host
// delivers them to an add-in's Notify entry point and the add-in matches on Type.

// OccurrenceEventPayload is the JSON shape of the five occurrence push events
// ([EventOccurrenceAdded], [EventOccurrenceDeleted], [EventOccurrenceReplaced],
// [EventOccurrenceTransformed], [EventOccurrenceSuppressed]). Document is the
// assembly document's id and Occurrence the affected occurrence's session id, with
// Name its instance name (e.g. "bracket:2"). The remaining fields are populated only
// for the events that carry them: Suppressed reports the new state on a suppressed
// event; Transform and Previous carry the new and prior placements on a transformed
// event (a coalesced solver drag reports the net move — Previous is the pre-drag
// placement). Add/delete/replace carry identity alone.
type OccurrenceEventPayload struct {
	Type       string        `json:"type"`
	Document   uint64        `json:"document"`
	Occurrence uint64        `json:"occurrence"`
	Name       string        `json:"name,omitempty"`
	Suppressed bool          `json:"suppressed,omitempty"`
	Transform  *types.Matrix `json:"transform,omitempty"`
	Previous   *types.Matrix `json:"previous,omitempty"`
}
