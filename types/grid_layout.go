// SPDX-License-Identifier: Apache-2.0

package types

// CSS-grid-like layout for nested panel controls (ADR-0019, ADR-0020). A PanelGrid
// container declares column tracks; its children fill cells, either by explicit Cell
// placement or by auto-flow (left-to-right, wrapping at the column count). Faithful to
// CSS Grid where it pays off — column tracks carry the full auto/fixed/fraction(+minmax)
// model — with one deliberate deviation: rows are content-height auto-flow, not tracks,
// because the host renders in immediate mode where column widths are first-class and row
// heights are content-driven. Column span ships; row span is reserved (see GridCell).

// GridTrackKind is how one grid column track is sized.
type GridTrackKind uint8

const (
	// GridTrackAuto sizes the track to its content (the zero value).
	GridTrackAuto GridTrackKind = 0
	// GridTrackFixed sizes the track to exactly Value pixels.
	GridTrackFixed GridTrackKind = 1
	// GridTrackFraction sizes the track to a share of the leftover space; Value is the
	// fraction weight (CSS "fr"): a 2fr track gets twice the leftover a 1fr track gets.
	GridTrackFraction GridTrackKind = 2
)

// GridTrack is one column track of a grid. Value is pixels when Fixed and the fraction
// weight when Fraction (ignored when Auto). MinPx/MaxPx (0 = unset) clamp the resolved
// width, expressing CSS minmax(): e.g. a 1fr track that never shrinks below 80px is
// {Kind: GridTrackFraction, Value: 1, MinPx: 80}.
//
// Example: a label/field form is two tracks — TrackAuto() for labels, TrackFr(1) for fields.
type GridTrack struct {
	Kind  GridTrackKind `json:"kind,omitempty"`
	Value float64       `json:"value,omitempty"`
	MinPx float64       `json:"minPx,omitempty"`
	MaxPx float64       `json:"maxPx,omitempty"`
}

// GridCell is the explicit placement of one child inside its parent grid. Col is the
// 0-based column; ColSpan (>=1; 0 is treated as 1) is how many columns the child covers.
// A child with no Cell (nil pointer on the wire) auto-flows into the next free cell.
//
// Row and RowSpan are reserved for a later slice: the wire shape accepts them so it stays
// forward-stable, but the host renders rows as content-height auto-flow, so RowSpan must be
// 1 in this version (the router rejects RowSpan > 1).
type GridCell struct {
	Col     int `json:"col,omitempty"`
	Row     int `json:"row,omitempty"`
	ColSpan int `json:"colSpan,omitempty"`
	RowSpan int `json:"rowSpan,omitempty"`
}
