// SPDX-License-Identifier: Apache-2.0

// Package wire is the JSON method contract: the canonical method-name constants and
// the request/response DTOs that travel across the host↔add-in boundary (today the
// in-process C ABI of ADR-0016; tomorrow gRPC or a socket — the DTOs are transport
// agnostic). These shapes ARE the public automation surface; the GPL host
// (/source/addin/router) marshals model state into them, and the typed client
// ([oblikovati/api/client]) marshals add-in calls out of them.
//
// Every type here is plain data with stable JSON tags — no behavior, no dependency
// on the implementation. Field renames are breaking changes to the contract.
package wire

// Method names. These string constants are the wire identity of each operation; the
// host router keys its dispatch table on them and the client sends them. Treat the
// string values as frozen (clients and saved automations depend on them).
const (
	MethodCommandsList     = "commands.list"
	MethodCommandsExecute  = "commands.execute"
	MethodCommandsCreate   = "commands.create"
	MethodCommandsSetState = "commands.setState"

	MethodRibbonList = "ribbon.list"

	MethodDocumentsList     = "documents.list"
	MethodDocumentsCreate   = "documents.create"
	MethodDocumentsActivate = "documents.activate"
	MethodDocumentsClose    = "documents.close"
	MethodDocumentsCloseAll = "documents.closeAll"
	MethodDocumentsImport   = "documents.import"
	MethodDocumentsExport   = "documents.export"

	MethodParametersList = "parameters.list"
	MethodParametersGet  = "parameters.get"
	MethodParametersAdd  = "parameters.add"
	MethodParametersSet  = "parameters.set"

	MethodModelTree          = "model.tree"
	MethodModelSelection     = "model.selection"
	MethodModelReferenceKeys = "model.referenceKeys"

	MethodSketchCreate    = "sketch.create"
	MethodSketchRectangle = "sketch.rectangle"

	MethodSketchList        = "sketch.list"
	MethodSketchGet         = "sketch.get"
	MethodSketchEdit        = "sketch.edit"
	MethodSketchExitEdit    = "sketch.exitEdit"
	MethodSketchSolve       = "sketch.solve"
	MethodSketchDelete      = "sketch.delete"
	MethodSketchEntities    = "sketch.entities"
	MethodSketchConstraints = "sketch.constraints"
	MethodSketchDimensions  = "sketch.dimensions"
	MethodSketchSetProperty = "sketch.setProperty"
	MethodSketchAddEntity   = "sketch.addEntity"

	MethodSketchAddConstraint    = "sketch.addConstraint"
	MethodSketchDeleteConstraint = "sketch.deleteConstraint"

	MethodSketchAddDimension   = "sketch.addDimension"
	MethodSketchDriveDimension = "sketch.driveDimension"

	MethodSketchConstraintStatus = "sketch.constraintStatus"
	MethodSketchProfiles         = "sketch.profiles"
	MethodSketchTransform        = "sketch.transform"
	MethodSketchAddPattern       = "sketch.addPattern"
	MethodSketchOffset           = "sketch.offset"
	MethodSketchAddImage         = "sketch.addImage"
	MethodSketchAddFillRegion    = "sketch.addFillRegion"
	MethodSketchAddText          = "sketch.addText"
	MethodSketchEditText         = "sketch.editText"
	MethodSketchGetText          = "sketch.getText"
	MethodSketchSetTextFont      = "sketch.setTextFont"
	MethodSketchAutoDimension    = "sketch.autoDimension"
	MethodSketchProject          = "sketch.project"

	MethodSketch3DCreate           = "sketch3d.create"
	MethodSketch3DList             = "sketch3d.list"
	MethodSketch3DGet              = "sketch3d.get"
	MethodSketch3DEdit             = "sketch3d.edit"
	MethodSketch3DExitEdit         = "sketch3d.exitEdit"
	MethodSketch3DSolve            = "sketch3d.solve"
	MethodSketch3DDelete           = "sketch3d.delete"
	MethodSketch3DSetProperty      = "sketch3d.setProperty"
	MethodSketch3DEntities         = "sketch3d.entities"
	MethodSketch3DConstraints      = "sketch3d.constraints"
	MethodSketch3DDimensions       = "sketch3d.dimensions"
	MethodSketch3DConstraintStatus = "sketch3d.constraintStatus"
	MethodSketch3DAddEntity        = "sketch3d.addEntity"
	MethodSketch3DAddConstraint    = "sketch3d.addConstraint"
	MethodSketch3DDeleteConstraint = "sketch3d.deleteConstraint"
	MethodSketch3DAddDimension     = "sketch3d.addDimension"
	MethodSketch3DDriveDimension   = "sketch3d.driveDimension"
	MethodSketch3DProfiles         = "sketch3d.profiles"
	MethodSketch3DPaths            = "sketch3d.paths"
	MethodSketch3DTransform        = "sketch3d.transform"
	MethodSketch3DInclude          = "sketch3d.include"
	MethodSketch3DIncludeSketch    = "sketch3d.includeSketch"
	MethodSketch3DAddSurfaceCurve  = "sketch3d.addSurfaceCurve"

	MethodFeaturesList = "features.list"
	MethodFeaturesAdd  = "features.add"

	MethodWorkPlanesList   = "workPlanes.list"
	MethodWorkPlanesCreate = "workPlanes.create"

	MethodThemeActive = "theme.active"
	MethodThemeList   = "theme.list"

	MethodFontsList = "fonts.list"

	MethodViewGetDisplayMode   = "view.getDisplayMode"
	MethodViewSetDisplayMode   = "view.setDisplayMode"
	MethodViewListDisplayModes = "view.listDisplayModes"

	MethodViewGetShadows = "view.getShadows"
	MethodViewSetShadows = "view.setShadows"

	MethodViewGetCamera = "view.getCamera"
	MethodViewSetCamera = "view.setCamera"

	MethodViewportCapture        = "viewport.capture"
	MethodViewportSetNormalDebug = "viewport.setNormalDebug"
	MethodViewportSetMeshColors  = "viewport.setMeshColors"

	MethodViewsList      = "views.list"
	MethodViewsAdd       = "views.add"
	MethodViewsActivate  = "views.activate"
	MethodViewsClose     = "views.close"
	MethodViewsRename    = "views.rename"
	MethodViewsGetLayout = "views.getLayout"
	MethodViewsSetLayout = "views.setLayout"

	MethodLightingGetStyle   = "lighting.getStyle"
	MethodLightingSetStyle   = "lighting.setStyle"
	MethodLightingListStyles = "lighting.listStyles"
	MethodLightingListLights = "lighting.listLights"
	MethodLightingAddLight   = "lighting.addLight"
	MethodLightingSetLight   = "lighting.setLight"

	MethodEnvironmentGet         = "environment.get"
	MethodEnvironmentSet         = "environment.set"
	MethodEnvironmentListPresets = "environment.listPresets"
	MethodEnvironmentLoadImage   = "environment.loadImage"

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

	MethodClientGraphicsSet        = "clientGraphics.set"
	MethodClientGraphicsList       = "clientGraphics.list"
	MethodClientGraphicsDelete     = "clientGraphics.delete"
	MethodClientGraphicsSetVisible = "clientGraphics.setVisible"

	MethodInteractionGraphicsUpdate = "interactionGraphics.update"
	MethodInteractionGraphicsClear  = "interactionGraphics.clear"

	MethodTransactionUndo  = "transaction.undo"
	MethodTransactionRedo  = "transaction.redo"
	MethodTransactionState = "transaction.state"
	MethodTransactionBegin = "transaction.begin"
	MethodTransactionEnd   = "transaction.end"

	MethodInteractionState     = "interaction.state"
	MethodInteractionSetNotice = "interaction.setNotice"

	MethodScriptRun = "scripts.run"

	MethodLogsTail = "logs.tail"
)

// Push-event type tags. These name host→add-in events delivered to the add-in's Notify
// entry point (ADR-0016), not callable request/response methods — there is no client
// method for them; an add-in matches on the event's "type" field. See the DTOs in this
// package (e.g. [EditCommittedEvent]).
const (
	EventEditCommitted = "edit.committed"
)

// OKResult is the trivial success payload for mutating methods with no return value.
type OKResult struct {
	OK bool `json:"ok"`
}
