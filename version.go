// SPDX-License-Identifier: Apache-2.0

// Package api is the root of the Oblikovati public API contract module
// (oblikovati.org/api). The contract itself lives in the subpackages — types,
// contract, wire, and client; this root package carries only the module's
// semantic version, the single in-repo source of truth a release is tagged from.
package api

import (
	"strconv"
	"strings"
)

// Version is the module's version, per Semantic Versioning 2.0.0
// (https://semver.org). It is the semver string WITHOUT the leading "v" that the
// git tag carries (the "v" is a Go-module tag convention, not part of semver), so
// the release tag for this value is "v" + Version.
//
// While the major version is 0 the API is in initial development: the public
// surface MAY change in any minor release and there is no backward-compatibility
// guarantee yet (semver §4). See RELEASING.md for how a release is cut.
const Version = "0.2.0"

// Major is the major component of [Version], per Semantic Versioning: the breaking
// -change boundary (semver §8). The host↔add-in load-time handshake refuses to load
// an add-in whose compiled-against major differs from the host's (see
// include/oblikovati_addin.h, ObkAddInApiMajor). Both sides derive it from [Version].
//
//	if addinMajor != api.Major() { /* incompatible — do not load */ }
func Major() int { return component(0) }

// Minor is the minor component of [Version]. Within a matching major, minor bumps are
// additive and backward-compatible (semver §7): a host satisfies any add-in built
// against the SAME OR AN OLDER minor, but an add-in built against a NEWER minor than
// the host needs API the host lacks, so the handshake refuses to load it
// (ObkAddInApiMinor).
//
//	if addinMajor == api.Major() && addinMinor > api.Minor() { /* too new — do not load */ }
func Minor() int { return component(1) }

// component parses the i-th dot-separated number of [Version] (0=major, 1=minor). It
// panics on a non-numeric component, which is unreachable while TestVersionIsSemver
// passes — better than silently returning 0, a real version number.
func component(i int) int {
	parts := strings.SplitN(Version, ".", 3)
	n, err := strconv.Atoi(parts[i])
	if err != nil {
		panic("api: Version " + strconv.Quote(Version) + " has a non-numeric component: " + err.Error())
	}
	return n
}
