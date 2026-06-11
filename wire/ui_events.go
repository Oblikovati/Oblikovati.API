// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The UI shell surface of M05-F12 (#619): command search, the radial marking
// menu, context-menu injection, object-visibility toggles, and the app-level
// input/UI push events.

// SearchCommandsArgs is the request of [MethodUISearch]: find registered commands
// whose display name, id, or alias matches the query (case-insensitive substring).
type SearchCommandsArgs struct {
	Query string `json:"query"`
}

// SearchCommandsResult is the response of [MethodUISearch].
type SearchCommandsResult struct {
	Commands []CommandInfo `json:"commands"`
}

// MarkingMenuItem is one radial slot: the command it runs, placed at a quadrant.
type MarkingMenuItem struct {
	Quadrant  types.ScreenQuadrant `json:"quadrant"`
	CommandID string               `json:"commandId"`
}

// MarkingMenuView is one environment's radial marking menu: up to eight quadrant
// slots plus the linear overflow items below them.
type MarkingMenuView struct {
	Environment types.Environment `json:"environment"`
	Quadrants   []MarkingMenuItem `json:"quadrants,omitempty"`
	Overflow    []string          `json:"overflow,omitempty"`
}

// GetMarkingMenuArgs is the request of [MethodUIGetMarkingMenu].
type GetMarkingMenuArgs struct {
	Environment types.Environment `json:"environment"`
}

// SetMarkingMenuArgs is the request of [MethodUISetMarkingMenu]: replace one
// environment's radial menu (the customization hook behind OnRadialMarkingMenu).
type SetMarkingMenuArgs struct {
	Menu MarkingMenuView `json:"menu"`
}

// ContextMenuItemSpec is one injected context-menu entry: the label shown and the
// command it runs.
type ContextMenuItemSpec struct {
	Label     string `json:"label"`
	CommandID string `json:"commandId"`
}

// SetContextMenuArgs is the request of [MethodUISetContextMenu]: replace an
// add-in's injected entries for one browser node kind (e.g. "feature", "sketch",
// "body"; "" injects into every node's menu) — the OnContextMenu equivalent,
// declarative so no round-trip happens while the menu is open.
type SetContextMenuArgs struct {
	AddIn string                `json:"addin"`
	Kind  string                `json:"kind"`
	Items []ContextMenuItemSpec `json:"items"`
}

// ObjectVisibilityView is the View ▸ Object-visibility toggle group: which datum
// and sketch overlays draw (hidden geometry is also not pickable).
type ObjectVisibilityView struct {
	WorkPlanes bool `json:"workPlanes"`
	WorkAxes   bool `json:"workAxes"`
	WorkPoints bool `json:"workPoints"`
	Sketches   bool `json:"sketches"`
}

// SetObjectVisibilityArgs is the request of [MethodUISetObjectVisibility].
type SetObjectVisibilityArgs struct {
	Visibility ObjectVisibilityView `json:"visibility"`
}

// CommandStartedEvent is the push event (type [EventCommandStarted]) fired when a
// command begins — the OnStartCommand observation; command.ended already reports
// completion.
type CommandStartedEvent struct {
	Type    string `json:"type"` // always EventCommandStarted
	Command string `json:"command"`
}

// SelectionChangedEvent is the push event (type [EventSelectionChanged]) fired
// when the selection set changes — the app-level OnSelect/OnUnSelect observation.
type SelectionChangedEvent struct {
	Type  string `json:"type"` // always EventSelectionChanged
	Count int    `json:"count"`
}

// EnvironmentChangedEvent is the push event (type [EventEnvironmentChanged]) fired
// when the UI environment switches (base ↔ sketch) — the OnEnvironmentChange
// equivalent, so add-ins re-aim their contextual UI.
type EnvironmentChangedEvent struct {
	Type        string            `json:"type"` // always EventEnvironmentChanged
	Environment types.Environment `json:"environment"`
}

// RegisterEnvironmentArgs is the request of [MethodUIRegisterEnvironment]: an
// add-in declaring its own contextual UI environment (M05-F16, Oblikovati#667).
// Environment must be ≥ 2 (0 and 1 are the built-in base/sketch); commands
// registered with the value form the environment's contextual tabs.
type RegisterEnvironmentArgs struct {
	Environment types.Environment `json:"environment"`
	Name        string            `json:"name"`
}

// ActivateEnvironmentArgs is the request of [MethodUIActivateEnvironment]: enter a
// REGISTERED add-in environment (its contextual tabs appear); the base value (0)
// leaves it.
type ActivateEnvironmentArgs struct {
	Environment types.Environment `json:"environment"`
}
