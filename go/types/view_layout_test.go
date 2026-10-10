// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestViewLayoutTilesAndValidity(t *testing.T) {
	cases := map[ViewLayout]int{
		LayoutSingle: 1,
		LayoutTwoH:   2,
		LayoutTwoV:   2,
		LayoutThree:  3,
		LayoutFour:   4,
	}
	for l, want := range cases {
		if !l.IsValid() {
			t.Errorf("%v should be valid", l)
		}
		if got := l.Tiles(); got != want {
			t.Errorf("%v.Tiles() = %d, want %d", l, got, want)
		}
		if l.String() == "viewLayout(?)" {
			t.Errorf("%v has no name", l)
		}
	}
	if ViewLayout(99).IsValid() {
		t.Error("undefined layout should be invalid")
	}
	if ViewLayout(99).Tiles() != 1 {
		t.Error("undefined layout should default to 1 tile")
	}
	if got := len(AllViewLayouts()); got != len(cases) {
		t.Errorf("AllViewLayouts len = %d, want %d", got, len(cases))
	}
}
