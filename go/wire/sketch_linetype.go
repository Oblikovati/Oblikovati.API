// SPDX-License-Identifier: Apache-2.0

package wire

// Custom sketch line types loaded from industry-standard .lin definition files
// (issue Oblikovati#161). A loaded definition switches the sketch's lineType
// override to "custom" and is persisted with the document, so the .lin file is
// only needed at load time.

// SetSketchCustomLineTypeArgs is the request of [MethodSketchSetCustomLineType]:
// load LineTypeName from the .lin file at FullFileName onto the sketch.
// ReplaceExisting replaces an already-loaded definition of the same name; when
// false, re-loading an existing name is an error.
type SetSketchCustomLineTypeArgs struct {
	SketchIndex     int    `json:"sketchIndex"`
	FullFileName    string `json:"fullFileName"`
	LineTypeName    string `json:"lineTypeName"`
	ReplaceExisting bool   `json:"replaceExisting,omitempty"`
}

// SketchCustomLineTypeResult is the response of [MethodSketchSetCustomLineType] and
// [MethodSketchGetCustomLineType]. Loaded reports whether the sketch holds a custom
// definition; Pattern is the .lin dash sequence (>0 dash, <0 gap, 0 dot), lengths
// in cm.
type SketchCustomLineTypeResult struct {
	Loaded       bool      `json:"loaded"`
	LineTypeName string    `json:"lineTypeName,omitempty"`
	FullFileName string    `json:"fullFileName,omitempty"`
	Pattern      []float64 `json:"pattern,omitempty"`
}
