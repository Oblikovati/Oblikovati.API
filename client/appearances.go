// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Appearances is the appearance operation group: list/read the PBR appearances available
// to the active document and create/edit custom ones.
type Appearances struct{ c *Client }

// Appearances returns the appearance operation group.
func (c *Client) Appearances() Appearances { return Appearances{c} }

// List returns every appearance available to the active document (built-in, project, and
// document-embedded).
//
// mcp:tool list_appearances
// mcp:summary List the document's appearances (visual styles).
// mcp:digest summarizeAppearances
func (a Appearances) List() (wire.ListAppearancesResult, error) {
	return call[wire.ListAppearancesResult](a.c, wire.MethodAppearancesList, nil)
}

// Get returns one appearance by id.
//
// mcp:tool get_appearance
// mcp:summary Get one appearance by id.
func (a Appearances) Get(id string) (wire.AppearanceInfo, error) {
	return call[wire.AppearanceInfo](a.c, wire.MethodAppearancesGet, wire.AssetRefArgs{ID: id})
}

// Create duplicates an existing appearance into a new editable one under name.
//
// mcp:tool create_appearance
// mcp:summary Duplicate an existing appearance into a new editable one under a name.
func (a Appearances) Create(args wire.DuplicateAssetArgs) (wire.AppearanceInfo, error) {
	return call[wire.AppearanceInfo](a.c, wire.MethodAppearancesCreate, args)
}

// Update writes the editable fields of an appearance (by its id) and returns the result.
//
// mcp:tool update_appearance
// mcp:summary Update an appearance's editable fields (identified by its id).
func (a Appearances) Update(info wire.AppearanceInfo) (wire.AppearanceInfo, error) {
	return call[wire.AppearanceInfo](a.c, wire.MethodAppearancesUpdate, info)
}

// Assign overrides the appearance at a scope ("part", "body", or "face"); Key is the hex
// reference key of the target (empty for the part default).
//
// mcp:tool assign_appearance
// mcp:summary Assign an appearance to the active part (or a selected body).
func (a Appearances) Assign(args wire.AssignAppearanceArgs) (wire.OKResult, error) {
	return call[wire.OKResult](a.c, wire.MethodModelAssignAppearance, args)
}
