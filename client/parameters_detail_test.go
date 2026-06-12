// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

// TestParametersGetDetailDecodesFullShape drives GetDetail end-to-end: the
// embedded ParameterInfo fields decode from the same flat JSON object as the
// detail-only fields.
func TestParametersGetDetailDecodesFullShape(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{
		"name":"od","kind":"user","expression":"4 cm","value":"40 mm",
		"units":"mm","comment":"outer diameter","isKey":true,"visible":true,
		"inUse":true,"precision":3,"displayFormat":"decimal",
		"modelValue":4.05,"modelValueType":"upper",
		"tolerance":{"type":"deviation","upper":0.05,"lower":-0.02},
		"expressionList":{"expressions":["4 cm","5 cm"],"allowCustomValues":true},
		"drivenBy":["base"],"dependents":["wall"]}`)}
	c := New(ft)

	got, err := c.Parameters().GetDetail("od")
	if err != nil {
		t.Fatalf("GetDetail: %v", err)
	}
	if ft.gotMethod != wire.MethodParametersGetDetail {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodParametersGetDetail)
	}
	if got.Name != "od" || got.Expression != "4 cm" {
		t.Errorf("embedded info = %+v, want od / 4 cm", got.ParameterInfo)
	}
	if got.Tolerance == nil || got.Tolerance.Type != "deviation" || got.Tolerance.Upper != 0.05 {
		t.Errorf("tolerance = %+v, want deviation +0.05", got.Tolerance)
	}
	if got.ModelValue != 4.05 || got.ModelValueType != "upper" {
		t.Errorf("model value = (%v, %s), want (4.05, upper)", got.ModelValue, got.ModelValueType)
	}
	if got.ExpressionList == nil || len(got.ExpressionList.Expressions) != 2 || !got.ExpressionList.AllowCustomValues {
		t.Errorf("expression list = %+v, want 2 choices + custom", got.ExpressionList)
	}
	if len(got.DrivenBy) != 1 || got.DrivenBy[0] != "base" || len(got.Dependents) != 1 {
		t.Errorf("graph neighborhood = (%v, %v), want ([base], [wall])", got.DrivenBy, got.Dependents)
	}
}

// TestParametersUpdateSendsOnlyChangedFields asserts nil optional fields stay
// off the wire, so the host can distinguish "unchanged" from "set to zero".
func TestParametersUpdateSendsOnlyChangedFields(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"name":"od","kind":"user","visible":true,"precision":2,"displayFormat":"decimal","modelValue":4,"modelValueType":"nominal"}`)}
	c := New(ft)

	comment := "hub bore"
	if _, err := c.Parameters().Update(wire.ParameterUpdateArgs{Name: "od", Comment: &comment}); err != nil {
		t.Fatalf("Update: %v", err)
	}
	var sent map[string]json.RawMessage
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if len(sent) != 2 {
		t.Errorf("request fields = %v, want exactly {name, comment}", sent)
	}
}

// TestParametersToleranceExpressionListAndGraphCalls pins the wire methods of
// the remaining group members.
func TestParametersToleranceExpressionListAndGraphCalls(t *testing.T) {
	detail := []byte(`{"name":"od","kind":"user","visible":true,"precision":2,"displayFormat":"decimal","modelValue":4,"modelValueType":"nominal"}`)
	ft := &fakeTransport{reply: detail}
	c := New(ft)

	if _, err := c.Parameters().SetTolerance(wire.ParameterToleranceArgs{Name: "od", Mode: "symmetric", Upper: "0.1 mm"}); err != nil {
		t.Fatalf("SetTolerance: %v", err)
	}
	if ft.gotMethod != wire.MethodParametersSetTolerance {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodParametersSetTolerance)
	}
	if _, err := c.Parameters().SetExpressionList(wire.ParameterExpressionListArgs{Name: "od", Expressions: []string{"4 cm"}}); err != nil {
		t.Fatalf("SetExpressionList: %v", err)
	}
	if ft.gotMethod != wire.MethodParametersSetExpressionList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodParametersSetExpressionList)
	}

	ft.reply = []byte(`{"names":["base"]}`)
	deps, err := c.Parameters().DrivenBy("od")
	if err != nil || len(deps.Names) != 1 || deps.Names[0] != "base" {
		t.Fatalf("DrivenBy = (%+v, %v), want [base]", deps, err)
	}
	if ft.gotMethod != wire.MethodParametersDrivenBy {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodParametersDrivenBy)
	}
	if _, err := c.Parameters().Dependents("od"); err != nil {
		t.Fatalf("Dependents: %v", err)
	}
	if ft.gotMethod != wire.MethodParametersDependents {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodParametersDependents)
	}

	ft.reply = []byte(`{}`)
	if err := c.Parameters().Delete("od"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if ft.gotMethod != wire.MethodParametersDelete {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodParametersDelete)
	}
}
