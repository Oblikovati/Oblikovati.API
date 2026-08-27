// SPDX-License-Identifier: Apache-2.0

package wire

// CreateWorkPointArgs is the request of [MethodWorkPointsCreate]: which constructor (Kind, a
// [oblikovati.org/api/types.WorkPointKind] value) and its inputs. An empty Kind means the
// position constructor, so the original position-only request (just At) is unchanged.
//
//   - position (default): At is the point [x, y, z] in model units.
//   - plane-axis-intersection: Refs = [plane, axis] — the point where the axis pierces the
//     plane. At is ignored.
//   - curve-and-entity: Refs = [curve, entity]; Proximity [x,y,z] cm picks the nearest
//     intersection when the curve crosses the entity more than once.
//   - centroid: Refs = [edge, edge, …] — the length-weighted centroid of the edges.
//   - cloud-point: Refs = [cloudID]; At is the frozen [x,y,z] captured from that cloud.
type CreateWorkPointArgs struct {
	At   []float64 `json:"at,omitempty"`   // position kind: [x, y, z] in model units
	Kind string    `json:"kind,omitempty"` // a types.WorkPointKind value (empty = position)
	Refs []string  `json:"refs,omitempty"` // reference-model kinds: plane-axis-intersection = [plane, axis]
	// Proximity is the solution-selection point [x, y, z] (cm) for the curve-and-entity kind: when
	// the curve pierces the entity at more than one point, the intersection nearest Proximity is
	// chosen (the reference CAD API's AddByCurveAndEntity ProximityPoint). Omitting it takes the first solution. #1842.
	Proximity []float64 `json:"proximity,omitempty"`
	// Construction, when true, creates the point as a construction (hidden, consumer-tied) work
	// feature — the reference CAD API's WorkPoints.Add* Construction parameter; excluded from the browser and
	// auto-deleted with its last consumer. #1849.
	Construction bool `json:"construction,omitempty"`
}

// CreateWorkPointResult is the response of [MethodWorkPointsCreate]: the new point's index in
// the work-points collection, its stable reference (usable as a point input to a work plane —
// e.g. a three-point plane — or to [MethodWorkPlanesRedefine]'s Repick), and its name. Healthy
// is false (with Reason) for a well-formed but unsatisfiable definition — e.g. a
// plane-axis-intersection whose axis is parallel to the plane; a position point is always
// healthy.
type CreateWorkPointResult struct {
	Index   int    `json:"index"`
	Ref     string `json:"ref"`
	Name    string `json:"name"`
	Healthy bool   `json:"healthy"`
	Reason  string `json:"reason,omitempty"` // empty when healthy
}

// WorkPointInfo is one row of [MethodWorkPointsList]: a datum point's identity, current position
// ([x, y, z] in model units), whether it is the origin centre point, its visibility, health, and
// constructor kind. #1842.
type WorkPointInfo struct {
	Index        int       `json:"index"`
	Name         string    `json:"name"`
	Ref          string    `json:"ref"`
	Position     []float64 `json:"position"`
	IsOrigin     bool      `json:"isOrigin"`
	Visible      bool      `json:"visible"`
	Construction bool      `json:"construction,omitempty"` // hidden, consumer-tied datum (#1849)
	Healthy      bool      `json:"healthy"`
	Reason       string    `json:"reason,omitempty"` // why Healthy is false (empty when healthy)
	Kind         string    `json:"kind,omitempty"`   // the point's constructor kind
}

// ListWorkPointsArgs is the (optional) request of [MethodWorkPointsList]. An empty request lists
// the active part/assembly's own datum points, hiding construction points. Occurrence, when set,
// is an assembly occurrence path (instance names, top-down) whose component's datum points are
// listed instead, each returned as an occurrence-qualified ref ("occ/<path>/point/N") resolved
// through that occurrence's context transform (#1857). IncludeConstruction, when true, also lists
// construction (hidden, consumer-tied) points (#1849).
type ListWorkPointsArgs struct {
	Occurrence          []string `json:"occurrence,omitempty"`
	IncludeConstruction bool     `json:"includeConstruction,omitempty"`
}

// ListWorkPointsResult is the response of [MethodWorkPointsList].
type ListWorkPointsResult struct {
	Points []WorkPointInfo `json:"points"`
}
