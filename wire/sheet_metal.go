// SPDX-License-Identifier: Apache-2.0

package wire

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
