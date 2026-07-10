// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// TestAddLineExprSendsPointExprs checks the expression-form line constructor sends pointExprs
// (not points) so the endpoints are parameter-driven at construction (#189).
func TestAddLineExprSendsPointExprs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":4,"kind":"line","pointIds":[2,3]}`)}
	c := New(ft)

	if _, err := c.Sketch().AddLineExpr(0, []string{"0", "0"}, []string{"bore_r", "0"}, false); err != nil {
		t.Fatalf("AddLineExpr: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchAddEntity {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchAddEntity)
	}
	var sent wire.AddSketchEntityArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != string(types.SketchEntityLine) {
		t.Errorf("kind = %q, want line", sent.Kind)
	}
	if len(sent.Points) != 0 {
		t.Errorf("points = %v, want none (expression form)", sent.Points)
	}
	if len(sent.PointExprs) != 2 || sent.PointExprs[1][0] != "bore_r" {
		t.Errorf("pointExprs = %v, want [[0 0] [bore_r 0]]", sent.PointExprs)
	}
}

// TestAddArcByCenterStartEndExprCarriesVariantAndExprs covers the expression-form arc.
func TestAddArcByCenterStartEndExprCarriesVariantAndExprs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":9,"kind":"arc","pointIds":[6,7,8]}`)}
	c := New(ft)

	if _, err := c.Sketch().AddArcByCenterStartEndExpr(0,
		[]string{"0", "0"}, []string{"r", "0"}, []string{"0", "r"}, true, false); err != nil {
		t.Fatalf("AddArcByCenterStartEndExpr: %v", err)
	}
	var sent wire.AddSketchEntityArgs
	_ = json.Unmarshal(ft.gotReq, &sent)
	if sent.Variant != "centerStartEnd" || !sent.CCW {
		t.Errorf("sent = %+v, want centerStartEnd / ccw", sent)
	}
	if len(sent.PointExprs) != 3 || sent.PointExprs[2][1] != "r" {
		t.Errorf("pointExprs = %v, want 3 expression points ending [0 r]", sent.PointExprs)
	}
}

// TestEventDispatcherRoutesFeatureAndSketchEvents covers the #148 wave: each feature-lifecycle
// and sketch-edit tag reaches its typed callback.
func TestEventDispatcherRoutesFeatureAndSketchEvents(t *testing.T) {
	d := NewEventDispatcher()
	var feats []wire.FeatureLifecycleEvent
	var sketches []wire.SketchEditEvent
	d.OnFeature(func(e wire.FeatureLifecycleEvent) { feats = append(feats, e) })
	d.OnSketchEdit(func(e wire.SketchEditEvent) { sketches = append(sketches, e) })

	events := [][]byte{
		[]byte(`{"type":"feature.added","document":3,"feature":7,"name":"Extrusion1","kind":"extrude"}`),
		[]byte(`{"type":"feature.edited","document":3,"feature":7,"name":"Extrusion1","kind":"extrude"}`),
		[]byte(`{"type":"feature.deleted","document":3,"feature":7}`),
		[]byte(`{"type":"sketch.editEntered","document":3,"sketch":4,"name":"Sketch1"}`),
		[]byte(`{"type":"sketch.editExited","document":3,"sketch":4,"name":"Sketch1"}`),
	}
	for _, ev := range events {
		if !d.Dispatch(ev) {
			t.Errorf("Dispatch(%s) = false, want routed", ev)
		}
	}
	if len(feats) != 3 || feats[0].Type != wire.EventFeatureAdded || feats[0].Feature != 7 || feats[0].Kind != "extrude" {
		t.Errorf("feature callbacks = %+v, want 3 with first added/feature 7/extrude", feats)
	}
	if len(sketches) != 2 || sketches[0].Type != wire.EventSketchEditEntered || sketches[0].Sketch != 4 {
		t.Errorf("sketch callbacks = %+v, want entered then exited on sketch 4", sketches)
	}
}

// TestAddCenterlineMarksCenterlineAndConstruction checks the centerline constructor sends a line
// with both centerline and construction set — so a procedural add-in can revolve about the
// sketch's internal axis (the tapered-roller domed body, PartDesigner #54).
func TestAddCenterlineMarksCenterlineAndConstruction(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":5,"kind":"line","pointIds":[3,4]}`)}
	c := New(ft)

	if _, err := c.Sketch().AddCenterline(0, []float64{0, 0}, []float64{0, 1}); err != nil {
		t.Fatalf("AddCenterline: %v", err)
	}
	if ft.gotMethod != wire.MethodSketchAddEntity {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodSketchAddEntity)
	}
	var sent wire.AddSketchEntityArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != string(types.SketchEntityLine) {
		t.Errorf("kind = %q, want line", sent.Kind)
	}
	if !sent.Centerline || !sent.Construction {
		t.Errorf("sent = %+v, want centerline=true construction=true", sent)
	}
}
