// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"fmt"

	"oblikovati.org/api/contract"
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// The concrete utility value collections (M00-F05, #601). They live HERE, in
// the Apache-2.0 module, because they are ownerless pure data an add-in
// assembles locally and ships as one wire payload — the same split ADR-0018's
// geometry addendum records for the transient geometry values. No transport
// is involved; [TransientObjects] needs no [Caller].

// TransientObjects is the factory for the utility collections, mirroring the
// reference factory so ported add-in code reads naturally.
//
//	opts := client.TransientObjects{}.CreateNameValueMap()
//	opts.Set("tolerance", types.UnitVariant(0.1, "mm"))
type TransientObjects struct{}

var _ contract.TransientObjects = TransientObjects{}

// CreateNameValueMap returns an empty ordered options bag.
func (TransientObjects) CreateNameValueMap() contract.NameValueMap { return &NameValueMap{} }

// CreateObjectCollection returns a collection seeded with refs, in order.
func (TransientObjects) CreateObjectCollection(refs ...types.ObjectRef) contract.ObjectCollection {
	return &ObjectCollection{refs: append([]types.ObjectRef(nil), refs...)}
}

// CreateObjectCollectionByVariant returns an empty keyed collection.
func (TransientObjects) CreateObjectCollectionByVariant() contract.ObjectCollectionByVariant {
	return &ObjectCollectionByVariant{}
}

// NameValueMap is the ordered string→variant options bag. Its JSON form is the
// canonical [wire.NameValueMap] entry list.
type NameValueMap struct {
	entries wire.NameValueMap
}

var _ contract.NameValueMap = (*NameValueMap)(nil)

// Count returns the number of entries.
func (m *NameValueMap) Count() int { return len(m.entries) }

// NameAt returns the name at the 0-based index, erroring out of range.
func (m *NameValueMap) NameAt(index int) (string, error) {
	e, err := boundsCheckedAt(m.entries, index, "NameValueMap")
	return e.Name, err
}

// ValueAt returns the value at the 0-based index, erroring out of range.
func (m *NameValueMap) ValueAt(index int) (types.Variant, error) {
	e, err := boundsCheckedAt(m.entries, index, "NameValueMap")
	return e.Value, err
}

// Value returns the value for name, ok=false when absent.
func (m *NameValueMap) Value(name string) (types.Variant, bool) {
	if i := m.indexOf(name); i >= 0 {
		return m.entries[i].Value, true
	}
	return types.Variant{}, false
}

// Set adds the entry, or replaces the value when name already exists.
func (m *NameValueMap) Set(name string, value types.Variant) {
	if i := m.indexOf(name); i >= 0 {
		m.entries[i].Value = value
		return
	}
	m.entries = append(m.entries, wire.NameValueEntry{Name: name, Value: value})
}

// Insert places a new entry before (or after) the entry at targetIndex,
// erroring when the name already exists or the index is out of range.
func (m *NameValueMap) Insert(name string, value types.Variant, targetIndex int, before bool) error {
	if m.indexOf(name) >= 0 {
		return fmt.Errorf("client: NameValueMap already has an entry named %q", name)
	}
	if _, err := boundsCheckedAt(m.entries, targetIndex, "NameValueMap insert"); err != nil {
		return err
	}
	at := targetIndex
	if !before {
		at++
	}
	entry := wire.NameValueEntry{Name: name, Value: value}
	m.entries = append(m.entries[:at], append(wire.NameValueMap{entry}, m.entries[at:]...)...)
	return nil
}

// Remove deletes the named entry, reporting whether it existed.
func (m *NameValueMap) Remove(name string) bool {
	i := m.indexOf(name)
	if i < 0 {
		return false
	}
	m.entries = append(m.entries[:i], m.entries[i+1:]...)
	return true
}

// Clear removes every entry.
func (m *NameValueMap) Clear() { m.entries = nil }

// Names returns the entry names in order.
func (m *NameValueMap) Names() []string {
	names := make([]string, len(m.entries))
	for i, e := range m.entries {
		names[i] = e.Name
	}
	return names
}

// Entries returns the canonical wire form (a copy; mutating it is safe).
func (m *NameValueMap) Entries() wire.NameValueMap {
	return append(wire.NameValueMap(nil), m.entries...)
}

// MarshalJSON encodes the canonical entry list.
func (m *NameValueMap) MarshalJSON() ([]byte, error) { return json.Marshal(m.entries) }

// UnmarshalJSON decodes the canonical entry list.
func (m *NameValueMap) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &m.entries) }

// indexOf returns the position of name, or -1.
func (m *NameValueMap) indexOf(name string) int {
	return indexOfFunc(m.entries, func(e wire.NameValueEntry) bool { return e.Name == name })
}

// ObjectCollection is the ordered, mutable object-reference collection. Its
// JSON form is the canonical [wire.ObjectRefList].
type ObjectCollection struct {
	refs []types.ObjectRef
}

var _ contract.ObjectCollection = (*ObjectCollection)(nil)

// Count returns the number of references.
func (c *ObjectCollection) Count() int { return len(c.refs) }

// At returns the reference at the 0-based index, erroring out of range.
func (c *ObjectCollection) At(index int) (types.ObjectRef, error) {
	return boundsCheckedAt(c.refs, index, "ObjectCollection")
}

// Refs returns the references in order (a copy; mutating it is safe).
func (c *ObjectCollection) Refs() []types.ObjectRef {
	return append([]types.ObjectRef(nil), c.refs...)
}

// Add appends a reference (duplicates are allowed, as in the reference).
func (c *ObjectCollection) Add(ref types.ObjectRef) { c.refs = append(c.refs, ref) }

// RemoveAt deletes the entry at the 0-based index, erroring out of range.
func (c *ObjectCollection) RemoveAt(index int) error {
	if _, err := boundsCheckedAt(c.refs, index, "ObjectCollection"); err != nil {
		return err
	}
	c.refs = append(c.refs[:index], c.refs[index+1:]...)
	return nil
}

// RemoveRef deletes the first entry equal to ref, reporting whether one existed.
func (c *ObjectCollection) RemoveRef(ref types.ObjectRef) bool {
	i := indexOfFunc(c.refs, func(r types.ObjectRef) bool { return r == ref })
	if i < 0 {
		return false
	}
	c.refs = append(c.refs[:i], c.refs[i+1:]...)
	return true
}

// Clear removes every reference.
func (c *ObjectCollection) Clear() { c.refs = nil }

// MarshalJSON encodes the canonical reference list.
func (c *ObjectCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(wire.ObjectRefList(c.refs))
}

// UnmarshalJSON decodes the canonical reference list.
func (c *ObjectCollection) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &c.refs) }
