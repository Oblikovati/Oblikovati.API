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
