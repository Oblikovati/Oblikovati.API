// SPDX-License-Identifier: Apache-2.0

package wire

import "testing"

// TestCommandLifecycleEventNames pins the on-the-wire spelling of the command-lifecycle event pair.
// Both names are the source of truth the host emits and add-ins subscribe to; a drift here silently
// breaks every add-in listening for command begin/end, so the strings are asserted exactly. The pair
// must also be distinct (a copy-paste that left them equal would route both lifecycle phases the same).
func TestCommandLifecycleEventNames(t *testing.T) {
	if EventCommandStarted != "command.started" {
		t.Errorf("EventCommandStarted = %q, want %q", EventCommandStarted, "command.started")
	}
	if EventCommandEnded != "command.ended" {
		t.Errorf("EventCommandEnded = %q, want %q", EventCommandEnded, "command.ended")
	}
	if EventCommandStarted == EventCommandEnded {
		t.Error("command-started and command-ended must be distinct event names")
	}
}
