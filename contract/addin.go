// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// AddInRegistry is the in-process contract for the host's add-in registry — the
// session-free slice of the ApplicationAddIns surface: what is installed, what is
// running, and when each entry loads. The GPL implementation (app.AddInManager)
// satisfies it (compile-time asserted there); lifecycle calls that need the live
// session cross the wire instead (addins.activate / addins.deactivate).
//
// Example:
//
//	for _, id := range reg.Registered() { fmt.Println(id, reg.IsActive(id)) }
type AddInRegistry interface {
	// Registered returns the registered add-in ids in registration order.
	Registered() []string
	// IsActive reports whether an add-in is currently active.
	IsActive(id string) bool
	// LoadBehavior returns when the host activates the add-in on startup.
	LoadBehavior(id string) types.AddInLoadBehavior
}

// AddInAutomation is the in-process contract for an add-in's automation surface —
// the ApplicationAddIn.Automation equivalent. An add-in that wants to be callable by
// other add-ins implements it (a shared-library add-in via the optional
// ObkAddInAutomation export, a first-party in-process add-in directly); the host
// routes addins.callAutomation to it. Method and the byte payloads are the target's
// own contract — the host passes them through opaquely.
//
// An automation handler runs on the host's session goroutine: it must return
// promptly and must NOT make synchronous host calls (they would wait on the very
// dispatcher this call is occupying).
type AddInAutomation interface {
	// CallAutomation runs one automation method with an opaque JSON payload and
	// returns the opaque JSON reply.
	CallAutomation(method string, args []byte) ([]byte, error)
}
