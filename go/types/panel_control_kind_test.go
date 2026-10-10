// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestPanelReferenceListKindName(t *testing.T) {
	if PanelReferenceList != 12 {
		t.Fatalf("PanelReferenceList = %d, want 12 (next free kind after PanelTabs=11)", PanelReferenceList)
	}
	if got := PanelReferenceList.String(); got != "referenceList" {
		t.Fatalf("PanelReferenceList.String() = %q, want %q", got, "referenceList")
	}
}
