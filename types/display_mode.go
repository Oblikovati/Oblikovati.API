// SPDX-License-Identifier: Apache-2.0

package types

// DisplayModeEnum is the viewport's display mode — Inventor's DisplayModeEnum, the way a View
// of a document is drawn. The numeric ids are Inventor's exact values (8706–8716) and are
// frozen; two names share 8707 (Inventor aliases kHiddenEdgeRendering and
// kShadedWithHiddenEdgesRendering). This is the canonical Apache-2.0 definition; the GPL
// implementation aliases it (app.DisplayModeEnum) and maps it onto the renderer's VisualStyle.
type DisplayModeEnum int32

const (
	// WireframeRendering — every edge, no shaded faces, no hidden-line removal (8706).
	WireframeRendering DisplayModeEnum = 8706
	// HiddenEdgeRendering / ShadedWithHiddenEdgesRendering — shaded faces with visible edges
	// solid and occluded edges dashed (8707; the two names are aliases, per Inventor).
	HiddenEdgeRendering            DisplayModeEnum = 8707
	ShadedWithHiddenEdgesRendering DisplayModeEnum = 8707
	// ShadedRendering — lit faces, no edges (8708).
	ShadedRendering DisplayModeEnum = 8708
	// RealisticRendering — physically based (PBR) shading (8709).
	RealisticRendering DisplayModeEnum = 8709
	// ShadedWithEdgesRendering — lit faces with the edge wireframe (8710).
	ShadedWithEdgesRendering DisplayModeEnum = 8710
	// WireframeNoHiddenEdges — wireframe with the occluded edges removed (8711).
	WireframeNoHiddenEdges DisplayModeEnum = 8711
	// WireframeWithHiddenEdgesRendering — wireframe with occluded edges drawn dashed (8712).
	WireframeWithHiddenEdgesRendering DisplayModeEnum = 8712
	// MonochromeRendering — desaturated, posterized NPR (8713).
	MonochromeRendering DisplayModeEnum = 8713
	// WatercolorRendering — soft pigment washes on paper, NPR (8714).
	WatercolorRendering DisplayModeEnum = 8714
	// IllustrationRendering — flat/cel shading with outlines, NPR (8715).
	IllustrationRendering DisplayModeEnum = 8715
	// TechnicalIllustrationRendering — Gooch cool-warm shading with emphasized edges (8716).
	TechnicalIllustrationRendering DisplayModeEnum = 8716
)

var displayModeNames = map[DisplayModeEnum]string{
	WireframeRendering:                "Wireframe",
	HiddenEdgeRendering:               "Shaded with Hidden Edges",
	ShadedRendering:                   "Shaded",
	RealisticRendering:                "Realistic",
	ShadedWithEdgesRendering:          "Shaded with Edges",
	WireframeNoHiddenEdges:            "Wireframe with Visible Edges Only",
	WireframeWithHiddenEdgesRendering: "Wireframe with Hidden Edges",
	MonochromeRendering:               "Monochrome",
	WatercolorRendering:               "Watercolor",
	IllustrationRendering:             "Illustration",
	TechnicalIllustrationRendering:    "Technical Illustration",
}

// String returns the mode's stable, user-facing name (the Visual Style gallery label).
func (m DisplayModeEnum) String() string {
	if name, ok := displayModeNames[m]; ok {
		return name
	}
	return "displayMode(?)"
}

// IsValid reports whether m is a defined display mode.
func (m DisplayModeEnum) IsValid() bool {
	_, ok := displayModeNames[m]
	return ok
}

// AllDisplayModes returns every display mode in Inventor gallery order — the source list for a
// display-mode picker. The 8707 alias appears once (as the canonical 8707 entry).
func AllDisplayModes() []DisplayModeEnum {
	return []DisplayModeEnum{
		RealisticRendering, ShadedRendering, ShadedWithEdgesRendering, HiddenEdgeRendering,
		WireframeRendering, WireframeNoHiddenEdges, WireframeWithHiddenEdgesRendering,
		MonochromeRendering, WatercolorRendering, IllustrationRendering,
		TechnicalIllustrationRendering,
	}
}
