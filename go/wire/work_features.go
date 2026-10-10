// SPDX-License-Identifier: Apache-2.0

package wire

// SetWorkFeatureVisibleArgs is the request of [MethodWorkFeaturesSetVisible]: show or hide the
// datum work plane, axis, or point named by Ref (a "plane/N" / "axis/N" / "point/N" reference, or
// an origin-frame ref like "origin/plane/xy"). Visible=false hides it in the viewport. #1856.
type SetWorkFeatureVisibleArgs struct {
	Ref     string `json:"ref"`
	Visible bool   `json:"visible"`
}

// DeleteWorkFeatureArgs is the request of [MethodWorkFeaturesDelete]: remove the user datum work
// feature (plane, axis, or point) named by Ref (a "plane/N" / "axis/N" / "point/N" reference).
//
// RetainDependents mirrors the reference CAD API's WorkFeature.Delete(RetainDependents): when false (the reference CAD API's
// default) every user work feature that references the datum, directly or transitively, is deleted
// with it; when true those dependents are left in place and go unhealthy, since their reference no
// longer resolves. Deleting an origin / coordinate-system datum, an unknown ref, or an
// already-deleted datum fails with a reason. #1855.
type DeleteWorkFeatureArgs struct {
	Ref              string `json:"ref"`
	RetainDependents bool   `json:"retainDependents,omitempty"`
}

// DeleteWorkFeatureResult is the response of [MethodWorkFeaturesDelete]: the references of every
// datum the call removed — the named datum plus, when RetainDependents was false, its cascaded
// dependents. Deletion is a tombstone: the datum's slot is kept so surviving datums keep their
// stable "plane/N" / "axis/N" / "point/N" references, but a deleted datum no longer appears in any
// list and no longer resolves as a reference. #1855.
type DeleteWorkFeatureResult struct {
	Deleted []string `json:"deleted"`
}
