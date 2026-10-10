// SPDX-License-Identifier: Apache-2.0

package types

// DocumentInterestType is the strength of a client's registered interest in a
// document (M03-F10, Oblikovati/Oblikovati#611).
//
// The values are a frozen block matching the reference API's document-interest
// enum; never renumber them.
type DocumentInterestType int32

const (
	// InterestNone marks an interest record as withdrawn while keeping it
	// addressable (the reference API's "not interested").
	InterestNone DocumentInterestType = 68865
	// Interested declares the client has data in / depends on the document.
	Interested DocumentInterestType = 68866
)

// documentInterestTypeNames are the frozen wire spellings.
var documentInterestTypeNames = map[DocumentInterestType]string{
	InterestNone: "notInterested",
	Interested:   "interested",
}

// String returns the interest type's wire spelling.
func (t DocumentInterestType) String() string {
	return enumName(documentInterestTypeNames, t, "enum(?)")
}

// ParseDocumentInterestType resolves a wire spelling back to its type.
func ParseDocumentInterestType(s string) (DocumentInterestType, bool) {
	return enumFromName(documentInterestTypeNames, s)
}

// DocumentInterestRecord is one entry of a document's add-in data registry:
// client X has data in / depends on this document. Unlike attribute sets
// (arbitrary payload on arbitrary objects), an interest is a lightweight
// discovery contract readable without activating the owning add-in.
type DocumentInterestRecord struct {
	// ClientID uniquely identifies the owning client (an add-in id).
	ClientID string `json:"clientId"`
	// Name names this interest, typically the data's sub-type; (ClientID,
	// Name) is the record's identity.
	Name string `json:"name"`
	// InterestType is the interest's strength.
	InterestType DocumentInterestType `json:"interestType"`
	// DataVersion is the client-managed migration version; 0 declares a
	// non-migrating interest.
	DataVersion int `json:"dataVersion,omitempty"`
	// ClientData is an uninterpreted client payload — use with extreme
	// economy (attribute sets are the payload layer).
	ClientData string `json:"clientData,omitempty"`
}
