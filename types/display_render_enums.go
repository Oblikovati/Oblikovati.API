// SPDX-License-Identifier: Apache-2.0

package types

// DisplayModeSourceTypeEnum is whether a view's display mode comes from the application
// default or an explicit per-view override (the default-vs-override chain a new window
// resolves). Frozen ids 54273–54274.
type DisplayModeSourceTypeEnum int32

const (
	// DefaultDisplayModeSource takes the display mode from the application default (54273).
	DefaultDisplayModeSource DisplayModeSourceTypeEnum = 54273
	// OverrideDisplayModeSource takes the display mode from a per-view override (54274).
	OverrideDisplayModeSource DisplayModeSourceTypeEnum = 54274
)

var displayModeSourceNames = map[DisplayModeSourceTypeEnum]string{
	DefaultDisplayModeSource:  "Default",
	OverrideDisplayModeSource: "Override",
}

// String returns the display-mode-source's user-facing name.
func (d DisplayModeSourceTypeEnum) String() string {
	return enumName(displayModeSourceNames, d, "displayModeSource(?)")
}

// IsValid reports whether d is a defined display-mode source.
func (d DisplayModeSourceTypeEnum) IsValid() bool {
	return enumValid(displayModeSourceNames, d)
}

// AllDisplayModeSources returns every defined display-mode source.
func AllDisplayModeSources() []DisplayModeSourceTypeEnum {
	return []DisplayModeSourceTypeEnum{DefaultDisplayModeSource, OverrideDisplayModeSource}
}

// DisplaySeparateColorsTypeEnum is the granularity at which the "display separate colors"
// option assigns a distinct color: none, per unique top-level component, per top-level
// component, per unique part, or per part. Frozen ids 127489–127493.
type DisplaySeparateColorsTypeEnum int32

const (
	// NoneDisplaySeparateColors assigns no separate colors (127489).
	NoneDisplaySeparateColors DisplaySeparateColorsTypeEnum = 127489
	// EachUniqueTopLevelComponentColors colors each unique top-level component (127490).
	EachUniqueTopLevelComponentColors DisplaySeparateColorsTypeEnum = 127490
	// EachTopLevelComponentColors colors each top-level component instance (127491).
	EachTopLevelComponentColors DisplaySeparateColorsTypeEnum = 127491
	// EachUniquePartColors colors each unique part (127492).
	EachUniquePartColors DisplaySeparateColorsTypeEnum = 127492
	// EachPartColors colors each part instance (127493).
	EachPartColors DisplaySeparateColorsTypeEnum = 127493
)

var displaySeparateColorsNames = map[DisplaySeparateColorsTypeEnum]string{
	NoneDisplaySeparateColors:         "None",
	EachUniqueTopLevelComponentColors: "Each Unique Top-Level Component",
	EachTopLevelComponentColors:       "Each Top-Level Component",
	EachUniquePartColors:              "Each Unique Part",
	EachPartColors:                    "Each Part",
}

// String returns the display-separate-colors mode's user-facing name.
func (d DisplaySeparateColorsTypeEnum) String() string {
	return enumName(displaySeparateColorsNames, d, "displaySeparateColors(?)")
}

// IsValid reports whether d is a defined display-separate-colors mode.
func (d DisplaySeparateColorsTypeEnum) IsValid() bool {
	return enumValid(displaySeparateColorsNames, d)
}

// AllDisplaySeparateColors returns every defined display-separate-colors mode.
func AllDisplaySeparateColors() []DisplaySeparateColorsTypeEnum {
	return []DisplaySeparateColorsTypeEnum{
		NoneDisplaySeparateColors, EachUniqueTopLevelComponentColors,
		EachTopLevelComponentColors, EachUniquePartColors, EachPartColors,
	}
}

// BackFaceCullingEnum is which triangle winding the renderer culls: none, clockwise, or
// counter-clockwise. Frozen ids 96769–96771.
type BackFaceCullingEnum int32

const (
	// CullNone draws both faces of every triangle (96769).
	CullNone BackFaceCullingEnum = 96769
	// CullClockwise culls clockwise-wound faces (96770).
	CullClockwise BackFaceCullingEnum = 96770
	// CullCounterClockwise culls counter-clockwise-wound faces (96771).
	CullCounterClockwise BackFaceCullingEnum = 96771
)

