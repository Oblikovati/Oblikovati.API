// SPDX-License-Identifier: Apache-2.0

package types

// Drawing sheet value types (M14-F01). A drawing document holds a set of sheets;
// each sheet has either a standard size (ISO 216 A-series or ANSI/ASME Y14.1) whose
// dimensions come from the table below, or a custom width×height the sheet carries
// itself. These are the canonical, Apache-2.0 definitions; the GPL model aliases them
// (ADR-0018) and renders the sheet at the resulting size.

// SheetSize names a standard drawing sheet size. The zero value is SheetSizeCustom,
// whose dimensions come from the sheet's explicit width/height rather than the table.
//
// Values are part of the wire contract via [SheetSize.String]; the names — not the
// ordinals — are what cross the wire, so reordering is safe but renaming is not.
type SheetSize int32

const (
	// SheetSizeCustom takes its dimensions from the sheet's explicit width/height.
	SheetSizeCustom SheetSize = iota
	// SheetSizeA0 is ISO 216 A0 (841×1189 mm).
	SheetSizeA0
	// SheetSizeA1 is ISO 216 A1 (594×841 mm).
	SheetSizeA1
	// SheetSizeA2 is ISO 216 A2 (420×594 mm).
	SheetSizeA2
	// SheetSizeA3 is ISO 216 A3 (297×420 mm).
	SheetSizeA3
	// SheetSizeA4 is ISO 216 A4 (210×297 mm).
	SheetSizeA4
	// SheetSizeAnsiA is ANSI/ASME Y14.1 A (8.5×11 in).
	SheetSizeAnsiA
	// SheetSizeAnsiB is ANSI/ASME Y14.1 B (11×17 in).
	SheetSizeAnsiB
	// SheetSizeAnsiC is ANSI/ASME Y14.1 C (17×22 in).
	SheetSizeAnsiC
	// SheetSizeAnsiD is ANSI/ASME Y14.1 D (22×34 in).
	SheetSizeAnsiD
	// SheetSizeAnsiE is ANSI/ASME Y14.1 E (34×44 in).
	SheetSizeAnsiE
)

var sheetSizeNames = map[SheetSize]string{
	SheetSizeCustom: "custom",
	SheetSizeA0:     "a0",
	SheetSizeA1:     "a1",
	SheetSizeA2:     "a2",
	SheetSizeA3:     "a3",
	SheetSizeA4:     "a4",
	SheetSizeAnsiA:  "ansiA",
	SheetSizeAnsiB:  "ansiB",
	SheetSizeAnsiC:  "ansiC",
	SheetSizeAnsiD:  "ansiD",
	SheetSizeAnsiE:  "ansiE",
}

// String returns the sheet size's wire spelling ("a3", "ansiC", "custom").
func (s SheetSize) String() string { return enumName(sheetSizeNames, s, "enum(?)") }

// ParseSheetSize resolves a wire spelling back to its sheet size.
//
//	sz, ok := types.ParseSheetSize("a3") // SheetSizeA3, true
func ParseSheetSize(s string) (SheetSize, bool) { return enumFromName(sheetSizeNames, s) }

// sheetDimsMM holds each standard size's PORTRAIT dimensions (width ≤ height) in
// millimetres. ANSI sizes are their inch dimensions converted at 25.4 mm/in (e.g.
// A = 8.5×11 in → 215.9×279.4 mm). SheetSizeCustom is absent: its dimensions are not
// table-driven. Drawings are dimensioned in mm regardless of standard; the model
// converts to the kernel's centimetre unit when placing sheet geometry.
var sheetDimsMM = map[SheetSize][2]float64{
	SheetSizeA0:    {841, 1189},
	SheetSizeA1:    {594, 841},
	SheetSizeA2:    {420, 594},
	SheetSizeA3:    {297, 420},
	SheetSizeA4:    {210, 297},
	SheetSizeAnsiA: {215.9, 279.4},
	SheetSizeAnsiB: {279.4, 431.8},
	SheetSizeAnsiC: {431.8, 558.8},
	SheetSizeAnsiD: {558.8, 863.6},
	SheetSizeAnsiE: {863.6, 1117.6},
}

