// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// SketchInferenceScope is the scalar view of a 2D sketch's session inference scope.
// Example: `mode := scope.InferenceScope()`.
type SketchInferenceScope interface {
	// InferenceScope reports the active candidate-geometry mode.
	InferenceScope() types.SketchInferenceScope
	// InferenceScopeEntityCount reports the number of selected candidates.
	InferenceScopeEntityCount() int
}
