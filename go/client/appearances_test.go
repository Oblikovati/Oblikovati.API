// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

// TestAppearancesList asserts List hits appearances.list and decodes the reply.
func TestAppearancesList(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"appearances":[{"id":"brushed-steel","displayName":"Brushed Steel","source":"builtin"}]}`)}
	c := New(ft)

	got, err := c.Appearances().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodAppearancesList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAppearancesList)
	}
	if len(got.Appearances) != 1 || got.Appearances[0].ID != "brushed-steel" {
		t.Errorf("decoded = %+v, want one brushed-steel entry", got)
	}
}

// TestAppearancesGet asserts Get addresses the id and hits appearances.get.
func TestAppearancesGet(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":"brushed-steel","displayName":"Brushed Steel","source":"builtin"}`)}
	c := New(ft)

	got, err := c.Appearances().Get("brushed-steel")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if ft.gotMethod != wire.MethodAppearancesGet {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAppearancesGet)
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

// TestAppearancesCreate asserts Create sends BaseID/Name and hits appearances.create.
func TestAppearancesCreate(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":"my-steel","displayName":"My Steel","source":"project"}`)}
	c := New(ft)

	got, err := c.Appearances().Create(wire.CreateAppearanceArgs{BaseID: "brushed-steel", Name: "My Steel"})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if ft.gotMethod != wire.MethodAppearancesCreate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAppearancesCreate)
	}
	var sent wire.CreateAppearanceArgs
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

// TestAppearancesUpdate asserts Update sends the full args and hits appearances.update.
func TestAppearancesUpdate(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":"my-steel","displayName":"My Steel","source":"project"}`)}
	c := New(ft)

	args := wire.UpdateAppearanceArgs{ID: "my-steel"}
	got, err := c.Appearances().Update(args)
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
	if ft.gotMethod != wire.MethodAppearancesUpdate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAppearancesUpdate)
	}
	var sent wire.UpdateAppearanceArgs
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

// TestAppearancesAssign asserts Assign sends the scope/key/id and hits
// model.assignAppearance.
func TestAppearancesAssign(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)

	args := wire.AssignAppearanceArgs{Scope: "body", Key: "body/0", AppearanceID: "my-steel"}
	if _, err := c.Appearances().Assign(args); err != nil {
		t.Fatalf("Assign: %v", err)
	}
	if ft.gotMethod != wire.MethodModelAssignAppearance {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodModelAssignAppearance)
	}
	var sent wire.AssignAppearanceArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent != args {
		t.Errorf("sent = %+v, want %+v", sent, args)
	}
}
