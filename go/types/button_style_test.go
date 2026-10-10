// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestButtonStyleString(t *testing.T) {
	cases := map[ButtonStyle]string{
		TextOnlyButton:    "text",
		SmallIconButton:   "small-icon",
		LargeIconButton:   "large-icon",
		CompactIconButton: "compact-icon",
		ButtonStyle(99):   "buttonStyle(?)",
	}
	for s, want := range cases {
		if got := s.String(); got != want {
			t.Errorf("ButtonStyle(%d).String() = %q, want %q", uint8(s), got, want)
		}
	}
}

func TestButtonStyleShowsIcon(t *testing.T) {
	if TextOnlyButton.ShowsIcon() {
		t.Error("TextOnlyButton.ShowsIcon() = true, want false")
	}
	for _, s := range []ButtonStyle{SmallIconButton, LargeIconButton, CompactIconButton} {
		if !s.ShowsIcon() {
			t.Errorf("%s.ShowsIcon() = false, want true", s)
		}
	}
}
