// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// EntityFormat reads one sketch entity's formatting overrides — its line type, colour and stroke
// width (Oblikovati/Oblikovati#2015). The result's HasFormat is false when the entity inherits
// everything from the sketch.
//
// mcp:tool sketch_get_entity_format
// mcp:summary Reads one sketch entity's line type, colour and stroke width overrides.
func (s Sketch) EntityFormat(entityID uint64) (wire.SketchEntityFormatView, error) {
	return call[wire.SketchEntityFormatView](s.c, wire.MethodSketchGetEntityFormat,
		wire.SketchEntityFormatArgs{EntityID: entityID})
}

// SetEntityFormat sets one sketch entity's formatting overrides. A format that overrides nothing
// clears them, so writing back a default is how an entity is returned to the sketch's attributes.
//
// mcp:tool sketch_set_entity_format
// mcp:summary Sets one sketch entity's line type, colour and stroke width overrides.
func (s Sketch) SetEntityFormat(args wire.SetSketchEntityFormatArgs) (wire.SketchEntityFormatView, error) {
	return call[wire.SketchEntityFormatView](s.c, wire.MethodSketchSetEntityFormat, args)
}

// FormatModes reads the Format panel's armed creation modes — what newly drawn geometry becomes.
//
// mcp:tool sketch_get_format_modes
// mcp:summary Reads the Format panel's armed creation modes (construction, centerline, centre point, driven dimension, show format).
func (s Sketch) FormatModes() (wire.SketchFormatModesView, error) {
	return call[wire.SketchFormatModesView](s.c, wire.MethodSketchGetFormatModes, nil)
}

// SetFormatModes replaces the Format panel's armed creation modes and returns the stored value.
// Every field is replaced, so read with [Sketch.FormatModes] first to change just one.
//
// mcp:tool sketch_set_format_modes
// mcp:summary Replaces the Format panel's armed creation modes.
func (s Sketch) SetFormatModes(view wire.SketchFormatModesView) (wire.SketchFormatModesView, error) {
	return call[wire.SketchFormatModesView](s.c, wire.MethodSketchSetFormatModes, view)
}
