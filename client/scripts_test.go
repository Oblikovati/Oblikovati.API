// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestScriptsRunMarshalsSourceAndDecodesResult(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"output":"hi\n","error":"","durationMs":12,"ops":0}`)}
	c := New(ft)

	res, err := c.Scripts().Run(`print("hi")`, 5000)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if ft.gotMethod != wire.MethodScriptRun {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodScriptRun)
	}
	var sent wire.ScriptRunArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Source != `print("hi")` || sent.WallMs != 5000 {
		t.Errorf("sent = %+v, want source+wallMs preserved", sent)
	}
	if res.Output != "hi\n" || res.DurationMs != 12 {
		t.Errorf("decoded = %+v, want output+duration", res)
	}
}

// TestScriptsRunSurfacesScriptError: a script failure rides in the result's Error field,
// not as a transport error (the call itself succeeds).
func TestScriptsRunSurfacesScriptError(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"output":"","error":"runtime error: boom","durationMs":3}`)}
	c := New(ft)

	res, err := c.Scripts().Run(`error("boom")`, 0)
	if err != nil {
		t.Fatalf("Run should not fail at transport level on a script error: %v", err)
	}
	if res.Error == "" {
		t.Error("a script error should be reported in the result")
	}
}
