// SPDX-License-Identifier: Apache-2.0

package wire

// AddSketchEntityArgs is the request of [MethodSketchAddEntity] — the discriminated
// entity constructor. Kind is the base entity
// ([github.com/Oblikovati/api/types.SketchEntityKind]: line | point | circle | arc | …);
// Variant selects the overload within that kind ("centerRadius" | "threePoint" |
// "centerStartEnd"; empty ⇒ the kind's default). Points are the defining points, each
// [x,y] in sketch-plane database units (cm), in the constructor's expected order. Radius
// is a unit-bearing expression ("10 mm") for the center-radius circle. CCW orients a
// center-start-end arc. Construction marks the entity as reference geometry.
type AddSketchEntityArgs struct {
	SketchIndex  int         `json:"sketchIndex"`
	Kind         string      `json:"kind"`
	Variant      string      `json:"variant,omitempty"`
	Points       [][]float64 `json:"points,omitempty"`
	Radius       string      `json:"radius,omitempty"`
	CCW          bool        `json:"ccw,omitempty"`
	Construction bool        `json:"construction,omitempty"`
}

// AddSketchEntityResult is the response of [MethodSketchAddEntity]: the new entity's
// session id, its base kind, and the session ids of its defining points (for use as
// constraint/dimension references).
type AddSketchEntityResult struct {
	EntityID uint64   `json:"entityId"`
	Kind     string   `json:"kind"`
	PointIDs []uint64 `json:"pointIds"`
}
