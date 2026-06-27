// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Add-in attribute sets (#155): named, typed values an add-in attaches to a document and that
// persist with it — the sanctioned way for an add-in to store its own data and tag the model.
// An attribute lives in a named SET (a namespace, conventionally the add-in's id) under a NAME,
// and carries a [types.Variant] value (integer/double/string/bytes/boolean).
//
// An attribute is anchored to a TARGET: the document itself (the default, an empty target) or a
// specific entity addressed by its persistent reference key — a body, face, edge, or vertex key
// as returned by body.list (BodyInfo.Key) or model.referenceKeys (TopologyRef.Key). Anchoring by
// reference key (not by index) is what lets a tag survive recompute: after the B-rep is rebuilt
// the same lineage re-mints an equal key, so the same attributes are found again. The target
// rides every DTO; omit it for document-scoped attributes.

// AttributeInfo is the JSON shape of one attribute: the set it lives in, its name, its value, and
// the target it is anchored to (empty for the document itself).
type AttributeInfo struct {
	Set    string        `json:"set"`
	Name   string        `json:"name"`
	Value  types.Variant `json:"value"`
	Target string        `json:"target,omitempty"`
}

// SetAttributeArgs is the request of [MethodAttributesSet]: create or replace the named attribute
// in the named set on the document (by session id from documents.list) with the typed value,
// anchored to Target (empty = the document itself). The set is created on first use.
type SetAttributeArgs struct {
	Document uint64        `json:"document"`
	Set      string        `json:"set"`
	Name     string        `json:"name"`
	Value    types.Variant `json:"value"`
	Target   string        `json:"target,omitempty"`
}

// GetAttributeArgs is the request of [MethodAttributesGet]: address one attribute by its set and
// name on the document, anchored to Target (empty = the document itself).
type GetAttributeArgs struct {
	Document uint64 `json:"document"`
	Set      string `json:"set"`
	Name     string `json:"name"`
	Target   string `json:"target,omitempty"`
}

// AttributeResult is the response of [MethodAttributesGet] / [MethodAttributesSet]: the addressed
// attribute and whether it was found (false for a get that missed).
type AttributeResult struct {
	Attribute AttributeInfo `json:"attribute"`
	Found     bool          `json:"found"`
}

// ListAttributesArgs is the request of [MethodAttributesList]: every attribute on the document, or
// only those in Set when it is non-empty. By default it lists the document-scoped attributes;
// set Target to list a specific entity's attributes, or set AllTargets to list every attribute on
// every target (each carrying its Target in the result).
type ListAttributesArgs struct {
	Document   uint64 `json:"document"`
	Set        string `json:"set,omitempty"`
	Target     string `json:"target,omitempty"`
	AllTargets bool   `json:"allTargets,omitempty"`
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
// set, or the whole set when Name is empty, on the target (empty = the document itself).
type DeleteAttributeArgs struct {
	Document uint64 `json:"document"`
	Set      string `json:"set"`
	Name     string `json:"name,omitempty"`
	Target   string `json:"target,omitempty"`
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
