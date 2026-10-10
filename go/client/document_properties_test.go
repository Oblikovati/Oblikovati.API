// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// TestDocumentsSetPropertySendsTypedValue asserts SetProperty marshals the set/name and the
// typed value (a string variant) and decodes the returned property (#156).
func TestDocumentsSetPropertySendsTypedValue(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"property":{"set":"Design Tracking Properties","name":"Part Number","value":{"type":"string","value":"BRK-001"}}}`)}
	c := New(ft)

	got, err := c.Documents().SetProperty(7, "Design Tracking Properties", "Part Number", types.StringVariant("BRK-001"))
	if err != nil {
		t.Fatalf("SetProperty: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentsSetProperty {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentsSetProperty)
	}
	var sent wire.SetPropertyArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Document != 7 || sent.Set != "Design Tracking Properties" || sent.Name != "Part Number" {
		t.Errorf("sent = %+v, want document 7 / Design Tracking / Part Number", sent)
	}
	if v, ok := sent.Value.Str(); !ok || v != "BRK-001" {
		t.Errorf("sent value = %v (ok=%v), want string BRK-001", v, ok)
	}
	if got.Property.Name != "Part Number" {
		t.Errorf("decoded property name = %q, want Part Number", got.Property.Name)
	}
}

// TestDocumentsGetPropertyAddressesProperty asserts GetProperty hits the right method with the
// document id, set, and name, and decodes the returned typed value (#156).
func TestDocumentsGetPropertyAddressesProperty(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"property":{"set":"Design Tracking Properties","name":"Part Number","value":{"type":"string","value":"BRK-001"}}}`)}
	c := New(ft)

	got, err := c.Documents().GetProperty(7, "Design Tracking Properties", "Part Number")
	if err != nil {
		t.Fatalf("GetProperty: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentsGetProperty {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentsGetProperty)
	}
	var sent wire.GetPropertyArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Document != 7 || sent.Set != "Design Tracking Properties" || sent.Name != "Part Number" {
		t.Errorf("sent = %+v, want document 7 / Design Tracking / Part Number", sent)
	}
	if got.Property.Name != "Part Number" {
		t.Errorf("decoded property name = %q, want Part Number", got.Property.Name)
	}
	if v, ok := got.Property.Value.Str(); !ok || v != "BRK-001" {
		t.Errorf("decoded value = %v (ok=%v), want string BRK-001", v, ok)
	}
}

// TestDocumentsListPropertiesAddressesDocument asserts ListProperties hits the right method with
// the document id and decodes the property list.
func TestDocumentsListPropertiesAddressesDocument(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"properties":[{"set":"Summary Information","name":"Title","value":{"type":"string","value":"Bracket"}}]}`)}
	c := New(ft)

	got, err := c.Documents().ListProperties(4)
	if err != nil {
		t.Fatalf("ListProperties: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentsListProperties {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentsListProperties)
	}
	var sent wire.ListPropertiesArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Document != 4 {
		t.Errorf("sent document = %d, want 4", sent.Document)
	}
	if len(got.Properties) != 1 || got.Properties[0].Name != "Title" {
		t.Errorf("decoded = %+v, want one Title property", got.Properties)
	}
}
