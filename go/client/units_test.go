// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"strings"
	"testing"

	"oblikovati.org/api/wire"
)

// TestDocumentUnitsRoundTrip drives GetUnits/SetUnits: the right wire constants
// go out, only changed fields ride the wire on a partial update, and the units
// DTO decodes back.
func TestDocumentUnitsRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{
		"lengthUnit":"mm","angleUnit":"deg","massUnit":"kg","timeUnit":"s",
		"lengthDisplayPrecision":3,"angleDisplayPrecision":2,
		"lengthDisplayFormat":"decimal"}`)}
	c := New(ft)

	got, err := c.Documents().GetUnits()
	if err != nil || ft.gotMethod != wire.MethodDocumentsGetUnits {
		t.Fatalf("GetUnits = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodDocumentsGetUnits)
	}
	if got.LengthUnit != "mm" || got.LengthDisplayPrecision != 3 || got.LengthDisplayFormat != "decimal" {
		t.Errorf("units = %+v, want the fake's values decoded", got)
	}

	in := "in"
	if _, err := c.Documents().SetUnits(wire.SetDocumentUnitsArgs{LengthUnit: &in}); err != nil {
		t.Fatalf("SetUnits: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentsSetUnits {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentsSetUnits)
	}
	var sent map[string]json.RawMessage
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("sent args not JSON: %v", err)
	}
	if len(sent) != 1 || string(sent["lengthUnit"]) != `"in"` {
		t.Errorf("sent args = %s, want only the changed lengthUnit on the wire", ft.gotReq)
	}
}

// TestUnitsServiceMethods covers the conversion / formatting / expression
// methods: each sends its own wire constant with the expected request and
// decodes its reply.
func TestUnitsServiceMethods(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"value":2.54}`)}
	c := New(ft)

	conv, err := c.Units().Convert(wire.ConvertUnitsArgs{Value: 1, From: "in", To: "cm"})
	if err != nil || ft.gotMethod != wire.MethodUnitsConvert {
		t.Fatalf("Convert = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodUnitsConvert)
	}
	if conv.Value != 2.54 || !strings.Contains(string(ft.gotReq), `"from":"in"`) {
		t.Errorf("convert = %+v sent %s", conv, ft.gotReq)
	}

	ft.reply = []byte(`{"value":"1-1/4 in"}`)
	str, err := c.Units().GetStringFromValue(3.175, "length")
	if err != nil || ft.gotMethod != wire.MethodUnitsGetStringFromValue || str.Value != "1-1/4 in" {
		t.Fatalf("GetStringFromValue = (%q, %+v, %v)", ft.gotMethod, str, err)
	}

	ft.reply = []byte(`{"value":2.5}`)
	val, err := c.Units().GetValueFromExpression("25 mm", "length")
	if err != nil || ft.gotMethod != wire.MethodUnitsGetValueFromExpression || val.Value != 2.5 {
		t.Fatalf("GetValueFromExpression = (%q, %+v, %v)", ft.gotMethod, val, err)
	}

	ft.reply = []byte(`{"unitsType":"length"}`)
	typ, err := c.Units().GetTypeFromString("mm")
	if err != nil || ft.gotMethod != wire.MethodUnitsGetTypeFromString || typ.UnitsType != "length" {
		t.Fatalf("GetTypeFromString = (%q, %+v, %v)", ft.gotMethod, typ, err)
	}

	ft.reply = []byte(`{"valid":true}`)
	v, err := c.Units().IsExpressionValid("3 mm + 1 cm", "length")
	if err != nil || ft.gotMethod != wire.MethodUnitsIsExpressionValid || !v.Valid {
		t.Fatalf("IsExpressionValid = (%q, %+v, %v)", ft.gotMethod, v, err)
	}

	ft.reply = []byte(`{"names":["width","height"]}`)
	dp, err := c.Units().GetDrivingParameters("width + height")
	if err != nil || ft.gotMethod != wire.MethodUnitsGetDrivingParameters || len(dp.Names) != 2 {
		t.Fatalf("GetDrivingParameters = (%q, %+v, %v)", ft.gotMethod, dp, err)
	}
}
