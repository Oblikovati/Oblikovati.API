// SPDX-License-Identifier: Apache-2.0

package featureargs

// GeomFaceSel selects a face by its GEOMETRY — centroid and outward normal in model-space
// (cm / unit) — instead of by a reference key. An external author (e.g. an exporter) cannot
// mint the host's lineage reference keys, and a key located at author time does not survive
// the feature's own recompute (the base solid re-mints lineage), so geometry is the stable
// selector: the host rebinds it to a body face on every recompute (ADR-0040). Centroid is
// [x, y, z] cm; Normal is the outward unit normal [x, y, z].
type GeomFaceSel struct {
	Centroid []float64 `json:"centroid"`
	Normal   []float64 `json:"normal"`
}

// GeomEdgeSel selects an edge by its GEOMETRY — midpoint and direction in model-space
// (cm / unit) — the edge counterpart of [GeomFaceSel] (see there for why geometry beats a
// reference key). Midpoint is [x, y, z] cm; Direction is the unit tangent [x, y, z].
type GeomEdgeSel struct {
	Midpoint  []float64 `json:"midpoint"`
	Direction []float64 `json:"direction"`
}
