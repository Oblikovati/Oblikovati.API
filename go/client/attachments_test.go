// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestDocumentsAttachmentLifecycle(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"name":"loads","kind":3331,"fullFileName":"/w/loads.csv","status":49666,"browserVisible":true}`)}
	c := New(ft)
	rec, err := c.Documents().AddAttachment(wire.AddAttachmentArgs{
		Document: 4, Name: "loads", Kind: types.AttachmentLinked, FullFileName: "/w/loads.csv",
	})
	if err != nil || rec.Kind != types.AttachmentLinked || rec.Status != types.ReferenceUpToDate {
		t.Fatalf("AddAttachment = (%+v, %v), want the linked record", rec, err)
	}
	var sent wire.AddAttachmentArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Name != "loads" || sent.Document != 4 {
		t.Errorf("sent = %s, want the attach args", ft.gotReq)
	}

	ft.reply = []byte(`{"attachments":[{"name":"loads","kind":3331,"status":49668}]}`)
	lst, err := c.Documents().Attachments(4)
	if err != nil || len(lst.Attachments) != 1 || lst.Attachments[0].Status != types.ReferenceMissing {
		t.Fatalf("Attachments = (%+v, %v), want the now-missing record", lst, err)
	}

	ft.reply = []byte(`{"ok":true}`)
	if _, err := c.Documents().RemoveAttachment(4, "loads"); err != nil {
		t.Fatalf("RemoveAttachment: %v", err)
	}
	var rm wire.RemoveAttachmentArgs
	if err := json.Unmarshal(ft.gotReq, &rm); err != nil || rm.Name != "loads" {
		t.Errorf("sent = %s, want the removal of loads", ft.gotReq)
	}
}
