// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"

	"oblikovati.org/api/wire"
)

// EventDispatcher routes the JSON an add-in's Notify entry point receives to
// typed callbacks, so subscribers program against the wire DTOs instead of
// hand-decoding event JSON (M04-F05, Oblikovati/Oblikovati#613). Register
// callbacks with the On* helpers, then feed every incoming event into Dispatch:
//
//	events := client.NewEventDispatcher()
//	events.OnTransaction(func(e wire.TransactionEventPayload) { cache.Invalidate(e.Document) })
//	// in the add-in's Notify entry point:
//	events.Dispatch(eventJSON)
//
// It is not safe for concurrent registration after dispatching begins; wire the
// callbacks up during add-in activation.
type EventDispatcher struct {
	onTransaction     []func(wire.TransactionEventPayload)
	onFileResolution  []func(wire.FileResolutionEventPayload)
	onFileDirty       []func(wire.FileDirtyEventPayload)
	onFileDialogHook  []func(wire.FileDialogHookPayload)
	onOccurrence      []func(wire.OccurrenceEventPayload)
	onAssemblyFeature []func(wire.AssemblyFeaturesChangedEvent)
	onFeature         []func(wire.FeatureLifecycleEvent)
	onSketchEdit      []func(wire.SketchEditEvent)
}

// NewEventDispatcher returns a dispatcher with no callbacks registered.
func NewEventDispatcher() *EventDispatcher { return &EventDispatcher{} }

// OnTransaction subscribes to the five transaction lifecycle events
// (transaction.committed/.undone/.redone/.aborted/.deleted); the payload's Type
// says which fired.
func (d *EventDispatcher) OnTransaction(fn func(wire.TransactionEventPayload)) {
	d.onTransaction = append(d.onTransaction, fn)
}

// OnFileResolution subscribes to file.resolution: a referenced document name
// failed to resolve (ResolvedName carries the substitute, if any).
func (d *EventDispatcher) OnFileResolution(fn func(wire.FileResolutionEventPayload)) {
	d.onFileResolution = append(d.onFileResolution, fn)
}

// OnFileDirty subscribes to file.dirty: a document's clean→dirty transition.
func (d *EventDispatcher) OnFileDirty(fn func(wire.FileDirtyEventPayload)) {
	d.onFileDirty = append(d.onFileDirty, fn)
}

// OnFileDialogHook subscribes to the file-UI hook events (file.new, the
// new/open/save-as dialog hooks, file.openFromMRU, file.populateMetadata); the
// payload's Type says which fired.
func (d *EventDispatcher) OnFileDialogHook(fn func(wire.FileDialogHookPayload)) {
	d.onFileDialogHook = append(d.onFileDialogHook, fn)
}

// OnOccurrence subscribes to the five assembly occurrence-lifecycle events
// (occurrence.added/.deleted/.replaced/.transformed/.suppressed); the payload's
// Type says which fired and carries the affected occurrence's identity (M11-F07).
func (d *EventDispatcher) OnOccurrence(fn func(wire.OccurrenceEventPayload)) {
	d.onOccurrence = append(d.onOccurrence, fn)
}

// OnAssemblyFeatures subscribes to assemblyFeatures.changed: the active assembly's
// machining-feature program was re-evaluated (the payload carries each feature's
// resulting health; re-read detail with assemblyFeatures.list) (M11-F08).
func (d *EventDispatcher) OnAssemblyFeatures(fn func(wire.AssemblyFeaturesChangedEvent)) {
	d.onAssemblyFeature = append(d.onAssemblyFeature, fn)
}

// OnFeature subscribes to the three feature-lifecycle events (feature.added/.edited/.deleted);
// the payload's Type says which fired and carries the affected feature's identity (#148).
func (d *EventDispatcher) OnFeature(fn func(wire.FeatureLifecycleEvent)) {
	d.onFeature = append(d.onFeature, fn)
}

