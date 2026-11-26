// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Sketch blocks (M06-F07, Oblikovati/Oblikovati#622): reusable entity groups
// owned by the part's component definition, instanced into sketches with a
// placement transform.

// CreateBlockDefinition creates a named block definition. With a source sketch
// and entity ids, the selection moves into the definition and one instance
// replaces it in place; without them, an empty definition is created.
//
// mcp:tool sketch_block_definitions_create
// mcp:summary Creates a named block definition.
func (s Sketch) CreateBlockDefinition(args wire.CreateBlockDefinitionArgs) (wire.SketchBlockDefinitionInfo, error) {
	return call[wire.SketchBlockDefinitionInfo](s.c, wire.MethodSketchBlockDefinitionCreate, args)
}

// BlockDefinitions enumerates the part's block definitions.
//
// mcp:tool sketch_block_definitions_list
// mcp:summary Enumerates the part's block definitions.
func (s Sketch) BlockDefinitions() (wire.ListBlockDefinitionsResult, error) {
	return call[wire.ListBlockDefinitionsResult](s.c, wire.MethodSketchBlockDefinitionList, nil)
}

// DeleteBlockDefinition removes a block definition by name. A definition that
// still has instances is rejected; the error names the consuming sketches.
//
// mcp:tool sketch_block_definitions_delete
// mcp:summary Removes a block definition by name.
func (s Sketch) DeleteBlockDefinition(name string) (wire.OKResult, error) {
	return call[wire.OKResult](s.c, wire.MethodSketchBlockDefinitionDelete, wire.DeleteBlockDefinitionArgs{Name: name})
}

// AddBlockInstance places an instance of a block definition in a sketch.
//
// mcp:tool sketch_add_block_instance
// mcp:summary Places an instance of a block definition in a sketch.
func (s Sketch) AddBlockInstance(args wire.AddSketchBlockArgs) (wire.AddSketchBlockResult, error) {
	return call[wire.AddSketchBlockResult](s.c, wire.MethodSketchAddBlockInstance, args)
}

// ExplodeBlockInstance replaces one instance with independent sketch entities.
// Example: result, err := sketch.ExplodeBlockInstance(wire.ExplodeSketchBlockArgs{SketchIndex: 0, InstanceID: 42}).
//
// mcp:tool sketch_explode_block_instance
// mcp:summary Replaces one sketch block instance with independent entities.
func (s Sketch) ExplodeBlockInstance(args wire.ExplodeSketchBlockArgs) (wire.ExplodeSketchBlockResult, error) {
	return call[wire.ExplodeSketchBlockResult](s.c, wire.MethodSketchExplodeBlockInstance, args)
}

// SetBlockDefinitionColor changes or clears a definition color override.
// Example: sketch.SetBlockDefinitionColor(wire.SetBlockDefinitionColorArgs{Name: "Rib"}).
//
// mcp:tool set_sketch_block_definition_color
// mcp:summary Set the color override for a sketch block definition. Use null to restore inheritance.
func (s Sketch) SetBlockDefinitionColor(args wire.SetBlockDefinitionColorArgs) (wire.SketchBlockDefinitionInfo, error) {
	return call[wire.SketchBlockDefinitionInfo](s.c, wire.MethodSketchSetBlockDefinitionColor, args)
}

// SetBlockInstanceColor changes or clears an instance color override.
// Example: sketch.SetBlockInstanceColor(wire.SetBlockInstanceColorArgs{InstanceID: 42}).
//
// mcp:tool set_sketch_block_instance_color
// mcp:summary Set the color override for a placed sketch block instance. Use null to restore inheritance.
func (s Sketch) SetBlockInstanceColor(args wire.SetBlockInstanceColorArgs) (wire.SketchBlockInfo, error) {
	return call[wire.SketchBlockInfo](s.c, wire.MethodSketchSetBlockInstanceColor, args)
}

// BlockInstances enumerates a sketch's placed block instances.
//
// mcp:tool sketch_block_instances
// mcp:summary Enumerates a sketch's placed block instances.
func (s Sketch) BlockInstances(index int) (wire.ListBlockInstancesResult, error) {
	return call[wire.ListBlockInstancesResult](s.c, wire.MethodSketchListBlockInstances, wire.SketchArgs{SketchIndex: index})
}
