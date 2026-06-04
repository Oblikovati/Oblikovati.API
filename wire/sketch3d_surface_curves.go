// SPDX-License-Identifier: Apache-2.0

package wire

// AddSketch3DSurfaceCurveArgs is the request of [MethodSketch3DAddSurfaceCurve] — the
// discriminated surface-derived curve constructor. Kind is the curve
// ([github.com/Oblikovati/api/types.Sketch3DEntityKind]: intersection | silhouette).
// FaceRefs are the reference keys of the referenced part faces (intersection: two faces;
// silhouette: one face). ViewDir [x,y,z] is the silhouette view direction. The Grid*
// fields window the base surface's parameter domain for the tracer (required for an
// unbounded base such as a plane); zero values default to the surface's finite domain.
type AddSketch3DSurfaceCurveArgs struct {
	SketchIndex int       `json:"sketchIndex"`
	Kind        string    `json:"kind"`
	FaceRefs    []string  `json:"faceRefs"`
	ViewDir     []float64 `json:"viewDir,omitempty"`
	GridUMin    float64   `json:"gridUMin,omitempty"`
	GridUMax    float64   `json:"gridUMax,omitempty"`
	GridVMin    float64   `json:"gridVMin,omitempty"`
	GridVMax    float64   `json:"gridVMax,omitempty"`
	GridUSteps  int       `json:"gridUSteps,omitempty"`
	GridVSteps  int       `json:"gridVSteps,omitempty"`
}

// AddSketch3DSurfaceCurveResult is the response of [MethodSketch3DAddSurfaceCurve]: the
// created entity's session id, its kind, and whether every face reference resolved.
type AddSketch3DSurfaceCurveResult struct {
	EntityID uint64 `json:"entityId"`
	Kind     string `json:"kind"`
	Healthy  bool   `json:"healthy"`
}
