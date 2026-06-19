// SPDX-License-Identifier: Apache-2.0

package wire

// Body, shell and wire enumeration DTOs (M07-F06, Oblikovati/Oblikovati#629).
// Bodies are addressed by their index in the active part's body collection;
// coordinates are [x, y, z] in database units (cm).

// BodyIndexArgs addresses one body of the active part.
type BodyIndexArgs struct {
	BodyIndex int `json:"bodyIndex"`
}

// BodyInfo is one body's summary in [MethodBodyList]'s result. Name is the body's display
// name (e.g. "Solid1"); Visible reports whether it is shown, toggled with
// [MethodBodySetVisible] — the body-level API multi-body workflows (Combine, Split, Mold)
// need to enumerate and show/hide their results.
type BodyInfo struct {
	Index    int    `json:"index"`
	Name     string `json:"name"`
	Solid    bool   `json:"solid"`
	Visible  bool   `json:"visible"`
	Faces    int    `json:"faces"`
	Edges    int    `json:"edges"`
	Vertices int    `json:"vertices"`
	Shells   int    `json:"shells"`
	Wires    int    `json:"wires"`
}

// BodyListResult is the response of [MethodBodyList].
type BodyListResult struct {
	Bodies []BodyInfo `json:"bodies,omitempty"`
}

// BodySetVisibleArgs is the request of [MethodBodySetVisible]: show or hide the body at
// BodyIndex (from [MethodBodyList]). Visible is set, not toggled, so the call is idempotent.
type BodySetVisibleArgs struct {
	BodyIndex int  `json:"bodyIndex"`
	Visible   bool `json:"visible"`
}

// BodyInfoResult is the response of [MethodBodySetVisible]: the body's refreshed summary.
type BodyInfoResult struct {
	Body BodyInfo `json:"body"`
}

// FaceShellInfo is one shell in [MethodBodyShells]'s result.
type FaceShellInfo struct {
	Index int `json:"index"`
	// Closed: the shell bounds a region; Void: it is an inner cavity skin.
	Closed bool `json:"closed"`
	Void   bool `json:"void,omitempty"`
	Faces  int  `json:"faces"`
	Edges  int  `json:"edges"`
	// Volume is the magnitude of the bounded region's volume in cm³.
	Volume float64 `json:"volume"`
	// RangeBox is [minX, minY, minZ, maxX, maxY, maxZ].
	RangeBox []float64 `json:"rangeBox,omitempty"`
	// Key is the persistent reference key; TransientKey the session id.
	Key          string `json:"key"`
	TransientKey uint64 `json:"transientKey"`
}

// BodyShellsResult is the response of [MethodBodyShells].
type BodyShellsResult struct {
	Shells []FaceShellInfo `json:"shells,omitempty"`
}

// WireInfo is one wire in [MethodBodyWires]'s result.
type WireInfo struct {
	Index  int  `json:"index"`
	Closed bool `json:"closed"`
	Planar bool `json:"planar"`
	Edges  int  `json:"edges"`
	// Key is the persistent reference key; TransientKey the session id.
	Key          string `json:"key"`
	TransientKey uint64 `json:"transientKey"`
}

// BodyWiresResult is the response of [MethodBodyWires].
type BodyWiresResult struct {
	Wires []WireInfo `json:"wires,omitempty"`
}

// OffsetPlanarWireArgs is the request of [MethodWireOffsetPlanar]. Positive
// distance offsets toward normal × tangent (the left of travel about the
// plane normal); CornerClosure is a
// [oblikovati.org/api/types.OffsetCornerClosureType] wire spelling. The wire
// lives either on a document body (BodyIndex) or on a transient body
// (Handle > 0 — sections and silhouettes produce those).
type OffsetPlanarWireArgs struct {
	BodyIndex     int       `json:"bodyIndex,omitempty"`
	Handle        int       `json:"handle,omitempty"`
	WireIndex     int       `json:"wireIndex"`
	Normal        []float64 `json:"normal"`
	Distance      float64   `json:"distance"`
	CornerClosure string    `json:"cornerClosure,omitempty"`
}

// WirePolyline is one wire's sampled polyline: flattened xyz triplets.
type WirePolyline struct {
	Points []float64 `json:"points"`
	Closed bool      `json:"closed,omitempty"`
}

// OffsetPlanarWireResult is the response of [MethodWireOffsetPlanar]: the
// offset wires land on a transient body (addressable via the brep.* methods)
// and are also returned sampled for immediate consumption.
type OffsetPlanarWireResult struct {
	Handle int            `json:"handle"`
	Wires  []WirePolyline `json:"wires,omitempty"`
}
