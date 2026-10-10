// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Attributes is the add-in attribute-set operation group (#155): named, typed values an add-in
// attaches to a document and that persist with it — the sanctioned way to store add-in data and
// tag the model. An attribute lives in a named set (a namespace, conventionally the add-in id).
type Attributes struct{ c *Client }

// Attributes returns the attribute-set operation group.
func (c *Client) Attributes() Attributes { return Attributes{c} }

// Set creates or replaces the named document-scoped attribute in the set with the typed value.
func (a Attributes) Set(document uint64, set, name string, value types.Variant) (wire.AttributeResult, error) {
	return a.SetOn(document, "", set, name, value)
}

// SetOn creates or replaces the named attribute in the set on the given target (a body/face/edge
// reference key from body.list or model.referenceKeys; empty target = the document itself) with
// the typed value. Anchoring by reference key lets the tag survive recompute.
//
// mcp:tool set_attribute
// mcp:summary Store a typed value under a name in a named set on a document, optionally anchored to an entity by reference key.
func (a Attributes) SetOn(document uint64, target, set, name string, value types.Variant) (wire.AttributeResult, error) {
	return call[wire.AttributeResult](a.c, wire.MethodAttributesSet, wire.SetAttributeArgs{Document: document, Set: set, Name: name, Value: value, Target: target})
}

// Get reads one document-scoped attribute by set and name; Found is false when it is absent.
func (a Attributes) Get(document uint64, set, name string) (wire.AttributeResult, error) {
	return a.GetOn(document, "", set, name)
}

// GetOn reads one attribute by set and name on the given target (empty = the document itself);
// Found is false when it is absent.
//
// mcp:tool get_attribute
// mcp:summary Read a stored attribute by set and name on a document (optionally anchored to an entity by reference key).
func (a Attributes) GetOn(document uint64, target, set, name string) (wire.AttributeResult, error) {
	return call[wire.AttributeResult](a.c, wire.MethodAttributesGet, wire.GetAttributeArgs{Document: document, Set: set, Name: name, Target: target})
}

// List returns the document-scoped attributes, or only those in set when set is non-empty.
//
// mcp:tool list_attributes
// mcp:summary List the document-scoped attributes on a document (optionally filtered to one set).
func (a Attributes) List(document uint64, set string) (wire.ListAttributesResult, error) {
	return call[wire.ListAttributesResult](a.c, wire.MethodAttributesList, wire.ListAttributesArgs{Document: document, Set: set})
}

// ListOn returns the attributes anchored to the given target (empty = the document itself),
// optionally filtered to one set.
func (a Attributes) ListOn(document uint64, target, set string) (wire.ListAttributesResult, error) {
	return call[wire.ListAttributesResult](a.c, wire.MethodAttributesList, wire.ListAttributesArgs{Document: document, Set: set, Target: target})
}

// ListAll returns every attribute on every target in the document (each result carries its own
// Target), optionally filtered to one set. Use it to read back all of an add-in's per-entity tags.
func (a Attributes) ListAll(document uint64, set string) (wire.ListAttributesResult, error) {
	return call[wire.ListAttributesResult](a.c, wire.MethodAttributesList, wire.ListAttributesArgs{Document: document, Set: set, AllTargets: true})
}

// ListSets returns the document's attribute set names, sorted.
//
// mcp:tool list_attribute_sets
// mcp:summary List the attribute set names on a document.
func (a Attributes) ListSets(document uint64) (wire.ListAttributeSetsResult, error) {
	return call[wire.ListAttributeSetsResult](a.c, wire.MethodAttributesListSets, wire.ListAttributeSetsArgs{Document: document})
}

// Delete removes the named document-scoped attribute in the set, or the whole set when name is
// empty; the result reports how many attributes were removed.
func (a Attributes) Delete(document uint64, set, name string) (wire.DeleteAttributeResult, error) {
	return a.DeleteOn(document, "", set, name)
}

// DeleteOn removes the named attribute in the set on the given target (empty = the document
// itself), or the whole set on that target when name is empty; the result reports how many were
// removed.
//
// mcp:tool delete_attribute
// mcp:summary Delete an attribute (or a whole set when name is empty) on a document, optionally anchored to an entity by reference key.
func (a Attributes) DeleteOn(document uint64, target, set, name string) (wire.DeleteAttributeResult, error) {
	return call[wire.DeleteAttributeResult](a.c, wire.MethodAttributesDelete, wire.DeleteAttributeArgs{Document: document, Set: set, Name: name, Target: target})
}

// Find locates the open documents carrying an attribute in set; restrict to a name when non-empty.
//
// mcp:tool find_by_attribute
// mcp:summary Find the open documents carrying an attribute in a given set (optionally by name).
func (a Attributes) Find(set, name string) (wire.FindByAttributeResult, error) {
	return call[wire.FindByAttributeResult](a.c, wire.MethodAttributesFind, wire.FindByAttributeArgs{Set: set, Name: name})
}
