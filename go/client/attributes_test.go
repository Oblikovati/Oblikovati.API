// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// mustJSON marshals v for use as a fake transport reply.
func mustJSON(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal reply: %v", err)
	}
	return b
}

// decodeReq unmarshals a captured request body for assertions.
func decodeReq(t *testing.T, b []byte, v any) {
	t.Helper()
	if err := json.Unmarshal(b, v); err != nil {
		t.Fatalf("decode request %s: %v", b, err)
	}
}

func TestAttributesLifecycle(t *testing.T) {
	ft := &fakeTransport{}
	c := New(ft)

	// Set stores a typed value and echoes it back.
	stored := wire.AttributeInfo{Set: "acme.bom", Name: "partNo", Value: types.StringVariant("X-12")}
	ft.reply = mustJSON(t, wire.AttributeResult{Attribute: stored, Found: true})
	res, err := c.Attributes().Set(7, "acme.bom", "partNo", types.StringVariant("X-12"))
	if err != nil || !res.Found || res.Attribute.Name != "partNo" {
		t.Fatalf("Set = (%+v, %v), want the stored attribute", res, err)
	}
	var sent wire.SetAttributeArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Document != 7 || sent.Set != "acme.bom" {
		t.Errorf("sent = %s, want the set args for document 7", ft.gotReq)
	}
	if v, ok := sent.Value.Str(); !ok || v != "X-12" {
		t.Errorf("sent value = %+v, want string X-12", sent.Value)
	}

	// Get reads it back.
	ft.reply = mustJSON(t, wire.AttributeResult{Attribute: stored, Found: true})
	got, err := c.Attributes().Get(7, "acme.bom", "partNo")
	if err != nil || !got.Found {
		t.Fatalf("Get = (%+v, %v), want found", got, err)
	}

	// List filtered to a set.
	ft.reply = mustJSON(t, wire.ListAttributesResult{Attributes: []wire.AttributeInfo{stored}})
	lst, err := c.Attributes().List(7, "acme.bom")
	if err != nil || len(lst.Attributes) != 1 {
		t.Fatalf("List = (%+v, %v), want one attribute", lst, err)
	}

	// ListSets.
	ft.reply = mustJSON(t, wire.ListAttributeSetsResult{Sets: []string{"acme.bom"}})
	sets, err := c.Attributes().ListSets(7)
	if err != nil || len(sets.Sets) != 1 || sets.Sets[0] != "acme.bom" {
		t.Fatalf("ListSets = (%+v, %v), want [acme.bom]", sets, err)
	}

	// Delete one attribute.
	ft.reply = mustJSON(t, wire.DeleteAttributeResult{Removed: 1})
	del, err := c.Attributes().Delete(7, "acme.bom", "partNo")
	if err != nil || del.Removed != 1 {
		t.Fatalf("Delete = (%+v, %v), want Removed 1", del, err)
	}

	// Find across documents.
	ft.reply = mustJSON(t, wire.FindByAttributeResult{Matches: []wire.AttributeMatch{{Document: 7, Attribute: stored}}})
	found, err := c.Attributes().Find("acme.bom", "partNo")
	if err != nil || len(found.Matches) != 1 || found.Matches[0].Document != 7 {
		t.Fatalf("Find = (%+v, %v), want one match on document 7", found, err)
	}
	var fa wire.FindByAttributeArgs
	if err := json.Unmarshal(ft.gotReq, &fa); err != nil || fa.Set != "acme.bom" || fa.Name != "partNo" {
		t.Errorf("sent find = %s, want set/name filter", ft.gotReq)
	}
}

// TestAttributesTarget checks the target-aware methods thread the entity reference key through
// the wire args (and the document-scoped wrappers send an empty target).
func TestAttributesTarget(t *testing.T) {
	ft := &fakeTransport{}
	c := New(ft)
	const bodyKey = "body-ref-key-abc"

	// SetOn anchors to the target.
	ft.reply = mustJSON(t, wire.AttributeResult{Attribute: wire.AttributeInfo{Value: types.DoubleVariant(1000)}, Found: true})
	if _, err := c.Attributes().SetOn(7, bodyKey, "traceon", "voltage", types.DoubleVariant(1000)); err != nil {
		t.Fatalf("SetOn: %v", err)
	}
	var set wire.SetAttributeArgs
	decodeReq(t, ft.gotReq, &set)
	if set.Target != bodyKey {
		t.Errorf("SetOn target = %q, want %q", set.Target, bodyKey)
	}

	// GetOn / DeleteOn carry the target.
	ft.reply = mustJSON(t, wire.AttributeResult{Attribute: wire.AttributeInfo{Value: types.StringVariant("")}, Found: false})
	_, _ = c.Attributes().GetOn(7, bodyKey, "traceon", "voltage")
	var get wire.GetAttributeArgs
	decodeReq(t, ft.gotReq, &get)
	if get.Target != bodyKey {
		t.Errorf("GetOn target = %q, want %q", get.Target, bodyKey)
	}
	ft.reply = mustJSON(t, wire.DeleteAttributeResult{Removed: 1})
	_, _ = c.Attributes().DeleteOn(7, bodyKey, "traceon", "voltage")
	var del wire.DeleteAttributeArgs
	decodeReq(t, ft.gotReq, &del)
	if del.Target != bodyKey {
		t.Errorf("DeleteOn target = %q, want %q", del.Target, bodyKey)
	}

	// ListOn carries the target; ListAll sets AllTargets; the document-scoped wrappers do neither.
	ft.reply = mustJSON(t, wire.ListAttributesResult{})
	_, _ = c.Attributes().ListOn(7, bodyKey, "traceon")
	var lon wire.ListAttributesArgs
	decodeReq(t, ft.gotReq, &lon)
	if lon.Target != bodyKey || lon.AllTargets {
		t.Errorf("ListOn = %+v, want target=%q allTargets=false", lon, bodyKey)
	}
	ft.reply = mustJSON(t, wire.ListAttributesResult{})
	_, _ = c.Attributes().ListAll(7, "traceon")
	var lall wire.ListAttributesArgs
	decodeReq(t, ft.gotReq, &lall)
	if !lall.AllTargets || lall.Target != "" {
		t.Errorf("ListAll = %+v, want allTargets=true target=empty", lall)
	}
	ft.reply = mustJSON(t, wire.AttributeResult{Attribute: wire.AttributeInfo{Value: types.StringVariant("")}, Found: false})
	_, _ = c.Attributes().Get(7, "traceon", "voltage")
	var getDoc wire.GetAttributeArgs
	decodeReq(t, ft.gotReq, &getDoc)
	if getDoc.Target != "" {
		t.Errorf("document-scoped Get target = %q, want empty", getDoc.Target)
	}
}
