// SPDX-License-Identifier: Apache-2.0

package contract

// ModelStatePropertyTable exposes the currently displayed property-column count.
// Example: `count := table.PropertyColumnCount()`.
type ModelStatePropertyTable interface {
	// PropertyColumnCount reports the number of displayed property columns.
	PropertyColumnCount() int
}
