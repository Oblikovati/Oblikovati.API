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
	writeFile(t, dir, "version.go", sampleVersionGo)
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
	assertFileContains(t, filepath.Join(dir, "version.go"), `const Version = "0.2.0"`)
	assertFileContains(t, filepath.Join(dir, "CHANGELOG.md"), "## [0.2.0] - 2026-06-15")
	assertFileContains(t, notesPath, "- feat(assembly): batch placement")
}

// TestRunNoReleaseLeavesFiles: a docs/chore-only set bumps nothing and writes no output.
func TestRunNoReleaseLeavesFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "version.go", sampleVersionGo)
	writeFile(t, dir, "CHANGELOG.md", sampleChangelog)

	in := strings.NewReader("docs: readme\x00chore: deps")
	var out strings.Builder
	if err := run(true, dir, "2026-06-15", "", in, &out); err != nil {
		t.Fatal(err)
	}
	if out.String() != "" {
		t.Fatalf("expected no stdout for a no-release set, got %q", out.String())
	}
	assertFileContains(t, filepath.Join(dir, "version.go"), `const Version = "0.1.0"`)
}

func writeFile(t *testing.T, dir, name, content string) {
	t.Helper()
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
	writeFile(t, dir, "version.go", sampleVersionGo)
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
