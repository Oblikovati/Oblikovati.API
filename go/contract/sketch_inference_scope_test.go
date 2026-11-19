// SPDX-License-Identifier: Apache-2.0

package contract

import (
	"testing"

	"oblikovati.org/api/types"
)

type fakeSketchInferenceScope struct {
	mode  types.SketchInferenceScope
	count int
}

func (scope fakeSketchInferenceScope) InferenceScope() types.SketchInferenceScope {
	return scope.mode
}

func (scope fakeSketchInferenceScope) InferenceScopeEntityCount() int {
	return scope.count
}

var _ SketchInferenceScope = fakeSketchInferenceScope{}

func TestSketchInferenceScopeScalarView(t *testing.T) {
	var scope SketchInferenceScope = fakeSketchInferenceScope{mode: types.InferenceScopeSelectedGeometry, count: 2}
	if scope.InferenceScope() != types.InferenceScopeSelectedGeometry || scope.InferenceScopeEntityCount() != 2 {
		t.Fatalf("scope %q count %d: expected selected geometry with two entities", scope.InferenceScope(), scope.InferenceScopeEntityCount())
	}
}
