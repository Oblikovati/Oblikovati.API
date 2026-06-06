// SPDX-License-Identifier: Apache-2.0

package wire

// CreateWorkPlaneArgs is the request of [MethodWorkPlanesCreate]: which constructor
// (Kind, a [oblikovati/api/types.WorkPlaneKind] value) and its inputs.
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
// origin coordinate-system planes, and its health.
type WorkPlaneInfo struct {
	Index    int       `json:"index"`
	Name     string    `json:"name"`
	Ref      string    `json:"ref"`
	Origin   []float64 `json:"origin"`
	Normal   []float64 `json:"normal"`
	IsOrigin bool      `json:"isOrigin"`
	Healthy  bool      `json:"healthy"`
}

// ListWorkPlanesResult is the response of [MethodWorkPlanesList].
type ListWorkPlanesResult struct {
	Planes []WorkPlaneInfo `json:"planes"`
}
