// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"reflect"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestSetModelStateTablePreservesOneMixedDraft(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"table":{},"createdRows":[{"draftKey":"new-1","memberId":7}]}`)}
	primary := wire.ModelStateTableRowRef{MemberID: 1}
	newRow := wire.ModelStateTableRowRef{DraftKey: "new-1"}
	property := types.ModelStatePropertyKey{Set: "Summary Information", Name: "Category"}
	value := types.StringVariant("")
	suppressed := true
	expression := "2 in"
	request := wire.SetModelStateTableArgs{
		ExpectedGeneration: 0,
		Rows:               []wire.ModelStateTableRow{{Ref: primary, Name: "Primary"}, {Ref: newRow, Name: "New", Source: &primary}},
		Columns:            []wire.ModelStateTableColumnRef{{Kind: types.ModelStateTablePropertyColumn, Property: &property}, {Kind: types.ModelStateTableSuppressionColumn, FeatureID: 4}, {Kind: types.ModelStateTableParameterColumn, ParameterID: 9}},
		Properties:         []wire.ModelStateTableProperty{{Property: property, Cells: []wire.ModelStateTablePropertyCell{{Row: primary}, {Row: newRow, Value: &value}}}},
		Suppressions:       []wire.ModelStateTableSuppression{{FeatureID: 4, Cells: []wire.ModelStateTableSuppressionCell{{Row: primary}, {Row: newRow, Suppressed: &suppressed}}}},
		Parameters:         []wire.ModelStateTableParameter{{ParameterID: 9, Base: "1 in", Cells: []wire.ModelStateTableParameterCell{{Row: primary}, {Row: newRow, Expression: &expression}}}},
	}
	result, err := New(transport).ModelStates().SetTable(request)
	if err != nil {
		t.Fatal(err)
	}
	if transport.gotMethod != wire.MethodModelStatesSetTable || len(result.CreatedRows) != 1 || result.CreatedRows[0].MemberID != 7 {
		t.Fatalf("method %q, result %+v: expected the created row mapping", transport.gotMethod, result)
	}
	var decoded wire.SetModelStateTableArgs
	if err := json.Unmarshal(transport.gotReq, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded, request) {
		t.Fatalf("request %+v: expected %+v", decoded, request)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(transport.gotReq, &fields); err != nil {
		t.Fatal(err)
	}
	if string(fields["expectedGeneration"]) != "0" {
		t.Fatalf("generation %s: expected explicit zero", fields["expectedGeneration"])
	}
}

func TestGetModelStateTableDecodesMixedMemberCells(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"generation":4,"activeMemberId":1,"rows":[{"ref":{"memberId":1},"name":"Primary"}],"columns":[],"properties":[{"property":{"set":"Summary Information","name":"Category"},"base":null,"cells":[{"row":{"memberId":1},"value":{"type":"string","value":""}}]}],"suppressions":[{"featureId":4,"base":false,"cells":[{"row":{"memberId":1},"suppressed":true}]}],"parameters":[{"parameterId":9,"base":"1 in","cells":[{"row":{"memberId":1},"expression":null}]}],"availableProperties":[],"availableFeatures":[],"availableParameters":[],"evaluations":[]}`)}
	view, err := New(transport).ModelStates().GetTable()
	if err != nil {
		t.Fatal(err)
	}
	if transport.gotMethod != wire.MethodModelStatesGetTable || view.Generation != 4 || view.ActiveMemberID != 1 {
		t.Fatalf("method %q, view %+v: expected generation 4 and Primary", transport.gotMethod, view)
	}
	value, ok := view.Properties[0].Cells[0].Value.Str()
	if !ok || value != "" || view.Properties[0].Base != nil {
		t.Fatalf("property %+v: expected an authored empty string and absent base", view.Properties[0])
	}
	if view.Suppressions[0].Cells[0].Suppressed == nil || !*view.Suppressions[0].Cells[0].Suppressed || view.Parameters[0].Cells[0].Expression != nil {
		t.Fatalf("suppression %+v, parameter %+v: expected suppression and inherited expression", view.Suppressions, view.Parameters)
	}
}

func TestVerifyModelStateTableReturnsCellDiagnostics(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"valid":false,"diagnostics":[{"row":{"memberId":1},"parameterId":7,"expression":"abc","message":"unknown name"}]}`)}
	request := wire.SetModelStateTableArgs{ExpectedGeneration: 4, Rows: []wire.ModelStateTableRow{}, Columns: []wire.ModelStateTableColumnRef{}, Properties: []wire.ModelStateTableProperty{}, Suppressions: []wire.ModelStateTableSuppression{}, Parameters: []wire.ModelStateTableParameter{}}
	result, err := New(transport).ModelStates().VerifyTable(request)
	if err != nil {
		t.Fatal(err)
	}
	if transport.gotMethod != wire.MethodModelStatesVerifyTable || result.Valid || len(result.Diagnostics) != 1 || result.Diagnostics[0].ParameterID != 7 || result.Diagnostics[0].Row.MemberID != 1 {
		t.Fatalf("method %q result %+v: expected the invalid cell", transport.gotMethod, result)
	}
	var decoded wire.SetModelStateTableArgs
	if err := json.Unmarshal(transport.gotReq, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded, request) {
		t.Fatalf("request %+v: expected %+v", decoded, request)
	}
}
