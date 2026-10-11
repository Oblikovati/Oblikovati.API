// SPDX-License-Identifier: Apache-2.0
package ellipse

import (
	"encoding/json"
	"fmt"
	"oblikovati.org/api/wire"
	"testing"
)

type minorRadiusHost struct{}

func (minorRadiusHost) Call(method string, payload []byte) ([]byte, error) {
	var request wire.AddDimensionArgs
	if err := json.Unmarshal(payload, &request); err != nil {
		return nil, err
	}
	if method != wire.MethodSketchAddDimension || request.Kind != "ellipseMinorRadius" || request.SketchIndex != 2 || len(request.Entities) != 1 || request.Entities[0] != 42 || request.Expression != "5 mm" {
		return nil, fmt.Errorf("call %q(%+v): expected minor radius of ellipse 42 in sketch 2 with 5 mm", method, request)
	}
	return []byte(`{"kind":"ellipseMinorRadius","value":0.5,"dof":0}`), nil
}
func TestAddMinorRadiusExample(t *testing.T) {
	result, err := AddMinorRadius(minorRadiusHost{}, 2, 42, "5 mm")
	if err != nil {
		t.Fatal(err)
	}
	if result.Value != 0.5 || result.DOF != 0 {
		t.Fatalf("result %+v: expected 0.5 cm and zero DOF", result)
	}
}
