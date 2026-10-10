// SPDX-License-Identifier: Apache-2.0

package wire

// Spline tangency handles (M06-F11, Oblikovati/Oblikovati#626): every fit
// point of an interpolation spline owns a latent handle; activating it adds a
// handle entity (a constrainable tangent direction + weight) that shapes the
// curve through that point.

// SetSplineHandleArgs is the request of [MethodSketchSetSplineHandle] and
// [MethodSketch3DSetSplineHandle]. Spline is the spline's session id and
// FitPointIndex which fit point's handle to address (0-based). Active
// activates or deactivates the handle. Tangent is the handle direction —
// [x,y] in sketch-plane coordinates for the 2D method, [x,y,z] for the 3D one
// (empty keeps the current/natural direction). Weight scales the tangent
// magnitude's pull on the curve (0 keeps the current weight).
type SetSplineHandleArgs struct {
	SketchIndex   int       `json:"sketchIndex"`
	Spline        uint64    `json:"spline"`
	FitPointIndex int       `json:"fitPointIndex"`
	Active        bool      `json:"active"`
	Tangent       []float64 `json:"tangent,omitempty"`
	Weight        float64   `json:"weight,omitempty"`
}

// SplineHandleInfo is the response of [MethodSketchSetSplineHandle] and
// [MethodSketch3DSetSplineHandle]: the handle's session id (0 when the call
// deactivated it), its current tangent direction and weight.
type SplineHandleInfo struct {
	HandleID uint64    `json:"handleId,omitempty"`
	Tangent  []float64 `json:"tangent,omitempty"`
	Weight   float64   `json:"weight,omitempty"`
}
