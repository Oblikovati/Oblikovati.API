// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Document-level parameter settings, tolerance sweeps, and parameter-set
// exchange (M02-F07, Oblikovati/Oblikovati#606), on the Parameters group.

// GetSettings returns the document's parameter settings.
func (p Parameters) GetSettings() (wire.ParameterSettingsInfo, error) {
	var r wire.ParameterSettingsInfo
	return r, p.c.call(wire.MethodParametersGetSettings, nil, &r)
}

// SetSettings applies the non-nil settings mutations and returns the updated
// settings.
func (p Parameters) SetSettings(args wire.ParameterSettingsUpdateArgs) (wire.ParameterSettingsInfo, error) {
	var r wire.ParameterSettingsInfo
	return r, p.c.call(wire.MethodParametersSetSettings, args, &r)
}

// SetAllModelValueType drives every toleranced parameter's model-value
// selection to one bound (nominal/lower/upper/median) in a single undo step,
// for limit-stack studies.
func (p Parameters) SetAllModelValueType(modelValueType string) (wire.ParameterSweepResult, error) {
	var r wire.ParameterSweepResult
	return r, p.c.call(wire.MethodParametersSetAllModelValueType, wire.ParameterSweepArgs{ModelValueType: modelValueType}, &r)
}

// Export returns the document's user parameters as the documented
// parameter-set XML for exchange with spreadsheet/PDM tooling.
func (p Parameters) Export() (wire.ParameterExportResult, error) {
	var r wire.ParameterExportResult
	return r, p.c.call(wire.MethodParametersExport, nil, &r)
}

// Import applies a parameter-set XML document: new names are created, known
// names updated. The whole set is validated first; a bad entry rejects the
// import naming the offending value.
func (p Parameters) Import(xml string) (wire.ParameterImportResult, error) {
	var r wire.ParameterImportResult
	return r, p.c.call(wire.MethodParametersImport, wire.ParameterImportArgs{XML: xml}, &r)
}
