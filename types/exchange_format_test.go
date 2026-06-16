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
	if FormatDWG.IsMesh() {
		t.Errorf("FormatDWG.IsMesh() = true, want false (DWG is a sketch/drawing format)")
	}
}

func TestExchangeFormatIsSketch(t *testing.T) {
	if !FormatDWG.IsSketch() {
		t.Errorf("FormatDWG.IsSketch() = false, want true")
	}
	for _, f := range []ExchangeFormat{FormatSTL, FormatOBJ, Format3MF, FormatSTEP} {
		if f.IsSketch() {
			t.Errorf("%q.IsSketch() = true, want false", f)
		}
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
