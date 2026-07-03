// SPDX-License-Identifier: Apache-2.0

package featureargs

import (
	"encoding/json"
	"reflect"
	"testing"
)

// TestEveryArgKindIsNonEmptyAndUnique guards the struct-to-kind binding: a blank or
// duplicated Kind would make the host parity guard ambiguous or the add-in envelope
// untaggable (#1616).
func TestEveryArgKindIsNonEmptyAndUnique(t *testing.T) {
	seen := map[string]string{}
	for _, a := range All() {
		kind := a.Kind()
		if kind == "" {
			t.Errorf("%T.Kind() is empty — every promoted arg type must name its feature kind", a)
		}
		if prev, dup := seen[kind]; dup {
			t.Errorf("kind %q is returned by both %s and %T — kinds must be unique", kind, prev, a)
		}
		seen[kind] = reflect.TypeOf(a).Name()
	}
}

// TestArgsMarshalRoundTrip marshals each promoted arg type and decodes it back into a
// fresh value of the same type, asserting equality. The host (addin/opregistry) decodes
// the wire bytes into THESE types, so a field the JSON codec cannot round-trip here is a
// silent drift there; this catches it in the Apache module (#1616).
func TestArgsMarshalRoundTrip(t *testing.T) {
	for _, a := range All() {
		raw, err := json.Marshal(a)
		if err != nil {
			t.Fatalf("marshal %T: %v", a, err)
		}
		fresh := reflect.New(reflect.TypeOf(a)).Interface()
		if err := json.Unmarshal(raw, fresh); err != nil {
			t.Fatalf("unmarshal %T from %s: %v", a, raw, err)
		}
		got := reflect.ValueOf(fresh).Elem().Interface()
		if !reflect.DeepEqual(a, got) {
			t.Errorf("%T did not round-trip: sent %#v, got %#v (json %s)", a, a, got, raw)
		}
	}
}

// TestKindsMatchAll keeps the Kinds() convenience in lockstep with All().
func TestKindsMatchAll(t *testing.T) {
	if got, want := len(Kinds()), len(All()); got != want {
		t.Fatalf("Kinds() has %d entries, All() has %d", got, want)
	}
	for i, a := range All() {
		if Kinds()[i] != a.Kind() {
			t.Errorf("Kinds()[%d]=%q but All()[%d].Kind()=%q", i, Kinds()[i], i, a.Kind())
		}
	}
}
