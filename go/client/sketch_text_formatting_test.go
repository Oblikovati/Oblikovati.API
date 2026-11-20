// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"strings"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

const formattedSketchTextReply = `{"entityId":42,"style":{"content":"A🦀界","height":0.5,"formattingRuns":[{"start":1,"end":2,"bold":true}]}}`

func TestSketchTextFormattingRoundTrips(t *testing.T) {
	runs := []types.SketchTextRun{{Start: 1, End: 2, Bold: true}}
	text := "A🦀界"
	ft := &fakeTransport{reply: []byte(`{"entityId":42}`)}
	sketch := New(ft).Sketch()
	if _, err := sketch.AddTextWith(wire.AddTextArgs{Text: text, FormattingRuns: runs}); err != nil {
		t.Fatal(err)
	}
	assertTextRunsRequest(t, ft.gotReq, `[{"start":1,"end":2,"bold":true}]`)
	ft.reply = []byte(formattedSketchTextReply)
	result, err := sketch.EditText(wire.EditTextArgs{EntityID: 42, FormattingRuns: &runs})
	assertBoldSketchTextResult(t, result, err)
	assertTextRunsRequest(t, ft.gotReq, `[{"start":1,"end":2,"bold":true}]`)
	result, err = sketch.GetText(0, 42)
	assertBoldSketchTextResult(t, result, err)
}

func assertBoldSketchTextResult(t *testing.T, result wire.SketchTextResult, err error) {
	t.Helper()
	runs := result.Style.FormattingRuns
	if err != nil || result.Style.Content != "A🦀界" || len(runs) != 1 || !runs[0].Bold {
		t.Fatalf("result %+v: expected one bold run in A🦀界, got %v", result, err)
	}
}

func TestSketchTextFormattingClearAndLegacyOmission(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":42,"style":{"content":"abc","height":0.5}}`)}
	sketch := New(ft).Sketch()
	empty := []types.SketchTextRun{}
	result, err := sketch.EditText(wire.EditTextArgs{EntityID: 42, FormattingRuns: &empty})
	if err != nil || len(result.Style.FormattingRuns) != 0 {
		t.Fatalf("legacy style %+v: expected no runs, got %v", result, err)
	}
	assertTextRunsRequest(t, ft.gotReq, `[]`)
	if _, err := sketch.EditText(wire.EditTextArgs{EntityID: 42}); err != nil {
		t.Fatal(err)
	}
	assertTextRunsRequest(t, ft.gotReq, "")
}

func assertTextRunsRequest(t *testing.T, request []byte, expected string) {
	t.Helper()
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(request, &payload); err != nil {
		t.Fatal(err)
	}
	if actual := string(payload["formattingRuns"]); actual != expected {
		t.Fatalf("formattingRuns %q: expected %q in %s", actual, expected, request)
	}
}

func TestSketchTextClientRejectsInvalidRuns(t *testing.T) {
	ft := &fakeTransport{reply: []byte(strings.Replace(formattedSketchTextReply, `"end":2`, `"end":4`, 1))}
	sketch := New(ft).Sketch()
	invalid := []types.SketchTextRun{{End: 2}}
	if _, err := sketch.AddTextWith(wire.AddTextArgs{Text: "🦀", FormattingRuns: invalid}); err == nil || ft.gotMethod != "" {
		t.Fatalf("invalid run must fail before transport, method %q, error %v", ft.gotMethod, err)
	}
	if _, err := sketch.GetText(0, 42); err == nil || !strings.Contains(err.Error(), "3 Unicode scalars") {
		t.Fatalf("got %v: expected invalid host result to be rejected", err)
	}
}

func TestSketchTextLineSpacingPreserveAndReset(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":42,"style":{"content":"a\nb","height":0.5,"lineSpacing":{"mode":"double"}}}`)}
	sketch := New(ft).Sketch()
	result, err := sketch.EditText(wire.EditTextArgs{EntityID: 42})
	if err != nil || result.Style.LineSpacing == nil || result.Style.LineSpacing.Mode != types.TextLineSpacingDouble {
		t.Fatalf("style %+v: expected stored double spacing, got %v", result.Style, err)
	}
	if strings.Contains(string(ft.gotReq), `"lineSpacing"`) {
		t.Fatalf("request %s: expected omitted spacing for a partial edit", ft.gotReq)
	}
	reset := &types.SketchTextLineSpacing{Mode: types.TextLineSpacingSingle}
	if _, err := sketch.EditText(wire.EditTextArgs{EntityID: 42, LineSpacing: reset}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(ft.gotReq), `"lineSpacing":{"mode":"single"}`) {
		t.Fatalf("request %s: expected explicit single spacing", ft.gotReq)
	}
}

func TestSketchTextLineSpacingRejectsMalformedPayloads(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"entityId":42,"style":{"content":"a","height":0.5,"lineSpacing":{"mode":"exactly"}}}`)}
	sketch := New(ft).Sketch()
	invalid := &types.SketchTextLineSpacing{Mode: types.TextLineSpacingMultiple}
	if _, err := sketch.AddTextWith(wire.AddTextArgs{Text: "a", LineSpacing: invalid}); err == nil || ft.gotMethod != "" {
		t.Fatalf("invalid spacing must fail before transport, method %q, error %v", ft.gotMethod, err)
	}
	if _, err := sketch.GetText(0, 42); err == nil || !strings.Contains(err.Error(), "positive finite") {
		t.Fatalf("got %v: expected malformed host spacing to be rejected", err)
	}
}
