// SPDX-License-Identifier: Apache-2.0

package api

import (
	"regexp"
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
