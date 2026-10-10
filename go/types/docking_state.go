// SPDX-License-Identifier: Apache-2.0

package types

// DockingState is where a dockable window sits — the DockingStateEnum equivalent,
// narrowed to the positions the host's dockspace honors (M05-F03, #247). It is the
// window's INITIAL placement; the user may re-dock it afterwards and the host's
// layout persistence keeps their arrangement.
type DockingState uint8

const (
	// DockFloating is a free-floating window (the zero value).
	DockFloating DockingState = 0
	// DockLeft docks into the left band (beside the model browser).
	DockLeft DockingState = 1
	// DockRight docks into the right band.
	DockRight DockingState = 2
	// DockBottom docks into the bottom band (above the status bar).
	DockBottom DockingState = 3
)

var dockingStateNames = map[DockingState]string{
	DockFloating: "floating", DockLeft: "left", DockRight: "right", DockBottom: "bottom",
}

// String returns the state's stable name.
func (d DockingState) String() string {
	return enumName(dockingStateNames, d, "dockingState(?)")
}
