// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Parameters is the parameter operation group for the active document — a part OR an
// assembly (both are parameter holders).
type Parameters struct{ c *Client }

// Parameters returns the parameter operation group.
func (c *Client) Parameters() Parameters { return Parameters{c} }

// List returns the active document's parameters (part or assembly).
//
// mcp:tool list_parameters
// mcp:summary List the active document's parameters (expression + evaluated value).
func (p Parameters) List() (wire.ListParametersResult, error) {
	return call[wire.ListParametersResult](p.c, wire.MethodParametersList, nil)
}

// Get returns one parameter by name.
//
// mcp:tool get_parameter
// mcp:summary Get one parameter of the active document (part or assembly) by name.
func (p Parameters) Get(name string) (wire.ParameterInfo, error) {
	return call[wire.ParameterInfo](p.c, wire.MethodParametersGet, wire.ParameterNameArgs{Name: name})
}

// Add creates a new parameter. By default it is a numeric user parameter from a unit-bearing
// expression (e.g. name="height" expression="3 cm"). Set args.ValueType to "text" or "boolean"
// for a non-numeric value, and args.Kind to "model" for a model parameter (#1845).
//
// mcp:tool add_parameter
// mcp:summary Add a parameter: numeric (default), text, or boolean; user (default) or model kind.
func (p Parameters) Add(args wire.ParameterSetArgs) (wire.ParameterInfo, error) {
	return call[wire.ParameterInfo](p.c, wire.MethodParametersAdd, args)
}

// Rename changes a parameter's name, keeping its identity and rewriting every expression that
// references it (#1847).
//
// mcp:tool rename_parameter
// mcp:summary Rename a parameter, rewriting expressions that reference it.
func (p Parameters) Rename(name, newName string) (wire.ParameterInfo, error) {
	return call[wire.ParameterInfo](p.c, wire.MethodParametersRename, wire.ParameterRenameArgs{Name: name, NewName: newName})
}

// Convert changes a parameter's category (user/model/reference) in place, keeping its identity and
// dependency edges; converting to reference makes it read-only (#1850).
//
// mcp:tool convert_parameter
// mcp:summary Convert a parameter's kind (user/model/reference), preserving its name, expression and dependents. Converting to "reference" makes it read-only. A built-in/auto or derived parameter cannot be converted.
func (p Parameters) Convert(name, targetKind string) (wire.ParameterInfo, error) {
	return call[wire.ParameterInfo](p.c, wire.MethodParametersConvert, wire.ParameterConvertArgs{Name: name, TargetKind: targetKind})
}

// Set changes an existing parameter's expression and recomputes the model.
//
// mcp:tool set_parameter
// mcp:summary Change a parameter's expression and recompute the model.
func (p Parameters) Set(args wire.ParameterSetArgs) (wire.ParameterInfo, error) {
	return call[wire.ParameterInfo](p.c, wire.MethodParametersSet, args)
}

// GetDetail returns the full member-level view of one parameter: units,
// presentation, tolerance, expression list, custom-property exposure and the
// dependency neighborhood.
//
// mcp:tool parameters_get_detail
// mcp:summary Returns the full member-level view of one parameter: units, presentation, tolerance, expression list, custom-property exposure and the dependency neighborhood.
func (p Parameters) GetDetail(name string) (wire.ParameterDetail, error) {
	return call[wire.ParameterDetail](p.c, wire.MethodParametersGetDetail, wire.ParameterNameArgs{Name: name})
}

// Update applies the non-nil presentation/exposure mutations and returns the
// updated detail.
//
// mcp:tool parameters_update
// mcp:summary Applies the non-nil presentation/exposure mutations and returns the updated detail.
func (p Parameters) Update(args wire.ParameterUpdateArgs) (wire.ParameterDetail, error) {
	return call[wire.ParameterDetail](p.c, wire.MethodParametersUpdate, args)
}

// SetTolerance sets the parameter's engineering tolerance (see
// wire.ParameterToleranceArgs for the modes) and returns the updated detail.
//
// mcp:tool parameters_set_tolerance
// mcp:summary Sets the parameter's engineering tolerance (see wire.ParameterToleranceArgs for the modes) and returns the updated detail.
func (p Parameters) SetTolerance(args wire.ParameterToleranceArgs) (wire.ParameterDetail, error) {
	return call[wire.ParameterDetail](p.c, wire.MethodParametersSetTolerance, args)
}

// SetExpressionList replaces the parameter's multi-value choices (empty
// expressions clear the list) and returns the updated detail.
//
// mcp:tool parameters_set_expression_list
// mcp:summary Replaces the parameter's multi-value choices (empty expressions clear the list) and returns the updated detail.
func (p Parameters) SetExpressionList(args wire.ParameterExpressionListArgs) (wire.ParameterDetail, error) {
	return call[wire.ParameterDetail](p.c, wire.MethodParametersSetExpressionList, args)
}

// Delete removes a parameter by name. The host rejects the call (naming the
// offending dependents) while the parameter is in use.
//
// mcp:tool parameters_delete
// mcp:summary Removes a parameter by name.
func (p Parameters) Delete(name string) error {
	return p.c.call(wire.MethodParametersDelete, wire.ParameterNameArgs{Name: name}, nil)
}

// DrivenBy returns the names of the parameters this parameter's expression reads.
//
// mcp:tool parameters_driven_by
// mcp:summary Returns the names of the parameters this parameter's expression reads.
func (p Parameters) DrivenBy(name string) (wire.ParameterNamesResult, error) {
	return call[wire.ParameterNamesResult](p.c, wire.MethodParametersDrivenBy, wire.ParameterNameArgs{Name: name})
}

// Dependents returns the names of the parameters whose expressions read this one.
//
// mcp:tool parameters_dependents
// mcp:summary Returns the names of the parameters whose expressions read this one.
func (p Parameters) Dependents(name string) (wire.ParameterNamesResult, error) {
	return call[wire.ParameterNamesResult](p.c, wire.MethodParametersDependents, wire.ParameterNameArgs{Name: name})
}