// SheetDimensionsMM returns a standard size's portrait width and height in
// millimetres (width ≤ height), and true. For SheetSizeCustom it returns (0, 0,
// false): a custom sheet's dimensions come from the sheet itself. Apply
// [SheetOrientation] to swap width and height for landscape.
//
//	w, h, _ := types.SheetDimensionsMM(types.SheetSizeA4) // 210, 297
func SheetDimensionsMM(s SheetSize) (width, height float64, ok bool) {
	d, ok := sheetDimsMM[s]
	if !ok {
		return 0, 0, false
	}
	return d[0], d[1], true
}

// SheetOrientation is how a sheet's standard dimensions are laid out. The zero value
// is SheetPortrait (width ≤ height); SheetLandscape swaps them.
type SheetOrientation int32

const (
	// SheetPortrait lays the sheet out with width ≤ height (the table's natural form).
	SheetPortrait SheetOrientation = iota
	// SheetLandscape swaps the standard width and height (the usual drafting layout).
	SheetLandscape
)

var sheetOrientationNames = map[SheetOrientation]string{
	SheetPortrait:  "portrait",
	SheetLandscape: "landscape",
}

// String returns the orientation's wire spelling ("portrait" / "landscape").
func (o SheetOrientation) String() string { return enumName(sheetOrientationNames, o, "enum(?)") }

// ParseSheetOrientation resolves a wire spelling back to its orientation.
func ParseSheetOrientation(s string) (SheetOrientation, bool) {
	return enumFromName(sheetOrientationNames, s)
}

// BorderLabelMode is how a zoned border labels one axis of zones (Inventor's BorderLabelModeEnum).
// The zero value is AlphabeticalBorderLabel (A, B, C …), the common default for the vertical axis.
type BorderLabelMode int32

const (
	// AlphabeticalBorderLabel labels zones A, B, C, … (the default).
	AlphabeticalBorderLabel BorderLabelMode = iota
	// NumericBorderLabel labels zones 1, 2, 3, ….
	NumericBorderLabel
	// NoBorderLabel draws the zone divisions without labels.
	NoBorderLabel
)

var borderLabelModeNames = map[BorderLabelMode]string{
	AlphabeticalBorderLabel: "alphabetical",
	NumericBorderLabel:      "numeric",
	NoBorderLabel:           "none",
}

// String returns the label mode's wire spelling.
func (m BorderLabelMode) String() string { return enumName(borderLabelModeNames, m, "enum(?)") }

// ParseBorderLabelMode resolves a wire spelling back to its label mode; "" ⇒ AlphabeticalBorderLabel.
func ParseBorderLabelMode(s string) (BorderLabelMode, bool) {
	if s == "" {
		return AlphabeticalBorderLabel, true
	}
	return enumFromName(borderLabelModeNames, s)
}

// TitleBlockLocation is the sheet corner a title block sits in (Inventor's TitleBlockLocationEnum).
// The zero value is BottomRightTitleBlock, the drafting default.
type TitleBlockLocation int32

const (
	// BottomRightTitleBlock places the title block in the lower-right corner (the default).
	BottomRightTitleBlock TitleBlockLocation = iota
	// BottomLeftTitleBlock places it in the lower-left corner.
	BottomLeftTitleBlock
	// TopLeftTitleBlock places it in the upper-left corner.
	TopLeftTitleBlock
	// TopRightTitleBlock places it in the upper-right corner.
	TopRightTitleBlock
)

var titleBlockLocationNames = map[TitleBlockLocation]string{
	BottomRightTitleBlock: "bottomRight",
	BottomLeftTitleBlock:  "bottomLeft",
	TopLeftTitleBlock:     "topLeft",
	TopRightTitleBlock:    "topRight",
}

// String returns the location's wire spelling.
func (l TitleBlockLocation) String() string { return enumName(titleBlockLocationNames, l, "enum(?)") }

// ParseTitleBlockLocation resolves a wire spelling back to its location; "" ⇒ BottomRightTitleBlock.
func ParseTitleBlockLocation(s string) (TitleBlockLocation, bool) {
	if s == "" {
		return BottomRightTitleBlock, true
	}
	return enumFromName(titleBlockLocationNames, s)
}