var backFaceCullingNames = map[BackFaceCullingEnum]string{
	CullNone:             "None",
	CullClockwise:        "Clockwise",
	CullCounterClockwise: "Counter-Clockwise",
}

// String returns the back-face-culling mode's user-facing name.
func (b BackFaceCullingEnum) String() string {
	return enumName(backFaceCullingNames, b, "backFaceCulling(?)")
}

// IsValid reports whether b is a defined back-face-culling mode.
func (b BackFaceCullingEnum) IsValid() bool {
	return enumValid(backFaceCullingNames, b)
}

// AllBackFaceCullings returns every defined back-face-culling mode.
func AllBackFaceCullings() []BackFaceCullingEnum {
	return []BackFaceCullingEnum{CullNone, CullClockwise, CullCounterClockwise}
}

// TransparencyTypeEnum is how transparency is rendered: per-pixel alpha blending or a
// screen-door (stipple) approximation. Frozen ids 58625–58626.
type TransparencyTypeEnum int32

const (
	// BlendingTransparency renders transparency with alpha blending (58625).
	BlendingTransparency TransparencyTypeEnum = 58625
	// ScreenDoorTransparency renders transparency with a screen-door stipple (58626).
	ScreenDoorTransparency TransparencyTypeEnum = 58626
)

var transparencyTypeNames = map[TransparencyTypeEnum]string{
	BlendingTransparency:   "Blending",
	ScreenDoorTransparency: "Screen Door",
}

// String returns the transparency mode's user-facing name.
func (t TransparencyTypeEnum) String() string {
	return enumName(transparencyTypeNames, t, "transparencyType(?)")
}

// IsValid reports whether t is a defined transparency mode.
func (t TransparencyTypeEnum) IsValid() bool {
	return enumValid(transparencyTypeNames, t)
}

// AllTransparencyTypes returns every defined transparency mode.
func AllTransparencyTypes() []TransparencyTypeEnum {
	return []TransparencyTypeEnum{BlendingTransparency, ScreenDoorTransparency}
}

// RayTracingQualityEnum is the quality/speed tier of the ray-traced realistic display, from
// draft (fastest) to best (slowest). Frozen ids 95745–95750.
type RayTracingQualityEnum int32

const (
	// InteractiveRayTracingQuality is the interactive (fast preview) tier (95745).
	InteractiveRayTracingQuality RayTracingQualityEnum = 95745
	// GoodRayTracingQuality is the good tier (95746).
	GoodRayTracingQuality RayTracingQualityEnum = 95746
	// BestRayTracingQuality is the best (slowest) tier (95747).
	BestRayTracingQuality RayTracingQualityEnum = 95747
	// LowRayTracingQuality is the low tier (95748).
	LowRayTracingQuality RayTracingQualityEnum = 95748
	// DraftRayTracingQuality is the draft (fastest) tier (95749).
	DraftRayTracingQuality RayTracingQualityEnum = 95749
	// HighRayTracingQuality is the high tier (95750).
	HighRayTracingQuality RayTracingQualityEnum = 95750
)

var rayTracingQualityNames = map[RayTracingQualityEnum]string{
	InteractiveRayTracingQuality: "Interactive",
	GoodRayTracingQuality:        "Good",
	BestRayTracingQuality:        "Best",
	LowRayTracingQuality:         "Low",
	DraftRayTracingQuality:       "Draft",
	HighRayTracingQuality:        "High",
}

// String returns the ray-tracing-quality tier's user-facing name.
func (r RayTracingQualityEnum) String() string {
	return enumName(rayTracingQualityNames, r, "rayTracingQuality(?)")
}

// IsValid reports whether r is a defined ray-tracing-quality tier.
func (r RayTracingQualityEnum) IsValid() bool {
	return enumValid(rayTracingQualityNames, r)
}

// AllRayTracingQualities returns every defined ray-tracing-quality tier, draft→best.
func AllRayTracingQualities() []RayTracingQualityEnum {
	return []RayTracingQualityEnum{
		DraftRayTracingQuality, LowRayTracingQuality, InteractiveRayTracingQuality,
		GoodRayTracingQuality, HighRayTracingQuality, BestRayTracingQuality,
	}
}
