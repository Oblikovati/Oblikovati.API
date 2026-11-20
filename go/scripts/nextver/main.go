// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// main reads the merged commit messages from stdin (NUL-separated, as
// `git log --format=%B%x00` emits), computes the next version from their scope, and
// prints "next=X.Y.Z" + "bump=<scope>" when a release is due. With -apply it also
// rewrites version.go and CHANGELOG.md, and writes the release notes to -notes-out.
// No output (exit 0) means nothing releasable changed.
func main() {
	apply := flag.Bool("apply", false, "rewrite version.go and CHANGELOG.md in place")
	dir := flag.String("dir", ".", "repository root holding go/version.go and CHANGELOG.md")
	date := flag.String("date", time.Now().UTC().Format("2006-01-02"), "release date (YYYY-MM-DD)")
	notesOut := flag.String("notes-out", "", "also write the rendered release notes to this file")
	flag.Parse()

	if err := run(*apply, *dir, *date, *notesOut, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "nextver:", err)
		os.Exit(1)
	}
}

// run is main's testable core: it computes (and optionally applies) the next version,
// writing the result to out.
func run(apply bool, dir, date, notesOut string, in io.Reader, out io.Writer) error {
	messages := readMessages(in)
	if err := validateReleaseScopes(messages); err != nil {
		return err
	}
	verPath := filepath.Join(dir, "go", "version.go")
	verSrc, err := os.ReadFile(verPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", verPath, err)
	}
	curStr, err := readVersionConst(string(verSrc))
	if err != nil {
		return err
	}
	cur, err := parseSemver(curStr)
	if err != nil {
		return err
	}
	scope := classify(messages)
	next, release := nextVersion(cur, scope)
	if !release {
		fmt.Fprintln(os.Stderr, "nextver: no releasable commits; current version stands")
		return nil
	}
	if apply {
		if err := applyEdits(dir, string(verSrc), curStr, next.String(), date, messages, notesOut); err != nil {
			return err
		}
	}
	// The release workflow parses this line to decide the tag it pushes, so a failed write has to
	// fail the step. Swallowing it would let the job carry on and read an empty version.
	if _, err := fmt.Fprintf(out, "next=%s\nbump=%s\n", next, scope); err != nil {
		return fmt.Errorf("nextver: writing the version output failed: %w", err)
	}
	return nil
}

// applyEdits rewrites version.go and CHANGELOG.md to newVer and, when notesOut is set,
// writes the rendered notes there for the release body.
func applyEdits(dir, verSrc, prevVer, newVer, date string, messages []string, notesOut string) error {
	bumped, err := bumpVersionGo(verSrc, newVer)
	if err != nil {
		return err
	}
	notes := releaseNotes(messages)
	rolled, err := prepareChangelog(dir, prevVer, newVer, date, notes)
	if err != nil {
		return err
	}
	published, err := releaseSection(rolled, prevVer, newVer, date)
	if err != nil {
		return err
	}
	return writeReleaseFiles(dir, bumped, rolled, notesOut, published)
}

// writeReleaseFiles persists prevalidated release inputs and the public notes.
func writeReleaseFiles(dir, bumped, rolled, notesOut, published string) error {
	if err := os.WriteFile(filepath.Join(dir, "go", "version.go"), []byte(bumped), 0o644); err != nil {
		return fmt.Errorf("write version.go: %w", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "CHANGELOG.md"), []byte(rolled), 0o644); err != nil {
		return fmt.Errorf("write CHANGELOG.md: %w", err)
	}
	if notesOut != "" {
		if err := os.WriteFile(notesOut, []byte(published+"\n"), 0o644); err != nil {
			return fmt.Errorf("write notes %s: %w", notesOut, err)
		}
	}
	return nil
}

// prepareChangelog rolls the changelog without changing the on-disk source.
func prepareChangelog(dir, prevVer, newVer, date, notes string) (string, error) {
	path := filepath.Join(dir, "CHANGELOG.md")
	src, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", path, err)
	}
	rolled, err := rollChangelog(string(src), prevVer, newVer, date, notes)
	if err != nil {
		return "", err
	}
	return rolled, nil
}

// releaseSection selects the exact new section without including an older release.
func releaseSection(rolled, prevVer, newVer, date string) (string, error) {
	heading := fmt.Sprintf("## [%s] - %s\n", newVer, date)
	previous := "\n## [" + prevVer + "]"
	if strings.Count(rolled, heading) != 1 || strings.Count(rolled, previous) != 1 {
		return "", fmt.Errorf("changelog releases %q and %q: expected one heading for each", newVer, prevVer)
	}
	_, after, _ := strings.Cut(rolled, heading)
	end := strings.Index(after, previous)
	if end < 0 {
		return "", fmt.Errorf("changelog release %q: expected previous release %q after it", newVer, prevVer)
	}
	return strings.TrimSpace(after[:end]), nil
}

// readMessages splits NUL-separated commit messages from in, dropping empties.
func readMessages(in io.Reader) []string {
	raw, _ := io.ReadAll(in)
	var msgs []string
	for _, m := range strings.Split(string(raw), "\x00") {
		if strings.TrimSpace(m) != "" {
			msgs = append(msgs, m)
		}
	}
	return msgs
}
