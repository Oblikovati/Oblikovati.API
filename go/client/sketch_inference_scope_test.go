// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"reflect"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestInferenceScopeRoundTrip(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"sketchIndex":2,"scope":"selectedGeometry","entities":[4,9]}`)}
	client := New(transport)
	request := wire.SketchInferenceScopeView{SketchIndex: 2, Scope: types.InferenceScopeSelectedGeometry, Entities: []uint64{4, 9}}
	result, err := client.Sketch().SetInferenceScope(request)
	if err != nil {
		t.Fatalf("SetInferenceScope: %v", err)
	}
	if transport.gotMethod != wire.MethodSketchSetInferenceScope || !reflect.DeepEqual(result, request) {
		t.Fatalf("method %q result %+v: expected set method and %+v", transport.gotMethod, result, request)
	}
	var sent wire.SketchInferenceScopeView
	if err := json.Unmarshal(transport.gotReq, &sent); err != nil || !reflect.DeepEqual(sent, request) {
		t.Fatalf("request %+v, error %v: expected %+v", sent, err, request)
	}
	if _, err := client.Sketch().InferenceScope(2); err != nil {
		t.Fatalf("InferenceScope: %v", err)
	}
	var address wire.SketchInferenceScopeArgs
	if err := json.Unmarshal(transport.gotReq, &address); err != nil || transport.gotMethod != wire.MethodSketchGetInferenceScope || address.SketchIndex != 2 {
		t.Fatalf("method %q request %+v error %v: expected get for sketch 2", transport.gotMethod, address, err)
	}
}
