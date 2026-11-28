// SPDX-License-Identifier: Apache-2.0

package wasmaddin

import (
	"errors"
	"testing"

	"oblikovati.org/api"
	"oblikovati.org/api/client"
)

type fakeBridge struct {
	lastMethod string
	lastReq    []byte
	reply      []byte
	ok         bool
}

func (f *fakeBridge) Call(method string, req []byte) ([]byte, bool) {
	f.lastMethod = method
	f.lastReq = append([]byte(nil), req...)
	return f.reply, f.ok
}

type fakeHandler struct {
	id            string
	manifest      string
	activateErr   error
	deactivateErr error
	activateCalls int
	deactCalls    int
	caller        client.Caller
	events        [][]byte
}

func (f *fakeHandler) ID() string       { return f.id }
func (f *fakeHandler) Manifest() string { return f.manifest }
func (f *fakeHandler) Activate(caller client.Caller) error {
	f.activateCalls++
	f.caller = caller
	return f.activateErr
}
func (f *fakeHandler) Deactivate() error {
	f.deactCalls++
	return f.deactivateErr
}
func (f *fakeHandler) Notify(ev []byte) {
	f.events = append(f.events, append([]byte(nil), ev...))
}

type fakeAutomationHandler struct {
	fakeHandler
	autoReply []byte
	autoErr   error
	autoCalls int
}

func (f *fakeAutomationHandler) Automation(_ string, _ []byte) ([]byte, error) {
	f.autoCalls++
	return f.autoReply, f.autoErr
}

func TestDispatcherValidation(t *testing.T) {
	if _, err := NewDispatcher(nil, &fakeBridge{}); err == nil {
		t.Fatal("expected error for nil handler")
	}
	if _, err := NewDispatcher(&fakeHandler{}, nil); err == nil {
		t.Fatal("expected error for nil bridge")
	}
}

func TestDispatcherLifecycleAndCaller(t *testing.T) {
	bridge := &fakeBridge{reply: []byte(`{"ok":true}`), ok: true}
	h := &fakeHandler{id: "com.oblikovati.test", manifest: `{"id":"com.oblikovati.test"}`}
	d, err := NewDispatcher(h, bridge)
	if err != nil {
		t.Fatalf("NewDispatcher: %v", err)
	}
	if d.ID() != h.id || d.Manifest() != h.manifest {
		t.Fatalf("metadata mismatch: id=%q manifest=%q", d.ID(), d.Manifest())
	}
	if d.APIMajor() != api.Major() || d.APIMinor() != api.Minor() {
		t.Fatalf("version mismatch: %d.%d != %d.%d", d.APIMajor(), d.APIMinor(), api.Major(), api.Minor())
	}
	if d.HasAutomation() {
		t.Fatal("HasAutomation = true, want false")
	}

	// Notify before activation is a no-op.
	if rc := d.Notify([]byte("ignored")); rc != StatusOK || len(h.events) != 0 {
		t.Fatalf("inactive notify: rc=%d events=%d", rc, len(h.events))
	}
	// Activate is idempotent.
	if rc := d.Activate(); rc != StatusOK {
		t.Fatalf("first Activate = %d, want %d", rc, StatusOK)
	}
	if rc := d.Activate(); rc != StatusOK || h.activateCalls != 1 {
		t.Fatalf("second Activate: rc=%d calls=%d", rc, h.activateCalls)
	}

	// Caller forwards through bridge.
	resp, err := h.caller.Call("documents.list", []byte(`{}`))
	if err != nil || string(resp) != `{"ok":true}` {
		t.Fatalf("Call ok: resp=%q err=%v", resp, err)
	}
	if bridge.lastMethod != "documents.list" || string(bridge.lastReq) != `{}` {
		t.Fatalf("bridge received method=%q req=%q", bridge.lastMethod, bridge.lastReq)
	}

	// Bridge error maps to Go error.
	bridge.ok = false
	bridge.reply = []byte("invalid request")
	if _, err := h.caller.Call("documents.list", nil); err == nil || err.Error() != "invalid request" {
		t.Fatalf("Call err = %v, want 'invalid request'", err)
	}
	bridge.reply = nil
	if _, err := h.caller.Call("documents.list", nil); err == nil {
		t.Fatal("expected error for empty error reply")
	}

	// Active Notify forwards bytes.
	if rc := d.Notify([]byte("event-1")); rc != StatusOK || len(h.events) != 1 || string(h.events[0]) != "event-1" {
		t.Fatalf("active notify: rc=%d events=%q", rc, h.events)
	}

	// Deactivate is idempotent and blocks subsequent host calls.
	if rc := d.Deactivate(); rc != StatusOK {
		t.Fatalf("first Deactivate = %d", rc)
	}
	if rc := d.Deactivate(); rc != StatusOK || h.deactCalls != 1 {
		t.Fatalf("second Deactivate: rc=%d calls=%d", rc, h.deactCalls)
	}
	if _, err := h.caller.Call("documents.list", nil); err == nil {
		t.Fatal("expected error calling host after deactivation")
	}
}

func TestDispatcherErrorsAndAutomation(t *testing.T) {
	h := &fakeAutomationHandler{
		fakeHandler: fakeHandler{id: "com.oblikovati.auto", activateErr: errors.New("boom")},
		autoReply:   []byte(`{"v":1}`),
	}
	d, err := NewDispatcher(h, &fakeBridge{ok: true})
	if err != nil {
		t.Fatalf("NewDispatcher: %v", err)
	}
	if !d.HasAutomation() {
		t.Fatal("HasAutomation = false, want true")
	}
	if rc := d.Activate(); rc != StatusErr {
		t.Fatalf("Activate error = %d, want %d", rc, StatusErr)
	}
	if _, err := d.Automation("ping", nil); err == nil {
		t.Fatal("expected Automation error while inactive")
	}

	h.activateErr = nil
	if rc := d.Activate(); rc != StatusOK {
		t.Fatalf("Activate retry = %d", rc)
	}
	out, err := d.Automation("ping", []byte(`{}`))
	if err != nil || string(out) != `{"v":1}` || h.autoCalls != 1 {
		t.Fatalf("Automation: out=%q err=%v calls=%d", out, err, h.autoCalls)
	}

	// Non-automation handler rejects Automation when active.
	plain, _ := NewDispatcher(&fakeHandler{id: "com.oblikovati.plain"}, &fakeBridge{ok: true})
	_ = plain.Activate()
	if _, err := plain.Automation("ping", nil); err == nil {
		t.Fatal("expected unsupported automation error")
	}

	// Deactivate error returns StatusErr.
	h.deactivateErr = errors.New("stop failed")
	if rc := d.Deactivate(); rc != StatusErr {
		t.Fatalf("Deactivate error = %d, want %d", rc, StatusErr)
	}
}
