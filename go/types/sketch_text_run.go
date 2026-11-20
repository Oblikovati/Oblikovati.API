// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"math"
	"unicode/utf8"
)

// SketchTextRun applies one style to a half-open range of zero-based Unicode
// scalar positions in the original content. Combining marks each occupy one
// position. Omitted settings use the base style and gaps retain plain text.
// For example, SketchTextRun{Start: 1, End: 2, Bold: true} styles 界 in "a界🙂".
// Runs are sorted by Start, do not overlap, and may be adjacent.
type SketchTextRun struct {
	Start        uint32   `json:"start"`
	End          uint32   `json:"end"`
	Bold         bool     `json:"bold,omitempty"`
	Italic       bool     `json:"italic,omitempty"`
	Underline    bool     `json:"underline,omitempty"`
	WidthPercent *float64 `json:"widthPercent,omitempty"` // positive percent; 100 leaves the width unchanged
	Color        *Color   `json:"color,omitempty"`        // explicit override with opacity in [0,1]
}

// ValidateSketchTextRuns verifies ordered ranges and their formatting. A nil
// content defers the final length check until the host resolves a partial edit.
// For example, ValidateSketchTextRuns(&content, runs) checks a complete value.
func ValidateSketchTextRuns(content *string, runs []SketchTextRun) error {
	if len(runs) == 0 {
		return nil
	}
	var limit *int
	if content != nil {
		if !utf8.ValidString(*content) {
			return fmt.Errorf("formatted sketch text %q: expected valid UTF-8", *content)
		}
		count := utf8.RuneCountInString(*content)
		limit = &count
	}
	var previousEnd uint32
	for index, run := range runs {
		if err := validateSketchTextRun(run, index, previousEnd, limit); err != nil {
			return err
		}
		previousEnd = run.End
	}
	return nil
}

func validateSketchTextRun(run SketchTextRun, index int, previousEnd uint32, limit *int) error {
	if run.Start < previousEnd || run.Start >= run.End {
		return fmt.Errorf("sketch text run %d [%d,%d): expected %d <= start < end", index, run.Start, run.End, previousEnd)
	}
	if limit != nil && uint64(run.End) > uint64(*limit) {
		return fmt.Errorf("sketch text run %d end %d: expected at most %d Unicode scalars", index, run.End, *limit)
	}
	if run.WidthPercent != nil {
		width := *run.WidthPercent
		if width <= 0 || math.IsNaN(width) || math.IsInf(width, 0) {
			return fmt.Errorf("sketch text run %d width %g: expected a positive finite percentage", index, width)
		}
	}
	if run.Color != nil && !validSketchTextRunColor(*run.Color) {
		return fmt.Errorf("sketch text run %d color %+v: expected override color with opacity in [0,1]", index, *run.Color)
	}
	return nil
}

func validSketchTextRunColor(color Color) bool {
	return color.IsOverride() && color.Opacity >= 0 && color.Opacity <= 1
}
