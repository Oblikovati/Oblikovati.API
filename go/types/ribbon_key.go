// SPDX-License-Identifier: Apache-2.0

package types

// RibbonKey is the internal name of one of the host's ribbons. There is one ribbon per
// document type plus ZeroDoc (shown when no document is open); the active ribbon is selected
// by the active document. An add-in targets a ribbon by this name when placing a control.
// The names are stable wire values and must not be renamed.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases it
// (app.RibbonKey) so existing call sites are unaffected.
type RibbonKey string

const (
	// ZeroDocRibbon is shown when no document is open (the Get Started ribbon).
	ZeroDocRibbon RibbonKey = "ZeroDoc"
	// PartRibbon is shown for a part document.
	PartRibbon RibbonKey = "Part"
	// AssemblyRibbon is shown for an assembly document.
	AssemblyRibbon RibbonKey = "Assembly"
	// DrawingRibbon is shown for a drawing document.
	DrawingRibbon RibbonKey = "Drawing"
	// PresentationRibbon is shown for a presentation document.
	PresentationRibbon RibbonKey = "Presentation"
	// IFeaturesRibbon is shown for an iFeature authoring document.
	IFeaturesRibbon RibbonKey = "iFeatures"
	// UnknownDocumentRibbon is shown for a document whose type is not resolved.
	UnknownDocumentRibbon RibbonKey = "UnknownDocument"
)

// Valid reports whether k is one of the seven known ribbon names — used by the host to
// reject an add-in placing a control on a nonexistent ribbon.
func (k RibbonKey) Valid() bool {
	switch k {
	case ZeroDocRibbon, PartRibbon, AssemblyRibbon, DrawingRibbon,
		PresentationRibbon, IFeaturesRibbon, UnknownDocumentRibbon:
		return true
	default:
		return false
	}
}
