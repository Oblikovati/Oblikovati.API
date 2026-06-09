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
