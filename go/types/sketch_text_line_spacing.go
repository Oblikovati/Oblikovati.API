// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"math"
)

// TextLineSpacingMode selects the distance between successive Sketch Text baselines.
type TextLineSpacingMode string

const (
	TextLineSpacingSingle     TextLineSpacingMode = "single"
	TextLineSpacingOneAndHalf TextLineSpacingMode = "oneAndHalf"
	TextLineSpacingDouble     TextLineSpacingMode = "double"
	TextLineSpacingMultiple   TextLineSpacingMode = "multiple"
	TextLineSpacingExactly    TextLineSpacingMode = "exactly"
)

// SketchTextLineSpacing specifies how successive text baselines are placed.
// Single, OneAndHalf and Double use the font's normal baseline advance scaled
// by 1, 1.5 and 2. Multiple scales that advance by Value. Exactly uses Value
// centimetres between baselines. Omitting the spacing uses Single.
type SketchTextLineSpacing struct {
	Mode  TextLineSpacingMode `json:"mode"`
	Value float64             `json:"value,omitempty"`
}

// Validate checks a present spacing value. A nil spacing uses the caller's default.
// For example, set Mode to TextLineSpacingDouble and leave Value zero.
func (spacing *SketchTextLineSpacing) Validate() error {
	if spacing == nil {
		return nil
	}
	switch spacing.Mode {
	case TextLineSpacingSingle, TextLineSpacingOneAndHalf, TextLineSpacingDouble:
		if spacing.Value == 0 {
			return nil
		}
		return fmt.Errorf("line spacing %q value %g: expected zero for a fixed mode", spacing.Mode, spacing.Value)
	case TextLineSpacingMultiple, TextLineSpacingExactly:
		if spacing.Value > 0 && !math.IsInf(spacing.Value, 0) {
			return nil
		}
		return fmt.Errorf("line spacing %q value %g: expected a positive finite value", spacing.Mode, spacing.Value)
	default:
		return fmt.Errorf("line spacing mode %q: expected single, oneAndHalf, double, multiple, or exactly", spacing.Mode)
	}
}
