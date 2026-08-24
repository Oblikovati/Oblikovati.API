// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

// TestOpenPBRAppearancesList asserts List hits openpbrAppearances.list and decodes the
// reply.
func TestOpenPBRAppearancesList(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"appearances":[{"id":"brushed-steel","displayName":"Brushed Steel","source":"builtin"}]}`)}
	c := New(ft)

	got, err := c.OpenPBRAppearances().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodOpenPBRAppearancesList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodOpenPBRAppearancesList)
	}
	if len(got.Appearances) != 1 || got.Appearances[0].ID != "brushed-steel" {
		t.Errorf("decoded = %+v, want one brushed-steel entry", got)
	}
}

// TestOpenPBRAppearancesGet asserts Get addresses the id and hits openpbrAppearances.get.
func TestOpenPBRAppearancesGet(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":"brushed-steel","displayName":"Brushed Steel","source":"builtin"}`)}
	c := New(ft)

	got, err := c.OpenPBRAppearances().Get("brushed-steel")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if ft.gotMethod != wire.MethodOpenPBRAppearancesGet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodOpenPBRAppearancesGet)
	}
	var sent wire.AssetRefArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ID != "brushed-steel" {
		t.Errorf("sent id = %q, want brushed-steel", sent.ID)
	}
	if got.DisplayName != "Brushed Steel" {
		t.Errorf("decoded display name = %q, want Brushed Steel", got.DisplayName)
	}
}

// TestOpenPBRAppearancesCreate asserts Create sends BaseID/Name and hits
// openpbrAppearances.create.
func TestOpenPBRAppearancesCreate(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":"my-steel","displayName":"My Steel","source":"project"}`)}
	c := New(ft)

	got, err := c.OpenPBRAppearances().Create(wire.CreateOpenPBRAppearanceArgs{BaseID: "brushed-steel", Name: "My Steel"})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if ft.gotMethod != wire.MethodOpenPBRAppearancesCreate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodOpenPBRAppearancesCreate)
	}
	var sent wire.CreateOpenPBRAppearanceArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.BaseID != "brushed-steel" || sent.Name != "My Steel" {
		t.Errorf("sent = %+v, want baseId brushed-steel / name My Steel", sent)
	}
	if got.ID != "my-steel" {
		t.Errorf("decoded id = %q, want my-steel", got.ID)
	}
}

// TestOpenPBRAppearancesUpdate asserts Update sends the full args and hits
// openpbrAppearances.update.
func TestOpenPBRAppearancesUpdate(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":"my-steel","displayName":"My Steel","source":"project"}`)}
	c := New(ft)

	args := wire.UpdateOpenPBRAppearanceArgs{ID: "my-steel"}
	got, err := c.OpenPBRAppearances().Update(args)
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
	if ft.gotMethod != wire.MethodOpenPBRAppearancesUpdate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodOpenPBRAppearancesUpdate)
	}
	var sent wire.UpdateOpenPBRAppearanceArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ID != "my-steel" {
		t.Errorf("sent id = %q, want my-steel", sent.ID)
	}
	if got.ID != "my-steel" {
		t.Errorf("decoded id = %q, want my-steel", got.ID)
	}
}

// TestOpenPBRAppearancesAssign asserts Assign sends the scope/key/id and hits
// model.assignOpenPBRAppearance.
func TestOpenPBRAppearancesAssign(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	args := wire.AssignOpenPBRAppearanceArgs{Scope: "body", Key: "body/0", AppearanceID: "my-steel"}
	if _, err := c.OpenPBRAppearances().Assign(args); err != nil {
		t.Fatalf("Assign: %v", err)
	}
	if ft.gotMethod != wire.MethodModelAssignOpenPBRAppearance {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodModelAssignOpenPBRAppearance)
	}
	var sent wire.AssignOpenPBRAppearanceArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent != args {
		t.Errorf("sent = %+v, want %+v", sent, args)
	}
}
