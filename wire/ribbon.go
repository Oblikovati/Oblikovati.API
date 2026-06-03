// SPDX-License-Identifier: Apache-2.0

package wire

import "github.com/Oblikovati/api/types"

// RibbonControlInfo is one command control (button) on a panel — its command id and display
// name, so an add-in can discover what is already there before inserting next to it.
type RibbonControlInfo struct {
	CommandID   string `json:"commandId"`
	DisplayName string `json:"displayName"`
}

// RibbonPanelInfo is one panel within a tab (Inventor's RibbonPanel) and its controls.
type RibbonPanelInfo struct {
	Name     string              `json:"name"`
	Controls []RibbonControlInfo `json:"controls,omitempty"`
}

// RibbonTabInfo is one tab within a ribbon (Inventor's RibbonTab) and its panels.
type RibbonTabInfo struct {
	Name   string            `json:"name"`
	Panels []RibbonPanelInfo `json:"panels,omitempty"`
}

// ListRibbonResult is the response of [MethodRibbonList]: the ribbon currently shown for the
// active document (Key is ZeroDoc when none is open), with its tabs/panels/controls. It is
// the discovery surface an add-in reads to find the internal names to insert into — Inventor's
// "list the contents of the ribbon" (RibbonUI_Overview). Contextual tabs (e.g. Sketch) appear
// only when their environment is active.
type ListRibbonResult struct {
	Key  types.RibbonKey `json:"key"`
	Tabs []RibbonTabInfo `json:"tabs,omitempty"`
}
