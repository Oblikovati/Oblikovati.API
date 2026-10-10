// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// WorkFeatures groups the operations that act on any datum work feature (plane, axis, or point) of
// the active part or assembly, addressed by its reference.
type WorkFeatures struct{ c *Client }

// WorkFeatures returns the datum work-feature operation group.
func (c *Client) WorkFeatures() WorkFeatures { return WorkFeatures{c} }

// SetVisible shows or hides the datum work plane, axis, or point named by ref (#1856).
//
// mcp:tool set_work_feature_visible
// mcp:summary Show or hide a datum work plane, axis, or point by its ref (e.g. "plane/1", "axis/0", "point/2", or an origin ref like "origin/plane/xy").
func (w WorkFeatures) SetVisible(ref string, visible bool) (wire.OKResult, error) {
	return call[wire.OKResult](w.c, wire.MethodWorkFeaturesSetVisible, wire.SetWorkFeatureVisibleArgs{Ref: ref, Visible: visible})
}

// Delete removes the user datum work plane, axis, or point named by ref (#1855). With
// retainDependents=false (the reference CAD API's default) every user work feature that references the datum,
// directly or transitively, is deleted with it; with retainDependents=true those dependents are
// left in place and go unhealthy. Deletion is a tombstone — surviving datums keep their stable
// refs. Deleting an origin datum, an unknown ref, or an already-deleted datum is an error. Returns
// the refs of every datum removed.
//
// mcp:tool delete_work_feature
// mcp:summary Delete a user datum work plane, axis, or point by its ref (e.g. "plane/3", "axis/1", "point/2"). retainDependents=false (default) also deletes any work feature built on it; true leaves those dependents in place (they go unhealthy). Origin datums cannot be deleted. Returns the refs removed.
func (w WorkFeatures) Delete(ref string, retainDependents bool) (wire.DeleteWorkFeatureResult, error) {
	return call[wire.DeleteWorkFeatureResult](w.c, wire.MethodWorkFeaturesDelete, wire.DeleteWorkFeatureArgs{Ref: ref, RetainDependents: retainDependents})
}
