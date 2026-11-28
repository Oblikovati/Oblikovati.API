// SPDX-License-Identifier: Apache-2.0

// Package wasmaddin defines the browser WebAssembly (`js-wasm`) host-to-add-in
// bridge for Go add-ins compiled with `GOOS=js GOARCH=wasm`.
//
// Native add-ins expose C-ABI symbols via cgo (`include/oblikovati_addin.h`).
// Browser add-ins cannot use cgo or dlopen; instead, each add-in runs inside a
// dedicated Web Worker and calls [Serve] from `main`. The worker host supplies
// `globalThis.__obkHost` backed by a SharedArrayBuffer RPC channel to the
// Oblikovati Orchestrator Worker.
package wasmaddin

import (
	"errors"
	"fmt"
	"sync"

	"oblikovati.org/api"
	"oblikovati.org/api/client"
)

const (
	// Platform is the add-in catalogue platform key for browser WebAssembly builds.
	Platform = "js-wasm"

	// StatusOK matches OBK_OK in include/oblikovati_addin.h.
	StatusOK = 0

	// StatusErr matches OBK_ERR in include/oblikovati_addin.h.
	StatusErr = 1
)

// Handler is the lifecycle contract an add-in implements for [Serve].
type Handler interface {
	ID() string
	Manifest() string
	Activate(caller client.Caller) error
	Deactivate() error
	Notify(event []byte)
}

// AutomationHandler is an optional extension matching ObkAddInAutomation.
type AutomationHandler interface {
	Automation(method string, req []byte) ([]byte, error)
}

// Bridge is the low-level synchronous host RPC transport injected by the
// WebAssembly worker host (or a named fake in unit tests).
type Bridge interface {
	Call(method string, req []byte) (resp []byte, ok bool)
}

// Dispatcher adapts a [Handler] and [Bridge] to the host's lifecycle calls.
//
//	d, err := wasmaddin.NewDispatcher(myHandler, myBridge)
//	if d.Activate() != wasmaddin.StatusOK { /* report activation error */ }
type Dispatcher struct {
	handler Handler
	bridge  Bridge

	mu     sync.Mutex
	active bool
}

// NewDispatcher validates inputs and returns a lifecycle dispatcher.
//
//	d, err := wasmaddin.NewDispatcher(handler, bridge)
func NewDispatcher(handler Handler, bridge Bridge) (*Dispatcher, error) {
	if handler == nil {
		return nil, errors.New("wasmaddin: handler is nil, expected a non-nil Handler")
	}
	if bridge == nil {
		return nil, errors.New("wasmaddin: bridge is nil, expected a non-nil Bridge")
	}
	return &Dispatcher{handler: handler, bridge: bridge}, nil
}

// ID returns the add-in's stable identifier.
func (d *Dispatcher) ID() string { return d.handler.ID() }

// Manifest returns the add-in's raw manifest.json text.
func (d *Dispatcher) Manifest() string { return d.handler.Manifest() }

// APIMajor returns the oblikovati.org/api major version compiled into the add-in.
func (d *Dispatcher) APIMajor() int { return api.Major() }

// APIMinor returns the oblikovati.org/api minor version compiled into the add-in.
func (d *Dispatcher) APIMinor() int { return api.Minor() }

// HasAutomation reports whether the handler implements [AutomationHandler].
func (d *Dispatcher) HasAutomation() bool {
	_, ok := d.handler.(AutomationHandler)
	return ok
}

// Activate starts the add-in idempotently and returns [StatusOK] or [StatusErr].
func (d *Dispatcher) Activate() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.active {
		return StatusOK
	}
	d.active = true
	if err := d.handler.Activate(&bridgeCaller{disp: d}); err != nil {
		d.active = false
		return StatusErr
	}
	return StatusOK
}

// Deactivate stops the add-in idempotently and returns [StatusOK] or [StatusErr].
func (d *Dispatcher) Deactivate() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	if !d.active {
		return StatusOK
	}
	d.active = false
	if err := d.handler.Deactivate(); err != nil {
		return StatusErr
	}
	return StatusOK
}

// Notify forwards a host event payload when active; inactive calls are no-ops.
func (d *Dispatcher) Notify(event []byte) int {
	d.mu.Lock()
	active := d.active
	d.mu.Unlock()
	if !active {
		return StatusOK
	}
	d.handler.Notify(event)
	return StatusOK
}

// Automation invokes the optional add-in automation surface when active.
func (d *Dispatcher) Automation(method string, req []byte) ([]byte, error) {
	d.mu.Lock()
	active := d.active
	d.mu.Unlock()
	if !active {
		return nil, fmt.Errorf("wasmaddin: add-in %q is not active", d.handler.ID())
	}
	auto, ok := d.handler.(AutomationHandler)
	if !ok {
		return nil, fmt.Errorf("wasmaddin: add-in %q does not implement automation", d.handler.ID())
	}
	return auto.Automation(method, req)
}

type bridgeCaller struct {
	disp *Dispatcher
}

func (c *bridgeCaller) Call(method string, req []byte) ([]byte, error) {
	c.disp.mu.Lock()
	active := c.disp.active
	c.disp.mu.Unlock()
	if !active {
		return nil, fmt.Errorf("wasmaddin: host call %q rejected while %q is inactive", method, c.disp.handler.ID())
	}
	resp, ok := c.disp.bridge.Call(method, req)
	if !ok {
		if len(resp) == 0 {
			return nil, fmt.Errorf("wasmaddin: host call %q failed", method)
		}
		return nil, errors.New(string(resp))
	}
	return resp, nil
}
