// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

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

// EdgesOfTypeArgs filters the flat's classified edges to a single type ("bendUp", "bendDown"
// or "tangent"); an empty Type returns all classified edges.
type EdgesOfTypeArgs struct {
	Type string `json:"type,omitempty"`
}

// FlatEdgeInfo is one classified edge of the developed flat: the fold/tangent line segment
// (in flat 2D, database units cm) and its type, plus the bend angle for a fold line.
type FlatEdgeInfo struct {
	Start types.Point2d `json:"start"`
	End   types.Point2d `json:"end"`
	Type  string        `json:"type"`
	Angle float64       `json:"angle,omitempty"`
}

// EdgesResult is the reply of edgesOfType: the matching classified edges.
type EdgesResult struct {
	Edges []FlatEdgeInfo `json:"edges"`
}

// FlatFaceInfo is one classified flat face: its type ("front"/"back"/…) and developed area
// (cm²).
type FlatFaceInfo struct {
	Type string  `json:"type"`
	Area float64 `json:"area"`
}

// FacesResult is the reply of faces: the developed flat's classified faces (front and back).
type FacesResult struct {
	Faces []FlatFaceInfo `json:"faces"`
}

// MapEntityArgs maps one topology entity between the folded model and the developed flat by
// reference key. Key is a topology reference key as model.referenceKeys reports it (so a key
// can be fed straight back in); ToFlat maps folded→flat (false maps flat→folded). The mapping
// is face-level: a folded top/bottom face maps to the flat front/back face and back.
type MapEntityArgs struct {
	Key    string `json:"key"`
	ToFlat bool   `json:"toFlat,omitempty"`
}

// MapEntityResult is the reply of mapEntity: the corresponding entity's reference key, its
// kind ("face"), and whether a counterpart was found.
type MapEntityResult struct {
	Key   string `json:"key,omitempty"`
	Kind  string `json:"kind,omitempty"`
	Found bool   `json:"found"`
}

// PlateInfo is one developed flat plate — a connected flat region of a (possibly multi-body)
// sheet-metal part: its index, and its extents/area under the active orientation (database
// units cm; cm²).
type PlateInfo struct {
	Index  int     `json:"index"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Area   float64 `json:"area"`
}

// PlatesResult is the reply of listPlates: the developed flat's plates (one per connected
// region).
type PlatesResult struct {
	Plates []PlateInfo `json:"plates"`
}

// FlatPatternSettings is the per-document flat-pattern settings. DeferUpdate suppresses the
// automatic flat-pattern recompute, so a heavy flat is only developed on demand.
type FlatPatternSettings struct {
	DeferUpdate bool `json:"deferUpdate"`
}

// SettingsResult is the reply of getSettings/setSettings: the settings after the call.
type SettingsResult struct {
	Settings FlatPatternSettings `json:"settings"`
}

// SetSettingsArgs edits the flat-pattern settings.
type SetSettingsArgs struct {
	DeferUpdate bool `json:"deferUpdate,omitempty"`
}
