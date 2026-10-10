// SPDX-License-Identifier: Apache-2.0

package main

import (
	"strings"
	"testing"
)

const sampleVersionGo = `package api

// Version is the module's version.
const Version = "0.1.0"

func Major() int { return 0 }
`

func TestReadAndBumpVersionGo(t *testing.T) {
	cur, err := readVersionConst(sampleVersionGo)
	if err != nil || cur != "0.1.0" {
		t.Fatalf("readVersionConst = %q, %v; want 0.1.0", cur, err)
	}
	out, err := bumpVersionGo(sampleVersionGo, "0.2.0")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, `const Version = "0.2.0"`) {
		t.Errorf("bumped source missing new version:\n%s", out)
	}
	if strings.Contains(out, "0.1.0") {
		t.Errorf("old version still present:\n%s", out)
	}
	if _, err := bumpVersionGo("package api\n", "0.2.0"); err == nil {
		t.Error("bumpVersionGo on source without the constant should error")
	}
}

const sampleChangelog = `# Changelog

## [Unreleased]

## [0.1.0] - 2026-06-14

First release.

[Unreleased]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/Oblikovati/Oblikovati.API/releases/tag/v0.1.0
`

func TestRollChangelog(t *testing.T) {
	out, err := rollChangelog(sampleChangelog, "0.1.0", "0.2.0", "2026-06-15", "### Added\n\n- feat: x")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"## [Unreleased]",
		"## [0.2.0] - 2026-06-15",
		"- feat: x",
		"[Unreleased]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.2.0...HEAD",
		"[0.2.0]: https://github.com/Oblikovati/Oblikovati.API/compare/v0.1.0...v0.2.0",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("rolled changelog missing %q:\n%s", want, out)
		}
	}
	// The new version heading must sit below Unreleased, above the old release.
	if strings.Index(out, "## [0.2.0]") > strings.Index(out, "## [0.1.0]") {
		t.Errorf("0.2.0 section is not above 0.1.0:\n%s", out)
	}
	if _, err := rollChangelog("# Changelog\n", "0.1.0", "0.2.0", "d", ""); err == nil {
		t.Error("rollChangelog without an Unreleased heading should error")
	}
}
