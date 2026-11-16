// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"strings"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestSetPartRejectsUnknownStartupPlaneBeforeTransport(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"ok":true}`)}
	invalid := types.StartupSketchPlane("diagonal")
	_, err := New(transport).Options().SetPart(wire.PartOptionsView{StartupSketchPlane: &invalid})
	if err == nil || !strings.Contains(err.Error(), "diagonal") || transport.gotMethod != "" {
		t.Fatalf("error %v, method %q: expected rejection before transport", err, transport.gotMethod)
	}
}

func TestSetPartPreservesOmissionAndFalse(t *testing.T) {
	for _, facing := range []*bool{nil, new(bool)} {
		transport := &fakeTransport{reply: []byte(`{"ok":true}`)}
		_, err := New(transport).Options().SetPart(wire.PartOptionsView{LookAtSketchPlane: facing})
		if err != nil {
			t.Fatal(err)
		}
		var request wire.OptionGroupView
		if err := json.Unmarshal(transport.gotReq, &request); err != nil || request.Part == nil {
			t.Fatalf("request %s, error %v: expected part", transport.gotReq, err)
		}
		if (request.Part.LookAtSketchPlane == nil) != (facing == nil) {
			t.Fatalf("request %s: expected facing %v", transport.gotReq, facing)
		}
		if facing != nil && *request.Part.LookAtSketchPlane {
			t.Fatalf("request %s: expected false", transport.gotReq)
		}
	}
}

func TestSetPartAcceptsEveryStartupPlane(t *testing.T) {
	for _, plane := range []types.StartupSketchPlane{types.StartupSketchNone, types.StartupSketchXY, types.StartupSketchYZ, types.StartupSketchXZ} {
		transport := &fakeTransport{reply: []byte(`{"ok":true}`)}
		_, err := New(transport).Options().SetPart(wire.PartOptionsView{StartupSketchPlane: &plane})
		if err != nil || transport.gotMethod != wire.MethodOptionsSetGroup {
			t.Fatalf("plane %q: method %q, error %v", plane, transport.gotMethod, err)
		}
	}
}
