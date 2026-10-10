// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Container-control constructors for building nested, CSS-grid-like panel layouts
// declaratively (ADR-0019). A grid declares column tracks and holds children that either
// auto-flow or carry an explicit Cell; a group is a titled vertical stack; tabs treats
// each child as one tab. These compose with the flat leaf constructors in
// panel_controls.go — a grid cell can be any control, including another grid.

// TrackAuto sizes a column to its content.
func TrackAuto() types.GridTrack {
	return types.GridTrack{Kind: types.GridTrackAuto}
}

// TrackFixed sizes a column to exactly px pixels.
func TrackFixed(px float64) types.GridTrack {
	return types.GridTrack{Kind: types.GridTrackFixed, Value: px}
}

// TrackFr sizes a column to weight shares of the leftover space (CSS "fr").
func TrackFr(weight float64) types.GridTrack {
	return types.GridTrack{Kind: types.GridTrackFraction, Value: weight}
}

// TrackMinMax clamps a base track's resolved width to [minPx, maxPx] (CSS minmax()).
// A zero bound means unset, e.g. TrackMinMax(TrackFr(1), 80, 0) only enforces a floor.
func TrackMinMax(base types.GridTrack, minPx, maxPx float64) types.GridTrack {
	base.MinPx = minPx
	base.MaxPx = maxPx
	return base
}

// PanelGrid is a grid container: columns are the column tracks, colGap/rowGap the px
// spacing, children the cells. Children with no Cell auto-flow left-to-right, wrapping at
// len(columns); use PlaceAt to span or position a child explicitly.
//
// Example: a two-column form is
//
//	PanelGrid("form", []types.GridTrack{TrackAuto(), TrackFr(1)}, 6, 4,
//	    PanelLabel("l", "Start depth"), PanelTextBox("start", "", "0"))
func PanelGrid(id string, columns []types.GridTrack, colGap, rowGap float64, children ...wire.PanelControlSpec) wire.PanelControlSpec {
	return wire.PanelControlSpec{
		Kind: types.PanelGrid, ID: id, Columns: columns,
		ColumnGap: colGap, RowGap: rowGap, Children: children,
	}
}

// PanelGroup is a titled box that stacks its children vertically (the QGroupBox of this
// vocabulary). As a tabs child, title doubles as the tab caption.
func PanelGroup(id, title string, children ...wire.PanelControlSpec) wire.PanelControlSpec {
	return wire.PanelControlSpec{Kind: types.PanelGroup, ID: id, Title: title, Children: children}
}

// PanelTabs is a tab strip; each pane is a container child whose Title is its tab caption.
// Build panes with PanelTab, or pass any titled container.
func PanelTabs(id string, panes ...wire.PanelControlSpec) wire.PanelControlSpec {
	return wire.PanelControlSpec{Kind: types.PanelTabs, ID: id, Children: panes}
}

// PanelTab is one pane for PanelTabs: a group titled title stacking content.
func PanelTab(title string, content ...wire.PanelControlSpec) wire.PanelControlSpec {
	return PanelGroup("", title, content...)
}

// PlaceAt stamps an explicit grid placement on a child: column col, spanning colSpan
// columns. Without it a child auto-flows into the next free cell.
func PlaceAt(control wire.PanelControlSpec, col, colSpan int) wire.PanelControlSpec {
	control.Cell = &types.GridCell{Col: col, ColSpan: colSpan}
	return control
}
