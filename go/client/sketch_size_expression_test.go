// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"errors"
	"testing"

	"oblikovati.org/api/wire"
)

// legacySketchHost simulates a host that has no associative constructor.
type legacySketchHost struct{}

func (legacySketchHost) Call(method string, _ []byte) ([]byte, error) {
	if method == wire.MethodSketchAddAssociativeEntity {
		return nil, errors.New("unknown method")
	}
	return []byte(`{"entityId":1,"kind":"circle","pointIds":[]}`), nil
}

func TestAssociativeEntityRejectsLegacyHost(t *testing.T) {
	_, err := New(legacySketchHost{}).Sketch().AddAssociativeEntity(wire.AddSketchEntityArgs{
		Kind: "circle", Points: [][]float64{{0, 0}}, Radius: "bore_r",
	})
	if err == nil {
		t.Fatal("expected an unknown-method error from a legacy host")
	}
}

func TestAssociativeEntitySendsExpressionToItsOwnMethod(t *testing.T) {
	host := &fakeTransport{reply: []byte(`{"entityId":1,"kind":"circle","pointIds":[]}`)}
	_, err := New(host).Sketch().AddAssociativeEntity(wire.AddSketchEntityArgs{
		Kind: "circle", Points: [][]float64{{0, 0}}, Radius: "bore_r + 2 mm",
	})
	if err != nil {
		t.Fatal(err)
	}
	var sent wire.AddSketchEntityArgs
	if err := json.Unmarshal(host.gotReq, &sent); err != nil {
		t.Fatal(err)
	}
	if host.gotMethod != wire.MethodSketchAddAssociativeEntity || sent.Radius != "bore_r + 2 mm" {
		t.Fatalf("method %q, radius %q: expected associative method and unchanged expression", host.gotMethod, sent.Radius)
	}
}

// checkSketchSizeRequest verifies the typed request at the host boundary.
func checkSketchSizeRequest(t *testing.T, call func(Sketch) error, want map[string]string) {
	t.Helper()
	host := &fakeTransport{reply: []byte(`{"entityId":1,"kind":"line","pointIds":[]}`)}
	if err := call(New(host).Sketch()); err != nil {
		t.Fatal(err)
	}
	if host.gotMethod != wire.MethodSketchAddEntity {
		t.Fatalf("method %q, want %q", host.gotMethod, wire.MethodSketchAddEntity)
	}
	var sent map[string]json.RawMessage
	if err := json.Unmarshal(host.gotReq, &sent); err != nil {
		t.Fatal(err)
	}
	for field, value := range want {
		checkExpressionField(t, sent[field], field, value)
	}
}

func checkExpressionField(t *testing.T, raw json.RawMessage, field, expected string) {
	t.Helper()
	var actual string
	if err := json.Unmarshal(raw, &actual); err != nil {
		t.Fatalf("field %q: expected a string: %v", field, err)
	}
	if actual != expected {
		t.Errorf("field %q = %q, want %q", field, actual, expected)
	}
}

func TestCircleHelperPreservesExpression(t *testing.T) {
	checkSketchSizeRequest(t, func(s Sketch) error {
		_, err := s.AddCircleByCenterRadius(0, []float64{0, 0}, "bore_r + 2 mm", false)
		return err
	}, map[string]string{"radius": "bore_r + 2 mm"})
}

func TestCircleHelperPreservesLiteralRequest(t *testing.T) {
	checkSketchSizeRequest(t, func(s Sketch) error {
		_, err := s.AddCircleByCenterRadius(0, []float64{0, 0}, "10 mm", false)
		return err
	}, map[string]string{"radius": "10 mm"})
}

func TestCircleLiteralWireShapeRemainsCompatible(t *testing.T) {
	host := &fakeTransport{reply: []byte(`{"entityId":1,"kind":"circle","pointIds":[]}`)}
	_, err := New(host).Sketch().AddCircleByCenterRadius(0, []float64{0, 0}, "10 mm", false)
	if err != nil {
		t.Fatal(err)
	}
	const expected = `{"sketchIndex":0,"kind":"circle","variant":"centerRadius","points":[[0,0]],"radius":"10 mm"}`
	if string(host.gotReq) != expected {
		t.Fatalf("request %s: expected %s", host.gotReq, expected)
	}
}
func TestEllipseHelperPreservesExpressions(t *testing.T) {
	checkSketchSizeRequest(t, func(s Sketch) error {
		_, err := s.AddEllipse(0, []float64{0, 0}, []float64{1, 0}, "major_r", "minor_r", false)
		return err
	}, map[string]string{"majorRadius": "major_r", "minorRadius": "minor_r"})
}
func TestSlotHelperPreservesExpression(t *testing.T) {
	checkSketchSizeRequest(t, func(s Sketch) error {
		_, err := s.AddSlot(0, []float64{0, 0}, []float64{1, 0}, "slot_width", false)
		return err
	}, map[string]string{"width": "slot_width"})
}
func TestFilletHelperPreservesExpression(t *testing.T) {
	checkSketchSizeRequest(t, func(s Sketch) error {
		_, err := s.AddFillet(0, 1, 2, "corner_radius")
		return err
	}, map[string]string{"radius": "corner_radius"})
}
func TestChamferHelperPreservesExpressions(t *testing.T) {
	checkSketchSizeRequest(t, func(s Sketch) error {
		_, err := s.AddChamfer(0, 1, 2, "width_a", "width_b")
		return err
	}, map[string]string{"radius": "width_a", "distance2": "width_b"})
}
func TestOffsetSplineHelperPreservesExpression(t *testing.T) {
	checkSketchSizeRequest(t, func(s Sketch) error {
		_, err := s.AddOffsetSpline(0, 4, "offset_width")
		return err
	}, map[string]string{"radius": "offset_width"})
}
