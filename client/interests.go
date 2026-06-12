// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// The add-in data registry on documents (M03-F10).

// Interests returns a document's registered interest records.
func (d Documents) Interests(id uint64) (wire.ListDocumentInterestsResult, error) {
	var r wire.ListDocumentInterestsResult
	return r, d.c.call(wire.MethodDocumentsListInterests, wire.ListDocumentInterestsArgs{Document: id}, &r)
}

// AddInterest registers (or updates) an interest record on a document.
//
//	c.Documents().AddInterest(doc.ID, types.DocumentInterestRecord{
//	    ClientID: "com.x.toolpaths", Name: "toolpath-recipes",
//	    InterestType: types.Interested, DataVersion: 2,
//	})
func (d Documents) AddInterest(id uint64, record types.DocumentInterestRecord) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.AddDocumentInterestArgs{Document: id, Interest: record}
	return r, d.c.call(wire.MethodDocumentsAddInterest, args, &r)
}

// RemoveInterest deletes the (clientID, name) interest record.
func (d Documents) RemoveInterest(id uint64, clientID, name string) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.RemoveDocumentInterestArgs{Document: id, ClientID: clientID, Name: name}
	return r, d.c.call(wire.MethodDocumentsRemoveInterest, args, &r)
}

// HasInterest reports whether any interest record's client id or name matches
// client — discovery without enumerating.
func (d Documents) HasInterest(id uint64, client string) (wire.HasDocumentInterestResult, error) {
	var r wire.HasDocumentInterestResult
	args := wire.HasDocumentInterestArgs{Document: id, Client: client}
	return r, d.c.call(wire.MethodDocumentsHasInterest, args, &r)
}
