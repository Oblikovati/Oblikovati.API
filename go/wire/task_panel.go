// SPDX-License-Identifier: Apache-2.0

package wire

// TaskPanelSpec is a modal task panel that takes over the host task area (FreeCAD Task-dialog
// semantics): declarative PanelControlSpec content with OK/Cancel. Showing is asynchronous — like
// file dialogs it must never block the session goroutine — so the user's accept/cancel arrives as a
// [TaskPanelClosedEvent]. While open, edits to its controls push the same PanelValueChangedEvent /
// PanelReferencesChangedEvent keyed on this panel's ID. Unlike WebDialogSpec it is built from the
// shared control vocabulary, so a referenceList composes inside it.
type TaskPanelSpec struct {
	ID          string             `json:"id"`
	Title       string             `json:"title"`
	Controls    []PanelControlSpec `json:"controls,omitempty"`
	OKLabel     string             `json:"okLabel,omitempty"`     // default "OK"
	CancelLabel string             `json:"cancelLabel,omitempty"` // default "Cancel"
}

// ShowTaskPanelArgs is the request of [MethodTaskPanelShow].
type ShowTaskPanelArgs struct {
	Panel TaskPanelSpec `json:"panel"`
}

// CloseTaskPanelArgs is the request of [MethodTaskPanelClose] (programmatic dismissal).
type CloseTaskPanelArgs struct {
	ID string `json:"id"`
}

// TaskPanelClosedEvent is the push event (type [EventTaskPanelClosed]) delivering the user's
// accept/cancel. Accepted is true for OK, false for Cancel/close. Control values are NOT echoed —
// they already arrived incrementally via the value/references events while the panel was open.
type TaskPanelClosedEvent struct {
	Type     string `json:"type"` // always EventTaskPanelClosed
	ID       string `json:"id"`
	Accepted bool   `json:"accepted"`
}
