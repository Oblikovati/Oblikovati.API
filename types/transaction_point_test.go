// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestTransactionPointFrozenBlock pins the reference ids and wire spellings —
// the block is frozen and a renumber or respell is a breaking wire change.
func TestTransactionPointFrozenBlock(t *testing.T) {
	want := map[TransactionPoint]string{
		3585: "unknown", 3586: "next", 3587: "previous",
		3588: "current", 3589: "upToSpecified",
	}
	if len(want) != len(transactionPointNames) {
		t.Fatalf("transaction point count = %d, want %d", len(transactionPointNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("TransactionPoint(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseTransactionPoint(name); !ok || parsed != v {
			t.Errorf("ParseTransactionPoint(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	if _, ok := ParseTransactionPoint("noSuchPoint"); ok {
		t.Error("ParseTransactionPoint must reject unknown spellings")
	}
}
