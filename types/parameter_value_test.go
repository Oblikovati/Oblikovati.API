// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestToleranceKind checks the zero Tolerance maps onto ToleranceDefault (so a
// `t != Tolerance{}` has-explicit-tolerance check keeps working) while an explicit flavor is
// returned as-is (M39-F06, #1562).
func TestToleranceKind(t *testing.T) {
	if got := (Tolerance{}).Kind(); got != ToleranceDefault {
		t.Errorf("zero Tolerance.Kind() = %v, want ToleranceDefault", got)
	}
	explicit := Tolerance{Type: ToleranceSymmetric, Upper: 0.05, Lower: -0.05}
	if got := explicit.Kind(); got != ToleranceSymmetric {
		t.Errorf("explicit Tolerance.Kind() = %v, want ToleranceSymmetric", got)
	}
	if explicit == (Tolerance{}) {
		t.Error("an explicit tolerance must not equal the zero value")
	}
}

// TestToleranceFitsStaysComparable: a fits tolerance carries its ISO class strings and
// still compares by value against the zero guard (#1848 — the struct must stay comparable
// so `t != Tolerance{}` keeps distinguishing an explicit tolerance).
func TestToleranceFitsStaysComparable(t *testing.T) {
	fit := Tolerance{Type: ToleranceLimitsFitsStacked, Upper: 0.0025, HoleTolerance: "H7", ShaftTolerance: "g6"}
	if fit == (Tolerance{}) {
		t.Error("a fits tolerance must not equal the zero value")
	}
	if fit != (Tolerance{Type: ToleranceLimitsFitsStacked, Upper: 0.0025, HoleTolerance: "H7", ShaftTolerance: "g6"}) {
		t.Error("identical fits tolerances must compare equal (struct stays comparable)")
	}
}
