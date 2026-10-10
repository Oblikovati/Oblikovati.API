// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"
	"testing"
)

// TestActionTypeNamesRoundTrip: a mask splits into stable-order spellings and
// re-combines to the same mask.
func TestActionTypeNamesRoundTrip(t *testing.T) {
	mask := ActionEdit | ActionDelete
	names := mask.Names()
	if !reflect.DeepEqual(names, []string{"edit", "delete"}) {
		t.Fatalf("Names() = %v, want [edit delete] (emission order)", names)
	}
	got, ok := ActionTypeMask(names)
	if !ok || got != mask {
		t.Errorf("ActionTypeMask(%v) = %v,%v; want %v,true", names, got, ok, mask)
	}
}

// TestActionNoneHasNoNames: the empty mask carries no spellings and reports every
// bit clear.
func TestActionNoneHasNoNames(t *testing.T) {
	if names := ActionNone.Names(); names != nil {
		t.Errorf("ActionNone.Names() = %v, want nil", names)
	}
	if ActionNone.Has(ActionEdit) {
		t.Error("ActionNone.Has(ActionEdit) = true, want false")
	}
}

// TestActionTypeMaskRejectsUnknown: an unknown spelling reports false and keeps
// the bits parsed before it.
func TestActionTypeMaskRejectsUnknown(t *testing.T) {
	got, ok := ActionTypeMask([]string{"rename", "bogus"})
	if ok {
		t.Error("ActionTypeMask with unknown spelling = ok true, want false")
	}
	if got != ActionRename {
		t.Errorf("partial mask = %v, want %v (rename only)", got, ActionRename)
	}
}

// TestParseActionType resolves each spelling to its single bit.
func TestParseActionType(t *testing.T) {
	for name, want := range map[string]ActionType{"edit": ActionEdit, "rename": ActionRename, "delete": ActionDelete} {
		if got, ok := ParseActionType(name); !ok || got != want {
			t.Errorf("ParseActionType(%q) = %v,%v; want %v,true", name, got, ok, want)
		}
	}
	if _, ok := ParseActionType("suppress"); ok {
		t.Error("ParseActionType(suppress) = ok true, want false")
	}
}
