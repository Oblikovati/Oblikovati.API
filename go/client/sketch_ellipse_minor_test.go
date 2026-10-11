// SPDX-License-Identifier: Apache-2.0
package client

import (
	"encoding/json"
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
	"reflect"
	"testing"
)

func TestEllipseMinorRadiusUsesCanonicalKindAndOwnerIndex(t *testing.T) {
	host := &fakeTransport{reply: []byte(`{"index":3,"kind":"ellipseMinorRadius","parameter":"d3","value":0.5,"dof":0}`)}
	result, err := New(host).Sketch().Dimension(2).EllipseMinorRadius(42, "5 mm")
	if err != nil {
		t.Fatal(err)
	}
	if host.gotMethod != wire.MethodSketchAddDimension {
		t.Fatalf("method %q: expected %q", host.gotMethod, wire.MethodSketchAddDimension)
	}
	assertMinorRadiusRequest(t, host.gotReq)
	if result.Kind != string(types.DimConstraintEllipseMinorRadius) || result.DOF != 0 {
		t.Fatalf("reply %+v: expected minor radius and zero DOF", result)
	}
}
func assertMinorRadiusRequest(t *testing.T, payload []byte) {
	t.Helper()
	var request wire.AddDimensionArgs
	if err := json.Unmarshal(payload, &request); err != nil {
		t.Fatal(err)
	}
	expected := wire.AddDimensionArgs{SketchIndex: 2, Kind: string(types.DimConstraintEllipseMinorRadius), Entities: []uint64{42}, Expression: "5 mm"}
	if !reflect.DeepEqual(request, expected) {
		t.Fatalf("request %+v: expected %+v", request, expected)
	}
	if types.DimConstraintEllipseRadius != "ellipseRadius" {
		t.Fatalf("major radius kind %q: expected ellipseRadius", types.DimConstraintEllipseRadius)
	}
}
