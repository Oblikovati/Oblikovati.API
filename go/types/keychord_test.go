// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestKeyChordStringCanonical(t *testing.T) {
	cases := []struct {
		chord KeyChord
		want  string
	}{
		{KeyChord{Key: "E"}, "E"},
		{KeyChord{Key: "e"}, "E"}, // single letters upper-cased
		{KeyChord{Key: "z", Ctrl: true}, "Ctrl+Z"},
		{KeyChord{Key: "z", Ctrl: true, Shift: true}, "Ctrl+Shift+Z"},
		{KeyChord{Key: "Escape"}, "Escape"}, // multi-char token kept verbatim
		{KeyChord{Key: "F5", Alt: true}, "Alt+F5"},
		{KeyChord{}, ""}, // unbound
	}
	for _, c := range cases {
		if got := c.chord.String(); got != c.want {
			t.Errorf("%#v.String() = %q, want %q", c.chord, got, c.want)
		}
	}
}

func TestParseChordRoundTrip(t *testing.T) {
	for _, s := range []string{"E", "Ctrl+Z", "Ctrl+Shift+Z", "Alt+F5", "Escape", "Ctrl+Alt+Shift+A"} {
		chord, err := ParseChord(s)
		if err != nil {
			t.Fatalf("ParseChord(%q): %v", s, err)
		}
		if got := chord.String(); got != s {
			t.Errorf("ParseChord(%q).String() = %q, want round-trip", s, got)
		}
	}
}

func TestParseChordCaseInsensitiveModifiers(t *testing.T) {
	chord, err := ParseChord("control+shift+e")
	if err != nil {
		t.Fatalf("ParseChord: %v", err)
	}
	if chord.String() != "Ctrl+Shift+E" {
		t.Errorf("got %q, want Ctrl+Shift+E", chord.String())
	}
}

func TestParseChordEmptyIsUnbound(t *testing.T) {
	chord, err := ParseChord("")
	if err != nil {
		t.Fatalf("ParseChord(\"\"): %v", err)
	}
	if !chord.IsZero() {
		t.Errorf("empty string should parse to the unbound zero chord, got %#v", chord)
	}
}

func TestParseChordRejectsUnknownModifier(t *testing.T) {
	_, err := ParseChord("Meta+E")
	if err == nil {
		t.Fatal("unknown modifier should error")
	}
	if !containsAll(err.Error(), "Meta", "Ctrl") {
		t.Errorf("error %q should name the offending modifier and the expected set", err)
	}
}

func TestParseChordRejectsEmptyKey(t *testing.T) {
	_, err := ParseChord("Ctrl+")
	if err == nil {
		t.Fatal("a chord with no key should error")
	}
	if !containsAll(err.Error(), "Ctrl+", "empty key") {
		t.Errorf("error %q should name the offending chord and the empty-key cause", err)
	}
}

// containsAll reports whether s contains every substring.
func containsAll(s string, subs ...string) bool {
	for _, sub := range subs {
		found := false
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
