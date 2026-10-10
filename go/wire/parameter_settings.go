// SPDX-License-Identifier: Apache-2.0

package wire

// Document-level parameter settings, tolerance sweeps, and parameter-set
// exchange (M02-F07, Oblikovati/Oblikovati#606).

// ParameterSettingsInfo is the JSON shape of the document's parameter
// settings (response of [MethodParametersGetSettings] and of
// [MethodParametersSetSettings]). The standard tolerances are unit-bearing
// expressions applied to dimensions without an explicit tolerance when
// UseStandardTolerances is on; the precisions are decimal places for linear
// and angular dimension display; DimensionDisplayType carries the wire
// spelling of types.DimensionDisplayType.
type ParameterSettingsInfo struct {
	LinearStandardTolerance      string `json:"linearStandardTolerance,omitempty"`
	AngularStandardTolerance     string `json:"angularStandardTolerance,omitempty"`
	UseStandardTolerances        bool   `json:"useStandardTolerances,omitempty"`
	ExportStandardTolerances     bool   `json:"exportStandardTolerances,omitempty"`
	LinearDimensionPrecision     int    `json:"linearDimensionPrecision"`
	AngularDimensionPrecision    int    `json:"angularDimensionPrecision"`
	DimensionDisplayType         string `json:"dimensionDisplayType"`
	DisplayParameterAsExpression bool   `json:"displayParameterAsExpression,omitempty"`
}

// ParameterSettingsUpdateArgs is the request of [MethodParametersSetSettings].
// Nil fields are left unchanged (pointer fields distinguish unchanged from
// zero); DimensionDisplayType takes a wire spelling.
type ParameterSettingsUpdateArgs struct {
	LinearStandardTolerance      *string `json:"linearStandardTolerance,omitempty"`
	AngularStandardTolerance     *string `json:"angularStandardTolerance,omitempty"`
	UseStandardTolerances        *bool   `json:"useStandardTolerances,omitempty"`
	ExportStandardTolerances     *bool   `json:"exportStandardTolerances,omitempty"`
	LinearDimensionPrecision     *int    `json:"linearDimensionPrecision,omitempty"`
	AngularDimensionPrecision    *int    `json:"angularDimensionPrecision,omitempty"`
	DimensionDisplayType         *string `json:"dimensionDisplayType,omitempty"`
	DisplayParameterAsExpression *bool   `json:"displayParameterAsExpression,omitempty"`
}

// ParameterSweepArgs is the request of [MethodParametersSetAllModelValueType]:
// drive every toleranced parameter's model-value selection to one bound
// (nominal/lower/upper/median wire spellings) for limit-stack studies.
type ParameterSweepArgs struct {
	ModelValueType string `json:"modelValueType"`
}

// ParameterSweepResult reports how many toleranced parameters the sweep moved.
type ParameterSweepResult struct {
	Affected int `json:"affected"`
}

// ParameterExportResult is the response of [MethodParametersExport]: the
// document's user parameters as the documented parameter-set XML.
type ParameterExportResult struct {
	XML string `json:"xml"`
}

// ParameterImportArgs is the request of [MethodParametersImport]: a
// parameter-set XML document. Import validates names, expressions, and units
// before touching the model and rejects the whole set naming the offending
// entry.
type ParameterImportArgs struct {
	XML string `json:"xml"`
}

// ParameterImportResult reports what the import did: parameters newly created
// and existing parameters whose expression or metadata changed.
type ParameterImportResult struct {
	Added   int `json:"added"`
	Updated int `json:"updated"`
}
