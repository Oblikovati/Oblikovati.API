// SPDX-License-Identifier: Apache-2.0

package wire

// CreateWorkPlaneArgs is the request of [MethodWorkPlanesCreate]: which constructor
// (Kind, a [oblikovati.org/api/types.WorkPlaneKind] value) and its inputs.
//
//   - Refs are work-feature references the plane is built on — origin constants
//     (types.WorkRefXYPlane …), refs returned by [MethodWorkPlanesList], or a face
//     reference for the tangent kinds. Each kind expects a fixed count/order (e.g.
//     "plane-offset" wants one plane ref; "three-points" wants three point refs).
//   - Offset and Angle are unit-bearing expressions ("10 mm", "45 deg") for the
//     offset and line-plane-angle kinds.
//   - Origin, XAxis, YAxis give the AddFixed frame: the origin point [x,y,z] (model
//     units) and two in-plane axis direction components.
type CreateWorkPlaneArgs struct {
	Kind   string    `json:"kind"`
	Refs   []string  `json:"refs,omitempty"`
	Offset string    `json:"offset,omitempty"`
	Angle  string    `json:"angle,omitempty"`
	Origin []float64 `json:"origin,omitempty"`
	XAxis  []float64 `json:"xaxis,omitempty"`
	YAxis  []float64 `json:"yaxis,omitempty"`
}

// CreateWorkPlaneResult is the response of [MethodWorkPlanesCreate]: the new plane's
// index in the work-planes collection, its stable reference (for building further
// datums), its name, and whether it resolved (an unsatisfiable definition reports
// healthy=false rather than failing the call).
type CreateWorkPlaneResult struct {
	Index   int    `json:"index"`
	Ref     string `json:"ref"`
	Name    string `json:"name"`
	Healthy bool   `json:"healthy"`
}

// WorkPlaneInfo is one row of [MethodWorkPlanesList]: a datum plane's identity and
// current geometry (origin and unit normal, in model units), whether it is one of the
// origin coordinate-system planes, its health, its constructor Kind, and — for a user
// plane — the editable inputs [MethodWorkPlanesRedefine] accepts (its scalars and its
// re-pickable reference slots). Origin planes report no scalars/slots (not redefinable).
type WorkPlaneInfo struct {
	Index    int                `json:"index"`
	Name     string             `json:"name"`
	Ref      string             `json:"ref"`
	Origin   []float64          `json:"origin"`
	Normal   []float64          `json:"normal"`
	IsOrigin bool               `json:"isOrigin"`
	Healthy  bool               `json:"healthy"`
	Reason   string             `json:"reason,omitempty"`  // why Healthy is false (empty when healthy)
	Kind     string             `json:"kind,omitempty"`    // a types.WorkPlaneKind value
	Scalars  []WorkPlaneScalar  `json:"scalars,omitempty"` // editable distance/angle inputs
	Slots    []WorkPlaneRefSlot `json:"slots,omitempty"`   // re-pickable reference inputs
}

// WorkPlaneScalar describes one editable scalar of a work plane (offset distance, line-plane
// angle): its slot Index, Label, the Unit its value is shown in ("mm", "deg", …), and the
// current Value in that unit. A redefine sets it via [ScalarEdit] keyed on Index.
type WorkPlaneScalar struct {
	Index int     `json:"index"`
	Label string  `json:"label"`
	Unit  string  `json:"unit,omitempty"`
	Value float64 `json:"value"`
}

// WorkPlaneRefSlot describes one re-pickable reference of a work plane: its slot Index, Label,
// and Kind — "plane" | "axis" | "point" | "face" — so a client knows what reference string a
// [SlotRepick] for this slot accepts (an origin constant, a List ref, or a face key).
type WorkPlaneRefSlot struct {
	Index int    `json:"index"`
	Label string `json:"label"`
	Kind  string `json:"kind"`
}

// ListWorkPlanesResult is the response of [MethodWorkPlanesList].
type ListWorkPlanesResult struct {
	Planes []WorkPlaneInfo `json:"planes"`
}

// RedefineWorkPlaneArgs is the request of [MethodWorkPlanesRedefine]: edit a placed user work
// plane in place. Index selects the plane (its position in List). Scalars sets editable
// distances/angles; Repick re-points reference slots at new geometry. Both are optional and
// applied together, then the part recomputes.
type RedefineWorkPlaneArgs struct {
	Index   int          `json:"index"`
	Scalars []ScalarEdit `json:"scalars,omitempty"`
	Repick  []SlotRepick `json:"repick,omitempty"`
}

// ScalarEdit sets the work plane's scalar at slot Index to a unit-bearing Value ("30 mm",
// "60 deg") — the index comes from [WorkPlaneScalar].
type ScalarEdit struct {
	Index int    `json:"index"`
	Value string `json:"value"`
}

// SlotRepick re-points the work plane's reference Slot at a new reference Ref — an origin
// constant (types.WorkRefXYPlane …), a ref returned by List, or a face key. Slot is the
// [WorkPlaneRefSlot] index.
type SlotRepick struct {
	Slot int    `json:"slot"`
	Ref  string `json:"ref"`
}

// RedefineWorkPlaneResult is the response of [MethodWorkPlanesRedefine]: the plane's refreshed
// info (new geometry/health and its still-editable scalars/slots). An unsatisfiable edit
// reports healthy=false rather than failing the call.
type RedefineWorkPlaneResult struct {
	Plane WorkPlaneInfo `json:"plane"`
}
