// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Derived parameter tables (M02-F06, Oblikovati/Oblikovati#605), on the
// Parameters operation group.

// ListDerivedTables returns the active part's derived parameter tables with
// their links, candidates, and health.
func (p Parameters) ListDerivedTables() (wire.ListDerivedParameterTablesResult, error) {
	var r wire.ListDerivedParameterTablesResult
	return r, p.c.call(wire.MethodParametersDerivedTablesList, nil, &r)
}

// AddDerivedTable links parameters from another document into this one,
// returning the created table.
func (p Parameters) AddDerivedTable(args wire.DerivedParameterTableAddArgs) (wire.DerivedParameterTableInfo, error) {
	var r wire.DerivedParameterTableInfo
	return r, p.c.call(wire.MethodParametersDerivedTablesAdd, args, &r)
}

// SetDerivedTableLinked replaces a table's linked subset — newly linked names
// gain derived parameters, unlinked ones lose theirs — and returns the
// updated table.
func (p Parameters) SetDerivedTableLinked(args wire.DerivedParameterTableSetLinkedArgs) (wire.DerivedParameterTableInfo, error) {
	var r wire.DerivedParameterTableInfo
	return r, p.c.call(wire.MethodParametersDerivedTablesSetLinked, args, &r)
}

// DeleteDerivedTable removes a table and its derived parameters. A table
// owned by a derived component cannot be deleted directly.
func (p Parameters) DeleteDerivedTable(id int) error {
	return p.c.call(wire.MethodParametersDerivedTablesDelete, wire.DerivedParameterTableDeleteArgs{ID: id}, nil)
}
