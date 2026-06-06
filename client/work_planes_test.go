// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati/api/types"
	"oblikovati/api/wire"
)

func TestWorkPlanesOffsetMarshalsKindRefsAndOffset(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":3,"ref":"plane/3","name":"Work Plane1","healthy":true}`)}
	c := New(ft)

	got, err := c.WorkPlanes().Offset(types.WorkRefXYPlane, "10 mm")
	if err != nil {
		t.Fatalf("Offset: %v", err)
	}
	if ft.gotMethod != wire.MethodWorkPlanesCreate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodWorkPlanesCreate)
	}
	var sent wire.CreateWorkPlaneArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != string(types.WorkPlaneOffset) ||
		len(sent.Refs) != 1 || sent.Refs[0] != types.WorkRefXYPlane || sent.Offset != "10 mm" {
		t.Errorf("sent = %+v, want offset of XY by 10 mm", sent)
	}
	if got.Index != 3 || got.Ref != "plane/3" || !got.Healthy {
		t.Errorf("decoded = %+v, want index 3 / plane/3 / healthy", got)
	}
}

func TestWorkPlanesTangentHelperSetsKindAndRefs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":4,"ref":"plane/4","name":"Work Plane2","healthy":false}`)}
	c := New(ft)

	if _, err := c.WorkPlanes().PointAndTangent("point/1", "face/abc"); err != nil {
		t.Fatalf("PointAndTangent: %v", err)
	}
	var sent wire.CreateWorkPlaneArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != string(types.WorkPlanePointAndTangent) ||
		len(sent.Refs) != 2 || sent.Refs[0] != "point/1" || sent.Refs[1] != "face/abc" {
		t.Errorf("sent = %+v, want point-tangent with [point/1, face/abc]", sent)
	}
}

func TestWorkPlanesListSendsNilBodyAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"planes":[{"index":0,"name":"XY Plane","ref":"origin/plane/xy","isOrigin":true,"healthy":true}]}`)}
	c := New(ft)

	res, err := c.WorkPlanes().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodWorkPlanesList || ft.gotReq != nil {
		t.Errorf("List sent method=%q body=%q, want %q / nil", ft.gotMethod, ft.gotReq, wire.MethodWorkPlanesList)
	}
	if len(res.Planes) != 1 || res.Planes[0].Ref != types.WorkRefXYPlane || !res.Planes[0].IsOrigin {
		t.Errorf("decoded = %+v, want the XY origin plane", res.Planes)
	}
}
