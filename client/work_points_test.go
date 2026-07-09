// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestWorkPointsAtMarshalsPosition(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":1,"ref":"point/0","name":"Work Point1"}`)}
	c := New(ft)

	got, err := c.WorkPoints().At(1, 2, 3)
	if err != nil {
		t.Fatalf("At: %v", err)
	}
	if ft.gotMethod != wire.MethodWorkPointsCreate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodWorkPointsCreate)
	}
	var sent wire.CreateWorkPointArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if len(sent.At) != 3 || sent.At[0] != 1 || sent.At[1] != 2 || sent.At[2] != 3 {
		t.Errorf("sent At = %v, want [1,2,3]", sent.At)
	}
	if got.Ref != "point/0" || got.Index != 1 {
		t.Errorf("decoded = %+v, want index 1 / point/0", got)
	}
}

func TestWorkPointsPlaneAxisIntersectionMarshalsKindAndRefs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":2,"ref":"point/1","name":"Work Point2","healthy":true}`)}
	c := New(ft)

	got, err := c.WorkPoints().PlaneAxisIntersection("origin/plane/xy", "axis/0")
	if err != nil {
		t.Fatalf("PlaneAxisIntersection: %v", err)
	}
	var sent wire.CreateWorkPointArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != "plane-axis-intersection" {
		t.Errorf("sent Kind = %q, want plane-axis-intersection", sent.Kind)
	}
	if len(sent.Refs) != 2 || sent.Refs[0] != "origin/plane/xy" || sent.Refs[1] != "axis/0" {
		t.Errorf("sent Refs = %v, want [origin/plane/xy axis/0]", sent.Refs)
	}
	if len(sent.At) != 0 {
		t.Errorf("intersection request should not send at, got %v", sent.At)
	}
	if got.Ref != "point/1" || !got.Healthy {
		t.Errorf("decoded = %+v, want point/1 healthy", got)
	}
}
