// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestExchangeFormatIsMesh(t *testing.T) {
	mesh := []ExchangeFormat{FormatSTL, FormatOBJ, Format3MF}
	for _, f := range mesh {
		if !f.IsMesh() {
			t.Errorf("%q.IsMesh() = false, want true", f)
		}
	}
	if FormatSTEP.IsMesh() {
		t.Errorf("FormatSTEP.IsMesh() = true, want false (STEP is a B-rep format)")
	}
}

func TestMeshResolutionNormalizedDefaultsToMedium(t *testing.T) {
	if got := MeshResolution("").Normalized(); got != ResolutionMedium {
		t.Errorf("empty.Normalized() = %q, want medium", got)
	}
	if got := ResolutionHigh.Normalized(); got != ResolutionHigh {
		t.Errorf("high.Normalized() = %q, want high", got)
	}
}
