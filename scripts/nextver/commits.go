// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"regexp"
	"strings"
)

// change is the release-relevant scope of a set of commits, ordered by precedence:
// a stronger scope in any one commit wins for the whole set.
type change int

const (
	changeNone     change = iota // docs/chore/ci/test/style/build only — no release
	changePatch                  // fix/perf/revert, or an unrecognized subject (safe floor)
	changeFeature                // feat — an additive surface change
	changeBreaking               // a "!" marker or a "BREAKING CHANGE:" footer
)

func (c change) String() string {
	switch c {
	case changeBreaking:
		return "breaking"
	case changeFeature:
		return "feature"
	case changePatch:
		return "patch"
	default:
		return "none"
	}
}

// headerRe matches a Conventional Commits subject "type(scope)!: description",
// capturing the type and the optional "!" breaking marker. The repo also uses the
// package name as the type (e.g. "wire: …"), which falls through to the default below.
var headerRe = regexp.MustCompile(`^([A-Za-z]+)(?:\([^)]*\))?(!)?:`)

// breakingFooterRe matches the Conventional Commits breaking-change footer on its own
// line ("BREAKING CHANGE:" or "BREAKING-CHANGE:").
var breakingFooterRe = regexp.MustCompile(`(?m)^BREAKING[ -]CHANGE:`)

// patchTypes are the conventional types that warrant a PATCH; featureTypes a feature;
// silentTypes nothing. An unrecognized type defaults to changePatch so the version
// never silently stalls on a real change — but additive surface changes should use
// "feat:" to earn the (minor) bump the policy gives them (see RELEASING.md).
var (
	featureTypes = map[string]bool{"feat": true}
	patchTypes   = map[string]bool{"fix": true, "perf": true, "revert": true}
	silentTypes  = map[string]bool{
		"docs": true, "chore": true, "ci": true, "test": true,
		"style": true, "build": true, "refactor": true,
	}
)

// classifyOne returns the release scope of a single commit message (full message, so
// the breaking-change footer in the body is seen).
func classifyOne(message string) change {
	if breakingFooterRe.MatchString(message) {
		return changeBreaking
	}
	m := headerRe.FindStringSubmatch(strings.TrimSpace(message))
	if m == nil {
		return changePatch // non-conventional subject — conservative floor, never None
	}
	if m[2] == "!" {
		return changeBreaking
	}
	t := strings.ToLower(m[1])
	switch {
	case featureTypes[t]:
		return changeFeature
	case patchTypes[t]:
		return changePatch
	case silentTypes[t]:
		return changeNone
	default:
		return changePatch // unknown type (e.g. "wire:") — treat as a patch, not a no-op
	}
}

// classify returns the strongest scope across all commit messages.
func classify(messages []string) change {
	strongest := changeNone
	for _, m := range messages {
		if c := classifyOne(m); c > strongest {
			strongest = c
		}
	}
	return strongest
}

// releaseNotes renders the commit subjects as a Keep a Changelog section body, grouping
// breaking/feat under "Changed"/"Added" and fixes under "Fixed". Unrecognized subjects
// land in "Changed" so nothing is dropped. Empty groups are omitted.
func releaseNotes(messages []string) string {
	added, changed, fixed := notesGroups(messages)
	var b strings.Builder
	writeGroup(&b, "Added", added)
	writeGroup(&b, "Changed", changed)
	writeGroup(&b, "Fixed", fixed)
	return strings.TrimRight(b.String(), "\n")
}

// notesGroups buckets each commit's subject line by its scope.
func notesGroups(messages []string) (added, changed, fixed []string) {
	for _, m := range messages {
		subject := strings.TrimSpace(strings.SplitN(strings.TrimSpace(m), "\n", 2)[0])
		if subject == "" {
			continue
		}
		switch classifyOne(m) {
		case changeFeature:
			added = append(added, subject)
		case changePatch:
			fixed = append(fixed, subject)
		case changeBreaking:
			changed = append(changed, subject)
		}
	}
	return added, changed, fixed
}

func writeGroup(b *strings.Builder, heading string, items []string) {
	if len(items) == 0 {
		return
	}
	fmt.Fprintf(b, "### %s\n\n", heading)
	for _, it := range items {
		fmt.Fprintf(b, "- %s\n", it)
	}
	b.WriteString("\n")
}
