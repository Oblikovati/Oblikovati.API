// SPDX-License-Identifier: Apache-2.0

package client

import (
	"slices"
	"strings"
	"testing"

	"oblikovati.org/api/wire"
)

// TestDerivedParameterTablesRoundTrip drives every table method against a
// fake transport: the right wire constant goes out and the table DTO decodes
// back.
func TestDerivedParameterTablesRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{
		"id":1,"sourceDocument":"base.obk","linked":["od","wall"],
		"available":["od","wall","len"],"health":""}`)}
	c := New(ft)

	got, err := c.Parameters().AddDerivedTable(wire.DerivedParameterTableAddArgs{
		SourceDocument: "base.obk", Linked: []string{"od", "wall"},
	})
	if err != nil || ft.gotMethod != wire.MethodParametersDerivedTablesAdd {
		t.Fatalf("AddDerivedTable = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodParametersDerivedTablesAdd)
	}
	if got.ID != 1 || got.SourceDocument != "base.obk" || !slices.Equal(got.Linked, []string{"od", "wall"}) {
		t.Errorf("table = %+v, want the linked base.obk table back", got)
	}
	if len(got.Available) != 3 {
		t.Errorf("available = %v, want the 3 candidates", got.Available)
	}

	if _, err := c.Parameters().SetDerivedTableLinked(wire.DerivedParameterTableSetLinkedArgs{
		ID: 1, Linked: []string{"od"},
	}); err != nil || ft.gotMethod != wire.MethodParametersDerivedTablesSetLinked {
		t.Errorf("SetDerivedTableLinked = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodParametersDerivedTablesSetLinked)
	}
	if !strings.Contains(string(ft.gotReq), `"id":1`) {
		t.Errorf("sent args = %s, want the table id", ft.gotReq)
	}

	ft.reply = []byte(`{"tables":[{"id":1,"sourceDocument":"base.obk"}]}`)
	list, err := c.Parameters().ListDerivedTables()
	if err != nil || ft.gotMethod != wire.MethodParametersDerivedTablesList || len(list.Tables) != 1 {
		t.Errorf("ListDerivedTables = (%+v, %q, %v), want one table back", list, ft.gotMethod, err)
	}

	ft.reply = []byte(`{}`)
	if err := c.Parameters().DeleteDerivedTable(1); err != nil || ft.gotMethod != wire.MethodParametersDerivedTablesDelete {
		t.Errorf("DeleteDerivedTable = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodParametersDerivedTablesDelete)
	}
}
