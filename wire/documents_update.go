// SPDX-License-Identifier: Apache-2.0

package wire

// Document update/rebuild (Oblikovati/Oblikovati#139). After batch edits a client needs to
// force a recompute and learn the result. update recomputes only the out-of-date features
// (Inventor PartDocument.Update); rebuild recomputes everything as if all driving entities were
// dirtied (PartDocument.Rebuild); requiresUpdate is the read-only "has stale entities" flag.

// UpdateDocumentArgs is the request of [MethodDocumentsUpdate] and [MethodDocumentsRebuild].
// AcceptErrorsAndContinue mirrors Update2/Rebuild2: when true the call succeeds and reports any
// sick features in [UpdateDocumentResult.Errors]; when false (the default) a sick feature makes
// the call fail instead.
type UpdateDocumentArgs struct {
	AcceptErrorsAndContinue bool `json:"acceptErrorsAndContinue,omitempty"`
}

// FeatureError is one feature that failed to evaluate during an update/rebuild: its stable id
// (from model.tree), display name, kind, and the reason it is sick.
type FeatureError struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Kind    string `json:"kind"`
	Message string `json:"message,omitempty"`
}

// UpdateDocumentResult is the response of [MethodDocumentsUpdate] and [MethodDocumentsRebuild]:
// whether the document still requires an update (false after a successful recompute) and the
// features that ended up sick (only populated when AcceptErrorsAndContinue was set).
type UpdateDocumentResult struct {
	RequiresUpdate bool           `json:"requiresUpdate"`
	Errors         []FeatureError `json:"errors,omitempty"`
}

// RequiresUpdateResult is the response of [MethodDocumentsRequiresUpdate]: the read-only flag
// telling a client the active document has out-of-date features a recompute would change.
type RequiresUpdateResult struct {
	RequiresUpdate bool `json:"requiresUpdate"`
}
