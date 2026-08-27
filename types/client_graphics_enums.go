// SPDX-License-Identifier: Apache-2.0

package types

// Client-graphics object-model enums (the typed, frozen-id reference vocabulary that the
// retained-mode object model uses). These are the canonical Apache-2.0 definitions; the GPL
// host aliases each. They sit alongside the compact string enums in graphics.go: the strings
// are the bulk-group wire transport; these typed enums are the object-model identity.

// ClientGraphicsTypeEnum is whether a client-graphics collection is transient (rebuilt every
// interaction) or a preview. Frozen ids 45313–45314.
type ClientGraphicsTypeEnum int32

const (
	// TransientClientGraphics is rebuilt each interaction (45313).
	TransientClientGraphics ClientGraphicsTypeEnum = 45313
	// PreviewClientGraphics is a command preview (45314).
	PreviewClientGraphics ClientGraphicsTypeEnum = 45314
)

var clientGraphicsTypeNames = map[ClientGraphicsTypeEnum]string{
	TransientClientGraphics: "Transient", PreviewClientGraphics: "Preview",
}

// String returns the client-graphics-type's user-facing name.
func (c ClientGraphicsTypeEnum) String() string {
	return enumName(clientGraphicsTypeNames, c, "clientGraphicsType(?)")
}

// IsValid reports whether c is a defined client-graphics type.
func (c ClientGraphicsTypeEnum) IsValid() bool { return enumValid(clientGraphicsTypeNames, c) }

// AllClientGraphicsTypes returns every defined client-graphics type.
func AllClientGraphicsTypes() []ClientGraphicsTypeEnum {
	return []ClientGraphicsTypeEnum{TransientClientGraphics, PreviewClientGraphics}
}

// ColorBindingEnum is how a primitive's colors bind to its geometry: per-vertex, per-strip,
// per-item (per line/triangle), or one overall color. Frozen ids 19457–19460.
type ColorBindingEnum int32

const (
	// PerVertexColors binds one color per vertex (19457).
	PerVertexColors ColorBindingEnum = 19457
	// PerStripColors binds one color per strip (19458).
	PerStripColors ColorBindingEnum = 19458
	// PerItemColors binds one color per line/triangle (19459).
	PerItemColors ColorBindingEnum = 19459
	// OverallColor binds one color to the whole primitive (19460).
	OverallColor ColorBindingEnum = 19460
)

var colorBindingNames = map[ColorBindingEnum]string{
	PerVertexColors: "Per Vertex", PerStripColors: "Per Strip",
	PerItemColors: "Per Item", OverallColor: "Overall",
}

// String returns the color-binding's user-facing name.
func (c ColorBindingEnum) String() string {
	return enumName(colorBindingNames, c, "colorBinding(?)")
}

// IsValid reports whether c is a defined color binding.
func (c ColorBindingEnum) IsValid() bool { return enumValid(colorBindingNames, c) }

// AllColorBindings returns every defined color binding.
func AllColorBindings() []ColorBindingEnum {
	return []ColorBindingEnum{PerVertexColors, PerStripColors, PerItemColors, OverallColor}
}

// NormalBindingEnum is how a primitive's normals bind to its geometry: per-vertex (smooth),
// per-strip, per-item, or one overall normal. Frozen ids 19713–19716.
type NormalBindingEnum int32

const (
	// PerVertexNormals binds one normal per vertex — smooth shading (19713).
	PerVertexNormals NormalBindingEnum = 19713
	// PerStripNormals binds one normal per strip (19714).
	PerStripNormals NormalBindingEnum = 19714
	// PerItemNormals binds one normal per item (19715).
	PerItemNormals NormalBindingEnum = 19715
	// OverallNormal binds one normal to the whole primitive (19716).
	OverallNormal NormalBindingEnum = 19716
)

var normalBindingNames = map[NormalBindingEnum]string{
	PerVertexNormals: "Per Vertex", PerStripNormals: "Per Strip",
	PerItemNormals: "Per Item", OverallNormal: "Overall",
}

// String returns the normal-binding's user-facing name.
func (n NormalBindingEnum) String() string {
	return enumName(normalBindingNames, n, "normalBinding(?)")
}

// IsValid reports whether n is a defined normal binding.
func (n NormalBindingEnum) IsValid() bool { return enumValid(normalBindingNames, n) }

// AllNormalBindings returns every defined normal binding.
func AllNormalBindings() []NormalBindingEnum {
	return []NormalBindingEnum{PerVertexNormals, PerStripNormals, PerItemNormals, OverallNormal}
}

