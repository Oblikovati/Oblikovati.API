// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"

	"oblikovati.org/api/wire"
)

// Features is the feature-creation operation group for the active part.
type Features struct{ c *Client }

// Features returns the feature operation group.
func (c *Client) Features() Features { return Features{c} }

// List returns every feature operation Add can create, each with its args schema.
//
// mcp:tool list_feature_kinds
// mcp:summary List the feature operations add_feature can create, each with its JSON args schema.
func (f Features) List() (wire.ListFeatureKindsResult, error) {
	return call[wire.ListFeatureKindsResult](f.c, wire.MethodFeaturesList, nil)
}

// Add applies a feature operation. The result shape is operation-specific (see the
// kind's schema from List), so it is returned as raw JSON for the caller to decode.
//
// mcp:tool add_feature
// mcp:summary Create a feature on the active part. Get the kind and its args schema from list_feature_kinds.
// mcp:input addFeatureArg
func (f Features) Add(args wire.AddFeatureArgs) (json.RawMessage, error) {
	return call[json.RawMessage](f.c, wire.MethodFeaturesAdd, args)
}

// Get returns one placed feature's state and editable scalars by its stable id
// (from model.tree), e.g. Get(7).
//
// mcp:tool features_get
// mcp:summary Returns one placed feature's state and editable scalars by its stable id (from model.tree), e.g.
func (f Features) Get(id uint64) (wire.FeatureDetailResult, error) {
	return call[wire.FeatureDetailResult](f.c, wire.MethodFeaturesGet, wire.FeatureRefArgs{ID: id})
}

// Edit edits a placed feature in place and recomputes: set editable scalars and/or re-pick
// its geometric references (a fillet's edges, an extrude's profile, a mirror's plane). E.g.
// Edit(wire.EditFeatureArgs{ID: 7, Scalars: []wire.ScalarEdit{{Index: 0, Value: "5 mm"}}}) or
// Edit(wire.EditFeatureArgs{ID: 7, Repick: []wire.FeatureRepick{{Slot: 0, Ref: edgeKey}}}).
//
// mcp:tool features_edit
// mcp:summary Edit a placed feature in place — set scalars and/or re-pick its geometric references (edges/faces/profile/plane) — then recompute.
func (f Features) Edit(args wire.EditFeatureArgs) (wire.FeatureDetailResult, error) {
	return call[wire.FeatureDetailResult](f.c, wire.MethodFeaturesEdit, args)
}

// Delete removes a placed feature from the history and recomputes, e.g. Delete(7).
//
// mcp:tool features_delete
// mcp:summary Removes a placed feature from the history and recomputes, e.g.
func (f Features) Delete(id uint64) (wire.DeleteFeatureResult, error) {
	return call[wire.DeleteFeatureResult](f.c, wire.MethodFeaturesDelete, wire.FeatureRefArgs{ID: id})
}

// Rename sets a feature's display name (the id stays stable), e.g. Rename(7, "Boss").
//
// mcp:tool features_rename
// mcp:summary Sets a feature's display name (the id stays stable), e.g.
func (f Features) Rename(id uint64, name string) (wire.FeatureDetailResult, error) {
	return call[wire.FeatureDetailResult](f.c, wire.MethodFeaturesRename, wire.RenameFeatureArgs{ID: id, Name: name})
}

// SetSuppressed sets explicit suppression and recomputes, e.g. SetSuppressed(7, true).
//
// mcp:tool features_set_suppressed
// mcp:summary Sets explicit suppression and recomputes, e.g.
func (f Features) SetSuppressed(id uint64, suppressed bool) (wire.FeatureDetailResult, error) {
	args := wire.SetFeatureSuppressedArgs{ID: id, Suppressed: suppressed}
	return call[wire.FeatureDetailResult](f.c, wire.MethodFeaturesSetSuppressed, args)
}

// Reorder moves a feature to a new history index and recomputes, e.g. Reorder(7, 0).
//
// mcp:tool features_reorder
// mcp:summary Moves a feature to a new history index and recomputes, e.g.
func (f Features) Reorder(id uint64, newIndex int) (wire.FeatureDetailResult, error) {
	args := wire.ReorderFeatureArgs{ID: id, NewIndex: newIndex}
	return call[wire.FeatureDetailResult](f.c, wire.MethodFeaturesReorder, args)
}