// OnSketchEdit subscribes to the sketch-edit-mode events (sketch.editEntered/.editExited); the
// payload's Type says which fired and carries the sketch's identity (#148).
func (d *EventDispatcher) OnSketchEdit(fn func(wire.SketchEditEvent)) {
	d.onSketchEdit = append(d.onSketchEdit, fn)
}

// transactionEventTypes are the type tags decoded as [wire.TransactionEventPayload].
var transactionEventTypes = map[string]bool{
	wire.EventTransactionCommitted: true,
	wire.EventTransactionUndone:    true,
	wire.EventTransactionRedone:    true,
	wire.EventTransactionAborted:   true,
	wire.EventTransactionDeleted:   true,
}

// fileDialogHookEventTypes are the type tags decoded as [wire.FileDialogHookPayload].
var fileDialogHookEventTypes = map[string]bool{
	wire.EventFileNew:              true,
	wire.EventFileNewDialog:        true,
	wire.EventFileOpenDialog:       true,
	wire.EventFileSaveAsDialog:     true,
	wire.EventFileOpenFromMRU:      true,
	wire.EventFilePopulateMetadata: true,
}

// occurrenceEventTypes are the type tags decoded as [wire.OccurrenceEventPayload].
var occurrenceEventTypes = map[string]bool{
	wire.EventOccurrenceAdded:       true,
	wire.EventOccurrenceDeleted:     true,
	wire.EventOccurrenceReplaced:    true,
	wire.EventOccurrenceTransformed: true,
	wire.EventOccurrenceSuppressed:  true,
}

// featureEventTypes are the type tags decoded as [wire.FeatureLifecycleEvent].
var featureEventTypes = map[string]bool{
	wire.EventFeatureAdded:   true,
	wire.EventFeatureEdited:  true,
	wire.EventFeatureDeleted: true,
}

// sketchEditEventTypes are the type tags decoded as [wire.SketchEditEvent].
var sketchEditEventTypes = map[string]bool{
	wire.EventSketchEditEntered: true,
	wire.EventSketchEditExited:  true,
}

// Dispatch decodes one Notify event and fires the matching callbacks, reporting
// whether the event's type tag is one this dispatcher understands (false lets
// the caller route other event families elsewhere). Malformed JSON is false.
func (d *EventDispatcher) Dispatch(eventJSON []byte) bool {
	var tag struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(eventJSON, &tag); err != nil {
		return false
	}
	return d.dispatchByType(tag.Type, eventJSON)
}

// dispatchByType routes a decoded type tag to its payload family.
func (d *EventDispatcher) dispatchByType(eventType string, eventJSON []byte) bool {
	switch {
	case transactionEventTypes[eventType]:
		return fireDecoded(eventJSON, d.onTransaction)
	case eventType == wire.EventFileResolution:
		return fireDecoded(eventJSON, d.onFileResolution)
	case eventType == wire.EventFileDirty:
		return fireDecoded(eventJSON, d.onFileDirty)
	case fileDialogHookEventTypes[eventType]:
		return fireDecoded(eventJSON, d.onFileDialogHook)
	case occurrenceEventTypes[eventType]:
		return fireDecoded(eventJSON, d.onOccurrence)
	case eventType == wire.EventAssemblyFeaturesChanged:
		return fireDecoded(eventJSON, d.onAssemblyFeature)
	case featureEventTypes[eventType]:
		return fireDecoded(eventJSON, d.onFeature)
	case sketchEditEventTypes[eventType]:
		return fireDecoded(eventJSON, d.onSketchEdit)
	}
	return false
}

// fireDecoded unmarshals the payload once and hands it to every callback.
func fireDecoded[P wire.TransactionEventPayload | wire.FileResolutionEventPayload |
	wire.FileDirtyEventPayload | wire.FileDialogHookPayload |
	wire.OccurrenceEventPayload | wire.AssemblyFeaturesChangedEvent |
	wire.FeatureLifecycleEvent | wire.SketchEditEvent](eventJSON []byte, fns []func(P)) bool {
	var payload P
	if err := json.Unmarshal(eventJSON, &payload); err != nil {
		return false
	}
	for _, fn := range fns {
		fn(payload)
	}
	return true
}
