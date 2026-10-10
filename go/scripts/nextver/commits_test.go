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
	}
	for _, c := range cases {
		if got := classifyOne(c.msg); got != c.want {
			t.Errorf("classifyOne(%q) = %s, want %s", c.msg, got, c.want)
		}
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
