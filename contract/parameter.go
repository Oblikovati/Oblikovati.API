// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati/api/types"

// Parameter is the in-process contract for one parametric variable: the authored
// expression, its category, and the value the model consumes. The GPL
// implementation's model/param.Parameter satisfies it (compile-time asserted
// there).
//
// The evaluated [Quantity], unit, tolerance, and health are not on the contract
// yet — they return implementation value types that will move into
// [oblikovati/api/types] as the contract grows. The JSON view of a
// parameter (expression + formatted value + health) is available now via
// [oblikovati/api/wire.ParameterInfo].
type Parameter interface {
	// Name is the display label.
	Name() string
	// Kind is the parameter category (model/user/reference/derived/table).
	Kind() types.ParameterKind
	// Expression is the authored expression source (e.g. "width * 2").
	Expression() string
	// ModelValue is the value the model consumes after the tolerance is applied
	// (database units).
	ModelValue() float64
	// SetExpression replaces the expression; it errors for read-only kinds and for
	// malformed source (see [types.ParameterKind.Editable]).
	SetExpression(src string) error
}
