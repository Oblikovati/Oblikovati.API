// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The scalar read surface of the drawing document (M14-F01, Oblikovati/Oblikovati#384):
// a drawing holds an ordered set of sheets; each sheet has a size/orientation, an
// optional border, and an optional title block whose fields resolve against the
// drawing's primary referenced model. An in-proc consumer reads this structure
// directly; every mutation (addSheet/setActiveSheet/setModelReference/…) travels over
// api/wire (drawing.*). The host implementations live in /source (model/drawing).

// DrawingSheet is one sheet of a drawing: its name, size and orientation (which fix
// its width and height in millimetres), and its border/title-block presence.
type DrawingSheet interface {
	// Name is the sheet's display name (unique within the drawing).
	Name() string
	// Size is the standard sheet size, or types.SheetSizeCustom for an explicit
	// width/height.
	Size() types.SheetSize
	// Orientation is how the standard dimensions are laid out (portrait/landscape).
	Orientation() types.SheetOrientation
	// WidthMM and HeightMM are the laid-out dimensions in millimetres (orientation
	// already applied).
	WidthMM() float64
	HeightMM() float64
	// Border returns the sheet's border, or nil if it has none.
	Border() DrawingBorder
	// TitleBlock returns the sheet's title block, or nil if it has none.
	TitleBlock() DrawingTitleBlock
}

// DrawingBorder is a sheet's border: the printable-area margins (millimetres) inset
// from the sheet edge.
type DrawingBorder interface {
	// Margins returns the left, right, top and bottom inset in millimetres.
	Margins() (left, right, top, bottom float64)
}

// DrawingTitleBlock is a sheet's title block: a named definition and the resolved
// field values (each field's static text or the value pulled from the referenced
// model's iProperties).
type DrawingTitleBlock interface {
	// DefinitionName is the title-block definition this block instantiates.
	DefinitionName() string
	// FieldValue returns the resolved text of the named field, and whether the field
	// exists.
	FieldValue(name string) (string, bool)
}
