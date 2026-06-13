// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// TestEventDispatcherRoutesTransactionEvents drives every transaction type tag
// through Dispatch and checks the typed payload reaches the callback.
func TestEventDispatcherRoutesTransactionEvents(t *testing.T) {
	d := NewEventDispatcher()
	var got []wire.TransactionEventPayload
	d.OnTransaction(func(e wire.TransactionEventPayload) { got = append(got, e) })

	events := [][]byte{
		[]byte(`{"type":"transaction.committed","document":3,"label":"Extrude","point":3588}`),
		[]byte(`{"type":"transaction.undone","document":3,"label":"Extrude","point":3587}`),
		[]byte(`{"type":"transaction.redone","document":3,"label":"Extrude","point":3586}`),
		[]byte(`{"type":"transaction.aborted","document":3,"label":"batch"}`),
		[]byte(`{"type":"transaction.deleted","document":3,"point":3585}`),
	}
	for _, ev := range events {
		if !d.Dispatch(ev) {
			t.Errorf("Dispatch(%s) = false, want it routed", ev)
		}
	}
	if len(got) != len(events) {
		t.Fatalf("callback fired %d times, want %d", len(got), len(events))
	}
	if got[0].Label != "Extrude" || got[0].Point != types.TransactionPointCurrent || got[0].Document != 3 {
		t.Errorf("committed payload = %+v, want Extrude at the current point on document 3", got[0])
	}
	if got[1].Point != types.TransactionPointPrevious || got[2].Point != types.TransactionPointNext {
		t.Errorf("undo/redo points = %v/%v, want previous/next", got[1].Point, got[2].Point)
	}
}

// TestEventDispatcherRoutesFileEvents covers the file-access and file-UI hook
// families, each to its own typed callback.
func TestEventDispatcherRoutesFileEvents(t *testing.T) {
	d := NewEventDispatcher()
	var resolution wire.FileResolutionEventPayload
	var dirty wire.FileDirtyEventPayload
	var hooks []wire.FileDialogHookPayload
	d.OnFileResolution(func(e wire.FileResolutionEventPayload) { resolution = e })
	d.OnFileDirty(func(e wire.FileDirtyEventPayload) { dirty = e })
	d.OnFileDialogHook(func(e wire.FileDialogHookPayload) { hooks = append(hooks, e) })

	if !d.Dispatch([]byte(`{"type":"file.resolution","requestedName":"base.obk","resolvedName":"/lib/base.obk"}`)) {
		t.Error("file.resolution must route")
	}
	if resolution.RequestedName != "base.obk" || resolution.ResolvedName != "/lib/base.obk" {
		t.Errorf("resolution payload = %+v, want base.obk resolved to /lib/base.obk", resolution)
	}

	if !d.Dispatch([]byte(`{"type":"file.dirty","document":7,"fullDocumentName":"Part1"}`)) {
		t.Error("file.dirty must route")
	}
	if dirty.Document != 7 || dirty.FullDocumentName != "Part1" {
		t.Errorf("dirty payload = %+v, want document 7 Part1", dirty)
	}

	hookEvents := [][]byte{
		[]byte(`{"type":"file.new","documentType":1}`),
		[]byte(`{"type":"file.openDialog","fileName":"/models/bracket.obk"}`),
		[]byte(`{"type":"file.saveAsDialog","saveCopyAs":true}`),
		[]byte(`{"type":"file.openFromMRU","fileName":"/models/gear.obk"}`),
		[]byte(`{"type":"file.populateMetadata","metadata":[{"name":"author","value":"vm"}]}`),
		[]byte(`{"type":"file.newDialog","templateFile":"/templates/part.obk"}`),
	}
	for _, ev := range hookEvents {
		if !d.Dispatch(ev) {
			t.Errorf("Dispatch(%s) = false, want it routed", ev)
		}
	}
	if len(hooks) != len(hookEvents) {
		t.Fatalf("hook callback fired %d times, want %d", len(hooks), len(hookEvents))
	}
	if hooks[0].DocumentType != types.DocumentPart || hooks[1].FileName != "/models/bracket.obk" {
		t.Errorf("hook payloads = %+v / %+v, want part new + open path decoded", hooks[0], hooks[1])
	}
	if !hooks[2].SaveCopyAs || hooks[4].Metadata[0].Name != "author" {
		t.Errorf("hook payloads = %+v / %+v, want saveCopyAs + metadata decoded", hooks[2], hooks[4])
	}
}

// TestEventDispatcherRoutesOccurrenceEvents drives every occurrence type tag through
// Dispatch and checks the typed payload — identity, suppression, and the new/prior
// placement on a transformed event — reaches the callback.
func TestEventDispatcherRoutesOccurrenceEvents(t *testing.T) {
	d := NewEventDispatcher()
	var got []wire.OccurrenceEventPayload
	d.OnOccurrence(func(e wire.OccurrenceEventPayload) { got = append(got, e) })

	events := [][]byte{
		[]byte(`{"type":"occurrence.added","document":4,"occurrence":7,"name":"bracket:1"}`),
		[]byte(`{"type":"occurrence.deleted","document":4,"occurrence":7,"name":"bracket:1"}`),
		[]byte(`{"type":"occurrence.replaced","document":4,"occurrence":8,"name":"pin:2"}`),
		[]byte(`{"type":"occurrence.suppressed","document":4,"occurrence":8,"name":"pin:2","suppressed":true}`),
		[]byte(`{"type":"occurrence.transformed","document":4,"occurrence":8,"name":"pin:2",` +
			`"transform":[1,0,0,5,0,1,0,0,0,0,1,0,0,0,0,1],` +
			`"previous":[1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1]}`),
	}
	for _, ev := range events {
		if !d.Dispatch(ev) {
			t.Errorf("Dispatch(%s) = false, want it routed", ev)
		}
	}
	if len(got) != len(events) {
		t.Fatalf("callback fired %d times, want %d", len(got), len(events))
	}
	if got[0].Occurrence != 7 || got[0].Name != "bracket:1" || got[0].Document != 4 {
		t.Errorf("added payload = %+v, want occurrence 7 bracket:1 on document 4", got[0])
	}
	if !got[3].Suppressed {
		t.Errorf("suppressed payload = %+v, want Suppressed=true", got[3])
	}
	move := got[4]
	if move.Transform == nil || move.Previous == nil {
		t.Fatalf("transformed payload = %+v, want both placements decoded", move)
	}
	if move.Transform.Cells[3] != 5 || move.Previous.Cells[3] != 0 {
		t.Errorf("transformed placements = %v / %v, want new X=5 and prior X=0", move.Transform.Cells[3], move.Previous.Cells[3])
	}
}

// TestEventDispatcherIgnoresForeignAndMalformedEvents pins the false returns:
// other event families and broken JSON are left to the caller.
func TestEventDispatcherIgnoresForeignAndMalformedEvents(t *testing.T) {
	d := NewEventDispatcher()
	fired := false
	d.OnTransaction(func(wire.TransactionEventPayload) { fired = true })

	if d.Dispatch([]byte(`{"type":"edit.committed","document":1}`)) {
		t.Error("a foreign event family must not be claimed")
	}
	if d.Dispatch([]byte(`{not json`)) {
		t.Error("malformed JSON must not be claimed")
	}
	if fired {
		t.Error("no callback may fire for unrouted events")
	}
}
