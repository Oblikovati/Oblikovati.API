// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"strings"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

const fractionText = "⌀1/2"
const fractionJSON = `[{"start":1,"end":4}]`
const fractionReply = `{"entityId":42,"style":{"content":"⌀1/2","height":0.5,"stackedFractions":` + fractionJSON + `}}`

func TestSketchTextFractionTypedRoundTrip(t *testing.T) {
	ranges := []types.SketchTextFraction{{Start: 1, End: 4}}
	host := &fakeTransport{reply: []byte(`{"entityId":42}`)}
	sketch := New(host).Sketch()
	if _, err := sketch.AddTextWith(wire.AddTextArgs{Text: fractionText, StackedFractions: ranges}); err != nil {
		t.Fatal(err)
	}
	assertFractionRequest(t, host, fractionJSON)
	host.reply = []byte(fractionReply)
	edited, err := sketch.EditText(wire.EditTextArgs{EntityID: 42, StackedFractions: &ranges})
	if err != nil || len(edited.Style.StackedFractions) != 1 || edited.Style.StackedFractions[0] != ranges[0] {
		t.Fatalf("style %+v: expected stacked range %v, got %v", edited.Style, ranges[0], err)
	}
	assertFractionRequest(t, host, fractionJSON)
	actual, err := sketch.GetText(0, 42)
	if err != nil || actual.Style.Content != fractionText || len(actual.Style.StackedFractions) != 1 {
		t.Fatalf("style %+v: expected unchanged source and one stacked range, got %v", actual.Style, err)
	}
}

func TestSketchTextFractionEditOmissionAndClear(t *testing.T) {
	host := &fakeTransport{reply: []byte(fractionReply)}
	sketch := New(host).Sketch()
	if _, err := sketch.EditText(wire.EditTextArgs{EntityID: 42}); err != nil {
		t.Fatal(err)
	}
	assertFractionRequest(t, host, "")
	empty := []types.SketchTextFraction{}
	if _, err := sketch.EditText(wire.EditTextArgs{EntityID: 42, StackedFractions: &empty}); err != nil {
		t.Fatal(err)
	}
	assertFractionRequest(t, host, "[]")
}

func TestSketchTextFractionRejectsMalformedClientPayloads(t *testing.T) {
	host := &fakeTransport{reply: []byte(strings.Replace(fractionReply, `"end":4`, `"end":6`, 1))}
	sketch := New(host).Sketch()
	ranges := []types.SketchTextFraction{{Start: 1, End: 4}}
	if _, err := sketch.AddTextWith(wire.AddTextArgs{Text: "⌀123", StackedFractions: ranges}); err == nil || host.gotMethod != "" {
		t.Fatalf("method %q, error %v: expected invalid content to fail before transport", host.gotMethod, err)
	}
	if _, err := sketch.GetText(0, 42); err == nil || !strings.Contains(err.Error(), "4 Unicode scalars") {
		t.Fatalf("got %v: expected the invalid host range to fail validation", err)
	}
}

func assertFractionRequest(t *testing.T, host *fakeTransport, want string) {
	t.Helper()
	var body map[string]json.RawMessage
	if err := json.Unmarshal(host.gotReq, &body); err != nil {
		t.Fatal(err)
	}
	if got := string(body["stackedFractions"]); got != want {
		t.Fatalf("stackedFractions %q: expected %q in %s", got, want, host.gotReq)
	}
}
