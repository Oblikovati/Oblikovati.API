// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Sheet-metal rule/style DTOs (M13-F01). The active sheet-metal part carries a rule that
// fixes constant thickness, default bend radius, relief geometry and the unfold method; a
// flat pattern develops each bend using that method's bend allowance. These DTOs move the
// rule across the wire and let an add-in read the developed length of a single bend.

// SheetMetalStyleInfo is the active sheet-metal rule, as reported by getStyle and accepted
// (field-by-field, omitempty = leave unchanged) by setStyle. Lengths are unit expressions
// ("0.8 mm", "t*1.5") so they stay parameter-backed end to end; KFactor is the neutral-axis
// ratio used by the K-factor unfold method.
type SheetMetalStyleInfo struct {
	Name          string  `json:"name"`
	Thickness     string  `json:"thickness"`
	BendRadius    string  `json:"bendRadius"`
	ReliefShape   string  `json:"reliefShape"`
	ReliefWidth   string  `json:"reliefWidth"`
	ReliefDepth   string  `json:"reliefDepth"`
	MindGap       string  `json:"minimumGap"`
	UnfoldMethod  string  `json:"unfoldMethod"`
	KFactor       float64 `json:"kFactor"`
	BendAllowance float64 `json:"bendAllowance,omitempty"` // reported convenience; not an input
	// The CORNER relief is a separate property from the bend relief above: it is the cut made
	// where two flanges meet, with its own shape, size and placement, plus a distinct shape and
	// size for the three-bend corner (#1960). Inventor's Default style trims the corner to the
	// bend at four times the thickness, and rounds a three-bend corner at the bend radius.
	CornerReliefShape     string `json:"cornerReliefShape,omitempty"`
	CornerReliefSize      string `json:"cornerReliefSize,omitempty"`
	CornerReliefPlacement string `json:"cornerReliefPlacement,omitempty"`
	ThreeBendReliefShape  string `json:"threeBendReliefShape,omitempty"`
	ThreeBendReliefSize   string `json:"threeBendReliefSize,omitempty"`
}

// SheetMetalStyleResult is the reply of getStyle/setStyle: the active rule after the call.
type SheetMetalStyleResult struct {
	Style SheetMetalStyleInfo `json:"style"`
}

// SetSheetMetalStyleArgs edits the active rule. Every field is optional: an empty length
// expression (or zero KFactor) leaves that property unchanged, so a caller can nudge a
// single property without restating the whole rule.
type SetSheetMetalStyleArgs struct {
	Thickness    string  `json:"thickness,omitempty"`
	BendRadius   string  `json:"bendRadius,omitempty"`
	ReliefShape  string  `json:"reliefShape,omitempty"`
	ReliefWidth  string  `json:"reliefWidth,omitempty"`
	ReliefDepth  string  `json:"reliefDepth,omitempty"`
	MinimumGap   string  `json:"minimumGap,omitempty"`
	UnfoldMethod string  `json:"unfoldMethod,omitempty"`
	KFactor      float64 `json:"kFactor,omitempty"`
	// The corner-relief properties (#1960); empty leaves each unchanged, like the rest.
	CornerReliefShape     string `json:"cornerReliefShape,omitempty"`
	CornerReliefSize      string `json:"cornerReliefSize,omitempty"`
	CornerReliefPlacement string `json:"cornerReliefPlacement,omitempty"`
	ThreeBendReliefShape  string `json:"threeBendReliefShape,omitempty"`
	ThreeBendReliefSize   string `json:"threeBendReliefSize,omitempty"`
}

// BendAllowanceArgs requests the developed flat length of one bend under the active rule's
// unfold method. Radius defaults to the rule's bend radius when its expression is empty;
// Angle is the bend angle (the swept angle of the arc, e.g. "90 deg").
type BendAllowanceArgs struct {
	Angle  string `json:"angle"`
	Radius string `json:"radius,omitempty"`
}

// BendAllowanceResult reports the computed bend allowance (developed arc length, in the
// document's length unit) plus the bend deduction (setback) for the same bend.
type BendAllowanceResult struct {
	BendAllowance float64 `json:"bendAllowance"`
	BendDeduction float64 `json:"bendDeduction"`
}

// BendInfo is one bend in the part's bend lineage (M13-F04): the feature that introduced
// it and the unfold values the flat pattern develops it by. Angle is in degrees; the
// lengths (radius/thickness/allowance/deduction) are in database units (cm), matching the
// rest of the sheet-metal surface. The allowance is the developed neutral-axis arc length
// the flat must include; the deduction is the setback subtracted from outside flange
// lengths. Together they let an add-in predict the flat extents before the flat is built.
type BendInfo struct {
	Feature   string  `json:"feature"`
	Angle     float64 `json:"angle"`
	Radius    float64 `json:"radius"`
	Thickness float64 `json:"thickness"`
	Allowance float64 `json:"allowance"`
	Deduction float64 `json:"deduction"`
}

// BendsResult is the reply of bends: every bend in the folded part, in creation order, plus
// the summed bend allowance (the total developed length the flat adds for all bends).
type BendsResult struct {
	Bends          []BendInfo `json:"bends"`
	TotalAllowance float64    `json:"totalAllowance"`
}

// FlatBendLineInfo is one fold line in the flat pattern: the segment (in base-plane 2D, cm)
// and the bend angle in degrees — what a DXF export draws on the bend layer.
type FlatBendLineInfo struct {
	Start types.Point2d `json:"start"`
	End   types.Point2d `json:"end"`
	Angle float64       `json:"angle"`
}

// FlatPatternInfo is the developed flat: its 2D extents (the footprint bounding box in
// base-plane coordinates), the gauge, the developed footprint area, and the fold lines. All
// lengths are in database units (cm), matching the rest of the sheet-metal surface; areas in
// cm². It lets an add-in size stock and place bend lines before cutting.
type FlatPatternInfo struct {
	Extents   types.Box2d        `json:"extents"`
	Thickness float64            `json:"thickness"`
	Area      float64            `json:"area"`
	Bends     []FlatBendLineInfo `json:"bends"`
	// Punches are the punch instances developed into this flat (#1963), the same list
	// flatPattern.listPunches reports.
	Punches []FlatPunchInfo `json:"punches,omitempty"`
}

// UnfoldResult is the reply of unfold: the developed flat pattern of the active part.
type UnfoldResult struct {
	Flat FlatPatternInfo `json:"flat"`
}
