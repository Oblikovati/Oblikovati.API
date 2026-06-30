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

	MethodCommandLineSubmit = "commandLine.submit"

	MethodRibbonList = "ribbon.list"

	MethodDocumentsList           = "documents.list"
	MethodDocumentsUpdate         = "documents.update"
	MethodDocumentsRebuild        = "documents.rebuild"
	MethodDocumentsRequiresUpdate = "documents.requiresUpdate"
	MethodDocumentsCreate         = "documents.create"
	MethodDocumentsActivate       = "documents.activate"
	MethodDocumentsClose          = "documents.close"
	MethodDocumentsCloseAll       = "documents.closeAll"
	MethodDocumentsImport         = "documents.import"
	MethodDocumentsExport         = "documents.export"
	// Flavored document subtypes (M05-F15, Oblikovati#665).
	MethodDocumentsRegisterSubType = "documents.registerSubType"
	MethodDocumentsListSubTypes    = "documents.listSubTypes"

	// Document iProperties — metadata sets read/written by add-ins, BOMs, title blocks (#156).
	MethodDocumentsListProperties = "documents.listProperties"
	MethodDocumentsGetProperty    = "documents.getProperty"
	MethodDocumentsSetProperty    = "documents.setProperty"

	// Per-document settings (#147) — the Document Settings dialog's persisted defaults. Starts with
	// the Sketch tab (the constraint-inference preferences the sketch tools read).
	MethodDocumentGetSketchSettings = "document.getSketchSettings"
	MethodDocumentSetSketchSettings = "document.setSketchSettings"

	// Part end-of-part rollback marker (#141) — inspect / move how far down the feature program the
	// active part evaluates (authoring features mid-history, design inspection).
	MethodDocumentGetEndOfPart = "document.getEndOfPart"
	MethodDocumentSetEndOfPart = "document.setEndOfPart"

	// Add-in attribute sets (#155) — named, typed values an add-in attaches to a document and
	// that persist with it: the sanctioned mechanism for add-ins to store their own data and tag
	// the model. set/get/list address a named attribute in a named set on a document; listSets
	// enumerates the sets; delete removes an attribute (or a whole set); find locates the open
	// documents carrying a given set/attribute.
	MethodAttributesSet      = "attributes.set"
	MethodAttributesGet      = "attributes.get"
	MethodAttributesList     = "attributes.list"
	MethodAttributesListSets = "attributes.listSets"
	MethodAttributesDelete   = "attributes.delete"
	MethodAttributesFind     = "attributes.find"

	// Document open/save lifecycle (#138) and the save policy layer around it:
	// SaveCopyAs and batch save (M03-F09, #610).
	MethodDocumentsOpen       = "documents.open"
	MethodDocumentsSave       = "documents.save"
	MethodDocumentsSaveAs     = "documents.saveAs"
	MethodDocumentsSaveCopyAs = "documents.saveCopyAs"
	MethodDocumentsBatchSave  = "documents.batchSave"

	// The file as an object distinct from its documents, with file-level
	// reference descriptors and repair (M03-F07, #608).
	MethodFilesGet                    = "files.get"
	MethodFilesListReferences         = "files.listReferences"
	MethodFilesReplaceReference       = "files.replaceReference"
	MethodDocumentsListFileReferences = "documents.listFileReferences"

	// Linked/embedded external-file attachments on a document (M03-F08, #609).
	MethodDocumentsListAttachments  = "documents.listAttachments"
	MethodDocumentsAddAttachment    = "documents.addAttachment"
	MethodDocumentsRemoveAttachment = "documents.removeAttachment"

	// The add-in data registry on documents (M03-F10, #611).
	MethodDocumentsListInterests  = "documents.listInterests"
	MethodDocumentsAddInterest    = "documents.addInterest"
	MethodDocumentsRemoveInterest = "documents.removeInterest"
	MethodDocumentsHasInterest    = "documents.hasInterest"

	// Document units of measure + unit/expression service (Oblikovati#146).
	MethodDocumentsGetUnits = "documents.getUnits"
	MethodDocumentsSetUnits = "documents.setUnits"

	MethodUnitsConvert                        = "units.convert"
	MethodUnitsGetStringFromValue             = "units.getStringFromValue"
	MethodUnitsGetPreciseStringFromValue      = "units.getPreciseStringFromValue"
	MethodUnitsGetValueFromExpression         = "units.getValueFromExpression"
	MethodUnitsGetDatabaseUnitsFromExpression = "units.getDatabaseUnitsFromExpression"
	MethodUnitsIsExpressionValid              = "units.isExpressionValid"
	MethodUnitsCompatibleUnits                = "units.compatibleUnits"
	MethodUnitsGetTypeFromString              = "units.getTypeFromString"
	MethodUnitsGetStringFromType              = "units.getStringFromType"
	MethodUnitsGetLocaleCorrectedExpression   = "units.getLocaleCorrectedExpression"
	MethodUnitsGetDrivingParameters           = "units.getDrivingParameters"

	MethodParametersList = "parameters.list"
	MethodParametersGet  = "parameters.get"
	MethodParametersAdd  = "parameters.add"
	MethodParametersSet  = "parameters.set"
	// Member-level parameter surface (M02-F08, Oblikovati#607).
	MethodParametersGetDetail         = "parameters.getDetail"
	MethodParametersUpdate            = "parameters.update"
	MethodParametersSetTolerance      = "parameters.setTolerance"
	MethodParametersSetExpressionList = "parameters.setExpressionList"
	MethodParametersDelete            = "parameters.delete"
	MethodParametersDrivenBy          = "parameters.drivenBy"
	MethodParametersDependents        = "parameters.dependents"
	// Parameter settings, tolerance sweeps & exchange (M02-F07, Oblikovati#606).
	MethodParametersGetSettings          = "parameters.getSettings"
	MethodParametersSetSettings          = "parameters.setSettings"
	MethodParametersSetAllModelValueType = "parameters.setAllModelValueType"
	MethodParametersExport               = "parameters.export"
	MethodParametersImport               = "parameters.import"
	// Derived parameter tables (M02-F06, Oblikovati#605).
	MethodParametersDerivedTablesList      = "parameters.derivedTables.list"
	MethodParametersDerivedTablesAdd       = "parameters.derivedTables.add"
	MethodParametersDerivedTablesSetLinked = "parameters.derivedTables.setLinked"
	MethodParametersDerivedTablesDelete    = "parameters.derivedTables.delete"
	// Custom parameter groups (M02-F05, Oblikovati#604).
	MethodParametersGroupsList           = "parameters.groups.list"
	MethodParametersGroupsAdd            = "parameters.groups.add"
	MethodParametersGroupsDelete         = "parameters.groups.delete"
	MethodParametersGroupsSetDisplayName = "parameters.groups.setDisplayName"
	MethodParametersGroupsAddMember      = "parameters.groups.addMember"
	MethodParametersGroupsRemoveMember   = "parameters.groups.removeMember"

	MethodModelTree          = "model.tree"
	MethodModelSelection     = "model.selection"
	MethodModelReferenceKeys = "model.referenceKeys"

	// Selection mutation (#157) — make the previously read-only selection writable: select the
	// entities named by their reference strings (from a SelectionResult), remove them, or clear
	// the whole set. The reply is the new SelectionResult.
	MethodModelSelect         = "model.select"
	MethodModelDeselect       = "model.deselect"
	MethodModelClearSelection = "model.clearSelection"

	// Highlight sets (#157): named, colored emphasis groups the viewport outlines without
	// selecting — an add-in guides the user. create/delete/addItems/setColor/list.
	MethodModelHighlightSetCreate   = "model.highlightSets.create"
	MethodModelHighlightSetDelete   = "model.highlightSets.delete"
	MethodModelHighlightSetAddItems = "model.highlightSets.addItems"
	MethodModelHighlightSetSetColor = "model.highlightSets.setColor"
	MethodModelHighlightSetList     = "model.highlightSets.list"

	MethodImportDWG = "import.dwg"
	MethodImportDXF = "import.dxf"
	MethodImportPDF = "import.pdf"
	MethodExportDXF = "export.dxf"

	MethodSketchCreate    = "sketch.create"
	MethodSketchRectangle = "sketch.rectangle"

	MethodSketchList        = "sketch.list"
	MethodSketchGet         = "sketch.get"
	MethodSketchDependents  = "sketch.dependents"
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
	MethodSketchCopyTo           = "sketch.copyTo"
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

	// Sketch blocks: reusable entity groups instanced with a placement
	// transform (M06-F07, Oblikovati/Oblikovati#622).
	MethodSketchBlockDefinitionCreate = "sketch.blockDefinitions.create"
	MethodSketchBlockDefinitionList   = "sketch.blockDefinitions.list"
	MethodSketchBlockDefinitionDelete = "sketch.blockDefinitions.delete"
	MethodSketchAddBlockInstance      = "sketch.addBlockInstance"
	MethodSketchListBlockInstances    = "sketch.blockInstances"

	// Region properties of a closed profile (M06-F08, #623).
	MethodSketchRegionProperties = "sketch.regionProperties"

	// Sketch inference options: enable/disable + constraint-family priority
	// (M06-F10, #625).
	MethodSketchSetInferenceOptions = "sketch.setInferenceOptions"
	MethodSketchGetInferenceOptions = "sketch.getInferenceOptions"

	// Spline tangency handles (M06-F11, #626).
	MethodSketchSetSplineHandle = "sketch.setSplineHandle"

	// Persistent reference keys for a sketch and its entities (#153):
	// referenceKey returns a sketch's own durable key; resolveReference rebinds a
	// stored sketch/entity key to its current location.
	MethodSketchReferenceKey     = "sketch.referenceKey"
	MethodSketchResolveReference = "sketch.resolveReference"

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

	// Region properties of a planar closed 3D profile (M06-F08, #623).
	MethodSketch3DRegionProperties = "sketch3d.regionProperties"
	// Helical curve definition edit — constant or variable shape rows plus
	// end conditions (M06-F09, #624).
	MethodSketch3DEditHelix = "sketch3d.editHelix"
	// 3D spline tangency handles (M06-F11, #626).
	MethodSketch3DSetSplineHandle = "sketch3d.setSplineHandle"

	// A 3D sketch's own persistent reference key (#153). 3D entity keys are reported by
	// MethodSketch3DEntities, and any key (2D or 3D) rebinds through
	// MethodSketchResolveReference.
	MethodSketch3DReferenceKey = "sketch3d.referenceKey"

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

	// Sheet-metal rule/style surface (M13-F01, Oblikovati#373/#369): the active part's
	// sheet-metal environment — its rule (thickness/bend-radius/relief/gap) and unfold
	// method. getStyle reports the active rule; setStyle edits it and recomputes (the rule
	// is parameter-backed, so a thickness/K-factor change repropagates to every wall);
	// bendAllowance previews the developed flat length of one bend under the active method.
	MethodSheetMetalGetStyle      = "sheetMetal.getStyle"
	MethodSheetMetalSetStyle      = "sheetMetal.setStyle"
	MethodSheetMetalBendAllowance = "sheetMetal.bendAllowance"
	// bends reports the part's bend lineage (M13-F04, Oblikovati#377): every bend the
	// wall/bend features introduced, with the unfold values (allowance/deduction) the flat
	// pattern develops it by. This is the flat pattern's prerequisite — the architecture
	// requires every bend to record its unfold parameters.
	MethodSheetMetalBends = "sheetMetal.bends"
	// unfold develops the folded part into its flat pattern (M13-F04, Oblikovati#377) and
	// reports the flat: its 2D extents, gauge, developed area and fold lines. The flat is a
	// derived body, so a thickness/K-factor edit changes its extents through the bend
	// allowance with no model recompute.
	MethodSheetMetalUnfold = "sheetMetal.unfold"

	// Flat-pattern orientations (M13-F05, Oblikovati#635): named alignment states that frame
	// the developed flat for drawing views and export. The active orientation drives the flat's
	// reported length/width; orientations persist in the document.
	MethodFlatPatternListOrientations    = "flatPattern.listOrientations"
	MethodFlatPatternAddOrientation      = "flatPattern.addOrientation"
	MethodFlatPatternActivateOrientation = "flatPattern.activateOrientation"
	MethodFlatPatternDeleteOrientation   = "flatPattern.deleteOrientation"
	// Flat-pattern edge/face classification (M13-F05, Oblikovati#635): the bend-up/bend-down
	// fold lines and the front/back faces of the developed flat — the manufacturing/drawing
	// classification of the flat's topology (the DXF layer-mapping input).
	MethodFlatPatternEdgesOfType = "flatPattern.edgesOfType"
	MethodFlatPatternFaces       = "flatPattern.faces"
	// mapEntity maps a folded-model topology entity to its developed-flat counterpart (or back)
	// by reference key (M13-F05, Oblikovati#635), so a drawing dimension or selection on the flat
	// survives recompute.
	MethodFlatPatternMapEntity = "flatPattern.mapEntity"
	// Flat-pattern plates + settings (M13-F05, Oblikovati#635): the disjoint developed regions
	// (one plate per connected flat region) and the per-document settings (deferred flat-pattern
	// update so a heavy flat only recomputes on demand).
	MethodFlatPatternListPlates  = "flatPattern.listPlates"
	MethodFlatPatternGetSettings = "flatPattern.getSettings"
	MethodFlatPatternSetSettings = "flatPattern.setSettings"
	// Bend-order annotation (M13-F06, Oblikovati#809): number/sequence the part's bends for
	// press-brake sequencing; the order is editable and persists, and is shown on the flat
	// pattern and in drawing views.
	MethodFlatPatternListBendOrder = "flatPattern.listBendOrder"
	MethodFlatPatternSetBendOrder  = "flatPattern.setBendOrder"
	// Cosmetic centerlines (M13-F06, Oblikovati#809): annotation lines drawn on the flat
	// pattern (e.g. a hole-pattern centerline) for manufacturing reference. They persist and
	// appear on the flat and in drawing views.
	MethodFlatPatternAddCenterline    = "flatPattern.addCenterline"
	MethodFlatPatternListCenterlines  = "flatPattern.listCenterlines"
	MethodFlatPatternDeleteCenterline = "flatPattern.deleteCenterline"

	// Drawing document sheets (M14-F01, Oblikovati#384): the active drawing's sheets
	// (sizes/orientation, borders, title blocks), the active-sheet selection, the
	// primary referenced model (whose iProperties feed title-block fields), and a title
	// block's resolved field values.
	MethodDrawingListSheets        = "drawing.listSheets"
	MethodDrawingAddSheet          = "drawing.addSheet"
	MethodDrawingRemoveSheet       = "drawing.removeSheet"
	MethodDrawingSetActiveSheet    = "drawing.setActiveSheet"
	MethodDrawingSetModelReference = "drawing.setModelReference"
	MethodDrawingTitleBlockFields  = "drawing.titleBlockFields"
	// Drawing sheet DXF export (M14-F05 PBI-145, Oblikovati#392): write the active sheet —
	// its views' visible/hidden edges, border and title block — to a DXF file, on named layers.
	MethodDrawingExportDXF = "drawing.exportDXF"

	// Drawing drafting standards + styles (M14-F01 PBI-138, Oblikovati#385): the active
	// drawing's drafting standard and its dimension/text/line style preset. Switching the
	// standard re-points the preset so every annotation re-renders to it.
	MethodDrawingStylesListStandards  = "drawingStyles.listStandards"
	MethodDrawingStylesGetActiveStyle = "drawingStyles.getActiveStyle"
	MethodDrawingStylesSetStandard    = "drawingStyles.setStandard"

	// Drawing views (M14-F02 PBI-139, Oblikovati#386): project the referenced model onto the
	// active sheet — a base view from a standard orientation and projected views off it — with
	// hidden-line removal producing visible/hidden drawing curves.
	MethodDrawingViewsList         = "drawingViews.list"
	MethodDrawingViewsAddBase      = "drawingViews.addBase"
	MethodDrawingViewsAddProjected = "drawingViews.addProjected"
	MethodDrawingViewsAddAuxiliary = "drawingViews.addAuxiliary"
	MethodDrawingViewsAddSection   = "drawingViews.addSection"
	MethodDrawingViewsAddDetail    = "drawingViews.addDetail"
	MethodDrawingViewsAddBreak     = "drawingViews.addBreak"
	MethodDrawingViewsAddSlice     = "drawingViews.addSlice"
	MethodDrawingViewsAddBreakout  = "drawingViews.addBreakout"
	MethodDrawingViewsAddDraft     = "drawingViews.addDraft"
	MethodDrawingViewsDelete       = "drawingViews.delete"
	MethodDrawingViewsCurves       = "drawingViews.curves"

	// Drawing annotations (M14-F02 #813): the centre-of-gravity marker (driven by the
	// referenced model's mass properties) and revision-cloud sheet markup.
	MethodDrawingAnnotationsList             = "drawingAnnotations.list"
	MethodDrawingAnnotationsAddCoG           = "drawingAnnotations.addCoG"
	MethodDrawingAnnotationsAddRevisionCloud = "drawingAnnotations.addRevisionCloud"
	MethodDrawingAnnotationsAddCenterMarks   = "drawingAnnotations.addCenterMarks"
	MethodDrawingAnnotationsAddCenterlines   = "drawingAnnotations.addCenterlines"
	MethodDrawingAnnotationsAddFCF           = "drawingAnnotations.addFeatureControlFrame"
	MethodDrawingAnnotationsAddDatum         = "drawingAnnotations.addDatumFeature"
	MethodDrawingAnnotationsAddSurfaceText   = "drawingAnnotations.addSurfaceTexture"
	MethodDrawingAnnotationsAddPartsList     = "drawingAnnotations.addPartsList"
	MethodDrawingAnnotationsAddBalloon       = "drawingAnnotations.addBalloon"
	MethodDrawingAnnotationsAddHoleTable     = "drawingAnnotations.addHoleTable"
	MethodDrawingAnnotationsAddRevTable      = "drawingAnnotations.addRevisionTable"
	MethodDrawingAnnotationsAddRevTag        = "drawingAnnotations.addRevisionTag"
	MethodDrawingAnnotationsAddNote          = "drawingAnnotations.addNote"
	MethodDrawingAnnotationsAddCustomTable   = "drawingAnnotations.addCustomTable"
	MethodDrawingAnnotationsAddHoleNotes     = "drawingAnnotations.addHoleNotes"
	MethodDrawingAnnotationsDelete           = "drawingAnnotations.delete"

	// Drawing sketches (M14-F08 #638): 2D geometry drawn directly in sheet space (millimetres) on a
	// sheet — linework and boundaries that hatch regions can fill.
	MethodDrawingSketchesList      = "drawingSketches.list"
	MethodDrawingSketchesAdd       = "drawingSketches.add"
	MethodDrawingSketchesAddEntity = "drawingSketches.addEntity"
	MethodDrawingSketchesAddHatch  = "drawingSketches.addHatchRegion"

	// Drawing dimensions (M14-F03 PBI-141 #388): associative linear dimensions on a view,
	// snapped to projected model vertices so the measured value tracks the model.
	MethodDrawingDimensionsList         = "drawingDimensions.list"
	MethodDrawingDimensionsAddLinear    = "drawingDimensions.addLinear"
	MethodDrawingDimensionsAddRadial    = "drawingDimensions.addRadial"
	MethodDrawingDimensionsAddAngular   = "drawingDimensions.addAngular"
	MethodDrawingDimensionsAddBaseline  = "drawingDimensions.addBaseline"
	MethodDrawingDimensionsAddChain     = "drawingDimensions.addChain"
	MethodDrawingDimensionsAddOrdinate  = "drawingDimensions.addOrdinate"
	MethodDrawingDimensionsAddArcLength = "drawingDimensions.addArcLength"
	MethodDrawingDimensionsDelete       = "drawingDimensions.delete"

	// Thread table query + designation resolution (M09-F01 PBI-101, #325):
	// one source of truth for thread data across tapping and drawings.
	MethodThreadsTableQuery = "threads.tableQuery"
	MethodThreadsResolve    = "threads.resolve"

	// Freeform cage editing after placement (M10-F03 PBI-114,
	// Oblikovati#699): the feature is addressed by its stable id, like the
	// features.* lifecycle methods.
	MethodFreeformSetLevel     = "freeform.setLevel"
	MethodFreeformMoveVertices = "freeform.moveVertices"
	MethodFreeformCreaseEdges  = "freeform.creaseEdges"

	// Assembly derive/shrinkwrap (M11-F06, Oblikovati#631/#716): derive a source
	// assembly into the active part as a base body, or simplify it into a lightweight
	// shrinkwrap body; break the link to freeze the current result. The derived feature
	// is addressed by its stable id, like the features.* lifecycle methods.
	MethodAssemblyDeriveCreate     = "assembly.deriveCreate"
	MethodAssemblyDeriveBreakLink  = "assembly.deriveBreakLink"
	MethodAssemblyDeriveStatus     = "assembly.deriveStatus"
	MethodAssemblyDeriveUpdate     = "assembly.deriveUpdate"
	MethodAssemblyShrinkwrapCreate = "assembly.shrinkwrapCreate"

	// Assembly occurrences (M11-F01/F02, Oblikovati#728): read the active assembly's
	// occurrence tree and place/transform/ground/suppress/replace/remove components.
	// Occurrences are addressed by session id (the ids the occurrence push events carry).
	MethodAssemblyOccurrences            = "assembly.occurrences"
	MethodAssemblyPlace                  = "assembly.place"
	MethodAssemblyPlaceByDefinition      = "assembly.placeByDefinition"
	MethodAssemblyPlaceByDefinitionBatch = "assembly.placeByDefinitionBatch"
	MethodAssemblyTransform              = "assembly.transform"
	MethodAssemblyGround                 = "assembly.ground"
	MethodAssemblySuppress               = "assembly.suppress"
	MethodAssemblySetFlexible            = "assembly.setFlexible"      // M12-F06
	MethodAssemblySetFlexibleChild       = "assembly.setFlexibleChild" // M12-F06 independent solve
	MethodAssemblyReplace                = "assembly.replace"
	MethodAssemblyRemove                 = "assembly.remove"

	// Assembly replication (M11-F04, Oblikovati#729): replicate placed components —
	// pattern (circular/rectangular), mirror across a plane, independent copy, and
	// substitute a set of components with one simplified representation.
	MethodAssemblyPatternCreate  = "assembly.patternCreate"
	MethodAssemblyMirror         = "assembly.mirror"
	MethodAssemblyMirrorIntoPart = "assembly.mirrorIntoPart"
	MethodAssemblyCopy           = "assembly.copy"
	MethodAssemblySubstitute     = "assembly.substitute"

	// Assembly bill of materials (M11-F05, Oblikovati#730): read a structured or
	// parts-only BOM view of the active assembly, and export a view to CSV with optional
	// custom property columns.
	MethodAssemblyBOMView   = "assembly.bomView"
	MethodAssemblyBOMExport = "assembly.bomExport"

	// Assembly feature program (M11-F08, Oblikovati#633/#725): the machining features
	// authored in the assembly, their per-occurrence participation and suppression, and
	// the end-of-features rollback marker.
	MethodAssemblyFeaturesList                = "assemblyFeatures.list"
	MethodAssemblyFeaturesAdd                 = "assemblyFeatures.add"
	MethodAssemblyFeaturesAddProxyCut         = "assemblyFeatures.addProxyCut"
	MethodAssemblyFeaturesAddHole             = "assemblyFeatures.addHole"
	MethodAssemblyFeaturesAddExtrude          = "assemblyFeatures.addExtrude"
	MethodAssemblyFeaturesAddRevolve          = "assemblyFeatures.addRevolve"
	MethodAssemblyFeaturesAddChamfer          = "assemblyFeatures.addChamfer"
	MethodAssemblyFeaturesAddFillet           = "assemblyFeatures.addFillet"
	MethodAssemblyFeaturesAddMoveFace         = "assemblyFeatures.addMoveFace"
	MethodAssemblyFeaturesAddSweep            = "assemblyFeatures.addSweep"
	MethodAssemblyFeaturesEdit                = "assemblyFeatures.edit"
	MethodAssemblyFeaturesSetParticipants     = "assemblyFeatures.setParticipants"
	MethodAssemblyFeaturesSetParticipantPaths = "assemblyFeatures.setParticipantPaths"
	MethodAssemblyFeaturesSetSuppressed       = "assemblyFeatures.setSuppressed"
	MethodAssemblyGetEndOfFeatures            = "assembly.getEndOfFeatures"
	MethodAssemblySetEndOfFeatures            = "assembly.setEndOfFeatures"

	// Assembly constraints (M12-F01, Oblikovati#358/#363): the relationships that
	// position one occurrence relative to another, the solve that applies them, and the
	// assembly's health / per-occurrence degrees-of-freedom report.
	MethodAssemblyConstraintsList                  = "assemblyConstraints.list"
	MethodAssemblyConstraintsAddMate               = "assemblyConstraints.addMate"
	MethodAssemblyConstraintsAddFlush              = "assemblyConstraints.addFlush"
	MethodAssemblyConstraintsAddAngle              = "assemblyConstraints.addAngle"
	MethodAssemblyConstraintsAddTangent            = "assemblyConstraints.addTangent"
	MethodAssemblyConstraintsAddInsert             = "assemblyConstraints.addInsert"
	MethodAssemblyConstraintsSnap                  = "assemblyConstraints.snap"
	MethodAssemblyConstraintsAddSymmetry           = "assemblyConstraints.addSymmetry"
	MethodAssemblyConstraintsAddRotateRotate       = "assemblyConstraints.addRotateRotate"
	MethodAssemblyConstraintsAddRotateTranslate    = "assemblyConstraints.addRotateTranslate"
	MethodAssemblyConstraintsAddTranslateTranslate = "assemblyConstraints.addTranslateTranslate"
	MethodAssemblyConstraintsAddTransitional       = "assemblyConstraints.addTransitional"
	MethodAssemblyConstraintsAddCustom             = "assemblyConstraints.addCustom"
	MethodAssemblyConstraintsDelete                = "assemblyConstraints.delete"
	MethodAssemblyConstraintsSetLimits             = "assemblyConstraints.setLimits"
	MethodAssemblyConstraintsSolve                 = "assemblyConstraints.solve"
	MethodAssemblyConstraintsHealth                = "assemblyConstraints.health"

	// Assembly joints (M12-F02, Oblikovati#359/#364): the simplified joints that establish
	// a degree-of-freedom set between two occurrences, and the DS-joint (DOF/imposed-motion)
	// view. Joints and constraints solve together (assemblyConstraints.solve/health).
	MethodAssemblyJointsList           = "assemblyJoints.list"
	MethodAssemblyJointsAddRigid       = "assemblyJoints.addRigid"
	MethodAssemblyJointsAddRotational  = "assemblyJoints.addRotational"
	MethodAssemblyJointsAddSlider      = "assemblyJoints.addSlider"
	MethodAssemblyJointsAddCylindrical = "assemblyJoints.addCylindrical"
	MethodAssemblyJointsAddPlanar      = "assemblyJoints.addPlanar"
	MethodAssemblyJointsAddBall        = "assemblyJoints.addBall"
	MethodAssemblyJointsDelete         = "assemblyJoints.delete"
	MethodAssemblyJointsSetLimits      = "assemblyJoints.setLimits"
	MethodAssemblyJointsSetFlip        = "assemblyJoints.setFlip"

	MethodDSJointsList             = "dsJoints.list"
	MethodDSJointsAdd              = "dsJoints.add"
	MethodDSJointsSetImposedMotion = "dsJoints.setImposedMotion"
	MethodDSJointsDelete           = "dsJoints.delete"

	// Assembly drive (M12-F03, Oblikovati#366): sweep a joint's driven variable through a
	// range, re-solving each step, to animate the assembly (kinematic motion study) with an
	// optional collision-stop.
	MethodAssemblyDrivePreview = "assemblyDrive.preview"

	// Assembly representations (M12-F04, Oblikovati#361/#367): the three override-layer
	// families — design-view, positional, level-of-detail — plus model states selecting one
	// of each. Capture snapshots the current scene; activate applies a representation.
	MethodDesignRepsCapture       = "designReps.capture"
	MethodDesignRepsActivate      = "designReps.activate"
	MethodDesignRepsList          = "designReps.list"
	MethodDesignRepsDelete        = "designReps.delete"
	MethodDesignRepsSetVisibility = "designReps.setVisibility"
	MethodDesignRepsSetAppearance = "designReps.setAppearance"
	MethodDesignRepsAddSection    = "designReps.addSection"

	MethodPositionalRepsCapture     = "positionalReps.capture"
	MethodPositionalRepsActivate    = "positionalReps.activate"
	MethodPositionalRepsList        = "positionalReps.list"
	MethodPositionalRepsDelete      = "positionalReps.delete"
	MethodPositionalRepsSetOverride = "positionalReps.setOverride"
	MethodPositionalRepsSetFlexible = "positionalReps.setFlexible"

	MethodLODRepsCapture       = "lodReps.capture"
	MethodLODRepsActivate      = "lodReps.activate"
	MethodLODRepsList          = "lodReps.list"
	MethodLODRepsDelete        = "lodReps.delete"
	MethodLODRepsSetSuppressed = "lodReps.setSuppressed"

	MethodModelStatesCreate   = "modelStates.create"
	MethodModelStatesActivate = "modelStates.activate"
	MethodModelStatesList     = "modelStates.list"
	MethodModelStatesDelete   = "modelStates.delete"

	// Assembly contact & interference (M12-F05, Oblikovati#362/#368): contact sets (resist
	// interpenetration when dragged), the contact-solver toggle, and static interference
	// analysis (overlapping volumes between occurrences).
	MethodContactSetsCreate       = "contactSets.create"
	MethodContactSetsList         = "contactSets.list"
	MethodContactSetsDelete       = "contactSets.delete"
	MethodContactSetsAddMember    = "contactSets.addMember"
	MethodContactSetsRemoveMember = "contactSets.removeMember"
	MethodContactSolverSetEnabled = "contactSolver.setEnabled"
	MethodContactSolverStatus     = "contactSolver.status"
	MethodInterferenceAnalyze     = "interference.analyze"

	MethodWorkPlanesList     = "workPlanes.list"
	MethodWorkPlanesCreate   = "workPlanes.create"
	MethodWorkPlanesRedefine = "workPlanes.redefine"

	MethodWorkPointsCreate = "workPoints.create"

	MethodWorkSurfacesList       = "workSurfaces.list"
	MethodWorkSurfacesGet        = "workSurfaces.get"
	MethodWorkSurfacesSetVisible = "workSurfaces.setVisible"
	MethodWorkSurfacesRename     = "workSurfaces.rename"

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
	MethodViewportCaptureWindow  = "viewport.captureWindow"
	MethodViewportSetNormalDebug = "viewport.setNormalDebug"
	MethodViewportSetMeshColors  = "viewport.setMeshColors"

	MethodViewsList      = "views.list"
	MethodViewsAdd       = "views.add"
	MethodViewsActivate  = "views.activate"
	MethodViewsClose     = "views.close"
	MethodViewsRename    = "views.rename"
	MethodViewsGetLayout = "views.getLayout"
	MethodViewsSetLayout = "views.setLayout"

	// Named views & standard orientations (M16-F03, Oblikovati#404/#409): capture the active
	// camera under a name and restore it exactly; jump to a standard orientation (front/top/iso).
	MethodViewsCaptureNamed  = "views.captureNamed"
	MethodViewsListNamed     = "views.listNamed"
	MethodViewsRestoreNamed  = "views.restoreNamed"
	MethodViewsDeleteNamed   = "views.deleteNamed"
	MethodViewSetOrientation = "view.setOrientation"

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

	// Body topology and queries (M07-F06/F07, Oblikovati/Oblikovati#629/#630).
	MethodBodyList             = "body.list"
	MethodBodySetVisible       = "body.setVisible"
	MethodBodyRename           = "body.rename"
	MethodBodyDelete           = "body.delete"
	MethodBodyPhysicalProps    = "body.physicalProperties"
	MethodBodyShells           = "body.shells"
	MethodBodyWires            = "body.wires"
	MethodWireOffsetPlanar     = "wire.offsetPlanar"
	MethodBodyLocateUsingPoint = "body.locateUsingPoint"
	MethodBodyFindUsingRay     = "body.findUsingRay"
	MethodBodyIsPointInside    = "body.isPointInside"
	MethodBodyConvexityEdges   = "body.convexityEdges"
	MethodBodyMinimumDistance  = "body.minimumDistance"
	MethodBodyValidate         = "body.validate"
	MethodBodyRangeBox         = "body.rangeBox"
	MethodBodyBindTransientKey = "body.bindTransientKey"

	// Facet/stroke calculation and retrieval (M07-F03 remainder,
	// Oblikovati/Oblikovati#293): the calculate variants cache per tolerance,
	// the existing variants retrieve without re-faceting.
	MethodBodyCalculateFacets  = "body.calculateFacets"
	MethodBodyExistingFacets   = "body.existingFacets"
	MethodBodyFacetTolerances  = "body.facetTolerances"
	MethodBodyCalculateStrokes = "body.calculateStrokes"
	MethodBodyExistingStrokes  = "body.existingStrokes"
	MethodBodyStrokeTolerances = "body.strokeTolerances"
	MethodFaceCalculateFacets  = "face.calculateFacets"
	MethodFaceCalculateStrokes = "face.calculateStrokes"

	// Batched surface evaluation of one document face by reference key (point,
	// normal, tangents, and point projection) — the out-of-process projection of
	// the in-proc surface-evaluator contract, for surface-following toolpaths and
	// point-projection queries.
	MethodBodyFaceEvaluate = "body.faceEvaluate"

	// The transient B-rep factory (M07-F05, Oblikovati/Oblikovati#628):
	// ownerless bodies addressed by session handles.
	MethodBrepCreatePrimitive      = "brep.createPrimitive"
	MethodBrepBoolean              = "brep.boolean"
	MethodBrepTransform            = "brep.transform"
	MethodBrepCopy                 = "brep.copy"
	MethodBrepSectionWithPlane     = "brep.sectionWithPlane"
	MethodBrepDeleteFaces          = "brep.deleteFaces"
	MethodBrepSilhouette           = "brep.silhouette"
	MethodBrepRuledSurface         = "brep.ruledSurface"
	MethodBrepOffsetFaces          = "brep.offsetFaces"
	MethodBrepImprint              = "brep.imprint"
	MethodBrepIdenticalBodies      = "brep.identicalBodies"
	MethodBrepCreateFromDefinition = "brep.createFromDefinition"
	MethodBrepDescribe             = "brep.describe"
	MethodBrepList                 = "brep.list"
	MethodBrepDelete               = "brep.delete"

	MethodClientGraphicsSet        = "clientGraphics.set"
	MethodClientGraphicsList       = "clientGraphics.list"
	MethodClientGraphicsDelete     = "clientGraphics.delete"
	MethodClientGraphicsSetVisible = "clientGraphics.setVisible"

	// Client-graphics object model — targeted retained-mode mutations and the named
	// color-mapper registry (M16-F05, Oblikovati#641). These move/toggle/flag a node without
	// resubmitting its (possibly large) mesh; the group transport stays the bulk clientGraphics.set.
	MethodGraphicsNodeSetTransform     = "graphicsNode.setTransform"
	MethodGraphicsNodeSetVisible       = "graphicsNode.setVisible"
	MethodGraphicsNodeSetSelectable    = "graphicsNode.setSelectable"
	MethodClientGraphicsRegisterMapper = "clientGraphics.registerColorMapper"
	MethodClientGraphicsListMappers    = "clientGraphics.listColorMappers"

	MethodInteractionGraphicsUpdate = "interactionGraphics.update"
	MethodInteractionGraphicsClear  = "interactionGraphics.clear"

	MethodTransactionUndo  = "transaction.undo"
	MethodTransactionRedo  = "transaction.redo"
	MethodTransactionState = "transaction.state"
	MethodTransactionBegin = "transaction.begin"
	MethodTransactionEnd   = "transaction.end"
	// MethodTransactionAbort discards the innermost open bounded transaction —
	// the model reverts to the group's pre-Begin state instead of committing it
	// (M04-F05, Oblikovati#613). Returns the resulting [UndoState].
	MethodTransactionAbort = "transaction.abort"
	// MethodTransactionHistory reads one open document's whole undo stream for a history
	// browser (every step since open, with save checkpoints flagged), without activating it.
	// Returns [TransactionHistory].
	MethodTransactionHistory = "transaction.history"
	// MethodTransactionJumpTo moves one document's undo cursor to an absolute position,
	// undoing or redoing as many steps as needed in one call. Returns [TransactionHistory].
	MethodTransactionJumpTo = "transaction.jumpTo"

	MethodInteractionState     = "interaction.state"
	MethodInteractionSetNotice = "interaction.setNotice"

	MethodScriptRun = "scripts.run"

	MethodLogsTail = "logs.tail"

	// Analysis & measurement (M18-F01 #423): engineering analysis on the model. Mass properties
	// (volume, surface area, centre of mass, mass) of the active part; measurement of entities.
	MethodAnalysisMassProperties = "analysis.massProperties"
	MethodAnalysisMeasure        = "analysis.measure"
	// Model health aggregation (M18-F02 #430): enumerate the active part's unhealthy features.
	MethodAnalysisModelHealth = "analysis.modelHealth"

	// Add-in registry & automation (M05-F01: #245, #251, #252).
	MethodAddInsList            = "addins.list"
	MethodAddInsGet             = "addins.get"
	MethodAddInsActivate        = "addins.activate"
	MethodAddInsDeactivate      = "addins.deactivate"
	MethodAddInsSetLoadBehavior = "addins.setLoadBehavior"
	MethodAddInsCallAutomation  = "addins.callAutomation"

	// Host application info an add-in queries at runtime. apiVersion reports the
	// full api.Version the host implements, so an add-in compatible at the major
	// boundary can still adapt to minor/patch differences.
	MethodApplicationApiVersion = "application.apiVersion"

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
	MethodDockableWindowsSetValue   = "dockableWindows.setValue"
	MethodDockableWindowsDelete     = "dockableWindows.delete"
	MethodDockableWindowsList       = "dockableWindows.list"

	// UI environments (M05-F03, #247; add-in environments: Oblikovati#667).
	MethodUIListEnvironments = "ui.listEnvironments"

	// Application options (M05-F11, #618).
	MethodOptionsListGroups = "options.listGroups"
	MethodOptionsGetGroup   = "options.getGroup"
	MethodOptionsSetGroup   = "options.setGroup"

	// Command alias & keyboard-shortcut customization (M05-F17, #831).
	MethodKeymapList     = "keymap.list"
	MethodKeymapSetChord = "keymap.setChord"
	MethodKeymapSetAlias = "keymap.setAlias"
	MethodKeymapReset    = "keymap.reset"
	MethodKeymapResetAll = "keymap.resetAll"
	MethodKeymapExport   = "keymap.export"
	MethodKeymapImport   = "keymap.import"

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

	// Application color schemes (M16-F06, Oblikovati#642): the named palettes (background,
	// highlight, selection colors) the viewport and selection pipeline traffic in.
	MethodColorSchemesList      = "colorSchemes.list"
	MethodColorSchemesGetActive = "colorSchemes.getActive"
	MethodColorSchemesSetActive = "colorSchemes.setActive"

	// Display options & settings (M16-F07, Oblikovati#643): the application-level display
	// options and the per-document display settings (background, edges, ground plane, shadows)
	// that parameterize the M23 display modes.
	MethodDisplayGetOptions          = "display.getOptions"
	MethodDisplaySetOptions          = "display.setOptions"
	MethodDocumentGetDisplaySettings = "document.getDisplaySettings"
	MethodDocumentSetDisplaySettings = "document.setDisplaySettings"

	// Styles & standards (M16-F02, Oblikovati#403/#408): the document's color styles, the
	// style-library cascade, and library import. (Lighting styles use the lighting.* methods.)
	MethodStylesList          = "styles.list"
	MethodStylesGet           = "styles.get"
	MethodStylesSet           = "styles.set"
	MethodStylesDelete        = "styles.delete"
	MethodStylesListLibraries = "styles.listLibraries"
	MethodStylesImportLibrary = "styles.importLibrary"

	// Attached point clouds (M17-F06, Oblikovati/Oblikovati#645): laser-scan / photogrammetry
	// references on the active part — attach, query, place, and budget their display. They
	// operate on the active part's cloud collection, keyed by unique name.
	MethodPointCloudsAttach         = "pointClouds.attach"
	MethodPointCloudsList           = "pointClouds.list"
	MethodPointCloudsGet            = "pointClouds.get"
	MethodPointCloudsDelete         = "pointClouds.delete"
	MethodPointCloudsSetVisible     = "pointClouds.setVisible"
	MethodPointCloudsSetTransform   = "pointClouds.setTransform"
	MethodPointCloudsSetScale       = "pointClouds.setScale"
	MethodPointCloudsSetDensity     = "pointClouds.setDensity"
	MethodPointCloudsToModelSpace   = "pointClouds.toModelSpace"
	MethodPointCloudsFromModelSpace = "pointClouds.fromModelSpace"

	// Crop volumes that limit a cloud's display to a working region (M17-F06, #645).
	MethodPointCloudsAddCrop       = "pointClouds.addCrop"
	MethodPointCloudsListCrops     = "pointClouds.listCrops"
	MethodPointCloudsDeleteCrop    = "pointClouds.deleteCrop"
	MethodPointCloudsSetCropActive = "pointClouds.setCropActive"

	// Derive model geometry from a cloud's scanned points (M17-F06, #645).
	MethodPointCloudsFitPlane     = "pointClouds.fitPlane"
	MethodPointCloudsNearestPoint = "pointClouds.nearestPoint"
)

