// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestKeymapListReadsCatalog(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"bindings":[{"actionId":"Feature.Extrude","displayName":"Extrude","kind":"command","chord":"E","defaultChord":"E"}]}`)}
	r, err := New(ft).Keymap().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodKeymapList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodKeymapList)
	}
	if len(r.Bindings) != 1 || r.Bindings[0].ActionID != "Feature.Extrude" || r.Bindings[0].Chord != "E" {
		t.Errorf("bindings = %+v, want one Extrude/E entry", r.Bindings)
	}
}

func TestKeymapSetChordSendsCanonicalString(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	if _, err := New(ft).Keymap().SetChord("Feature.Extrude", types.KeyChord{Key: "e", Ctrl: true}); err != nil {
		t.Fatalf("SetChord: %v", err)
	}
	if ft.gotMethod != wire.MethodKeymapSetChord {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodKeymapSetChord)
	}
	var sent wire.SetChordArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ActionID != "Feature.Extrude" || sent.Chord != "Ctrl+E" {
		t.Errorf("sent = %+v, want Feature.Extrude / Ctrl+E (canonical)", sent)
	}
}

func TestKeymapSetAliasAndReset(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.Keymap().SetAlias("Feature.Extrude", "EXT"); err != nil {
		t.Fatalf("SetAlias: %v", err)
	}
	if ft.gotMethod != wire.MethodKeymapSetAlias {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodKeymapSetAlias)
	}
	if _, err := c.Keymap().Reset("Feature.Extrude"); err != nil {
		t.Fatalf("Reset: %v", err)
	}
	if ft.gotMethod != wire.MethodKeymapReset {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodKeymapReset)
	}
}

func TestKeymapExportImportRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"chords":{"Feature.Extrude":"Ctrl+E"},"aliases":{"Feature.Hole":"HOL"}}`)}
	c := New(ft)
	exp, err := c.Keymap().Export()
	if err != nil {
		t.Fatalf("Export: %v", err)
	}
	if ft.gotMethod != wire.MethodKeymapExport || exp.Chords["Feature.Extrude"] != "Ctrl+E" {
		t.Errorf("export = (%q, %+v), want keymap.export with Ctrl+E", ft.gotMethod, exp)
	}

	ft.reply = []byte(`{"ok":true}`)
	if _, err := c.Keymap().Import(exp); err != nil {
		t.Fatalf("Import: %v", err)
	}
	if ft.gotMethod != wire.MethodKeymapImport {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodKeymapImport)
	}
	var sent wire.KeymapExport
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Aliases["Feature.Hole"] != "HOL" {
		t.Errorf("sent = %+v, want the exported delta echoed back", sent)
	}
}

func TestKeymapResetAll(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	if _, err := New(ft).Keymap().ResetAll(); err != nil {
		t.Fatalf("ResetAll: %v", err)
	}
	if ft.gotMethod != wire.MethodKeymapResetAll {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodKeymapResetAll)
	}
}
