// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati/api/wire"
)

func TestSketch3DCreateSendsNameAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"sketchIndex":0}`)}
	c := New(ft)

	got, err := c.Sketch3D().Create(wire.CreateSketch3DArgs{Name: "Path"})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DCreate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DCreate)
	}
	var sent wire.CreateSketch3DArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Name != "Path" {
		t.Errorf("sent name = %q, want Path", sent.Name)
	}
	if got.SketchIndex != 0 {
		t.Errorf("decoded index = %d, want 0", got.SketchIndex)
	}
}

func TestSketch3DGetSendsIndexAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":1,"name":"3D Sketch2","visible":true,"dimensionsVisible":true,"entityCount":3,"dof":9,"healthy":true}`)}
	c := New(ft)

	got, err := c.Sketch3D().Get(1)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DGet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DGet)
	}
	var sent wire.Sketch3DArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.SketchIndex != 1 {
		t.Errorf("sent index = %d, want 1", sent.SketchIndex)
	}
	if got.Name != "3D Sketch2" || got.EntityCount != 3 || got.DOF != 9 {
		t.Errorf("decoded = %+v, want 3D Sketch2 / 3 entities / DOF 9", got)
	}
}

func TestSketch3DSetDimensionsVisibleSendsProperty(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"name":"3D Sketch1","dimensionsVisible":false}`)}
	c := New(ft)

	got, err := c.Sketch3D().SetDimensionsVisible(0, false)
	if err != nil {
		t.Fatalf("SetDimensionsVisible: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DSetProperty {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DSetProperty)
	}
	var sent wire.SetSketch3DPropertyArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Property != "dimensionsVisible" || sent.Value != "false" {
		t.Errorf("sent = %+v, want property dimensionsVisible / value false", sent)
	}
	if got.DimensionsVisible {
		t.Error("decoded dimensionsVisible should be false")
	}
}

func TestSketch3DListSendsNilBodyAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"sketches":[{"index":0,"name":"3D Sketch1","visible":true,"dof":0,"healthy":true}]}`)}
	c := New(ft)

	res, err := c.Sketch3D().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DList)
	}
	if ft.gotReq != nil {
		t.Errorf("List should send a nil body, got %q", ft.gotReq)
	}
	if len(res.Sketches) != 1 || res.Sketches[0].Name != "3D Sketch1" {
		t.Errorf("decoded = %+v, want one 3D Sketch1", res.Sketches)
	}
}

func TestSketch3DParallelSendsKindAndEntities(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"kind":"parallel","dof":4}`)}
	c := New(ft)

	got, err := c.Sketch3D().Parallel(0, 11, 22)
	if err != nil {
		t.Fatalf("Parallel: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DAddConstraint {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DAddConstraint)
	}
	var sent wire.AddSketch3DConstraintArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "parallel" || len(sent.Entities) != 2 || sent.Entities[0] != 11 || sent.Entities[1] != 22 {
		t.Errorf("sent = %+v, want parallel over [11,22]", sent)
	}
	if got.DOF != 4 {
		t.Errorf("decoded DOF = %d, want 4", got.DOF)
	}
}

func TestSketch3DDeleteConstraintSendsIndices(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	if _, err := c.Sketch3D().DeleteConstraint(1, 3); err != nil {
		t.Fatalf("DeleteConstraint: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DDeleteConstraint {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DDeleteConstraint)
	}
	var sent wire.DeleteSketch3DConstraintArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.SketchIndex != 1 || sent.ConstraintIndex != 3 {
		t.Errorf("sent = %+v, want sketch 1 / constraint 3", sent)
	}
}

func TestSketch3DRadiusSendsKindAndExpression(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"kind":"radius","parameter":"d3_0","value":3,"dof":3}`)}
	c := New(ft)

	got, err := c.Sketch3D().Radius(0, 42, "3 cm")
	if err != nil {
		t.Fatalf("Radius: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DAddDimension {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DAddDimension)
	}
	var sent wire.AddSketch3DDimensionArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "radius" || sent.Expression != "3 cm" || len(sent.Entities) != 1 || sent.Entities[0] != 42 {
		t.Errorf("sent = %+v, want radius over [42] = 3 cm", sent)
	}
	if got.Value != 3 {
		t.Errorf("decoded value = %v, want 3", got.Value)
	}
}

