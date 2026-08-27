// SPDX-License-Identifier: Apache-2.0

package types

// Hole note value types (M14-F07, #637). A hole note is a feature note: a leadered diameter callout
// on a base view's holes, computed from the hole geometry so it re-resolves with the model. These
// are the canonical Apache-2.0 definitions; the GPL model draws the callouts.

// HoleNoteQuantity is how hole notes treat multiple holes of the same size. The zero value is
// HoleNotePerHole (one callout per hole).
type HoleNoteQuantity int32

const (
	// HoleNotePerHole gives every hole its own diameter callout.
	HoleNotePerHole HoleNoteQuantity = iota
	// HoleNoteCombined groups holes by diameter into one "<n>x Ø<d>" callout per distinct diameter.
	HoleNoteCombined
)

var holeNoteQuantityNames = map[HoleNoteQuantity]string{
	HoleNotePerHole:  "perHole",
	HoleNoteCombined: "combined",
}

// String returns the quantity mode's wire spelling ("perHole", "combined").
func (q HoleNoteQuantity) String() string { return enumName(holeNoteQuantityNames, q, "enum(?)") }

// ParseHoleNoteQuantity resolves a wire spelling back to its quantity mode.
//
//	q, ok := types.ParseHoleNoteQuantity("combined") // HoleNoteCombined, true
func ParseHoleNoteQuantity(s string) (HoleNoteQuantity, bool) {
	return enumFromName(holeNoteQuantityNames, s)
}
