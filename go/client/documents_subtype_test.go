// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestDocumentsRegisterSubType(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"ok":true}`)}
	c := New(ft)
	if _, err := c.Documents().RegisterSubType(wire.RegisterDocumentSubTypeArgs{
		ID: "com.x.sim.study", BaseType: "part", DisplayName: "Simulation Study",
	}); err != nil {
		t.Fatalf("RegisterSubType: %v", err)
	}
	var sent wire.RegisterDocumentSubTypeArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil || sent.BaseType != "part" {
		t.Errorf("sent = %s, want the part-based study", ft.gotReq)
	}

	ft.reply = []byte(`{"subTypes":[{"id":"com.x.sim.study","baseType":"part"}]}`)
	lst, err := c.Documents().SubTypes()
	if err != nil || len(lst.SubTypes) != 1 || lst.SubTypes[0].ID != "com.x.sim.study" {
		t.Fatalf("SubTypes = (%+v, %v), want the study", lst, err)
	}
}
