// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"
	"testing"
)

// TestValueTypeValuesAreFrozen pins the reference ValueTypeEnum block.
func TestValueTypeValuesAreFrozen(t *testing.T) {
	if IntegerValue != 14593 || DoubleValue != 14594 || StringValue != 14595 ||
		ByteArrayValue != 14596 || BooleanValue != 14597 {
		t.Error("ValueType values drifted from the frozen block 14593…")
	}
}

// TestVariantJSONIsCanonical pins the wire encoding of every tag.
func TestVariantJSONIsCanonical(t *testing.T) {
	cases := []struct {
		name string
		v    Variant
		want string
	}{
		{"integer", IntegerVariant(42), `{"type":"integer","value":42}`},
		{"double", DoubleVariant(2.5), `{"type":"double","value":2.5}`},
		{"unit double", UnitVariant(2.5, "mm"), `{"type":"double","value":2.5,"unit":"mm"}`},
		{"string", StringVariant("steel"), `{"type":"string","value":"steel"}`},
		{"boolean", BoolVariant(true), `{"type":"boolean","value":true}`},
		{"bytes", BytesVariant([]byte{1, 2}), `{"type":"bytes","value":"AQI="}`},
	}
	for _, tc := range cases {
		got, err := json.Marshal(tc.v)
		if err != nil {
			t.Fatalf("%s marshal: %v", tc.name, err)
		}
		if string(got) != tc.want {
			t.Errorf("%s JSON = %s, want %s", tc.name, got, tc.want)
		}
	}
}

func TestVariantJSONRoundTrip(t *testing.T) {
	for _, v := range []Variant{
		IntegerVariant(-7), UnitVariant(0.1, "deg"), StringVariant(""),
		BoolVariant(false), BytesVariant([]byte("raw")),
	} {
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("marshal %s: %v", v.Type(), err)
		}
		var back Variant
		if err := json.Unmarshal(b, &back); err != nil {
			t.Fatalf("unmarshal %s (%s): %v", v.Type(), b, err)
		}
		if back.Type() != v.Type() || back.Unit() != v.Unit() {
			t.Errorf("round trip changed tag/unit: %s → %s", b, back.Type())
		}
	}
	var v Variant
	if err := json.Unmarshal([]byte(`{"type":"color","value":1}`), &v); err == nil {
		t.Error("an unknown type tag must be rejected with the offending name")
	}
	if err := json.Unmarshal([]byte(`{"type":"integer","value":"nope"}`), &v); err == nil {
		t.Error("a tag/value mismatch must be rejected")
	}
}

func TestVariantAccessorsEnforceTag(t *testing.T) {
	v := UnitVariant(3.5, "mm")
	if f, ok := v.Double(); !ok || f != 3.5 || v.Unit() != "mm" {
		t.Errorf("Double() = (%v, %v), Unit() = %q", f, ok, v.Unit())
	}
	if _, ok := v.Str(); ok {
		t.Error("Str() on a double must report ok=false")
	}
	if _, ok := IntegerVariant(1).Bool(); ok {
		t.Error("Bool() on an integer must report ok=false")
	}
	if i, ok := IntegerVariant(9).Integer(); !ok || i != 9 {
		t.Errorf("Integer() = (%v, %v)", i, ok)
	}
}
