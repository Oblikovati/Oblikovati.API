// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestRepresentationKindStringAndValid(t *testing.T) {
	cases := []struct {
		k     RepresentationKind
		name  string
		valid bool
	}{
		{RepresentationUnknown, "unknown", false},
		{RepresentationDesignView, "designView", true},
		{RepresentationPositional, "positional", true},
		{RepresentationLevelOfDetail, "levelOfDetail", true},
		{RepresentationLevelOfDetail + 1, "unknown", false},
	}
	for _, c := range cases {
		if got := c.k.String(); got != c.name {
			t.Errorf("RepresentationKind(%d).String() = %q, want %q", c.k, got, c.name)
		}
		if got := c.k.IsValid(); got != c.valid {
			t.Errorf("RepresentationKind(%d).IsValid() = %v, want %v", c.k, got, c.valid)
		}
	}
}
