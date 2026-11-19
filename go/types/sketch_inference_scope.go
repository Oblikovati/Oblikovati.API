// SPDX-License-Identifier: Apache-2.0

package types

// SketchInferenceScope identifies the geometry eligible for 2D sketch inference.
// Example: `scope := types.InferenceScopeAllGeometry`.
type SketchInferenceScope string

const (
	// InferenceScopeCurrentCommand uses geometry made by the active command.
	InferenceScopeCurrentCommand SketchInferenceScope = "currentCommand"
	// InferenceScopeAllGeometry uses all eligible geometry in the edited sketch.
	InferenceScopeAllGeometry SketchInferenceScope = "allGeometry"
	// InferenceScopeSelectedGeometry uses the explicitly selected sketch entities.
	InferenceScopeSelectedGeometry SketchInferenceScope = "selectedGeometry"
)
