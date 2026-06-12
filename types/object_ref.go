// SPDX-License-Identifier: Apache-2.0

package types

// ObjectRef identifies one host-owned object across the contract: the numeric
// id every wire surface already hands out, qualified by the stable camelCase
// kind name ("sketchLine", "extrudeFeature", …) so heterogeneous collections
// stay self-describing (M00-F05, #601). The kind is an open string, not a
// frozen enum: each surface documents its kind names, and a collection can mix
// kinds the way the reference ObjectCollection mixes object types.
type ObjectRef struct {
	Kind string `json:"kind"`
	ID   uint64 `json:"id"`
}

// NewObjectRef builds a reference from a kind name and host id.
func NewObjectRef(kind string, id uint64) ObjectRef { return ObjectRef{Kind: kind, ID: id} }
