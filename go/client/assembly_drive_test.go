// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestAssemblyDrivePreviewRoutesAndMarshals(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"frames":[{"value":0,"placements":[{"occurrence":2,"transform":[1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1]}]}],"stoppedByCollision":true,"stoppedAtStep":1}`)}
	c := New(ft)

	args := wire.DriveJointArgs{Joint: 7, Settings: wire.DriveSettingsDTO{
		Variable: "angular", Start: 0, End: 3.14, Step: 0.5, CollisionDetection: true}}
	r, err := c.AssemblyDrive().Preview(args)
	if err != nil {
		t.Fatalf("Preview: %v", err)
	}
	if ft.gotMethod != wire.MethodAssemblyDrivePreview {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAssemblyDrivePreview)
	}
	if len(r.Frames) != 1 || !r.StoppedByCollision || r.StoppedAtStep != 1 {
		t.Fatalf("decoded = %+v, want one frame stopped by collision at step 1", r)
	}

	var sent wire.DriveJointArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Joint != 7 || sent.Settings.Variable != "angular" || sent.Settings.Step != 0.5 {
		t.Errorf("sent = %+v, want joint 7 / angular / step 0.5", sent)
	}
}
