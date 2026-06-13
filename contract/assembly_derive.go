// SPDX-License-Identifier: Apache-2.0

package contract

// DerivedAssemblyComponent is a part feature that pulls a source assembly's geometry
// into the part as a base body and tracks it associatively — the reference API's
// DerivedAssemblyComponent (M11-F06, Oblikovati/Oblikovati#631). A shrinkwrap is the
// simplified flavor of the same contract. The host implementation lives in /source
// (model/feature); this is the scalar surface a consumer reads without the geometry.
type DerivedAssemblyComponent interface {
	// Kind names the feature type ("derivedAssembly" or "shrinkwrap").
	Kind() string
	// Linked reports whether the feature still pulls from its source assembly.
	Linked() bool
	// SourceVersion returns the source assembly's current geometry version, so a
	// consumer can tell when the derived result is stale.
	SourceVersion() string
	// BreakLink freezes the current derived geometry and severs the source link, so
	// the part keeps the result without further updates.
	BreakLink() error
}
