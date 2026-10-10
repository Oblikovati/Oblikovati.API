// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestMoveOperationTypeValidity(t *testing.T) {
	for _, op := range AllMoveOperationTypes() {
		if !op.IsValid() {
			t.Errorf("AllMoveOperationTypes returned %q which IsValid rejects", op)
		}
	}
	if (MoveOperationType("spin")).IsValid() {
		t.Error("MoveOperationType.IsValid must reject unknown spellings")
	}
}

func TestMoveOperationTypeFrozenSpellings(t *testing.T) {
	want := map[MoveOperationType]bool{"freeDrag": true, "alongRay": true, "rotateAboutLine": true}
	if got := AllMoveOperationTypes(); len(got) != len(want) {
		t.Fatalf("AllMoveOperationTypes count = %d, want %d", len(got), len(want))
	}
	for _, op := range AllMoveOperationTypes() {
		if !want[op] {
			t.Errorf("unexpected move-operation spelling %q (frozen set changed?)", op)
		}
	}
}
