// SPDX-License-Identifier: Apache-2.0

package contract

// SketchBlockDefinition is the scalar view of a reusable sketch entity group
// owned by a part's component definition (M06-F07,
// Oblikovati/Oblikovati#622). The host's model/sketch.BlockDefinition
// satisfies this via a compile-time assertion.
type SketchBlockDefinition interface {
	// Name is the definition's unique name within the part.
	Name() string
	// EntityCount is the number of entities the definition holds.
	EntityCount() int
	// InstanceCount is how many placed instances reference this definition.
	InstanceCount() int
}

// SketchBlock is the scalar view of one placed block instance in a sketch.
// The host's model/sketch.BlockInstance satisfies this.
type SketchBlock interface {
	// DefinitionName names the block definition this instance places.
	DefinitionName() string
	// EntityCount is the live entity count of the instanced definition.
	EntityCount() int
}
