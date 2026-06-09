// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"oblikovati.org/api/wire"
)

func TestModelReferenceKeysCallsMethod(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"bodies":[{"faces":[{"key":"k0","point":[1,2,3]}],"edges":[],"vertices":[]}]}`)}
	c := New(ft)

	res, err := c.Model().ReferenceKeys()
	if err != nil {
		t.Fatalf("ReferenceKeys: %v", err)
	}
	if ft.gotMethod != wire.MethodModelReferenceKeys {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodModelReferenceKeys)
	}
	if len(res.Bodies) != 1 || len(res.Bodies[0].Faces) != 1 || res.Bodies[0].Faces[0].Key != "k0" {
		t.Errorf("decoded = %+v, want one body with face key k0", res)
	}
}
