// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Assembly is the assembly derive/shrinkwrap operation group (M11-F06,
// Oblikovati/Oblikovati#631/#716): pull a source assembly into the active part as a
// base body — optionally simplified — and break the link to freeze it. Each create
// returns the new feature's detail; break-link returns the refreshed detail.
type Assembly struct{ c *Client }

// Assembly returns the assembly derive/shrinkwrap operation group.
func (c *Client) Assembly() Assembly { return Assembly{c} }

// DeriveCreate derives the open assembly document source into the active part as a
// base body, e.g. DeriveCreate(assemblyDocID).
func (a Assembly) DeriveCreate(source uint64) (wire.FeatureDetailResult, error) {
	var r wire.FeatureDetailResult
	return r, a.c.call(wire.MethodAssemblyDeriveCreate, wire.DeriveCreateArgs{Source: source}, &r)
}

// ShrinkwrapCreate derives the open assembly document into the active part as a
// simplified, lightweight base body per the given removal/envelope options, e.g.
// ShrinkwrapCreate(wire.ShrinkwrapCreateArgs{Source: id, EnvelopeStyle: types.EnvelopeWhole}).
func (a Assembly) ShrinkwrapCreate(args wire.ShrinkwrapCreateArgs) (wire.FeatureDetailResult, error) {
	var r wire.FeatureDetailResult
	return r, a.c.call(wire.MethodAssemblyShrinkwrapCreate, args, &r)
}

// DeriveBreakLink freezes and severs the source link of the derived-assembly or
// shrinkwrap feature with the given id, e.g. DeriveBreakLink(featureID).
func (a Assembly) DeriveBreakLink(id uint64) (wire.FeatureDetailResult, error) {
	var r wire.FeatureDetailResult
	return r, a.c.call(wire.MethodAssemblyDeriveBreakLink, wire.DeriveBreakLinkArgs{ID: id}, &r)
}

// DeriveStatus reports whether the derive feature with the given id is out of date
// relative to its source document (its drive state), e.g. DeriveStatus(featureID).
func (a Assembly) DeriveStatus(id uint64) (wire.DeriveStatusResult, error) {
	var r wire.DeriveStatusResult
	return r, a.c.call(wire.MethodAssemblyDeriveStatus, wire.DeriveStatusArgs{ID: id}, &r)
}

// DeriveUpdate re-syncs the derive feature with the given id to its source's current
// revision, clearing its out-of-date state, e.g. DeriveUpdate(featureID).
func (a Assembly) DeriveUpdate(id uint64) (wire.DeriveStatusResult, error) {
	var r wire.DeriveStatusResult
	return r, a.c.call(wire.MethodAssemblyDeriveUpdate, wire.DeriveStatusArgs{ID: id}, &r)
}
