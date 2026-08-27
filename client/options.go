// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"

	"oblikovati.org/api/wire"
)

// Options is the application-options operation group (M05-F11): typed per-user
// option groups — general (startup), display (color scheme + ViewCube), sketch
// (grid/snapping), part (modeling defaults) — read and written as whole groups.
type Options struct{ c *Client }

// Options returns the application-options operation group.
func (c *Client) Options() Options { return Options{c} }

// Groups returns the available option group names.
//
// mcp:tool options_list_groups
// mcp:summary Returns the available option group names.
func (o Options) Groups() (wire.ListOptionGroupsResult, error) {
	return call[wire.ListOptionGroupsResult](o.c, wire.MethodOptionsListGroups, nil)
}

// getGroup fetches one group and returns the union view.
//
// mcp:tool options_get_group
// mcp:summary Fetches one group and returns the union view.
func (o Options) getGroup(group string) (wire.OptionGroupView, error) {
	return call[wire.OptionGroupView](o.c, wire.MethodOptionsGetGroup, wire.GetOptionGroupArgs{Group: group})
}

// optionGroupField fetches group and returns the field sel selects from its
// union view, erroring when the reply carries no payload for it. group also
// names the payload in that error (wire.OptionGroupXxx are already the
// lowercase group words: "general", "display", …).
func optionGroupField[V any](o Options, group string, sel func(wire.OptionGroupView) *V) (V, error) {
	r, err := o.getGroup(group)
	if err != nil {
		var zero V
		return zero, err
	}
	if f := sel(r); f != nil {
		return *f, nil
	}
	var zero V
	return zero, fmt.Errorf("client: options.getGroup(%q) reply carries no %s payload", group, group)
}

// General returns the general options (startup behavior).
func (o Options) General() (wire.GeneralOptionsView, error) {
	return optionGroupField(o, wire.OptionGroupGeneral, func(r wire.OptionGroupView) *wire.GeneralOptionsView { return r.General })
}

// SetGeneral writes the general options.
//
//	client.Options().SetGeneral(wire.GeneralOptionsView{StartupAction: types.StartupEmptyWorkspace})
//
// mcp:tool options_set_group
// mcp:summary Writes the display options.
func (o Options) SetGeneral(v wire.GeneralOptionsView) (wire.OKResult, error) {
	args := wire.OptionGroupView{Group: wire.OptionGroupGeneral, General: &v}
	return call[wire.OKResult](o.c, wire.MethodOptionsSetGroup, args)
}

// Display returns the display options (color scheme + ViewCube).
func (o Options) Display() (wire.DisplayOptionsView, error) {
	return optionGroupField(o, wire.OptionGroupDisplay, func(r wire.OptionGroupView) *wire.DisplayOptionsView { return r.Display })
}

// SetDisplay writes the display options.
//
// mcp:tool options_set_group
// mcp:summary Writes the display options.
func (o Options) SetDisplay(v wire.DisplayOptionsView) (wire.OKResult, error) {
	args := wire.OptionGroupView{Group: wire.OptionGroupDisplay, Display: &v}
	return call[wire.OKResult](o.c, wire.MethodOptionsSetGroup, args)
}

// Sketch returns the sketch options (grid + snapping).
func (o Options) Sketch() (wire.SketchOptionsView, error) {
	return optionGroupField(o, wire.OptionGroupSketch, func(r wire.OptionGroupView) *wire.SketchOptionsView { return r.Sketch })
}

// SetSketch writes the sketch options.
//
// mcp:tool options_set_group
// mcp:summary Writes the display options.
func (o Options) SetSketch(v wire.SketchOptionsView) (wire.OKResult, error) {
	args := wire.OptionGroupView{Group: wire.OptionGroupSketch, Sketch: &v}
	return call[wire.OKResult](o.c, wire.MethodOptionsSetGroup, args)
}

// Part returns the part-modeling defaults.
func (o Options) Part() (wire.PartOptionsView, error) {
	return optionGroupField(o, wire.OptionGroupPart, func(r wire.OptionGroupView) *wire.PartOptionsView { return r.Part })
}

// SetPart writes the part-modeling defaults.
//
// mcp:tool options_set_group
// mcp:summary Writes the display options.
func (o Options) SetPart(v wire.PartOptionsView) (wire.OKResult, error) {
	args := wire.OptionGroupView{Group: wire.OptionGroupPart, Part: &v}
	return call[wire.OKResult](o.c, wire.MethodOptionsSetGroup, args)
}

// Save returns the save policy (thumbnail capture, dependents, old-version
// retention) (M03-F09).
func (o Options) Save() (wire.SaveOptionsView, error) {
	return optionGroupField(o, wire.OptionGroupSave, func(r wire.OptionGroupView) *wire.SaveOptionsView { return r.Save })
}

// SetSave writes the save policy; the host rejects capture modes it cannot
// perform rather than persisting a dead setting.
//
// mcp:tool options_set_group
// mcp:summary Writes the display options.
func (o Options) SetSave(v wire.SaveOptionsView) (wire.OKResult, error) {
	args := wire.OptionGroupView{Group: wire.OptionGroupSave, Save: &v}
	return call[wire.OKResult](o.c, wire.MethodOptionsSetGroup, args)
}
