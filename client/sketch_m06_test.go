// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestCreateBlockDefinitionSendsSelectionAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":0,"name":"fastener","entityCount":3,"instanceCount":1}`)}
	c := New(ft)

	got, err := c.Sketch().CreateBlockDefinition(wire.CreateBlockDefinitionArgs{
		Name: "fastener", SourceSketchIndex: 2, EntityRefs: []uint64{7, 8, 9},
	})
	if err != nil {
		t.Fatalf("CreateBlockDefinition: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchBlockDefinitionCreate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchBlockDefinitionCreate)
	}
	var sent wire.CreateBlockDefinitionArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Name != "fastener" || len(sent.EntityRefs) != 3 {
		t.Errorf("sent = %+v, want fastener with 3 entity refs", sent)
	}
	if got.EntityCount != 3 || got.InstanceCount != 1 {
		t.Errorf("decoded = %+v, want 3 entities / 1 instance", got)
	}
}

func TestAddBlockInstanceSendsPlacement(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":42}`)}
	c := New(ft)

	got, err := c.Sketch().AddBlockInstance(wire.AddSketchBlockArgs{
		SketchIndex: 1, Definition: "fastener",
		Position: []float64{2, 3}, RotationAngle: "45 deg", Scale: 2,
	})
	if err != nil {
		t.Fatalf("AddBlockInstance: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchAddBlockInstance {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchAddBlockInstance)
	}
	var sent wire.AddSketchBlockArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Definition != "fastener" || sent.RotationAngle != "45 deg" || sent.Scale != 2 {
		t.Errorf("sent = %+v, want the fastener placement", sent)
	}
	if got.EntityID != 42 {
		t.Errorf("decoded entity id = %d, want 42", got.EntityID)
	}
}

func TestBlockDefinitionListAndInstances(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"definitions":[{"index":0,"name":"fastener"}]}`)}
	c := New(ft)
	defs, err := c.Sketch().BlockDefinitions()
	if err != nil || len(defs.Definitions) != 1 || defs.Definitions[0].Name != "fastener" {
		t.Fatalf("BlockDefinitions = (%+v, %v), want the fastener row", defs, err)
	}
	if ft.gotMethod != wire.MethodSketchBlockDefinitionList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchBlockDefinitionList)
	}

	ft.reply = []byte(`{"instances":[{"index":0,"id":42,"definition":"fastener","position":[2,3],"rotation":0.5,"scale":2,"entityCount":3}]}`)
	inst, err := c.Sketch().BlockInstances(1)
	if err != nil || len(inst.Instances) != 1 || inst.Instances[0].ID != 42 {
		t.Fatalf("BlockInstances = (%+v, %v), want the placed instance", inst, err)
	}
	if ft.gotMethod != wire.MethodSketchListBlockInstances {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchListBlockInstances)
	}

	ft.reply = []byte(`{"ok":true}`)
	if _, err := c.Sketch().DeleteBlockDefinition("fastener"); err != nil {
		t.Fatalf("DeleteBlockDefinition: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchBlockDefinitionDelete {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchBlockDefinitionDelete)
	}
}

func TestRegionPropertiesSendsAccuracySpelling(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"area":12,"perimeter":14,"centroid":[1,2],"momentsOfInertia":[4,5,0],"principalMoments":[5,4],"rotationAngle":1.5707963,"principalAxes":[[0,1],[-1,0]],"accuracy":"veryHigh"}`)}
	c := New(ft)

	got, err := c.Sketch().RegionProperties(0, 1, types.AccuracyVeryHigh)
	if err != nil {
		t.Fatalf("RegionProperties: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchRegionProperties {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchRegionProperties)
	}
	var sent wire.RegionPropertiesArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ProfileIndex != 1 || sent.Accuracy != "veryHigh" {
		t.Errorf("sent = %+v, want profile 1 at veryHigh", sent)
	}
	if got.Area != 12 || len(got.MomentsOfInertia) != 3 {
		t.Errorf("decoded = %+v, want area 12 with 3 moments", got)
	}

	if _, err := c.Sketch3D().RegionProperties(0, 1, types.AccuracyLow); err != nil {
		t.Fatalf("Sketch3D RegionProperties: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DRegionProperties {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DRegionProperties)
	}
}

func TestInferenceOptionsRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"inferEnabled":true,"constrainEnabled":false,"priority":"parallelPerpendicular"}`)}
	c := New(ft)

	got, err := c.Sketch().SetInferenceOptions(wire.InferenceOptionsView{
		InferEnabled: true, Priority: types.PriorityParallelPerpendicular.String(),
	})
	if err != nil {
		t.Fatalf("SetInferenceOptions: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchSetInferenceOptions {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchSetInferenceOptions)
	}
	if !got.InferEnabled || got.ConstrainEnabled || got.Priority != "parallelPerpendicular" {
		t.Errorf("decoded = %+v, want infer-on constrain-off parallelPerpendicular", got)
	}

	if _, err := c.Sketch().InferenceOptions(); err != nil {
		t.Fatalf("InferenceOptions: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchGetInferenceOptions {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchGetInferenceOptions)
	}
}

func TestSetSplineHandleSendsTangentBothDomains(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"handleId":9,"tangent":[1,0],"weight":1.5}`)}
	c := New(ft)

	got, err := c.Sketch().SetSplineHandle(wire.SetSplineHandleArgs{
		SketchIndex: 0, Spline: 5, FitPointIndex: 1, Active: true,
		Tangent: []float64{1, 0}, Weight: 1.5,
	})
	if err != nil {
		t.Fatalf("SetSplineHandle: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchSetSplineHandle {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchSetSplineHandle)
	}
	if got.HandleID != 9 || got.Weight != 1.5 {
		t.Errorf("decoded = %+v, want handle 9 at weight 1.5", got)
	}

	if _, err := c.Sketch3D().SetSplineHandle(wire.SetSplineHandleArgs{Spline: 5, Active: true}); err != nil {
		t.Fatalf("Sketch3D SetSplineHandle: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DSetSplineHandle {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DSetSplineHandle)
	}
}

func TestVariableHelixAndEditDefinition(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":7,"kind":"helical"}`)}
	c := New(ft)

	rows := []wire.HelixShapeRow{
		{Diameter: "10 mm", Pitch: "2 mm", Revolution: 0},
		{Diameter: "20 mm", Pitch: "4 mm", Revolution: 5},
	}
	_, err := c.Sketch3D().AddVariableHelix(0, []float64{0, 0, 0}, nil, "5 mm", rows,
		wire.AddSketch3DEntityArgs{Revolutions: 5})
	if err != nil {
		t.Fatalf("AddVariableHelix: %v", err)
	}
	var sent wire.AddSketch3DEntityArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "helical" || len(sent.Rows) != 2 || sent.Rows[1].Diameter != "20 mm" {
		t.Errorf("sent = %+v, want a helical kind with 2 rows", sent)
	}

	ft.reply = []byte(`{"shapeKind":"pitchRevolution","variable":true,"rows":[{"diameter":"10 mm"}],"start":{"kind":"flat","transitionAngle":"90 deg"}}`)
	view, err := c.Sketch3D().EditHelixDefinition(wire.EditHelixArgs{
		SketchIndex: 0, Entity: 7,
		Start: &wire.HelixEndCondition{Kind: "flat", TransitionAngle: "90 deg"},
	})
	if err != nil {
		t.Fatalf("EditHelixDefinition: %v", err)
	}
	if ft.gotMethod != wire.MethodSketch3DEditHelix {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketch3DEditHelix)
	}
	if !view.Variable || view.Start == nil || view.Start.Kind != "flat" {
		t.Errorf("decoded = %+v, want a variable definition with a flat start", view)
	}
}
