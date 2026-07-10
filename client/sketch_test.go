// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestSketchProjectCutEdgesSendsSketchIndex(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"created":[3,4],"healthy":true}`)}
	c := New(ft)

	got, err := c.Sketch().ProjectCutEdges(2)
	if err != nil {
		t.Fatalf("ProjectCutEdges: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchProjectCutEdges {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchProjectCutEdges)
	}
	var sent wire.ProjectCutEdgesArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.SketchIndex != 2 || len(got.Created) != 2 {
		t.Errorf("sent=%+v got=%+v, want sketch 2 / 2 created", sent, got)
	}
}

func TestSketchProjectSilhouetteSendsFaceRefAndProximity(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"created":[7],"healthy":true}`)}
	c := New(ft)

	if _, err := c.Sketch().ProjectSilhouette(0, "face-key", []float64{1, 2, 3}, true); err != nil {
		t.Fatalf("ProjectSilhouette: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchProjectSilhouette {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchProjectSilhouette)
	}
	var sent wire.ProjectSilhouetteArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.FaceRef != "face-key" || len(sent.ProximityPoint) != 3 || sent.ProximityPoint[2] != 3 || !sent.IncludeBoundary {
		t.Errorf("sent = %+v, want face-key / proximity [1,2,3] / includeBoundary", sent)
	}
}

func TestSketchSetLineTypeSendsPropertyAndValue(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"name":"Sketch1","plane":"XY","lineType":"center"}`)}
	c := New(ft)

	got, err := c.Sketch().SetLineType(0, types.SketchLineCenter)
	if err != nil {
		t.Fatalf("SetLineType: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchSetProperty {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchSetProperty)
	}
	var sent wire.SetSketchPropertyArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Property != "lineType" || sent.Value != "center" {
		t.Errorf("sent = %+v, want property lineType / value center", sent)
	}
	if got.LineType != "center" {
		t.Errorf("decoded lineType = %q, want center", got.LineType)
	}
}

func TestSketchGetSendsIndexAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":2,"name":"Sketch3","plane":"XY","visible":true,"entityCount":8,"dof":8,"editing":false,"healthy":true}`)}
	c := New(ft)

	got, err := c.Sketch().Get(2)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchGet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchGet)
	}
	var sent wire.SketchArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.SketchIndex != 2 {
		t.Errorf("sent index = %d, want 2", sent.SketchIndex)
	}
	if got.Name != "Sketch3" || got.EntityCount != 8 || got.DOF != 8 {
		t.Errorf("decoded = %+v, want Sketch3 / 8 entities / DOF 8", got)
	}
}

func TestSketchListSendsNilBodyAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"sketches":[{"index":0,"name":"Sketch1","plane":"XY","dof":0,"healthy":true}]}`)}
	c := New(ft)

	res, err := c.Sketch().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchList)
	}
	if ft.gotReq != nil {
		t.Errorf("List sent a body %q, want nil", string(ft.gotReq))
	}
	if len(res.Sketches) != 1 || res.Sketches[0].Name != "Sketch1" {
		t.Errorf("decoded = %+v, want one sketch named Sketch1", res.Sketches)
	}
}

func TestSketchAddCircleByCenterRadiusMarshalsKindVariant(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":7,"kind":"circle","pointIds":[6]}`)}
	c := New(ft)

	got, err := c.Sketch().AddCircleByCenterRadius(0, []float64{0, 0}, "10 mm", false)
	if err != nil {
		t.Fatalf("AddCircleByCenterRadius: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchAddEntity {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchAddEntity)
	}
	var sent wire.AddSketchEntityArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "circle" || sent.Variant != "centerRadius" || sent.Radius != "10 mm" || len(sent.Points) != 1 {
		t.Errorf("sent = %+v, want circle/centerRadius/10 mm/1 point", sent)
	}
	if got.EntityID != 7 || len(got.PointIDs) != 1 {
		t.Errorf("decoded = %+v, want entity 7 / 1 point id", got)
	}
}

func TestSketchAddArcByThreePointsSetsVariant(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":9,"kind":"arc","pointIds":[6,7,8]}`)}
	c := New(ft)

	if _, err := c.Sketch().AddArcByThreePoints(0, []float64{2, 0}, []float64{0, 2}, []float64{-2, 0}, false); err != nil {
		t.Fatalf("AddArcByThreePoints: %v", err)
	}
	var sent wire.AddSketchEntityArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "arc" || sent.Variant != "threePoint" || len(sent.Points) != 3 {
		t.Errorf("sent = %+v, want arc/threePoint/3 points", sent)
	}
}

func TestSketchAddEllipseMarshalsAxisAndRadii(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":5,"kind":"ellipse","pointIds":[4]}`)}
	c := New(ft)

	if _, err := c.Sketch().AddEllipse(0, []float64{0, 0}, []float64{1, 0}, "20 mm", "10 mm", false); err != nil {
		t.Fatalf("AddEllipse: %v", err)
	}
	var sent wire.AddSketchEntityArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "ellipse" || len(sent.Axis) != 2 || sent.MajorRadius != "20 mm" || sent.MinorRadius != "10 mm" {
		t.Errorf("sent = %+v, want ellipse with axis + 20mm/10mm radii", sent)
	}
}

