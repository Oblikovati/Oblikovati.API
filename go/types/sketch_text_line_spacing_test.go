// SPDX-License-Identifier: Apache-2.0

package types

import (
	"math"
	"testing"
)

func TestSketchTextLineSpacingModes(t *testing.T) {
	for _, spacing := range []*SketchTextLineSpacing{
		nil,
		{Mode: TextLineSpacingSingle},
		{Mode: TextLineSpacingOneAndHalf},
		{Mode: TextLineSpacingDouble},
		{Mode: TextLineSpacingMultiple, Value: 1.75},
		{Mode: TextLineSpacingExactly, Value: 0.61},
	} {
		if err := spacing.Validate(); err != nil {
			t.Fatalf("spacing %+v: expected valid mode, got %v", spacing, err)
		}
	}
}

func TestSketchTextLineSpacingRejectsInvalidValues(t *testing.T) {
	for _, spacing := range []SketchTextLineSpacing{
		{Mode: "unknown"},
		{Mode: TextLineSpacingDouble, Value: 2},
		{Mode: TextLineSpacingMultiple},
		{Mode: TextLineSpacingExactly, Value: -1},
		{Mode: TextLineSpacingMultiple, Value: math.Inf(1)},
		{Mode: TextLineSpacingExactly, Value: math.NaN()},
	} {
		if err := spacing.Validate(); err == nil {
			t.Fatalf("spacing %+v: expected validation failure", spacing)
		}
	}
}
