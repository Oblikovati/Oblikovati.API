// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"slices"
	"testing"

	"oblikovati.org/api/wire"
)

// TestParameterGroupsRoundTrip drives every group method against a fake
// transport: the right wire constant goes out, args marshal by the DTO shape,
// and ParameterGroupInfo decodes back.
func TestParameterGroupsRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{
		"internalName":"com.example.gears:ratios","displayName":"Gear Ratios",
		"clientId":"com.example.gears","members":["module","teeth"]}`)}
	c := New(ft)

	got, err := c.Parameters().AddGroup(wire.ParameterGroupAddArgs{
		InternalName: "com.example.gears:ratios", DisplayName: "Gear Ratios", ClientID: "com.example.gears",
	})
	if err != nil {
		t.Fatalf("AddGroup: %v", err)
	}
	if ft.gotMethod != wire.MethodParametersGroupsAdd {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodParametersGroupsAdd)
	}
	if got.InternalName != "com.example.gears:ratios" || got.DisplayName != "Gear Ratios" {
		t.Errorf("group = %+v, want the gears group back", got)
	}
	if got.ClientID != "com.example.gears" || !slices.Equal(got.Members, []string{"module", "teeth"}) {
		t.Errorf("provenance/members = %q %v, want com.example.gears [module teeth]", got.ClientID, got.Members)
	}

	if _, err := c.Parameters().AddGroupMember(wire.ParameterGroupMemberArgs{
		InternalName: "com.example.gears:ratios", Parameter: "module",
	}); err != nil || ft.gotMethod != wire.MethodParametersGroupsAddMember {
		t.Errorf("AddGroupMember = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodParametersGroupsAddMember)
	}
	if _, err := c.Parameters().RemoveGroupMember(wire.ParameterGroupMemberArgs{
		InternalName: "com.example.gears:ratios", Parameter: "module",
	}); err != nil || ft.gotMethod != wire.MethodParametersGroupsRemoveMember {
		t.Errorf("RemoveGroupMember = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodParametersGroupsRemoveMember)
	}
	if _, err := c.Parameters().SetGroupDisplayName(wire.ParameterGroupDisplayNameArgs{
		InternalName: "com.example.gears:ratios", DisplayName: "Ratios",
	}); err != nil || ft.gotMethod != wire.MethodParametersGroupsSetDisplayName {
		t.Errorf("SetGroupDisplayName = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodParametersGroupsSetDisplayName)
	}
}

// TestParameterGroupsListAndDelete covers the two remaining methods: list
// decodes the collection, delete sends the opt-in cascade flag.
func TestParameterGroupsListAndDelete(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"groups":[{"internalName":"g1","displayName":"G1"}]}`)}
	c := New(ft)

	list, err := c.Parameters().ListGroups()
	if err != nil || ft.gotMethod != wire.MethodParametersGroupsList {
		t.Fatalf("ListGroups = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodParametersGroupsList)
	}
	if len(list.Groups) != 1 || list.Groups[0].InternalName != "g1" {
		t.Errorf("groups = %+v, want [g1]", list.Groups)
	}

	ft.reply = []byte(`{}`)
	if err := c.Parameters().DeleteGroup(wire.ParameterGroupDeleteArgs{
		InternalName: "g1", DeleteParameters: true,
	}); err != nil || ft.gotMethod != wire.MethodParametersGroupsDelete {
		t.Fatalf("DeleteGroup = (%q, %v), want %q sent", ft.gotMethod, err, wire.MethodParametersGroupsDelete)
	}
	var sent map[string]json.RawMessage
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("sent args not JSON: %v", err)
	}
	if string(sent["deleteParameters"]) != "true" {
		t.Errorf("sent args = %s, want the deleteParameters cascade flag", ft.gotReq)
	}
}
