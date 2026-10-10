// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The canonical collection encodings (M00-F05, #601). Every options-bag,
// object-list or paged-query parameter on the wire uses THESE shapes — they
// are defined once here and never re-declared per method (ADR-0018).

// NameValueEntry is one named bag value: the canonical JSON shape of a
// NameValueMap element, {"name": …, "value": {"type": …, "value": …}}.
type NameValueEntry struct {
	Name  string        `json:"name"`
	Value types.Variant `json:"value"`
}

// NameValueMap is the ordered options/context bag: entry order is meaningful
// and preserved (insert-before/after semantics live on the client collection).
type NameValueMap []NameValueEntry

// ObjectRefList is the ordered object-reference list accepted wherever the
// reference API takes an ObjectCollection (feature inputs, pattern element
// lists, graphics inputs).
type ObjectRefList []types.ObjectRef

// KeyedObjectRef is one string-keyed object reference: the canonical element
// of an ObjectCollectionByVariant over the wire.
type KeyedObjectRef struct {
	Key string          `json:"key"`
	Ref types.ObjectRef `json:"ref"`
}

// KeyedObjectRefList is the ordered, string-keyed object-reference list.
type KeyedObjectRefList []KeyedObjectRef

// EnumeratorPage is the canonical paged-result envelope for query methods that
// enumerate large sets: embed it next to the items slice in the result DTO.
// NextCursor is the opaque position to pass back to fetch the following page;
// Done reports that no further pages exist. (logs.tail predates this envelope
// and keeps its frozen NextSeq spelling.)
type EnumeratorPage struct {
	NextCursor uint64 `json:"nextCursor,omitempty"`
	Done       bool   `json:"done,omitempty"`
}
