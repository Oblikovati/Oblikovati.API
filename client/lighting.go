// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Lighting is the scene-lighting operation group: it lets an add-in read and switch the
// active lighting style and read/edit the discrete lights, so
// automations can drive how the model is illuminated.
type Lighting struct{ c *Client }

// Lighting returns the scene-lighting operation group.
func (c *Client) Lighting() Lighting { return Lighting{c} }

// Style returns the active lighting style's global controls.
//
//	s, _ := client.Lighting().Style()
//	_ = s.Exposure
//
// mcp:tool get_lighting_style
// mcp:summary Read the active lighting style (the named set of lights + exposure).
func (l Lighting) Style() (wire.LightingStyleView, error) {
	var r wire.LightingStyleView
	return r, l.c.call(wire.MethodLightingGetStyle, nil, &r)
}

// SetStyle activates the named lighting style, returning the resulting style.
//
//	client.Lighting().SetStyle("Outdoors")
//
// mcp:tool set_lighting_style
// mcp:summary Switch to a lighting style by id; see list_lighting_styles.
func (l Lighting) SetStyle(name string) (wire.LightingStyleView, error) {
	var r wire.LightingStyleView
	return r, l.c.call(wire.MethodLightingSetStyle, wire.SetLightingStyleArgs{Name: name}, &r)
}

// ListStyles returns every selectable lighting style, flagging the active one.
//
// mcp:tool list_lighting_styles
// mcp:summary List the available lighting styles (the ids set_lighting_style accepts).
func (l Lighting) ListStyles() (wire.LightingStyleListResult, error) {
	var r wire.LightingStyleListResult
	return r, l.c.call(wire.MethodLightingListStyles, nil, &r)
}

// Lights returns the active style's discrete lights.
//
// mcp:tool list_lights
// mcp:summary List the lights in the active lighting style (id, type, direction/position, intensity, color).
func (l Lighting) Lights() (wire.LightListResult, error) {
	var r wire.LightListResult
	return r, l.c.call(wire.MethodLightingListLights, nil, &r)
}

// AddLight adds a light of the given emission shape with neutral defaults, returning it.
//
// mcp:tool add_light
// mcp:summary Add a light to the active lighting style (type directional|point|spot, with direction/position, color, intensity).
func (l Lighting) AddLight(def types.LightDefinitionTypeEnum) (wire.LightInfo, error) {
	var r wire.LightInfo
	return r, l.c.call(wire.MethodLightingAddLight, wire.AddLightArgs{DefinitionType: def}, &r)
}

// SetLight replaces the state of the light at index, returning the resulting light.
//
// mcp:tool set_light
// mcp:summary Update a light's properties by id (direction/position, color, intensity, on/off).
func (l Lighting) SetLight(index int, light wire.LightInfo) (wire.LightInfo, error) {
	var r wire.LightInfo
	return r, l.c.call(wire.MethodLightingSetLight, wire.SetLightArgs{Index: index, Light: light}, &r)
}

// Environment is the image-based-lighting operation group: it lets an add-in read, switch,
// and load the HDR environment the scene reflects.
type Environment struct{ c *Client }

// Environment returns the image-based-lighting operation group.
func (c *Client) Environment() Environment { return Environment{c} }

// Get returns the active environment.
//
// mcp:tool get_environment
// mcp:summary Read the HDR/IBL environment (sky preset or loaded image, rotation, intensity, background).
func (e Environment) Get() (wire.EnvironmentView, error) {
	var r wire.EnvironmentView
	return r, e.c.call(wire.MethodEnvironmentGet, nil, &r)
}

// Set activates the named built-in preset with the given display parameters.
//
//	client.Environment().Set(wire.SetEnvironmentArgs{Preset: "Studio", Intensity: 1, ShowImage: true})
//
// mcp:tool set_environment
// mcp:summary Set the environment: choose a preset or tune rotation/intensity/background visibility.
func (e Environment) Set(args wire.SetEnvironmentArgs) (wire.EnvironmentView, error) {
	var r wire.EnvironmentView
	return r, e.c.call(wire.MethodEnvironmentSet, args, &r)
}

// ListPresets returns every built-in environment, flagging the active one.
//
// mcp:tool list_environment_presets
// mcp:summary List the built-in environment (sky/IBL) presets.
func (e Environment) ListPresets() (wire.EnvironmentPresetListResult, error) {
	var r wire.EnvironmentPresetListResult
	return r, e.c.call(wire.MethodEnvironmentListPresets, nil, &r)
}

// LoadImage loads an equirectangular HDR file (.hdr) as the environment.
//
// mcp:tool load_environment_image
// mcp:summary Load an HDR/EXR image as the lighting environment (image-based lighting) by reference.
func (e Environment) LoadImage(filePath string) (wire.EnvironmentView, error) {
	var r wire.EnvironmentView
	return r, e.c.call(wire.MethodEnvironmentLoadImage, wire.LoadEnvironmentImageArgs{FilePath: filePath}, &r)
}

// Shadows returns the viewport's current shadow settings.
//
// mcp:tool get_shadows
// mcp:summary Read the viewport's shadow settings (ground/object shadows, direction, softness).
func (v View) Shadows() (wire.ShadowSettings, error) {
	var r wire.ShadowSettings
	return r, v.c.call(wire.MethodViewGetShadows, nil, &r)
}

// SetShadows applies the viewport's shadow settings, returning the resulting settings.
//
// mcp:tool set_shadows
// mcp:summary Apply the viewport's shadow settings and return the resulting settings.
func (v View) SetShadows(s wire.ShadowSettings) (wire.ShadowSettings, error) {
	var r wire.ShadowSettings
	return r, v.c.call(wire.MethodViewSetShadows, s, &r)
}
