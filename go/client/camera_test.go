// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestViewSetCameraMarshalsFrameAndDecodesReply(t *testing.T) {
	reply := `{"eye":[0,0,10],"target":[0,0,0],"up":[0,1,0],"fov":0.7853981633974483}`
	ft := &fakeTransport{reply: []byte(reply)}
	c := New(ft)

	got, err := c.View().SetCamera(wire.SetCameraArgs{
		Eye: types.NewPoint(5, 6, 7), Target: types.NewPoint(1, 2, 3), Up: types.NewVector(0, 1, 0), FOV: 0.9,
	})
	if err != nil {
		t.Fatalf("SetCamera: %v", err)
	}
	if ft.gotMethod != wire.MethodViewSetCamera {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodViewSetCamera)
	}
	var sent wire.SetCameraArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Eye != types.NewPoint(5, 6, 7) || sent.Target != types.NewPoint(1, 2, 3) || sent.FOV != 0.9 {
		t.Errorf("sent = %+v, want eye=[5 6 7] target=[1 2 3] fov=0.9", sent)
	}
	if got.Eye != types.NewPoint(0, 0, 10) {
		t.Errorf("decoded eye = %v, want [0 0 10]", got.Eye)
	}
}

func TestViewCameraSendsNilBodyAndDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"eye":[1,1,1],"target":[0,0,0],"up":[0,1,0],"fov":1.0}`)}
	c := New(ft)

	got, err := c.View().Camera()
	if err != nil {
		t.Fatalf("Camera: %v", err)
	}
	if ft.gotMethod != wire.MethodViewGetCamera {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodViewGetCamera)
	}
	if ft.gotReq != nil {
		t.Errorf("no-arg getter should send nil body, got %q", ft.gotReq)
	}
	if got.FOV != 1.0 || got.Eye != types.NewPoint(1, 1, 1) {
		t.Errorf("decoded = %+v, want eye=[1 1 1] fov=1", got)
	}
}
