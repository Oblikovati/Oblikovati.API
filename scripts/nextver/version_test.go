// SPDX-License-Identifier: Apache-2.0

package main

import "testing"

func TestParseSemverRejectsNonClean(t *testing.T) {
	for _, bad := range []string{"v1.2.3", "1.2", "1.2.3-rc1", "1.2.x", "1.2.3.4", ""} {
		if _, err := parseSemver(bad); err == nil {
			t.Errorf("parseSemver(%q) = nil error, want a rejection", bad)
		}
	}
	v, err := parseSemver("0.1.0")
	if err != nil || v != (semver{0, 1, 0}) {
		t.Fatalf("parseSemver(0.1.0) = %v, %v; want {0 1 0}, nil", v, err)
	}
}

// TestNextVersionPolicy pins the RELEASING.md rules: 0.x has no major bump (feat and
// breaking both move MINOR), and from 1.0.0 on the three scopes map to major/minor/patch.
func TestNextVersionPolicy(t *testing.T) {
	cases := []struct {
		cur     semver
		scope   change
		want    string
		release bool
	}{
		{semver{0, 1, 0}, changeNone, "0.1.0", false},
		{semver{0, 1, 0}, changePatch, "0.1.1", true},
		{semver{0, 1, 0}, changeFeature, "0.2.0", true},
		{semver{0, 1, 5}, changeBreaking, "0.2.0", true}, // 0.x breaking == minor, patch resets
		{semver{1, 4, 2}, changePatch, "1.4.3", true},
		{semver{1, 4, 2}, changeFeature, "1.5.0", true},
		{semver{1, 4, 2}, changeBreaking, "2.0.0", true},
	}
	for _, c := range cases {
		got, rel := nextVersion(c.cur, c.scope)
		if got.String() != c.want || rel != c.release {
			t.Errorf("nextVersion(%v, %s) = %s,%v; want %s,%v", c.cur, c.scope, got, rel, c.want, c.release)
		}
	}
}
