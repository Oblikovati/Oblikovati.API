// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"
	"strings"
	"testing"

	"oblikovati.org/api/types"
)

func TestMovePlanRetainsExplicitAndUnknownPolicies(t *testing.T) {
	input := PlanSketchMoveArgs{Entities: []uint64{1}, Vector: []float64{1, 2}, DimensionPolicy: types.SketchMoveDimensionIfNoEquation, GeometricPolicy: types.SketchMoveGeometricPolicy("future")}
	encoded, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	var output PlanSketchMoveArgs
	if err := json.Unmarshal(encoded, &output); err != nil {
		t.Fatal(err)
	}
	if output.DimensionPolicy != input.DimensionPolicy || output.GeometricPolicy != input.GeometricPolicy {
		t.Fatalf("request %s decoded to %+v: expected both policy values unchanged", encoded, output)
	}
}

func TestMovePlanEmptyPolicyIsNotConvertedToNever(t *testing.T) {
	encoded, err := json.Marshal(PlanSketchMoveArgs{})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(encoded), `"dimensionPolicy":""`) || !strings.Contains(string(encoded), `"geometricPolicy":""`) {
		t.Fatalf("request %s: expected explicit empty policies for host rejection", encoded)
	}
}

func TestMovePlanResponseKeepsAllImpactsAndApprovals(t *testing.T) {
	input := PlanSketchMoveResult{Token: "opaque", Dimensions: []SketchMoveDimensionChange{{Index: 2, BeforeExpression: "10 mm", AfterExpression: "12 mm", ApprovalRequired: true}}, Geometric: []SketchMoveGeometricChange{{Index: 4, Kind: "coincident", Entities: []uint64{1, 2}, ApprovalRequired: true}}}
	encoded, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	var output PlanSketchMoveResult
	if err := json.Unmarshal(encoded, &output); err != nil {
		t.Fatal(err)
	}
	if output.Token != input.Token || len(output.Dimensions) != 1 || len(output.Geometric) != 1 || !output.Dimensions[0].ApprovalRequired || !output.Geometric[0].ApprovalRequired {
		t.Fatalf("plan %s decoded to %+v: expected each approval-required impact", encoded, output)
	}
}
