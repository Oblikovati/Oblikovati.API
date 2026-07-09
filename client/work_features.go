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
