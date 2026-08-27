// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"
	"fmt"
)

// The core utility value shapes (M00-F05, #601). Variant is the typed
// name-value-bag currency of the options/events surfaces: implemented HERE, in
// the Apache-2.0 contract module, because (like the transient geometry) it is
// ownerless immutable pure data an add-in builds locally at zero wire cost
// (ADR-0018 addendum). The wire bag shapes live in api/wire and embed it.

// ValueType is the type tag of a [Variant]. Numeric values are FROZEN to the
// reference ValueTypeEnum — clients and saved automations depend on them.
type ValueType int32

const (
	IntegerValue   ValueType = 14593
	DoubleValue    ValueType = 14594
	StringValue    ValueType = 14595
	ByteArrayValue ValueType = 14596
	BooleanValue   ValueType = 14597
)

var valueTypeNames = map[ValueType]string{
	IntegerValue: "integer", DoubleValue: "double", StringValue: "string",
	ByteArrayValue: "bytes", BooleanValue: "boolean",
}

// String returns the tag's stable name — also its wire spelling.
func (t ValueType) String() string {
	return enumName(valueTypeNames, t, "valueType(?)")
}

// Variant is a typed bag value: exactly one underlying value, selected by its
// [ValueType] tag. The typed accessors return ok=false on a tag mismatch so a
// wrong read is never silent. A double may carry a unit expression (the
// unit-bearing attribute shape — the full unit-display surface is #146).
//
//	opts := types.StringVariant("color"), types.UnitVariant(2.5, "mm")
type Variant struct {
	typ  ValueType
	b    bool
	i    int64
	f    float64
	s    string
	raw  []byte
	unit string
}

// IntegerVariant wraps an integer value.
func IntegerVariant(v int64) Variant { return Variant{typ: IntegerValue, i: v} }

// DoubleVariant wraps a floating-point value.
func DoubleVariant(v float64) Variant { return Variant{typ: DoubleValue, f: v} }

// UnitVariant wraps a floating-point value carrying a unit expression ("mm",
// "deg", …) — the unit-bearing bag value used by unit-aware options.
func UnitVariant(v float64, unit string) Variant {
	return Variant{typ: DoubleValue, f: v, unit: unit}
}

// StringVariant wraps a string value.
func StringVariant(v string) Variant { return Variant{typ: StringValue, s: v} }

// BoolVariant wraps a boolean value.
func BoolVariant(v bool) Variant { return Variant{typ: BooleanValue, b: v} }

// BytesVariant wraps an opaque byte-array value (base64 on the wire).
func BytesVariant(v []byte) Variant { return Variant{typ: ByteArrayValue, raw: v} }

// Type returns the tag selecting which accessor is meaningful.
func (v Variant) Type() ValueType { return v.typ }

// Integer returns the integer value, ok=false when v holds another type.
func (v Variant) Integer() (int64, bool) { return v.i, v.typ == IntegerValue }

// Double returns the floating-point value, ok=false when v holds another type.
func (v Variant) Double() (float64, bool) { return v.f, v.typ == DoubleValue }

// Str returns the string value, ok=false when v holds another type.
func (v Variant) Str() (string, bool) { return v.s, v.typ == StringValue }

// Bool returns the boolean value, ok=false when v holds another type.
func (v Variant) Bool() (bool, bool) { return v.b, v.typ == BooleanValue }

// Bytes returns the byte-array value, ok=false when v holds another type.
func (v Variant) Bytes() ([]byte, bool) { return v.raw, v.typ == ByteArrayValue }

// Unit returns the unit expression of a unit-bearing double ("" when none).
func (v Variant) Unit() string { return v.unit }

// variantJSON is the canonical wire shape: a stable type-tag name, the value
// under that tag, and the optional unit expression.
type variantJSON struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
	Unit  string          `json:"unit,omitempty"`
}

// MarshalJSON encodes the canonical bag-value shape, e.g.
// {"type":"double","value":2.5,"unit":"mm"}.
func (v Variant) MarshalJSON() ([]byte, error) {
	payload, err := json.Marshal(v.payload())
	if err != nil {
		return nil, err
	}
	return json.Marshal(variantJSON{Type: v.typ.String(), Value: payload, Unit: v.unit})
}

// payload returns the Go value selected by the tag.
func (v Variant) payload() interface{} {
	switch v.typ {
	case IntegerValue:
		return v.i
	case DoubleValue:
		return v.f
	case StringValue:
		return v.s
	case ByteArrayValue:
		return v.raw
	default:
		return v.b
	}
}

// UnmarshalJSON decodes the canonical bag-value shape, rejecting unknown tags
// with the offending name.
func (v *Variant) UnmarshalJSON(b []byte) error {
	var enc variantJSON
	if err := json.Unmarshal(b, &enc); err != nil {
		return fmt.Errorf("types: Variant expects {type, value}: %w", err)
	}
	typ, ok := valueTypeByName(enc.Type)
	if !ok {
		return fmt.Errorf("types: Variant type %q is not one of integer|double|string|bytes|boolean", enc.Type)
	}
	decoded, err := decodeVariantPayload(typ, enc.Value)
	if err != nil {
		return err
	}
	*v = decoded
	v.unit = enc.Unit
	return nil
}

// valueTypeByName inverts the stable tag names.
func valueTypeByName(name string) (ValueType, bool) {
	for typ, n := range valueTypeNames {
		if n == name {
			return typ, true
		}
	}
	return 0, false
}

// decodeVariantPayload parses the value field under the already-validated tag.
func decodeVariantPayload(typ ValueType, raw json.RawMessage) (Variant, error) {
	v := Variant{typ: typ}
	var err error
	switch typ {
	case IntegerValue:
		err = json.Unmarshal(raw, &v.i)
	case DoubleValue:
		err = json.Unmarshal(raw, &v.f)
	case StringValue:
		err = json.Unmarshal(raw, &v.s)
	case ByteArrayValue:
		err = json.Unmarshal(raw, &v.raw)
	default:
		err = json.Unmarshal(raw, &v.b)
	}
	if err != nil {
		return Variant{}, fmt.Errorf("types: Variant %s value %s: %w", typ, raw, err)
	}
	return v, nil
}