// Push-event type tags. These name host→add-in events delivered to the add-in's Notify
// entry point (ADR-0016), not callable request/response methods — there is no client
// method for them; an add-in matches on the event's "type" field. See the DTOs in this
// package (e.g. [EditCommittedEvent]).
const (
	EventEditCommitted = "edit.committed"

	// Document lifecycle + modeling events (#148): the DocumentEvents/ApplicationEvents/
	// ModelingEvents wave. Created/Opened/Saved/Closed/Activated fire as a document moves
	// through its life; ModelChanged carries a committed batch of model changes (features,
	// sketches, parameters) on a document — an add-in re-queries the affected document.
	EventDocumentCreated   = "document.created"
	EventDocumentOpened    = "document.opened"
	EventDocumentSaved     = "document.saved"
	EventDocumentClosed    = "document.closed"
	EventDocumentActivated = "document.activated"
	EventModelChanged      = "model.changed"
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
	// EventCommandEnded reports a command finishing (M05-F12 #619); the payload carries the command id
	// and whether it failed. The host emits it as the after-pair of EventCommandStarted.
	EventCommandEnded = "command.ended"
	// EventPanelValueChanged reports the user editing an editable dockable-window control
	// (M05-F03): the add-in receives the window id, control id, and new value.
	EventPanelValueChanged = "panel.valueChanged"
	// EventSelectionChanged reports the selection set changing (M05-F12 #619).
	EventSelectionChanged = "selection.changed"
	// EventParameterChanged reports a parameter's expression/value changing (#148) — the granular
	// notification beyond the generic edit.committed; the payload carries the parameter's new state.
	EventParameterChanged = "parameters.changed"
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

	// Transaction lifecycle events (see [TransactionEventPayload], M04-F05
	// Oblikovati#613): every move of a document's transaction stream.
	EventTransactionCommitted = "transaction.committed"
	EventTransactionUndone    = "transaction.undone"
	EventTransactionRedone    = "transaction.redone"
	EventTransactionAborted   = "transaction.aborted"
	EventTransactionDeleted   = "transaction.deleted"

	// File-access events (see [FileResolutionEventPayload] and
	// [FileDirtyEventPayload], M04-F05 Oblikovati#613).
	EventFileResolution = "file.resolution"
	EventFileDirty      = "file.dirty"

	// File-UI hook events (see [FileDialogHookPayload], M04-F05
	// Oblikovati#613): the new/open/save-as flows and their dialogs.
	EventFileNew              = "file.new"
	EventFileNewDialog        = "file.newDialog"
	EventFileOpenDialog       = "file.openDialog"
	EventFileSaveAsDialog     = "file.saveAsDialog"
	EventFileOpenFromMRU      = "file.openFromMRU"
	EventFilePopulateMetadata = "file.populateMetadata"

	// Assembly occurrence-lifecycle events (see [OccurrenceEventPayload], M11-F07
	// Oblikovati#723): a component was placed, removed, replaced, moved, or had its
	// suppression toggled. The host coalesces a solver drag into one transformed event.
	EventOccurrenceAdded       = "occurrence.added"
	EventOccurrenceDeleted     = "occurrence.deleted"
	EventOccurrenceReplaced    = "occurrence.replaced"
	EventOccurrenceTransformed = "occurrence.transformed"
	EventOccurrenceSuppressed  = "occurrence.suppressed"

	// Assembly feature-program event (see [AssemblyFeaturesChangedEvent], M11-F08
	// Oblikovati#725): the assembly's machining-feature program was re-evaluated.
	EventAssemblyFeaturesChanged = "assemblyFeatures.changed"

	// Assembly relationship events (see [ConstraintEventPayload], M12-F01
	// Oblikovati#358/#363): a constraint was added or deleted, or the assembly was
	// re-solved (occurrence placements may have changed — also signalled by the
	// occurrence.transformed events the solve raises).
	EventAssemblyConstraintAdded   = "assemblyConstraints.added"
	EventAssemblyConstraintDeleted = "assemblyConstraints.deleted"
	EventAssemblyResolved          = "assembly.resolved"

	// Assembly joint events (see [JointEventPayload], M12-F02 Oblikovati#359/#364): a joint
	// was added or removed (positions also change via the occurrence.transformed events the
	// combined solve raises).
	EventAssemblyJointAdded   = "assemblyJoints.added"
	EventAssemblyJointDeleted = "assemblyJoints.deleted"

	// Representation/model-state events (M12-F04, Oblikovati#361/#367).
	EventRepresentationCaptured  = "representations.captured"
	EventRepresentationActivated = "representations.activated"
	EventModelStateActivated     = "modelStates.activated"

	// Style events (see [StyleChangedEvent], M16-F02 Oblikovati#403/#408): a color or lighting
	// style was added, edited, or deleted — consumers re-resolve their styling.
	EventStyleAdded   = "style.added"
	EventStyleChanged = "style.changed"
	EventStyleDeleted = "style.deleted"

	// Camera event (see [CameraChangedEvent], M16-F03 Oblikovati#404/#409): the active view's
	// camera moved (orbit/pan/zoom/fit/named-view restore) — collaboration and overlay add-ins
	// re-sync to the new frame.
	EventCameraChanged = "camera.changed"

	// Modeling events (#148): the ModelingEvents/SketchEvents wave, granular beyond the batched
	// model.changed. A feature was created, edited, or deleted (see [FeatureLifecycleEvent]); the
	// document's sketch-edit mode was entered or exited (see [SketchEditEvent]). Like edit.committed
	// (ADR-0004), the v1 scope is the host method router: feature events fire for features.add/edit/
	// delete, and sketch-edit events fire whenever the host enters/exits a sketch (UI or add-in driven).
	EventFeatureAdded      = "feature.added"
	EventFeatureEdited     = "feature.edited"
	EventFeatureDeleted    = "feature.deleted"
	EventSketchEditEntered = "sketch.editEntered"
	EventSketchEditExited  = "sketch.editExited"
)

// OKResult is the trivial success payload for mutating methods with no return value.
type OKResult struct {
	OK bool `json:"ok"`
}
