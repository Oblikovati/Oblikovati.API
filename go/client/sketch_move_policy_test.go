// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"errors"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestPlanMoveSendsExplicitPoliciesAndReceivesImpacts(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"token":"opaque","dimensions":[{"index":2,"beforeExpression":"10 mm","afterExpression":"12 mm","approvalRequired":true}],"geometric":[]}`)}
	args := wire.PlanSketchMoveArgs{SketchIndex: 1, Entities: []uint64{3}, Vector: []float64{1, 2}, DimensionPolicy: types.SketchMoveDimensionPrompt, GeometricPolicy: types.SketchMoveGeometricNever}
	plan, err := New(ft).Sketch().PlanMove(args)
	if err != nil {
		t.Fatal(err)
	}
	var sent wire.PlanSketchMoveArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatal(err)
	}
	if ft.gotMethod != wire.MethodSketchPlanMove || sent.DimensionPolicy != args.DimensionPolicy || sent.GeometricPolicy != args.GeometricPolicy || plan.Token != "opaque" || len(plan.Dimensions) != 1 || !plan.Dimensions[0].ApprovalRequired {
		t.Fatalf("method %q request %+v response %+v: expected explicit policies and approval impact", ft.gotMethod, sent, plan)
	}
}

func TestCommitMoveSendsTokenAndApprovals(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"committed":true}`)}
	result, err := New(ft).Sketch().CommitMove(wire.CommitSketchMoveArgs{Token: "opaque", ApproveDimensions: true})
	if err != nil {
		t.Fatal(err)
	}
	var sent wire.CommitSketchMoveArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatal(err)
	}
	if ft.gotMethod != wire.MethodSketchCommitMove || sent.Token != "opaque" || !sent.ApproveDimensions || sent.ApproveGeometric || !result.Committed {
		t.Fatalf("method %q request %+v response %+v: expected committed approved plan", ft.gotMethod, sent, result)
	}
}

func TestPlanMovePropagatesUnsupportedHost(t *testing.T) {
	unsupported := errors.New("unknown method")
	ft := &fakeTransport{err: unsupported}
	_, err := New(ft).Sketch().PlanMove(wire.PlanSketchMoveArgs{})
	if !errors.Is(err, unsupported) || ft.gotMethod != wire.MethodSketchPlanMove {
		t.Fatalf("method %q error %v: expected unsupported new method", ft.gotMethod, err)
	}
}

func TestCommitMovePropagatesStalePlan(t *testing.T) {
	stale := errors.New("stale plan")
	ft := &fakeTransport{err: stale}
	result, err := New(ft).Sketch().CommitMove(wire.CommitSketchMoveArgs{Token: "expired"})
	if !errors.Is(err, stale) || result.Committed || ft.gotMethod != wire.MethodSketchCommitMove {
		t.Fatalf("method %q result %+v error %v: expected rejected stale plan", ft.gotMethod, result, err)
	}
}
