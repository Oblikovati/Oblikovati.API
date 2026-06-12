// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Parameters is the parameter operation group for the active part.
type Parameters struct{ c *Client }

// Parameters returns the parameter operation group.
func (c *Client) Parameters() Parameters { return Parameters{c} }

// List returns the active part's parameters.
func (p Parameters) List() (wire.ListParametersResult, error) {
	var r wire.ListParametersResult
	return r, p.c.call(wire.MethodParametersList, nil, &r)
}

// Get returns one parameter by name.
func (p Parameters) Get(name string) (wire.ParameterInfo, error) {
	var r wire.ParameterInfo
	return r, p.c.call(wire.MethodParametersGet, wire.ParameterNameArgs{Name: name}, &r)
}

// Add creates a new user parameter from a name and a unit-bearing expression.
func (p Parameters) Add(args wire.ParameterSetArgs) (wire.ParameterInfo, error) {
	var r wire.ParameterInfo
	return r, p.c.call(wire.MethodParametersAdd, args, &r)
}

// Set changes an existing parameter's expression and recomputes the model.
func (p Parameters) Set(args wire.ParameterSetArgs) (wire.ParameterInfo, error) {
	var r wire.ParameterInfo
	return r, p.c.call(wire.MethodParametersSet, args, &r)
}

// GetDetail returns the full member-level view of one parameter: units,
// presentation, tolerance, expression list, custom-property exposure and the
// dependency neighborhood.
func (p Parameters) GetDetail(name string) (wire.ParameterDetail, error) {
	var r wire.ParameterDetail
	return r, p.c.call(wire.MethodParametersGetDetail, wire.ParameterNameArgs{Name: name}, &r)
}

// Update applies the non-nil presentation/exposure mutations and returns the
// updated detail.
func (p Parameters) Update(args wire.ParameterUpdateArgs) (wire.ParameterDetail, error) {
	var r wire.ParameterDetail
	return r, p.c.call(wire.MethodParametersUpdate, args, &r)
}

// SetTolerance sets the parameter's engineering tolerance (see
// wire.ParameterToleranceArgs for the modes) and returns the updated detail.
func (p Parameters) SetTolerance(args wire.ParameterToleranceArgs) (wire.ParameterDetail, error) {
	var r wire.ParameterDetail
	return r, p.c.call(wire.MethodParametersSetTolerance, args, &r)
}

// SetExpressionList replaces the parameter's multi-value choices (empty
// expressions clear the list) and returns the updated detail.
func (p Parameters) SetExpressionList(args wire.ParameterExpressionListArgs) (wire.ParameterDetail, error) {
	var r wire.ParameterDetail
	return r, p.c.call(wire.MethodParametersSetExpressionList, args, &r)
}

// Delete removes a parameter by name. The host rejects the call (naming the
// offending dependents) while the parameter is in use.
func (p Parameters) Delete(name string) error {
	return p.c.call(wire.MethodParametersDelete, wire.ParameterNameArgs{Name: name}, nil)
}

// DrivenBy returns the names of the parameters this parameter's expression reads.
func (p Parameters) DrivenBy(name string) (wire.ParameterNamesResult, error) {
	var r wire.ParameterNamesResult
	return r, p.c.call(wire.MethodParametersDrivenBy, wire.ParameterNameArgs{Name: name}, &r)
}

// Dependents returns the names of the parameters whose expressions read this one.
func (p Parameters) Dependents(name string) (wire.ParameterNamesResult, error) {
	var r wire.ParameterNamesResult
	return r, p.c.call(wire.MethodParametersDependents, wire.ParameterNameArgs{Name: name}, &r)
}
