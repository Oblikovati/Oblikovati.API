// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// RibbonItemInfo is one entry of a control's dropdown — a split button's variant or
// a popup control's menu item — with its live enabled state.
type RibbonItemInfo struct {
	CommandID string `json:"commandId"`
	Label     string `json:"label"`
	Tooltip   string `json:"tooltip,omitempty"`
	Enabled   bool   `json:"enabled"`
}

// RibbonControlInfo is one command control on a panel: the typed ribbon object
// model's leaf (M05-F03, #247). Beyond the command id and display name (the
// original discovery surface), it carries the control kind, look, and live state,
// so an add-in can mirror or extend the ribbon without guessing. Items holds a
// split button's variants or a popup control's menu entries.
type RibbonControlInfo struct {
	CommandID   string            `json:"commandId"`
	DisplayName string            `json:"displayName"`
	Kind        types.ControlKind `json:"kind,omitempty"`
	ButtonStyle types.ButtonStyle `json:"buttonStyle,omitempty"`
	Icon        string            `json:"icon,omitempty"`
	Tooltip     string            `json:"tooltip,omitempty"`
	Alias       string            `json:"alias,omitempty"`
	Enabled     bool              `json:"enabled"`
	Active      bool              `json:"active,omitempty"`
	Items       []RibbonItemInfo  `json:"items,omitempty"`
}

// RibbonSelectorInfo is a panel rendered as one drop-down selection box (a panel of
// combo controls): its options and which is currently selected.
type RibbonSelectorInfo struct {
	Options       []RibbonItemInfo `json:"options"`
	SelectedIndex int              `json:"selectedIndex"`
}

// RibbonPanelInfo is one panel within a tab and its controls. Selector is set
// instead of Controls when the panel renders as a selection box.
type RibbonPanelInfo struct {
	Name     string              `json:"name"`
	Controls []RibbonControlInfo `json:"controls,omitempty"`
	Selector *RibbonSelectorInfo `json:"selector,omitempty"`
}

// RibbonTabInfo is one tab within a ribbon and its panels.
type RibbonTabInfo struct {
	Name   string            `json:"name"`
	Panels []RibbonPanelInfo `json:"panels,omitempty"`
}

// ListRibbonResult is the response of [MethodRibbonList]: the ribbon currently shown for the
// active document (Key is ZeroDoc when none is open), with its tabs/panels/controls. It is
// the discovery surface an add-in reads to find the internal names to insert into (it lists
// the contents of the ribbon). Contextual tabs (e.g. Sketch) appear only when their
// environment is active.
type ListRibbonResult struct {
	Key  types.RibbonKey `json:"key"`
	Tabs []RibbonTabInfo `json:"tabs,omitempty"`
}

// EnvironmentInfo is one entry of [MethodUIListEnvironments]: a UI environment the
// command framework scopes by (base = always shown; others are contextual tab
// sets), flagging the active one. Add-in-created environments are tracked by
// Oblikovati/Oblikovati#667.
type EnvironmentInfo struct {
	Environment types.Environment `json:"environment"`
	Name        string            `json:"name"`
	Active      bool              `json:"active"`
}

// ListEnvironmentsResult is the response of [MethodUIListEnvironments].
type ListEnvironmentsResult struct {
	Environments []EnvironmentInfo `json:"environments"`
}
