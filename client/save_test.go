// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestDocumentsOpenSaveLifecycle(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"id":4,"name":"bracket","type":"part","active":true}`)}
	c := New(ft)
	info, err := c.Documents().Open(wire.OpenDocumentArgs{FullDocumentName: "/w/bracket.obk", Visible: true})
	if err != nil || info.ID != 4 {
		t.Fatalf("Open = (%+v, %v), want document 4", info, err)
	}
	var sent wire.OpenDocumentArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || !sent.Visible || sent.FullDocumentName != "/w/bracket.obk" {
		t.Errorf("sent = %s, want a visible open of /w/bracket.obk", ft.gotReq)
	}

	ft.reply = []byte(`{"fullDocumentName":"/w/bracket.obk"}`)
	if saved, err := c.Documents().Save(4); err != nil || saved.FullDocumentName != "/w/bracket.obk" {
		t.Fatalf("Save = (%+v, %v), want the saved path", saved, err)
	}
	if saved, err := c.Documents().SaveAs(4, "/w/bracket-v2.obk"); err != nil || saved.FullDocumentName == "" {
		t.Fatalf("SaveAs = (%+v, %v), want a saved path", saved, err)
	}
	var as wire.SaveDocumentAsArgs
	if err := json.Unmarshal(ft.gotReq, &as); err != nil || as.NewFullDocumentName != "/w/bracket-v2.obk" {
		t.Errorf("sent = %s, want the new identity", ft.gotReq)
	}
}

func TestDocumentsSaveCopyAsKeepsMetadata(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"fullDocumentName":"/exports/bracket-rev3.obk"}`)}
	c := New(ft)
	_, err := c.Documents().SaveCopyAs(wire.SaveCopyAsArgs{
		Document:       4,
		TargetFileName: "/exports/bracket-rev3.obk",
		Metadata:       &types.NewFileMetadata{DisplayName: "Bracket rev3"},
	})
	if err != nil {
		t.Fatalf("SaveCopyAs: %v", err)
	}
	var sent wire.SaveCopyAsArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Metadata == nil || sent.Metadata.DisplayName != "Bracket rev3" {
		t.Errorf("sent = %s, want the copy metadata", ft.gotReq)
	}
}

func TestDocumentsBatchSaveReportsPerFileOutcomes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"saved":1,"results":[{"document":4,"ok":true},{"document":5,"ok":false,"error":"doc: no store"}]}`)}
	c := New(ft)
	r, err := c.Documents().BatchSave(wire.BatchSaveArgs{
		Operation: "saveCopyAs",
		Items: []wire.BatchSaveItem{
			{Document: 4, TargetFileName: "/exports/a.obk"},
			{Document: 5, TargetFileName: "/exports/b.obk"},
		},
	})
	if err != nil || r.Saved != 1 || len(r.Results) != 2 || r.Results[1].Error == "" {
		t.Fatalf("BatchSave = (%+v, %v), want one success and one carried failure", r, err)
	}
}

func TestOptionsSaveGroupRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"group":"save","save":{"thumbnail":79875,"saveDependents":true,"oldVersionsToKeep":3}}`)}
	c := New(ft)
	v, err := c.Options().Save()
	if err != nil || v.Thumbnail != types.ThumbnailActiveWindowOnSave || !v.SaveDependents || v.OldVersionsToKeep != 3 {
		t.Fatalf("Options().Save() = (%+v, %v), want the save policy", v, err)
	}

	ft.reply = []byte(`{"ok":true}`)
	if _, err := c.Options().SetSave(wire.SaveOptionsView{Thumbnail: types.ThumbnailNone}); err != nil {
		t.Fatalf("SetSave: %v", err)
	}
	var sent wire.OptionGroupView
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Group != wire.OptionGroupSave || sent.Save == nil {
		t.Errorf("sent = %s, want the save group write", ft.gotReq)
	}
}
