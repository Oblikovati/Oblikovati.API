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

// Set creates or replaces the named attribute in the set on the document with the typed value.
//
// mcp:tool set_attribute
// mcp:summary Store a typed value (the add-in's own data) under a name in a named set on a document.
func (a Attributes) Set(document uint64, set, name string, value types.Variant) (wire.AttributeResult, error) {
	var r wire.AttributeResult
	return r, a.c.call(wire.MethodAttributesSet, wire.SetAttributeArgs{Document: document, Set: set, Name: name, Value: value}, &r)
}

// Get reads one attribute by set and name on the document; Found is false when it is absent.
//
// mcp:tool get_attribute
// mcp:summary Read a stored attribute by set and name on a document.
func (a Attributes) Get(document uint64, set, name string) (wire.AttributeResult, error) {
	var r wire.AttributeResult
	return r, a.c.call(wire.MethodAttributesGet, wire.GetAttributeArgs{Document: document, Set: set, Name: name}, &r)
}

// List returns every attribute on the document, or only those in set when set is non-empty.
//
// mcp:tool list_attributes
// mcp:summary List the attributes on a document (optionally filtered to one set).
func (a Attributes) List(document uint64, set string) (wire.ListAttributesResult, error) {
	var r wire.ListAttributesResult
	return r, a.c.call(wire.MethodAttributesList, wire.ListAttributesArgs{Document: document, Set: set}, &r)
}

// ListSets returns the document's attribute set names, sorted.
//
// mcp:tool list_attribute_sets
// mcp:summary List the attribute set names on a document.
func (a Attributes) ListSets(document uint64) (wire.ListAttributeSetsResult, error) {
	var r wire.ListAttributeSetsResult
	return r, a.c.call(wire.MethodAttributesListSets, wire.ListAttributeSetsArgs{Document: document}, &r)
}

// Delete removes the named attribute in the set, or the whole set when name is empty; the result
// reports how many attributes were removed.
//
// mcp:tool delete_attribute
// mcp:summary Delete an attribute (or a whole set when name is empty) on a document.
func (a Attributes) Delete(document uint64, set, name string) (wire.DeleteAttributeResult, error) {
	var r wire.DeleteAttributeResult
	return r, a.c.call(wire.MethodAttributesDelete, wire.DeleteAttributeArgs{Document: document, Set: set, Name: name}, &r)
}

// Find locates the open documents carrying an attribute in set; restrict to a name when non-empty.
//
// mcp:tool find_by_attribute
// mcp:summary Find the open documents carrying an attribute in a given set (optionally by name).
func (a Attributes) Find(set, name string) (wire.FindByAttributeResult, error) {
	var r wire.FindByAttributeResult
	return r, a.c.call(wire.MethodAttributesFind, wire.FindByAttributeArgs{Set: set, Name: name}, &r)
}
