// SPDX-License-Identifier: Apache-2.0

package types

// Client/interaction graphics enums (Inventor's ClientGraphics + InteractionGraphics).
// These string values are the stable wire vocabulary — treat them as frozen. They are
// the canonical, Apache-2.0 definitions; the GPL host aliases each with `type X = X`
// and maps them to its renderer.

// GraphicsPrimitiveKind names the topology of one graphics primitive (Inventor's
// LineGraphics/PointGraphics/TriangleGraphics/TextGraphics …). It is the discriminator
// of a [github.com/Oblikovati/api/wire.GraphicsPrimitive].
type GraphicsPrimitiveKind string

const (
	// GraphicsLines is an indexed line list (segment pairs).
	GraphicsLines GraphicsPrimitiveKind = "lines"
	// GraphicsLineStrip is a connected polyline (each vertex joins the previous).
	GraphicsLineStrip GraphicsPrimitiveKind = "lineStrip"
	// GraphicsPoints is a set of point markers drawn with a [GraphicsPointStyle] glyph.
	GraphicsPoints GraphicsPrimitiveKind = "points"
	// GraphicsTriangles is an indexed triangle mesh (the heatmap/surface primitive).
	GraphicsTriangles GraphicsPrimitiveKind = "triangles"
	// GraphicsTriangleStrip is a triangle strip (each vertex extends the previous two).
	GraphicsTriangleStrip GraphicsPrimitiveKind = "triangleStrip"
	// GraphicsText is a world-anchored text label.
	GraphicsText GraphicsPrimitiveKind = "text"
)

// GraphicsColorBinding says how a primitive's colors map onto its geometry (Inventor's
// ColorBindingEnum). With kPerVertex each vertex carries its own color (the heatmap
// case); kOverall uses the single primitive color; kPerItem colors each line/triangle.
type GraphicsColorBinding string

const (
	GraphicsColorOverall   GraphicsColorBinding = "overall"
	GraphicsColorPerVertex GraphicsColorBinding = "perVertex"
	GraphicsColorPerItem   GraphicsColorBinding = "perItem"
)

// GraphicsNormalBinding says how normals map onto a triangle primitive (Inventor's
// NormalBindingEnum). kPerVertex gives smooth shading; kOverall a single face normal.
type GraphicsNormalBinding string

const (
	GraphicsNormalOverall   GraphicsNormalBinding = "overall"
	GraphicsNormalPerVertex GraphicsNormalBinding = "perVertex"
)

// GraphicsLineType is a line primitive's dash pattern (Inventor's LineTypeEnum, subset).
type GraphicsLineType string

const (
	GraphicsLineContinuous GraphicsLineType = "continuous"
	GraphicsLineDashed     GraphicsLineType = "dashed"
	GraphicsLineDotted     GraphicsLineType = "dotted"
)

// GraphicsPointStyle is the glyph drawn for a points primitive (Inventor's
// PointRenderStyleEnum, subset). The host expands each point into screen-constant glyph
// geometry.
type GraphicsPointStyle string

const (
	GraphicsPointDot    GraphicsPointStyle = "dot"
	GraphicsPointCross  GraphicsPointStyle = "cross" // diagonal X
	GraphicsPointPlus   GraphicsPointStyle = "plus"  // axis-aligned +
	GraphicsPointSquare GraphicsPointStyle = "square"
	GraphicsPointCircle GraphicsPointStyle = "circle"
)

// GraphicsLane selects which display lane a graphics group lives in — the bridge between
// Inventor's two surfaces. "persistent" is ClientGraphics (document-owned, lives until
// deleted); "overlay" and "preview" are the InteractionGraphics lanes (command-scoped,
// cleared when the interaction ends). Overlay graphics draw on top of the scene
// (depth-test off); preview graphics draw depth-tested with the scene.
type GraphicsLane string

const (
	GraphicsLanePersistent GraphicsLane = "persistent"
	GraphicsLaneOverlay    GraphicsLane = "overlay"
	GraphicsLanePreview    GraphicsLane = "preview"
)
