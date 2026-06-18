// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Add-in attribute sets (#155): named, typed values an add-in attaches to a document and that
// persist with it — the sanctioned way for an add-in to store its own data and tag the model.
// An attribute lives in a named SET (a namespace, conventionally the add-in's id) under a NAME,
// and carries a [types.Variant] value (integer/double/string/bytes/boolean). This first surface
// targets the document itself; finer targets (features, bodies, sketch entities) are a follow-up
// that rides the same DTOs with an added target selector.

// AttributeInfo is the JSON shape of one attribute: the set it lives in, its name, and its value.
type AttributeInfo struct {
	Set   string        `json:"set"`
	Name  string        `json:"name"`
	Value types.Variant `json:"value"`
}

// SetAttributeArgs is the request of [MethodAttributesSet]: create or replace the named attribute
// in the named set on the document (by session id from documents.list) with the typed value. The
// set is created on first use.
type SetAttributeArgs struct {
	Document uint64        `json:"document"`
	Set      string        `json:"set"`
	Name     string        `json:"name"`
	Value    types.Variant `json:"value"`
}

// GetAttributeArgs is the request of [MethodAttributesGet]: address one attribute by its set and
// name on the document.
type GetAttributeArgs struct {
	Document uint64 `json:"document"`
	Set      string `json:"set"`
	Name     string `json:"name"`
}

// AttributeResult is the response of [MethodAttributesGet] / [MethodAttributesSet]: the addressed
// attribute and whether it was found (false for a get that missed).
type AttributeResult struct {
	Attribute AttributeInfo `json:"attribute"`
	Found     bool          `json:"found"`
}

// ListAttributesArgs is the request of [MethodAttributesList]: every attribute on the document, or
// only those in Set when it is non-empty.
type ListAttributesArgs struct {
	Document uint64 `json:"document"`
	Set      string `json:"set,omitempty"`
}

// ListAttributesResult is the response of [MethodAttributesList]: the matching attributes in set
// then insertion order.
type ListAttributesResult struct {
	Attributes []AttributeInfo `json:"attributes"`
}

// ListAttributeSetsArgs is the request of [MethodAttributesListSets]: the document whose attribute
// set names to enumerate.
type ListAttributeSetsArgs struct {
	Document uint64 `json:"document"`
}

// ListAttributeSetsResult is the response of [MethodAttributesListSets]: the set names, sorted.
type ListAttributeSetsResult struct {
	Sets []string `json:"sets"`
}

// DeleteAttributeArgs is the request of [MethodAttributesDelete]: remove the named attribute in the
// set, or the whole set when Name is empty.
type DeleteAttributeArgs struct {
	Document uint64 `json:"document"`
	Set      string `json:"set"`
	Name     string `json:"name,omitempty"`
}

// DeleteAttributeResult is the response of [MethodAttributesDelete]: how many attributes were
// removed (0 when nothing matched).
type DeleteAttributeResult struct {
	Removed int `json:"removed"`
}

// FindByAttributeArgs is the request of [MethodAttributesFind]: locate the open documents carrying
// an attribute in Set; restrict to a given Name when it is non-empty.
type FindByAttributeArgs struct {
	Set  string `json:"set"`
	Name string `json:"name,omitempty"`
}

// AttributeMatch is one hit of [MethodAttributesFind]: the document session id and the matching
// attribute on it.
type AttributeMatch struct {
	Document  uint64        `json:"document"`
	Attribute AttributeInfo `json:"attribute"`
}

// FindByAttributeResult is the response of [MethodAttributesFind]: every match across open documents.
type FindByAttributeResult struct {
	Matches []AttributeMatch `json:"matches"`
}
