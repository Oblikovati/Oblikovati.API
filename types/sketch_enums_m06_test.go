// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestAccuracyFrozenBlock pins the reference ids and wire spellings.
func TestAccuracyFrozenBlock(t *testing.T) {
	want := map[Accuracy]string{
		69377: "low", 69378: "medium", 69379: "high", 69380: "veryHigh",
	}
	assertFrozenBlock(t, "Accuracy", want, accuracyNames, ParseAccuracy)
}

// TestGeometryMoveableStatusFrozenBlock pins the reference ids and wire spellings.
func TestGeometryMoveableStatusFrozenBlock(t *testing.T) {
	want := map[GeometryMoveableStatus]string{
		53505: "freeToMove", 53506: "byDimensionChange", 53507: "fixed", 53508: "unknown",
	}
	assertFrozenBlock(t, "GeometryMoveableStatus", want, geometryMoveableStatusNames,
		ParseGeometryMoveableStatus)
}

// TestSketchPointInferenceFrozenBlock pins the reference ids and wire spellings.
func TestSketchPointInferenceFrozenBlock(t *testing.T) {
	want := map[SketchPointInferenceKind]string{
		22273: "atIntersection", 22274: "onCurve", 22275: "onPoint", 22276: "atMidpoint",
	}
	assertFrozenBlock(t, "SketchPointInferenceKind", want, sketchPointInferenceNames,
		ParseSketchPointInferenceKind)
}

// TestConstraintInferenceFrozenBlock pins the reference bit flags and spellings.
func TestConstraintInferenceFrozenBlock(t *testing.T) {
	want := map[ConstraintInferenceKind]string{
		1: "coincident", 2: "horizontal", 4: "intersection", 8: "midpoint",
		16: "onCurve", 32: "parallel", 64: "perpendicular", 128: "tangent", 256: "vertical",
	}
	assertFrozenBlock(t, "ConstraintInferenceKind", want, constraintInferenceNames,
		ParseConstraintInferenceKind)
}

// TestConstraintInferencePriorityFrozenBlock pins the reference ids and spellings.
func TestConstraintInferencePriorityFrozenBlock(t *testing.T) {
	want := map[ConstraintInferencePriority]string{
		50433: "parallelPerpendicular", 50434: "horizontalVertical", 50435: "none",
	}
	assertFrozenBlock(t, "ConstraintInferencePriority", want, constraintInferencePriorityNames,
		ParseConstraintInferencePriority)
}

// TestHelicalShapeDefinitionFrozenBlock pins the reference ids and wire
// spellings — which must also stay equal to the helix Mode strings accepted
// by sketch3d.addEntity since M22-F04.
func TestHelicalShapeDefinitionFrozenBlock(t *testing.T) {
	want := map[HelicalShapeDefinitionKind]string{
		115713: "pitchRevolution", 115714: "pitchHeight",
		115715: "revolutionHeight", 115716: "spiral",
	}
	assertFrozenBlock(t, "HelicalShapeDefinitionKind", want, helicalShapeDefinitionNames,
		ParseHelicalShapeDefinitionKind)
}

// TestHelixEndFrozenBlock pins the reference ids and wire spellings.
func TestHelixEndFrozenBlock(t *testing.T) {
	want := map[HelixEndKind]string{115969: "natural", 115970: "flat"}
	assertFrozenBlock(t, "HelixEndKind", want, helixEndNames, ParseHelixEndKind)
}

// TestSplineFitMethodFrozenBlock pins the reference ids and wire spellings.
func TestSplineFitMethodFrozenBlock(t *testing.T) {
	want := map[SplineFitMethod]string{26369: "smooth", 26370: "sweet", 26371: "chord"}
	assertFrozenBlock(t, "SplineFitMethod", want, splineFitMethodNames, ParseSplineFitMethod)
}

// assertFrozenBlock checks a frozen enum block: exact member set, String()
// spellings, and Parse round-trip including unknown rejection.
func assertFrozenBlock[E interface {
	comparable
	String() string
}](t *testing.T, label string, want, names map[E]string, parse func(string) (E, bool)) {
	t.Helper()
	if len(want) != len(names) {
		t.Fatalf("%s member count = %d, want %d", label, len(names), len(want))
	}
	for v, name := range want {
		if got := v.String(); got != name {
			t.Errorf("%s.String() for %v = %q, want %q", label, v, got, name)
		}
		if parsed, ok := parse(name); !ok || parsed != v {
			t.Errorf("Parse%s(%q) = (%v, %v), want %v", label, name, parsed, ok, v)
		}
	}
	if _, ok := parse("noSuchSpelling"); ok {
		t.Errorf("Parse%s must reject unknown spellings", label)
	}
}
