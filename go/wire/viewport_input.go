// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Synthesised user input — [MethodViewportClick] and [MethodViewportKey].
//
// Every other method here edits the model directly. That cannot reach the behaviour that only
// exists on the way IN: which constraints a tool infers from where you clicked, what a command
// previews between clicks, how a multi-click chain builds up. A client automating or testing the
// application needs to drive those the way a person does.

// ClickViewportArgs is the request of [MethodViewportClick]: where to click, and with what.
//
// Give the position EITHER as viewport pixels (X, Y from the top-left of the 3D viewport) or as
// Point, a model-space position the host projects to pixels for you. Point is usually what a
// caller wants — it needs no knowledge of the camera — and is required when clicking a spot on a
// sketch plane, where a pixel would have to be derived from the current view.
type ClickViewportArgs struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`

	// Point is a model-space position to click, projected to the viewport by the host. It takes
	// precedence over X/Y.
	Point *types.Point `json:"point,omitempty"`

	// Button is "left" (the default), "right", "middle", or one of the extended buttons a mouse may
	// carry: "back", "forward", "button5", "button6", "button7" (#1822). The extended buttons deliver
	// but carry no default binding, so a host tool or add-in reads them without any built-in effect.
	Button string `json:"button,omitempty"`

	// Held modifiers, as the selection and snapping paths read them.
	Shift bool `json:"shift,omitempty"`
	Ctrl  bool `json:"ctrl,omitempty"`
	Alt   bool `json:"alt,omitempty"`
}

// ClickViewportResult is the response of [MethodViewportClick]: the pixel actually clicked (useful
// when Point was projected), and the command still running afterwards — empty once a command has
// finished, which is how a caller knows a shape was created rather than another click awaited.
type ClickViewportResult struct {
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	ActiveTool string  `json:"activeTool"`
}

// ScrollViewportArgs is the request of [MethodViewportScroll]: a mouse-wheel tick at a viewport
// pixel, in notches (a wheel detent is one notch; fractional for high-resolution wheels). DY is the
// vertical axis (positive scrolls up — the zoom-in direction the host has always used). DX is the
// horizontal axis a two-axis wheel or thumbwheel reports (#1822); it is delivered but carries no
// default binding, so it changes nothing until a tool or add-in reads it. X/Y locate the cursor so
// the host can zoom toward it, defaulting to the viewport centre when omitted.
type ScrollViewportArgs struct {
	DX float64 `json:"dx,omitempty"`
	DY float64 `json:"dy,omitempty"`
	X  float64 `json:"x,omitempty"`
	Y  float64 `json:"y,omitempty"`

	Shift bool `json:"shift,omitempty"`
	Ctrl  bool `json:"ctrl,omitempty"`
	Alt   bool `json:"alt,omitempty"`
}

// ScrollViewportResult is the response of [MethodViewportScroll]: the command still running after
// the scroll (empty once none is), mirroring the click/key results.
type ScrollViewportResult struct {
	ActiveTool string `json:"activeTool"`
}

// PressKeyArgs is the request of [MethodViewportKey]: a key to deliver to the running command.
//
// Key is the key's name — "Escape", "Enter", "Delete" — matching what the host's own input layer
// reports. Escape and Enter are the two that matter most: they are how a variable-length command
// (a continuous line chain, a spline) says it is finished.
type PressKeyArgs struct {
	Key   string `json:"key"`
	Shift bool   `json:"shift,omitempty"`
	Ctrl  bool   `json:"ctrl,omitempty"`
	Alt   bool   `json:"alt,omitempty"`
}

// PressKeyResult is the response of [MethodViewportKey]: the command still running after the key,
// empty when the key ended it.
type PressKeyResult struct {
	ActiveTool string `json:"activeTool"`
}
