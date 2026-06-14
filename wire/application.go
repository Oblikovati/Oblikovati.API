// SPDX-License-Identifier: Apache-2.0

package wire

// ApplicationApiVersionResult is the response of [MethodApplicationApiVersion]: the
// semantic version of the api contract the running host implements. Version is the
// full "MAJOR.MINOR.PATCH" string (== api.Version, no leading "v"); Major and Minor
// are broken out so an add-in can gate compatibility without parsing.
//
// The load-time handshake (ObkAddInApiMajor/ObkAddInApiMinor in
// include/oblikovati_addin.h) already guarantees the add-in's major matches and its
// minor is not newer than the host before it is loaded; this runtime query lets a
// loaded add-in additionally adapt to the exact minor/patch within that major.
type ApplicationApiVersionResult struct {
	Version string `json:"version"`
	Major   int    `json:"major"`
	Minor   int    `json:"minor"`
}
