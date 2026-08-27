// SPDX-License-Identifier: Apache-2.0

package types

// ASideFaceStatus reports how a face of a mold parting was assigned to the A-side of
// the tooling split (parity: ASideFaceStatusEnum). The A-side faces are those that go
// with the cavity half; "default" means the assignment was made automatically, while
// "sick" flags a face whose assignment could not be resolved. The numeric values are
// frozen at the reference API's ids and must never be renumbered.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases it
// (ADR-0018) and maps it onto the core/cavity split result.
type ASideFaceStatus int32

const (
	// ASideDefault — the face's side was assigned automatically.
	ASideDefault ASideFaceStatus = 106241
	// ASideSick — the face is on the A-side list but its assignment is invalid.
	ASideSick ASideFaceStatus = 106242
	// ASideUpToDate — the face's side assignment is explicit and current.
	ASideUpToDate ASideFaceStatus = 106243
)

var aSideFaceStatusNames = map[ASideFaceStatus]string{
	ASideDefault:  "default",
	ASideSick:     "sick",
	ASideUpToDate: "upToDate",
}

// String returns the status's wire spelling.
func (s ASideFaceStatus) String() string { return enumName(aSideFaceStatusNames, s, "enum(?)") }

// ParseASideFaceStatus resolves a wire spelling back to its status.
func ParseASideFaceStatus(s string) (ASideFaceStatus, bool) {
	return enumFromName(aSideFaceStatusNames, s)
}
