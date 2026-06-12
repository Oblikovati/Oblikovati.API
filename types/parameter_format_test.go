// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestParameterDisplayFormatFrozenBlock pins the reference ids and wire spellings.
func TestParameterDisplayFormatFrozenBlock(t *testing.T) {
	want := map[ParameterDisplayFormat]string{
		92417: "decimal", 92418: "fractional", 92419: "architectural",
	}
	if len(want) != len(parameterDisplayFormatNames) {
		t.Fatalf("display format count = %d, want %d", len(parameterDisplayFormatNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("ParameterDisplayFormat(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseParameterDisplayFormat(name); !ok || parsed != v {
			t.Errorf("ParseParameterDisplayFormat(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
}

// TestCustomPropertyTypeFrozenBlock pins the reference ids and wire spellings.
func TestCustomPropertyTypeFrozenBlock(t *testing.T) {
	want := map[CustomPropertyType]string{85249: "text", 85250: "number"}
	if len(want) != len(customPropertyTypeNames) {
		t.Fatalf("custom property type count = %d, want %d", len(customPropertyTypeNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("CustomPropertyType(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseCustomPropertyType(name); !ok || parsed != v {
			t.Errorf("ParseCustomPropertyType(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
}

// TestCustomPropertyPrecisionFrozenBlock pins the 24-value reference block:
// decimal places 85505–85513, fractional lengths 85514–85521, angular forms
// 85522–85528.
func TestCustomPropertyPrecisionFrozenBlock(t *testing.T) {
	want := map[CustomPropertyPrecision]string{
		85505: "zeroDecimalPlace", 85506: "oneDecimalPlace", 85507: "twoDecimalPlaces",
		85508: "threeDecimalPlaces", 85509: "fourDecimalPlaces", 85510: "fiveDecimalPlaces",
		85511: "sixDecimalPlaces", 85512: "sevenDecimalPlaces", 85513: "eightDecimalPlaces",
		85514: "zeroFractional", 85515: "halfFractional", 85516: "quarterFractional",
		85517: "eighthsFractional", 85518: "sixteenthsFractional",
		85519: "thirtySecondsFractional", 85520: "sixtyFourthsFractional",
		85521: "oneTwentyEighthsFractional", 85522: "degrees", 85523: "minutes",
		85524: "seconds", 85525: "secondsOneDecimalPlace", 85526: "secondsTwoDecimalPlaces",
		85527: "secondsThreeDecimalPlaces", 85528: "secondsFourDecimalPlaces",
	}
	if len(want) != len(customPropertyPrecisionNames) {
		t.Fatalf("precision count = %d, want %d", len(customPropertyPrecisionNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("CustomPropertyPrecision(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseCustomPropertyPrecision(name); !ok || parsed != v {
			t.Errorf("ParseCustomPropertyPrecision(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
}
