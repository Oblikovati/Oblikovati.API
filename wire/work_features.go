// SPDX-License-Identifier: Apache-2.0

package wire

// SetWorkFeatureVisibleArgs is the request of [MethodWorkFeaturesSetVisible]: show or hide the
// datum work plane, axis, or point named by Ref (a "plane/N" / "axis/N" / "point/N" reference, or
// an origin-frame ref like "origin/plane/xy"). Visible=false hides it in the viewport. #1856.
type SetWorkFeatureVisibleArgs struct {
	Ref     string `json:"ref"`
	Visible bool   `json:"visible"`
}
