// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

const (
	// MethodModelStatesGetTable reads the active part's complete model-state table.
	MethodModelStatesGetTable = "modelStates.getTable"
	// MethodModelStatesVerifyTable checks a draft without changing document state.
	MethodModelStatesVerifyTable = "modelStates.verifyTable"
	// MethodModelStatesSetTable applies a complete model-state table as one edit.
	MethodModelStatesSetTable = "modelStates.setTable"
)

// ModelStateTableRowRef identifies an existing member or a new draft row, never both.
// DraftKey is nonempty and unique within one set request; MemberID is nonzero for
// an existing member. Example: `wire.ModelStateTableRowRef{MemberID: 1}`.
type ModelStateTableRowRef struct {
	MemberID uint64 `json:"memberId,omitempty"`
	DraftKey string `json:"draftKey,omitempty"`
}

// ModelStateTableRow is an ordered member row. New rows require Source to refer
// to an existing row or an earlier new row; existing rows require Source to be
// nil. The built-in Primary member stays first and cannot be renamed or removed. Example:
// `wire.ModelStateTableRow{Ref: wire.ModelStateTableRowRef{MemberID: 1}, Name: "Primary"}`.
type ModelStateTableRow struct {
	Ref    ModelStateTableRowRef  `json:"ref"`
	Name   string                 `json:"name"`
	Source *ModelStateTableRowRef `json:"source,omitempty"`
}

// ModelStateTableColumnRef identifies one visible column. Exactly one identity
// must match Kind. Example: `wire.ModelStateTableColumnRef{Kind: types.ModelStateTableSuppressionColumn, FeatureID: 1}`.
type ModelStateTableColumnRef struct {
	Kind        types.ModelStateTableColumnKind `json:"kind"`
	Property    *types.ModelStatePropertyKey    `json:"property,omitempty"`
	FeatureID   uint64                          `json:"featureId,omitempty"`
	ParameterID uint64                          `json:"parameterId,omitempty"`
}

// ModelStateTablePropertyCell holds a property override; nil inherits the base.
// An authored empty string is distinct from nil. Example:
// `wire.ModelStateTablePropertyCell{Row: wire.ModelStateTableRowRef{MemberID: 1}}`.
type ModelStateTablePropertyCell struct {
	Row   ModelStateTableRowRef `json:"row"`
	Value *types.Variant        `json:"value"`
}

// ModelStateTableProperty holds one visible or retained property and one cell per row.
// Base is the inherited document value, or nil when absent. Example:
// `wire.ModelStateTableProperty{Property: key, Cells: cells}`.
type ModelStateTableProperty struct {
	Property types.ModelStatePropertyKey   `json:"property"`
	Base     *types.Variant                `json:"base"`
	Cells    []ModelStateTablePropertyCell `json:"cells"`
}

// ModelStateTableSuppressionCell holds a feature suppression override; nil
// inherits the authored feature state. Example:
// `wire.ModelStateTableSuppressionCell{Row: wire.ModelStateTableRowRef{MemberID: 1}}`.
type ModelStateTableSuppressionCell struct {
	Row        ModelStateTableRowRef `json:"row"`
	Suppressed *bool                 `json:"suppressed"`
}

// ModelStateTableSuppression holds a feature and one suppression cell per row.
// Base is the authored feature suppression state. Example:
// `wire.ModelStateTableSuppression{FeatureID: 1, Cells: cells}`.
type ModelStateTableSuppression struct {
	FeatureID uint64                           `json:"featureId"`
	Base      bool                             `json:"base"`
	Cells     []ModelStateTableSuppressionCell `json:"cells"`
}

// ModelStateTableParameterCell holds an expression override; nil inherits Base.
// Expressions are authored text, not evaluated numbers. Example:
// `wire.ModelStateTableParameterCell{Row: wire.ModelStateTableRowRef{MemberID: 1}}`.
type ModelStateTableParameterCell struct {
	Row        ModelStateTableRowRef `json:"row"`
	Expression *string               `json:"expression"`
}

// ModelStateTableParameter holds a parameter and one expression cell per row.
// Base is the inherited authored expression. Example:
// `wire.ModelStateTableParameter{ParameterID: 1, Base: "10 mm", Cells: cells}`.
type ModelStateTableParameter struct {
	ParameterID uint64                         `json:"parameterId"`
	Base        string                         `json:"base"`
	Cells       []ModelStateTableParameterCell `json:"cells"`
}

// ModelStateTablePropertyChoice names a selectable document property and its
// inherited base, or nil when the property is absent. Example:
// `wire.ModelStateTablePropertyChoice{Property: key}`.
type ModelStateTablePropertyChoice struct {
	Property types.ModelStatePropertyKey `json:"property"`
	Base     *types.Variant              `json:"base"`
}

