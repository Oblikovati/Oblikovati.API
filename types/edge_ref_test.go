// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestGeometricEdgeRefRoundTrip: Ref encodes and ParseGeometricEdgeRef decodes exactly (float64
// bits are preserved), and the direction sign is carried verbatim (the host normalises it).
func TestGeometricEdgeRefRoundTrip(t *testing.T) {
	want := GeometricEdgeRef{Midpoint: [3]float64{1.5, -2.25, 3.0}, Direction: [3]float64{0, 0, -1}}
	got, ok := ParseGeometricEdgeRef(want.Ref())
	if !ok {
		t.Fatalf("ParseGeometricEdgeRef(%q) reported not-an-edge-ref", want.Ref())
	}
	if got != want {
		t.Errorf("round-trip = %+v, want %+v", got, want)
	}
}

// TestParseGeometricEdgeRefRejectsOthers: a lineage-key edge ref, an unrelated ref, and a malformed
// payload are all reported as not-a-geometric-edge-ref (false), so callers fall through.
func TestParseGeometricEdgeRefRejectsOthers(t *testing.T) {
	for _, s := range []string{"edge/AAAA", "plane/3", "edge-geom/not-base64!!", "edge-geom/AAAA", ""} {
		if _, ok := ParseGeometricEdgeRef(s); ok {
			t.Errorf("ParseGeometricEdgeRef(%q) = true, want false", s)
		}
	}
}
