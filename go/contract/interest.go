// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The add-in data registry on documents (M03-F10, Oblikovati/Oblikovati#611).
// Interests are the discovery/versioning layer; attribute sets are the
// payload layer.

// DocumentInterest is one registered interest record on a document.
type DocumentInterest interface {
	// ClientID returns the owning client's id.
	ClientID() string
	// Name returns the interest's name; (ClientID, Name) is its identity.
	Name() string
	// InterestType returns the interest's strength.
	InterestType() types.DocumentInterestType
	// DataVersion returns the client-managed migration version, 0 for a
	// non-migrating interest.
	DataVersion() int
	// ClientData returns the uninterpreted client payload.
	ClientData() string
}

// DocumentInterests is a document's interest registry.
type DocumentInterests interface {
	// Count returns the number of interest records.
	Count() int
	// Records returns the interest records in insertion order.
	Records() []types.DocumentInterestRecord
	// HasInterest reports whether any record's ClientID or Name matches
	// clientIDOrName — the cheap discovery probe.
	HasInterest(clientIDOrName string) bool
	// Add registers the record, updating an existing (ClientID, Name) entry
	// in place; it errors when ClientID or Name is empty.
	Add(record types.DocumentInterestRecord) error
	// Remove deletes the (clientID, name) record, reporting whether it existed.
	Remove(clientID, name string) bool
}
