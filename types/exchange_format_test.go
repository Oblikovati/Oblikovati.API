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
	if FormatDXF.IsMesh() {
		t.Errorf("FormatDXF.IsMesh() = true, want false (DXF is a sketch/drawing format)")
	}
}

func TestExchangeFormatIsSketch(t *testing.T) {
	for _, f := range []ExchangeFormat{FormatDWG, FormatDXF} {
		if !f.IsSketch() {
			t.Errorf("%q.IsSketch() = false, want true", f)
		}
	}
	for _, f := range []ExchangeFormat{FormatSTL, FormatOBJ, Format3MF, FormatSTEP} {
		if f.IsSketch() {
			t.Errorf("%q.IsSketch() = true, want false", f)
		}
	}
}

func TestDXFVersionNormalizedDefaultsToR2000(t *testing.T) {
	if got := DXFVersion("").Normalized(); got != DXFR2000 {
		t.Errorf("empty.Normalized() = %q, want r2000", got)
	}
	if got := DXFR2018.Normalized(); got != DXFR2018 {
		t.Errorf("r2018.Normalized() = %q, want r2018", got)
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
