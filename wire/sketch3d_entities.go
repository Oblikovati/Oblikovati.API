// SPDX-License-Identifier: Apache-2.0

package wire

// AddSketch3DEntityArgs is the request of [MethodSketch3DAddEntity] — the discriminated
// 3D entity constructor. Kind is the base entity
// ([github.com/Oblikovati/api/types.Sketch3DEntityKind]: point | line | circle | arc | …).
// Points are the defining points, each [x,y,z] in model database units (cm), in the
// constructor's expected order (line: A,B; circle: center; arc: center,start,end). Radius
// is a unit-bearing expression ("10 mm") for circular kinds. Axis is the circle plane's
// normal [x,y,z] (defaults to +Z when empty). CCW orients an arc. Construction marks the
// entity as reference geometry.
type AddSketch3DEntityArgs struct {
	SketchIndex  int         `json:"sketchIndex"`
	Kind         string      `json:"kind"`
	Points       [][]float64 `json:"points,omitempty"`
	Radius       string      `json:"radius,omitempty"`
	Axis         []float64   `json:"axis,omitempty"`
	CCW          bool        `json:"ccw,omitempty"`
	Construction bool        `json:"construction,omitempty"`
}

// AddSketch3DEntityResult is the response of [MethodSketch3DAddEntity]: the created
// entity's session id, its base kind, and the session ids of its defining points.
type AddSketch3DEntityResult struct {
	EntityID uint64   `json:"entityId"`
	Kind     string   `json:"kind"`
	PointIDs []uint64 `json:"pointIds,omitempty"`
}
