// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The add-in data registry on documents (M03-F10, Oblikovati/Oblikovati#611):
// interests declare "client X has data in / depends on this document", are
// persisted with the document, and are readable without loading the add-in —
// the sanctioned way for an add-in to mark documents it has augmented.

// ListDocumentInterestsArgs is the request of [MethodDocumentsListInterests].
type ListDocumentInterestsArgs struct {
	Document uint64 `json:"document"`
}

// ListDocumentInterestsResult is the response of [MethodDocumentsListInterests].
type ListDocumentInterestsResult struct {
	Interests []types.DocumentInterestRecord `json:"interests"`
}

// AddDocumentInterestArgs is the request of [MethodDocumentsAddInterest]: an
// existing (ClientID, Name) record is updated in place.
type AddDocumentInterestArgs struct {
	Document uint64                       `json:"document"`
	Interest types.DocumentInterestRecord `json:"interest"`
}

// RemoveDocumentInterestArgs is the request of [MethodDocumentsRemoveInterest]:
// delete the record identified by (ClientID, Name).
type RemoveDocumentInterestArgs struct {
	Document uint64 `json:"document"`
	ClientID string `json:"clientId"`
	Name     string `json:"name"`
}

// HasDocumentInterestArgs is the request of [MethodDocumentsHasInterest]:
// Client matches either a record's ClientID or its Name — the cheap discovery
// probe ("does anyone hold data here?").
type HasDocumentInterestArgs struct {
	Document uint64 `json:"document"`
	Client   string `json:"client"`
}

// HasDocumentInterestResult is the response of [MethodDocumentsHasInterest].
type HasDocumentInterestResult struct {
	HasInterest bool `json:"hasInterest"`
}