// GraphicsSelectabilityEnum is whether none, some, or all of a graphics node's geometry can be
// picked. Frozen ids 25345–25347.
type GraphicsSelectabilityEnum int32

const (
	// NoGraphicsSelectable makes none of the node selectable (25345).
	NoGraphicsSelectable GraphicsSelectabilityEnum = 25345
	// SomeGraphicsSelectable makes part of the node selectable (25346).
	SomeGraphicsSelectable GraphicsSelectabilityEnum = 25346
	// AllGraphicsSelectable makes all of the node selectable (25347).
	AllGraphicsSelectable GraphicsSelectabilityEnum = 25347
)

var graphicsSelectabilityNames = map[GraphicsSelectabilityEnum]string{
	NoGraphicsSelectable: "None", SomeGraphicsSelectable: "Some", AllGraphicsSelectable: "All",
}

// String returns the selectability's user-facing name.
func (g GraphicsSelectabilityEnum) String() string {
	return enumName(graphicsSelectabilityNames, g, "graphicsSelectability(?)")
}

// IsValid reports whether g is a defined selectability.
func (g GraphicsSelectabilityEnum) IsValid() bool { return enumValid(graphicsSelectabilityNames, g) }

// AllGraphicsSelectabilities returns every defined selectability.
func AllGraphicsSelectabilities() []GraphicsSelectabilityEnum {
	return []GraphicsSelectabilityEnum{NoGraphicsSelectable, SomeGraphicsSelectable, AllGraphicsSelectable}
}

// GraphicsVisibilityEnum is whether none, some, or all of a graphics node is visible. Frozen
// ids 25089–25091.
type GraphicsVisibilityEnum int32

const (
	// NoGraphicsVisible hides the whole node (25089).
	NoGraphicsVisible GraphicsVisibilityEnum = 25089
	// SomeGraphicsVisible shows part of the node (25090).
	SomeGraphicsVisible GraphicsVisibilityEnum = 25090
	// AllGraphicsVisible shows the whole node (25091).
	AllGraphicsVisible GraphicsVisibilityEnum = 25091
)

var graphicsVisibilityNames = map[GraphicsVisibilityEnum]string{
	NoGraphicsVisible: "None", SomeGraphicsVisible: "Some", AllGraphicsVisible: "All",
}

// String returns the visibility's user-facing name.
func (g GraphicsVisibilityEnum) String() string {
	return enumName(graphicsVisibilityNames, g, "graphicsVisibility(?)")
}

// IsValid reports whether g is a defined visibility.
func (g GraphicsVisibilityEnum) IsValid() bool { return enumValid(graphicsVisibilityNames, g) }

// AllGraphicsVisibilities returns every defined visibility.
func AllGraphicsVisibilities() []GraphicsVisibilityEnum {
	return []GraphicsVisibilityEnum{NoGraphicsVisible, SomeGraphicsVisible, AllGraphicsVisible}
}

// PointRenderStyleEnum is the glyph a point primitive draws (the full reference set; the
// compact [GraphicsPointStyle] strings are the bulk-transport subset). Frozen ids 20225–20235.
type PointRenderStyleEnum int32

const (
	// XPointStyle draws a diagonal X (20225).
	XPointStyle PointRenderStyleEnum = 20225
	// CirclePointStyle draws an open circle (20226).
	CirclePointStyle PointRenderStyleEnum = 20226
	// DimCircularPointStyle draws a dim circle (20227).
	DimCircularPointStyle PointRenderStyleEnum = 20227
	// FilledCirclePointStyle draws a filled circle (20228).
	FilledCirclePointStyle PointRenderStyleEnum = 20228
	// FilledCircleSelectPointStyle draws a filled circle for selection (20229).
	FilledCircleSelectPointStyle PointRenderStyleEnum = 20229
	// CrossPointStyle draws an axis-aligned cross (20230).
	CrossPointStyle PointRenderStyleEnum = 20230
	// FilledCrossPointStyle draws a filled cross (20231).
	FilledCrossPointStyle PointRenderStyleEnum = 20231
	// OnCurvePointStyle draws the on-curve glyph (20232).
	OnCurvePointStyle PointRenderStyleEnum = 20232
	// NoSnapPointStyle draws the no-snap glyph (20233).
	NoSnapPointStyle PointRenderStyleEnum = 20233
	// EndPointStyle draws the endpoint glyph (20234).
	EndPointStyle PointRenderStyleEnum = 20234
	// CustomImagePointStyle draws a custom image (20235).
	CustomImagePointStyle PointRenderStyleEnum = 20235
)

