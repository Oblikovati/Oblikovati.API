// SPDX-License-Identifier: Apache-2.0

package operations

import (
	"oblikovati.org/api/wire"
	"testing"
)

func TestListAllAxes(t *testing.T) {
	host := recordingHost{wire.MethodWorkAxesList, `{"includeConstruction":true}`, `{"axes":[{"index":4,"construction":true}]}`}
	result, err := ListAllAxes(host)
	if err != nil || len(result.Axes) != 1 || !result.Axes[0].Construction {
		t.Fatalf("result=%+v error=%v: expected one construction axis", result, err)
	}
}

func TestListAllPoints(t *testing.T) {
	host := recordingHost{wire.MethodWorkPointsList, `{"includeConstruction":true}`, `{"points":[{"index":2,"construction":true}]}`}
	result, err := ListAllPoints(host)
	if err != nil || len(result.Points) != 1 || !result.Points[0].Construction {
		t.Fatalf("result=%+v error=%v: expected one construction point", result, err)
	}
}

func TestListAllPlanes(t *testing.T) {
	host := recordingHost{wire.MethodWorkPlanesList, `{"includeConstruction":true}`, `{"planes":[{"index":4,"construction":true}]}`}
	result, err := ListAllPlanes(host)
	if err != nil || len(result.Planes) != 1 || !result.Planes[0].Construction {
		t.Fatalf("result=%+v error=%v: expected one construction plane", result, err)
	}
}
