// SPDX-License-Identifier: Apache-2.0

// Command nextver derives the next Semantic Version of oblikovati.org/api from the
// scope of the conventional-commit messages merged since the last release tag, and
// (with -apply) rewrites version.go and CHANGELOG.md to it. The release workflow runs
// it on every push to develop so a merge auto-bumps the version per its scope instead
// of relying on a human to edit version.go (see RELEASING.md).
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// semver is a parsed MAJOR.MINOR.PATCH triple. Pre-release / build metadata is not
// used by the bump policy, so the parser rejects it rather than silently dropping it.
type semver struct{ major, minor, patch int }

// parseSemver parses a clean "MAJOR.MINOR.PATCH" string (no leading "v", no
// pre-release/build suffix — the shape api.Version always carries).
func parseSemver(s string) (semver, error) {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return semver{}, fmt.Errorf("version %q is not MAJOR.MINOR.PATCH (got %d dot-separated parts)", s, len(parts))
	}
	nums := [3]int{}
	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil || n < 0 {
			return semver{}, fmt.Errorf("version %q has a non-numeric component %q; expected MAJOR.MINOR.PATCH", s, p)
		}
		nums[i] = n
	}
	return semver{nums[0], nums[1], nums[2]}, nil
}

func (v semver) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

// nextVersion applies the project's versioning policy (RELEASING.md) to bump cur for a
// change of the given scope. The second return is false when nothing releasable changed
// (changeNone), so the caller cuts no release.
//
// While MAJOR is 0 the surface is in initial development: an additive OR breaking change
// bumps MINOR and a fix bumps PATCH (semver §4). From 1.0.0 on, breaking→MAJOR,
// additive→MINOR, fix→PATCH.
func nextVersion(cur semver, c change) (semver, bool) {
	switch {
	case c == changeNone:
		return cur, false
	case cur.major == 0:
		if c == changePatch {
			return semver{0, cur.minor, cur.patch + 1}, true
		}
		return semver{0, cur.minor + 1, 0}, true // feat or breaking — 0.x has no major bump
	case c == changeBreaking:
		return semver{cur.major + 1, 0, 0}, true
	case c == changeFeature:
		return semver{cur.major, cur.minor + 1, 0}, true
	default: // changePatch
		return semver{cur.major, cur.minor, cur.patch + 1}, true
	}
}
