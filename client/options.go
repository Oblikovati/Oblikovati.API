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

// General returns the general options (startup behavior).
func (o Options) General() (wire.GeneralOptionsView, error) {
	r, err := o.getGroup(wire.OptionGroupGeneral)
	if err != nil {
		return wire.GeneralOptionsView{}, err
	}
	if r.General == nil {
		return wire.GeneralOptionsView{}, fmt.Errorf("client: options.getGroup(%q) reply carries no general payload", wire.OptionGroupGeneral)
	}
	return *r.General, nil
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
	r, err := o.getGroup(wire.OptionGroupDisplay)
	if err != nil {
		return wire.DisplayOptionsView{}, err
	}
	if r.Display == nil {
		return wire.DisplayOptionsView{}, fmt.Errorf("client: options.getGroup(%q) reply carries no display payload", wire.OptionGroupDisplay)
	}
	return *r.Display, nil
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
	r, err := o.getGroup(wire.OptionGroupSketch)
	if err != nil {
		return wire.SketchOptionsView{}, err
	}
	if r.Sketch == nil {
		return wire.SketchOptionsView{}, fmt.Errorf("client: options.getGroup(%q) reply carries no sketch payload", wire.OptionGroupSketch)
	}
	return *r.Sketch, nil
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
	r, err := o.getGroup(wire.OptionGroupPart)
	if err != nil {
		return wire.PartOptionsView{}, err
	}
	if r.Part == nil {
		return wire.PartOptionsView{}, fmt.Errorf("client: options.getGroup(%q) reply carries no part payload", wire.OptionGroupPart)
	}
	return *r.Part, nil
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
	r, err := o.getGroup(wire.OptionGroupSave)
	if err != nil {
		return wire.SaveOptionsView{}, err
	}
	if r.Save == nil {
		return wire.SaveOptionsView{}, fmt.Errorf("client: options.getGroup(%q) reply carries no save payload", wire.OptionGroupSave)
	}
	return *r.Save, nil
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
