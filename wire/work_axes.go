// SPDX-License-Identifier: Apache-2.0

package wire

// CreateWorkAxisArgs is the request of [MethodWorkAxesCreate]: which constructor
// (Kind, a [oblikovati.org/api/types.WorkAxisKind] value) and its inputs.
//
//   - Origin and Direction give the grounded "line" axis: the origin point [x,y,z]
//     and the axis direction [dx,dy,dz], both in database units (centimetres; Direction
//     is a direction, so only its orientation matters). This is the raw axis a revolve
//     or sweep spins about when the axis is not one of the origin axes.
//   - Refs are work-feature references for the reference-model kinds — refs returned by
//     [MethodWorkPointsCreate] / [MethodWorkPlanesList] or origin constants
//     (types.WorkRefCenter, types.WorkRefXYPlane …). Each kind expects a fixed count/order:
//     "two-points" wants two point refs; "plane-intersection" wants two plane refs.
//   - Construction, when true, creates the axis as a construction (hidden, consumer-tied) work
//     feature — Inventor's WorkAxes.Add* Construction parameter; excluded from the browser and
//     auto-deleted with its last consumer. #1849.
type CreateWorkAxisArgs struct {
	Kind         string    `json:"kind"`
	Origin       []float64 `json:"origin,omitempty"`
	Direction    []float64 `json:"direction,omitempty"`
	Refs         []string  `json:"refs,omitempty"`
	Construction bool      `json:"construction,omitempty"` // #1849
}

// CreateWorkAxisResult is the response of [MethodWorkAxesCreate]: the new axis's index
// in the work-axes collection, its stable reference (for building further datums — a
// two-plane bisector, a normal-to-curve plane, or a revolve axis), its name, and whether
// it resolved (an unsatisfiable definition reports healthy=false, with Reason, rather
// than failing the call).
type CreateWorkAxisResult struct {
	Index   int    `json:"index"`
	Ref     string `json:"ref"`
	Name    string `json:"name"`
	Healthy bool   `json:"healthy"`
	Reason  string `json:"reason,omitempty"` // why Healthy is false (empty when healthy)
}

// WorkAxisInfo is one row of [MethodWorkAxesList]: a datum axis's identity and current
// geometry (Origin point and unit Direction, in database units), whether it is one of the
// origin coordinate-system axes, its health, and its constructor Kind. Origin axes report
// IsOrigin=true.
type WorkAxisInfo struct {
	Index        int       `json:"index"`
	Name         string    `json:"name"`
	Ref          string    `json:"ref"`
	Kind         string    `json:"kind,omitempty"` // a types.WorkAxisKind value
	Origin       []float64 `json:"origin"`
	Direction    []float64 `json:"direction"`
	IsOrigin     bool      `json:"isOrigin"`
	Construction bool      `json:"construction,omitempty"` // hidden, consumer-tied datum (#1849)
	Healthy      bool      `json:"healthy"`
	Reason       string    `json:"reason,omitempty"` // why Healthy is false (empty when healthy)
}

// ListWorkAxesArgs is the (optional) request of [MethodWorkAxesList]. An empty request lists the
// active part/assembly's own datum axes, hiding construction axes. Occurrence, when set, is an
// assembly occurrence path (instance names, top-down) whose component's datum axes are listed
// instead, each returned as an occurrence-qualified ref ("occ/<path>/axis/N") resolved through
// that occurrence's context transform (#1857). IncludeConstruction, when true, also lists
// construction (hidden, consumer-tied) axes (#1849).
type ListWorkAxesArgs struct {
	Occurrence          []string `json:"occurrence,omitempty"`
	IncludeConstruction bool     `json:"includeConstruction,omitempty"`
}

// ListWorkAxesResult is the response of [MethodWorkAxesList].
type ListWorkAxesResult struct {
	Axes []WorkAxisInfo `json:"axes"`
}
