// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

const (
	// MethodModelStatesGetPropertyTable reads the active part's property table.
	MethodModelStatesGetPropertyTable = "modelStates.getPropertyTable"
	// MethodModelStatesSetPropertyTable atomically replaces its displayed property columns and values.
	MethodModelStatesSetPropertyTable = "modelStates.setPropertyTable"
)

// ModelStatePropertyCell is one member's authored value. A nil Value inherits
// the column's DefaultValue; it is distinct from an empty string. MemberID must be
// a nonzero member identity in the active part. Example:
// `wire.ModelStatePropertyCell{MemberID: 1, Value: nil}`.
type ModelStatePropertyCell struct {
	MemberID uint64         `json:"memberId"`
	Value    *types.Variant `json:"value"`
}

// ModelStatePropertyColumn addresses a property and has exactly one cell per
// current member, in member order. DefaultValue is the document value inherited
// by members without an override; nil means the underlying property is absent.
// Columns are ordered and have unique keys.
// Example: `wire.ModelStatePropertyColumn{Property: key, Cells: cells}`.
type ModelStatePropertyColumn struct {
	Property     types.ModelStatePropertyKey `json:"property"`
	DefaultValue *types.Variant              `json:"defaultValue"`
	Cells        []ModelStatePropertyCell    `json:"cells"`
}

// ModelStatePropertyTableView is the active part's displayed property table.
// Generation identifies the sampled document generation; Members are in table
// order, AvailableProperties lists properties that can be selected, and Columns
// contains the displayed properties. RetainedColumns contains hidden properties
// with their preserved defaults and member values. A cell without a value is not a blank
// string. Example: `view, err := client.ModelStates().GetPropertyTable()`.
type ModelStatePropertyTableView struct {
	Generation          uint64                        `json:"generation"`
	Members             []ModelStateInfo              `json:"members"`
	AvailableProperties []types.ModelStatePropertyKey `json:"availableProperties"`
	Columns             []ModelStatePropertyColumn    `json:"columns"`
	RetainedColumns     []ModelStatePropertyColumn    `json:"retainedColumns"`
}

// SetModelStatePropertyTableArgs replaces the displayed property columns and
// their member values as one undoable document edit. If ExpectedGeneration is
// supplied it must equal the current document generation. Duplicate columns,
// missing or duplicate member cells, unknown properties, invalid typed values,
// stale generations, unit-bearing values, and altered default values are rejected
// without mutation. Values use the document-property types and carry no unit string.
// Removing a displayed
// column does not delete existing member property values. A canceled UI draft
// sends no set request. Example: `wire.SetModelStatePropertyTableArgs{Columns: columns}`.
type SetModelStatePropertyTableArgs struct {
	ExpectedGeneration *uint64                    `json:"expectedGeneration,omitempty"`
	Columns            []ModelStatePropertyColumn `json:"columns"`
}
