// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The client-graphics object model (M16-F05, #641) is the retained-mode, in-process view of a
// client-graphics group: a tree of [GraphicsNode]s, each owning typed primitives that index
// into shared, anti-duplicated [GraphicsDataSets]. It composes onto the declarative bulk-group
// wire transport (the geometry travels as wire DTOs); these interfaces are the scalar surface
// in-process callers read.
//
// STATUS: forward-declared, NOT yet implemented. No host type satisfies these interfaces and
// there is no compile-time assertion for them — they describe the intended object model ahead
// of the host build (the retained-mode group travels and is read as wire DTOs today). The host
// keeps them on an explicit pending-implementation allowlist (archguard's
// pendingContractAssertions), guarded so this notice cannot silently go stale: an interface
// gains a real assertion here only when a host type implements it (#1613, audit B2).

// GraphicsCoordinateSet is a shared pool of vertex positions (cm). Primitives index into it so
// shared vertices are stored once.
type GraphicsCoordinateSet interface {
	Count() int
	At(i int) types.Point
}

// GraphicsColorSet is a shared pool of colors with a binding mode (per-vertex/per-item/overall).
type GraphicsColorSet interface {
	Count() int
	At(i int) types.Color
	Binding() types.ColorBindingEnum
}

// GraphicsIndexSet is a shared pool of 0-based indices into a coordinate/color/normal set.
type GraphicsIndexSet interface {
	Count() int
	At(i int) int
}

// GraphicsNormalSet is a shared pool of vertex normals with a binding mode.
type GraphicsNormalSet interface {
	Count() int
	At(i int) types.Vector
	Binding() types.NormalBindingEnum
}

// GraphicsTextureCoordinateSet is a shared pool of uv texture coordinates for image-mapped meshes.
type GraphicsTextureCoordinateSet interface {
	Count() int
	At(i int) (u, v float64)
}

// GraphicsImageSet is a shared pool of images referenced by image-billboard primitives.
type GraphicsImageSet interface {
	Count() int
	Path(i int) string
}

// GraphicsColorMapper maps scalar vertex values to colors (a heatmap legend). It is either
// inline on a primitive or registered by name and shared.
type GraphicsColorMapper interface {
	Name() string
	StopCount() int
	// ColorAt resolves the color for a scalar value, interpolating between stops and clamping.
	ColorAt(value float64) types.Color
}

// GraphicsDataSets are a node's shared, anti-duplicated data pools the primitives index into.
type GraphicsDataSets interface {
	Coordinates() GraphicsCoordinateSet
	Colors() GraphicsColorSet
	Indices() GraphicsIndexSet
	Normals() GraphicsNormalSet
	TextureCoordinates() GraphicsTextureCoordinateSet
	Images() GraphicsImageSet
}

// GraphicsPrimitive is the base scalar view shared by every typed primitive: its kind, the
// overall color, and its selectability/depth behavior.
type GraphicsPrimitive interface {
	// Kind is the primitive's topology/type (a types.GraphicsPrimitiveKind value).
	Kind() types.GraphicsPrimitiveKind
	// OverallColor is the primitive's broadcast color (when ColorBinding is overall).
	OverallColor() types.Color
	// Selectable reports whether the primitive participates in picking.
	Selectable() bool
	// OnTop reports whether the primitive draws through occluders (depth test off).
	OnTop() bool
}

// LineGraphics is an indexed line-list primitive.
type LineGraphics interface {
	GraphicsPrimitive
	// LineSpace is the space the line's width/pattern is defined in.
	LineSpace() types.LineDefinitionSpaceEnum
}

// LineStripGraphics is a connected-polyline primitive.
type LineStripGraphics interface{ LineGraphics }

// PointGraphics is a point-marker primitive drawn with a render style.
type PointGraphics interface {
	GraphicsPrimitive
	// RenderStyle is the glyph each point draws.
	RenderStyle() types.PointRenderStyleEnum
}

// TextGraphics is a world- or screen-anchored text label.
type TextGraphics interface {
	GraphicsPrimitive
	// Text is the label string.
	Text() string
	// Anchor is the label's world anchor point.
	Anchor() types.Point
	// Behavior is the view-relative behavior (billboarding / pixel scaling).
	Behavior() types.DisplayTransformBehaviorEnum
}

// TriangleGraphics is an indexed triangle-mesh primitive.
type TriangleGraphics interface {
	GraphicsPrimitive
	// Mapper is the color mapper resolving per-vertex colors from scalars, or nil.
	Mapper() GraphicsColorMapper
}

// TriangleStripGraphics is a triangle-strip primitive.
type TriangleStripGraphics interface{ TriangleGraphics }

// TriangleFanGraphics is a triangle-fan primitive.
type TriangleFanGraphics interface{ TriangleGraphics }

// SurfaceGraphics renders an existing B-rep body/face by reference key with an override
// color/transform — its geometry stays host-side.
type SurfaceGraphics interface {
	GraphicsPrimitive
	// BodyKey is the persistent reference key of the rendered body/face.
	BodyKey() string
}

// CurveGraphics renders an existing B-rep edge by reference key.
type CurveGraphics interface {
	GraphicsPrimitive
	// EdgeKey is the persistent reference key of the rendered edge.
	EdgeKey() string
}

// ImageGraphics is an image-billboard primitive (a textured quad).
type ImageGraphics interface {
	GraphicsPrimitive
	// ImagePath is the source image file.
	ImagePath() string
	// Anchor is the billboard's world anchor.
	Anchor() types.Point
	// Behavior is the view-relative behavior (front-facing / pixel scaling).
	Behavior() types.DisplayTransformBehaviorEnum
}

// SweepGraphics renders a profile swept along a path (a tube/ribbon overlay).
type SweepGraphics interface {
	GraphicsPrimitive
	// Radius is the swept profile radius (cm).
	Radius() float64
}

// GraphicsObjectNode is the object-model view of one graphics node: its transform, flags, the
// owner it is anchored to, the typed primitives it holds, and the data sets they index into.
type GraphicsObjectNode interface {
	// Id is the node's id within its group.
	Id() string
	// Transform places the node's geometry (identity when unset).
	Transform() types.Matrix
	// Visibility reports whether none/some/all of the node is visible.
	Visibility() types.GraphicsVisibilityEnum
	// Selectability reports whether none/some/all of the node is pickable.
	Selectability() types.GraphicsSelectabilityEnum
	// ComponentKey is the reference key of the component/feature this node is anchored to, or "".
	ComponentKey() string
	// DataSets are the shared data pools the primitives index into.
	DataSets() GraphicsDataSets
	// Primitives are the typed primitives the node holds.
	Primitives() []GraphicsPrimitive
}

// ComponentGraphics is a client-graphics collection owned by and transformed with a component
// (rather than the document root). It is the persistent, cached counterpart of the transient
// interaction graphics.
type ComponentGraphics interface {
	// ComponentKey is the reference key of the owning component.
	ComponentKey() string
	// Status is the freshness of the cached graphics (none/out-of-date/up-to-date).
	Status() types.CachedGraphicsStatusEnum
	// Nodes are the collection's object-model nodes.
	Nodes() []GraphicsObjectNode
}
