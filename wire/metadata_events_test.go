// SPDX-License-Identifier: Apache-2.0

package wire

import "testing"

// TestMetadataEventNames pins the on-the-wire spelling of the two metadata-mutation events. The
// host emits and add-ins subscribe to these exact strings, so a drift silently breaks every add-in
// listening for renames or property changes; the pair must also stay distinct.
func TestMetadataEventNames(t *testing.T) {
	if EventObjectRenamed != "object.renamed" {
		t.Errorf("EventObjectRenamed = %q, want %q", EventObjectRenamed, "object.renamed")
	}
	if EventPropertyChanged != "property.changed" {
		t.Errorf("EventPropertyChanged = %q, want %q", EventPropertyChanged, "property.changed")
	}
	if EventObjectRenamed == EventPropertyChanged {
		t.Error("object.renamed and property.changed must be distinct event names")
	}
}
