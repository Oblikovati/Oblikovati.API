// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"
	"testing"
)

func TestTaskPanelSpecRoundTrip(t *testing.T) {
	spec := TaskPanelSpec{
		ID: "fix", Title: "Fixed Constraint",
		Controls: []PanelControlSpec{{ID: "faces"}},
		OKLabel:  "OK", CancelLabel: "Cancel",
	}
	b, _ := json.Marshal(ShowTaskPanelArgs{Panel: spec})
	var back ShowTaskPanelArgs
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if back.Panel.ID != "fix" || back.Panel.Title != "Fixed Constraint" || len(back.Panel.Controls) != 1 {
		t.Fatalf("round-trip lost data: %+v", back.Panel)
	}
}

func TestTaskPanelConstants(t *testing.T) {
	if MethodTaskPanelShow != "taskPanel.show" || MethodTaskPanelClose != "taskPanel.close" {
		t.Fatalf("methods: %q %q", MethodTaskPanelShow, MethodTaskPanelClose)
	}
	if EventTaskPanelClosed != "taskPanel.closed" {
		t.Fatalf("event: %q", EventTaskPanelClosed)
	}
	ev := TaskPanelClosedEvent{Type: EventTaskPanelClosed, ID: "fix", Accepted: true}
	if !ev.Accepted {
		t.Fatalf("event lost Accepted: %+v", ev)
	}
}
