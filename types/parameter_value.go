// SPDX-License-Identifier: Apache-2.0

package types

// Tolerance is the engineering-tolerance value type shared across the contract boundary
// (M39-F06, #1562). It is the canonical Apache-2.0 definition; the GPL implementation's
// model/param aliases it (ADR-0018), so the host's rich parameter type satisfies the
// contract.Parameter interface without re-declaring it.
//
// Parameter health is intentionally NOT defined here: the host's ParameterHealth enum is a
// distinct, source-internal concept (only its reason text crosses the boundary, never the
// enum — #1501), so the contract projects it leanly as IsHealthy()+HealthReason() rather than
// a shared health type.

// Tolerance is an engineering tolerance: a flavor plus the deviation band from the nominal
// value (Upper, Lower, database units). The zero Tolerance means "standard/default tolerance,
// no explicit band"; [Tolerance.Kind] maps the zero Type onto [ToleranceDefault], so a
// `t != Tolerance{}` has-explicit-tolerance check keeps working. Which value within the band
// the model consumes is the parameter's [ModelValueType] (the reference API splits
// Tolerance.ToleranceType from Parameter.ModelValueType).
type Tolerance struct {
	Type  ToleranceType
	Upper float64
	Lower float64
}

// Kind returns the tolerance flavor, mapping the zero value to ToleranceDefault.
func (t Tolerance) Kind() ToleranceType {
	if t.Type == 0 {
		return ToleranceDefault
	}
	return t.Type
}
