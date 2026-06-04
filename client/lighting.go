// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/Oblikovati/api/types"
	"github.com/Oblikovati/api/wire"
)

// Lighting is the scene-lighting operation group: it lets an add-in read and switch the
// active lighting style (Inventor's LightingStyle) and read/edit the discrete lights, so
// automations can drive how the model is illuminated.
type Lighting struct{ c *Client }

// Lighting returns the scene-lighting operation group.
func (c *Client) Lighting() Lighting { return Lighting{c} }

// Style returns the active lighting style's global controls.
//
//	s, _ := client.Lighting().Style()
//	_ = s.Exposure
func (l Lighting) Style() (wire.LightingStyleView, error) {
	var r wire.LightingStyleView
	return r, l.c.call(wire.MethodLightingGetStyle, nil, &r)
}

// SetStyle activates the named lighting style, returning the resulting style.
//
//	client.Lighting().SetStyle("Outdoors")
func (l Lighting) SetStyle(name string) (wire.LightingStyleView, error) {
	var r wire.LightingStyleView
	return r, l.c.call(wire.MethodLightingSetStyle, wire.SetLightingStyleArgs{Name: name}, &r)
}

// ListStyles returns every selectable lighting style, flagging the active one.
func (l Lighting) ListStyles() (wire.LightingStyleListResult, error) {
	var r wire.LightingStyleListResult
	return r, l.c.call(wire.MethodLightingListStyles, nil, &r)
}

// Lights returns the active style's discrete lights.
func (l Lighting) Lights() (wire.LightListResult, error) {
	var r wire.LightListResult
	return r, l.c.call(wire.MethodLightingListLights, nil, &r)
}

// AddLight adds a light of the given emission shape with neutral defaults, returning it.
func (l Lighting) AddLight(def types.LightDefinitionTypeEnum) (wire.LightInfo, error) {
	var r wire.LightInfo
	return r, l.c.call(wire.MethodLightingAddLight, wire.AddLightArgs{DefinitionType: def}, &r)
}

// SetLight replaces the state of the light at index, returning the resulting light.
func (l Lighting) SetLight(index int, light wire.LightInfo) (wire.LightInfo, error) {
	var r wire.LightInfo
	return r, l.c.call(wire.MethodLightingSetLight, wire.SetLightArgs{Index: index, Light: light}, &r)
}

// Environment is the image-based-lighting operation group: it lets an add-in read, switch,
// and load the HDR environment the scene reflects (Inventor's LightingStyle image).
type Environment struct{ c *Client }

// Environment returns the image-based-lighting operation group.
func (c *Client) Environment() Environment { return Environment{c} }

// Get returns the active environment.
func (e Environment) Get() (wire.EnvironmentView, error) {
	var r wire.EnvironmentView
	return r, e.c.call(wire.MethodEnvironmentGet, nil, &r)
}

// Set activates the named built-in preset with the given display parameters.
//
//	client.Environment().Set(wire.SetEnvironmentArgs{Preset: "Studio", Intensity: 1, ShowImage: true})
func (e Environment) Set(args wire.SetEnvironmentArgs) (wire.EnvironmentView, error) {
	var r wire.EnvironmentView
	return r, e.c.call(wire.MethodEnvironmentSet, args, &r)
}

// ListPresets returns every built-in environment, flagging the active one.
func (e Environment) ListPresets() (wire.EnvironmentPresetListResult, error) {
	var r wire.EnvironmentPresetListResult
	return r, e.c.call(wire.MethodEnvironmentListPresets, nil, &r)
}

// LoadImage loads an equirectangular HDR file (.hdr) as the environment.
func (e Environment) LoadImage(filePath string) (wire.EnvironmentView, error) {
	var r wire.EnvironmentView
	return r, e.c.call(wire.MethodEnvironmentLoadImage, wire.LoadEnvironmentImageArgs{FilePath: filePath}, &r)
}

// Shadows returns the viewport's current shadow settings.
func (v View) Shadows() (wire.ShadowSettings, error) {
	var r wire.ShadowSettings
	return r, v.c.call(wire.MethodViewGetShadows, nil, &r)
}

// SetShadows applies the viewport's shadow settings, returning the resulting settings.
func (v View) SetShadows(s wire.ShadowSettings) (wire.ShadowSettings, error) {
	var r wire.ShadowSettings
	return r, v.c.call(wire.MethodViewSetShadows, s, &r)
}
