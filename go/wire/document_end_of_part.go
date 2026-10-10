// SPDX-License-Identifier: Apache-2.0

package wire

// Part end-of-part marker (#141): the rollback marker that controls how far down a part's feature
// program the model evaluates. document.getEndOfPart inspects it; document.setEndOfPart moves it
// (before/after a feature is expressed as the feature's index). It addresses the ACTIVE part — the
// marker is part-document state, so no document id rides the request.

// EndOfPartResult is the reply of [MethodDocumentGetEndOfPart] and [MethodDocumentSetEndOfPart]: the
// marker position (-1 at the end) and whether the part is currently rolled back.
type EndOfPartResult struct {
	Position   int  `json:"position"`
	RolledBack bool `json:"rolledBack"`
}

// SetEndOfPartArgs is the request of [MethodDocumentSetEndOfPart]: move the marker to Position
// (a negative index restores it to the end, re-including every feature).
type SetEndOfPartArgs struct {
	Position int `json:"position"`
}
