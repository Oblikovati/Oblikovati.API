// SPDX-License-Identifier: Apache-2.0

package types

// ModelStateTableColumnKind selects the owner of a model-state table column.
// Example: `types.ModelStateTablePropertyColumn` identifies a property column.
type ModelStateTableColumnKind int32

const (
	// ModelStateTablePropertyColumn identifies a document-property column.
	ModelStateTablePropertyColumn ModelStateTableColumnKind = 1
	// ModelStateTableSuppressionColumn identifies a feature-suppression column.
	ModelStateTableSuppressionColumn ModelStateTableColumnKind = 2
	// ModelStateTableParameterColumn identifies a parameter-expression column.
	ModelStateTableParameterColumn ModelStateTableColumnKind = 3
)
