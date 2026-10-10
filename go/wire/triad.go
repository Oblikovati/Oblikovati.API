// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The interactive-gizmo surface of M05-F13 (#620): the move/rotate triad and
// add-in manipulator handles. Both stream drags as push events carrying a
// DragContext; a wire call never blocks on user input.

// TriadSpec places the triad: a position and an orientation (three column axes;
// identity when omitted), with the allowed segments (empty ⇒ all). Command, when
// set, ties the triad's lifetime to that command like interaction graphics.
type TriadSpec struct {
	Position types.Point          `json:"position"`
	AxisX    *types.UnitVector    `json:"axisX,omitempty"`
	AxisY    *types.UnitVector    `json:"axisY,omitempty"`
	AxisZ    *types.UnitVector    `json:"axisZ,omitempty"`
	Allowed  []types.TriadSegment `json:"allowed,omitempty"`
	Visible  bool                 `json:"visible"`
	Command  string               `json:"command,omitempty"`
}

// ShowTriadArgs is the request of [MethodTriadShow] (and [MethodTriadUpdate]).
type ShowTriadArgs struct {
	Triad TriadSpec `json:"triad"`
}

// DragContext rides every drag event: where the gesture started, the current
// pointer ray, the modifiers, and any point inference the position snapped to.
type DragContext struct {
	Start     types.Point              `json:"start"`
	RayOrigin types.Point              `json:"rayOrigin"`
	RayDir    types.Vector             `json:"rayDir"`
	Shift     bool                     `json:"shift,omitempty"`
	Ctrl      bool                     `json:"ctrl,omitempty"`
	Inference types.PointInferenceKind `json:"inference,omitempty"`
}

// TriadDragEvent is the push event (type [EventTriadDrag]) streaming a triad
// gesture: Phase is "start", "move" or "end"; Delta is the row-major 4×4 transform
// accumulated since the drag started (identity at start).
type TriadDragEvent struct {
	Type     string              `json:"type"` // always EventTriadDrag
	Phase    string              `json:"phase"`
	Segment  types.TriadSegment  `json:"segment"`
	MoveType types.TriadMoveType `json:"moveType"`
	Delta    types.Matrix        `json:"delta"`
	Context  DragContext         `json:"context"`
}

// TriadSegmentEvent is the push event (type [EventTriadSegment]) fired when the
// hovered/selected segment changes (OnSegmentSelectionChange).
type TriadSegmentEvent struct {
	Type    string             `json:"type"` // always EventTriadSegment
	Segment types.TriadSegment `json:"segment"`
	Hovered bool               `json:"hovered"`
}

// ManipulatorHandleSpec is one add-in drag handle: a world-space hotspot of
// RadiusPx pixels, typically placed over the add-in's client graphics — the
// custom-gizmo building block (ManipulatorEvents).
type ManipulatorHandleSpec struct {
	ID       string      `json:"id"`
	Position types.Point `json:"position"`
	RadiusPx float64     `json:"radiusPx,omitempty"`
}

// SetManipulatorsArgs is the request of [MethodManipulatorsSet]: replace one
// add-in gizmo's handle set (the declared-bulk precedent).
type SetManipulatorsArgs struct {
	ID      string                  `json:"id"`
	Handles []ManipulatorHandleSpec `json:"handles"`
	Command string                  `json:"command,omitempty"`
}

// RemoveManipulatorsArgs is the request of [MethodManipulatorsRemove].
type RemoveManipulatorsArgs struct {
	ID string `json:"id"`
}

// ManipulatorDragEvent is the push event (type [EventManipulatorDrag]) streaming a
// handle gesture: Phase is "start", "move" or "end"; Position is the handle's
// dragged world position (on the view plane through its start, snapped per the
// context's inference).
type ManipulatorDragEvent struct {
	Type     string      `json:"type"` // always EventManipulatorDrag
	Gizmo    string      `json:"gizmo"`
	Handle   string      `json:"handle"`
	Phase    string      `json:"phase"`
	Position types.Point `json:"position"`
	Context  DragContext `json:"context"`
}
