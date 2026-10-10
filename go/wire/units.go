// SPDX-License-Identifier: Apache-2.0

package wire

// Document units of measure + the unit conversion / expression / formatting
// service (Oblikovati/Oblikovati#146). The document owns the display-unit
// preferences and precision; the units service parses, formats, converts, and
// validates unit-bearing values and expressions against them.

// DocumentUnitsInfo is the document's display-unit preferences and precision
// (response of [MethodDocumentsGetUnits], request body of partial updates via
// [SetDocumentUnitsArgs]). Unit fields carry unit-name spellings ("mm", "deg",
// "kg", "s"); the precisions are decimal places; LengthDisplayFormat is a
// types.ParameterDisplayFormat spelling ("decimal"/"fractional"/"architectural").
type DocumentUnitsInfo struct {
	LengthUnit             string `json:"lengthUnit"`
	AngleUnit              string `json:"angleUnit"`
	MassUnit               string `json:"massUnit"`
	TimeUnit               string `json:"timeUnit"`
	LengthDisplayPrecision int    `json:"lengthDisplayPrecision"`
	AngleDisplayPrecision  int    `json:"angleDisplayPrecision"`
	LengthDisplayFormat    string `json:"lengthDisplayFormat,omitempty"`
	// WorkingScaleCm is the centimetre size of one stored (working) length unit (ADR-0042
	// Phase 2). The kernel stores and reports geometry in working units; 1.0 means they are
	// centimetres (the default). A document centred on an extreme unit (µm/pm, km) reports a
	// different value, so an add-in reading raw geometry quantities multiplies by the
	// appropriate power of this to recover centimetres. Omitted (0) means the centimetre default.
	WorkingScaleCm float64 `json:"workingScaleCm,omitempty"`
}

// SetDocumentUnitsArgs is the request of [MethodDocumentsSetUnits]: only the
// non-nil fields are applied, so a caller can change one preference without
// restating the rest. The response is the updated [DocumentUnitsInfo].
type SetDocumentUnitsArgs struct {
	LengthUnit             *string `json:"lengthUnit,omitempty"`
	AngleUnit              *string `json:"angleUnit,omitempty"`
	MassUnit               *string `json:"massUnit,omitempty"`
	TimeUnit               *string `json:"timeUnit,omitempty"`
	LengthDisplayPrecision *int    `json:"lengthDisplayPrecision,omitempty"`
	AngleDisplayPrecision  *int    `json:"angleDisplayPrecision,omitempty"`
	LengthDisplayFormat    *string `json:"lengthDisplayFormat,omitempty"`
}

// ConvertUnitsArgs is the request of [MethodUnitsConvert]: a numeric value to
// convert From one unit name To another (both must name the same category).
type ConvertUnitsArgs struct {
	Value float64 `json:"value"`
	From  string  `json:"from"`
	To    string  `json:"to"`
}

// ConvertUnitsResult is the converted value (response of [MethodUnitsConvert]).
type ConvertUnitsResult struct {
	Value float64 `json:"value"`
}

// StringFromValueArgs is the request of [MethodUnitsGetStringFromValue] and
// [MethodUnitsGetPreciseStringFromValue]: a Value in database units of the
// given UnitsType (a types.UnitsType spelling), formatted in the document's
// display unit — honoring display precision (GetStringFromValue) or at full
// precision (GetPreciseStringFromValue).
type StringFromValueArgs struct {
	Value     float64 `json:"value"`
	UnitsType string  `json:"unitsType"`
}

// StringResult is a single formatted/normalized string (response of the
// value-formatting and locale/type-name methods).
type StringResult struct {
	Value string `json:"value"`
}

// ExpressionWithTypeArgs is the request of the expression methods that need a
// target category: [MethodUnitsGetValueFromExpression],
// [MethodUnitsIsExpressionValid], and [MethodUnitsCompatibleUnits]. UnitsType
// is a types.UnitsType spelling.
type ExpressionWithTypeArgs struct {
	Expression string `json:"expression"`
	UnitsType  string `json:"unitsType"`
}

// ValueResult is a single numeric value in database units (response of
// [MethodUnitsGetValueFromExpression]).
type ValueResult struct {
	Value float64 `json:"value"`
}

// ExpressionArgs is the request of the expression methods that auto-detect the
// unit: [MethodUnitsGetDatabaseUnitsFromExpression],
// [MethodUnitsGetLocaleCorrectedExpression], and
// [MethodUnitsGetDrivingParameters].
type ExpressionArgs struct {
	Expression string `json:"expression"`
}

// DatabaseUnitsResult is the evaluated database-unit value of an expression and
// the category it resolved to (response of
// [MethodUnitsGetDatabaseUnitsFromExpression]).
type DatabaseUnitsResult struct {
	Value     float64 `json:"value"`
	UnitsType string  `json:"unitsType"`
}

// ExpressionValidResult reports whether an expression is valid for the target
// category, with the parse/dimension error when not (response of
// [MethodUnitsIsExpressionValid]).
type ExpressionValidResult struct {
	Valid bool   `json:"valid"`
	Error string `json:"error,omitempty"`
}

// CompatibleUnitsResult reports whether an expression's resolved unit is
// dimensionally compatible with the target category (response of
// [MethodUnitsCompatibleUnits]).
type CompatibleUnitsResult struct {
	Compatible bool `json:"compatible"`
}

// UnitStringArgs is the request of [MethodUnitsGetTypeFromString]: a unit-name
// spelling ("mm") whose category is wanted.
type UnitStringArgs struct {
	UnitString string `json:"unitString"`
}

// UnitsTypeResult is a single category spelling (response of
// [MethodUnitsGetTypeFromString]).
type UnitsTypeResult struct {
	UnitsType string `json:"unitsType"`
}

// UnitsTypeArgs is the request of [MethodUnitsGetStringFromType]: a category
// whose document-preferred unit name is wanted.
type UnitsTypeArgs struct {
	UnitsType string `json:"unitsType"`
}

// DrivingParametersResult lists the parameter names an expression references
// (response of [MethodUnitsGetDrivingParameters]).
type DrivingParametersResult struct {
	Names []string `json:"names"`
}
