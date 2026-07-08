// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
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

func TestWorkPlanesOffsetHiddenSetsVisibleFalse(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"index":5,"ref":"plane/5","name":"Work Plane3","healthy":true}`)}
	c := New(ft)

	if _, err := c.WorkPlanes().OffsetHidden(types.WorkRefXYPlane, "-8 mm"); err != nil {
		t.Fatalf("OffsetHidden: %v", err)
	}
	var sent wire.CreateWorkPlaneArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Kind != string(types.WorkPlaneOffset) || sent.Offset != "-8 mm" {
		t.Errorf("sent = %+v, want an XY offset by -8 mm", sent)
	}
	if sent.Visible == nil || *sent.Visible {
		t.Errorf("Visible = %v, want a pointer to false (a hidden construction datum)", sent.Visible)
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

func TestWorkPlanesSetScalarMarshalsRedefine(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"plane":{"index":3,"name":"Work Plane1","healthy":true}}`)}
	c := New(ft)

	got, err := c.WorkPlanes().SetScalar(3, 0, "50 mm")
	if err != nil {
		t.Fatalf("SetScalar: %v", err)
	}
	if ft.gotMethod != wire.MethodWorkPlanesRedefine {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodWorkPlanesRedefine)
	}
	var sent wire.RedefineWorkPlaneArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Index != 3 || len(sent.Scalars) != 1 || sent.Scalars[0].Index != 0 || sent.Scalars[0].Value != "50 mm" {
		t.Errorf("sent = %+v, want index 3 scalar 0 = 50 mm", sent)
	}
	if got.Plane.Index != 3 || !got.Plane.Healthy {
		t.Errorf("decoded = %+v, want plane index 3 healthy", got.Plane)
	}
}

func TestWorkPlanesRepickMarshalsRedefine(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"plane":{"index":3,"name":"Work Plane1","healthy":true}}`)}
	c := New(ft)

	if _, err := c.WorkPlanes().Repick(3, 0, types.WorkRefXZPlane); err != nil {
		t.Fatalf("Repick: %v", err)
	}
	var sent wire.RedefineWorkPlaneArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Index != 3 || len(sent.Repick) != 1 || sent.Repick[0].Slot != 0 || sent.Repick[0].Ref != types.WorkRefXZPlane {
		t.Errorf("sent = %+v, want index 3 slot 0 -> XZ plane", sent)
	}
}
