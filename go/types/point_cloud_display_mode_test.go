// SPDX-License-Identifier: Apache-2.0

package types

import (
	"strings"
	"testing"
)

// TestPointCloudDisplayModeTokensAreStable pins the on-the-wire / persisted string tokens so a
// stored `.obk` document or a peer's JSON never silently changes meaning (#645).
func TestPointCloudDisplayModeTokensAreStable(t *testing.T) {
	want := map[PointCloudDisplayMode]string{
		PointCloudDisplayModeDefault:   "default",
		PointCloudDisplayModeRGB:       "rgb",
		PointCloudDisplayModeIntensity: "intensity",
	}
	for mode, token := range want {
		if string(mode) != token {
			t.Errorf("%s token = %q, want %q", mode, string(mode), token)
		}
	}
}

// TestPointCloudDisplayModeLabels pins the user-facing labels (used for the selector and the host's
// command ids) and the round-trip invariant the persistence layer relies on: lower-casing the label
// yields the underlying token, so a mode written via String and read back via strings.ToLower is
// unchanged.
func TestPointCloudDisplayModeLabels(t *testing.T) {
	want := map[PointCloudDisplayMode]string{
		PointCloudDisplayModeDefault:   "Default",
		PointCloudDisplayModeRGB:       "RGB",
		PointCloudDisplayModeIntensity: "Intensity",
	}
	for mode, label := range want {
		if mode.String() != label {
			t.Errorf("%s.String() = %q, want %q", string(mode), mode.String(), label)
		}
		if got := PointCloudDisplayMode(strings.ToLower(mode.String())); got != mode {
			t.Errorf("round-trip of %q via label %q = %q, want identity", string(mode), label, string(got))
		}
	}
	if got := PointCloudDisplayMode("bogus").String(); got != "pointCloudDisplayMode(?)" {
		t.Errorf("unknown mode String() = %q, want placeholder", got)
	}
}

// TestPointCloudDisplayModeValidity checks every AllPointCloudDisplayModes entry is valid and that
// an unknown token — including the zero value — is rejected, which is what the host's SetDisplayMode
// leans on to refuse a bogus mode.
func TestPointCloudDisplayModeValidity(t *testing.T) {
	for _, mode := range AllPointCloudDisplayModes() {
		if !mode.IsValid() {
			t.Errorf("AllPointCloudDisplayModes returned %q but IsValid is false", mode)
		}
	}
	if PointCloudDisplayMode("").IsValid() {
		t.Errorf("empty display mode must be invalid")
	}
	if PointCloudDisplayMode("bogus").IsValid() {
		t.Errorf("unknown display mode must be invalid")
	}
}

// TestAllPointCloudDisplayModesIsComplete guards against adding a constant without listing it in
// AllPointCloudDisplayModes (which feeds the picker and the validation message).
func TestAllPointCloudDisplayModesIsComplete(t *testing.T) {
	if got := len(AllPointCloudDisplayModes()); got != 3 {
		t.Errorf("AllPointCloudDisplayModes length = %d, want 3 — update this test when adding a mode", got)
	}
}
