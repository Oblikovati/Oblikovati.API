// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// The M07 B-rep enum blocks are FROZEN to the reference ids and wire
// spellings (Oblikovati/Oblikovati#628/#629/#630) — saved automations and
// add-ins depend on both. These pins fail on any renumber or respell.

// TestBooleanTypeFrozenBlock pins the reference ids and wire spellings.
func TestBooleanTypeFrozenBlock(t *testing.T) {
	want := map[BooleanType]string{74241: "difference", 74242: "union", 74243: "intersect"}
	assertFrozenBlock(t, "BooleanType", want, booleanTypeNames, ParseBooleanType)
}

// TestOffsetCornerClosureFrozenBlock pins the reference ids and spellings.
func TestOffsetCornerClosureFrozenBlock(t *testing.T) {
	want := map[OffsetCornerClosureType]string{96257: "circular", 96258: "linear", 96259: "extend"}
	assertFrozenBlock(t, "OffsetCornerClosureType", want, offsetCornerClosureNames, ParseOffsetCornerClosureType)
}

// TestEdgeCollectionKindFrozenBlock pins the reference ids and spellings.
func TestEdgeCollectionKindFrozenBlock(t *testing.T) {
	want := map[EdgeCollectionKind]string{
		27649: "tangentiallyConnected", 27650: "allConcave", 27651: "allConvex", 27652: "undefined",
	}
	assertFrozenBlock(t, "EdgeCollectionKind", want, edgeCollectionKindNames, ParseEdgeCollectionKind)
}

// TestParseContainmentRoundTrip pins the (pre-existing) containment block's
// new Parse path, which the point-inside queries reply through.
func TestParseContainmentRoundTrip(t *testing.T) {
	want := map[Containment]string{30977: "unknown", 30978: "inside", 30979: "on", 30980: "outside"}
	assertFrozenBlock(t, "Containment", want, containmentNames, ParseContainment)
}
