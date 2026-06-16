// SPDX-License-Identifier: Apache-2.0

package wire

// Flat-pattern orientation DTOs (M13-F05, Oblikovati/Oblikovati#635). An orientation is a
// saved alignment state of the developed flat: an alignment axis (a reference key, or empty
// for the part's natural axes) laid horizontal or vertical, an extra alignment rotation, and
// flips for the alignment direction and the base face. The active orientation frames the flat
// for drawing views and export and drives its reported length/width.

// FlatPatternOrientationInfo is one orientation, as reported by listOrientations and returned
// by addOrientation/activateOrientation. AlignmentRotation is in degrees; Length and Width are
// the flat's extents under this orientation (database units, cm). AlignmentAxis is a topology
// reference key, empty when the orientation uses the part's natural axes.
type FlatPatternOrientationInfo struct {
	Name              string  `json:"name"`
	AlignmentType     string  `json:"alignmentType"`
	AlignmentRotation float64 `json:"alignmentRotation"`
	AlignmentAxis     string  `json:"alignmentAxis,omitempty"`
	FlipAlignmentAxis bool    `json:"flipAlignmentAxis,omitempty"`
	FlipBaseFace      bool    `json:"flipBaseFace,omitempty"`
	Active            bool    `json:"active,omitempty"`
	Length            float64 `json:"length"`
	Width             float64 `json:"width"`
}

// OrientationsResult is the reply of listOrientations: every orientation in creation order
// (the active one flagged).
type OrientationsResult struct {
	Orientations []FlatPatternOrientationInfo `json:"orientations"`
}

// OrientationResult is the reply of addOrientation/activateOrientation: the affected
// orientation after the call.
type OrientationResult struct {
	Orientation FlatPatternOrientationInfo `json:"orientation"`
}

// AddOrientationArgs creates a named orientation. AlignmentType defaults to "horizontal"; an
// empty AlignmentAxis uses the part's natural axes; AlignmentRotation is in degrees.
type AddOrientationArgs struct {
	Name              string  `json:"name"`
	AlignmentType     string  `json:"alignmentType,omitempty"`
	AlignmentRotation float64 `json:"alignmentRotation,omitempty"`
	AlignmentAxis     string  `json:"alignmentAxis,omitempty"`
	FlipAlignmentAxis bool    `json:"flipAlignmentAxis,omitempty"`
	FlipBaseFace      bool    `json:"flipBaseFace,omitempty"`
	Activate          bool    `json:"activate,omitempty"`
}

// ActivateOrientationArgs / DeleteOrientationArgs name the target orientation.
type ActivateOrientationArgs struct {
	Name string `json:"name"`
}
type DeleteOrientationArgs struct {
	Name string `json:"name"`
}
