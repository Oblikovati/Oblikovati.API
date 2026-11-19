// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestPropertyTableGetRoutesAndDecodesMemberValues(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"generation":3,"members":[{"id":1,"name":"Primary","active":true}],"availableProperties":[{"set":"Summary Information","name":"Category"}],"columns":[{"property":{"set":"Summary Information","name":"Category"},"defaultValue":null,"cells":[{"memberId":1,"value":{"type":"string","value":"Prototype"}}]}]}`)}
	view, err := New(transport).ModelStates().GetPropertyTable()
	if err != nil {
		t.Fatal(err)
	}
	if transport.gotMethod != wire.MethodModelStatesGetPropertyTable || view.Generation != 3 || len(view.Columns) != 1 {
		t.Fatalf("method %q, view %+v: expected property table at generation 3", transport.gotMethod, view)
	}
	value, ok := view.Columns[0].Cells[0].Value.Str()
	if !ok || value != "Prototype" {
		t.Fatalf("value %q, string %v: expected Prototype", value, ok)
	}
}

func TestPropertyTableSetPreservesGenerationAndUnsetCells(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"generation":4,"members":[],"availableProperties":[],"columns":[]}`)}
	generation := uint64(0)
	value := types.StringVariant("")
	base := types.StringVariant("Design")
	args := wire.SetModelStatePropertyTableArgs{ExpectedGeneration: &generation, Columns: []wire.ModelStatePropertyColumn{{Property: types.ModelStatePropertyKey{Set: "Summary Information", Name: "Category"}, DefaultValue: &base, Cells: []wire.ModelStatePropertyCell{{MemberID: 1, Value: nil}, {MemberID: 2, Value: &value}}}}}
	if _, err := New(transport).ModelStates().SetPropertyTable(args); err != nil {
		t.Fatal(err)
	}
	if transport.gotMethod != wire.MethodModelStatesSetPropertyTable {
		t.Fatalf("method %q: expected property table set", transport.gotMethod)
	}
	var decoded wire.SetModelStatePropertyTableArgs
	if err := json.Unmarshal(transport.gotReq, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.ExpectedGeneration == nil || *decoded.ExpectedGeneration != 0 || decoded.Columns[0].Cells[0].Value != nil {
		t.Fatalf("decoded %+v: expected generation zero and unset first cell", decoded)
	}
	inherited, ok := decoded.Columns[0].DefaultValue.Str()
	if !ok || inherited != "Design" {
		t.Fatalf("default %q, string %v: expected Design", inherited, ok)
	}
	blank, ok := decoded.Columns[0].Cells[1].Value.Str()
	if !ok || blank != "" {
		t.Fatalf("second cell %q, string %v: expected authored empty string", blank, ok)
	}
}

func TestPropertyTableGetPreservesHiddenMemberValues(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"generation":5,"members":[],"availableProperties":[],"columns":[],"retainedColumns":[{"property":{"set":"Summary Information","name":"Category"},"defaultValue":null,"cells":[{"memberId":1,"value":{"type":"string","value":"Review"}}]}]}`)}
	view, err := New(transport).ModelStates().GetPropertyTable()
	if err != nil {
		t.Fatal(err)
	}
	if len(view.RetainedColumns) != 1 || len(view.RetainedColumns[0].Cells) != 1 {
		t.Fatalf("retained columns %+v: expected one hidden member value", view.RetainedColumns)
	}
	value, ok := view.RetainedColumns[0].Cells[0].Value.Str()
	if !ok || value != "Review" {
		t.Fatalf("value %q, string %v: expected Review", value, ok)
	}
}
