// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestFilesGetAndReferences(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"fullFileName":"/w/bracket.obk","internalName":"guid-1","revisionId":"rev-2","saveCounter":3,"loaded":true}`)}
	c := New(ft)
	info, err := c.Files().Get("/w/bracket.obk")
	if err != nil || info.InternalName != "guid-1" || info.SaveCounter != 3 {
		t.Fatalf("Get = (%+v, %v), want the identity of /w/bracket.obk", info, err)
	}
	var sent wire.GetFileArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.FullFileName != "/w/bracket.obk" {
		t.Errorf("sent = %s, want the full file name", ft.gotReq)
	}

	ft.reply = []byte(`{"references":[{"fullFileName":"/w/pin.obk","status":49668,"missing":true}]}`)
	refs, err := c.Files().References("/w/bracket.obk")
	if err != nil || len(refs.References) != 1 || refs.References[0].Status != types.ReferenceMissing {
		t.Fatalf("References = (%+v, %v), want one missing reference", refs, err)
	}
}

func TestFilesReplaceReference(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"fullFileName":"/w/pin-v2.obk","status":49669,"replaced":true}`)}
	c := New(ft)
	rec, err := c.Files().ReplaceReference(wire.ReplaceFileReferenceArgs{
		FullFileName: "/w/bracket.obk", RequestedName: "/w/pin.obk", NewFileName: "/w/pin-v2.obk",
	})
	if err != nil || rec.Status != types.ReferenceReplaced || !rec.Replaced {
		t.Fatalf("ReplaceReference = (%+v, %v), want the replaced record", rec, err)
	}
	var sent wire.ReplaceFileReferenceArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.NewFileName != "/w/pin-v2.obk" {
		t.Errorf("sent = %s, want the repair args", ft.gotReq)
	}
}

func TestDocumentsFileReferences(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"references":[{"fullFileName":"/w/pin.obk","status":49666,"documentFound":true}]}`)}
	c := New(ft)
	refs, err := c.Documents().FileReferences(7)
	if err != nil || len(refs.References) != 1 || !refs.References[0].DocumentFound {
		t.Fatalf("FileReferences = (%+v, %v), want one resolved reference", refs, err)
	}
	var sent wire.ListDocumentFileReferencesArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Document != 7 {
		t.Errorf("sent = %s, want document 7", ft.gotReq)
	}
}
