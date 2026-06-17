// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Document units of measure + the unit conversion / expression / formatting
// service (Oblikovati/Oblikovati#146). GetUnits/SetUnits read and edit the
// document's display preferences; the Units group parses, formats, converts,
// and validates unit-bearing values and expressions against them.

// Units is the unit conversion / expression / formatting service group.
type Units struct{ c *Client }

// Units returns the unit-of-measure service group.
func (c *Client) Units() Units { return Units{c} }

// GetUnits returns the active document's display-unit preferences and precision.
//
// mcp:tool get_document_units
// mcp:summary Returns the active document's display-unit preferences and precision.
func (d Documents) GetUnits() (wire.DocumentUnitsInfo, error) {
	var r wire.DocumentUnitsInfo
	return r, d.c.call(wire.MethodDocumentsGetUnits, nil, &r)
}

// SetUnits applies the non-nil unit/precision preferences and returns the
// updated units.
//
// mcp:tool set_document_units
// mcp:summary Applies the non-nil document unit/precision preferences and returns the updated units.
func (d Documents) SetUnits(args wire.SetDocumentUnitsArgs) (wire.DocumentUnitsInfo, error) {
	var r wire.DocumentUnitsInfo
	return r, d.c.call(wire.MethodDocumentsSetUnits, args, &r)
}

// Convert converts a value From one unit name To another within the same
// category (e.g. 25 "mm" → "in").
//
// mcp:tool units_convert
// mcp:summary Converts a value from one unit name to another within the same category.
func (u Units) Convert(args wire.ConvertUnitsArgs) (wire.ConvertUnitsResult, error) {
	var r wire.ConvertUnitsResult
	return r, u.c.call(wire.MethodUnitsConvert, args, &r)
}

// GetStringFromValue formats a database-unit value of the given category in the
// document's display unit, honoring the document's display precision.
//
// mcp:tool units_get_string_from_value
// mcp:summary Formats a database-unit value in the document display unit, honoring display precision.
func (u Units) GetStringFromValue(value float64, unitsType string) (wire.StringResult, error) {
	var r wire.StringResult
	return r, u.c.call(wire.MethodUnitsGetStringFromValue, wire.StringFromValueArgs{Value: value, UnitsType: unitsType}, &r)
}

// GetPreciseStringFromValue formats a database-unit value in the document's
// display unit at full precision (ignoring the display-precision rounding).
//
// mcp:tool units_get_precise_string_from_value
// mcp:summary Formats a database-unit value in the document display unit at full precision.
func (u Units) GetPreciseStringFromValue(value float64, unitsType string) (wire.StringResult, error) {
	var r wire.StringResult
	return r, u.c.call(wire.MethodUnitsGetPreciseStringFromValue, wire.StringFromValueArgs{Value: value, UnitsType: unitsType}, &r)
}

// GetValueFromExpression evaluates a unit-bearing expression to a database-unit
// value of the given target category.
//
// mcp:tool units_get_value_from_expression
// mcp:summary Evaluates a unit-bearing expression to a database-unit value of the given category.
func (u Units) GetValueFromExpression(expression, unitsType string) (wire.ValueResult, error) {
	var r wire.ValueResult
	return r, u.c.call(wire.MethodUnitsGetValueFromExpression, wire.ExpressionWithTypeArgs{Expression: expression, UnitsType: unitsType}, &r)
}

// GetDatabaseUnitsFromExpression evaluates an expression to a database-unit
// value, auto-detecting its category.
//
// mcp:tool units_get_database_units_from_expression
// mcp:summary Evaluates an expression to a database-unit value, auto-detecting its category.
func (u Units) GetDatabaseUnitsFromExpression(expression string) (wire.DatabaseUnitsResult, error) {
	var r wire.DatabaseUnitsResult
	return r, u.c.call(wire.MethodUnitsGetDatabaseUnitsFromExpression, wire.ExpressionArgs{Expression: expression}, &r)
}

// IsExpressionValid reports whether an expression parses and is dimensionally
// valid for the target category.
//
// mcp:tool units_is_expression_valid
// mcp:summary Reports whether an expression parses and is dimensionally valid for the target category.
func (u Units) IsExpressionValid(expression, unitsType string) (wire.ExpressionValidResult, error) {
	var r wire.ExpressionValidResult
	return r, u.c.call(wire.MethodUnitsIsExpressionValid, wire.ExpressionWithTypeArgs{Expression: expression, UnitsType: unitsType}, &r)
}

// CompatibleUnits reports whether an expression's resolved unit is
// dimensionally compatible with the target category.
//
// mcp:tool units_compatible_units
// mcp:summary Reports whether an expression's resolved unit is compatible with the target category.
func (u Units) CompatibleUnits(expression, unitsType string) (wire.CompatibleUnitsResult, error) {
	var r wire.CompatibleUnitsResult
	return r, u.c.call(wire.MethodUnitsCompatibleUnits, wire.ExpressionWithTypeArgs{Expression: expression, UnitsType: unitsType}, &r)
}

// GetTypeFromString returns the category a unit name belongs to (e.g. "mm" →
// "length").
//
// mcp:tool units_get_type_from_string
// mcp:summary Returns the category a unit name belongs to (e.g. "mm" → "length").
func (u Units) GetTypeFromString(unitString string) (wire.UnitsTypeResult, error) {
	var r wire.UnitsTypeResult
	return r, u.c.call(wire.MethodUnitsGetTypeFromString, wire.UnitStringArgs{UnitString: unitString}, &r)
}

// GetStringFromType returns the document-preferred unit name for a category
// (e.g. "length" → "mm").
//
// mcp:tool units_get_string_from_type
// mcp:summary Returns the document-preferred unit name for a category (e.g. "length" → "mm").
func (u Units) GetStringFromType(unitsType string) (wire.StringResult, error) {
	var r wire.StringResult
	return r, u.c.call(wire.MethodUnitsGetStringFromType, wire.UnitsTypeArgs{UnitsType: unitsType}, &r)
}

// GetLocaleCorrectedExpression normalizes an expression's number formatting
// (e.g. decimal separators) to the canonical form the evaluator accepts.
//
// mcp:tool units_get_locale_corrected_expression
// mcp:summary Normalizes an expression's number formatting to the canonical evaluator form.
func (u Units) GetLocaleCorrectedExpression(expression string) (wire.StringResult, error) {
	var r wire.StringResult
	return r, u.c.call(wire.MethodUnitsGetLocaleCorrectedExpression, wire.ExpressionArgs{Expression: expression}, &r)
}

// GetDrivingParameters returns the parameter names an expression references.
//
// mcp:tool units_get_driving_parameters
// mcp:summary Returns the parameter names an expression references.
func (u Units) GetDrivingParameters(expression string) (wire.DrivingParametersResult, error) {
	var r wire.DrivingParametersResult
	return r, u.c.call(wire.MethodUnitsGetDrivingParameters, wire.ExpressionArgs{Expression: expression}, &r)
}
