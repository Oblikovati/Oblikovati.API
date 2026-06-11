// SPDX-License-Identifier: Apache-2.0

// Package wire is the JSON method contract: the canonical method-name constants and
// the request/response DTOs that travel across the host↔add-in boundary (today the
// in-process C ABI of ADR-0016; tomorrow gRPC or a socket — the DTOs are transport
// agnostic). These shapes ARE the public automation surface; the GPL host
// (/source/addin/router) marshals model state into them, and the typed client
// ([oblikovati.org/api/client]) marshals add-in calls out of them.
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
	// Flavored document subtypes (M05-F15, Oblikovati#665).
	MethodDocumentsRegisterSubType = "documents.registerSubType"
	MethodDocumentsListSubTypes    = "documents.listSubTypes"

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

	// Custom line types loaded from industry-standard .lin definition files
	// (issue Oblikovati#161).
	MethodSketchGetCustomLineType = "sketch.getCustomLineType"
	MethodSketchSetCustomLineType = "sketch.setCustomLineType"

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

	// Feature lifecycle after placement (issue Oblikovati#140): features are
	// addressed by the stable id reported in [FeatureInfo] (model.tree).
	MethodFeaturesGet           = "features.get"
	MethodFeaturesEdit          = "features.edit"
	MethodFeaturesDelete        = "features.delete"
	MethodFeaturesRename        = "features.rename"
	MethodFeaturesSetSuppressed = "features.setSuppressed"
	MethodFeaturesReorder       = "features.reorder"

	MethodWorkPlanesList     = "workPlanes.list"
	MethodWorkPlanesCreate   = "workPlanes.create"
	MethodWorkPlanesRedefine = "workPlanes.redefine"

	MethodWorkPointsCreate = "workPoints.create"

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

	// Add-in registry & automation (M05-F01: #245, #251, #252).
	MethodAddInsList            = "addins.list"
	MethodAddInsGet             = "addins.get"
	MethodAddInsActivate        = "addins.activate"
	MethodAddInsDeactivate      = "addins.deactivate"
	MethodAddInsSetLoadBehavior = "addins.setLoadBehavior"
	MethodAddInsCallAutomation  = "addins.callAutomation"

	// External client applications driving the session (M05-F01, #245).
	MethodClientAppsRegister   = "clientApps.register"
	MethodClientAppsUnregister = "clientApps.unregister"
	MethodClientAppsList       = "clientApps.list"

	// Add-in browser panes (M05-F03: #247, #256).
	MethodBrowserSetPane    = "browser.setPane"
	MethodBrowserDeletePane = "browser.deletePane"
	MethodBrowserListPanes  = "browser.listPanes"

	// Add-in dockable windows (M05-F03, #247).
	MethodDockableWindowsSet        = "dockableWindows.set"
	MethodDockableWindowsSetVisible = "dockableWindows.setVisible"
	MethodDockableWindowsDelete     = "dockableWindows.delete"
	MethodDockableWindowsList       = "dockableWindows.list"

	// UI environments (M05-F03, #247; add-in environments: Oblikovati#667).
	MethodUIListEnvironments = "ui.listEnvironments"

	// Application options (M05-F11, #618).
	MethodOptionsListGroups = "options.listGroups"
	MethodOptionsGetGroup   = "options.getGroup"
	MethodOptionsSetGroup   = "options.setGroup"

	// Status, progress & user messaging (M05-F09, #616).
	MethodStatusSetText      = "status.setText"
	MethodStatusGetText      = "status.getText"
	MethodProgressBegin      = "progress.begin"
	MethodProgressUpdate     = "progress.update"
	MethodProgressEnd        = "progress.end"
	MethodBalloonTipRegister = "balloonTip.register"
	MethodBalloonTipShow     = "balloonTip.show"
	MethodPromptsShow        = "prompts.show"
	MethodErrorsAddMessage   = "errors.addMessage"
	MethodErrorsBeginSection = "errors.beginSection"
	MethodErrorsEndSection   = "errors.endSection"
	MethodErrorsList         = "errors.list"
	MethodErrorsClear        = "errors.clear"
	MethodErrorsShow         = "errors.show"

	// In-canvas mini-toolbars (M05-F07, #614).
	MethodMiniToolbarSet    = "miniToolbar.set"
	MethodMiniToolbarUpdate = "miniToolbar.update"
	MethodMiniToolbarRemove = "miniToolbar.remove"
	MethodMiniToolbarList   = "miniToolbar.list"

	// Host-provided modal dialogs (M05-F08, #615).
	MethodDialogsShowFileDialog = "dialogs.showFileDialog"
	MethodDialogsShowWebDialog  = "dialogs.showWebDialog"
	MethodDialogsCloseWebDialog = "dialogs.closeWebDialog"
	MethodDialogsListWebViews   = "dialogs.listWebViews"

	// Document windows: frames & tabs (M05-F10, #617).
	MethodWindowsListFrames  = "windows.listFrames"
	MethodWindowsListTabs    = "windows.listTabs"
	MethodWindowsActivateTab = "windows.activateTab"
	MethodWindowsCloseTab    = "windows.closeTab"

	// Help routing & language info (M05-F14, #621).
	MethodHelpRegisterContext = "help.registerContext"
	MethodHelpDisplay         = "help.display"
	MethodHelpPath            = "help.path"
	MethodLanguageInfo        = "language.info"

	// Interactive gizmos: the triad and manipulator handles (M05-F13, #620).
	MethodTriadShow          = "triad.show"
	MethodTriadUpdate        = "triad.update"
	MethodTriadHide          = "triad.hide"
	MethodTriadGet           = "triad.get"
	MethodManipulatorsSet    = "manipulators.set"
	MethodManipulatorsRemove = "manipulators.remove"

	// UI shell: search, marking menus, context menus, object visibility (M05-F12, #619).
	MethodUISearch              = "ui.search"
	MethodUIGetMarkingMenu      = "ui.getMarkingMenu"
	MethodUISetMarkingMenu      = "ui.setMarkingMenu"
	MethodUISetContextMenu      = "ui.setContextMenu"
	MethodUIGetObjectVisibility = "ui.getObjectVisibility"
	MethodUISetObjectVisibility = "ui.setObjectVisibility"
	// Add-in UI environments (M05-F16, Oblikovati#667).
	MethodUIRegisterEnvironment = "ui.registerEnvironment"
	MethodUIActivateEnvironment = "ui.activateEnvironment"
)