func TestSketch3DDriveDimensionSends(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	if _, err := c.Sketch3D().DriveDimension(wire.DriveSketch3DDimensionArgs{SketchIndex: 0, DimensionIndex: 2, Expression: "5 cm", SetDriven: true, Driven: true}); err != nil {
		t.Fatalf("DriveDimension: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DDriveDimension {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DDriveDimension)
	}
}

func TestSketch3DProfilesSendsIndex(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"profiles":[{"index":0,"area":12,"normal":[0,0,1],"vertices":4}]}`)}
	c := New(ft)

	res, err := c.Sketch3D().Profiles(0)
	if err != nil {
		t.Fatalf("Profiles: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DProfiles {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DProfiles)
	}
	if len(res.Profiles) != 1 || res.Profiles[0].Area != 12 {
		t.Errorf("decoded = %+v, want one profile of area 12", res.Profiles)
	}
}

func TestSketch3DPathsSendsIndex(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"paths":[{"index":0,"closed":true,"points":5}]}`)}
	c := New(ft)

	res, err := c.Sketch3D().Paths(0)
	if err != nil {
		t.Fatalf("Paths: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DPaths {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DPaths)
	}
	if len(res.Paths) != 1 || !res.Paths[0].Closed || res.Paths[0].Points != 5 {
		t.Errorf("decoded = %+v, want one closed path of 5 pts", res.Paths)
	}
}

func TestSketch3DAddSplineSendsKindAndClosed(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":7,"kind":"controlPointSpline"}`)}
	c := New(ft)

	if _, err := c.Sketch3D().AddSpline(0, [][]float64{{0, 0, 0}, {1, 0, 1}}, true, false); err != nil {
		t.Fatalf("AddSpline: %v", err)
	}
	var sent wire.AddSketch3DEntityArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "controlPointSpline" || !sent.Closed || len(sent.Points) != 2 {
		t.Errorf("sent = %+v, want closed controlPointSpline of 2 pts", sent)
	}
}

func TestSketch3DAddEquationCurveSendsExprs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":8,"kind":"equationCurve"}`)}
	c := New(ft)

	if _, err := c.Sketch3D().AddEquationCurve(0, "cos(t)", "sin(t)", "t", 0, 6.28); err != nil {
		t.Fatalf("AddEquationCurve: %v", err)
	}
	var sent wire.AddSketch3DEntityArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.XExpr != "cos(t)" || sent.ZExpr != "t" || sent.T1 != 6.28 {
		t.Errorf("sent = %+v, want the x/y/z exprs over [0,6.28]", sent)
	}
}

func TestSketch3DMoveAndCopySendOp(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"created":[9],"entityCount":2}`)}
	c := New(ft)

	res, err := c.Sketch3D().Copy(0, []uint64{3}, []float64{0, 5, 0})
	if err != nil {
		t.Fatalf("Copy: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DTransform {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DTransform)
	}
	var sent wire.Transform3DArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Op != "copy" || len(sent.Vector) != 3 || sent.Vector[1] != 5 {
		t.Errorf("sent = %+v, want copy by [0,5,0]", sent)
	}
	if len(res.Created) != 1 || res.EntityCount != 2 {
		t.Errorf("decoded = %+v, want 1 created / 2 entities", res)
	}
}

func TestSketch3DDeleteEntitiesSendsOp(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityCount":0}`)}
	c := New(ft)
	if _, err := c.Sketch3D().DeleteEntities(0, []uint64{1, 2}); err != nil {
		t.Fatalf("DeleteEntities: %v", err)
	}
	var sent wire.Transform3DArgs
	_ = json.Unmarshal(ft.gotReq, &sent)
	if sent.Op != "delete" || len(sent.Entities) != 2 {
		t.Errorf("sent = %+v, want delete of 2 entities", sent)
	}
}

