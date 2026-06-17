// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The drawing view surface (M14-F02 PBI-139, Oblikovati/Oblikovati#386): a drawing sheet
// carries views that project the referenced 3D model — a base view from a standard
// orientation and projected views derived from it. Each view holds drawing curves (visible /
// hidden edge segments) computed by the hidden-line engine. An in-proc consumer reads the
// view metadata directly; creation and the curve geometry travel over api/wire
// (drawingViews.*). The host owns the projection (/source, model/drawing + kernel/hlr).

// DrawingView is one view on a sheet: its name, the orientation it projects from, its scale,
// rendering style, sheet position (millimetres) and the number of drawing curves it produced.
type DrawingView interface {
	// Name is the view's display name (unique within the drawing).
	Name() string
	// Type discriminates the view kind (base, projected, auxiliary, section, detail, break).
	Type() types.DrawingViewType
	// ParentView is the name of the view this one derives from (projected/auxiliary/section/
	// detail), or "" for a base view.
	ParentView() string
	// IsProjected reports whether this is a projected view (derived from a base view) rather
	// than a base view.
	IsProjected() bool
	// Orientation is the standard orientation the view projects from.
	Orientation() types.BaseViewOrientation
	// Scale is the model→sheet scale factor (e.g. 0.5 for 1:2).
	Scale() float64
	// Style is the view's rendering style.
	Style() types.DrawingViewStyle
	// CenterMM is the view centre on the sheet, in millimetres.
	CenterMM() (x, y float64)
	// CurveCount is the number of drawing curves (visible + hidden) the view holds.
	CurveCount() int
}

// SectionDrawingView is a view cut from a parent view by a section line: the model is sliced by
// the plane through that line (perpendicular to the parent), the near half removed, the cut
// outline drawn bold and the exposed faces hatched. The reference contracts give section views a
// distinct interface (unlike auxiliary/projected, which are a plain DrawingView + a type tag).
type SectionDrawingView interface {
	DrawingView
	// SectionLineMM is the cut line on the parent view, in sheet millimetres.
	SectionLineMM() (x1, y1, x2, y2 float64)
}

// DetailDrawingView is a magnified view of a circular region of a parent view: the parent's
// projection clipped to the boundary circle and re-placed at a larger scale. Like section, the
// reference API gives it its own interface.
type DetailDrawingView interface {
	DrawingView
	// DetailBoundaryMM is the magnified region's circle on the parent view (sheet millimetres):
	// centre and radius.
	DetailBoundaryMM() (cx, cy, r float64)
}