var pointRenderStyleNames = map[PointRenderStyleEnum]string{
	XPointStyle: "X", CirclePointStyle: "Circle", DimCircularPointStyle: "Dim Circular",
	FilledCirclePointStyle: "Filled Circle", FilledCircleSelectPointStyle: "Filled Circle Select",
	CrossPointStyle: "Cross", FilledCrossPointStyle: "Filled Cross", OnCurvePointStyle: "On Curve",
	NoSnapPointStyle: "No Snap", EndPointStyle: "End", CustomImagePointStyle: "Custom Image",
}

// String returns the point-render-style's user-facing name.
func (p PointRenderStyleEnum) String() string {
	return enumName(pointRenderStyleNames, p, "pointRenderStyle(?)")
}

// IsValid reports whether p is a defined point render style.
func (p PointRenderStyleEnum) IsValid() bool { return enumValid(pointRenderStyleNames, p) }

// AllPointRenderStyles returns every defined point render style.
func AllPointRenderStyles() []PointRenderStyleEnum {
	return []PointRenderStyleEnum{
		XPointStyle, CirclePointStyle, DimCircularPointStyle, FilledCirclePointStyle,
		FilledCircleSelectPointStyle, CrossPointStyle, FilledCrossPointStyle, OnCurvePointStyle,
		NoSnapPointStyle, EndPointStyle, CustomImagePointStyle,
	}
}

// CachedGraphicsStatusEnum is the freshness of a document's cached (persistent) client
// graphics: none, out-of-date (needs rebuild), or up-to-date. Frozen ids 103169–103171.
type CachedGraphicsStatusEnum int32

const (
	// NoneCachedGraphics means no cached graphics exist (103169).
	NoneCachedGraphics CachedGraphicsStatusEnum = 103169
	// OutOfDateCachedGraphics means the cache needs a rebuild (103170).
	OutOfDateCachedGraphics CachedGraphicsStatusEnum = 103170
	// UpToDateCachedGraphics means the cache is current (103171).
	UpToDateCachedGraphics CachedGraphicsStatusEnum = 103171
)

var cachedGraphicsStatusNames = map[CachedGraphicsStatusEnum]string{
	NoneCachedGraphics: "None", OutOfDateCachedGraphics: "Out of Date", UpToDateCachedGraphics: "Up to Date",
}

// String returns the cached-graphics-status's user-facing name.
func (c CachedGraphicsStatusEnum) String() string {
	return enumName(cachedGraphicsStatusNames, c, "cachedGraphicsStatus(?)")
}

// IsValid reports whether c is a defined cached-graphics status.
func (c CachedGraphicsStatusEnum) IsValid() bool { return enumValid(cachedGraphicsStatusNames, c) }

// AllCachedGraphicsStatuses returns every defined cached-graphics status.
func AllCachedGraphicsStatuses() []CachedGraphicsStatusEnum {
	return []CachedGraphicsStatusEnum{NoneCachedGraphics, OutOfDateCachedGraphics, UpToDateCachedGraphics}
}

// LineDefinitionSpaceEnum is the coordinate space a line graphic's width/pattern is defined
// in: screen space (constant pixels), model space (scales with zoom), or hybrid. Frozen ids
// 49921–49923.
type LineDefinitionSpaceEnum int32

const (
	// ScreenSpace defines line width in constant screen pixels (49921).
	ScreenSpace LineDefinitionSpaceEnum = 49921
	// ModelSpace defines line width in model units (scales with zoom) (49922).
	ModelSpace LineDefinitionSpaceEnum = 49922
	// HybridSpace mixes screen and model space (49923).
	HybridSpace LineDefinitionSpaceEnum = 49923
)

var lineDefinitionSpaceNames = map[LineDefinitionSpaceEnum]string{
	ScreenSpace: "Screen", ModelSpace: "Model", HybridSpace: "Hybrid",
}

// String returns the line-definition-space's user-facing name.
func (l LineDefinitionSpaceEnum) String() string {
	return enumName(lineDefinitionSpaceNames, l, "lineDefinitionSpace(?)")
}

// IsValid reports whether l is a defined line-definition space.
func (l LineDefinitionSpaceEnum) IsValid() bool { return enumValid(lineDefinitionSpaceNames, l) }

// AllLineDefinitionSpaces returns every defined line-definition space.
func AllLineDefinitionSpaces() []LineDefinitionSpaceEnum {
	return []LineDefinitionSpaceEnum{ScreenSpace, ModelSpace, HybridSpace}
}
