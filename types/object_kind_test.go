// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestObjectKindSpellings pins the on-the-wire spelling of each metadata object kind: these strings
// are the source of truth the host emits and add-ins match on, so a drift silently misroutes events.
func TestObjectKindSpellings(t *testing.T) {
	cases := map[ObjectKind]string{
		ObjectKindBody:       "body",
		ObjectKindSketch:     "sketch",
		ObjectKindFeature:    "feature",
		ObjectKindOccurrence: "occurrence",
		ObjectKindDocument:   "document",
	}
	for got, want := range cases {
		if string(got) != want {
			t.Errorf("ObjectKind %q, want %q", string(got), want)
		}
	}
}
