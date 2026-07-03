// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
	"oblikovati.org/api/wire/featureargs"
)

// TestAddFeatureTagsEnvelopeFromKind proves the typed constructor tags the wire envelope
// with the arg's own Kind and marshals the struct into Args — so an add-in builds a feature
// with a compile-checked type, never a raw JSON blob (#1616).
func TestAddFeatureTagsEnvelopeFromKind(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"feature":"Extrusion1"}`)}
	c := New(ft)

	if _, err := AddFeature(c.Features(), featureargs.Extrude{SketchIndex: 2, Distance: "50 mm"}); err != nil {
		t.Fatalf("AddFeature: %v", err)
	}
	if ft.gotMethod != wire.MethodFeaturesAdd {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodFeaturesAdd)
	}
	var env wire.AddFeatureArgs
	if err := json.Unmarshal(ft.gotReq, &env); err != nil {
		t.Fatalf("envelope not valid JSON: %v", err)
	}
	if env.Kind != featureargs.KindExtrude {
		t.Errorf("envelope kind = %q, want %q", env.Kind, featureargs.KindExtrude)
	}
	var sent featureargs.Extrude
	if err := json.Unmarshal(env.Args, &sent); err != nil {
		t.Fatalf("args not valid Extrude JSON: %v", err)
	}
	if sent.SketchIndex != 2 || sent.Distance != "50 mm" {
		t.Errorf("sent args = %+v, want sketchIndex=2 distance=50 mm", sent)
	}
}
