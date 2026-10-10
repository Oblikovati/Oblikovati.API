// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The core utility value collections (M00-F05, #601): the options/object-set
// currency of the options, events and graphics surfaces. The concrete
// implementations are ownerless pure data and live in api/client (built by
// [TransientObjects]); the host converts to its typed structs at the boundary.

// NameValueMap is the ordered string→variant options/context bag. Names are
// unique; entry order is meaningful and preserved.
type NameValueMap interface {
	// Count returns the number of entries.
	Count() int
	// NameAt returns the name at the 0-based index, erroring out of range.
	NameAt(index int) (string, error)
	// ValueAt returns the value at the 0-based index, erroring out of range.
	ValueAt(index int) (types.Variant, error)
	// Value returns the value for name, ok=false when absent.
	Value(name string) (types.Variant, bool)
	// Set adds the entry, or replaces the value when name already exists.
	Set(name string, value types.Variant)
	// Insert places a new entry before (or after) the entry at targetIndex,
	// erroring when the name already exists or the index is out of range.
	Insert(name string, value types.Variant, targetIndex int, before bool) error
	// Remove deletes the named entry, reporting whether it existed.
	Remove(name string) bool
	// Clear removes every entry.
	Clear()
	// Names returns the entry names in order.
	Names() []string
}

// ObjectsEnumerator is the read-only ordered view over object references —
// the result shape of query methods.
type ObjectsEnumerator interface {
	// Count returns the number of references.
	Count() int
	// At returns the reference at the 0-based index, erroring out of range.
	At(index int) (types.ObjectRef, error)
	// Refs returns the references in order (a copy; mutating it is safe).
	Refs() []types.ObjectRef
}

// ObjectCollection is the ordered, mutable object-reference collection
// accepted wherever the reference API takes an ObjectCollection.
type ObjectCollection interface {
	ObjectsEnumerator
	// Add appends a reference (duplicates are allowed, as in the reference).
	Add(ref types.ObjectRef)
	// RemoveAt deletes the entry at the 0-based index, erroring out of range.
	RemoveAt(index int) error
	// RemoveRef deletes the first entry equal to ref, reporting whether one existed.
	RemoveRef(ref types.ObjectRef) bool
	// Clear removes every reference.
	Clear()
}

// ObjectsEnumeratorByVariant is the read-only ordered view over string-keyed
// object references.
type ObjectsEnumeratorByVariant interface {
	// Count returns the number of entries.
	Count() int
	// KeyAt returns the key at the 0-based index, erroring out of range.
	KeyAt(index int) (string, error)
	// At returns the reference at the 0-based index, erroring out of range.
	At(index int) (types.ObjectRef, error)
	// ByKey returns the reference for key, ok=false when absent.
	ByKey(key string) (types.ObjectRef, bool)
}

// ObjectCollectionByVariant is the ordered, mutable, string-keyed
// object-reference collection. Keys are unique.
type ObjectCollectionByVariant interface {
	ObjectsEnumeratorByVariant
	// Add appends a keyed reference, erroring when the key already exists.
	Add(key string, ref types.ObjectRef) error
	// Remove deletes the keyed entry, reporting whether it existed.
	Remove(key string) bool
	// RemoveAt deletes the entry at the 0-based index, erroring out of range.
	RemoveAt(index int) error
	// Clear removes every entry.
	Clear()
}

// TransientObjects is the discoverable construction point for the utility
// collections, mirroring the reference factory so add-in code ports naturally
// (geometry construction is [TransientGeometry]).
type TransientObjects interface {
	// CreateNameValueMap returns an empty ordered options bag.
	CreateNameValueMap() NameValueMap
	// CreateObjectCollection returns a collection seeded with refs, in order.
	CreateObjectCollection(refs ...types.ObjectRef) ObjectCollection
	// CreateObjectCollectionByVariant returns an empty keyed collection.
	CreateObjectCollectionByVariant() ObjectCollectionByVariant
}
