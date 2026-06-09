// SPDX-License-Identifier: Apache-2.0

package types

// TextHorizontalAlign is how sketch text is positioned horizontally about its anchor
// point (left/center/right, the standard MCAD horizontal text alignments). String values
// are frozen — they appear in the .obk document and on the wire.
type TextHorizontalAlign string

const (
	TextAlignLeft   TextHorizontalAlign = "left"
	TextAlignCenter TextHorizontalAlign = "center"
	TextAlignRight  TextHorizontalAlign = "right"
)

// TextVerticalAlign is how sketch text is positioned vertically about its anchor point
// (baseline/lower/middle/upper, the standard MCAD vertical text alignments). Baseline
// keeps the text's baseline on the anchor; the others measure from the text's cap box.
// String values are frozen.
type TextVerticalAlign string

const (
	TextAlignBaseline TextVerticalAlign = "baseline" // baseline on the anchor (sketch-text default)
	TextAlignLower    TextVerticalAlign = "lower"    // text below the anchor
	TextAlignMiddle   TextVerticalAlign = "middle"   // cap box centred on the anchor
	TextAlignUpper    TextVerticalAlign = "upper"    // text above the anchor
)

// SketchTextStyle is the renderable description of a sketch text entity: the content plus
// the type-setting parameters MCAD apps split across a text box + text style (font Family,
// FontSize in cm, the character Height in cm that scales the glyph em, the Rotation about
// the anchor in radians CCW, and horizontal/vertical alignment). It is the value carried
// by the wire DTOs and the typed client so an add-in can author/edit text without
// re-deriving the field set.
//
// Example: SketchTextStyle{Content: "PART A", Family: "Liberation Sans", FontSize: 0.5,
// Height: 0.5, HAlign: TextAlignCenter, VAlign: TextAlignBaseline}.
type SketchTextStyle struct {
	Content  string              `json:"content"`
	Family   string              `json:"family,omitempty"`   // font family ("" ⇒ document default)
	FontSize float64             `json:"fontSize,omitempty"` // cm; 0 ⇒ track Height
	Height   float64             `json:"height"`             // character (em) height in cm
	Rotation float64             `json:"rotation,omitempty"` // radians CCW about the anchor
	HAlign   TextHorizontalAlign `json:"hAlign,omitempty"`
	VAlign   TextVerticalAlign   `json:"vAlign,omitempty"`
}
