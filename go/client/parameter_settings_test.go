// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"strings"
	"testing"

	"oblikovati.org/api/wire"
)

// TestParameterSettingsRoundTrip drives GetSettings/SetSettings: the right
// wire constants go out, nil update fields stay off the wire, and the
// settings DTO decodes back.
func TestParameterSettingsRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{
		"linearStandardTolerance":"0.1 mm","useStandardTolerances":true,
		"linearDimensionPrecision":3,"angularDimensionPrecision":2,
		"dimensionDisplayType":"expression"}`)}
	c := New(ft)

	got, err := c.Parameters().GetSettings()
	if err != nil || ft.gotMethod != wire.MethodParametersGetSettings {
		t.Fatalf("GetSettings = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodParametersGetSettings)
	}
	if got.LinearStandardTolerance != "0.1 mm" || !got.UseStandardTolerances || got.DimensionDisplayType != "expression" {
		t.Errorf("settings = %+v, want the fake's values decoded", got)
	}

	prec := 4
	if _, err := c.Parameters().SetSettings(wire.ParameterSettingsUpdateArgs{LinearDimensionPrecision: &prec}); err != nil {
		t.Fatalf("SetSettings: %v", err)
	}
	var sent map[string]json.RawMessage
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("sent args not JSON: %v", err)
	}
	if len(sent) != 1 || string(sent["linearDimensionPrecision"]) != "4" {
		t.Errorf("sent args = %s, want only the changed precision on the wire", ft.gotReq)
	}
}

// TestParameterSweepAndExchange covers the sweep and the XML export/import
// methods.
func TestParameterSweepAndExchange(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"affected":3}`)}
	c := New(ft)

	sweep, err := c.Parameters().SetAllModelValueType("upper")
	if err != nil || ft.gotMethod != wire.MethodParametersSetAllModelValueType {
		t.Fatalf("SetAllModelValueType = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodParametersSetAllModelValueType)
	}
	if sweep.Affected != 3 || !strings.Contains(string(ft.gotReq), `"modelValueType":"upper"`) {
		t.Errorf("sweep = %+v sent %s, want 3 affected by an upper sweep", sweep, ft.gotReq)
	}

	ft.reply = []byte(`{"xml":"<parameters/>"}`)
	exp, err := c.Parameters().Export()
	if err != nil || ft.gotMethod != wire.MethodParametersExport || exp.XML != "<parameters/>" {
		t.Errorf("Export = (%+v, %q, %v), want the XML back", exp, ft.gotMethod, err)
	}

	ft.reply = []byte(`{"added":2,"updated":1}`)
	imp, err := c.Parameters().Import("<parameters/>")
	if err != nil || ft.gotMethod != wire.MethodParametersImport {
		t.Fatalf("Import = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodParametersImport)
	}
	if imp.Added != 2 || imp.Updated != 1 {
		t.Errorf("import result = %+v, want 2 added / 1 updated", imp)
	}
}
