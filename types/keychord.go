// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"strings"
)

// ActionID identifies a bindable action in the keymap (M05-F17, Oblikovati#831): a
// registered command's id verbatim (e.g. "Feature.Extrude"), or a reserved built-in
// action id (e.g. "edit.undo"). It is a string so existing command-id call sites are
// unaffected; the host's binding engine resolves it to a command or a built-in.
type ActionID = string

// KeyChord is a keyboard shortcut: a key token plus held modifiers. This is the
// canonical, Apache-2.0 definition shared by the typed client and the host's binding
// engine; the GPL implementation maps it onto its internal modifier bitmask. The
// canonical text form ([KeyChord.String]) is "Ctrl+Alt+Shift+Key" with the modifiers
// in that fixed order, so two chords are equal iff their String values match — that
// string is the comparison key the engine uses for conflict detection.
//
// Example: KeyChord{Key: "E"}.String() == "E"; KeyChord{Key: "z", Ctrl: true}.String() == "Ctrl+Z".
type KeyChord struct {
	Key   string `json:"key"`
	Ctrl  bool   `json:"ctrl,omitempty"`
	Alt   bool   `json:"alt,omitempty"`
	Shift bool   `json:"shift,omitempty"`
}

// CanonicalKey normalizes a raw key token so case never splits a binding: a single
// letter is upper-cased ("e" and "E" are the same chord), and a multi-character token
// ("Escape", "F5") is kept verbatim.
func CanonicalKey(key string) string {
	if len([]rune(key)) == 1 {
		return strings.ToUpper(key)
	}
	return key
}

// IsZero reports whether the chord is unbound (no key) — the sentinel for "this action
// has no shortcut".
func (k KeyChord) IsZero() bool { return k.Key == "" }

// String renders the chord in canonical form ("Ctrl+Shift+E"). An unbound chord (empty
// key) renders as "".
func (k KeyChord) String() string {
	var b strings.Builder
	if k.Ctrl {
		b.WriteString("Ctrl+")
	}
	if k.Alt {
		b.WriteString("Alt+")
	}
	if k.Shift {
		b.WriteString("Shift+")
	}
	b.WriteString(CanonicalKey(k.Key))
	return b.String()
}

// ParseChord parses a canonical chord string into a KeyChord. The final '+'-separated
// token is the key; the rest are modifiers (case-insensitive: Ctrl/Control, Alt/Option,
// Shift). An empty string parses to the unbound zero chord; an empty key after modifiers
// or an unknown modifier is an error naming the offending input and the expected
// "[Ctrl+][Alt+][Shift+]Key" shape.
func ParseChord(s string) (KeyChord, error) {
	if strings.TrimSpace(s) == "" {
		return KeyChord{}, nil
	}
	parts := strings.Split(s, "+")
	key := parts[len(parts)-1]
	if key == "" {
		return KeyChord{}, fmt.Errorf("types: key chord %q has empty key; expected \"[Ctrl+][Alt+][Shift+]Key\"", s)
	}
	chord := KeyChord{Key: CanonicalKey(key)}
	for _, mod := range parts[:len(parts)-1] {
		if err := chord.applyModifier(s, mod); err != nil {
			return KeyChord{}, err
		}
	}
	return chord, nil
}

// applyModifier sets the modifier named by mod, erroring on an unknown name (with the
// full chord string for context).
func (k *KeyChord) applyModifier(chord, mod string) error {
	switch strings.ToLower(mod) {
	case "ctrl", "control":
		k.Ctrl = true
	case "alt", "option":
		k.Alt = true
	case "shift":
		k.Shift = true
	default:
		return fmt.Errorf("types: key chord %q has unknown modifier %q; expected Ctrl, Alt or Shift", chord, mod)
	}
	return nil
}
