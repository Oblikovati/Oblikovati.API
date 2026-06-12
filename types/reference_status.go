// SPDX-License-Identifier: Apache-2.0

package types

// ReferenceStatus is the resolution state of one file-to-file reference — the
// single status vocabulary shared by file descriptors (M03-F07), external-file
// attachments (M03-F08) and, later, drawing/assembly references
// (Oblikovati/Oblikovati#608).
//
// The values are a frozen block matching the reference API's reference-status
// enum; never renumber them.
type ReferenceStatus int32

const (
	// ReferenceUnknown means the reference has not been resolved yet (e.g. the
	// owning file was opened deferred and nothing probed the target).
	ReferenceUnknown ReferenceStatus = 49665
	// ReferenceUpToDate means the referenced file was found and matches what
	// the owner saved against.
	ReferenceUpToDate ReferenceStatus = 49666
	// ReferenceOutOfDate means the referenced file was found but has been
	// saved since the owner last recorded it.
	ReferenceOutOfDate ReferenceStatus = 49667
	// ReferenceMissing means the referenced file cannot be found at any
	// resolvable location.
	ReferenceMissing ReferenceStatus = 49668
	// ReferenceReplaced means the reference was re-pointed at another file
	// (a repair) and the owner has not been saved since.
	ReferenceReplaced ReferenceStatus = 49669
)

// referenceStatusNames are the frozen wire spellings.
var referenceStatusNames = map[ReferenceStatus]string{
	ReferenceUnknown:   "unknown",
	ReferenceUpToDate:  "upToDate",
	ReferenceOutOfDate: "outOfDate",
	ReferenceMissing:   "missing",
	ReferenceReplaced:  "replaced",
}

// String returns the reference status's wire spelling.
func (s ReferenceStatus) String() string { return enumName(referenceStatusNames, s) }

// ParseReferenceStatus resolves a wire spelling back to its ReferenceStatus.
func ParseReferenceStatus(s string) (ReferenceStatus, bool) {
	return enumFromName(referenceStatusNames, s)
}
