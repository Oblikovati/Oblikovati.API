// SPDX-License-Identifier: Apache-2.0

package wire

// OffsetSketchArgs is the request of [MethodSketchOffset]: offset the entity Entity (a
// line/circle/arc id) by the unit-bearing Distance (signed — a parallel line to the left
// of A→B, or a concentric circle/arc of radius r+d, for a positive distance).
type OffsetSketchArgs struct {
	SketchIndex int    `json:"sketchIndex"`
	Entity      uint64 `json:"entity"`
	Distance    string `json:"distance"`
}

// OffsetSketchResult is the response of [MethodSketchOffset]: the new entity's id and kind.
type OffsetSketchResult struct {
	EntityID uint64 `json:"entityId"`
	Kind     string `json:"kind"`
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
