// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// TestColorSchemesSetActiveMarshals checks the color-scheme group sends the right method and
// name and decodes the active-scheme reply.
func TestColorSchemesSetActiveMarshals(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"name":"Presentation","active":true,"backgroundType":52738}`)}
	c := New(ft)
	got, err := c.ColorSchemes().SetActive("Presentation")
	if err != nil {
		t.Fatalf("SetActive: %v", err)
	}
	if ft.gotMethod != wire.MethodColorSchemesSetActive {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodColorSchemesSetActive)
	}
	var sent wire.SetColorSchemeArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not JSON: %v", err)
	}
	if sent.Name != "Presentation" {
		t.Errorf("sent name = %q, want Presentation", sent.Name)
	}
	if !got.Active || got.BackgroundType != types.GradientBackground {
		t.Errorf("decoded = %+v, want active gradient", got)
	}
}

// TestDisplaySetSettingsMarshals checks the display group wraps the settings view in the args.
func TestDisplaySetSettingsMarshals(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"backgroundType":52737,"groundShadow":69122}`)}
	c := New(ft)
	_, err := c.Display().SetSettings(wire.DisplaySettingsView{
		BackgroundType: types.OneColorBackground, GroundShadow: types.GroundShadow,
	})
	if err != nil {
		t.Fatalf("SetSettings: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentSetDisplaySettings {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentSetDisplaySettings)
	}
	var sent wire.SetDisplaySettingsArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not JSON: %v", err)
	}
	if sent.Settings.BackgroundType != types.OneColorBackground {
		t.Errorf("sent background = %v, want OneColor", sent.Settings.BackgroundType)
	}
}

// TestViewsCaptureNamedMarshals checks the named-view capture sends the name and document.
func TestViewsCaptureNamedMarshals(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"name":"iso","camera":{"eye":[1,1,1],"target":[0,0,0],"up":[0,1,0],"fov":1}}`)}
	c := New(ft)
	got, err := c.Views().CaptureNamed(wire.CaptureNamedViewArgs{Name: "iso"})
	if err != nil {
		t.Fatalf("CaptureNamed: %v", err)
	}
	if ft.gotMethod != wire.MethodViewsCaptureNamed {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodViewsCaptureNamed)
	}
	if got.Name != "iso" {
		t.Errorf("decoded name = %q, want iso", got.Name)
	}
}

// TestStylesSetMarshals checks the style group sends the color-style view.
func TestStylesSetMarshals(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"name":"Steel","opacity":1}`)}
	c := New(ft)
	_, err := c.Styles().Set(wire.ColorStyleView{Name: "Steel", Diffuse: types.NewColor(180, 180, 190), Opacity: 1})
	if err != nil {
		t.Fatalf("Set: %v", err)
	}
	if ft.gotMethod != wire.MethodStylesSet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodStylesSet)
	}
	var sent wire.ColorStyleView
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not JSON: %v", err)
	}
	if sent.Name != "Steel" || sent.Diffuse.R != 180 {
		t.Errorf("sent = %+v, want Steel diffuse.R=180", sent)
	}
}

// TestGraphicsAddBodyOverlayMarshals checks the body-overlay builder produces a surface
// primitive carrying the body key under the persistent lane.
func TestGraphicsAddBodyOverlayMarshals(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"clientId":"hi","nodeCount":1,"primitiveCount":1}`)}
	c := New(ft)
	_, err := c.Graphics().AddBodyOverlay("hi", "FACEKEY==", []float32{1, 0, 0, 1})
	if err != nil {
		t.Fatalf("AddBodyOverlay: %v", err)
	}
	if ft.gotMethod != wire.MethodClientGraphicsSet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodClientGraphicsSet)
	}
	var sent wire.SetClientGraphicsArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not JSON: %v", err)
	}
	p := sent.Nodes[0].Primitives[0]
	if p.Kind != string(types.GraphicsSurface) || p.BodyKey != "FACEKEY==" {
		t.Errorf("primitive = %+v, want surface with body key", p)
	}
}

// TestGraphicsSetNodeTransformMarshals checks the targeted node-move method bypasses a full
// group resubmit.
func TestGraphicsSetNodeTransformMarshals(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{}`)}
	c := New(ft)
	if err := c.Graphics().SetNodeTransform("g", "n1", []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}); err != nil {
		t.Fatalf("SetNodeTransform: %v", err)
	}
	if ft.gotMethod != wire.MethodGraphicsNodeSetTransform {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodGraphicsNodeSetTransform)
	}
	var sent wire.SetNodeTransformArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not JSON: %v", err)
	}
	if sent.NodeId != "n1" || len(sent.Transform) != 16 {
		t.Errorf("sent = %+v, want node n1 with 16-elem transform", sent)
	}
}
