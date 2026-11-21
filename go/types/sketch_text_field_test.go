// SPDX-License-Identifier: Apache-2.0

package types

import (
	"strings"
	"testing"
)

func TestSketchTextFieldUsesSourceScalarPositions(t *testing.T) {
	content := "🦀" + SketchTextFieldPlaceholder + " 1/2"
	field := SketchTextField{At: 1, Kind: SketchTextFieldModelParameter, Parameter: "d0"}
	if err := ValidateSketchTextFields(&content, []SketchTextField{field}, []SketchTextFraction{{Start: 3, End: 6}}); err != nil {
		t.Fatal(err)
	}
	if err := ValidateSketchTextFields(nil, []SketchTextField{field}, nil); err != nil {
		t.Fatal(err)
	}
	if err := ValidateSketchTextFields(&content, []SketchTextField{field, field}, nil); err == nil || !strings.Contains(err.Error(), "greater than") {
		t.Fatalf("error %v: expected duplicate position to be rejected", err)
	}
}

func TestSketchTextFieldRejectsStackedPlacementAndInvalidSource(t *testing.T) {
	content := "1/" + SketchTextFieldPlaceholder
	field := SketchTextField{At: 2, Kind: SketchTextFieldFileName}
	if err := ValidateSketchTextFields(&content, []SketchTextField{field}, []SketchTextFraction{{Start: 0, End: 3}}); err == nil || !strings.Contains(err.Error(), "outside all stacked fractions") {
		t.Fatalf("error %v: expected a field to be rejected inside a stacked fraction", err)
	}
	field.Parameter = "d0"
	if err := ValidateSketchTextFields(&content, []SketchTextField{field}, nil); err == nil || !strings.Contains(err.Error(), "no parameter name or precision") {
		t.Fatalf("error %v: expected file name fields to reject a parameter", err)
	}
}
