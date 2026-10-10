// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Client/interaction graphics DTOs (persistent client graphics + transient interaction graphics).
//
// The model is declarative bulk groups (not a chatty mutable object model): one
// [SetClientGraphicsArgs] submits or replaces a whole named group, its geometry shipped
// as flat arrays. The vocabulary is conventional — a group holds nodes, each node holds
// primitives, each primitive's geometry comes from
// coordinate/index/color/normal sets — but is one round-trip per group, the right shape
// for large simulation-result meshes over the wire.
//
// Flat-array conventions: Coordinates and Normals are xyz triples (len%3==0); Colors are
// rgba quads in 0..1 (len%4==0); Indices are 0-based into the primitive's own vertex set;
// Transform is a 16-element row-major 4x4 matrix.

// GraphicsColorMapper maps scalar vertex values to colors — the heatmap legend. Values is
// an ascending list of scalar stops; Colors is the rgba quad for each stop (len(Colors)
// == 4*len(Values)). The host interpolates piecewise-linearly between stops and clamps
// outside the range. A primitive carrying both Scalars and ColorMapper resolves per-vertex
// colors from them.
type GraphicsColorMapper struct {
	Values []float64 `json:"values"`
	Colors []float32 `json:"colors"`
}

// GraphicsPrimitive is one drawable primitive in a node (a line, point, triangle, or
// text primitive). Geometry travels inline as flat arrays.
//
//   - Kind is a [oblikovati.org/api/types.GraphicsPrimitiveKind] value.
//   - Coordinates/Indices define the vertices and topology.
//   - Per-vertex color is resolved in priority order: Colors (rgba quads) if present;
//     else Scalars+ColorMapper; else the overall Color broadcast to every vertex.
//   - ColorBinding/NormalBinding are the binding modes (types.GraphicsColorBinding/…).
//   - Text/Anchor/FontSize apply to the "text" kind (Anchor is the xyz world anchor).
//   - OnTop draws the primitive ignoring the depth test (always visible, drawn through
//     occluders); Opacity 0..1 (0 = use the lane/node default, i.e. opaque).
type GraphicsPrimitive struct {
	Kind          string               `json:"kind"`
	Coordinates   []float64            `json:"coordinates,omitempty"`
	Indices       []int                `json:"indices,omitempty"`
	Colors        []float32            `json:"colors,omitempty"`
	Normals       []float64            `json:"normals,omitempty"`
	Scalars       []float64            `json:"scalars,omitempty"`
	ColorMapper   *GraphicsColorMapper `json:"colorMapper,omitempty"`
	Color         []float32            `json:"color,omitempty"`
	ColorBinding  string               `json:"colorBinding,omitempty"`
	NormalBinding string               `json:"normalBinding,omitempty"`
	LineType      string               `json:"lineType,omitempty"`
	LineWeight    float64              `json:"lineWeight,omitempty"`
	PointStyle    string               `json:"pointStyle,omitempty"`
	PointSize     float64              `json:"pointSize,omitempty"`
	Text          string               `json:"text,omitempty"`
	Anchor        []float64            `json:"anchor,omitempty"`
	FontSize      float64              `json:"fontSize,omitempty"`
	Opacity       float32              `json:"opacity,omitempty"`
	OnTop         bool                 `json:"onTop,omitempty"`
	DepthPriority int                  `json:"depthPriority,omitempty"`

	// Body-derived primitives ("surface"/"curve"): render an existing B-rep body/face/edge by
	// its persistent reference key (BodyKey) or transient handle (TransientKey) with the override
	// Color/Transform — the geometry stays host-side (no mesh shipped over the wire).
	BodyKey      string `json:"bodyKey,omitempty"`
	TransientKey uint64 `json:"transientKey,omitempty"`

	// Image billboards ("image"): ImagePath is the source image; ImageWidth/ImageHeight size the
	// quad (model units, or pixels when Behavior is pixel-scaling); Anchor is its world point.
	ImagePath   string  `json:"imagePath,omitempty"`
	ImageWidth  float64 `json:"imageWidth,omitempty"`
	ImageHeight float64 `json:"imageHeight,omitempty"`

	// TextureCoords are uv pairs (len%2==0, one per vertex) for an image-mapped overlay mesh.
	TextureCoords []float64 `json:"textureCoords,omitempty"`

	// MapperName references a color mapper registered via [MethodClientGraphicsRegisterMapper]
	// (an alternative to the inline ColorMapper) so a legend is shared across primitives.
	MapperName string `json:"mapperName,omitempty"`

	// Selectable makes this primitive participate in picking (nil = inherit the node default);
	// Behavior is the view-relative behavior of an image/text billboard (front-facing / pixel
	// scaling); LineSpace is the coordinate space a line's width/pattern is defined in.
	Selectable *bool                              `json:"selectable,omitempty"`
	Behavior   types.DisplayTransformBehaviorEnum `json:"behavior,omitempty"`
	LineSpace  types.LineDefinitionSpaceEnum      `json:"lineSpace,omitempty"`
}

