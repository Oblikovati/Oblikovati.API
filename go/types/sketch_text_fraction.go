// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// SketchTextFraction stacks a slash-delimited range without altering the source
// text. Start and End are zero-based Unicode scalar offsets in a half-open
// range. For example, {Start: 1, End: 4} stacks "1/2" in "⌀1/2".
type SketchTextFraction struct {
	Start uint32 `json:"start"`
	End   uint32 `json:"end"`
}

// ValidateSketchTextFractions checks sorted, non-overlapping ranges. A nil
// content defers the source check to the host. For example, validate &text
// against the ranges before saving a text entity.
func ValidateSketchTextFractions(content *string, fractions []SketchTextFraction) error {
	if len(fractions) == 0 {
		return nil
	}
	scalars, err := fractionScalars(content)
	if err != nil {
		return err
	}
	var previousEnd uint32
	for index, fraction := range fractions {
		if err := fraction.validate(index, previousEnd, scalars); err != nil {
			return err
		}
		previousEnd = fraction.End
	}
	return nil
}

func fractionScalars(content *string) (*[]rune, error) {
	if content == nil {
		return nil, nil
	}
	if !utf8.ValidString(*content) {
		return nil, fmt.Errorf("stacked sketch text %q: expected valid UTF-8", *content)
	}
	scalars := []rune(*content)
	return &scalars, nil
}

func (fraction SketchTextFraction) validate(index int, previousEnd uint32, scalars *[]rune) error {
	if fraction.Start < previousEnd || fraction.End <= fraction.Start ||
		fraction.End-fraction.Start < 3 {
		return fmt.Errorf("stacked fraction %d [%d,%d): expected start >= %d and at least three Unicode scalars",
			index, fraction.Start, fraction.End, previousEnd)
	}
	if scalars == nil {
		return nil
	}
	if uint64(fraction.End) > uint64(len(*scalars)) {
		return fmt.Errorf("stacked fraction %d end %d: expected at most %d Unicode scalars", index, fraction.End, len(*scalars))
	}
	return validateFractionContent(index, string((*scalars)[fraction.Start:fraction.End]))
}

func validateFractionContent(index int, selected string) error {
	top, bottom, found := strings.Cut(selected, "/")
	if !found || strings.TrimSpace(top) == "" || strings.TrimSpace(bottom) == "" ||
		strings.Contains(bottom, "/") || strings.ContainsAny(selected, "\r\n") {
		return fmt.Errorf("stacked fraction %d content %q: expected one slash between two non-empty values on one line", index, selected)
	}
	return nil
}
