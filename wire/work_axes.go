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
type CreateWorkAxisArgs struct {
	Kind      string    `json:"kind"`
	Origin    []float64 `json:"origin,omitempty"`
	Direction []float64 `json:"direction,omitempty"`
	Refs      []string  `json:"refs,omitempty"`
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
	Index     int       `json:"index"`
	Name      string    `json:"name"`
	Ref       string    `json:"ref"`
	Kind      string    `json:"kind,omitempty"` // a types.WorkAxisKind value
	Origin    []float64 `json:"origin"`
	Direction []float64 `json:"direction"`
	IsOrigin  bool      `json:"isOrigin"`
	Healthy   bool      `json:"healthy"`
	Reason    string    `json:"reason,omitempty"` // why Healthy is false (empty when healthy)
}

// ListWorkAxesResult is the response of [MethodWorkAxesList].
type ListWorkAxesResult struct {
	Axes []WorkAxisInfo `json:"axes"`
}
