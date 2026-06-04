// SPDX-License-Identifier: Apache-2.0

package wire

// OffsetSketchArgs is the request of [MethodSketchOffset]: offset Entity (a line/circle/arc
// id) by the unit-bearing Distance (signed — a parallel line to the left of A→B, or a
// concentric circle/arc of radius r+d, for a positive distance). When Entities (a chain of
// connected line ids, in order) is set instead, the whole chain is offset and mitred.
type OffsetSketchArgs struct {
	SketchIndex int      `json:"sketchIndex"`
	Entity      uint64   `json:"entity,omitempty"`
	Entities    []uint64 `json:"entities,omitempty"`
	Distance    string   `json:"distance"`
}

// OffsetSketchResult is the response of [MethodSketchOffset]: the primary new entity's id
// and kind, plus all created entity ids (more than one for a chain offset).
type OffsetSketchResult struct {
	EntityID uint64   `json:"entityId"`
	Kind     string   `json:"kind"`
	Created  []uint64 `json:"created,omitempty"`
}

// AutoDimensionResult is the response of [MethodSketchAutoDimension]: how many constraints
// were added and the sketch's resulting DOF (0 when it is now fully constrained).
type AutoDimensionResult struct {
	Added int `json:"added"`
	DOF   int `json:"dof"`
}

// AddSketchImageArgs is the request of [MethodSketchAddImage]: place a raster image (Ref
// is a package-store reference/path) anchored at Anchor ([x,y] cm) with the given
// unit-bearing Width/Height, Rotation (unit-bearing angle, CCW about the anchor), and
// Opacity (0…1).
type AddSketchImageArgs struct {
	SketchIndex int       `json:"sketchIndex"`
	Ref         string    `json:"ref"`
	Anchor      []float64 `json:"anchor"`
	Width       string    `json:"width"`
	Height      string    `json:"height"`
	Rotation    string    `json:"rotation,omitempty"`
	Opacity     float64   `json:"opacity,omitempty"`
}

// AddSketchImageResult is the response of [MethodSketchAddImage]: the new image's id.
type AddSketchImageResult struct {
	EntityID uint64 `json:"entityId"`
}

// ProjectGeometryArgs is the request of [MethodSketchProject]: project the part topology
// referenced by Refs (edge/vertex reference-key strings) onto the sketch plane. Mode is
// "reference" (associative reference geometry, default) or "include" (projected as
// ordinary sketch geometry). Each ref becomes a projected point (vertex) or curve (edge).
type ProjectGeometryArgs struct {
	SketchIndex int      `json:"sketchIndex"`
	Refs        []string `json:"refs"`
	Mode        string   `json:"mode,omitempty"`
}

// ProjectGeometryResult is the response of [MethodSketchProject]: the ids of the created
// projected entities and whether every reference resolved.
type ProjectGeometryResult struct {
	Created []uint64 `json:"created"`
	Healthy bool     `json:"healthy"`
}

// AddFillRegionArgs is the request of [MethodSketchAddFillRegion]: fill the closed region
// containing Seed ([x,y] cm) with the named Style (empty ⇒ solid).
type AddFillRegionArgs struct {
	SketchIndex int       `json:"sketchIndex"`
	Seed        []float64 `json:"seed"`
	Style       string    `json:"style,omitempty"`
}

// AddTextArgs is the request of [MethodSketchAddText]: place Text anchored at Anchor
// ([x,y] cm), with a unit-bearing Height, an optional unit-bearing Rotation (CCW about
// the anchor), and a Justify ("left" | "center" | "right").
type AddTextArgs struct {
	SketchIndex int       `json:"sketchIndex"`
	Anchor      []float64 `json:"anchor"`
	Text        string    `json:"text"`
	Height      string    `json:"height"`
	Rotation    string    `json:"rotation,omitempty"`
	Justify     string    `json:"justify,omitempty"`
}

// AddEntityIDResult is the trivial response carrying just a created entity's id (used by
// [MethodSketchAddFillRegion] and [MethodSketchAddText]).
type AddEntityIDResult struct {
	EntityID uint64 `json:"entityId"`
}
