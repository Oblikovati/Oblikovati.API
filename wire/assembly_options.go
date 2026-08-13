// SPDX-License-Identifier: Apache-2.0

package wire

// Assembly editing options (#1981): the assembly-modeling option set that governs placement,
// adaptivity, update deferral, and section/opacity defaults (Inventor's AssemblyOptions). Two options
// change host behaviour — PlaceAndGroundFirstComponentAtOrigin grounds the first placed component, and
// DeferUpdate batches recomputes until it is cleared — the rest are stored defaults the head reads.

// AssemblyOptions is the assembly-editing option set.
type AssemblyOptions struct {
	PartFeaturesInitiallyAdaptive            bool   `json:"partFeaturesInitiallyAdaptive,omitempty"`
	DeferUpdate                              bool   `json:"deferUpdate,omitempty"`
	OnlyActiveComponentIsOpaque              bool   `json:"onlyActiveComponentIsOpaque,omitempty"`
	PlaceAndGroundFirstComponentAtOrigin     bool   `json:"placeAndGroundFirstComponentAtOrigin"`
	EnableConstraintRedundancyAnalysis       bool   `json:"enableConstraintRedundancyAnalysis"`
	DeleteComponentPatternSources            bool   `json:"deleteComponentPatternSources,omitempty"`
	SectionAllParts                          bool   `json:"sectionAllParts,omitempty"`
	UseLastOccurrenceOrientationForPlacement bool   `json:"useLastOccurrenceOrientationForPlacement,omitempty"`
	DefaultLevelOfDetail                     string `json:"defaultLevelOfDetail,omitempty"`
	DefaultDesignView                        string `json:"defaultDesignView,omitempty"`
}

// AssemblyOptionsResult is the reply of [MethodAssemblyOptionsGet] and [MethodAssemblyOptionsSet]:
// the assembly's current editing options.
type AssemblyOptionsResult struct {
	Options AssemblyOptions `json:"options"`
}

// SetAssemblyOptionsArgs is the request of [MethodAssemblyOptionsSet]: replace the assembly's editing
// options. Clearing DeferUpdate flushes any recompute deferred while it was set.
type SetAssemblyOptionsArgs struct {
	Options AssemblyOptions `json:"options"`
}
