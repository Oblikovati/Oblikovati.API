// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

func TestDocumentsInterestRegistry(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	record := types.DocumentInterestRecord{
		ClientID: "com.x.toolpaths", Name: "toolpath-recipes",
		InterestType: types.Interested, DataVersion: 2,
	}
	if _, err := c.Documents().AddInterest(4, record); err != nil {
		t.Fatalf("AddInterest: %v", err)
	}
	var sent wire.AddDocumentInterestArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.Interest.DataVersion != 2 {
		t.Errorf("sent = %s, want the migrating interest", ft.gotReq)
	}

	ft.reply = []byte(`{"interests":[{"clientId":"com.x.toolpaths","name":"toolpath-recipes","interestType":68866}]}`)
	lst, err := c.Documents().Interests(4)
	if err != nil || len(lst.Interests) != 1 || lst.Interests[0].InterestType != types.Interested {
		t.Fatalf("Interests = (%+v, %v), want the registered record", lst, err)
	}

	ft.reply = []byte(`{"hasInterest":true}`)
	has, err := c.Documents().HasInterest(4, "com.x.toolpaths")
	if err != nil || !has.HasInterest {
		t.Fatalf("HasInterest = (%+v, %v), want true", has, err)
	}

	ft.reply = []byte(`{"ok":true}`)
	if _, err := c.Documents().RemoveInterest(4, "com.x.toolpaths", "toolpath-recipes"); err != nil {
		t.Fatalf("RemoveInterest: %v", err)
	}
	var rm wire.RemoveDocumentInterestArgs
	if err := json.Unmarshal(ft.gotReq, &rm); err != nil || rm.ClientID != "com.x.toolpaths" {
		t.Errorf("sent = %s, want the removal args", ft.gotReq)
	}
}
