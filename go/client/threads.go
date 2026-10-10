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
//
// mcp:tool threads_table_query
// mcp:summary Lists the thread tables progressively: thread types always; a type's nominal sizes, a size's designations, and a designation's classes as each filter is given, e.g.
func (t Threads) TableQuery(args wire.ThreadTableQueryArgs) (wire.ThreadTableQueryResult, error) {
	return call[wire.ThreadTableQueryResult](t.c, wire.MethodThreadsTableQuery, args)
}

// Resolve resolves a designation (with optional class / handedness / tapered
// flag) to its thread data, e.g.
// Resolve(wire.ResolveThreadArgs{Designation: "M8x1.25", Class: "6H", Internal: true}).
//
// mcp:tool threads_resolve
// mcp:summary Resolves a designation (with optional class / handedness / tapered flag) to its thread data, e.g.
func (t Threads) Resolve(args wire.ResolveThreadArgs) (wire.ThreadInfoResult, error) {
	return call[wire.ThreadInfoResult](t.c, wire.MethodThreadsResolve, args)
}
