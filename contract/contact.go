// SPDX-License-Identifier: Apache-2.0

package contract

// The scalar read surface of the assembly contact & interference system (M12-F05, Oblikovati/
// Oblikovati#362/#368): contact sets (occurrences that resist interpenetration when dragged)
// and static interference analysis (the overlapping volumes between occurrences). An in-proc
// consumer reads these directly; membership and analysis travel over api/wire (contactSets.*/
// contactSolver.*/interference.*). The host implementations live in /source (model/assembly).

// ContactSet is a named group of occurrences that resist interpenetration: when the contact
// solver is enabled, dragging a member stops at contact with another member.
type ContactSet interface {
	// ID is the contact set's session id.
	ID() uint64
	// Name is the contact set's display name.
	Name() string
	// MemberCount returns the number of occurrences in the set.
	MemberCount() int
}

// ContactSets is an assembly's contact-set collection (host: assembly.ContactSets).
type ContactSets = Enumerable[ContactSet]

// ContactSolver reports whether contact enforcement is enabled and how many sets it governs —
// the reference API's ActiveContactSolver toggle.
type ContactSolver interface {
	// Enabled reports whether contact is enforced during a drag.
	Enabled() bool
	// SetCount returns the number of contact sets the solver governs.
	SetCount() int
}

// InterferenceResult is one overlapping pair from a static interference analysis: the two
// occurrences and the volume (and a representative point) of their overlap.
type InterferenceResult interface {
	// OccurrenceA / OccurrenceB are the session ids of the two interfering occurrences.
	OccurrenceA() uint64
	OccurrenceB() uint64
	// Volume is the overlap volume (cubic centimetres).
	Volume() float64
}

// InterferenceResults is the outcome of an interference analysis: the interfering pairs and
// the total overlap volume.
type InterferenceResults interface {
	Enumerable[InterferenceResult]
	// TotalVolume is the sum of every pair's overlap volume.
	TotalVolume() float64
}
