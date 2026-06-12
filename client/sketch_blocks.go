// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Sketch blocks (M06-F07, Oblikovati/Oblikovati#622): reusable entity groups
// owned by the part's component definition, instanced into sketches with a
// placement transform.

// CreateBlockDefinition creates a named block definition. With a source sketch
// and entity ids, the selection moves into the definition and one instance
// replaces it in place; without them, an empty definition is created.
func (s Sketch) CreateBlockDefinition(args wire.CreateBlockDefinitionArgs) (wire.SketchBlockDefinitionInfo, error) {
	var r wire.SketchBlockDefinitionInfo
	return r, s.c.call(wire.MethodSketchBlockDefinitionCreate, args, &r)
}

// BlockDefinitions enumerates the part's block definitions.
func (s Sketch) BlockDefinitions() (wire.ListBlockDefinitionsResult, error) {
	var r wire.ListBlockDefinitionsResult
	return r, s.c.call(wire.MethodSketchBlockDefinitionList, nil, &r)
}

// DeleteBlockDefinition removes a block definition by name. A definition that
// still has instances is rejected; the error names the consuming sketches.
func (s Sketch) DeleteBlockDefinition(name string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, s.c.call(wire.MethodSketchBlockDefinitionDelete,
		wire.DeleteBlockDefinitionArgs{Name: name}, &r)
}

// AddBlockInstance places an instance of a block definition in a sketch.
func (s Sketch) AddBlockInstance(args wire.AddSketchBlockArgs) (wire.AddSketchBlockResult, error) {
	var r wire.AddSketchBlockResult
	return r, s.c.call(wire.MethodSketchAddBlockInstance, args, &r)
}

// BlockInstances enumerates a sketch's placed block instances.
func (s Sketch) BlockInstances(index int) (wire.ListBlockInstancesResult, error) {
	var r wire.ListBlockInstancesResult
	return r, s.c.call(wire.MethodSketchListBlockInstances,
		wire.SketchArgs{SketchIndex: index}, &r)
}
