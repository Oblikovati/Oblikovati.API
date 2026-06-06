// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati/api/types"
	"oblikovati/api/wire"
)

func TestImportMarshalsRequestAndDecodesReply(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"bodyCount":1,"solid":true,"warnings":["w"]}`)}
	c := New(ft)

	got, err := c.Documents().Import(wire.ImportRequest{Path: "bolt.stl", Format: string(types.FormatSTL)})
	if err != nil {
		t.Fatalf("Import: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentsImport {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentsImport)
	}
	var sent wire.ImportRequest
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Path != "bolt.stl" || sent.Format != "stl" {
		t.Errorf("sent = %+v, want path=bolt.stl format=stl", sent)
	}
	if got.BodyCount != 1 || !got.Solid || len(got.Warnings) != 1 {
		t.Errorf("decoded = %+v, want bodyCount=1 solid warnings=[w]", got)
	}
}

func TestExportMarshalsResolutionAndDecodesTriangleCount(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"triangleCount":12}`)}
	c := New(ft)

	got, err := c.Documents().Export(wire.ExportRequest{
		Path: "p.stl", Format: string(types.FormatSTL), Resolution: string(types.ResolutionHigh),
	})
	if err != nil {
		t.Fatalf("Export: %v", err)
	}
	if ft.gotMethod != wire.MethodDocumentsExport {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodDocumentsExport)
	}
	var sent wire.ExportRequest
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Resolution != "high" {
		t.Errorf("sent resolution = %q, want high", sent.Resolution)
	}
	if got.TriangleCount != 12 {
		t.Errorf("decoded triangleCount = %d, want 12", got.TriangleCount)
	}
}