// ModelStateTableFeatureChoice names a selectable authored feature.
// Example: `wire.ModelStateTableFeatureChoice{ID: 1, Name: "Extrusion1"}`.
type ModelStateTableFeatureChoice struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	BaseSuppressed bool   `json:"baseSuppressed"`
}

// ModelStateTableParameterChoice names a selectable parameter and its feature,
// if any. Example: `wire.ModelStateTableParameterChoice{ID: 1, Name: "d0"}`.
type ModelStateTableParameterChoice struct {
	ID             uint64 `json:"id"`
	FeatureID      uint64 `json:"featureId,omitempty"`
	Name           string `json:"name"`
	Unit           string `json:"unit"`
	BaseExpression string `json:"baseExpression"`
}

// ModelStateTableEvaluation describes a read-only evaluated parameter cell.
// Failure is empty for a healthy value. Example:
// `wire.ModelStateTableEvaluation{ParameterID: 1, Row: row, Value: "10 mm"}`.
type ModelStateTableEvaluation struct {
	ParameterID uint64                `json:"parameterId"`
	Row         ModelStateTableRowRef `json:"row"`
	Value       string                `json:"value"`
	Failure     string                `json:"failure"`
}

// ModelStateTableView is the sampled complete table. Columns specifies visible
// grid order; property, suppression and parameter collections also retain hidden
// values. ActiveMemberID is read-only. Evaluations are read-only results and are
// never accepted as authored values. Example: `view, err := client.ModelStates().GetTable()`.
type ModelStateTableView struct {
	Generation          uint64                           `json:"generation"`
	ActiveMemberID      uint64                           `json:"activeMemberId"`
	Rows                []ModelStateTableRow             `json:"rows"`
	Columns             []ModelStateTableColumnRef       `json:"columns"`
	Properties          []ModelStateTableProperty        `json:"properties"`
	Suppressions        []ModelStateTableSuppression     `json:"suppressions"`
	Parameters          []ModelStateTableParameter       `json:"parameters"`
	AvailableProperties []ModelStateTablePropertyChoice  `json:"availableProperties"`
	AvailableFeatures   []ModelStateTableFeatureChoice   `json:"availableFeatures"`
	AvailableParameters []ModelStateTableParameterChoice `json:"availableParameters"`
	Evaluations         []ModelStateTableEvaluation      `json:"evaluations"`
}

// SetModelStateTableArgs is a complete staged table applied as one undoable
// transaction. ExpectedGeneration must equal the sampled generation. Every
// collection is required, including explicit empty collections. Visible column
// references are unique and point to typed records; retained hidden records are
// not in Columns. Each record has exactly one cell per row in row order. The
// client must echo unchanged bases. A stale generation, bad reference, invalid
// value or failed validation rejects the entire edit. The active member is
// retained unless its row is deleted, in which case Primary becomes active.
// Example: `wire.SetModelStateTableArgs{ExpectedGeneration: generation, Rows: rows}`.
type SetModelStateTableArgs struct {
	ExpectedGeneration uint64                       `json:"expectedGeneration"`
	Rows               []ModelStateTableRow         `json:"rows"`
	Columns            []ModelStateTableColumnRef   `json:"columns"`
	Properties         []ModelStateTableProperty    `json:"properties"`
	Suppressions       []ModelStateTableSuppression `json:"suppressions"`
	Parameters         []ModelStateTableParameter   `json:"parameters"`
}

// ModelStateTableCreatedRow maps a client draft key to its reserved member ID.
// Example: `wire.ModelStateTableCreatedRow{DraftKey: "new-1", MemberID: 7}`.
type ModelStateTableCreatedRow struct {
	DraftKey string `json:"draftKey"`
	MemberID uint64 `json:"memberId"`
}

// ModelStateTableResult contains the adopted table and created-row identity map.
// Example: `result, err := client.ModelStates().SetTable(args)`.
type ModelStateTableResult struct {
	Table       ModelStateTableView         `json:"table"`
	CreatedRows []ModelStateTableCreatedRow `json:"createdRows"`
}

// ModelStateTableDiagnostic identifies a rejected table cell or a table-wide error.
// A nil Row and zero ParameterID identify an error that is not specific to a cell.
// Example: `wire.ModelStateTableDiagnostic{Row: &row, ParameterID: 3, Message: "unknown name"}`.
type ModelStateTableDiagnostic struct {
	Row         *ModelStateTableRowRef `json:"row"`
	ParameterID uint64                 `json:"parameterId"`
	Expression  string                 `json:"expression"`
	Message     string                 `json:"message"`
}

// ModelStateTableVerification reports candidate validation without publishing it.
// Example: `wire.ModelStateTableVerification{Valid: false, Diagnostics: errors}`.
type ModelStateTableVerification struct {
	Valid       bool                        `json:"valid"`
	Diagnostics []ModelStateTableDiagnostic `json:"diagnostics"`
}
