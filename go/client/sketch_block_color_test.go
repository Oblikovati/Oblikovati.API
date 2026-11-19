// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestBlockColorsCarryOverrideAndExplicitDefault(t *testing.T) {
	cyan := types.NewColor(0, 255, 255)
	ft := &fakeTransport{reply: []byte(`{"name":"Rib","color":{"r":0,"g":255,"b":255,"opacity":1,"source":79105}}`)}
	client := New(ft).Sketch()
	definition, err := client.SetBlockDefinitionColor(wire.SetBlockDefinitionColorArgs{Name: "Rib", Color: &cyan})
	if err != nil || definition.Color == nil || *definition.Color != cyan {
		t.Fatalf("definition color = %+v, %v: expected cyan", definition, err)
	}
	if ft.gotMethod != wire.MethodSketchSetBlockDefinitionColor {
		t.Fatalf("method = %q: expected definition color method", ft.gotMethod)
	}
	var input wire.SetBlockDefinitionColorArgs
	if err := json.Unmarshal(ft.gotReq, &input); err != nil || input.Color == nil || *input.Color != cyan {
		t.Fatalf("definition input = %+v, %v: expected cyan", input, err)
	}
	ft.reply = []byte(`{"id":42,"definition":"Rib"}`)
	instance, err := client.SetBlockInstanceColor(wire.SetBlockInstanceColorArgs{SketchIndex: 2, InstanceID: 42})
	if err != nil || instance.Color != nil {
		t.Fatalf("instance color = %+v, %v: expected default", instance, err)
	}
	if ft.gotMethod != wire.MethodSketchSetBlockInstanceColor || string(ft.gotReq) != `{"sketchIndex":2,"instanceId":42,"color":null}` {
		t.Fatalf("instance request = %q %s: expected explicit null color", ft.gotMethod, ft.gotReq)
	}
}
