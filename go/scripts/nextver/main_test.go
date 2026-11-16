// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestRunAppliesBump drives the whole tool: a feature commit on a 0.1.0 tree should
// produce 0.2.0, rewrite version.go and CHANGELOG.md, and emit the release notes file.
func TestRunAppliesBump(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "go/version.go", sampleVersionGo)
	writeFile(t, dir, "CHANGELOG.md", sampleChangelog)
	notesPath := filepath.Join(dir, "notes.md")

	in := strings.NewReader("feat(assembly): batch placement\x00fix: tidy\x00docs: ignore")
	var out strings.Builder
	if err := run(true, dir, "2026-06-15", notesPath, in, &out); err != nil {
		t.Fatal(err)
	}
	if got := out.String(); !strings.Contains(got, "next=0.2.0") || !strings.Contains(got, "bump=feature") {
		t.Fatalf("run output = %q; want next=0.2.0 / bump=feature", got)
	}
	assertFileContains(t, filepath.Join(dir, "go", "version.go"), `const Version = "0.2.0"`)
	assertFileContains(t, filepath.Join(dir, "CHANGELOG.md"), "## [0.2.0] - 2026-06-15")
	assertFileContains(t, notesPath, "- feat(assembly): batch placement")
}

func TestRunPublishesTheCompleteRolledReleaseSection(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "go/version.go", sampleVersionGo)
	manual := "### Changed\n\n- Existing callers must supply the new field.\n"
	changelog := strings.Replace(sampleChangelog, "## [Unreleased]\n", "## [Unreleased]\n\n"+manual, 1)
	writeFile(t, dir, "CHANGELOG.md", changelog)
	notesPath := filepath.Join(dir, "notes.md")
	if err := run(true, dir, "2026-06-15", notesPath, strings.NewReader("feat: add field"), &strings.Builder{}); err != nil {
		t.Fatal(err)
	}
	notes, err := os.ReadFile(notesPath)
	if err != nil {
		t.Fatal(err)
	}
	for _, value := range []string{"- feat: add field", "- Existing callers must supply the new field."} {
		if strings.Count(string(notes), value) != 1 {
			t.Errorf("notes %q: expected %q exactly once", notes, value)
		}
	}
	if strings.Contains(string(notes), "First release.") || strings.Contains(string(notes), "[Unreleased]:") {
		t.Errorf("notes %q: expected only the new release content", notes)
	}
	rolled, err := os.ReadFile(filepath.Join(dir, "CHANGELOG.md"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(string(rolled), "- Existing callers must supply the new field.") != 1 {
		t.Errorf("rolled changelog %q: expected the manual note exactly once", rolled)
	}
	section, err := releaseSection(string(rolled), "0.1.0", "0.2.0", "2026-06-15")
	if err != nil || strings.TrimSpace(string(notes)) != section {
		t.Errorf("published notes %q do not match rolled section %q: %v", notes, section, err)
	}
}

func TestRunRejectsMissingPreviousSectionBeforeEditingVersion(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "go/version.go", sampleVersionGo)
	writeFile(t, dir, "CHANGELOG.md", strings.Replace(sampleChangelog, "## [0.1.0]", "## [0.0.9]", 1))
	var out strings.Builder
	err := run(true, dir, "2026-06-15", "", strings.NewReader("feat: add field"), &out)
	if err == nil || !strings.Contains(err.Error(), "0.1.0") {
		t.Fatalf("error %v: expected missing previous section", err)
	}
	assertFileContains(t, filepath.Join(dir, "go", "version.go"), `const Version = "0.1.0"`)
}

func TestReleaseSectionRejectsAmbiguousOrOutOfOrderHeadings(t *testing.T) {
	for _, changelog := range []string{
		"## [0.2.0] - 2026-06-15\nNew\n## [0.1.0] - old\n## [0.1.0] - duplicate\n",
		"## [0.1.0] - old\n## [0.2.0] - 2026-06-15\nNew\n",
	} {
		if section, err := releaseSection(changelog, "0.1.0", "0.2.0", "2026-06-15"); err == nil {
			t.Errorf("changelog %q produced %q: expected a heading error", changelog, section)
		}
	}
}

// TestRunNoReleaseLeavesFiles: a docs/chore-only set bumps nothing and writes no output.
func TestRunNoReleaseLeavesFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "go/version.go", sampleVersionGo)
	writeFile(t, dir, "CHANGELOG.md", sampleChangelog)

	in := strings.NewReader("docs: readme\x00chore: deps")
	var out strings.Builder
	if err := run(true, dir, "2026-06-15", "", in, &out); err != nil {
		t.Fatal(err)
	}
	if out.String() != "" {
		t.Fatalf("expected no stdout for a no-release set, got %q", out.String())
	}
	assertFileContains(t, filepath.Join(dir, "go", "version.go"), `const Version = "0.1.0"`)
}

func TestRunRejectsInvalidScopeBeforeEditingFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "go/version.go", sampleVersionGo)
	writeFile(t, dir, "CHANGELOG.md", sampleChangelog)
	var out strings.Builder
	err := run(true, dir, "2026-06-15", "", strings.NewReader("fix: x\n\nRelease-Scope: minor"), &out)
	if err == nil || !strings.Contains(err.Error(), "Release-Scope: minor") {
		t.Fatalf("error %v: expected invalid scope diagnostic", err)
	}
	assertFileContains(t, filepath.Join(dir, "go", "version.go"), `const Version = "0.1.0"`)
	if out.String() != "" {
		t.Fatalf("output %q: expected no version", out.String())
	}
}

func writeFile(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(filepath.Join(dir, name)), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func assertFileContains(t *testing.T, path, want string) {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), want) {
		t.Errorf("%s missing %q:\n%s", path, want, b)
	}
}

// failingWriter is a named fake stdout that always errors, standing in for a closed pipe or a
// full disk when the release workflow reads the tool's output.
type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) { return 0, errors.New("pipe closed") }

// TestRunReportsAFailedWrite: the release workflow parses "next=" off stdout to decide the tag it
// pushes. A swallowed write error would let the job carry on and read an empty version, tagging
// nothing or the wrong thing — so the failure has to surface.
func TestRunReportsAFailedWrite(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "go/version.go", sampleVersionGo)
	writeFile(t, dir, "CHANGELOG.md", sampleChangelog)

	in := strings.NewReader("feat(assembly): batch placement")
	err := run(true, dir, "2026-06-15", "", in, failingWriter{})

	if err == nil {
		t.Fatal("a failed write of the version output must be reported")
	}
	if !strings.Contains(err.Error(), "pipe closed") {
		t.Errorf("error %q should carry the underlying write failure", err)
	}
}
