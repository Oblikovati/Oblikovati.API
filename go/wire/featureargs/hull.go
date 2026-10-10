// SPDX-License-Identifier: Apache-2.0

package featureargs

// The convex-hull feature kind (#1709): replaces the part's running solids with their single
// convex hull (OpenSCAD hull()). It takes no arguments, so Hull is a zero-field arg type — an
// add-in still builds it with a compile-checked struct rather than an empty raw blob.

const KindHull = "hull"

// Hull convex-hulls the part's running solid bodies into one body (KindHull). It has no fields.
type Hull struct{}

// Kind reports the feature kind Hull creates.
func (Hull) Kind() string { return KindHull }

// hullArgs is the hull family's contribution to [All].
var hullArgs = []Arg{Hull{}}
