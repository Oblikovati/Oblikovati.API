// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Materials is the material operation group: list/read materials, create/edit custom
// ones, assign them to bodies, and read a part's physical properties.
type Materials struct{ c *Client }

// Materials returns the material operation group.
func (c *Client) Materials() Materials { return Materials{c} }

// List returns every material available to the active document.
//
// mcp:tool list_materials
// mcp:summary List the document's materials.
// mcp:digest summarizeMaterials
func (m Materials) List() (wire.ListMaterialsResult, error) {
	return call[wire.ListMaterialsResult](m.c, wire.MethodMaterialsList, nil)
}

// Get returns one material by id.
//
// mcp:tool get_material
// mcp:summary Get one material by id.
func (m Materials) Get(id string) (wire.MaterialInfo, error) {
	return call[wire.MaterialInfo](m.c, wire.MethodMaterialsGet, wire.AssetRefArgs{ID: id})
}

// Create duplicates an existing material into a new editable one under name.
//
// mcp:tool create_material
// mcp:summary Duplicate an existing material into a new editable one under a name.
func (m Materials) Create(args wire.DuplicateAssetArgs) (wire.MaterialInfo, error) {
	return call[wire.MaterialInfo](m.c, wire.MethodMaterialsCreate, args)
}

// Update writes the editable fields of a material (by its id) and returns the result.
//
// mcp:tool update_material
// mcp:summary Update a material's editable fields (identified by its id).
func (m Materials) Update(info wire.MaterialInfo) (wire.MaterialInfo, error) {
	return call[wire.MaterialInfo](m.c, wire.MethodMaterialsUpdate, info)
}

// Assign sets a body's material (or the part default when BodyKey is empty).
//
// mcp:tool assign_material
// mcp:summary Assign a material to the active part (or a selected body).
func (m Materials) Assign(args wire.AssignMaterialArgs) (wire.OKResult, error) {
	return call[wire.OKResult](m.c, wire.MethodModelAssignMaterial, args)
}

// PhysicalProperties returns the active part's computed mass/volume/area/centroid.
//
// mcp:tool get_physical_properties
// mcp:summary Read the active part's physical properties (mass, volume, area, center of mass).
func (m Materials) PhysicalProperties() (types.PhysicalProperties, error) {
	return call[types.PhysicalProperties](m.c, wire.MethodModelPhysicalProperties, nil)
}
