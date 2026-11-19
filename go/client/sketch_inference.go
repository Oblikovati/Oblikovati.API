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
	return call[wire.InferenceOptionsView](s.c, wire.MethodSketchSetInferenceOptions, view)
}

// InferenceOptions reads the current sketch inference configuration.
//
// mcp:tool sketch_get_inference_options
// mcp:summary Reads the current sketch inference configuration.
func (s Sketch) InferenceOptions() (wire.InferenceOptionsView, error) {
	return call[wire.InferenceOptionsView](s.c, wire.MethodSketchGetInferenceOptions, nil)
}

// InferenceScope reads the applied scope for a 2D sketch in the current document session.
// Example: `scope, err := c.Sketch().InferenceScope(0)`.
//
// mcp:tool sketch_get_inference_scope
// mcp:summary Reads the candidate geometry scope for a 2D sketch in the current document session.
func (s Sketch) InferenceScope(sketchIndex int) (wire.SketchInferenceScopeView, error) {
	return call[wire.SketchInferenceScopeView](s.c, wire.MethodSketchGetInferenceScope, wire.SketchInferenceScopeArgs{SketchIndex: sketchIndex})
}

// SetInferenceScope atomically replaces a 2D sketch's session inference scope.
// Example: `scope, err := c.Sketch().SetInferenceScope(view)`.
//
// mcp:tool sketch_set_inference_scope
// mcp:summary Replaces the candidate geometry scope for a 2D sketch in the current document session.
func (s Sketch) SetInferenceScope(view wire.SketchInferenceScopeView) (wire.SketchInferenceScopeView, error) {
	return call[wire.SketchInferenceScopeView](s.c, wire.MethodSketchSetInferenceScope, view)
}
