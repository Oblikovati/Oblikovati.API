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

// ToleranceInfo is the JSON shape of a parameter's engineering tolerance: its
// type spelling (types.ToleranceType.String()) and the deviation band from the
// nominal value, in database units.
type ToleranceInfo struct {
	Type  string  `json:"type"`
	Upper float64 `json:"upper,omitempty"`
	Lower float64 `json:"lower,omitempty"`
}

// ExpressionListInfo is the JSON shape of a parameter's multi-value choices.
// An absent/empty Expressions means the parameter is single-valued.
type ExpressionListInfo struct {
	Expressions       []string `json:"expressions,omitempty"`
	AllowCustomValues bool     `json:"allowCustomValues,omitempty"`
	CustomOrder       bool     `json:"customOrder,omitempty"`
}

// CustomPropertyFormatInfo is the JSON shape of how a parameter value is
// formatted when exposed as a custom document property. PropertyType and
// Precision carry the wire spellings of types.CustomPropertyType and
// types.CustomPropertyPrecision; Units empty means the document display unit.
type CustomPropertyFormatInfo struct {
	PropertyType      string `json:"propertyType"`
	Units             string `json:"units,omitempty"`
	Precision         string `json:"precision"`
	ShowLeadingZeros  bool   `json:"showLeadingZeros,omitempty"`
	ShowTrailingZeros bool   `json:"showTrailingZeros,omitempty"`
	ShowUnitsString   bool   `json:"showUnitsString,omitempty"`
}

// ParameterDetail is the full member-level view of one parameter (response of
// [MethodParametersGetDetail] and of the parameter mutation methods). It embeds
// [ParameterInfo], adding units, presentation, tolerance, the multi-value list,
// custom-property exposure and the dependency neighborhood. ModelValue is in
// database units; DisplayFormat and ModelValueType carry the wire spellings of
// types.ParameterDisplayFormat and types.ModelValueType. Tolerance is nil for
// text and true/false parameters.
type ParameterDetail struct {
	ParameterInfo
	Units                string                    `json:"units,omitempty"`
	Comment              string                    `json:"comment,omitempty"`
	IsKey                bool                      `json:"isKey,omitempty"`
	Visible              bool                      `json:"visible"`
	InUse                bool                      `json:"inUse,omitempty"`
	Precision            int                       `json:"precision"`
	DisplayFormat        string                    `json:"displayFormat"`
	ExposedAsProperty    bool                      `json:"exposedAsProperty,omitempty"`
	ModelValue           float64                   `json:"modelValue"`
	ModelValueType       string                    `json:"modelValueType"`
	Tolerance            *ToleranceInfo            `json:"tolerance,omitempty"`
	ExpressionList       *ExpressionListInfo       `json:"expressionList,omitempty"`
	CustomPropertyFormat *CustomPropertyFormatInfo `json:"customPropertyFormat,omitempty"`
	DrivenBy             []string                  `json:"drivenBy,omitempty"`
	Dependents           []string                  `json:"dependents,omitempty"`
}

// ParameterUpdateArgs is the request of [MethodParametersUpdate]: presentation
// and exposure mutations for one parameter. Nil fields are left unchanged.
// DisplayFormat and ModelValueType take wire spellings; CustomPropertyFormat
// replaces the whole format when present.
type ParameterUpdateArgs struct {
	Name                 string                    `json:"name"`
	Comment              *string                   `json:"comment,omitempty"`
	IsKey                *bool                     `json:"isKey,omitempty"`
	Visible              *bool                     `json:"visible,omitempty"`
	Precision            *int                      `json:"precision,omitempty"`
	DisplayFormat        *string                   `json:"displayFormat,omitempty"`
	ExposedAsProperty    *bool                     `json:"exposedAsProperty,omitempty"`
	ModelValueType       *string                   `json:"modelValueType,omitempty"`
	CustomPropertyFormat *CustomPropertyFormatInfo `json:"customPropertyFormat,omitempty"`
}

// ParameterToleranceArgs is the request of [MethodParametersSetTolerance].
// Mode is one of "default", "deviation", "symmetric", "limits", "min", "max".
// Upper/Lower are unit-bearing expressions in the parameter's unit (e.g.
// "0.1 mm"): deviation takes both as deviations from nominal, symmetric takes
// Upper as the ± band, limits takes both as absolute limit values, and the
// remaining modes take none.
type ParameterToleranceArgs struct {
	Name  string `json:"name"`
	Mode  string `json:"mode"`
	Upper string `json:"upper,omitempty"`
	Lower string `json:"lower,omitempty"`
}

// ParameterExpressionListArgs is the request of [MethodParametersSetExpressionList].
// An empty Expressions clears the list, returning the parameter to single-valued.
type ParameterExpressionListArgs struct {
	Name              string   `json:"name"`
	Expressions       []string `json:"expressions,omitempty"`
	AllowCustomValues bool     `json:"allowCustomValues,omitempty"`
	CustomOrder       bool     `json:"customOrder,omitempty"`
}

// ParameterNamesResult is the response of [MethodParametersDrivenBy] /
// [MethodParametersDependents]: the names of the directly linked parameters.
type ParameterNamesResult struct {
	Names []string `json:"names,omitempty"`
}
