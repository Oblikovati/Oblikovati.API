// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Sketch blocks (M06-F07, Oblikovati/Oblikovati#622): a block definition is a
// named, self-contained group of sketch entities owned by the part's component
// definition; sketches place it as block instances, each with its own
// placement transform. Editing the definition updates every instance.

// CreateBlockDefinitionArgs is the request of [MethodSketchBlockDefinitionCreate].
// Name must be unique among the part's block definitions. When SourceSketchIndex
// and EntityRefs are set, the definition is created from that selection: the
// referenced entities (with their defining points) move out of the sketch into
// the definition and one instance replaces them in place (the interactive
// "create block from selection"). Without them, an empty definition is created
// to be populated by later edits.
type CreateBlockDefinitionArgs struct {
	Name              string   `json:"name"`
	SourceSketchIndex int      `json:"sourceSketchIndex,omitempty"`
	EntityRefs        []uint64 `json:"entityRefs,omitempty"`
}

// SketchBlockDefinitionInfo is one row of [MethodSketchBlockDefinitionList] and
// the result of [MethodSketchBlockDefinitionCreate]: the definition's identity
// plus how much geometry it holds and how widely it is instanced.
type SketchBlockDefinitionInfo struct {
	Index         int          `json:"index"`
	Name          string       `json:"name"`
	EntityCount   int          `json:"entityCount"`
	InstanceCount int          `json:"instanceCount"`
	Color         *types.Color `json:"color,omitempty"`
}

// ListBlockDefinitionsResult is the response of [MethodSketchBlockDefinitionList]:
// every block definition of the active part's component definition.
type ListBlockDefinitionsResult struct {
	Definitions []SketchBlockDefinitionInfo `json:"definitions"`
}

// DeleteBlockDefinitionArgs is the request of [MethodSketchBlockDefinitionDelete].
// A definition that still has instances cannot be deleted; the error names the
// block and the sketches consuming it.
type DeleteBlockDefinitionArgs struct {
	Name string `json:"name"`
}

// AddSketchBlockArgs is the request of [MethodSketchAddBlockInstance]: place an
// instance of the named definition in a sketch. Position is the insertion
// point [x,y] in sketch-plane database units (cm); RotationAngle is a
// unit-bearing angle ("30 deg", empty ⇒ 0); Scale is a uniform scale factor
// (0 ⇒ 1).
type AddSketchBlockArgs struct {
	SketchIndex   int       `json:"sketchIndex"`
	Definition    string    `json:"definition"`
	Position      []float64 `json:"position"`
	RotationAngle string    `json:"rotationAngle,omitempty"`
	Scale         float64   `json:"scale,omitempty"`
}

// AddSketchBlockResult is the response of [MethodSketchAddBlockInstance]: the
// created instance's session id.
type AddSketchBlockResult struct {
	EntityID uint64 `json:"entityId"`
}

// ExplodeSketchBlockArgs identifies one placed instance in its owning sketch.
// The host applies its placement to the independent geometry in one undoable
// edit. Failure must leave the sketch and instance unchanged.
// The instance color override is removed. Member overrides remain; a member
// without one retains the definition color when it has an override.
// Example: ExplodeSketchBlockArgs{SketchIndex: 0, InstanceID: 42}.
type ExplodeSketchBlockArgs struct {
	SketchIndex int    `json:"sketchIndex"`
	InstanceID  uint64 `json:"instanceId"`
}

// ExplodeSketchBlockResult identifies the newly independent sketch entities in
// creation order. The definition and other instances remain. Example:
// ExplodeSketchBlockResult{EntityRefs: []uint64{7, 8}}.
type ExplodeSketchBlockResult struct {
	EntityRefs []uint64 `json:"entityRefs"`
}

// SetBlockDefinitionColorArgs changes a definition override. Color nil restores
// inheritance. A non-nil color must have OverrideColorSource and finite opacity
// in [0,1]. Example: SetBlockDefinitionColorArgs{Name: "Rib", Color: nil}.
type SetBlockDefinitionColorArgs struct {
	Name  string       `json:"name"`
	Color *types.Color `json:"color"`
}

// SetBlockInstanceColorArgs changes one placed instance override. Color nil
// restores inheritance. An instance override draws all its members in that
// color; otherwise each member inherits from its definition if it has no
// member override. Example: SetBlockInstanceColorArgs{InstanceID: 42}.
type SetBlockInstanceColorArgs struct {
	SketchIndex int          `json:"sketchIndex"`
	InstanceID  uint64       `json:"instanceId"`
	Color       *types.Color `json:"color"`
}

// SketchBlockInfo is one row of [MethodSketchListBlockInstances]: a placed
// block instance — its session id, the definition it instances, its placement
// (insertion point in sketch cm, rotation in radians CCW, uniform scale), and
// the live entity count of its definition.
type SketchBlockInfo struct {
	Index       int          `json:"index"`
	ID          uint64       `json:"id"`
	Definition  string       `json:"definition"`
	Position    []float64    `json:"position"`
	Rotation    float64      `json:"rotation"`
	Scale       float64      `json:"scale"`
	EntityCount int          `json:"entityCount"`
	Color       *types.Color `json:"color,omitempty"`
}

// ListBlockInstancesResult is the response of [MethodSketchListBlockInstances].
type ListBlockInstancesResult struct {
	Instances []SketchBlockInfo `json:"instances"`
}
