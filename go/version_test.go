// SPDX-License-Identifier: Apache-2.0

package api

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// semverRe is the official Semantic Versioning 2.0.0 grammar from
// https://semver.org (the "without v-prefix" form), anchored. api.Version must
// match it so the release tag "v"+Version is a valid Go module version.
var semverRe = regexp.MustCompile(
	`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)` +
		`(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?` +
		`(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// TestVersionIsSemver guards the release contract: a malformed Version would make
// the auto-tagged "v"+Version an invalid module version, so the build fails here
// first with the offending value.
func TestVersionIsSemver(t *testing.T) {
	if !semverRe.MatchString(Version) {
		t.Fatalf("api.Version = %q is not valid semver per https://semver.org (expected MAJOR.MINOR.PATCH with no leading \"v\"); the release tag is \"v\"+Version", Version)
	}
}

// TestMajorMinorMatchVersion guards the load-time handshake: Major()/Minor() must
// equal the leading components of Version, since both host and add-in gate on them.
func TestMajorMinorMatchVersion(t *testing.T) {
	parts := strings.SplitN(Version, ".", 3)
	wantMajor, err := strconv.Atoi(parts[0])
	if err != nil {
		t.Fatalf("Version %q has a non-numeric major: %v", Version, err)
	}
	wantMinor, err := strconv.Atoi(parts[1])
	if err != nil {
		t.Fatalf("Version %q has a non-numeric minor: %v", Version, err)
	}
	if got := Major(); got != wantMajor {
		t.Fatalf("Major() = %d, want %d (from Version %q)", got, wantMajor, Version)
	}
	if got := Minor(); got != wantMinor {
		t.Fatalf("Minor() = %d, want %d (from Version %q)", got, wantMinor, Version)
	}
}
