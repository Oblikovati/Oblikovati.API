// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Document-level parameter settings, tolerance sweeps, and parameter-set
// exchange (M02-F07, Oblikovati/Oblikovati#606), on the Parameters group.

// GetSettings returns the document's parameter settings.
//
// mcp:tool parameters_get_settings
// mcp:summary Returns the document's parameter settings.
func (p Parameters) GetSettings() (wire.ParameterSettingsInfo, error) {
	return call[wire.ParameterSettingsInfo](p.c, wire.MethodParametersGetSettings, nil)
}

// SetSettings applies the non-nil settings mutations and returns the updated
// settings.
//
// mcp:tool parameters_set_settings
// mcp:summary Applies the non-nil settings mutations and returns the updated settings.
func (p Parameters) SetSettings(args wire.ParameterSettingsUpdateArgs) (wire.ParameterSettingsInfo, error) {
	return call[wire.ParameterSettingsInfo](p.c, wire.MethodParametersSetSettings, args)
}

// SetAllModelValueType drives every toleranced parameter's model-value
// selection to one bound (nominal/lower/upper/median) in a single undo step,
// for limit-stack studies.
//
// mcp:tool parameters_set_all_model_value_type
// mcp:summary Drives every toleranced parameter's model-value selection to one bound (nominal/lower/upper/median) in a single undo step, for limit-stack studies.
func (p Parameters) SetAllModelValueType(modelValueType string) (wire.ParameterSweepResult, error) {
	return call[wire.ParameterSweepResult](p.c, wire.MethodParametersSetAllModelValueType, wire.ParameterSweepArgs{ModelValueType: modelValueType})
}

// Export returns the document's user parameters as the documented
// parameter-set XML for exchange with spreadsheet/PDM tooling.
//
// mcp:tool parameters_export
// mcp:summary Returns the document's user parameters as the documented parameter-set XML for exchange with spreadsheet/PDM tooling.
func (p Parameters) Export() (wire.ParameterExportResult, error) {
	return call[wire.ParameterExportResult](p.c, wire.MethodParametersExport, nil)
}

// Import applies a parameter-set XML document: new names are created, known
// names updated. The whole set is validated first; a bad entry rejects the
// import naming the offending value.
//
// mcp:tool parameters_import
// mcp:summary Applies a parameter-set XML document: new names are created, known names updated.
func (p Parameters) Import(xml string) (wire.ParameterImportResult, error) {
	return call[wire.ParameterImportResult](p.c, wire.MethodParametersImport, wire.ParameterImportArgs{XML: xml})
}
