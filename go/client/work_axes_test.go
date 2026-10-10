// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestWorkAxesLineMarshalsKindOriginAndDirection(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":3,"ref":"axis/3","name":"Work Axis1","healthy":true}`)}
	c := New(ft)

	got, err := c.WorkAxes().Line([]float64{0, 0, 0}, []float64{1, 0, 0})
	if err != nil {
		t.Fatalf("Line: %v", err)
	}
	if ft.gotMethod != wire.MethodWorkAxesCreate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodWorkAxesCreate)
	}
	var sent wire.CreateWorkAxisArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != string(types.WorkAxisLine) ||
		len(sent.Origin) != 3 || len(sent.Direction) != 3 || sent.Direction[0] != 1 {
		t.Errorf("sent = %+v, want a line axis at origin along +X", sent)
	}
	if sent.Refs != nil {
		t.Errorf("line axis leaked refs: %+v", sent.Refs)
	}
	if got.Index != 3 || got.Ref != "axis/3" || !got.Healthy {
		t.Errorf("decoded = %+v, want index 3 / axis/3 / healthy", got)
	}
}

func TestWorkAxesTwoPointsSetsKindAndRefs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":4,"ref":"axis/4","name":"Work Axis2","healthy":false,"reason":"points coincident"}`)}
	c := New(ft)

	got, err := c.WorkAxes().TwoPoints("point/0", types.WorkRefCenter)
	if err != nil {
		t.Fatalf("TwoPoints: %v", err)
	}
	var sent wire.CreateWorkAxisArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != string(types.WorkAxisTwoPoints) ||
		len(sent.Refs) != 2 || sent.Refs[0] != "point/0" || sent.Refs[1] != types.WorkRefCenter {
		t.Errorf("sent = %+v, want two-points with [point/0, origin center]", sent)
	}
	if got.Healthy || got.Reason != "points coincident" {
		t.Errorf("decoded = %+v, want unhealthy with a reason", got)
	}
}

func TestWorkAxesPlaneIntersectionSetsKindAndRefs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":5,"ref":"axis/5","name":"Work Axis3","healthy":true}`)}
	c := New(ft)

	if _, err := c.WorkAxes().PlaneIntersection(types.WorkRefXYPlane, types.WorkRefXZPlane); err != nil {
		t.Fatalf("PlaneIntersection: %v", err)
	}
	if ft.gotMethod != wire.MethodWorkAxesCreate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodWorkAxesCreate)
	}
	var sent wire.CreateWorkAxisArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != string(types.WorkAxisPlaneIntersection) ||
		len(sent.Refs) != 2 || sent.Refs[0] != types.WorkRefXYPlane || sent.Refs[1] != types.WorkRefXZPlane {
		t.Errorf("sent = %+v, want plane-intersection with [XY, XZ]", sent)
	}
}

func TestWorkAxesListSendsNilBodyAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"axes":[{"index":0,"name":"X Axis","ref":"origin/axis/x","kind":"line","origin":[0,0,0],"direction":[1,0,0],"isOrigin":true,"healthy":true}]}`)}
	c := New(ft)

	res, err := c.WorkAxes().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodWorkAxesList || ft.gotReq != nil {
		t.Errorf("List sent method=%q body=%q, want %q / nil", ft.gotMethod, ft.gotReq, wire.MethodWorkAxesList)
	}
	if len(res.Axes) != 1 || res.Axes[0].Ref != types.WorkRefXAxis || !res.Axes[0].IsOrigin {
		t.Errorf("decoded = %+v, want the X origin axis", res.Axes)
	}
}
