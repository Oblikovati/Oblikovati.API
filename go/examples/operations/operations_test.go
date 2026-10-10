// SPDX-License-Identifier: Apache-2.0

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"oblikovati.org/api/wire"
	"testing"
)

type recordingHost struct {
	method   string
	request  string
	response string
}

func (host recordingHost) Call(method string, request []byte) ([]byte, error) {
	var compact bytes.Buffer
	if err := json.Compact(&compact, request); err != nil {
		return nil, err
	}
	if method != host.method || compact.String() != host.request {
		return nil, fmt.Errorf("call %s(%s): expected %s(%s)", method, compact.String(), host.method, host.request)
	}
	return []byte(host.response), nil
}

func TestCreatePart(t *testing.T) {
	host := recordingHost{wire.MethodDocumentsCreate, `{"type":"part","name":"Bracket"}`, `{"id":7,"name":"Bracket","type":"part"}`}
	part, err := CreatePart(host, "Bracket")
	if err != nil || part.ID != 7 {
		t.Fatalf("part=%+v error=%v", part, err)
	}
}

func TestExtrudeProfile(t *testing.T) {
	host := recordingHost{wire.MethodFeaturesAdd, `{"kind":"extrude","args":{"sketchIndex":2,"profileIndex":0,"distance":"50 mm","operation":"new"}}`, `{}`}
	if _, err := ExtrudeProfile(host, 2, "50 mm"); err != nil {
		t.Fatal(err)
	}
}
