// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestModelSelectionMutation(t *testing.T) {
	ft := &fakeTransport{}
	c := New(ft)

	ft.reply = []byte(`{"count":2,"kinds":[0,0],"refs":["face/AAA","face/BBB"]}`)
	res, err := c.Model().Select([]string{"face/AAA", "face/BBB"}, "replace")
	if err != nil || res.Count != 2 {
		t.Fatalf("Select = (%+v, %v), want count 2", res, err)
	}
	var sa wire.SelectArgs
	if err := json.Unmarshal(ft.gotReq, &sa); err != nil || sa.Mode != "replace" || len(sa.Refs) != 2 {
		t.Errorf("sent select args = %s, want replace with 2 refs", ft.gotReq)
	}

	ft.reply = []byte(`{"count":1,"kinds":[0],"refs":["face/AAA"]}`)
	if res, err := c.Model().Deselect([]string{"face/BBB"}); err != nil || res.Count != 1 {
		t.Fatalf("Deselect = (%+v, %v), want count 1", res, err)
	}

	ft.reply = []byte(`{"count":0,"kinds":[],"refs":[]}`)
	if res, err := c.Model().ClearSelection(); err != nil || res.Count != 0 {
		t.Fatalf("ClearSelection = (%+v, %v), want count 0", res, err)
	}
}
