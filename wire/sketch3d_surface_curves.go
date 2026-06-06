// SPDX-License-Identifier: Apache-2.0

package wire

// AddSketch3DSurfaceCurveArgs is the request of [MethodSketch3DAddSurfaceCurve] — the
// discriminated surface-derived curve constructor. Kind is the curve
// ([oblikovati/api/types.Sketch3DEntityKind]: intersection | silhouette).
// FaceRefs are the reference keys of the referenced part faces (intersection: two faces;
// silhouette: one face; onFace: one face). ViewDir [x,y,z] is the silhouette view
// direction. UV is the flat [u0,v0,u1,v1,…] parameter-space polyline for an onFace curve
// (mapped onto the referenced face's surface). The Grid* fields window the base surface's
// parameter domain for the tracer (required for an unbounded base such as a plane); zero
// values default to the surface's finite domain.
// SourceEntityID is the in-sketch source curve for a projectToSurface or offset curve
// (resolved to its kernel curve). OffsetDistance + Normal [x,y,z] parameterize an offset
// curve (offset direction = normal × tangent). For projectToSurface, FaceRefs[0] is the
// target surface.
type AddSketch3DSurfaceCurveArgs struct {
	SketchIndex    int       `json:"sketchIndex"`
	Kind           string    `json:"kind"`
	FaceRefs       []string  `json:"faceRefs"`
	ViewDir        []float64 `json:"viewDir,omitempty"`
	UV             []float64 `json:"uv,omitempty"`
	SourceEntityID uint64    `json:"sourceEntityId,omitempty"`
	OffsetDistance float64   `json:"offsetDistance,omitempty"`
	Normal         []float64 `json:"normal,omitempty"`
	GridUMin       float64   `json:"gridUMin,omitempty"`
	GridUMax       float64   `json:"gridUMax,omitempty"`
	GridVMin       float64   `json:"gridVMin,omitempty"`
	GridVMax       float64   `json:"gridVMax,omitempty"`
	GridUSteps     int       `json:"gridUSteps,omitempty"`
	GridVSteps     int       `json:"gridVSteps,omitempty"`
}

// AddSketch3DSurfaceCurveResult is the response of [MethodSketch3DAddSurfaceCurve]: the
// created entity's session id, its kind, and whether every face reference resolved.
type AddSketch3DSurfaceCurveResult struct {
	EntityID uint64 `json:"entityId"`
	Kind     string `json:"kind"`
	Healthy  bool   `json:"healthy"`
}
