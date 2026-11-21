// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestSketchTextFractionRangeAndContent(t *testing.T) {
	for _, test := range []struct {
		name, content string
		ranges        []SketchTextFraction
		valid         bool
	}{
		{"Unicode offsets", "⌀1/2 3/4", []SketchTextFraction{{1, 4}, {5, 8}}, true},
		{"Unicode terms", "界/二", []SketchTextFraction{{0, 3}}, true},
		{"overlap", "1/2/3", []SketchTextFraction{{0, 3}, {2, 5}}, false},
		{"outside content", "1/2", []SketchTextFraction{{0, 4}}, false},
		{"short range", "1/2", []SketchTextFraction{{0, 2}}, false},
		{"blank term", " /2", []SketchTextFraction{{0, 3}}, false},
		{"two slashes", "1//2", []SketchTextFraction{{0, 4}}, false},
		{"multiline", "1/\n2", []SketchTextFraction{{0, 4}}, false},
		{"invalid UTF-8", string([]byte{255, '/', '2'}), []SketchTextFraction{{0, 3}}, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			if err := ValidateSketchTextFractions(&test.content, test.ranges); (err == nil) != test.valid {
				t.Fatalf("ranges %v in %q: expected valid=%v, got %v", test.ranges, test.content, test.valid, err)
			}
		})
	}
}

func TestSketchTextFractionPartialEditDefersContent(t *testing.T) {
	if err := ValidateSketchTextFractions(nil, []SketchTextFraction{{9, 12}}); err != nil {
		t.Fatal(err)
	}
	if err := ValidateSketchTextFractions(nil, []SketchTextFraction{{9, 10}}); err == nil {
		t.Fatal("expected a short range to be rejected even in a partial edit")
	}
}
