// SPDX-License-Identifier: Apache-2.0

package types

// Client/interaction graphics enums (persistent client graphics + transient interaction graphics).
// These string values are the stable wire vocabulary — treat them as frozen. They are
// the canonical, Apache-2.0 definitions; the GPL host aliases each with `type X = X`
// and maps them to its renderer.

// GraphicsPrimitiveKind names the topology of one graphics primitive (line, point,
// triangle, or text graphics). It is the discriminator
// of a [oblikovati.org/api/wire.GraphicsPrimitive].
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

// GraphicsColorBinding says how a primitive's colors map onto its geometry. With
// per-vertex each vertex carries its own color (the heatmap case); overall uses the
// single primitive color; per-item colors each line/triangle.
type GraphicsColorBinding string

const (
	GraphicsColorOverall   GraphicsColorBinding = "overall"
	GraphicsColorPerVertex GraphicsColorBinding = "perVertex"
	GraphicsColorPerItem   GraphicsColorBinding = "perItem"
)

// GraphicsNormalBinding says how normals map onto a triangle primitive. Per-vertex
// gives smooth shading; overall a single face normal.
type GraphicsNormalBinding string

const (
	GraphicsNormalOverall   GraphicsNormalBinding = "overall"
	GraphicsNormalPerVertex GraphicsNormalBinding = "perVertex"
)

// GraphicsLineType is a line primitive's dash pattern (a subset of the line-type styles).
type GraphicsLineType string

const (
	GraphicsLineContinuous GraphicsLineType = "continuous"
	GraphicsLineDashed     GraphicsLineType = "dashed"
	GraphicsLineDotted     GraphicsLineType = "dotted"
)

// GraphicsPointStyle is the glyph drawn for a points primitive (a subset of the point
// render styles). The host expands each point into screen-constant glyph geometry.
type GraphicsPointStyle string

const (
	GraphicsPointDot    GraphicsPointStyle = "dot"
	GraphicsPointCross  GraphicsPointStyle = "cross" // diagonal X
	GraphicsPointPlus   GraphicsPointStyle = "plus"  // axis-aligned +
	GraphicsPointSquare GraphicsPointStyle = "square"
	GraphicsPointCircle GraphicsPointStyle = "circle"
)

// GraphicsLane selects which display lane a graphics group lives in — the bridge between
// the two graphics surfaces. "persistent" is client graphics (document-owned, lives until
// deleted); "overlay" and "preview" are the interaction-graphics lanes (command-scoped,
// cleared when the interaction ends). Overlay graphics draw on top of the scene
// (depth-test off); preview graphics draw depth-tested with the scene.
type GraphicsLane string

const (
	GraphicsLanePersistent GraphicsLane = "persistent"
	GraphicsLaneOverlay    GraphicsLane = "overlay"
	GraphicsLanePreview    GraphicsLane = "preview"
)
