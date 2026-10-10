// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestToleranceTypeFrozenBlock pins the reference ids and wire spellings — the
// block is frozen and a renumber or respell is a breaking wire change.
func TestToleranceTypeFrozenBlock(t *testing.T) {
	want := map[ToleranceType]string{
		31233: "default", 31234: "override", 31235: "symmetric",
		31236: "deviation", 31237: "limitsStacked", 31238: "limitLinear",
		31239: "max", 31240: "min", 31241: "limitsFitsStacked",
		31242: "limitsFitsLinear", 31243: "limitsFitsShowSize",
		31244: "limitsFitsShowTolerance", 31245: "basic", 31246: "reference",
	}
	if len(want) != len(toleranceTypeNames) {
		t.Fatalf("tolerance type count = %d, want %d", len(toleranceTypeNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("ToleranceType(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseToleranceType(name); !ok || parsed != v {
			t.Errorf("ParseToleranceType(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
	if _, ok := ParseToleranceType("noSuchTolerance"); ok {
		t.Error("ParseToleranceType must reject unknown spellings")
	}
}

// TestModelValueTypeFrozenBlock pins the reference ids — note the reference
// orders lower (31490) before upper (31491).
func TestModelValueTypeFrozenBlock(t *testing.T) {
	want := map[ModelValueType]string{
		31489: "nominal", 31490: "lower", 31491: "upper", 31492: "median",
	}
	if len(want) != len(modelValueTypeNames) {
		t.Fatalf("model value type count = %d, want %d", len(modelValueTypeNames), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("ModelValueType(%d).String() = %q, want %q", int32(v), got, name)
		}
		if parsed, ok := ParseModelValueType(name); !ok || parsed != v {
			t.Errorf("ParseModelValueType(%q) = (%v, %v), want %d", name, parsed, ok, int32(v))
		}
	}
}
