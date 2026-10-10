// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"fmt"

	"oblikovati.org/api/contract"
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// ObjectCollectionByVariant is the ordered, string-keyed object-reference
// collection (M00-F05, #601). Keys are unique. Its JSON form is the canonical
// [wire.KeyedObjectRefList].
type ObjectCollectionByVariant struct {
	entries wire.KeyedObjectRefList
}

var _ contract.ObjectCollectionByVariant = (*ObjectCollectionByVariant)(nil)

// Count returns the number of entries.
func (c *ObjectCollectionByVariant) Count() int { return len(c.entries) }

// KeyAt returns the key at the 0-based index, erroring out of range.
func (c *ObjectCollectionByVariant) KeyAt(index int) (string, error) {
	e, err := boundsCheckedAt(c.entries, index, "ObjectCollectionByVariant")
	return e.Key, err
}

// At returns the reference at the 0-based index, erroring out of range.
func (c *ObjectCollectionByVariant) At(index int) (types.ObjectRef, error) {
	e, err := boundsCheckedAt(c.entries, index, "ObjectCollectionByVariant")
	return e.Ref, err
}

// ByKey returns the reference for key, ok=false when absent.
func (c *ObjectCollectionByVariant) ByKey(key string) (types.ObjectRef, bool) {
	if i := c.indexOf(key); i >= 0 {
		return c.entries[i].Ref, true
	}
	return types.ObjectRef{}, false
}

// Add appends a keyed reference, erroring when the key already exists.
func (c *ObjectCollectionByVariant) Add(key string, ref types.ObjectRef) error {
	if c.indexOf(key) >= 0 {
		return fmt.Errorf("client: ObjectCollectionByVariant already has an entry keyed %q", key)
	}
	c.entries = append(c.entries, wire.KeyedObjectRef{Key: key, Ref: ref})
	return nil
}

// Remove deletes the keyed entry, reporting whether it existed.
func (c *ObjectCollectionByVariant) Remove(key string) bool {
	i := c.indexOf(key)
	if i < 0 {
		return false
	}
	c.entries = append(c.entries[:i], c.entries[i+1:]...)
	return true
}

// RemoveAt deletes the entry at the 0-based index, erroring out of range.
func (c *ObjectCollectionByVariant) RemoveAt(index int) error {
	if _, err := boundsCheckedAt(c.entries, index, "ObjectCollectionByVariant"); err != nil {
		return err
	}
	c.entries = append(c.entries[:index], c.entries[index+1:]...)
	return nil
}

// Clear removes every entry.
func (c *ObjectCollectionByVariant) Clear() { c.entries = nil }

// Entries returns the canonical wire form (a copy; mutating it is safe).
func (c *ObjectCollectionByVariant) Entries() wire.KeyedObjectRefList {
	return append(wire.KeyedObjectRefList(nil), c.entries...)
}

// MarshalJSON encodes the canonical keyed-reference list.
func (c *ObjectCollectionByVariant) MarshalJSON() ([]byte, error) { return json.Marshal(c.entries) }

// UnmarshalJSON decodes the canonical keyed-reference list.
func (c *ObjectCollectionByVariant) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &c.entries)
}

// indexOf returns the position of key, or -1.
func (c *ObjectCollectionByVariant) indexOf(key string) int {
	return indexOfFunc(c.entries, func(e wire.KeyedObjectRef) bool { return e.Key == key })
}
