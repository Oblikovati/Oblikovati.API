// SPDX-License-Identifier: Apache-2.0

package wire

// ParameterInfo is the JSON shape of a parameter: its authored expression and its
// evaluated value formatted in the document's display units. Health is empty when
// the parameter is healthy.
type ParameterInfo struct {
	Name       string `json:"name"`
	Kind       string `json:"kind"`
	Expression string `json:"expression"`
	Value      string `json:"value"`
	Health     string `json:"health,omitempty"`
}

// ListParametersResult is the response of [MethodParametersList].
type ListParametersResult struct {
	Parameters []ParameterInfo `json:"parameters"`
}

// ParameterNameArgs identifies a parameter by name (request of [MethodParametersGet]).
type ParameterNameArgs struct {
	Name string `json:"name"`
}

// ParameterSetArgs is the request of [MethodParametersAdd] / [MethodParametersSet]:
// a parameter name and a unit-bearing expression (e.g. "4 cm").
type ParameterSetArgs struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}
