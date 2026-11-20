// SPDX-License-Identifier: Apache-2.0

package types

// SketchInferenceTypes is a bit set of selected inference relation types.
// Combine choices with bitwise OR; for example, SketchInferHorizontal | SketchInferVertical.
type SketchInferenceTypes uint32

const (
	// SketchInferHorizontal selects Horizontal.
	SketchInferHorizontal SketchInferenceTypes = 1
	// SketchInferVertical selects Vertical.
	SketchInferVertical SketchInferenceTypes = 2
	// SketchInferParallel selects Parallel.
	SketchInferParallel SketchInferenceTypes = 4
	// SketchInferPerpendicular selects Perpendicular.
	SketchInferPerpendicular SketchInferenceTypes = 8
	// SketchInferIntersection selects Intersection.
	SketchInferIntersection SketchInferenceTypes = 16
	// SketchInferMidpoint selects Midpoint.
	SketchInferMidpoint SketchInferenceTypes = 32
	// SketchInferOnCurve selects OnCurve.
	SketchInferOnCurve SketchInferenceTypes = 64
	// SketchInferTangent selects Tangent.
	SketchInferTangent SketchInferenceTypes = 128
	// SketchInferCoincident selects Coincident.
	SketchInferCoincident SketchInferenceTypes = 256
)

// SketchRelaxTypes is a bit set of selected relaxation relation types.
// Combine choices with bitwise OR; for example, SketchRelaxCoincident | SketchRelaxTangent.
type SketchRelaxTypes uint32

const (
	// SketchRelaxCoincident selects Coincident.
	SketchRelaxCoincident SketchRelaxTypes = 1
	// SketchRelaxTangent selects Tangent.
	SketchRelaxTangent SketchRelaxTypes = 2
	// SketchRelaxSmooth selects Smooth.
	SketchRelaxSmooth SketchRelaxTypes = 4
	// SketchRelaxSymmetric selects Symmetric.
	SketchRelaxSymmetric SketchRelaxTypes = 8
	// SketchRelaxCollinear selects Collinear.
	SketchRelaxCollinear SketchRelaxTypes = 16
	// SketchRelaxConcentric selects Concentric.
	SketchRelaxConcentric SketchRelaxTypes = 32
	// SketchRelaxHorizontal selects Horizontal.
	SketchRelaxHorizontal SketchRelaxTypes = 64
	// SketchRelaxVertical selects Vertical.
	SketchRelaxVertical SketchRelaxTypes = 128
	// SketchRelaxParallel selects Parallel.
	SketchRelaxParallel SketchRelaxTypes = 256
	// SketchRelaxPerpendicular selects Perpendicular.
	SketchRelaxPerpendicular SketchRelaxTypes = 512
	// SketchRelaxEqual selects Equal.
	SketchRelaxEqual SketchRelaxTypes = 1024
	// SketchRelaxFix selects Fix.
	SketchRelaxFix SketchRelaxTypes = 2048
)
