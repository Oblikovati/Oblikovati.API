// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"regexp"
	"strings"
)

// versionConstRe matches the `Version = "X.Y.Z"` constant in version.go, capturing the
// quoted value so it can be read and replaced in place.
var versionConstRe = regexp.MustCompile(`(Version = ")([0-9A-Za-z.+-]+)(")`)

// readVersionConst extracts the current Version string from version.go's source.
func readVersionConst(src string) (string, error) {
	m := versionConstRe.FindStringSubmatch(src)
	if m == nil {
		return "", fmt.Errorf(`version.go has no ` + "`Version = \"X.Y.Z\"`" + ` constant; cannot read the current version`)
	}
	return m[2], nil
}

// bumpVersionGo returns version.go's source with the Version constant set to newVer.
func bumpVersionGo(src, newVer string) (string, error) {
	if !versionConstRe.MatchString(src) {
		return "", fmt.Errorf(`version.go has no `+"`Version = \"X.Y.Z\"`"+` constant to bump to %q`, newVer)
	}
	return versionConstRe.ReplaceAllString(src, "${1}"+newVer+"${3}"), nil
}

const repoCompare = "https://github.com/Oblikovati/Oblikovati.API/compare"

// rollChangelog moves the "## [Unreleased]" marker down over a new "## [newVer] - date"
// section (filled with notes), and refreshes the link references at the bottom so
// [Unreleased] compares from the new tag and a [newVer] link is added.
func rollChangelog(src, prevVer, newVer, date, notes string) (string, error) {
	const marker = "## [Unreleased]"
	if !strings.Contains(src, marker) {
		return "", fmt.Errorf("CHANGELOG.md has no %q heading to roll to %s", marker, newVer)
	}
	section := fmt.Sprintf("%s\n\n## [%s] - %s", marker, newVer, date)
	if notes != "" {
		section += "\n\n" + notes
	}
	out := strings.Replace(src, marker, section, 1)
	return rewriteLinks(out, prevVer, newVer)
}

// rewriteLinks points the [Unreleased] compare reference at the new tag and inserts a
// [newVer] compare reference (prevVer...newVer) directly after it.
func rewriteLinks(src, prevVer, newVer string) (string, error) {
	unreleasedRe := regexp.MustCompile(`(?m)^\[Unreleased\]:.*$`)
	if !unreleasedRe.MatchString(src) {
		return "", fmt.Errorf("CHANGELOG.md has no [Unreleased] link reference to update for %s", newVer)
	}
	newLink := fmt.Sprintf("[Unreleased]: %s/v%s...HEAD\n[%s]: %s/v%s...v%s",
		repoCompare, newVer, newVer, repoCompare, prevVer, newVer)
	return unreleasedRe.ReplaceAllString(src, newLink), nil
}