func TestSketch3DIncludeSendsRefs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"created":[5],"healthy":true}`)}
	c := New(ft)

	res, err := c.Sketch3D().Include(0, []string{"edge-key-1"})
	if err != nil {
		t.Fatalf("Include: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DInclude {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DInclude)
	}
	var sent wire.IncludeSketch3DArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if len(sent.Refs) != 1 || sent.Refs[0] != "edge-key-1" {
		t.Errorf("sent = %+v, want one edge ref", sent)
	}
	if len(res.Created) != 1 || !res.Healthy {
		t.Errorf("decoded = %+v, want 1 created / healthy", res)
	}
}

func TestSketch3DAddIntersectionCurveSendsFaces(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":12,"kind":"intersection","healthy":true}`)}
	c := New(ft)

	res, err := c.Sketch3D().AddIntersectionCurve(0, "faceA", "faceB", wire.AddSketch3DSurfaceCurveArgs{GridUMax: 10})
	if err != nil {
		t.Fatalf("AddIntersectionCurve: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DAddSurfaceCurve {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DAddSurfaceCurve)
	}
	var sent wire.AddSketch3DSurfaceCurveArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "intersection" || len(sent.FaceRefs) != 2 || sent.GridUMax != 10 {
		t.Errorf("sent = %+v, want intersection over [faceA,faceB] with the grid window", sent)
	}
	if !res.Healthy || res.EntityID != 12 {
		t.Errorf("decoded = %+v, want healthy id 12", res)
	}
}

func TestSketch3DAddSilhouetteCurveSendsViewDir(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":13,"kind":"silhouette","healthy":true}`)}
	c := New(ft)
	if _, err := c.Sketch3D().AddSilhouetteCurve(0, "faceA", []float64{0, 0, 1}, wire.AddSketch3DSurfaceCurveArgs{}); err != nil {
		t.Fatalf("AddSilhouetteCurve: %v", err)
	}
	var sent wire.AddSketch3DSurfaceCurveArgs
	_ = json.Unmarshal(ft.gotReq, &sent)
	if sent.Kind != "silhouette" || len(sent.FaceRefs) != 1 || len(sent.ViewDir) != 3 {
		t.Errorf("sent = %+v, want silhouette of one face with a view dir", sent)
	}
}

func TestSketch3DAddOnFaceCurveSendsUV(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":14,"kind":"onFace","healthy":true}`)}
	c := New(ft)
	if _, err := c.Sketch3D().AddOnFaceCurve(0, "faceA", []float64{0, 0, 1, 0, 1, 1}); err != nil {
		t.Fatalf("AddOnFaceCurve: %v", err)
	}
	var sent wire.AddSketch3DSurfaceCurveArgs
	_ = json.Unmarshal(ft.gotReq, &sent)
	if sent.Kind != "onFace" || len(sent.FaceRefs) != 1 || len(sent.UV) != 6 {
		t.Errorf("sent = %+v, want onFace of one face with a 6-element uv", sent)
	}
}

func TestSketch3DAddProjectToSurfaceCurveSendsSource(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":15,"kind":"projectToSurface","healthy":true}`)}
	c := New(ft)
	if _, err := c.Sketch3D().AddProjectToSurfaceCurve(0, 7, "faceA"); err != nil {
		t.Fatalf("AddProjectToSurfaceCurve: %v", err)
	}
	var sent wire.AddSketch3DSurfaceCurveArgs
	_ = json.Unmarshal(ft.gotReq, &sent)
	if sent.Kind != "projectToSurface" || sent.SourceEntityID != 7 || len(sent.FaceRefs) != 1 {
		t.Errorf("sent = %+v, want projectToSurface of source 7 onto one face", sent)
	}
}

func TestSketch3DAddOffsetCurveSendsDistance(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":16,"kind":"offset","healthy":true}`)}
	c := New(ft)
	if _, err := c.Sketch3D().AddOffsetCurve(0, 7, 2.5, []float64{0, 0, 1}); err != nil {
		t.Fatalf("AddOffsetCurve: %v", err)
	}
	var sent wire.AddSketch3DSurfaceCurveArgs
	_ = json.Unmarshal(ft.gotReq, &sent)
	if sent.Kind != "offset" || sent.SourceEntityID != 7 || sent.OffsetDistance != 2.5 || len(sent.Normal) != 3 {
		t.Errorf("sent = %+v, want offset of source 7 by 2.5 with a normal", sent)
	}
}
