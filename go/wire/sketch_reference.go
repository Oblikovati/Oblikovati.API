// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// OffsetSketchArgs is the request of [MethodSketchOffset]: offset Entity (a line/circle/arc
// id) by the unit-bearing Distance (signed — a parallel line to the left of A→B, or a
// concentric circle/arc of radius r+d, for a positive distance). When Entities (a chain of
// connected line ids, in order) is set instead, the whole chain is offset and mitred. When
// ProfileIndex is set, the whole CLOSED region of that profile is offset (OpenSCAD offset(r):
// Distance>0 grows, <0 shrinks; convex corners rounded with arcs of radius |Distance|, sampled
// into ArcSegments spans per corner, default 8).
type OffsetSketchArgs struct {
	SketchIndex  int      `json:"sketchIndex"`
	Entity       uint64   `json:"entity,omitempty"`
	Entities     []uint64 `json:"entities,omitempty"`
	ProfileIndex *int     `json:"profileIndex,omitempty"`
	ArcSegments  int      `json:"arcSegments,omitempty"`
	Distance     string   `json:"distance"`
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
// projected entities and whether every reference resolved. Shared by the cut-edge and
// silhouette projection methods (#1873).
type ProjectGeometryResult struct {
	Created []uint64 `json:"created"`
	Healthy bool     `json:"healthy"`
}

// ProjectCutEdgesArgs is the request of [MethodSketchProjectCutEdges]: project the section
// curves where the sketch plane cuts the part solid, as associative reference geometry — one
// projected curve per section loop (the reference CAD API's PlanarSketch.ProjectedCuts, #1873).
type ProjectCutEdgesArgs struct {
	SketchIndex int `json:"sketchIndex"`
}

// ProjectSilhouetteArgs is the request of [MethodSketchProjectSilhouette]: project the
// silhouette of the face with reference key FaceRef onto the sketch plane, viewed along the
// plane normal (the reference CAD API's PlanarSketch.AddBySilhouette, #1873). ProximityPoint ([x,y,z] model
// cm) selects which silhouette loop when a face has several — the one nearest the point.
// IncludeBoundary keeps silhouette runs that coincide with the face's own edges.
type ProjectSilhouetteArgs struct {
	SketchIndex     int       `json:"sketchIndex"`
	FaceRef         string    `json:"faceRef"`
	ProximityPoint  []float64 `json:"proximityPoint"`
	IncludeBoundary bool      `json:"includeBoundary,omitempty"`
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
// the anchor), Justify (horizontal: "left" | "center" | "right"), VJustify (vertical:
// "baseline" | "lower" | "middle" | "upper"), and an optional Font family name and
// unit-bearing FontSize.
//
// The created entity is a single sketch TEXT entity (not baked line geometry): it keeps the
// content + font and DERIVES its glyph outlines on demand, so editing the text re-derives
// the geometry and an emboss that references the text recomputes from it — nothing is baked
// into the document. A letter's counter (the hole in A/O/B) becomes a profile hole.
type AddTextArgs struct {
	SketchIndex int       `json:"sketchIndex"`
	Anchor      []float64 `json:"anchor"`
	Text        string    `json:"text"`
	Height      string    `json:"height"`
	Rotation    string    `json:"rotation,omitempty"`
	Justify     string    `json:"justify,omitempty"`  // horizontal alignment
	VJustify    string    `json:"vJustify,omitempty"` // vertical alignment
	Font        string    `json:"font,omitempty"`     // font family name ("" ⇒ document default)
	FontSize    string    `json:"fontSize,omitempty"` // unit-bearing; "" ⇒ track Height
}

// EditTextArgs is the request of [MethodSketchEditText]: edit the existing sketch text
// entity EntityID in sketch SketchIndex. Only the non-empty/non-nil fields are applied
// (a partial edit), so an add-in can change just the content or just the font. Height and
// FontSize are unit-bearing; Justify/VJustify use the AddTextArgs vocabularies.
type EditTextArgs struct {
	SketchIndex int     `json:"sketchIndex"`
	EntityID    uint64  `json:"entityId"`
	Text        *string `json:"text,omitempty"`
	Height      string  `json:"height,omitempty"`
	Rotation    string  `json:"rotation,omitempty"`
	Justify     string  `json:"justify,omitempty"`
	VJustify    string  `json:"vJustify,omitempty"`
	Font        string  `json:"font,omitempty"`
	FontSize    string  `json:"fontSize,omitempty"`
}

// GetTextArgs is the request of [MethodSketchGetText]: read back the style of sketch text
// entity EntityID in sketch SketchIndex.
type GetTextArgs struct {
	SketchIndex int    `json:"sketchIndex"`
	EntityID    uint64 `json:"entityId"`
}

// SketchTextResult is the response of [MethodSketchGetText] (and of [MethodSketchEditText]):
// the entity id plus its resolved text style.
type SketchTextResult struct {
	EntityID uint64                `json:"entityId"`
	Style    types.SketchTextStyle `json:"style"`
}

// AddEntityIDResult is the trivial response carrying just a created entity's id (used by
// [MethodSketchAddFillRegion] and [MethodSketchAddText]).
type AddEntityIDResult struct {
	EntityID uint64 `json:"entityId"`
}
