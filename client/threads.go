// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Threads is the thread-table operation group (M09-F01 PBI-101, #325).
type Threads struct{ c *Client }

// Threads returns the thread-table operation group.
func (c *Client) Threads() Threads { return Threads{c} }

// TableQuery lists the thread tables progressively: thread types always; a
// type's nominal sizes, a size's designations, and a designation's classes as
// each filter is given, e.g.
// TableQuery(wire.ThreadTableQueryArgs{ThreadType: "ISO Metric profile"}).
func (t Threads) TableQuery(args wire.ThreadTableQueryArgs) (wire.ThreadTableQueryResult, error) {
	var r wire.ThreadTableQueryResult
	return r, t.c.call(wire.MethodThreadsTableQuery, args, &r)
}

// Resolve resolves a designation (with optional class / handedness / tapered
// flag) to its thread data, e.g.
// Resolve(wire.ResolveThreadArgs{Designation: "M8x1.25", Class: "6H", Internal: true}).
func (t Threads) Resolve(args wire.ResolveThreadArgs) (wire.ThreadInfoResult, error) {
	var r wire.ThreadInfoResult
	return r, t.c.call(wire.MethodThreadsResolve, args, &r)
}
