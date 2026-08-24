// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// OpenPBRAppearances is the OpenPBR Surface appearance operation group: list/read the
// full-lobe appearances available to the active document and create/edit custom ones —
// additive alongside [Appearances]' metallic-roughness subset.
type OpenPBRAppearances struct{ c *Client }

// OpenPBRAppearances returns the OpenPBR appearance operation group.
func (c *Client) OpenPBRAppearances() OpenPBRAppearances { return OpenPBRAppearances{c} }

// List returns every OpenPBR appearance available to the active document (built-in,
// project, and document-embedded).
//
// mcp:tool list_openpbr_appearances
// mcp:summary List the document's OpenPBR Surface appearances.
// mcp:digest summarizeOpenPBRAppearances
func (a OpenPBRAppearances) List() (wire.ListOpenPBRAppearancesResult, error) {
	return call[wire.ListOpenPBRAppearancesResult](a.c, wire.MethodOpenPBRAppearancesList, nil)
}

// Get returns one OpenPBR appearance by id.
//
// mcp:tool get_openpbr_appearance
// mcp:summary Get one OpenPBR Surface appearance by id.
func (a OpenPBRAppearances) Get(id string) (wire.OpenPBRAppearanceInfo, error) {
	return call[wire.OpenPBRAppearanceInfo](a.c, wire.MethodOpenPBRAppearancesGet, wire.AssetRefArgs{ID: id})
}

// Create duplicates an existing OpenPBR appearance into a new editable one under name.
//
// mcp:tool create_openpbr_appearance
// mcp:summary Duplicate an existing OpenPBR Surface appearance into a new editable one.
func (a OpenPBRAppearances) Create(args wire.CreateOpenPBRAppearanceArgs) (wire.OpenPBRAppearanceInfo, error) {
	return call[wire.OpenPBRAppearanceInfo](a.c, wire.MethodOpenPBRAppearancesCreate, args)
}

// Update writes every group of an OpenPBR appearance (by its id) and returns the result.
//
// mcp:tool update_openpbr_appearance
// mcp:summary Update an OpenPBR Surface appearance's groups (identified by its id).
func (a OpenPBRAppearances) Update(args wire.UpdateOpenPBRAppearanceArgs) (wire.OpenPBRAppearanceInfo, error) {
	return call[wire.OpenPBRAppearanceInfo](a.c, wire.MethodOpenPBRAppearancesUpdate, args)
}

// Assign overrides the OpenPBR appearance at a scope ("part", "body", or "face"); Key is
// the hex reference key of the target (empty for the part default).
//
// mcp:tool assign_openpbr_appearance
// mcp:summary Assign an OpenPBR Surface appearance to the active part (or a selected body).
func (a OpenPBRAppearances) Assign(args wire.AssignOpenPBRAppearanceArgs) (wire.OKResult, error) {
	return call[wire.OKResult](a.c, wire.MethodModelAssignOpenPBRAppearance, args)
}
