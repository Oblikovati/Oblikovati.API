// SPDX-License-Identifier: Apache-2.0

package types

// PointCloudDisplayMode is how an attached point cloud is coloured in the viewport
// (Oblikovati/Oblikovati#645). Unlike the numeric-id enums, this is a lowercase string token so it
// is stable in the on-the-wire JSON and in the persisted `.obk` document without a lookup table:
//
//	Default   — the neutral marker colour (the pre-#645 look).
//	RGB       — per-point colour decoded from the scan (LAS/PLY/E57), falling back to the marker
//	            colour where a point carries none.
//	Intensity — per-point greyscale mapped from the scan's intensity channel, normalised over the
//	            cloud's decoded range, falling back to the marker colour when the cloud has no
//	            usable intensity.
//
// This is the canonical Apache-2.0 definition; the GPL implementation aliases it
// (pointcloud.DisplayMode) and maps it onto the renderer.
type PointCloudDisplayMode string

const (
	// PointCloudDisplayModeDefault renders the cloud in the neutral marker colour (the pre-#645 look).
	PointCloudDisplayModeDefault PointCloudDisplayMode = "default"
	// PointCloudDisplayModeRGB renders per-point scan colour, marker colour where a point lacks it.
	PointCloudDisplayModeRGB PointCloudDisplayMode = "rgb"
	// PointCloudDisplayModeIntensity renders per-point greyscale from the scan's intensity channel.
	PointCloudDisplayModeIntensity PointCloudDisplayMode = "intensity"
)

var pointCloudDisplayModeNames = map[PointCloudDisplayMode]string{
	PointCloudDisplayModeDefault:   "Default",
	PointCloudDisplayModeRGB:       "RGB",
	PointCloudDisplayModeIntensity: "Intensity",
}

// String returns the mode's user-facing label (the display-mode selector label), distinct from the
// lowercase wire/persist token that is the underlying string value. The label lower-cases back to
// the token, so writing String() and reading it with strings.ToLower round-trips a stored mode.
func (m PointCloudDisplayMode) String() string {
	return enumName(pointCloudDisplayModeNames, m, "pointCloudDisplayMode(?)")
}

// IsValid reports whether m is one of the defined display modes.
func (m PointCloudDisplayMode) IsValid() bool {
	switch m {
	case PointCloudDisplayModeDefault, PointCloudDisplayModeRGB, PointCloudDisplayModeIntensity:
		return true
	default:
		return false
	}
}

// AllPointCloudDisplayModes returns every defined display mode — the source list for a display-mode
// picker and for the host's "expected one of …" validation message.
func AllPointCloudDisplayModes() []PointCloudDisplayMode {
	return []PointCloudDisplayMode{
		PointCloudDisplayModeDefault,
		PointCloudDisplayModeRGB,
		PointCloudDisplayModeIntensity,
	}
}
