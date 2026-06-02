// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

// Appearances is the appearance operation group: list/read the PBR appearances available
// to the active document and create/edit custom ones.
type Appearances struct{ c *Client }

// Appearances returns the appearance operation group.
func (c *Client) Appearances() Appearances { return Appearances{c} }

// List returns every appearance available to the active document (built-in, project, and
// document-embedded).
func (a Appearances) List() (wire.ListAppearancesResult, error) {
	var r wire.ListAppearancesResult
	return r, a.c.call(wire.MethodAppearancesList, nil, &r)
}

// Get returns one appearance by id.
func (a Appearances) Get(id string) (wire.AppearanceInfo, error) {
	var r wire.AppearanceInfo
	return r, a.c.call(wire.MethodAppearancesGet, wire.AssetRefArgs{ID: id}, &r)
}

// Create duplicates an existing appearance into a new editable one under name.
func (a Appearances) Create(args wire.DuplicateAssetArgs) (wire.AppearanceInfo, error) {
	var r wire.AppearanceInfo
	return r, a.c.call(wire.MethodAppearancesCreate, args, &r)
}

// Update writes the editable fields of an appearance (by its id) and returns the result.
func (a Appearances) Update(info wire.AppearanceInfo) (wire.AppearanceInfo, error) {
	var r wire.AppearanceInfo
	return r, a.c.call(wire.MethodAppearancesUpdate, info, &r)
}

// Assign overrides the appearance at a scope ("part", "body", or "face"); Key is the hex
// reference key of the target (empty for the part default).
func (a Appearances) Assign(args wire.AssignAppearanceArgs) (wire.OKResult, error) {
	var r wire.OKResult
	return r, a.c.call(wire.MethodModelAssignAppearance, args, &r)
}
