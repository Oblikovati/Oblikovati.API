// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// The add-in data registry on documents (M03-F10).

// Interests returns a document's registered interest records.
//
// mcp:tool documents_list_interests
// mcp:summary Returns a document's registered interest records.
func (d Documents) Interests(id uint64) (wire.ListDocumentInterestsResult, error) {
	return call[wire.ListDocumentInterestsResult](d.c, wire.MethodDocumentsListInterests, wire.ListDocumentInterestsArgs{Document: id})
}

// AddInterest registers (or updates) an interest record on a document.
//
//	c.Documents().AddInterest(doc.ID, types.DocumentInterestRecord{
//	    ClientID: "com.x.toolpaths", Name: "toolpath-recipes",
//	    InterestType: types.Interested, DataVersion: 2,
//	})
//
// mcp:tool documents_add_interest
// mcp:summary Registers (or updates) an interest record on a document.
func (d Documents) AddInterest(id uint64, record types.DocumentInterestRecord) (wire.OKResult, error) {
	args := wire.AddDocumentInterestArgs{Document: id, Interest: record}
	return call[wire.OKResult](d.c, wire.MethodDocumentsAddInterest, args)
}

// RemoveInterest deletes the (clientID, name) interest record.
//
// mcp:tool documents_remove_interest
// mcp:summary Deletes the (clientID, name) interest record.
func (d Documents) RemoveInterest(id uint64, clientID, name string) (wire.OKResult, error) {
	args := wire.RemoveDocumentInterestArgs{Document: id, ClientID: clientID, Name: name}
	return call[wire.OKResult](d.c, wire.MethodDocumentsRemoveInterest, args)
}

// HasInterest reports whether any interest record's client id or name matches
// client — discovery without enumerating.
//
// mcp:tool documents_has_interest
// mcp:summary Reports whether any interest record's client id or name matches client — discovery without enumerating.
func (d Documents) HasInterest(id uint64, client string) (wire.HasDocumentInterestResult, error) {
	args := wire.HasDocumentInterestArgs{Document: id, Client: client}
	return call[wire.HasDocumentInterestResult](d.c, wire.MethodDocumentsHasInterest, args)
}
