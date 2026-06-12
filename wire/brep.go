// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Transient B-rep factory DTOs (M07-F05, Oblikovati/Oblikovati#628).
// Transient bodies are ownerless (no document); the session addresses them by
// positive integer handles. A BrepBodyRef names either a transient body (by
// handle) or a document body (by index) as an operation source — operations
// never mutate document bodies, they copy.

// BrepBodyRef addresses an operation source: exactly one of Handle (transient,
// 1-based) or BodyIndex (document body) must be set.
type BrepBodyRef struct {
	Handle    int  `json:"handle,omitempty"`
	BodyIndex *int `json:"bodyIndex,omitempty"`
}

// BrepBodyStats summarizes a transient body after an operation.
type BrepBodyStats struct {
	Solid    bool    `json:"solid"`
	Faces    int     `json:"faces"`
	Edges    int     `json:"edges"`
	Vertices int     `json:"vertices"`
	Shells   int     `json:"shells"`
	Wires    int     `json:"wires"`
	Volume   float64 `json:"volume"`
}

// BrepHandleResult is the common "a transient body came back" payload.
type BrepHandleResult struct {
	Handle int           `json:"handle"`
	Stats  BrepBodyStats `json:"stats"`
}

// CreatePrimitiveArgs is the request of [MethodBrepCreatePrimitive]. Kind
// selects the shape and which fields apply:
//   - "block": Min, Max (opposite box corners);
//   - "cylinderCone": Bottom, Top (section centers), BottomRadius, TopRadius
//     (equal = cylinder, one zero = full cone);
//   - "sphere": Center, Radius;
//   - "torus": Center, Axis, MajorRadius, MinorRadius.
type CreatePrimitiveArgs struct {
	Kind         string    `json:"kind"`
	Min          []float64 `json:"min,omitempty"`
	Max          []float64 `json:"max,omitempty"`
	Bottom       []float64 `json:"bottom,omitempty"`
	Top          []float64 `json:"top,omitempty"`
	BottomRadius float64   `json:"bottomRadius,omitempty"`
	TopRadius    float64   `json:"topRadius,omitempty"`
	Center       []float64 `json:"center,omitempty"`
	Axis         []float64 `json:"axis,omitempty"`
	Radius       float64   `json:"radius,omitempty"`
	MajorRadius  float64   `json:"majorRadius,omitempty"`
	MinorRadius  float64   `json:"minorRadius,omitempty"`
}

// BrepBooleanArgs is the request of [MethodBrepBoolean]: the blank (a
// transient handle — it is modified in place) combined with the tool under
// Operation (a [oblikovati.org/api/types.BooleanType] wire spelling).
type BrepBooleanArgs struct {
	BlankHandle int         `json:"blankHandle"`
	Tool        BrepBodyRef `json:"tool"`
	Operation   string      `json:"operation"`
}

// BrepTransformArgs is the request of [MethodBrepTransform]: a 4×4 row-major
// matrix restricted to rotation/translation/reflection/uniform scale.
type BrepTransformArgs struct {
	Handle int       `json:"handle"`
	Matrix []float64 `json:"matrix"`
}

// BrepCopyArgs is the request of [MethodBrepCopy].
type BrepCopyArgs struct {
	Source BrepBodyRef `json:"source"`
}

// BrepSectionArgs is the request of [MethodBrepSectionWithPlane].
type BrepSectionArgs struct {
	Source      BrepBodyRef `json:"source"`
	PlaneOrigin []float64   `json:"planeOrigin"`
	PlaneNormal []float64   `json:"planeNormal"`
}

// BrepWiresResult is the payload of section/silhouette results: the wires
// land on a new transient body and are also returned sampled.
type BrepWiresResult struct {
	Handle int            `json:"handle"`
	Wires  []WirePolyline `json:"wires,omitempty"`
}

// BrepDeleteFacesArgs is the request of [MethodBrepDeleteFaces]: faces named
// by reference key are deleted (or with KeepInstead, everything BUT them);
// the openings stay open (the result may become a surface body).
type BrepDeleteFacesArgs struct {
	Handle      int      `json:"handle"`
	FaceKeys    []string `json:"faceKeys"`
	KeepInstead bool     `json:"keepInstead,omitempty"`
}

