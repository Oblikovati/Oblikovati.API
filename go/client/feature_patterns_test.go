// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

// TestPatternCircularSendsKindAndExpressionCount checks the typed circular-pattern helper
// dispatches features.add with the patternCircular kind and carries the count expression (#189).
func TestPatternCircularSendsKindAndExpressionCount(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":12,"name":"CircularPattern1"}`)}
	c := New(ft)

	if _, err := c.Features().PatternCircular(wire.CircularPatternFeatureArgs{
		SourceFeatures: []string{"Slot1"}, CountExpr: "slots", Angle: "360 deg",
	}); err != nil {
		t.Fatalf("PatternCircular: %v", err)
	}
	if ft.gotMethod != wire.MethodFeaturesAdd {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodFeaturesAdd)
	}
	var sent wire.AddFeatureArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != wire.FeatureKindPatternCircular {
		t.Errorf("kind = %q, want %q", sent.Kind, wire.FeatureKindPatternCircular)
	}
	var args wire.CircularPatternFeatureArgs
	if err := json.Unmarshal(sent.Args, &args); err != nil {
		t.Fatalf("args not valid JSON: %v", err)
	}
	if args.CountExpr != "slots" || args.Angle != "360 deg" || len(args.SourceFeatures) != 1 {
		t.Errorf("args = %+v, want countExpr slots / 360 deg / one source", args)
	}
}

// TestPatternRectangularAndMirrorSendTheirKinds covers the other two pattern helpers' routing.
func TestPatternRectangularAndMirrorSendTheirKinds(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":13}`)}
	c := New(ft)

	if _, err := c.Features().PatternRectangular(wire.RectangularPatternFeatureArgs{
		SourceFeatures: []string{"Hole1"}, CountXExpr: "cols", StepX: []float64{1, 0, 0},
	}); err != nil {
		t.Fatalf("PatternRectangular: %v", err)
	}
	var rect wire.AddFeatureArgs
	_ = json.Unmarshal(ft.gotReq, &rect)
	if rect.Kind != wire.FeatureKindPatternRectangular {
		t.Errorf("kind = %q, want %q", rect.Kind, wire.FeatureKindPatternRectangular)
	}

	if _, err := c.Features().MirrorFeatures(wire.MirrorFeatureArgs{
		SourceFeatures: []string{"Boss1"}, Normal: []float64{1, 0, 0},
	}); err != nil {
		t.Fatalf("MirrorFeatures: %v", err)
	}
	var mir wire.AddFeatureArgs
	_ = json.Unmarshal(ft.gotReq, &mir)
	if mir.Kind != wire.FeatureKindMirror {
		t.Errorf("kind = %q, want %q", mir.Kind, wire.FeatureKindMirror)
	}
}
