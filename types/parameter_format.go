// SPDX-License-Identifier: Apache-2.0

package types

// ParameterDisplayFormat is how a parameter's numeric value is rendered for
// display (parity: ParameterDisplayFormatEnum). It affects presentation only,
// never the stored or model value. The numeric values are frozen at the
// reference API's ids and must never be renumbered.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases
// it (ADR-0018).
type ParameterDisplayFormat int32

const (
	DisplayFormatDecimal       ParameterDisplayFormat = 92417
	DisplayFormatFractional    ParameterDisplayFormat = 92418
	DisplayFormatArchitectural ParameterDisplayFormat = 92419
)

// parameterDisplayFormatNames are the wire spellings (wire.ParameterDetail.DisplayFormat).
var parameterDisplayFormatNames = map[ParameterDisplayFormat]string{
	DisplayFormatDecimal:       "decimal",
	DisplayFormatFractional:    "fractional",
	DisplayFormatArchitectural: "architectural",
}

// String returns the display format's wire spelling.
func (f ParameterDisplayFormat) String() string {
	return enumName(parameterDisplayFormatNames, f, "enum(?)")
}

// ParseParameterDisplayFormat resolves a wire spelling back to its ParameterDisplayFormat.
func ParseParameterDisplayFormat(s string) (ParameterDisplayFormat, bool) {
	return enumFromName(parameterDisplayFormatNames, s)
}

// CustomPropertyType is whether a parameter exposed as a custom document
// property is published as text or as a number (parity: CustomPropertyTypeEnum).
// The numeric values are frozen at the reference API's ids.
type CustomPropertyType int32

const (
	CustomPropertyText   CustomPropertyType = 85249
	CustomPropertyNumber CustomPropertyType = 85250
)

// customPropertyTypeNames are the wire spellings (wire.CustomPropertyFormatInfo.PropertyType).
var customPropertyTypeNames = map[CustomPropertyType]string{
	CustomPropertyText:   "text",
	CustomPropertyNumber: "number",
}

// String returns the property type's wire spelling.
func (t CustomPropertyType) String() string { return enumName(customPropertyTypeNames, t, "enum(?)") }

// ParseCustomPropertyType resolves a wire spelling back to its CustomPropertyType.
func ParseCustomPropertyType(s string) (CustomPropertyType, bool) {
	return enumFromName(customPropertyTypeNames, s)
}

// CustomPropertyPrecision is the precision used when formatting a parameter
// value published as a custom document property (parity:
// CustomPropertyPrecisionEnum): decimal places, fractional denominators for
// lengths, or angular sexagesimal forms. The numeric values are frozen at the
// reference API's ids.
type CustomPropertyPrecision int32

const (
	PrecisionZeroDecimalPlace           CustomPropertyPrecision = 85505
	PrecisionOneDecimalPlace            CustomPropertyPrecision = 85506
	PrecisionTwoDecimalPlaces           CustomPropertyPrecision = 85507
	PrecisionThreeDecimalPlaces         CustomPropertyPrecision = 85508
	PrecisionFourDecimalPlaces          CustomPropertyPrecision = 85509
	PrecisionFiveDecimalPlaces          CustomPropertyPrecision = 85510
	PrecisionSixDecimalPlaces           CustomPropertyPrecision = 85511
	PrecisionSevenDecimalPlaces         CustomPropertyPrecision = 85512
	PrecisionEightDecimalPlaces         CustomPropertyPrecision = 85513
	PrecisionZeroFractional             CustomPropertyPrecision = 85514
	PrecisionHalfFractional             CustomPropertyPrecision = 85515
	PrecisionQuarterFractional          CustomPropertyPrecision = 85516
	PrecisionEighthsFractional          CustomPropertyPrecision = 85517
	PrecisionSixteenthsFractional       CustomPropertyPrecision = 85518
	PrecisionThirtySecondsFractional    CustomPropertyPrecision = 85519
	PrecisionSixtyFourthsFractional     CustomPropertyPrecision = 85520
	PrecisionOneTwentyEighthsFractional CustomPropertyPrecision = 85521
	PrecisionDegreesAngle               CustomPropertyPrecision = 85522
	PrecisionMinutesAngle               CustomPropertyPrecision = 85523
	PrecisionSecondsAngle               CustomPropertyPrecision = 85524
	PrecisionSecondsOneDecimalAngle     CustomPropertyPrecision = 85525
	PrecisionSecondsTwoDecimalAngle     CustomPropertyPrecision = 85526
	PrecisionSecondsThreeDecimalAngle   CustomPropertyPrecision = 85527
	PrecisionSecondsFourDecimalAngle    CustomPropertyPrecision = 85528
)

// customPropertyPrecisionNames are the wire spellings (wire.CustomPropertyFormatInfo.Precision).
var customPropertyPrecisionNames = map[CustomPropertyPrecision]string{
	PrecisionZeroDecimalPlace:           "zeroDecimalPlace",
	PrecisionOneDecimalPlace:            "oneDecimalPlace",
	PrecisionTwoDecimalPlaces:           "twoDecimalPlaces",
	PrecisionThreeDecimalPlaces:         "threeDecimalPlaces",
	PrecisionFourDecimalPlaces:          "fourDecimalPlaces",
	PrecisionFiveDecimalPlaces:          "fiveDecimalPlaces",
	PrecisionSixDecimalPlaces:           "sixDecimalPlaces",
	PrecisionSevenDecimalPlaces:         "sevenDecimalPlaces",
	PrecisionEightDecimalPlaces:         "eightDecimalPlaces",
	PrecisionZeroFractional:             "zeroFractional",
	PrecisionHalfFractional:             "halfFractional",
	PrecisionQuarterFractional:          "quarterFractional",
	PrecisionEighthsFractional:          "eighthsFractional",
	PrecisionSixteenthsFractional:       "sixteenthsFractional",
	PrecisionThirtySecondsFractional:    "thirtySecondsFractional",
	PrecisionSixtyFourthsFractional:     "sixtyFourthsFractional",
	PrecisionOneTwentyEighthsFractional: "oneTwentyEighthsFractional",
	PrecisionDegreesAngle:               "degrees",
	PrecisionMinutesAngle:               "minutes",
	PrecisionSecondsAngle:               "seconds",
	PrecisionSecondsOneDecimalAngle:     "secondsOneDecimalPlace",
	PrecisionSecondsTwoDecimalAngle:     "secondsTwoDecimalPlaces",
	PrecisionSecondsThreeDecimalAngle:   "secondsThreeDecimalPlaces",
	PrecisionSecondsFourDecimalAngle:    "secondsFourDecimalPlaces",
}

// String returns the precision's wire spelling.
func (p CustomPropertyPrecision) String() string {
	return enumName(customPropertyPrecisionNames, p, "enum(?)")
}

// ParseCustomPropertyPrecision resolves a wire spelling back to its CustomPropertyPrecision.
func ParseCustomPropertyPrecision(s string) (CustomPropertyPrecision, bool) {
	return enumFromName(customPropertyPrecisionNames, s)
}
