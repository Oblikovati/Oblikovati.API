// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/base64"
	"encoding/binary"
	"math"
	"strings"
)

// GeometricEdgeRef names a B-rep edge by WHERE it sits — its midpoint and a sign-agnostic
// direction — rather than by an Oblikovati lineage key (ADR-0040). It is the external author's
// path to a topology selection: an exporter reading another CAD system knows the edge's geometry,
// not Oblikovati's internally-minted lineage. The host re-binds it to the running body each
// recompute by geometric nearness (topo.FindEdgeByGeometry), returning it only when it is
// unambiguous. Coordinates are model space (database units, centimetres).
//
// A work-feature reference string carries this descriptor via [GeometricEdgeRef.Ref]; the same
// string can equally be a lineage-key edge reference ("edge/…") produced by an interactive pick —
// both resolve through the host's edge resolver. Use it wherever a work-axis/point create takes an
// edge reference (analytic-edge / line-by-entity axis, edge-midpoint point).
type GeometricEdgeRef struct {
	Midpoint  [3]float64 `json:"midpoint"`  // a point on the edge, model cm
	Direction [3]float64 `json:"direction"` // the edge tangent (orientation only; sign ignored)
}

// edgeGeomPrefix tags the reference string of a geometric edge descriptor, distinct from the
// lineage-key "edge/" form so the host resolver can route each to the right binder.
const edgeGeomPrefix = "edge-geom/"

// Ref encodes the descriptor as a work-feature reference string ("edge-geom/<base64>"): the six
// float64 components in little-endian, URL-safe-base64'd (no '/' or '+', so it never collides with
// the reference path separator). Pass the result as an edge reference to a work-axis/point create.
func (r GeometricEdgeRef) Ref() string {
	var buf [48]byte
	comps := [6]float64{r.Midpoint[0], r.Midpoint[1], r.Midpoint[2], r.Direction[0], r.Direction[1], r.Direction[2]}
	for i, c := range comps {
		binary.LittleEndian.PutUint64(buf[i*8:], math.Float64bits(c))
	}
	return edgeGeomPrefix + base64.RawURLEncoding.EncodeToString(buf[:])
}

// ParseGeometricEdgeRef decodes a reference string produced by [GeometricEdgeRef.Ref]. The bool is
// false when s is not a geometric edge reference (e.g. a lineage-key "edge/…" form or an unrelated
// ref) or its payload is malformed, so a caller can fall through to the other reference kinds.
func ParseGeometricEdgeRef(s string) (GeometricEdgeRef, bool) {
	if !strings.HasPrefix(s, edgeGeomPrefix) {
		return GeometricEdgeRef{}, false
	}
	raw, err := base64.RawURLEncoding.DecodeString(strings.TrimPrefix(s, edgeGeomPrefix))
	if err != nil || len(raw) != 48 {
		return GeometricEdgeRef{}, false
	}
	var c [6]float64
	for i := range c {
		c[i] = math.Float64frombits(binary.LittleEndian.Uint64(raw[i*8:]))
	}
	return GeometricEdgeRef{Midpoint: [3]float64{c[0], c[1], c[2]}, Direction: [3]float64{c[3], c[4], c[5]}}, true
}
