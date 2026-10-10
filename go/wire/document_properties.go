// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Document iProperties (#156): a document's metadata — the standard sets (Summary
// Information, Document Summary Information, Design Tracking Properties) plus user-defined
// custom properties — that feed BOMs (#718), drawing title blocks, and PDM. Each property is
// a typed named value in a named set; the value rides the shared [types.Variant] currency.

// PropertyInfo is the JSON shape of one document property: its set, name, and typed value.
// FromParameter names the model parameter a property was exposed from (the iProperty bridge),
// empty for a directly authored property.
type PropertyInfo struct {
	Set           string        `json:"set"`
	Name          string        `json:"name"`
	Value         types.Variant `json:"value"`
	FromParameter string        `json:"fromParameter,omitempty"`
}

// ListPropertiesArgs is the request of [MethodDocumentsListProperties]: the document whose
// properties to enumerate (by session id from documents.list).
type ListPropertiesArgs struct {
	Document uint64 `json:"document"`
}

// ListPropertiesResult is the response of [MethodDocumentsListProperties]: every property of
// the document, across all its sets, in set then insertion order.
type ListPropertiesResult struct {
	Properties []PropertyInfo `json:"properties"`
}

// GetPropertyArgs is the request of [MethodDocumentsGetProperty]: address one property by its
// set and name on the document.
type GetPropertyArgs struct {
	Document uint64 `json:"document"`
	Set      string `json:"set"`
	Name     string `json:"name"`
}

// SetPropertyArgs is the request of [MethodDocumentsSetProperty]: create or replace the named
// property in the set with the typed value (the set is created if it is a custom set name not
// among the standard sets). Address by document session id.
type SetPropertyArgs struct {
	Document uint64        `json:"document"`
	Set      string        `json:"set"`
	Name     string        `json:"name"`
	Value    types.Variant `json:"value"`
}

// PropertyResult is the response of [MethodDocumentsGetProperty] / [MethodDocumentsSetProperty]:
// the addressed property's current state.
type PropertyResult struct {
	Property PropertyInfo `json:"property"`
}
