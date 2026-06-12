// SPDX-License-Identifier: Apache-2.0

package wire

// Facet/stroke calculation and retrieval DTOs (M07-F03 remainder,
// Oblikovati/Oblikovati#293). Calculate methods facet at the chordal
// tolerance and cache the set under that exact tolerance; existing methods
// retrieve a cached set without re-faceting (an error when absent). Arrays
// are flattened: coordinates and normals are xyz triplets, texture
// coordinates uv pairs.

// CalculateFacetsArgs is the request of [MethodBodyCalculateFacets] and
// [MethodBodyExistingFacets].
type CalculateFacetsArgs struct {
	BodyIndex int     `json:"bodyIndex"`
	Tolerance float64 `json:"tolerance"`
	// IncludeTextureMap adds per-vertex surface-parameter (u, v) pairs.
	IncludeTextureMap bool `json:"includeTextureMap,omitempty"`
}

// FaceFacetsArgs is the request of [MethodFaceCalculateFacets] and
// [MethodFaceCalculateStrokes]: one face addressed by its reference key.
type FaceFacetsArgs struct {
	BodyIndex int     `json:"bodyIndex"`
	FaceKey   string  `json:"faceKey"`
	Tolerance float64 `json:"tolerance"`
}

// FacetSetResult is the facet payload: triangle mesh plus the per-face
// triangle-index counts (in body face order; one entry for a face-level set).
type FacetSetResult struct {
	VertexCount        int       `json:"vertexCount"`
	FacetCount         int       `json:"facetCount"`
	VertexCoordinates  []float64 `json:"vertexCoordinates,omitempty"`
	NormalVectors      []float64 `json:"normalVectors,omitempty"`
	VertexIndices      []int     `json:"vertexIndices,omitempty"`
	TextureCoordinates []float64 `json:"textureCoordinates,omitempty"`
	IndexCountPerFace  []int     `json:"indexCountPerFace,omitempty"`
}

// CalculateStrokesArgs is the request of [MethodBodyCalculateStrokes] and
// [MethodBodyExistingStrokes].
type CalculateStrokesArgs struct {
	BodyIndex int     `json:"bodyIndex"`
	Tolerance float64 `json:"tolerance"`
}

// StrokeSetResult is the wireframe payload: every edge sampled into a
// polyline; PolylineLengths gives each polyline's point count, in order,
// over the shared VertexCoordinates.
type StrokeSetResult struct {
	VertexCount       int       `json:"vertexCount"`
	VertexCoordinates []float64 `json:"vertexCoordinates,omitempty"`
	PolylineCount     int       `json:"polylineCount"`
	PolylineLengths   []int     `json:"polylineLengths,omitempty"`
}

// FacetTolerancesResult is the response of [MethodBodyFacetTolerances] and
// [MethodBodyStrokeTolerances]: the tolerances cached sets exist at, ascending.
type FacetTolerancesResult struct {
	Tolerances []float64 `json:"tolerances,omitempty"`
}