// Push-event type tags. These name host→add-in events delivered to the add-in's Notify
// entry point (ADR-0016), not callable request/response methods — there is no client
// method for them; an add-in matches on the event's "type" field. See the DTOs in this
// package (e.g. [EditCommittedEvent]).
const (
	EventEditCommitted = "edit.committed"
	// EventBrowserNode notifies an add-in of interaction with one of its browser
	// pane nodes (see [BrowserNodeEvent], M05-F03 #256).
	EventBrowserNode = "browser.node"
	// EventDockableWindowChanged notifies an add-in its dockable window was shown
	// or hidden (see [DockableWindowChangedEvent], M05-F03 #247).
	EventDockableWindowChanged = "dockableWindow.changed"
	// EventProgressCancelled notifies the owner its progress bar was cancelled
	// (see [ProgressCancelledEvent], M05-F09 #616).
	EventProgressCancelled = "progress.cancelled"
	// EventBalloonTipClicked notifies the owner its balloon was clicked
	// (see [BalloonTipClickedEvent], M05-F09 #616).
	EventBalloonTipClicked = "balloonTip.clicked"
	// EventPromptAnswered delivers a pending prompt's answer
	// (see [PromptAnsweredEvent], M05-F09 #616).
	EventPromptAnswered = "prompt.answered"
	// EventMiniToolbarChanged streams a mini-toolbar control edit
	// (see [MiniToolbarChangedEvent], M05-F07 #614).
	EventMiniToolbarChanged = "miniToolbar.changed"
	// EventMiniToolbarCommitted delivers a mini-toolbar's OK/Apply/Cancel
	// (see [MiniToolbarCommittedEvent], M05-F07 #614).
	EventMiniToolbarCommitted = "miniToolbar.committed"
	// EventFileDialogChosen delivers a file dialog's chosen paths
	// (see [FileDialogChosenEvent], M05-F08 #615).
	EventFileDialogChosen = "dialog.fileChosen"
	// EventWebDialogChanged notifies a web view was shown or closed
	// (see [WebDialogChangedEvent], M05-F08 #615).
	EventWebDialogChanged = "webDialog.changed"
	// EventCommandStarted reports a command beginning (M05-F12 #619).
	EventCommandStarted = "command.started"
	// EventSelectionChanged reports the selection set changing (M05-F12 #619).
	EventSelectionChanged = "selection.changed"
	// EventEnvironmentChanged reports the UI environment switching (M05-F12 #619).
	EventEnvironmentChanged = "ui.environmentChanged"
	// EventTriadDrag streams a triad gesture (see [TriadDragEvent], M05-F13 #620).
	EventTriadDrag = "triad.drag"
	// EventTriadSegment reports the hovered triad segment changing (M05-F13 #620).
	EventTriadSegment = "triad.segment"
	// EventManipulatorDrag streams a manipulator-handle gesture (M05-F13 #620).
	EventManipulatorDrag = "manipulator.drag"
	// EventClientOperation tells a subtype's owner its flavored document needs
	// servicing (see [ClientOperationEvent], M05-F15 Oblikovati#665).
	EventClientOperation = "client.operation"
)

// OKResult is the trivial success payload for mutating methods with no return value.
type OKResult struct {
	OK bool `json:"ok"`
}
