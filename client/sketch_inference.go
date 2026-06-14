// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// SetInferenceOptions configures sketch inference: whether point snapping and
// constraint auto-application run, and which constraint family wins when two
// could apply (M06-F10, Oblikovati/Oblikovati#625).
//
// mcp:tool sketch_set_inference_options
// mcp:summary Configures sketch inference: whether point snapping and constraint auto-application run, and which constraint family wins when two could apply (M06-F10, Oblikovati/Oblikovati#625).
func (s Sketch) SetInferenceOptions(view wire.InferenceOptionsView) (wire.InferenceOptionsView, error) {
	var r wire.InferenceOptionsView
	return r, s.c.call(wire.MethodSketchSetInferenceOptions, view, &r)
}

// InferenceOptions reads the current sketch inference configuration.
//
// mcp:tool sketch_get_inference_options
// mcp:summary Reads the current sketch inference configuration.
func (s Sketch) InferenceOptions() (wire.InferenceOptionsView, error) {
	var r wire.InferenceOptionsView
	return r, s.c.call(wire.MethodSketchGetInferenceOptions, nil, &r)
}
