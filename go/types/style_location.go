// SPDX-License-Identifier: Apache-2.0

package types

// StyleLocationEnum is where a style lives in the cascade: in both the document and the style
// library, only locally in the document, or only in the library. A local style overrides the
// library style of the same name. The numeric ids are stable, frozen values (51201–51203).
//
// This is the canonical Apache-2.0 definition; the GPL implementation aliases it.
type StyleLocationEnum int32

const (
	// BothStyleLocation: the style exists in both the document and the library (51201).
	BothStyleLocation StyleLocationEnum = 51201
	// LocalStyleLocation: the style exists only locally in the document (51202).
	LocalStyleLocation StyleLocationEnum = 51202
	// LibraryStyleLocation: the style exists only in the style library (51203).
	LibraryStyleLocation StyleLocationEnum = 51203
)

var styleLocationNames = map[StyleLocationEnum]string{
	BothStyleLocation:    "Both",
	LocalStyleLocation:   "Local",
	LibraryStyleLocation: "Library",
}

// String returns the style-location's user-facing name.
func (s StyleLocationEnum) String() string {
	return enumName(styleLocationNames, s, "styleLocation(?)")
}

// IsValid reports whether s is a defined style location.
func (s StyleLocationEnum) IsValid() bool {
	return enumValid(styleLocationNames, s)
}

// AllStyleLocations returns every defined style location.
func AllStyleLocations() []StyleLocationEnum {
	return []StyleLocationEnum{BothStyleLocation, LocalStyleLocation, LibraryStyleLocation}
}
