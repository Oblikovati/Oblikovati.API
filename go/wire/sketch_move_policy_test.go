// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"
	"strings"
	"testing"

	"oblikovati.org/api/types"
)

func TestMovePoliciesAreOptionalAndRoundTrip(t *testing.T) {
	old := TransformSketchArgs{Op: "move", Entities: []uint64{1}, Vector: []float64{1, 2}}
	encoded, err := json.Marshal(old)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), "Policy") {
		t.Fatalf("unexpected policy in %s", encoded)
	}
	for _, dimension := range []types.SketchMoveDimensionPolicy{types.SketchMoveDimensionNever, types.SketchMoveDimensionIfNoEquation, types.SketchMoveDimensionAlways, types.SketchMoveDimensionPrompt} {
		for _, geometric := range []types.SketchMoveGeometricPolicy{types.SketchMoveGeometricNever, types.SketchMoveGeometricAlways, types.SketchMoveGeometricPrompt} {
			request := old
			request.DimensionPolicy, request.GeometricPolicy = dimension, geometric
			encoded, err := json.Marshal(request)
			if err != nil {
				t.Fatal(err)
			}
			var decoded TransformSketchArgs
			if err := json.Unmarshal(encoded, &decoded); err != nil {
				t.Fatal(err)
			}
			if decoded.DimensionPolicy != dimension || decoded.GeometricPolicy != geometric {
				t.Fatalf("policies %q, %q became %q, %q", dimension, geometric, decoded.DimensionPolicy, decoded.GeometricPolicy)
			}
		}
	}
}

func TestMoveDecisionRoundTripsWithoutCreatedEntities(t *testing.T) {
	response := TransformSketchResult{DecisionRequired: &SketchMoveDecision{Dimension: true, Geometric: true}}
	encoded, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	var decoded TransformSketchResult
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.DecisionRequired == nil || !decoded.DecisionRequired.Dimension || !decoded.DecisionRequired.Geometric || len(decoded.Created) != 0 {
		t.Fatalf("decision response %s did not round trip: %+v", encoded, decoded)
	}
}
