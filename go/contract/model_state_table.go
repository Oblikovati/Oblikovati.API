// SPDX-License-Identifier: Apache-2.0

package contract

// ModelStateTable exposes the active document's ordered model-state table.
// Example: `count := table.ColumnCount()`.
type ModelStateTable interface {
	// ColumnCount reports the number of visible columns across all categories.
	ColumnCount() int
}
