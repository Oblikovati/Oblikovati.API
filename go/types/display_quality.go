// SPDX-License-Identifier: Apache-2.0

package types

// DisplayQualityEnum is the surface-display quality the viewport renders at: it binds to the
// tessellation tolerance, trading smoothness for triangle count. The numeric ids are stable,
// frozen values (58881–58884).
//
// This is the canonical Apache-2.0 definition; the GPL implementation aliases it
// (app.DisplayQualityEnum).
type DisplayQualityEnum int32

const (
	// SmoothDisplayQuality is the smoothest (finest tessellation) quality (58881).
	SmoothDisplayQuality DisplayQualityEnum = 58881
	// MediumDisplayQuality is the medium quality (58882).
	MediumDisplayQuality DisplayQualityEnum = 58882
	// RoughDisplayQuality is the roughest (coarsest tessellation) quality (58883).
	RoughDisplayQuality DisplayQualityEnum = 58883
	// SmootherDisplayQuality is finer than Smooth, the highest quality (58884).
	SmootherDisplayQuality DisplayQualityEnum = 58884
)

var displayQualityNames = map[DisplayQualityEnum]string{
	SmoothDisplayQuality:   "Smooth",
	MediumDisplayQuality:   "Medium",
	RoughDisplayQuality:    "Rough",
	SmootherDisplayQuality: "Smoother",
}

// String returns the display-quality's user-facing name.
func (d DisplayQualityEnum) String() string {
	return enumName(displayQualityNames, d, "displayQuality(?)")
}

// IsValid reports whether d is a defined display quality.
func (d DisplayQualityEnum) IsValid() bool {
	return enumValid(displayQualityNames, d)
}

// AllDisplayQualities returns every defined display quality, in picker order.
func AllDisplayQualities() []DisplayQualityEnum {
	return []DisplayQualityEnum{RoughDisplayQuality, MediumDisplayQuality, SmoothDisplayQuality, SmootherDisplayQuality}
}
