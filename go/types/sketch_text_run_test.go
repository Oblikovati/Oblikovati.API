// SPDX-License-Identifier: Apache-2.0

package types

import (
	"math"
	"strings"
	"testing"
)

func TestSketchTextRunUnicodeRanges(t *testing.T) {
	content := "A🦀e\u0301界"
	width := 50.0
	color := NewColor(80, 0, 110)
	runs := []SketchTextRun{
		{Start: 1, End: 2, Bold: true},
		{Start: 3, End: 4, Italic: true},
		{Start: 4, End: 5, Underline: true, WidthPercent: &width, Color: &color},
	}
	if err := ValidateSketchTextRuns(&content, runs); err != nil {
		t.Fatal(err)
	}
	if err := ValidateSketchTextRuns(nil, []SketchTextRun{{Start: 6, End: 7}}); err != nil {
		t.Fatalf("format-only edit must defer the length check: %v", err)
	}
}

func TestInvalidSketchTextRuns(t *testing.T) {
	content := "A🦀e\u0301界"
	for _, test := range invalidSketchTextRuns() {
		t.Run(test.message, func(t *testing.T) {
			err := ValidateSketchTextRuns(&content, test.runs)
			if err == nil || !strings.Contains(err.Error(), test.message) {
				t.Fatalf("runs %+v: got %v, expected %q", test.runs, err, test.message)
			}
		})
	}
	invalid := string([]byte{0xff})
	if err := ValidateSketchTextRuns(&invalid, []SketchTextRun{{End: 1}}); err == nil {
		t.Fatal("formatted invalid UTF-8 must be rejected")
	}
}

type invalidSketchTextCase struct {
	message string
	runs    []SketchTextRun
}

func invalidSketchTextRuns() []invalidSketchTextCase {
	zero, nonfinite := 0.0, math.Inf(1)
	inherited := Color{Opacity: 1}
	badOpacity := NewColor(1, 1, 1)
	badOpacity.Opacity = math.NaN()
	return []invalidSketchTextCase{
		{"start < end", []SketchTextRun{{Start: 2, End: 2}}},
		{"2 <= start", []SketchTextRun{{End: 2}, {Start: 1, End: 3}}},
		{"5 Unicode scalars", []SketchTextRun{{Start: 2, End: 6}}},
		{"positive finite percentage", []SketchTextRun{{End: 2, WidthPercent: &zero}}},
		{"width +Inf", []SketchTextRun{{End: 2, WidthPercent: &nonfinite}}},
		{"override color", []SketchTextRun{{End: 2, Color: &inherited}}},
		{"opacity in [0,1]", []SketchTextRun{{End: 2, Color: &badOpacity}}},
	}
}
