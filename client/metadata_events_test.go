// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// TestDispatchObjectRenamed checks the dispatcher decodes an object.renamed event to the typed
// OnObjectRenamed callback with the object kind, key, and names intact (#1644).
func TestDispatchObjectRenamed(t *testing.T) {
	var got wire.ObjectRenamedEvent
	d := NewEventDispatcher()
	d.OnObjectRenamed(func(e wire.ObjectRenamedEvent) { got = e })
	if !d.Dispatch([]byte(`{"type":"object.renamed","document":3,"kind":"body","key":"b7","oldName":"Body1","newName":"Housing"}`)) {
		t.Fatal("Dispatch did not recognise object.renamed")
	}
	if got.Kind != types.ObjectKindBody || got.Key != "b7" || got.OldName != "Body1" || got.NewName != "Housing" {
		t.Errorf("decoded %+v, want body b7 Body1->Housing", got)
	}
}

// TestDispatchPropertyChanged checks the dispatcher decodes a property.changed event to the typed
// OnPropertyChanged callback with the property name and old/new values intact (#1644).
func TestDispatchPropertyChanged(t *testing.T) {
	var got wire.PropertyChangedEvent
	d := NewEventDispatcher()
	d.OnPropertyChanged(func(e wire.PropertyChangedEvent) { got = e })
	if !d.Dispatch([]byte(`{"type":"property.changed","document":3,"kind":"feature","key":"f2","property":"suppressed","oldValue":"false","newValue":"true"}`)) {
		t.Fatal("Dispatch did not recognise property.changed")
	}
	if got.Property != "suppressed" || got.OldValue != "false" || got.NewValue != "true" {
		t.Errorf("decoded %+v, want suppressed false->true", got)
	}
}
