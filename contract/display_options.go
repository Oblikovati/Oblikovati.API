// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// GroundPlaneSettings is the in-process contract for a document's ground plane — the receiver
// the renderer drops shadows and reflections onto. Scoped to what the renderer consumes; the
// GPL app satisfies it (compile-time asserted there). Lengths are cm; Opacity/Reflectivity in
// [0,1].
type GroundPlaneSettings interface {
	// Visible reports whether the ground plane is drawn.
	Visible() bool
	// Color is the ground plane's color.
	Color() types.Color
	// HeightOffset is the plane's signed offset along its up direction (cm).
	HeightOffset() float64
	// DisplayGridLines reports whether the grid is drawn.
	DisplayGridLines() bool
	// MinorGridLineSpacing is the spacing between minor grid lines (cm).
	MinorGridLineSpacing() float64
	// MinorLinesPerMajorGridLine is how many minor lines fall between major lines.
	MinorLinesPerMajorGridLine() int
	// Opacity is the plane's opacity in [0,1].
	Opacity() float64
	// Reflectivity is the mirror reflection strength in [0,1].
	Reflectivity() float64
}

// ShadedDisplayOptions is the in-process contract for the shaded-mode display sub-options
// (edge overlay, transparency rendering). The GPL app satisfies it.
type ShadedDisplayOptions interface {
	// EdgeDisplay reports whether edges overlay the shaded faces.
	EdgeDisplay() bool
	// EdgeColor is the overlaid edge color.
	EdgeColor() types.Color
	// Silhouettes reports whether silhouette edges are drawn.
	Silhouettes() bool
	// TransparencyType is how transparency is rendered (blending / screen-door).
	TransparencyType() types.TransparencyTypeEnum
}

// WireframeDisplayOptions is the in-process contract for the wireframe-mode display
// sub-options. The GPL app satisfies it.
type WireframeDisplayOptions interface {
	// DepthDimming reports whether distant edges are dimmed.
	DepthDimming() bool
	// Silhouettes reports whether silhouette edges are drawn.
	Silhouettes() bool
	// DimmedHiddenEdges reports whether hidden edges are drawn dimmed (vs. removed).
	DimmedHiddenEdges() bool
}

// DisplayOptions is the in-process contract for the application-level display options — the
// preferences that parameterize the M23 display modes. Read-mostly here; mutation goes through
// the wire methods. The GPL app satisfies it.
type DisplayOptions interface {
	// DisplayQuality is the surface-display quality (binds to tessellation tolerance).
	DisplayQuality() types.DisplayQualityEnum
	// ViewTransitionTime is the animated camera-transition duration (seconds).
	ViewTransitionTime() float64
	// MinimumFrameRate is the interaction frame-rate floor (degrades quality to hold it).
	MinimumFrameRate() float64
	// HiddenLineDimmingPercent is the 0–100 dimming applied to hidden lines.
	HiddenLineDimmingPercent() int
	// EdgeColor is the global model-edge color.
	EdgeColor() types.Color
	// NewWindowDisplayMode is the display mode new views open in.
	NewWindowDisplayMode() types.DisplayModeEnum
	// NewWindowProjection is the projection new views open in.
	NewWindowProjection() types.ProjectionTypeEnum
	// BackFaceCulling is which triangle winding the renderer culls.
	BackFaceCulling() types.BackFaceCullingEnum
	// UseRayTracing reports whether realistic display uses ray tracing.
	UseRayTracing() bool
	// RayTracingQuality is the ray-tracing quality tier.
	RayTracingQuality() types.RayTracingQualityEnum
	// Shaded returns the shaded-mode sub-options.
	Shaded() ShadedDisplayOptions
	// Wireframe returns the wireframe-mode sub-options.
	Wireframe() WireframeDisplayOptions
}

// DisplaySettings is the in-process contract for a document's per-document display settings —
// the home for background, edge color, ground-plane, and shadow state. The GPL app satisfies
// it (compile-time asserted there).
type DisplaySettings interface {
	// BackgroundType is how the viewport background is painted (solid/gradient/image).
	BackgroundType() types.BackgroundTypeEnum
	// EdgeColor is the document's model-edge color override.
	EdgeColor() types.Color
	// DepthDimming reports whether distant geometry is dimmed.
	DepthDimming() bool
	// DisplaySilhouettes reports whether silhouette edges are drawn.
	DisplaySilhouettes() bool
	// HiddenLineDimmingPercent is the 0–100 dimming applied to hidden lines.
	HiddenLineDimmingPercent() int
	// NewWindowDisplayMode is the display mode new views of this document open in.
	NewWindowDisplayMode() types.DisplayModeEnum
	// DisplayModeSource is whether the display mode is the default or a per-view override.
	DisplayModeSource() types.DisplayModeSourceTypeEnum
	// NewWindowProjection is the projection new views of this document open in.
	NewWindowProjection() types.ProjectionTypeEnum
	// GroundPlane returns the document's ground-plane settings.
	GroundPlane() GroundPlaneSettings
	// GroundShadow is the ground-shadow style (none/standard/x-ray).
	GroundShadow() types.GroundShadowEnum
	// ShadowDirection is the light source shadows are cast from.
	ShadowDirection() types.ShadowDirectionEnum
	// ShowGroundReflections reports whether the ground plane mirrors the model.
	ShowGroundReflections() bool
	// ShowObjectShadows reports whether objects cast shadows on each other.
	ShowObjectShadows() bool
	// ShowAmbientShadows reports whether ambient-occlusion shadows are drawn.
	ShowAmbientShadows() bool
	// TexturesOn reports whether appearance textures are displayed.
	TexturesOn() bool
}
