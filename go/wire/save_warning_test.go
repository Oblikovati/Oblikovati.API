// SPDX-License-Identifier: Apache-2.0
package wire

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSaveWarningsIdentifySuccessfulPartialOutcome(t *testing.T) {
	warning := SaveWarning{Document: 7, Code: SaveWarningAssemblyCacheUnavailable, Detail: "component is dirty"}
	results := []any{
		SaveDocumentResult{FullDocumentName: "a.oad", Warnings: []SaveWarning{warning}},
		BatchSaveItemResult{Document: 7, OK: true, Warnings: []SaveWarning{warning}},
		CloseDocumentsResult{Closed: 1, Warnings: []SaveWarning{warning}},
	}
	for _, result := range results {
		encoded, err := json.Marshal(result)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(encoded), `"code":"assemblyCacheUnavailable"`) || !strings.Contains(string(encoded), `"document":7`) {
			t.Fatalf("warning absent from %s", encoded)
		}
	}
}

func TestSaveWarningsAreOptionalForExistingClients(t *testing.T) {
	var result SaveDocumentResult
	if err := json.Unmarshal([]byte(`{"fullDocumentName":"a.oad"}`), &result); err != nil {
		t.Fatal(err)
	}
	if len(result.Warnings) != 0 {
		t.Fatalf("unexpected warnings: %#v", result.Warnings)
	}
	encoded, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), "warnings") {
		t.Fatalf("empty warnings are not omitted: %s", encoded)
	}
}
