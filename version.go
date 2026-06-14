// SPDX-License-Identifier: Apache-2.0

// Package api is the root of the Oblikovati public API contract module
// (oblikovati.org/api). The contract itself lives in the subpackages — types,
// contract, wire, and client; this root package carries only the module's
// semantic version, the single in-repo source of truth a release is tagged from.
package api

// Version is the module's version, per Semantic Versioning 2.0.0
// (https://semver.org). It is the semver string WITHOUT the leading "v" that the
// git tag carries (the "v" is a Go-module tag convention, not part of semver), so
// the release tag for this value is "v" + Version.
//
// While the major version is 0 the API is in initial development: the public
// surface MAY change in any minor release and there is no backward-compatibility
// guarantee yet (semver §4). See RELEASING.md for how a release is cut.
const Version = "0.1.0"
