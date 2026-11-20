// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"reflect"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

type partViewCall struct {
	name, method, request, reply string
	invoke                       func(PartDesignViews) error
}

func partViewIdentityCalls() []partViewCall {
	return []partViewCall{
		{"list", wire.MethodPartDesignViewsList, `{}`, `{"views":[]}`, func(p PartDesignViews) error { _, e := p.List(); return e }},
		{"copy", wire.MethodPartDesignViewsCopy, `{"sourceId":1,"name":"Review"}`, `{"view":{"id":2}}`, func(p PartDesignViews) error {
			_, e := p.Copy(wire.CopyPartDesignViewArgs{SourceID: 1, Name: "Review"})
			return e
		}},
		{"rename", wire.MethodPartDesignViewsRename, `{"id":2,"name":"Front"}`, `{"view":{"id":2}}`, func(p PartDesignViews) error {
			_, e := p.Rename(wire.RenamePartDesignViewArgs{ID: 2, Name: "Front"})
			return e
		}},
		{"activate", wire.MethodPartDesignViewsActivate, `{"id":2}`, `{"view":{"id":2}}`, func(p PartDesignViews) error { _, e := p.Activate(2); return e }},
		{"delete", wire.MethodPartDesignViewsDelete, `{"id":2}`, `{"views":[]}`, func(p PartDesignViews) error { _, e := p.Delete(2); return e }},
	}
}

func partViewContentCalls() []partViewCall {
	locked := true
	scale := types.PartAnnotationScale{Mode: types.PartAnnotationScaleCustom, Numerator: 3, Denominator: 2}
	return []partViewCall{
		{"options", wire.MethodPartDesignViewsSetOptions, `{"id":2,"locked":true,"annotationScale":{"mode":"custom","numerator":3,"denominator":2}}`, `{"view":{"id":2}}`, func(p PartDesignViews) error {
			_, e := p.SetOptions(wire.SetPartDesignViewOptionsArgs{ID: 2, Locked: &locked, AnnotationScale: &scale})
			return e
		}},
		{"camera", wire.MethodPartDesignViewsGetCamera, `{"id":2}`, `{}`, func(p PartDesignViews) error { _, e := p.Camera(2); return e }},
		{"capture", wire.MethodPartDesignViewsCaptureCamera, `{"id":2,"camera":{"eye":[0,0,0],"target":[0,0,0],"up":[0,0,0],"scale":0,"near":0,"far":0,"projection":0}}`, `{"view":{"id":2}}`, func(p PartDesignViews) error {
			_, e := p.CaptureCamera(wire.CapturePartDesignViewCameraArgs{ID: 2})
			return e
		}},
		{"visibility", wire.MethodPartDesignViewsSetBodyVisibility, `{"id":2,"bodyKeys":["aabb"],"visible":false}`, `{"view":{"id":2}}`, func(p PartDesignViews) error {
			_, e := p.SetBodyVisibility(wire.SetPartDesignViewBodyVisibilityArgs{ID: 2, BodyKeys: []string{"aabb"}})
			return e
		}},
	}
}

func TestPartDesignViewMethodsUseTypedWireRequests(t *testing.T) {
	cases := append(partViewIdentityCalls(), partViewContentCalls()...)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkPartViewCall(t, tc) })
	}
}

func TestPartViewCameraRoundTripKeepsScaleAndClipDistances(t *testing.T) {
	want := types.PartDesignViewCamera{
		Eye: types.Point{Z: 8}, Target: types.Point{}, Up: types.Vector{Y: 1},
		Scale: 12, Near: 0.5, Far: 250, Projection: types.OrthographicProjection,
	}
	reply, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	transport := &fakeTransport{reply: reply}
	got, err := New(transport).PartDesignViews().Camera(2)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("camera %+v: expected %+v", got, want)
	}
}

func TestExistingCameraViewEncodingIsUnchanged(t *testing.T) {
	encoded, err := json.Marshal(wire.CameraView{})
	if err != nil {
		t.Fatal(err)
	}
	const want = `{"eye":[0,0,0],"target":[0,0,0],"up":[0,0,0],"fov":0}`
	if string(encoded) != want {
		t.Fatalf("camera encoding %s: expected %s", encoded, want)
	}
}

func checkPartViewCall(t *testing.T, tc partViewCall) {
	t.Helper()
	transport := &fakeTransport{reply: []byte(tc.reply)}
	if err := tc.invoke(New(transport).PartDesignViews()); err != nil {
		t.Fatal(err)
	}
	if transport.gotMethod != tc.method {
		t.Fatalf("method %q: expected %q", transport.gotMethod, tc.method)
	}
	var got, want map[string]json.RawMessage
	if err := json.Unmarshal(transport.gotReq, &got); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(tc.request), &want); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request %s: expected %s", transport.gotReq, tc.request)
	}
}
