// SPDX-License-Identifier: Apache-2.0

package main

import (
	"strings"
	"testing"
)

func TestClassifyOne(t *testing.T) {
	cases := []struct {
		msg  string
		want change
	}{
		{"feat(assembly): batch place-by-definition", changeFeature},
		{"feat: add thing", changeFeature},
		{"fix(api): mcp:input override", changePatch},
		{"perf: faster tessellation", changePatch},
		{"docs: tidy readme", changeNone},
		{"chore: bump deps", changeNone},
		{"ci: bump actions", changeNone},
		{"refactor: split function", changeNone},
		{"feat!: drop legacy field", changeBreaking},
		{"feat(wire)!: rename DTO", changeBreaking},
		{"fix: x\n\nBREAKING CHANGE: removed Foo", changeBreaking},
		{"types: add ChamferType", changePatch}, // package-name "type" → conservative floor
		{"Add M20 F16–F20 API parity contracts", changePatch},
		{"M26: add commandLine.submit", changePatch},
		{"fix(release): correct scope\n\nRelease-Scope: feature", changePatch},
		{"docs: mention Release-Scope: feature in prose", changeNone},
		{"fix: wrong trailer\n\nRelease-Scope: features", changePatch},
		{"fix: breaking change\n\nRelease-Scope: feature\nBREAKING CHANGE: removed Foo", changeBreaking},
	}
	for _, c := range cases {
		if got := classifyOne(c.msg); got != c.want {
			t.Errorf("classifyOne(%q) = %s, want %s", c.msg, got, c.want)
		}
	}
}

func TestExplicitAdditiveScopeCorrectsUnconventionalFeatureCommits(t *testing.T) {
	messages := []string{
		"Report nonfatal warnings for saved documents",
		"Define express assembly loading operations",
		"fix(release): correct additive scope\n\nRelease-Scope: feature",
	}
	scope := classify(messages)
	version, release := nextVersion(semver{0, 159, 0}, scope)
	if !release || version.String() != "0.160.0" {
		t.Fatalf("scope %s version %s release %t: expected feature release 0.160.0", scope, version, release)
	}
	for _, header := range []string{"feat!: remove field", "fix(api)!: remove field"} {
		if got := classify([]string{header + "\n\nRelease-Scope: feature"}); got != changeBreaking {
			t.Fatalf("header %q classified as %s: expected breaking", header, got)
		}
	}
}

func TestReleaseScopeRejectsUnsupportedValues(t *testing.T) {
	for _, message := range []string{"fix: x\n\nRelease-Scope: minor", "fix: x\n\nRelease-Scope:"} {
		if err := validateReleaseScopes([]string{message}); err == nil {
			t.Fatalf("release directive %q: expected rejection", message)
		}
	}
	if err := validateReleaseScopes([]string{"fix: x\n\nRelease-Scope: feature\r\n"}); err != nil {
		t.Fatal(err)
	}
}

func TestClassifyTakesStrongest(t *testing.T) {
	msgs := []string{"docs: x", "fix: y", "feat: z", "chore: w"}
	if got := classify(msgs); got != changeFeature {
		t.Fatalf("classify = %s, want feature", got)
	}
	if got := classify([]string{"docs: a", "ci: b"}); got != changeNone {
		t.Fatalf("classify(silent only) = %s, want none", got)
	}
	if got := classify(nil); got != changeNone {
		t.Fatalf("classify(nil) = %s, want none", got)
	}
}

func TestReleaseNotesGroups(t *testing.T) {
	notes := releaseNotes([]string{
		"feat(a): add A\n\nbody ignored",
		"fix(b): fix B",
		"feat!: break C",
		"docs: nope",
	})
	for _, want := range []string{
		"### Added", "- feat(a): add A",
		"### Fixed", "- fix(b): fix B",
		"### Changed", "- feat!: break C",
	} {
		if !strings.Contains(notes, want) {
			t.Errorf("notes missing %q:\n%s", want, notes)
		}
	}
	if strings.Contains(notes, "nope") {
		t.Errorf("docs commit leaked into notes:\n%s", notes)
	}
	if strings.Contains(notes, "body ignored") {
		t.Errorf("notes used body instead of subject:\n%s", notes)
	}
}

func TestReleaseScopeDoesNotRelabelCorrectionInNotes(t *testing.T) {
	notes := releaseNotes([]string{"fix(release): correct scope\n\nRelease-Scope: feature"})
	if !strings.Contains(notes, "### Fixed\n\n- fix(release): correct scope") || strings.Contains(notes, "### Added") {
		t.Fatalf("release notes %q: expected correction under Fixed", notes)
	}
}

func TestUnconventionalSubjectsAreChangedInReleaseNotes(t *testing.T) {
	messages := []string{"Define express assembly loading operations", "wire: extend save results"}
	if got := classify(messages); got != changePatch {
		t.Fatalf("version scope %s: expected conservative patch floor", got)
	}
	for _, message := range messages {
		if got := noteCategory(message); got != noteChanged {
			t.Fatalf("note category %v for %q: expected changed", got, message)
		}
	}
	notes := releaseNotes(messages)
	if !strings.Contains(notes, "### Changed\n\n- Define express assembly loading operations\n- wire: extend save results") || strings.Contains(notes, "### Fixed") {
		t.Fatalf("release notes %q: expected both subjects under Changed", notes)
	}
}
