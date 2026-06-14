// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestDriveVariableStringAndValid(t *testing.T) {
	cases := []struct {
		v     DriveVariable
		name  string
		valid bool
	}{
		{DriveNatural, "natural", true},
		{DriveAngular, "angular", true},
		{DriveLinear, "linear", true},
		{DriveLinear + 1, "unknown", false},
	}
	for _, c := range cases {
		if got := c.v.String(); got != c.name {
			t.Errorf("DriveVariable(%d).String() = %q, want %q", c.v, got, c.name)
		}
		if got := c.v.IsValid(); got != c.valid {
			t.Errorf("DriveVariable(%d).IsValid() = %v, want %v", c.v, got, c.valid)
		}
	}
}
