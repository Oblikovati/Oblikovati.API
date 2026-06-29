// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// Parameter is the in-process contract for one parametric variable: the authored
// expression, its category, the value the model consumes, the unit category, the
// engineering tolerance, and the evaluation health. The GPL implementation's
// model/param.Parameter satisfies it (compile-time asserted there).
//
// The unit-bearing value is projected leanly: [Parameter.NominalValue] +
// [Parameter.UnitName] give value and dimensional category without dragging the
// host's dimensional-arithmetic Quantity type across the boundary. The JSON view of
// a parameter (expression + formatted value + health) is also available via
// [oblikovati.org/api/wire.ParameterInfo].
type Parameter interface {
	// Name is the display label.
	Name() string
	// Kind is the parameter category (model/user/reference/derived/table).
	Kind() types.ParameterKind
	// Expression is the authored expression source (e.g. "width * 2").
	Expression() string
	// NominalValue is the evaluated value before the tolerance is applied (database
	// units); [Parameter.ModelValue] applies the tolerance.
	NominalValue() float64
	// ModelValue is the value the model consumes after the tolerance is applied
	// (database units).
	ModelValue() float64
	// UnitName is the parameter's dimensional unit category (e.g. "Length", "Angle",
	// "Unitless") — the reference API's Parameter.Units. It is the category, not the
	// document's display unit; format a display string via the host where needed.
	UnitName() string
	// Tolerance is the engineering-tolerance band; the zero value is the standard
	// tolerance (see [types.Tolerance]).
	Tolerance() types.Tolerance
	// IsHealthy reports whether the parameter evaluated cleanly. The full status enum is a
	// host-internal concept (#1501); the contract exposes only this flag plus the reason.
	IsHealthy() bool
	// HealthReason is the human-readable reason when the parameter is not healthy ("" when
	// healthy) — a failed expression, a cycle, or a removed derived source.
	HealthReason() string
	// SetExpression replaces the expression; it errors for read-only kinds and for
	// malformed source (see [types.ParameterKind.Editable]).
	SetExpression(src string) error
}
