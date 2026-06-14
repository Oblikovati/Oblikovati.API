// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"strings"
	"testing"

	"oblikovati.org/api/types"
)

func TestNameValueMapOrderingAndAccess(t *testing.T) {
	m := TransientObjects{}.CreateNameValueMap()
	m.Set("a", types.IntegerVariant(1))
	m.Set("c", types.StringVariant("last"))
	if err := m.Insert("b", types.BoolVariant(true), 1, true); err != nil {
		t.Fatalf("insert before c: %v", err)
	}
	if err := m.Insert("d", types.DoubleVariant(4), 2, false); err != nil {
		t.Fatalf("insert after c: %v", err)
	}
	got := m.Names()
	want := []string{"a", "b", "c", "d"}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Names() = %v, want %v (insert-before/after ordering)", got, want)
		}
	}
	if v, ok := m.Value("b"); !ok {
		t.Error("Value(b) missing")
	} else if b, ok := v.Bool(); !ok || !b {
		t.Errorf("Value(b) = %v", v)
	}
	if err := m.Insert("a", types.IntegerVariant(0), 0, true); err == nil {
		t.Error("inserting a duplicate name must error")
	}
	m.Set("a", types.IntegerVariant(2)) // replace keeps position
	if name, _ := m.NameAt(0); name != "a" {
		t.Errorf("Set must replace in place, NameAt(0) = %q", name)
	}
	if !m.Remove("b") || m.Remove("b") {
		t.Error("Remove must report existence exactly once")
	}
	if m.Count() != 3 {
		t.Errorf("Count = %d, want 3", m.Count())
	}
	m.Clear()
	if m.Count() != 0 {
		t.Error("Clear must empty the map")
	}
}

// TestNameValueMapJSONIsCanonical pins the wire encoding of the options bag.
func TestNameValueMapJSONIsCanonical(t *testing.T) {
	m := &NameValueMap{}
	m.Set("tolerance", types.UnitVariant(0.1, "mm"))
	m.Set("visible", types.BoolVariant(true))
	got, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	want := `[{"name":"tolerance","value":{"type":"double","value":0.1,"unit":"mm"}},` +
		`{"name":"visible","value":{"type":"boolean","value":true}}]`
	if string(got) != want {
		t.Errorf("JSON = %s, want %s", got, want)
	}
	var back NameValueMap
	if err := json.Unmarshal(got, &back); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if v, ok := back.Value("tolerance"); !ok || v.Unit() != "mm" {
		t.Errorf("round trip lost the unit: %+v", v)
	}
}

func TestObjectCollectionMutation(t *testing.T) {
	line := types.NewObjectRef("sketchLine", 7)
	arc := types.NewObjectRef("sketchArc", 9)
	c := TransientObjects{}.CreateObjectCollection(line, arc)
	c.Add(line) // duplicates allowed, as in the reference
	if c.Count() != 3 {
		t.Fatalf("Count = %d, want 3", c.Count())
	}
	if !c.RemoveRef(line) {
		t.Error("RemoveRef must remove the first match")
	}
	if got, _ := c.At(0); got != arc {
		t.Errorf("At(0) = %+v, want the arc", got)
	}
	if err := c.RemoveAt(5); err == nil {
		t.Error("out-of-range RemoveAt must error")
	}
	if err := c.RemoveAt(1); err != nil || c.Count() != 1 {
		t.Errorf("RemoveAt(1): %v, Count = %d", err, c.Count())
	}
	c.Clear()
	if c.Count() != 0 {
		t.Error("Clear must empty the collection")
	}
}

// TestObjectCollectionJSONIsCanonical pins the wire reference-list encoding.
func TestObjectCollectionJSONIsCanonical(t *testing.T) {
	c := &ObjectCollection{}
	c.Add(types.NewObjectRef("extrudeFeature", 12))
	got, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	want := `[{"kind":"extrudeFeature","id":12}]`
	if string(got) != want {
		t.Errorf("JSON = %s, want %s", got, want)
	}
}

func TestObjectCollectionByVariantKeyedAccess(t *testing.T) {
	c := TransientObjects{}.CreateObjectCollectionByVariant()
	face := types.NewObjectRef("face", 3)
	if err := c.Add("seed", face); err != nil {
		t.Fatalf("Add: %v", err)
	}
	if err := c.Add("seed", face); err == nil {
		t.Error("a duplicate key must error")
	}
	if got, ok := c.ByKey("seed"); !ok || got != face {
		t.Errorf("ByKey = (%+v, %v)", got, ok)
	}
	if key, _ := c.KeyAt(0); key != "seed" {
		t.Errorf("KeyAt(0) = %q", key)
	}
	if !c.Remove("seed") || c.Count() != 0 {
		t.Error("Remove must delete the keyed entry")
	}
}

// Guards the three index-addressed accessors that share errIndexRangeFmt: each
// must error out of range and report the offending index and the bound.
func TestObjectCollectionByVariantIndexOutOfRange(t *testing.T) {
	c := TransientObjects{}.CreateObjectCollectionByVariant()
	if err := c.Add("seed", types.NewObjectRef("face", 3)); err != nil {
		t.Fatalf("Add: %v", err)
	}
	if _, err := c.KeyAt(5); err == nil {
		t.Error("KeyAt out of range must error")
	}
	if _, err := c.At(5); err == nil {
		t.Error("At out of range must error")
	}
	if err := c.RemoveAt(5); err == nil {
		t.Error("RemoveAt out of range must error")
	} else if want := "index 5 out of range [0,1)"; !strings.Contains(err.Error(), want) {
		t.Errorf("RemoveAt error = %q, want it to contain %q", err, want)
	}
}