func TestSketchConstrainParallelMarshalsKindAndRefs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"kind":"parallel","dof":3}`)}
	c := New(ft)

	got, err := c.Sketch().Constrain(0).Parallel(5, 9)
	if err != nil {
		t.Fatalf("Parallel: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchAddConstraint {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchAddConstraint)
	}
	var sent wire.AddConstraintArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "parallel" || len(sent.Entities) != 2 || sent.Entities[0] != 5 || sent.Entities[1] != 9 {
		t.Errorf("sent = %+v, want parallel of [5,9]", sent)
	}
	if got.DOF != 3 {
		t.Errorf("decoded DOF = %d, want 3", got.DOF)
	}
}

// Symmetric carries 3 refs (point A, point B, mirror line) in that order — the order the
// router resolves and the enumerate path reports (#1574).
func TestSketchConstrainSymmetricMarshalsKindAndRefs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"kind":"symmetry","dof":0}`)}
	c := New(ft)

	got, err := c.Sketch().Constrain(0).Symmetric(3, 7, 11)
	if err != nil {
		t.Fatalf("Symmetric: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchAddConstraint {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchAddConstraint)
	}
	var sent wire.AddConstraintArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "symmetry" || len(sent.Entities) != 3 ||
		sent.Entities[0] != 3 || sent.Entities[1] != 7 || sent.Entities[2] != 11 {
		t.Errorf("sent = %+v, want symmetry of [3,7,11]", sent)
	}
	if got.DOF != 0 {
		t.Errorf("decoded DOF = %d, want 0", got.DOF)
	}
}

func TestSketchDimensionRadiusMarshalsExpression(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"kind":"radius","parameter":"d0","value":2,"dof":0}`)}
	c := New(ft)

	got, err := c.Sketch().Dimension(0).Radius(7, "20 mm")
	if err != nil {
		t.Fatalf("Radius: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchAddDimension {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchAddDimension)
	}
	var sent wire.AddDimensionArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "radius" || sent.Expression != "20 mm" || len(sent.Entities) != 1 || sent.Entities[0] != 7 {
		t.Errorf("sent = %+v, want radius of circle 7 @ 20 mm", sent)
	}
	if got.Parameter != "d0" || got.Value != 2 {
		t.Errorf("decoded = %+v, want parameter d0 / value 2", got)
	}
}

func TestSketchDimensionAddWithCarriesDrivenTextPointLinearDiameter(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"kind":"offsetDim","parameter":"d0","value":6,"dof":1}`)}
	c := New(ft)

	_, err := c.Sketch().Dimension(2).AddWith(wire.AddDimensionArgs{
		Kind: "offsetDim", Entities: []uint64{3, 9}, Expression: "3 mm",
		Driven: true, TextPoint: []float64{1.5, 2}, LinearDiameter: true,
	})
	if err != nil {
		t.Fatalf("AddWith: %v", err)
	}
	var sent wire.AddDimensionArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.SketchIndex != 2 || !sent.Driven || !sent.LinearDiameter ||
		len(sent.TextPoint) != 2 || sent.TextPoint[0] != 1.5 {
		t.Errorf("sent = %+v, want sketch 2 / driven / linearDiameter / textPoint [1.5,2]", sent)
	}
}

func TestSketchDimensionOffsetSplineNamesKind(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"kind":"offsetSplineDim","parameter":"d0","value":5,"dof":0}`)}
	c := New(ft)

	if _, err := c.Sketch().Dimension(0).OffsetSpline(12, "5 mm"); err != nil {
		t.Fatalf("OffsetSpline: %v", err)
	}
	var sent wire.AddDimensionArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "offsetSplineDim" || len(sent.Entities) != 1 || sent.Entities[0] != 12 {
		t.Errorf("sent = %+v, want offsetSplineDim of entity 12", sent)
	}
}

func TestSketchProfilesDecodesArea(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"profiles":[{"index":0,"area":64,"closed":true,"holes":1}]}`)}
	c := New(ft)

	got, err := c.Sketch().Profiles(0)
	if err != nil {
		t.Fatalf("Profiles: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchProfiles {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchProfiles)
	}
	if len(got.Profiles) != 1 || got.Profiles[0].Area != 64 || got.Profiles[0].Holes != 1 {
		t.Errorf("decoded = %+v, want one profile area 64 / 1 hole", got.Profiles)
	}
}

func TestSketchSolveDecodesStatus(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"sketchIndex":0,"dof":0,"status":"well","converged":true,"healthy":true}`)}
	c := New(ft)

	got, err := c.Sketch().Solve(0)
	if err != nil {
		t.Fatalf("Solve: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchSolve {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchSolve)
	}
	if got.DOF != 0 || got.Status != "well" || !got.Converged {
		t.Errorf("decoded = %+v, want DOF 0 / well / converged", got)
	}
}
