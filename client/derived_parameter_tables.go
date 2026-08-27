// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Derived parameter tables (M02-F06, Oblikovati/Oblikovati#605), on the
// Parameters operation group.

// ListDerivedTables returns the active document's derived parameter tables (part or assembly) with
// their links, candidates, and health.
//
// mcp:tool parameters_derived_tables_list
// mcp:summary Returns the active document's derived parameter tables (part or assembly) with their links, candidates, and health.
func (p Parameters) ListDerivedTables() (wire.ListDerivedParameterTablesResult, error) {
	return call[wire.ListDerivedParameterTablesResult](p.c, wire.MethodParametersDerivedTablesList, nil)
}

// AddDerivedTable links parameters from another document into this one,
// returning the created table.
//
// mcp:tool parameters_derived_tables_add
// mcp:summary Links parameters from another document into this one, returning the created table.
func (p Parameters) AddDerivedTable(args wire.DerivedParameterTableAddArgs) (wire.DerivedParameterTableInfo, error) {
	return call[wire.DerivedParameterTableInfo](p.c, wire.MethodParametersDerivedTablesAdd, args)
}

// SetDerivedTableLinked replaces a table's linked subset — newly linked names
// gain derived parameters, unlinked ones lose theirs — and returns the
// updated table.
//
// mcp:tool parameters_derived_tables_set_linked
// mcp:summary Replaces a table's linked subset — newly linked names gain derived parameters, unlinked ones lose theirs — and returns the updated table.
func (p Parameters) SetDerivedTableLinked(args wire.DerivedParameterTableSetLinkedArgs) (wire.DerivedParameterTableInfo, error) {
	return call[wire.DerivedParameterTableInfo](p.c, wire.MethodParametersDerivedTablesSetLinked, args)
}

// DeleteDerivedTable removes a table and its derived parameters. A table
// owned by a derived component cannot be deleted directly.
//
// mcp:tool parameters_derived_tables_delete
// mcp:summary Removes a table and its derived parameters.
func (p Parameters) DeleteDerivedTable(id int) error {
	return p.c.invoke(wire.MethodParametersDerivedTablesDelete, wire.DerivedParameterTableDeleteArgs{ID: id}, nil)
}
