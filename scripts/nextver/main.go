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
	dir := flag.String("dir", ".", "module root holding version.go and CHANGELOG.md")
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
	verPath := filepath.Join(dir, "version.go")
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
	fmt.Fprintf(out, "next=%s\nbump=%s\n", next, scope)
	return nil
}

// applyEdits rewrites version.go and CHANGELOG.md to newVer and, when notesOut is set,
// writes the rendered notes there for the release body.
func applyEdits(dir, verSrc, prevVer, newVer, date string, messages []string, notesOut string) error {
	bumped, err := bumpVersionGo(verSrc, newVer)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, "version.go"), []byte(bumped), 0o644); err != nil {
		return fmt.Errorf("write version.go: %w", err)
	}
	notes := releaseNotes(messages)
	if err := rollChangelogFile(dir, prevVer, newVer, date, notes); err != nil {
		return err
	}
	if notesOut != "" {
		if err := os.WriteFile(notesOut, []byte(notes+"\n"), 0o644); err != nil {
			return fmt.Errorf("write notes %s: %w", notesOut, err)
		}
	}
	return nil
}

// rollChangelogFile reads, rolls, and writes CHANGELOG.md in dir.
func rollChangelogFile(dir, prevVer, newVer, date, notes string) error {
	path := filepath.Join(dir, "CHANGELOG.md")
	src, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}
	rolled, err := rollChangelog(string(src), prevVer, newVer, date, notes)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(rolled), 0o644)
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
