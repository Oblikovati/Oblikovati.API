// SPDX-License-Identifier: Apache-2.0

// Package wire is the JSON method contract: the canonical method-name constants and
// the request/response DTOs that travel across the host↔add-in boundary (today the
// in-process C ABI of ADR-0016; tomorrow gRPC or a socket — the DTOs are transport
// agnostic). These shapes ARE the public automation surface; the GPL host
// (/source/addin/router) marshals model state into them, and the typed client
// ([github.com/Oblikovati/api/client]) marshals add-in calls out of them.
//
// Every type here is plain data with stable JSON tags — no behavior, no dependency
// on the implementation. Field renames are breaking changes to the contract.
package wire

// Method names. These string constants are the wire identity of each operation; the
// host router keys its dispatch table on them and the client sends them. Treat the
// string values as frozen (clients and saved automations depend on them).
const (
	MethodCommandsList    = "commands.list"
	MethodCommandsExecute = "commands.execute"
	MethodCommandsCreate  = "commands.create"

	MethodDocumentsList     = "documents.list"
	MethodDocumentsCreate   = "documents.create"
	MethodDocumentsActivate = "documents.activate"

	MethodParametersList = "parameters.list"
	MethodParametersGet  = "parameters.get"
	MethodParametersAdd  = "parameters.add"
	MethodParametersSet  = "parameters.set"

	MethodModelTree      = "model.tree"
	MethodModelSelection = "model.selection"

	MethodSketchCreate    = "sketch.create"
	MethodSketchRectangle = "sketch.rectangle"

	MethodFeaturesList = "features.list"
	MethodFeaturesAdd  = "features.add"

	MethodWorkPlanesList   = "workPlanes.list"
	MethodWorkPlanesCreate = "workPlanes.create"

	MethodThemeActive = "theme.active"
	MethodThemeList   = "theme.list"

	MethodAppearancesList   = "appearances.list"
	MethodAppearancesGet    = "appearances.get"
	MethodAppearancesCreate = "appearances.create"
	MethodAppearancesUpdate = "appearances.update"

	MethodMaterialsList   = "materials.list"
	MethodMaterialsGet    = "materials.get"
	MethodMaterialsCreate = "materials.create"
	MethodMaterialsUpdate = "materials.update"

	MethodModelAssignMaterial     = "model.assignMaterial"
	MethodModelAssignAppearance   = "model.assignAppearance"
	MethodModelPhysicalProperties = "model.physicalProperties"
)

// OKResult is the trivial success payload for mutating methods with no return value.
type OKResult struct {
	OK bool `json:"ok"`
}