// GraphicsNode groups primitives under one transform and visibility/opacity.
// Transform (optional 16-element row-major matrix) places the node's
// geometry; Visible nil means "inherit the group's visibility".
type GraphicsNode struct {
	Id         string              `json:"id,omitempty"`
	Transform  []float64           `json:"transform,omitempty"`
	Visible    *bool               `json:"visible,omitempty"`
	Opacity    float32             `json:"opacity,omitempty"`
	Primitives []GraphicsPrimitive `json:"primitives"`

	// Selectable makes the node's primitives pickable (default false: overlay graphics do not
	// intercept picks). ComponentKey anchors the node to a component/feature by reference key so
	// it transforms with that owner rather than the document root (the ComponentGraphics case).
	Selectable   bool   `json:"selectable,omitempty"`
	ComponentKey string `json:"componentKey,omitempty"`
}

// SetClientGraphicsArgs is the request of [MethodClientGraphicsSet]: submit or replace
// the whole graphics group named ClientId. Lane selects the display lane (a
// types.GraphicsLane; empty = "persistent"). Visible nil means visible. Re-sending the
// same ClientId replaces the previous group (idempotent).
type SetClientGraphicsArgs struct {
	ClientId string         `json:"clientId"`
	Lane     string         `json:"lane,omitempty"`
	Visible  *bool          `json:"visible,omitempty"`
	Nodes    []GraphicsNode `json:"nodes"`
}

// SetClientGraphicsResult is the response of [MethodClientGraphicsSet].
type SetClientGraphicsResult struct {
	ClientId       string `json:"clientId"`
	NodeCount      int    `json:"nodeCount"`
	PrimitiveCount int    `json:"primitiveCount"`
}

// ClientGraphicsInfo is one row of [MethodClientGraphicsList].
type ClientGraphicsInfo struct {
	ClientId       string `json:"clientId"`
	Lane           string `json:"lane"`
	Visible        bool   `json:"visible"`
	NodeCount      int    `json:"nodeCount"`
	PrimitiveCount int    `json:"primitiveCount"`
}

// ListClientGraphicsResult is the response of [MethodClientGraphicsList].
type ListClientGraphicsResult struct {
	Groups []ClientGraphicsInfo `json:"groups"`
}

// DeleteClientGraphicsArgs is the request of [MethodClientGraphicsDelete].
type DeleteClientGraphicsArgs struct {
	ClientId string `json:"clientId"`
}

// SetClientGraphicsVisibleArgs is the request of [MethodClientGraphicsSetVisible].
type SetClientGraphicsVisibleArgs struct {
	ClientId string `json:"clientId"`
	Visible  bool   `json:"visible"`
}

// UpdateInteractionGraphicsArgs is the request of [MethodInteractionGraphicsUpdate]: it
// replaces the transient nodes of one interaction lane (overlay or preview) — the
// rubber-band/manipulator update path that runs on mouse move. Lane is a
// types.GraphicsLane ("overlay" or "preview").
type UpdateInteractionGraphicsArgs struct {
	Lane  string         `json:"lane"`
	Nodes []GraphicsNode `json:"nodes"`
}

// SetNodeTransformArgs is the request of [MethodGraphicsNodeSetTransform]: replace one node's
// transform without resubmitting its (possibly large) geometry. Transform is a 16-element
// row-major 4x4 matrix; an empty Transform resets to identity.
type SetNodeTransformArgs struct {
	ClientId  string    `json:"clientId"`
	NodeId    string    `json:"nodeId"`
	Transform []float64 `json:"transform,omitempty"`
}

// SetNodeVisibleArgs is the request of [MethodGraphicsNodeSetVisible]: toggle one node's
// visibility within a group without resubmitting geometry.
type SetNodeVisibleArgs struct {
	ClientId string `json:"clientId"`
	NodeId   string `json:"nodeId"`
	Visible  bool   `json:"visible"`
}

// SetNodeSelectableArgs is the request of [MethodGraphicsNodeSetSelectable]: toggle whether
// one node's primitives participate in picking.
type SetNodeSelectableArgs struct {
	ClientId   string `json:"clientId"`
	NodeId     string `json:"nodeId"`
	Selectable bool   `json:"selectable"`
}

// RegisterColorMapperArgs is the request of [MethodClientGraphicsRegisterMapper]: store a named,
// reusable color mapper that primitives reference by Name (via GraphicsPrimitive.MapperName)
// instead of carrying an inline copy of the legend.
type RegisterColorMapperArgs struct {
	Name   string              `json:"name"`
	Mapper GraphicsColorMapper `json:"mapper"`
}

// ColorMapperInfo is one entry of [ColorMappersResult]: a registered named mapper.
type ColorMapperInfo struct {
	Name      string `json:"name"`
	StopCount int    `json:"stopCount"`
}

// ColorMappersResult is the response of [MethodClientGraphicsListMappers].
type ColorMappersResult struct {
	Mappers []ColorMapperInfo `json:"mappers"`
}
