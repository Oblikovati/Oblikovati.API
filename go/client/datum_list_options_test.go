// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"errors"
	"oblikovati.org/api/wire"
	"reflect"
	"testing"
)

func TestWorkAxesListWithOptionsUsesCanonicalRequest(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"axes":[{"index":4,"ref":"datum/v1/9","construction":true}]}`)}
	args := wire.ListWorkAxesArgs{Occurrence: []string{"assembly", "part"}, IncludeConstruction: true}
	result, err := New(ft).WorkAxes().ListWithOptions(args)
	if err != nil || len(result.Axes) != 1 || result.Axes[0].Index != 4 || !result.Axes[0].Construction {
		t.Fatalf("result=%+v error=%v: expected construction axis at index 4", result, err)
	}
	assertDatumListRequest(t, ft, wire.MethodWorkAxesList, args)
}

func TestWorkPointsListWithOptionsUsesCanonicalRequest(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"points":[{"index":2,"ref":"datum/v1/8","visible":false,"construction":true}]}`)}
	args := wire.ListWorkPointsArgs{Occurrence: []string{"assembly", "part"}, IncludeConstruction: true}
	result, err := New(ft).WorkPoints().ListWithOptions(args)
	if err != nil || len(result.Points) != 1 || result.Points[0].Index != 2 || result.Points[0].Visible {
		t.Fatalf("result=%+v error=%v: expected hidden point at index 2", result, err)
	}
	assertDatumListRequest(t, ft, wire.MethodWorkPointsList, args)
}

func assertDatumListRequest[T comparableRequest](t *testing.T, ft *fakeTransport, method string, args T) {
	t.Helper()
	var sent T
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request=%q: expected canonical JSON: %v", ft.gotReq, err)
	}
	if ft.gotMethod != method || !reflect.DeepEqual(sent, args) {
		t.Fatalf("method=%q request=%+v: expected %q %+v", ft.gotMethod, sent, method, args)
	}
}

type comparableRequest interface {
	wire.ListWorkAxesArgs | wire.ListWorkPointsArgs | wire.ListWorkPlanesArgs
}

func TestDatumListWithOptionsPropagatesTransportFailure(t *testing.T) {
	failure := errors.New("local part scope required")
	ft := &fakeTransport{err: failure}
	c := New(ft)
	_, axesErr := c.WorkAxes().ListWithOptions(wire.ListWorkAxesArgs{})
	_, pointsErr := c.WorkPoints().ListWithOptions(wire.ListWorkPointsArgs{})
	_, planesErr := c.WorkPlanes().ListWithOptions(wire.ListWorkPlanesArgs{})
	if !errors.Is(axesErr, failure) || !errors.Is(pointsErr, failure) || !errors.Is(planesErr, failure) {
		t.Fatalf("errors=(%v,%v): expected transport failure %v", axesErr, pointsErr, failure)
	}
}

func TestWorkPlanesListWithOptionsUsesCanonicalRequest(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"planes":[{"index":4,"ref":"datum/v1/9","construction":true}]}`)}
	args := wire.ListWorkPlanesArgs{Occurrence: []string{"assembly", "part"}, IncludeConstruction: true}
	result, err := New(ft).WorkPlanes().ListWithOptions(args)
	if err != nil || len(result.Planes) != 1 || result.Planes[0].Index != 4 || !result.Planes[0].Construction {
		t.Fatalf("result=%+v error=%v: expected construction plane at index 4", result, err)
	}
	assertDatumListRequest(t, ft, wire.MethodWorkPlanesList, args)
}