// BrepSilhouetteArgs is the request of [MethodBrepSilhouette]: the silhouette
// curves of one face viewed along ViewDirection.
type BrepSilhouetteArgs struct {
	Source        BrepBodyRef `json:"source"`
	FaceKey       string      `json:"faceKey"`
	ViewDirection []float64   `json:"viewDirection"`
	// IncludeBoundary keeps silhouette runs lying on the face's own edges.
	IncludeBoundary bool `json:"includeBoundary,omitempty"`
}

// BrepWireRef addresses one wire of a body for the ruled surface.
type BrepWireRef struct {
	Body      BrepBodyRef `json:"body"`
	WireIndex int         `json:"wireIndex"`
}

// BrepRuledSurfaceArgs is the request of [MethodBrepRuledSurface].
type BrepRuledSurfaceArgs struct {
	SectionOne BrepWireRef `json:"sectionOne"`
	SectionTwo BrepWireRef `json:"sectionTwo"`
}

// BrepImprintArgs is the request of [MethodBrepImprint].
type BrepImprintArgs struct {
	BodyOne BrepBodyRef `json:"bodyOne"`
	BodyTwo BrepBodyRef `json:"bodyTwo"`
	// ImprintCoincidentEdges is accepted for reference parity; coincident
	// boundary edges are always imprinted by the planar arrangement.
	ImprintCoincidentEdges bool `json:"imprintCoincidentEdges,omitempty"`
}

// BrepImprintResult is the response of [MethodBrepImprint]: imprinted copies
// plus the touched faces / new imprint edges of each, by reference key.
type BrepImprintResult struct {
	HandleOne            int      `json:"handleOne"`
	HandleTwo            int      `json:"handleTwo"`
	OneTouchedFaceKeys   []string `json:"oneTouchedFaceKeys,omitempty"`
	TwoTouchedFaceKeys   []string `json:"twoTouchedFaceKeys,omitempty"`
	OneImprintedEdgeKeys []string `json:"oneImprintedEdgeKeys,omitempty"`
	TwoImprintedEdgeKeys []string `json:"twoImprintedEdgeKeys,omitempty"`
}

// BrepIdenticalBodiesArgs is the request of [MethodBrepIdenticalBodies].
type BrepIdenticalBodiesArgs struct {
	Sources []BrepBodyRef `json:"sources"`
	// Tolerance is relative (0 → 1e-6, the reference default).
	Tolerance     float64 `json:"tolerance,omitempty"`
	MatchTopology bool    `json:"matchTopology,omitempty"`
	// MatchReflection defaults TRUE (mirrored bodies match) per the
	// reference; the pointer distinguishes "unset" from explicit false.
	MatchReflection *bool `json:"matchReflection,omitempty"`
}

// BrepIdenticalBodiesResult groups indices into Sources; identical bodies
// share a group.
type BrepIdenticalBodiesResult struct {
	Groups [][]int `json:"groups,omitempty"`
}

// BrepCreateFromDefinitionArgs is the request of
// [MethodBrepCreateFromDefinition]: the bottom-up definition graph.
type BrepCreateFromDefinitionArgs struct {
	Definition types.BrepBodyDefinition `json:"definition"`
}

// BrepCreateFromDefinitionResult is the response: a handle on success, the
// per-definition issues otherwise (no body is created when issues exist).
type BrepCreateFromDefinitionResult struct {
	Handle int                         `json:"handle,omitempty"`
	Stats  BrepBodyStats               `json:"stats,omitempty"`
	Issues []types.BrepDefinitionIssue `json:"issues,omitempty"`
}

// BrepHandleArgs addresses one transient body.
type BrepHandleArgs struct {
	Handle int `json:"handle"`
}

// BrepListResult is the response of [MethodBrepList].
type BrepListResult struct {
	Handles []int `json:"handles,omitempty"`
}
