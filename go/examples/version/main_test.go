// SPDX-License-Identifier: Apache-2.0

package main

import (
	"oblikovati.org/api"
	"oblikovati.org/api/client"
	"testing"
)

func TestVersionExampleUsesSharedContract(t *testing.T) {
	got, err := client.New(exampleHost{}).Application().ApiVersion()
	if err != nil || got.Version != api.Version {
		t.Fatalf("version = %+v, %v; expected %s", got, err, api.Version)
	}
	if _, err := (exampleHost{}).Call("bad", nil); err == nil {
		t.Fatal("bad method should fail")
	}
}
