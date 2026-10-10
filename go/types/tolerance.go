// SPDX-License-Identifier: Apache-2.0

package types

// ToleranceType is the engineering-tolerance flavor attached to a parameter or
// dimension (parity: ToleranceTypeEnum). The numeric values are frozen at the
// reference API's ids and must never be renumbered.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases
// it (ADR-0018).
type ToleranceType int32

const (
	ToleranceDefault                 ToleranceType = 31233
	ToleranceOverride                ToleranceType = 31234
	ToleranceSymmetric               ToleranceType = 31235
	ToleranceDeviation               ToleranceType = 31236
	ToleranceLimitsStacked           ToleranceType = 31237
	ToleranceLimitLinear             ToleranceType = 31238
	ToleranceMax                     ToleranceType = 31239
	ToleranceMin                     ToleranceType = 31240
	ToleranceLimitsFitsStacked       ToleranceType = 31241
	ToleranceLimitsFitsLinear        ToleranceType = 31242
	ToleranceLimitsFitsShowSize      ToleranceType = 31243
	ToleranceLimitsFitsShowTolerance ToleranceType = 31244
	ToleranceBasic                   ToleranceType = 31245
	ToleranceReference               ToleranceType = 31246
)

// toleranceTypeNames are the wire spellings (wire.ToleranceInfo.Type).
var toleranceTypeNames = map[ToleranceType]string{
	ToleranceDefault:                 "default",
	ToleranceOverride:                "override",
	ToleranceSymmetric:               "symmetric",
	ToleranceDeviation:               "deviation",
	ToleranceLimitsStacked:           "limitsStacked",
	ToleranceLimitLinear:             "limitLinear",
	ToleranceMax:                     "max",
	ToleranceMin:                     "min",
	ToleranceLimitsFitsStacked:       "limitsFitsStacked",
	ToleranceLimitsFitsLinear:        "limitsFitsLinear",
	ToleranceLimitsFitsShowSize:      "limitsFitsShowSize",
	ToleranceLimitsFitsShowTolerance: "limitsFitsShowTolerance",
	ToleranceBasic:                   "basic",
	ToleranceReference:               "reference",
}

// String returns the tolerance type's wire spelling.
func (t ToleranceType) String() string { return enumName(toleranceTypeNames, t, "enum(?)") }

// ParseToleranceType resolves a wire spelling back to its ToleranceType.
func ParseToleranceType(s string) (ToleranceType, bool) {
	return enumFromName(toleranceTypeNames, s)
}

// ModelValueType selects which value within a tolerance band the model actually
// consumes when recomputing (parity: ModelValueTypeEnum). The numeric values
// are frozen at the reference API's ids and must never be renumbered.
type ModelValueType int32

const (
	// ModelValueNominal uses the authored value, ignoring the tolerance band.
	ModelValueNominal ModelValueType = 31489
	// ModelValueLower uses nominal + the lower deviation.
	ModelValueLower ModelValueType = 31490
	// ModelValueUpper uses nominal + the upper deviation.
	ModelValueUpper ModelValueType = 31491
	// ModelValueMedian uses nominal + the midpoint of the deviation band.
	ModelValueMedian ModelValueType = 31492
)

// modelValueTypeNames are the wire spellings (wire.ParameterDetail.ModelValueType).
var modelValueTypeNames = map[ModelValueType]string{
	ModelValueNominal: "nominal",
	ModelValueLower:   "lower",
	ModelValueUpper:   "upper",
	ModelValueMedian:  "median",
}

// String returns the model-value type's wire spelling.
func (m ModelValueType) String() string { return enumName(modelValueTypeNames, m, "enum(?)") }

// ParseModelValueType resolves a wire spelling back to its ModelValueType.
func ParseModelValueType(s string) (ModelValueType, bool) {
	return enumFromName(modelValueTypeNames, s)
}

// enumName looks up an enum's wire spelling, returning fallback for values
// outside the map (e.g. "toleranceType(?)").
func enumName[E comparable](names map[E]string, v E, fallback string) string {
	if name, ok := names[v]; ok {
		return name
	}
	return fallback
}

// enumFromName is the inverse of [enumName]: it resolves a wire spelling to its
// enum value, reporting false for unknown spellings.
func enumFromName[E comparable](names map[E]string, s string) (E, bool) {
	for v, name := range names {
		if name == s {
			return v, true
		}
	}
	var zero E
	return zero, false
}

// enumValid reports whether v is a key of names — the shared body for every
// enum's IsValid method.
func enumValid[E comparable](names map[E]string, v E) bool {
	_, ok := names[v]
	return ok
}
